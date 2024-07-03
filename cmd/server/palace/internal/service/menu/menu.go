package menu

import (
	"github.com/aide-family/moon/cmd/server/palace/internal/biz"
)

type MenuService struct {
	menuBiz *biz.MenuBiz
}

func NewMenuService(menuBiz *biz.MenuBiz) *MenuService {
	return &MenuService{
		menuBiz: menuBiz,
	}
}
