package realtime

import (
	"context"

	pb "github.com/aide-family/moon/api/admin/realtime"
)

type DashboardService struct {
	pb.UnimplementedDashboardServer
}

func NewDashboardService() *DashboardService {
	return &DashboardService{}
}

func (s *DashboardService) CreateDashboard(ctx context.Context, req *pb.CreateDashboardRequest) (*pb.CreateDashboardReply, error) {
	return &pb.CreateDashboardReply{}, nil
}

func (s *DashboardService) UpdateDashboard(ctx context.Context, req *pb.UpdateDashboardRequest) (*pb.UpdateDashboardReply, error) {
	return &pb.UpdateDashboardReply{}, nil
}

func (s *DashboardService) DeleteDashboard(ctx context.Context, req *pb.DeleteDashboardRequest) (*pb.DeleteDashboardReply, error) {
	return &pb.DeleteDashboardReply{}, nil
}

func (s *DashboardService) GetDashboard(ctx context.Context, req *pb.GetDashboardRequest) (*pb.GetDashboardReply, error) {
	return &pb.GetDashboardReply{}, nil
}

func (s *DashboardService) ListDashboard(ctx context.Context, req *pb.ListDashboardRequest) (*pb.ListDashboardReply, error) {
	return &pb.ListDashboardReply{}, nil
}

func (s *DashboardService) ListDashboardSelect(ctx context.Context, req *pb.ListDashboardSelectRequest) (*pb.ListDashboardSelectReply, error) {
	return &pb.ListDashboardSelectReply{}, nil
}
