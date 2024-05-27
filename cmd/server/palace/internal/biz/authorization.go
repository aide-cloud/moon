package biz

import (
	"context"

	"github.com/aide-cloud/moon/pkg/vobj"
	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"

	"github.com/aide-cloud/moon/api/merr"
	"github.com/aide-cloud/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-cloud/moon/cmd/server/palace/internal/biz/do/model"
	"github.com/aide-cloud/moon/cmd/server/palace/internal/biz/repo"
	"github.com/aide-cloud/moon/pkg/helper/middleware"
	"github.com/aide-cloud/moon/pkg/types"
)

type AuthorizationBiz struct {
	userRepo  repo.UserRepo
	teamRepo  repo.TeamRepo
	cacheRepo repo.CacheRepo
}

func NewAuthorizationBiz(
	userRepo repo.UserRepo,
	teamRepo repo.TeamRepo,
	cacheRepo repo.CacheRepo,
) *AuthorizationBiz {
	return &AuthorizationBiz{
		userRepo:  userRepo,
		teamRepo:  teamRepo,
		cacheRepo: cacheRepo,
	}
}

// CheckPermission 检查用户是否有该资源权限
func (b *AuthorizationBiz) CheckPermission(ctx context.Context, req *bo.CheckPermissionParams) error {
	if req.JwtClaims.GetTeamRole().IsSuperadmin() {
		return nil
	}
	// 检查用户是否被团队禁用
	teamDo, err := b.teamRepo.GetUserTeamByID(ctx, req.JwtClaims.GetUser(), req.JwtClaims.GetTeam())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return merr.ErrorModal("用户不在该团队中")
		}
		return merr.ErrorNotification("系统错误")
	}
	if !vobj.Status(teamDo.Status).IsEnable() {
		return merr.ErrorModal("用户被禁用")
	}

	// TODO 查询用户角色

	return nil
}

func (b *AuthorizationBiz) CheckToken(ctx context.Context, req *bo.CheckTokenParams) error {
	// 检查token是否过期
	if types.IsNil(req) || types.IsNil(req.JwtClaims) {
		return bo.UnLoginErr
	}
	if middleware.IsExpire(req.JwtClaims) {
		return bo.UnLoginErr
	}
	// 检查token是否被登出
	if req.JwtClaims.IsLogout(ctx, b.cacheRepo.Cacher()) {
		return bo.UnLoginErr
	}

	// 检查用户是否被系统禁用
	userDo, err := b.userRepo.GetByID(ctx, req.JwtClaims.GetUser())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return merr.ErrorModal("用户不存在")
		}
		return merr.ErrorNotification("系统错误")
	}
	if !vobj.Status(userDo.Status).IsEnable() {
		return merr.ErrorModal("用户被禁用")
	}
	return nil
}

// Login 登录
func (b *AuthorizationBiz) Login(ctx context.Context, req *bo.LoginParams) (*bo.LoginReply, error) {
	// 检查用户是否存在
	userDo, err := b.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 统一包装成密码错误
			return nil, bo.PasswordErr
		}
		return nil, bo.SystemErr
	}
	// 检查用户密码是否正确
	if err = checkPassword(userDo, req.EnPassword); err != nil {
		return nil, err
	}

	// 生成token
	base := &middleware.JwtBaseInfo{}

	base.SetUserInfo(func() (userId uint32, role vobj.Role, err error) {
		return userDo.ID, vobj.Role(userDo.Role), nil
	})
	base.SetTeamInfo(func() (teamId uint32, teamRole vobj.Role, err error) {
		if req.Team <= 0 {
			return
		}
		// 查询用户所属团队是否存在，存在着set temId
		memberItem, err := b.teamRepo.GetUserTeamByID(ctx, userDo.ID, req.Team)
		if err != nil {
			return 0, 0, err
		}
		return req.Team, vobj.Role(memberItem.Role), nil
	})

	jwtClaims := middleware.NewJwtClaims(base)
	return &bo.LoginReply{
		JwtClaims: jwtClaims,
		User:      userDo,
	}, nil
}

// RefreshToken 刷新token
func (b *AuthorizationBiz) RefreshToken(ctx context.Context, req *bo.RefreshTokenParams) (*bo.RefreshTokenReply, error) {
	// 检查token是否过期
	if types.IsNil(req) || types.IsNil(req.JwtClaims) {
		return nil, bo.UnLoginErr
	}
	// 检查用户是否存在
	userDo, err := b.userRepo.GetByID(ctx, req.JwtClaims.GetUser())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 统一包装成密码错误
			return nil, merr.ErrorModal("用户不存在")
		}
		return nil, bo.SystemErr
	}
	if !vobj.Status(userDo.Status).IsEnable() {
		return nil, merr.ErrorRedirect("用户被禁用").WithMetadata(map[string]string{
			"redirect": "/login",
		})
	}

	// 查询用户所属团队是否存在，存在着set temId
	teamMemberDo, err := b.teamRepo.GetUserTeamByID(ctx, userDo.ID, req.Team)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, merr.ErrorNotification("用户不在该团队中")
		}
		return nil, bo.SystemErr
	}

	if !vobj.Status(teamMemberDo.Status).IsEnable() {
		return nil, merr.ErrorNotification("用户被禁用")
	}

	// 查询用户所属团队角色是否存在，存在着set teamRoleId
	memberItem, err := b.teamRepo.GetUserTeamByID(ctx, userDo.ID, req.Team)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, merr.ErrorNotification("用户此权限已被移除")
		}
		return nil, bo.SystemErr
	}

	// 生成token
	base := &middleware.JwtBaseInfo{}
	base.SetUserInfo(func() (userId uint32, role vobj.Role, err error) {
		return userDo.ID, vobj.Role(userDo.Role), nil
	})
	base.SetTeamInfo(func() (teamId uint32, teamRole vobj.Role, err error) {
		return req.Team, vobj.Role(memberItem.Role), nil
	})

	jwtClaims := middleware.NewJwtClaims(base)
	return &bo.RefreshTokenReply{
		User:      userDo,
		JwtClaims: jwtClaims,
	}, nil
}

// Logout 登出
func (b *AuthorizationBiz) Logout(ctx context.Context, params *bo.LogoutParams) error {
	return params.JwtClaims.Logout(ctx, b.cacheRepo.Cacher())
}

// 检查用户密码是否正确
func checkPassword(user *model.SysUser, password string) error {
	_, err := types.DecryptPassword(password, user.Salt)
	if err == nil {
		return nil
	}
	return bo.PasswordErr
}