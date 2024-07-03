package bo

import (
	"github.com/aide-family/moon/pkg/util/types"
	"github.com/aide-family/moon/pkg/vobj"
)

type (
	CreateMenuParams struct {
	}

	UpdateMenuParams struct {
	}

	QueryMenuListParams struct {
		Keyword  string           `json:"keyword"`
		Page     types.Pagination `json:"page"`
		Status   vobj.Status      `json:"status"`
		MenuType vobj.MenuType    `json:"menu_type"`
	}
)
