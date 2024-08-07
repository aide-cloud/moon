package build

import (
	"context"

	"github.com/aide-family/moon/api/admin"
	datasourceapi "github.com/aide-family/moon/api/admin/datasource"
	dictapi "github.com/aide-family/moon/api/admin/dict"
	menuapi "github.com/aide-family/moon/api/admin/menu"
	strategyapi "github.com/aide-family/moon/api/admin/strategy"
	teamapi "github.com/aide-family/moon/api/admin/team"
	userapi "github.com/aide-family/moon/api/admin/user"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-family/moon/pkg/palace/imodel"
	"github.com/aide-family/moon/pkg/palace/model"
	"github.com/aide-family/moon/pkg/palace/model/bizmodel"
)

// NewBuilder 创建公共构造器
func NewBuilder() Builder {
	return &builder{
		ctx: context.TODO(),
	}
}

type (
	builder struct {
		ctx context.Context
	}

	// Builder 构造器公共接口
	Builder interface {
		WithContext(ctx context.Context) Builder

		// TODO 注册新的数据转换方法写在这里

		WithDoDatasource(*bizmodel.Datasource) DatasourceModelBuilder
		WithCreateDatasourceBo(*datasourceapi.CreateDatasourceRequest) DatasourceRequestBuilder
		WithListDatasourceBo(*datasourceapi.ListDatasourceRequest) DatasourceRequestBuilder
		WithUpdateDatasourceBo(*datasourceapi.UpdateDatasourceRequest) DatasourceRequestBuilder
		WithBoDatasourceQueryData(*bo.DatasourceQueryData) DatasourceQueryDataBuilder

		WithAPITemplateStrategy(*model.StrategyTemplate) TemplateModelBuilder
		WithCreateBoTemplateStrategy(*strategyapi.CreateTemplateStrategyRequest) TemplateRequestBuilder
		WithUpdateBoTemplateStrategy(*strategyapi.UpdateTemplateStrategyRequest) TemplateRequestBuilder
		WithAPITemplateStrategyLevel(*model.StrategyLevelTemplate) TemplateLevelBuilder

		WithAPIStrategy(*bizmodel.Strategy) StrategyModelBuilder
		WithCreateBoStrategy(*strategyapi.CreateStrategyRequest) StrategyRequestBuilder
		WithUpdateBoStrategy(*strategyapi.UpdateStrategyRequest) StrategyRequestBuilder
		WithAPIStrategyLevel(*bizmodel.StrategyLevel) StrategyLevelModelBuilder

		WithAPIStrategyGroup(strategyGroup *bizmodel.StrategyGroup) StrategyGroupModelBuilder
		WithStrategyGroupList(strategyGroup []*bizmodel.StrategyGroup, strategyCountModel []*bo.StrategyCountModel, strategyEnableCountModel []*bo.StrategyCountModel) StrategyGroupModelBuilder

		WithCreateBoStrategyGroup(*strategyapi.CreateStrategyGroupRequest) StrategyGroupRequestBuilder
		WithUpdateBoStrategyGroup(*strategyapi.UpdateStrategyGroupRequest) StrategyGroupRequestBuilder
		WithListStrategyGroup(*strategyapi.ListStrategyGroupRequest) StrategyGroupRequestBuilder

		WithCreateBoDict(*dictapi.CreateDictRequest) DictRequestBuilder
		WithUpdateBoDict(*dictapi.UpdateDictRequest) DictRequestBuilder
		WithAPIDictSelect(*model.SysDict) DictModelBuilder
		WithDict(imodel.IDict) DictModelBuilder

		WithCreateMenuBo(*menuapi.CreateMenuRequest) MenuRequestBuilder
		WithUpdateMenuBo(*menuapi.UpdateMenuRequest) MenuRequestBuilder
		WithAPIMenu(*model.SysMenu) MenuModelBuilder
		WithBatchCreateMenuBo(*menuapi.BatchCreateMenuRequest) MenuRequestBuilder
		WithAPIMenuTree([]*admin.Menu, uint32) MenuTreeBuilder

		WithAPITeam(*model.SysTeam) TeamModelBuilder
		WithSelectTeamRole(*bizmodel.SysTeamRole) TeamRoleBuilder
		WithAPITeamRole(*bizmodel.SysTeamRole) TeamRoleBuilder
		WithCreateTeamBo(*teamapi.CreateTeamRequest) TeamRequestBuilder
		WithUpdateTeamBo(*teamapi.UpdateTeamRequest) TeamRequestBuilder
		WithListTeamBo(*teamapi.ListTeamRequest) TeamRequestBuilder

		WithListTeamTeamMemberBo(*teamapi.ListTeamMemberRequest) TeamRequestBuilder
		WithAddTeamMemberBo(*teamapi.AddTeamMemberRequest) TeamRequestBuilder
		WithAPITeamMember(*bizmodel.SysTeamMember) TeamMemberBuilder

		WithAPIUserBo(*model.SysUser) UserModelBuilder
		WithCreateUserBo(*userapi.CreateUserRequest) UserRequestBuilder
		WithUpdateUserBo(*userapi.UpdateUserRequest) UserRequestBuilder
		WithAPIDatasourceMetric(*bizmodel.DatasourceMetric) DatasourceMetricModelBuilder

		WithAPIDatasourceMetricLabel(metric *bizmodel.MetricLabel) DatasourceMetricLabelModelBuilder
		WithAPIDatasourceMetricLabelValue(metric *bizmodel.MetricLabelValue) DatasourceMetricLabelValueBuilder

		StrategyGroupModuleBuilder() StrategyGroupModuleBuilder

		RealTimeAlarmModule() RealtimeAlarmModuleBuilder
		DashboardModule() DashboardModuleBuilder
		AlarmPageModule() AlarmPageModuleBuilder
	}
)

