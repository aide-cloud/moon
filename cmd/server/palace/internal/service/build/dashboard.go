package build

import (
	adminapi "github.com/aide-family/moon/api/admin"
	realtimeapi "github.com/aide-family/moon/api/admin/realtime"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-family/moon/pkg/palace/model/bizmodel"
)

type (
	// ChartItemBuilder bo 图表构造器
	ChartItemBuilder interface {
		ToBo() *bo.ChartItem
	}

	// BoAddDashboardParamsBuilder 添加仪表盘参数构造器
	BoAddDashboardParamsBuilder interface {
		ToBo() *bo.AddDashboardParams
	}

	// BoUpdateDashboardParamsBuilder 更新仪表盘参数构造器
	BoUpdateDashboardParamsBuilder interface {
		ToBo() *bo.UpdateDashboardParams
	}

	// BoDeleteDashboardParamsBuilder 删除仪表盘参数构造器
	BoDeleteDashboardParamsBuilder interface {
		ToBo() *bo.DeleteDashboardParams
	}

	// BoListDashboardParams 列表构造器
	BoListDashboardParams interface {
		ToBo() *bo.ListDashboardParams
	}

	// DoDashboardBuilder 仪表盘构造器
	DoDashboardBuilder interface {
		ToAPI() *adminapi.DashboardItem
		ToSelect() *adminapi.SelectItem
	}

	// DoDashboardListBuilder 仪表盘列表构造器
	DoDashboardListBuilder interface {
		ToAPIs() []*adminapi.DashboardItem
		ToSelects() []*adminapi.SelectItem
	}

	// DashboardModuleBuilder 仪表盘模块构造器
	DashboardModuleBuilder interface {
		WithBoChart(*adminapi.ChartItem) ChartItemBuilder
		WithAddDashboardParams(*realtimeapi.CreateDashboardRequest) BoAddDashboardParamsBuilder
		WithUpdateDashboardParams(*realtimeapi.UpdateDashboardRequest) BoUpdateDashboardParamsBuilder
		WithDeleteDashboardParams(*realtimeapi.DeleteDashboardRequest) BoDeleteDashboardParamsBuilder
		WithQueryDashboardListParams(*realtimeapi.ListDashboardRequest) BoListDashboardParams
		WithQueryDashboardSelectParams(*realtimeapi.ListDashboardSelectRequest) BoListDashboardParams
		WithDoDashboard(*bizmodel.Dashboard) DoDashboardBuilder
		WithDoDashboardList([]*bizmodel.Dashboard) DoDashboardListBuilder
	}
)
