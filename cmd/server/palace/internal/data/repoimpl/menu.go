package repoimpl

import (
	"context"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/repository"
	"github.com/aide-family/moon/cmd/server/palace/internal/data"
	"github.com/aide-family/moon/pkg/palace/model"
	"github.com/aide-family/moon/pkg/palace/model/query"
	"github.com/aide-family/moon/pkg/util/types"
	"github.com/aide-family/moon/pkg/vobj"
	"gorm.io/gen"
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
	menuModel := createMenuParamsToModel(menu)
	menuModel.WithContext(cxt)
	queryWrapper := query.Use(m.data.GetMainDB(cxt)).WithContext(cxt).SysMenu
	if err := queryWrapper.Create(menuModel); !types.IsNil(err) {
		return nil, err
	}
	return menuModel, nil
}

func (m menuRepositoryImpl) BatchCreate(ctx context.Context, menus []*bo.CreateMenuParams) error {
	menuModels := types.SliceToWithFilter(menus, func(item *bo.CreateMenuParams) (*model.SysMenu, bool) {
		if types.IsNil(item) || types.TextIsNull(item.Name) {
			return nil, false
		}
		return createMenuParamsToModel(item), true
	})
	return query.Use(m.data.GetMainDB(ctx)).WithContext(ctx).SysMenu.CreateInBatches(menuModels, 10)
}

func (m menuRepositoryImpl) UpdateById(cxt context.Context, menu *bo.UpdateMenuParams) error {
	updateParam := menu.UpdateParam
	_, err := query.Use(m.data.GetMainDB(cxt)).WithContext(cxt).SysMenu.Where(query.SysMenu.ID.Eq(menu.ID)).UpdateSimple(
		query.SysMenu.Name.Value(updateParam.Name),
		query.SysMenu.Component.Value(updateParam.Component),
		query.SysMenu.Path.Value(updateParam.Path),
		query.SysMenu.Icon.Value(updateParam.Icon),
		query.SysMenu.Permission.Value(updateParam.Permission),
		query.SysMenu.Level.Value(updateParam.Level),
	)
	return err
}

func (m menuRepositoryImpl) DeleteById(cxt context.Context, id uint32) error {
	_, err := query.Use(m.data.GetMainDB(cxt)).WithContext(cxt).SysDict.Where(query.SysMenu.ID.Eq(id)).Delete()
	return err
}

func (m menuRepositoryImpl) GetById(cxt context.Context, id uint32) (*model.SysMenu, error) {
	return query.Use(m.data.GetMainDB(cxt)).SysMenu.WithContext(cxt).Where(query.SysMenu.ID.Eq(id)).First()
}

func (m menuRepositoryImpl) ListAll(cxt context.Context) ([]*model.SysMenu, error) {
	menus, err := query.Use(m.data.GetMainDB(cxt)).SysMenu.WithContext(cxt).Find()
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (m menuRepositoryImpl) FindByPage(ctx context.Context, params *bo.QueryMenuListParams) ([]*model.SysMenu, error) {
	queryWrapper := query.Use(m.data.GetMainDB(ctx)).SysMenu.WithContext(ctx)
	var wheres []gen.Condition
	if !params.Status.IsUnknown() {
		wheres = append(wheres, query.SysMenu.Status.Eq(params.Status.GetValue()))
	}

	if !params.MenuType.IsUnknown() {
		wheres = append(wheres, query.SysMenu.Type.Eq(params.MenuType.GetValue()))
	}

	if !types.TextIsNull(params.Keyword) {
		queryWrapper = queryWrapper.Or(
			query.SysMenu.Name.Like(params.Keyword),
			query.SysMenu.Path.Like(params.Keyword),
			query.SysMenu.EnName.Like(params.Keyword),
		)
	}
	queryWrapper = queryWrapper.Where(wheres...)
	if err := types.WithPageQuery[query.ISysMenuDo](queryWrapper, params.Page); err != nil {
		return nil, err
	}
	return queryWrapper.Order(query.SysMenu.ID.Desc()).Find()
}

func (m menuRepositoryImpl) UpdateStatusByIds(ctx context.Context, status vobj.Status, ids ...uint32) error {
	_, err := query.Use(m.data.GetMainDB(ctx)).WithContext(ctx).SysMenu.Where(query.SysMenu.ID.In(ids...)).Update(query.SysMenu.Status, status)
	return err
}

func (m menuRepositoryImpl) UpdateTypeByIds(ctx context.Context, menuType vobj.MenuType, ids ...uint32) error {
	_, err := query.Use(m.data.GetMainDB(ctx)).WithContext(ctx).SysMenu.Where(query.SysMenu.ID.In(ids...)).Update(query.SysMenu.Type, menuType)
	return err
}

func createMenuParamsToModel(menu *bo.CreateMenuParams) *model.SysMenu {
	if types.IsNil(menu) {
		return nil
	}
	return &model.SysMenu{
		Name:       menu.Name,
		Path:       menu.Path,
		Icon:       menu.Icon,
		Type:       menu.Type,
		Sort:       menu.Sort,
		ParentID:   menu.ParentId,
		Status:     menu.Status,
		Permission: menu.Permission,
		EnName:     menu.EnName,
	}
}
