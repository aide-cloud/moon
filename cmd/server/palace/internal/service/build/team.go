package build

import (
	"context"

	"github.com/aide-family/moon/api"
	"github.com/aide-family/moon/api/admin"
	teamapi "github.com/aide-family/moon/api/admin/team"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-family/moon/cmd/server/palace/internal/data/runtimecache"
	"github.com/aide-family/moon/pkg/palace/model"
	"github.com/aide-family/moon/pkg/palace/model/bizmodel"
	"github.com/aide-family/moon/pkg/util/types"
	"github.com/aide-family/moon/pkg/vobj"
)

type (
	TeamModelBuilder interface {
		ToApi() *admin.Team
	}

	TeamRequestBuilder interface {
		ToCreateTeamBO() *bo.CreateTeamParams

		ToUpdateRoleBO() *bo.UpdateTeamParams

		ToQueryTeamListBO() *bo.QueryTeamListParams

		ToListTeamMemberBO() *bo.ListTeamMemberParams

		ToTeamListBO() *bo.QueryTeamListParams

		ToAddTeamMemberBO() *bo.AddTeamMemberParams

		WithLeaderId(leaderId uint32) TeamRequestBuilder
	}

	teamBuilder struct {
		// model
		SystemModel *model.SysTeam

		// request
		CreateRoleRequest     *teamapi.CreateTeamRequest
		UpdateTeamRequest     *teamapi.UpdateTeamRequest
		ListTeamRequest       *teamapi.ListTeamRequest
		ListTeamMemberRequest *teamapi.ListTeamMemberRequest
		AddTeamMemberRequest  *teamapi.AddTeamMemberRequest
		LeaderId              uint32

		// context
		ctx context.Context
	}

	TeamRoleBuilder interface {
		ToSelect() *admin.SelectItem
		ToApi() *admin.TeamRole
	}

	teamRoleBuilder struct {
		SysTeamRole *bizmodel.SysTeamRole
		ctx         context.Context
	}
)

func (b *teamBuilder) ToAddTeamMemberBO() *bo.AddTeamMemberParams {
	return &bo.AddTeamMemberParams{
		ID: b.AddTeamMemberRequest.GetId(),
		Members: types.SliceTo(b.AddTeamMemberRequest.GetMembers(), func(member *teamapi.AddTeamMemberRequest_MemberItem) *bo.AddTeamMemberItem {
			return &bo.AddTeamMemberItem{
				UserID:  member.GetUserId(),
				Role:    vobj.Role(member.GetRole()),
				RoleIds: member.GetRoles(),
			}
		}),
	}
}

func (b *teamBuilder) ToQueryTeamListBO() *bo.QueryTeamListParams {
	//TODO implement me
	panic("implement me")
}

// ToApi 转换为API层数据
func (b *teamBuilder) ToApi() *admin.Team {
	if types.IsNil(b) || types.IsNil(b.SystemModel) {
		return nil
	}
	cache := runtimecache.GetRuntimeCache()
	return &admin.Team{
		Id:        b.SystemModel.ID,
		Name:      b.SystemModel.Name,
		Status:    api.Status(b.SystemModel.Status),
		Remark:    b.SystemModel.Remark,
		CreatedAt: b.SystemModel.CreatedAt.String(),
		UpdatedAt: b.SystemModel.UpdatedAt.String(),
		Leader:    NewBuilder().WithApiUserBo(cache.GetUser(b.ctx, b.SystemModel.LeaderID)).ToApi(),
		Creator:   NewBuilder().WithApiUserBo(cache.GetUser(b.ctx, b.SystemModel.CreatorID)).ToApi(),
		Logo:      b.SystemModel.Logo,
		// 从全局中取
		Admin: types.SliceTo(cache.GetTeamAdminList(b.ctx, b.SystemModel.ID), func(item *bizmodel.SysTeamMember) *admin.TeamMember {
			return NewBuilder().WithApiTeamMember(item).ToApi(b.ctx)
		}),
	}
}

func (b *teamBuilder) ToCreateTeamBO() *bo.CreateTeamParams {
	if types.IsNil(b) || types.IsNil(b.CreateRoleRequest) {
		return nil
	}
	return &bo.CreateTeamParams{
		Name:     b.CreateRoleRequest.GetName(),
		Remark:   b.CreateRoleRequest.GetRemark(),
		Logo:     b.CreateRoleRequest.GetLogo(),
		Status:   vobj.Status(b.CreateRoleRequest.GetStatus()),
		LeaderID: b.LeaderId,
		Admins:   b.CreateRoleRequest.GetAdminIds(),
	}
}

