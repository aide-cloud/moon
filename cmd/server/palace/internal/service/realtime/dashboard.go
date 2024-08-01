package realtime

import (
	"context"

	pb "github.com/aide-family/moon/api/admin/realtime"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz"
	"github.com/aide-family/moon/cmd/server/palace/internal/service/build"
)

// DashboardService 监控大盘服务
type DashboardService struct {
	pb.UnimplementedDashboardServer

	dashboardBiz *biz.DashboardBiz
}

// NewDashboardService 创建监控大盘服务
func NewDashboardService(dashboardBiz *biz.DashboardBiz) *DashboardService {
	return &DashboardService{
		dashboardBiz: dashboardBiz,
	}
}

// CreateDashboard 创建监控大盘
func (s *DashboardService) CreateDashboard(ctx context.Context, req *pb.CreateDashboardRequest) (*pb.CreateDashboardReply, error) {
	params := build.NewBuilder().WithContext(ctx).DashboardModule().WithAddDashboardParams(req).ToBo()
	if err := s.dashboardBiz.CreateDashboard(ctx, params); err != nil {
		return nil, err
	}
	return &pb.CreateDashboardReply{}, nil
}

// UpdateDashboard 更新监控大盘
func (s *DashboardService) UpdateDashboard(ctx context.Context, req *pb.UpdateDashboardRequest) (*pb.UpdateDashboardReply, error) {
	params := build.NewBuilder().
		WithContext(ctx).
		DashboardModule().
		WithUpdateDashboardParams(req).
		ToBo()
	if err := s.dashboardBiz.UpdateDashboard(ctx, params); err != nil {
		return nil, err
	}
	return &pb.UpdateDashboardReply{}, nil
}

// DeleteDashboard 删除监控大盘
func (s *DashboardService) DeleteDashboard(ctx context.Context, req *pb.DeleteDashboardRequest) (*pb.DeleteDashboardReply, error) {
	params := build.NewBuilder().
		WithContext(ctx).
		DashboardModule().
		WithDeleteDashboardParams(req).
		ToBo()
	if err := s.dashboardBiz.DeleteDashboard(ctx, params); err != nil {
		return nil, err
	}
	return &pb.DeleteDashboardReply{}, nil
}

// GetDashboard 获取监控大盘
func (s *DashboardService) GetDashboard(ctx context.Context, req *pb.GetDashboardRequest) (*pb.GetDashboardReply, error) {
	detail, err := s.dashboardBiz.GetDashboard(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	apiDetail := build.NewBuilder().
		WithContext(ctx).
		DashboardModule().
		WithDoDashboard(detail).
		ToAPI()
	return &pb.GetDashboardReply{
		Detail: apiDetail,
	}, nil
}

// ListDashboard 获取监控大盘列表
func (s *DashboardService) ListDashboard(ctx context.Context, req *pb.ListDashboardRequest) (*pb.ListDashboardReply, error) {
	params := build.NewBuilder().
		WithContext(ctx).
		DashboardModule().
		WithQueryDashboardListParams(req).
		ToBo()
	list, err := s.dashboardBiz.ListDashboard(ctx, params)
	if err != nil {
		return nil, err
	}
	apiList := build.NewBuilder().
		WithContext(ctx).
		DashboardModule().
		WithDoDashboardList(list).
		ToAPIs()
	return &pb.ListDashboardReply{
		List:       apiList,
		Pagination: build.NewPageBuilder(params.Page).ToAPI(),
	}, nil
}

// ListDashboardSelect 获取监控大盘下拉列表
func (s *DashboardService) ListDashboardSelect(ctx context.Context, req *pb.ListDashboardSelectRequest) (*pb.ListDashboardSelectReply, error) {
	params := build.NewBuilder().
		WithContext(ctx).
		DashboardModule().
		WithQueryDashboardSelectParams(req).
		ToBo()
	list, err := s.dashboardBiz.ListDashboard(ctx, params)
	if err != nil {
		return nil, err
	}
	apiList := build.NewBuilder().
		WithContext(ctx).
		DashboardModule().
		WithDoDashboardList(list).
		ToSelects()
	return &pb.ListDashboardSelectReply{
		List:       apiList,
		Pagination: build.NewPageBuilder(params.Page).ToAPI(),
	}, nil
}
