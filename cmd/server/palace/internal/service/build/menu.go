package build

import (
	"github.com/aide-family/moon/api"
	"github.com/aide-family/moon/api/admin"
	"github.com/aide-family/moon/pkg/palace/model"
	"github.com/aide-family/moon/pkg/palace/model/bizmodel"
	"github.com/aide-family/moon/pkg/util/types"
)

type MenuBuilder struct {
	Menu *model.SysMenu
}

func NewMenuBuilder(menu *model.SysMenu) *MenuBuilder {
	return &MenuBuilder{
		Menu: menu,
	}
}

func (b *MenuBuilder) ToApi() *admin.Menu {
	if types.IsNil(b) || types.IsNil(b.Menu) {
		return nil
	}
	return &admin.Menu{
		Id:         b.Menu.ID,
		Name:       b.Menu.Name,
		Path:       b.Menu.Path,
		Icon:       b.Menu.Icon,
		Status:     api.Status(b.Menu.Status),
		ParentId:   b.Menu.ParentID,
		Sort:       b.Menu.Sort,
		Type:       api.MenuType(b.Menu.Type),
		Level:      b.Menu.Level,
		Component:  b.Menu.Component,
		Permission: b.Menu.Permission,
		EnName:     b.Menu.EnName,
		CreatedAt:  b.Menu.CreatedAt.String(),
		UpdatedAt:  b.Menu.UpdatedAt.String(),
	}
}

type MenuTreeBuilder struct {
	menuMap  map[uint32][]*bizmodel.SysTeamMenu
	parentID uint32
}

func NewMenuTreeBuilder(menuList []*bizmodel.SysTeamMenu, parentID uint32) *MenuTreeBuilder {
	menuMap := make(map[uint32][]*bizmodel.SysTeamMenu)
	// 按照父级ID分组
	for _, menu := range menuList {
		if _, ok := menuMap[menu.ParentID]; !ok {
			menuMap[menu.ParentID] = make([]*bizmodel.SysTeamMenu, 0)
		}
		menuMap[menu.ParentID] = append(menuMap[menu.ParentID], menu)
	}
	return &MenuTreeBuilder{
		menuMap:  menuMap,
		parentID: parentID,
	}
}

// ToTree 转换为树形菜单
func (b *MenuTreeBuilder) ToTree() []*admin.Menu {
	if types.IsNil(b) || types.IsNil(b.menuMap) || len(b.menuMap) == 0 {
		return nil
	}
	list := make([]*admin.Menu, 0)
	// 递归遍历
	for _, menu := range b.menuMap[b.parentID] {
		if menu.ParentID == b.parentID {
			list = append(list, &admin.Menu{
				Id:        menu.ID,
				Name:      menu.Name,
				Path:      menu.Path,
				Icon:      menu.Icon,
				Status:    api.Status(menu.Status),
				ParentId:  menu.ParentID,
				CreatedAt: menu.CreatedAt.String(),
				UpdatedAt: menu.UpdatedAt.String(),
				Level:     menu.Level,
				//Children:  NewMenuTreeBuilder(b.menuMap[menu.ID], menu.ID).ToTree(),
			})
		}
	}
	return list
}