func (b *builder) AlarmPageModule() AlarmPageModuleBuilder {
	return newAlarmPageModuleBuilder(b.ctx)
}

func (b *builder) StrategyGroupModuleBuilder() StrategyGroupModuleBuilder {
	return NewStrategyGroupModuleBuilder(b.ctx)
}

func (b *builder) WithStrategyGroupList(strategyGroup []*bizmodel.StrategyGroup, strategyCountModel []*bo.StrategyCountModel, strategyEnableCountModel []*bo.StrategyCountModel) StrategyGroupModelBuilder {
	return &strategyGroupBuilder{
		StrategyGroups: strategyGroup,
	}
}

func (b *builder) DashboardModule() DashboardModuleBuilder {
	return NewDashboardModuleBuilder(b.ctx)
}

func (b *builder) RealTimeAlarmModule() RealtimeAlarmModuleBuilder {
	return newRealtimeAlarmModuleBuilder(b.ctx)
}

func (b *builder) WithDict(dict imodel.IDict) DictModelBuilder {
	return &dictBuilder{
		SysDict: dict,
		ctx:     b.ctx,
	}
}

func (b *builder) WithBoDatasourceQueryData(d *bo.DatasourceQueryData) DatasourceQueryDataBuilder {
	return &datasourceQueryDataBuilder{
		DatasourceQueryData: d,
		ctx:                 b.ctx,
	}
}

func (b *builder) WithDoDatasource(d *bizmodel.Datasource) DatasourceModelBuilder {
	return &datasourceBuilder{
		Datasource: d,
		ctx:        b.ctx,
	}
}
func (b *builder) WithAPITemplateStrategy(template *model.StrategyTemplate) TemplateModelBuilder {
	return &templateStrategyBuilder{
		StrategyTemplate: template,
		ctx:              b.ctx,
	}
}

func (b *builder) WithCreateBoTemplateStrategy(template *strategyapi.CreateTemplateStrategyRequest) TemplateRequestBuilder {
	return &templateStrategyBuilder{
		CreateStrategy: template,
		ctx:            b.ctx,
	}
}

func (b *builder) WithUpdateBoTemplateStrategy(template *strategyapi.UpdateTemplateStrategyRequest) TemplateRequestBuilder {
	return &templateStrategyBuilder{
		UpdateStrategy: template,
		ctx:            b.ctx,
	}
}

func (b *builder) WithAPITemplateStrategyLevel(template *model.StrategyLevelTemplate) TemplateLevelBuilder {
	return &templateStrategyLevelBuilder{
		StrategyLevelTemplate: template,
		ctx:                   b.ctx,
	}
}

