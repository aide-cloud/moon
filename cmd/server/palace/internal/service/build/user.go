package build

import (
	"context"

	"github.com/aide-family/moon/api"
	"github.com/aide-family/moon/api/admin"
	userapi "github.com/aide-family/moon/api/admin/user"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-family/moon/cmd/server/palace/internal/data/runtimecache"
	"github.com/aide-family/moon/pkg/palace/model"
	"github.com/aide-family/moon/pkg/palace/model/bizmodel"
	"github.com/aide-family/moon/pkg/util/types"
	"github.com/aide-family/moon/pkg/vobj"
)

type (
	// UserModelBuilder 用户模型转换
	UserModelBuilder interface {
		ToAPI() *admin.User

		GetUsername() string
	}

	// UserRequestBuilder 用户请求转换
	UserRequestBuilder interface {
		ToCreateUserBO(uint32, types.Password) *bo.CreateUserParams

		ToUpdateUserBO() *bo.UpdateUserParams
	}

	userBuilder struct {
		// model
		*model.SysUser

		// request
		CreateUserRequest *userapi.CreateUserRequest
		UpdateUserRequest *userapi.UpdateUserRequest

		// context
		ctx context.Context
	}

	// TeamMemberBuilder 团队成员构造器
	TeamMemberBuilder interface {
		ToAPI(ctx context.Context) *admin.TeamMember
	}

	teamMemberBuilder struct {
		*bizmodel.SysTeamMember
		ctx context.Context
	}
)

func (b *userBuilder) ToCreateUserBO(userID uint32, pass types.Password) *bo.CreateUserParams {
	if types.IsNil(b) || types.IsNil(b.CreateUserRequest) {
		return nil
	}
	return &bo.CreateUserParams{
		Name:      b.CreateUserRequest.GetName(),
		Password:  pass,
		Email:     b.CreateUserRequest.GetEmail(),
		Phone:     b.CreateUserRequest.GetPhone(),
		Nickname:  b.CreateUserRequest.GetNickname(),
		Remark:    b.CreateUserRequest.GetRemark(),
		Avatar:    b.CreateUserRequest.GetAvatar(),
		CreatorID: userID,
		Status:    vobj.Status(b.CreateUserRequest.GetStatus()),
		Gender:    vobj.Gender(b.CreateUserRequest.GetGender()),
		Role:      vobj.Role(b.CreateUserRequest.GetRole()),
	}
}

func (b *userBuilder) ToUpdateUserBO() *bo.UpdateUserParams {
	if types.IsNil(b) || types.IsNil(b.UpdateUserRequest) {
		return nil
	}
	createParams := bo.CreateUserParams{
		Name:     b.UpdateUserRequest.GetData().GetName(),
		Email:    b.UpdateUserRequest.GetData().GetEmail(),
		Phone:    b.UpdateUserRequest.GetData().GetPhone(),
		Nickname: b.UpdateUserRequest.GetData().GetNickname(),
		Remark:   b.UpdateUserRequest.GetData().GetRemark(),
		Avatar:   b.UpdateUserRequest.GetData().GetAvatar(),
		Status:   vobj.Status(b.UpdateUserRequest.GetData().GetStatus()),
		Gender:   vobj.Gender(b.UpdateUserRequest.GetData().GetGender()),
		Role:     vobj.Role(b.UpdateUserRequest.GetData().GetRole()),
	}
	return &bo.UpdateUserParams{
		ID:               b.UpdateUserRequest.GetId(),
		CreateUserParams: createParams,
	}
}

// ToAPI 转换成api
func (b *userBuilder) ToAPI() *admin.User {
	if types.IsNil(b) || types.IsNil(b.SysUser) {
		return nil
	}
	return &admin.User{
		Id:        b.ID,
		Name:      b.Username,
		Nickname:  b.Nickname,
		Email:     b.Email,
		Phone:     b.Phone,
		Status:    api.Status(b.Status),
		Gender:    api.Gender(b.Gender),
		Role:      api.Role(b.Role),
		Avatar:    b.Avatar,
		Remark:    b.Remark,
		CreatedAt: b.CreatedAt.String(),
		UpdatedAt: b.UpdatedAt.String(),
	}
}

func (b *userBuilder) GetUsername() string {
	if types.IsNil(b) || types.IsNil(b.SysUser) {
		return ""
	}
	return b.Username
}

func (b *teamMemberBuilder) ToAPI(ctx context.Context) *admin.TeamMember {
	if types.IsNil(b) || types.IsNil(b.SysTeamMember) {
		return nil
	}
	cache := runtimecache.GetRuntimeCache()
	return &admin.TeamMember{
		UserId:    b.UserID,
		Id:        b.ID,
		Role:      api.Role(b.Role),
		Status:    api.Status(b.Status),
		CreatedAt: b.CreatedAt.String(),
		UpdatedAt: b.UpdatedAt.String(),
		User:      NewBuilder().WithAPIUserBo(cache.GetUser(ctx, b.UserID)).ToAPI(),
	}
}