func (b *teamBuilder) ToUpdateRoleBO() *bo.UpdateTeamParams {
	if types.IsNil(b) || types.IsNil(b.UpdateTeamRequest) {
		return nil
	}
	return &bo.UpdateTeamParams{
		ID:     b.UpdateTeamRequest.GetId(),
		Name:   b.UpdateTeamRequest.GetName(),
		Remark: b.UpdateTeamRequest.GetRemark(),
		Logo:   b.UpdateTeamRequest.GetLogo(),
		Status: vobj.Status(b.UpdateTeamRequest.GetStatus()),
	}
}

func (b *teamBuilder) ToQueryTeamList() *bo.QueryTeamListParams {
	if types.IsNil(b) || types.IsNil(b.ListTeamRequest) {
		return nil
	}
	return &bo.QueryTeamListParams{
		Page:      types.NewPagination(b.ListTeamRequest.GetPagination()),
		Keyword:   b.ListTeamRequest.GetKeyword(),
		Status:    vobj.Status(b.ListTeamRequest.GetStatus()),
		CreatorID: b.ListTeamRequest.GetCreatorId(),
		LeaderID:  b.ListTeamRequest.GetLeaderId(),
	}
}

func (b *teamBuilder) ToListTeamMemberBO() *bo.ListTeamMemberParams {
	if types.IsNil(b) || types.IsNil(b.ListTeamMemberRequest) {
		return nil
	}
	return &bo.ListTeamMemberParams{
		Page:    types.NewPagination(b.ListTeamMemberRequest.GetPagination()),
		ID:      b.ListTeamMemberRequest.GetId(),
		Keyword: b.ListTeamMemberRequest.GetKeyword(),
		Role:    vobj.Role(b.ListTeamMemberRequest.GetRole()),
		Gender:  vobj.Gender(b.ListTeamMemberRequest.GetGender()),
		Status:  vobj.Status(b.ListTeamMemberRequest.GetStatus()),
	}
}

func (b *teamBuilder) ToTeamListBO() *bo.QueryTeamListParams {
	if types.IsNil(b) || types.IsNil(b.ListTeamRequest) {
		return nil
	}
	return &bo.QueryTeamListParams{
		Page:      types.NewPagination(b.ListTeamRequest.GetPagination()),
		Keyword:   b.ListTeamRequest.GetKeyword(),
		Status:    vobj.Status(b.ListTeamRequest.GetStatus()),
		CreatorID: b.ListTeamRequest.GetCreatorId(),
		LeaderID:  b.ListTeamRequest.GetLeaderId(),
	}
}

func (b *teamBuilder) WithLeaderId(leaderId uint32) TeamRequestBuilder {
	b.LeaderId = leaderId
	return b
}

func (b *teamRoleBuilder) ToApi() *admin.TeamRole {
	if types.IsNil(b) || types.IsNil(b.SysTeamRole) {
		return nil
	}
	return &admin.TeamRole{
		Id:        b.SysTeamRole.ID,
		Name:      b.SysTeamRole.Name,
		Remark:    b.SysTeamRole.Remark,
		CreatedAt: b.SysTeamRole.CreatedAt.String(),
		UpdatedAt: b.SysTeamRole.UpdatedAt.String(),
		Status:    api.Status(b.SysTeamRole.Status),
		Resources: types.SliceTo(b.SysTeamRole.Apis, func(item *bizmodel.SysTeamAPI) *admin.ResourceItem {
			return NewTeamResourceBuilder(item).ToApi()
		}),
	}
}

// ToSelect 转换为Select数据
func (b *teamRoleBuilder) ToSelect() *admin.SelectItem {
	if types.IsNil(b) || types.IsNil(b.SysTeamRole) {
		return nil
	}
	return &admin.SelectItem{
		Value:    b.SysTeamRole.ID,
		Label:    b.SysTeamRole.Name,
		Disabled: b.SysTeamRole.DeletedAt > 0 || !vobj.Status(b.SysTeamRole.Status).IsEnable(),
	}
}