func (b *builder) WithAPIStrategy(strategy *bizmodel.Strategy) StrategyModelBuilder {
	return &strategyBuilder{
		Strategy: strategy,
		ctx:      b.ctx,
	}
}

func (b *builder) WithCreateBoStrategy(strategy *strategyapi.CreateStrategyRequest) StrategyRequestBuilder {
	return &strategyBuilder{
		CreateStrategy: strategy,
		ctx:            b.ctx,
	}
}

func (b *builder) WithUpdateBoStrategy(strategy *strategyapi.UpdateStrategyRequest) StrategyRequestBuilder {
	return &strategyBuilder{
		UpdateStrategy: strategy,
		ctx:            b.ctx,
	}
}

func (b *builder) WithCreateBoDict(dict *dictapi.CreateDictRequest) DictRequestBuilder {
	return &dictBuilder{
		CreateDictRequest: dict,
		ctx:               b.ctx,
	}
}

func (b *builder) WithUpdateBoDict(dict *dictapi.UpdateDictRequest) DictRequestBuilder {
	return &dictBuilder{
		UpdateDictRequest: dict,
		ctx:               b.ctx,
	}
}

func (b *builder) WithAPIDictSelect(dict *model.SysDict) DictModelBuilder {
	return &dictBuilder{
		SysDict: dict,
		ctx:     b.ctx,
	}
}

func (b *builder) WithCreateMenuBo(menu *menuapi.CreateMenuRequest) MenuRequestBuilder {
	return &menuBuilder{
		CreateMenuRequest: menu,
		ctx:               b.ctx,
	}
}

func (b *builder) WithUpdateMenuBo(menu *menuapi.UpdateMenuRequest) MenuRequestBuilder {
	return &menuBuilder{
		UpdateMenuRequest: menu,
		ctx:               b.ctx,
	}
}

func (b *builder) WithAPIMenu(menu *model.SysMenu) MenuModelBuilder {
	return &menuBuilder{
		Menu: menu,
		ctx:  b.ctx,
	}
}

func (b *builder) WithBatchCreateMenuBo(menu *menuapi.BatchCreateMenuRequest) MenuRequestBuilder {
	return &menuBuilder{
		BatchCreateMenuRequest: menu,
		ctx:                    b.ctx,
	}
}

func (b *builder) WithAPIMenuTree(menuList []*admin.Menu, parentID uint32) MenuTreeBuilder {
	menuMap := make(map[uint32][]*admin.Menu)
	// 按照父级ID分组
	for _, menu := range menuList {
		if _, ok := menuMap[menu.GetParentId()]; !ok {
			menuMap[menu.GetParentId()] = make([]*admin.Menu, 0)
		}
		menuMap[menu.GetParentId()] = append(menuMap[menu.GetParentId()], menu)
	}
	return &menuTreeBuilder{
		MenuMap:  menuMap,
		ParentID: parentID,
		ctx:      b.ctx,
	}
}

func (b *builder) WithAPITeam(team *model.SysTeam) TeamModelBuilder {
	return &teamBuilder{
		SystemModel: team,
		ctx:         b.ctx,
	}
}

func (b *builder) WithSelectTeamRole(team *bizmodel.SysTeamRole) TeamRoleBuilder {
	return &teamRoleBuilder{
		SysTeamRole: team,
		ctx:         b.ctx,
	}
}

func (b *builder) WithAPITeamRole(team *bizmodel.SysTeamRole) TeamRoleBuilder {
	return &teamRoleBuilder{
		SysTeamRole: team,
		ctx:         b.ctx,
	}
}

func (b *builder) WithCreateTeamBo(req *teamapi.CreateTeamRequest) TeamRequestBuilder {
	return &teamBuilder{
		CreateRoleRequest: req,
		ctx:               b.ctx,
	}
}

func (b *builder) WithUpdateTeamBo(req *teamapi.UpdateTeamRequest) TeamRequestBuilder {
	return &teamBuilder{
		UpdateTeamRequest: req,
		ctx:               b.ctx,
	}
}

func (b *builder) WithListTeamBo(req *teamapi.ListTeamRequest) TeamRequestBuilder {
	return &teamBuilder{
		ListTeamRequest: req,
		ctx:             b.ctx,
	}
}

func (b *builder) WithListTeamTeamMemberBo(req *teamapi.ListTeamMemberRequest) TeamRequestBuilder {
	return &teamBuilder{
		ListTeamMemberRequest: req,
		ctx:                   b.ctx,
	}
}

