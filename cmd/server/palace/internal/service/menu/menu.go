package menu

import (
	"context"
	"github.com/aide-family/moon/api/admin"
	menuapi "github.com/aide-family/moon/api/admin/menu"
	"github.com/aide-family/moon/api/merr"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-family/moon/cmd/server/palace/internal/service/build"
	"github.com/aide-family/moon/pkg/palace/model"
	"github.com/aide-family/moon/pkg/util/types"
	"github.com/aide-family/moon/pkg/vobj"
	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
)

type MenuService struct {
	menuapi.UnimplementedMenuServer
	menuBiz *biz.MenuBiz
}

func NewMenuService(menuBiz *biz.MenuBiz) *MenuService {
	return &MenuService{
		menuBiz: menuBiz,
	}
}

func (m *MenuService) CreateMenu(ctx context.Context, req *menuapi.CreateMenuRequest) (*menuapi.CreateMenuReply, error) {
	createParams := bo.CreateMenuParams{
		Name:       req.GetName(),
		Path:       req.GetPath(),
		Component:  req.GetComponent(),
		Type:       vobj.MenuType(req.GetMenuType()),
		Status:     vobj.Status(req.GetStatus()),
		Icon:       req.GetIcon(),
		Permission: req.GetPermission(),
		ParentId:   req.GetParentId(),
		EnName:     req.GetEnName(),
		Sort:       req.GetSort(),
		Level:      req.GetLevel(),
	}
	_, err := m.menuBiz.CreateMenu(ctx, &createParams)
	if err != nil {
		return nil, err
	}
	if !types.IsNil(err) {
		return nil, err
	}
	return &menuapi.CreateMenuReply{}, nil
}

func (m *MenuService) BatchCreateMenu(ctx context.Context, request *menuapi.BatchCreateMenuRequest) (*menuapi.BatchCreateMenuReply, error) {
	return &menuapi.BatchCreateMenuReply{}, nil
}
func (m *MenuService) UpdateMenu(ctx context.Context, req *menuapi.UpdateMenuRequest) (*menuapi.UpdateMenuReply, error) {
	data := req.GetData()
	createParams := bo.CreateMenuParams{
		Name:       data.GetName(),
		Path:       data.GetPath(),
		Component:  data.GetComponent(),
		Type:       vobj.MenuType(data.GetMenuType()),
		Status:     vobj.Status(data.GetStatus()),
		Icon:       data.GetIcon(),
		Permission: data.GetPermission(),
		ParentId:   data.GetParentId(),
		EnName:     data.GetEnName(),
		Sort:       data.GetSort(),
		Level:      data.GetLevel(),
	}
	updateParams := bo.UpdateMenuParams{
		ID:          req.GetId(),
		UpdateParam: createParams,
	}
	if err := m.menuBiz.UpdateMenu(ctx, &updateParams); !types.IsNil(err) {
		return nil, err
	}
	return &menuapi.UpdateMenuReply{}, nil
}

func (m *MenuService) DeleteMenu(ctx context.Context, req *menuapi.DeleteMenuRequest) (*menuapi.DeleteMenuReply, error) {

	err := m.menuBiz.DeleteMenu(ctx, req.GetId())
	if err != nil {
		return nil, merr.ErrorI18nSystemErr(ctx).WithCause(err)
	}
	return &menuapi.DeleteMenuReply{}, nil
}

func (m *MenuService) GetMenu(ctx context.Context, req *menuapi.GetMenuRequest) (*menuapi.GetMenuReply, error) {
	data, err := m.menuBiz.GetMenu(ctx, req.GetId())
	if !types.IsNil(err) {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, merr.ErrorI18nMenuNotFoundErr(ctx)
		}
		return nil, err
	}
	resMenu := build.NewMenuBuilder(data).ToApi()
	return &menuapi.GetMenuReply{
		Menu: resMenu,
	}, nil
}

func (m *MenuService) TreeMenu(ctx context.Context, req *menuapi.TreeMenuRequest) (*menuapi.TreeMenuReply, error) {

	return nil, nil
}

func (m *MenuService) MenuListPage(ctx context.Context, req *menuapi.ListMenuRequest) (*menuapi.ListMenuReply, error) {
	queryParams := &bo.QueryMenuListParams{
		Keyword:  req.GetKeyword(),
		Page:     types.NewPagination(req.GetPagination()),
		Status:   vobj.Status(req.GetStatus()),
		MenuType: vobj.MenuType(req.GetMenuType()),
	}
	menuPage, err := m.menuBiz.ListMenuPage(ctx, queryParams)
	if !types.IsNil(err) {
		return nil, err
	}

	return &menuapi.ListMenuReply{
		Menu: types.SliceTo(menuPage, func(menu *model.SysMenu) *admin.Menu {
			return build.NewMenuBuilder(menu).ToApi()
		}),
		Pagination: build.NewPageBuilder(queryParams.Page).ToApi(),
	}, nil
}

func (m *MenuService) BatchUpdateDictStatus(ctx context.Context, req *menuapi.BatchUpdateMenuStatusRequest) (*menuapi.BatchUpdateMenuStatusReply, error) {
	params := &bo.UpdateMenuStatusParams{
		IDs:    req.Ids,
		Status: vobj.Status(req.Status),
	}
	err := m.menuBiz.UpdateMenuStatus(ctx, params)
	if err != nil {
		return nil, merr.ErrorI18nSystemErr(ctx).WithCause(err)
	}
	return &menuapi.BatchUpdateMenuStatusReply{}, nil
}

func (m *MenuService) BatchUpdateMenuType(ctx context.Context, req *menuapi.BatchUpdateMenuTypeRequest) (*menuapi.BatchUpdateMenuTypeReply, error) {

	params := &bo.UpdateMenuTypeParams{
		IDs:  req.Ids,
		Type: vobj.MenuType(req.GetType()),
	}

	err := m.menuBiz.UpdateMenuTypes(ctx, params)
	if err != nil {
		return nil, merr.ErrorI18nSystemErr(ctx).WithCause(err)
	}
	return &menuapi.BatchUpdateMenuTypeReply{}, nil
}
