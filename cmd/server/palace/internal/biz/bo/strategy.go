package bo

import (
	"github.com/aide-family/moon/pkg/util/types"
	"github.com/aide-family/moon/pkg/vobj"
)

type (
	CreateStrategyParams struct {
		// 策略组ID
		GroupId uint32 `json:"group_id"`
		// 策略模板id
		TemplateId uint32 `json:"template_id"`
		// 备注
		Remark string `json:"remark"`
		// 状态
		Status vobj.Status `json:"status"`
		// 采样率
		Step uint32 `json:"step"`
		// 数据源id
		DatasourceIds []uint32 `json:"datasource_ids"`
		// 模板来源
		SourceType vobj.TemplateSourceType `json:"source_type"`
		// 策略名称
		Name   string `json:"name"`
		TeamID uint32 `json:"teamID"`
		// 条件
		Condition string `json:"condition"`
		// 阈值
		Threshold float64 `json:"threshold"`
	}

	UpdateStrategyParams struct {
		ID          uint32 `json:"id"`
		UpdateParam CreateStrategyParams
		TeamID      uint32 `json:"teamID"`
	}

	QueryStrategyListParams struct {
		Keyword    string `json:"keyword"`
		Page       types.Pagination
		Alert      string
		Status     vobj.Status
		SourceType vobj.TemplateSourceType
		TeamID     uint32 `json:"teamID"`
	}

	GetStrategyDetailParams struct {
		ID     uint32 `json:"id"`
		TeamID uint32 `json:"teamID"`
	}

	DelStrategyParams struct {
		ID     uint32 `json:"id"`
		TeamID uint32 `json:"teamID"`
	}

	UpdateStrategyStatusParams struct {
		Ids    []uint32 `json:"ids"`
		TeamID uint32   `json:"teamID"`
		Status vobj.Status
	}
)
