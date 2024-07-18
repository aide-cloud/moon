package build

import (
	"context"

	"github.com/aide-family/moon/api/admin"
	dictapi "github.com/aide-family/moon/api/admin/dict"
	menuapi "github.com/aide-family/moon/api/admin/menu"
	strategyapi "github.com/aide-family/moon/api/admin/strategy"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-family/moon/pkg/palace/model"
	"github.com/aide-family/moon/pkg/palace/model/bizmodel"
)

func NewBuilder() *builder {
	return &builder{}
}

type (
	builder struct {
		ctx context.Context
	}

	Builder interface {
		WithContext(ctx context.Context) Builder

		// TODO 注册新的数据转换方法写在这里

		WithDoDatasource(d *bizmodel.Datasource) DatasourceBuilder
		WithBoDatasourceQueryData(d *bo.DatasourceQueryData) DatasourceQueryDataBuilder

		WithApiTemplateStrategy(template *model.StrategyTemplate) TemplateBuilder
		WithCreateBoTemplateStrategy(template *strategyapi.CreateTemplateStrategyRequest) TemplateBuilder
		WithUpdateBoTemplateStrategy(template *strategyapi.UpdateTemplateStrategyRequest) TemplateBuilder

		WithApiTemplateStrategyLevel(*model.StrategyLevelTemplate) TemplateLevelBuilder

		WithApiStrategy(strategy *bizmodel.Strategy) StrategyBuilder

		WithCreateBoStrategy(strategy *strategyapi.CreateStrategyRequest) StrategyBuilder

		WithUpdateBoStrategy(strategy *strategyapi.UpdateStrategyRequest) StrategyBuilder

		WithCreateBoDict(dict *dictapi.CreateDictRequest) DictBuilder

		WithUpdateBoDict(dict *dictapi.UpdateDictRequest) DictBuilder

		WithApiDict(dict *model.SysDict) DictBuilder

		WithApiDictSelect(dict *model.SysDict) DictBuilder

		WithCreateMenuBo(menu *menuapi.CreateMenuRequest) MenuBuilder

		WithUpdateMenuBo(menu *menuapi.UpdateMenuRequest) MenuBuilder

		WithApiMenu(menu *model.SysMenu) MenuBuilder

		WithBatchCreateMenuBo(menus *menuapi.BatchCreateMenuRequest) MenuBuilder

		WithApiMenuTree(menuList []*admin.Menu, parentID uint32) MenuTreeBuilder
	}
)

func (b *builder) WithBoDatasourceQueryData(d *bo.DatasourceQueryData) DatasourceQueryDataBuilder {
	return &datasourceQueryDataBuilder{
		DatasourceQueryData: d,
		ctx:                 b.ctx,
	}
}

func (b *builder) WithDoDatasource(d *bizmodel.Datasource) DatasourceBuilder {
	return &datasourceBuilder{
		Datasource: d,
		ctx:        b.ctx,
	}
}
func (b *builder) WithApiTemplateStrategy(template *model.StrategyTemplate) TemplateBuilder {
	return &templateStrategyBuilder{
		StrategyTemplate: template,
		ctx:              b.ctx,
	}
}

func (b *builder) WithCreateBoTemplateStrategy(template *strategyapi.CreateTemplateStrategyRequest) TemplateBuilder {
	return &templateStrategyBuilder{
		CreateStrategy: template,
		ctx:            b.ctx,
	}
}

func (b *builder) WithUpdateBoTemplateStrategy(template *strategyapi.UpdateTemplateStrategyRequest) TemplateBuilder {
	return &templateStrategyBuilder{
		UpdateStrategy: template,
		ctx:            b.ctx,
	}
}

func (b *builder) WithApiTemplateStrategyLevel(template *model.StrategyLevelTemplate) TemplateLevelBuilder {
	return &templateStrategyLevelBuilder{
		StrategyLevelTemplate: template,
		ctx:                   b.ctx,
	}
}

func (b *builder) WithApiStrategy(strategy *bizmodel.Strategy) StrategyBuilder {
	return &strategyBuilder{
		Strategy: strategy,
		ctx:      b.ctx,
	}
}

func (b *builder) WithCreateBoStrategy(strategy *strategyapi.CreateStrategyRequest) StrategyBuilder {
	return &strategyBuilder{
		CreateStrategy: strategy,
		ctx:            b.ctx,
	}
}

func (b *builder) WithUpdateBoStrategy(strategy *strategyapi.UpdateStrategyRequest) StrategyBuilder {
	return &strategyBuilder{
		UpdateStrategy: strategy,
		ctx:            b.ctx,
	}
}

func (b *builder) WithCreateBoDict(dict *dictapi.CreateDictRequest) DictBuilder {
	return &dictBuilder{
		CreateDictRequest: dict,
		ctx:               b.ctx,
	}
}

func (b *builder) WithUpdateBoDict(dict *dictapi.UpdateDictRequest) DictBuilder {
	return &dictBuilder{
		UpdateDictRequest: dict,
		ctx:               b.ctx,
	}
}

func (b *builder) WithApiDict(dict *model.SysDict) DictBuilder {
	return &dictBuilder{
		SysDict: dict,
		ctx:     b.ctx,
	}
}

func (b *builder) WithApiDictSelect(dict *model.SysDict) DictBuilder {
	return &dictBuilder{
		SysDict: dict,
		ctx:     b.ctx,
	}
}

func (b *builder) WithCreateMenuBo(menu *menuapi.CreateMenuRequest) MenuBuilder {
	return &menuBuilder{
		CreateMenuRequest: menu,
		ctx:               b.ctx,
	}
}

func (b *builder) WithUpdateMenuBo(menu *menuapi.UpdateMenuRequest) MenuBuilder {
	return &menuBuilder{
		UpdateMenuRequest: menu,
		ctx:               b.ctx,
	}
}

func (b *builder) WithApiMenu(menu *model.SysMenu) MenuBuilder {
	return &menuBuilder{
		Menu: menu,
		ctx:  b.ctx,
	}
}

func (b *builder) WithBatchCreateMenuBo(menu *menuapi.BatchCreateMenuRequest) MenuBuilder {
	return &menuBuilder{
		BatchCreateMenuRequest: menu,
		ctx:                    b.ctx,
	}
}

func (b *builder) WithApiMenuTree(menuList []*admin.Menu, parentID uint32) MenuTreeBuilder {
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

func (b *builder) WithContext(ctx context.Context) Builder {
	b.ctx = ctx
	return b
}
