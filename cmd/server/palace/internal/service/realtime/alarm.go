package realtime

import (
	"context"

	pb "github.com/aide-family/moon/api/admin/realtime"
)

type AlarmService struct {
	pb.UnimplementedAlarmServer
}

func NewAlarmService() *AlarmService {
	return &AlarmService{}
}

func (s *AlarmService) GetAlarm(ctx context.Context, req *pb.GetAlarmRequest) (*pb.GetAlarmReply, error) {
	return &pb.GetAlarmReply{}, nil
}

func (s *AlarmService) ListAlarm(ctx context.Context, req *pb.ListAlarmRequest) (*pb.ListAlarmReply, error) {
	return &pb.ListAlarmReply{}, nil
}
