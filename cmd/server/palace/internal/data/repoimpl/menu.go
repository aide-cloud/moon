package repoimpl

import (
	"context"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/repository"
	"github.com/aide-family/moon/cmd/server/palace/internal/data"
	"github.com/aide-family/moon/pkg/palace/model"
	"github.com/aide-family/moon/pkg/util/types"
	"github.com/aide-family/moon/pkg/vobj"
)

func NewMenuRepository(data *data.Data) repository.Menu {
	return &menuRepositoryImpl{
		data: data,
	}
}

type menuRepositoryImpl struct {
	data *data.Data
}

func (m menuRepositoryImpl) Create(cxt context.Context, menu *bo.CreateMenuParams) (*model.SysMenu, error) {

	return nil, nil
}

func (m menuRepositoryImpl) UpdateById(cxt context.Context, menu *bo.UpdateMenuParams) error {

	return nil
}

func (m menuRepositoryImpl) DeleteById(cxt context.Context, id int64) error {

	return nil
}

func (m menuRepositoryImpl) GetById(cxt context.Context, id int64) (*model.SysMenu, error) {

	return nil, nil
}

func (m menuRepositoryImpl) ListAll(cxt context.Context) ([]*model.SysMenu, error) {

	return nil, nil
}

func (m menuRepositoryImpl) FindByPage(ctx context.Context, params *bo.QueryMenuListParams) ([]*model.SysMenu, error) {

	return nil, nil
}

func (m menuRepositoryImpl) UpdateStatusByIds(ctx context.Context, status vobj.Status, ids ...uint32) error {

	return nil
}

func (m menuRepositoryImpl) UpdateTypeByIds(ctx context.Context, status vobj.MenuType, ids ...uint32) error {

	return nil
}

func createMenuParamsToModel(menu *bo.CreateMenuParams) *model.SysMenu {
	if types.IsNil(menu) {
		return nil
	}
	return &model.SysMenu{}
}
