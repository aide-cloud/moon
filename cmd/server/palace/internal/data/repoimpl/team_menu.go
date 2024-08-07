package repoimpl

import (
	"context"

	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/repository"
	"github.com/aide-family/moon/cmd/server/palace/internal/data"
	"github.com/aide-family/moon/pkg/palace/model/bizmodel"
	"github.com/aide-family/moon/pkg/palace/model/bizmodel/bizquery"
	"github.com/aide-family/moon/pkg/util/types"
)

// NewTeamMenuRepository 创建团队菜单仓库
func NewTeamMenuRepository(data *data.Data) repository.TeamMenu {
	return &teamMenuRepositoryImpl{
		data: data,
	}
}

type teamMenuRepositoryImpl struct {
	data *data.Data
}

func (l *teamMenuRepositoryImpl) GetTeamMenuList(ctx context.Context, params *bo.QueryTeamMenuListParams) ([]*bizmodel.SysTeamMenu, error) {
	bizDB, err := l.data.GetBizGormDB(params.TeamID)
	if !types.IsNil(err) {
		return nil, err
	}
	bizQuery := bizquery.Use(bizDB)
	return bizQuery.SysTeamMenu.WithContext(ctx).Find()
}