func (b *builder) WithAddTeamMemberBo(req *teamapi.AddTeamMemberRequest) TeamRequestBuilder {
	return &teamBuilder{
		AddTeamMemberRequest: req,
		ctx:                  b.ctx,
	}
}

func (b *builder) WithAPITeamMember(teamMember *bizmodel.SysTeamMember) TeamMemberBuilder {
	return &teamMemberBuilder{
		SysTeamMember: teamMember,
		ctx:           b.ctx,
	}
}

func (b *builder) WithAPIUserBo(user *model.SysUser) UserModelBuilder {
	return &userBuilder{
		SysUser: user,
		ctx:     b.ctx,
	}
}

func (b *builder) WithCreateUserBo(req *userapi.CreateUserRequest) UserRequestBuilder {
	return &userBuilder{
		CreateUserRequest: req,
		ctx:               b.ctx,
	}
}

func (b *builder) WithUpdateUserBo(req *userapi.UpdateUserRequest) UserRequestBuilder {
	return &userBuilder{
		UpdateUserRequest: req,
		ctx:               b.ctx,
	}
}

func (b *builder) WithCreateDatasourceBo(req *datasourceapi.CreateDatasourceRequest) DatasourceRequestBuilder {
	return &datasourceBuilder{
		CreateDatasourceRequest: req, ctx: b.ctx}
}

func (b *builder) WithUpdateDatasourceBo(req *datasourceapi.UpdateDatasourceRequest) DatasourceRequestBuilder {
	return &datasourceBuilder{
		UpdateDatasourceRequest: req,
		ctx:                     b.ctx,
	}
}

func (b *builder) WithListDatasourceBo(req *datasourceapi.ListDatasourceRequest) DatasourceRequestBuilder {
	return &datasourceBuilder{
		ListDatasourceRequest: req,
		ctx:                   b.ctx,
	}
}

func (b *builder) WithAPIDatasourceMetric(metric *bizmodel.DatasourceMetric) DatasourceMetricModelBuilder {
	return &datasourceMetricModelBuilder{
		DatasourceMetric: metric,
		ctx:              b.ctx,
	}
}

func (b *builder) WithAPIDatasourceMetricLabel(metric *bizmodel.MetricLabel) DatasourceMetricLabelModelBuilder {
	return &datasourceMetricLabelModelBuilder{
		MetricLabel: metric,
		ctx:         b.ctx,
	}
}

func (b *builder) WithAPIDatasourceMetricLabelValue(metric *bizmodel.MetricLabelValue) DatasourceMetricLabelValueBuilder {

	return &datasourceMetricLabelValueBuilder{
		MetricLabelValue: metric,
		ctx:              b.ctx,
	}
}

func (b *builder) WithAPIStrategyLevel(strategy *bizmodel.StrategyLevel) StrategyLevelModelBuilder {
	return &strategyLevelBuilder{
		StrategyLevel: strategy,
		ctx:           b.ctx,
	}

}

func (b *builder) WithAPIStrategyGroup(strategy *bizmodel.StrategyGroup) StrategyGroupModelBuilder {
	return &strategyGroupBuilder{
		StrategyGroup: strategy,
		ctx:           b.ctx,
	}
}

func (b *builder) WithCreateBoStrategyGroup(strategy *strategyapi.CreateStrategyGroupRequest) StrategyGroupRequestBuilder {
	return &strategyGroupBuilder{
		CreateStrategyGroupRequest: strategy,
		ctx:                        b.ctx,
	}
}

func (b *builder) WithUpdateBoStrategyGroup(strategy *strategyapi.UpdateStrategyGroupRequest) StrategyGroupRequestBuilder {
	return &strategyGroupBuilder{
		UpdateStrategyGroupRequest: strategy,
		ctx:                        b.ctx,
	}
}
func (b *builder) WithListStrategyGroup(strategy *strategyapi.ListStrategyGroupRequest) StrategyGroupRequestBuilder {
	return &strategyGroupBuilder{
		ListStrategyGroupRequest: strategy,
		ctx:                      b.ctx,
	}
}

func (b *builder) WithContext(ctx context.Context) Builder {
	b.ctx = ctx
	return b
}
