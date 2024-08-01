package realtime

import (
	"context"

	"github.com/aide-family/moon/api"
	adminapi "github.com/aide-family/moon/api/admin"
	pb "github.com/aide-family/moon/api/admin/realtime"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-family/moon/cmd/server/palace/internal/service/build"
	"github.com/aide-family/moon/pkg/palace/model/bizmodel"
	"github.com/aide-family/moon/pkg/util/types"
	"github.com/aide-family/moon/pkg/vobj"
)

// AlarmService 实时告警数据服务
type AlarmService struct {
	pb.UnimplementedAlarmServer

	alarmBiz *biz.AlarmBiz
}

// NewAlarmService 实时告警数据服务
func NewAlarmService(alarmBiz *biz.AlarmBiz) *AlarmService {
	return &AlarmService{
		alarmBiz: alarmBiz,
	}
}

// GetAlarm 获取实时告警数据
func (s *AlarmService) GetAlarm(ctx context.Context, req *pb.GetAlarmRequest) (*pb.GetAlarmReply, error) {
	realtimeAlarmDetail, err := s.alarmBiz.GetRealTimeAlarm(ctx, &bo.GetRealTimeAlarmParams{
		RealtimeAlarmID: req.GetId(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.GetAlarmReply{
		Detail: build.NewBuilder().WithContext(ctx).WithRealTimeAlarm(realtimeAlarmDetail).ToAPI(),
	}, nil
}

// ListAlarm 获取实时告警数据列表
func (s *AlarmService) ListAlarm(ctx context.Context, req *pb.ListAlarmRequest) (*pb.ListAlarmReply, error) {
	params := &bo.GetRealTimeAlarmsParams{
		Pagination:      types.NewPagination(req.GetPagination()),
		EventAtStart:    req.GetEventAtStart(),
		EventAtEnd:      req.GetEventAtEnd(),
		ResolvedAtStart: req.GetRecoverAtStart(),
		ResolvedAtEnd:   req.GetRecoverAtEnd(),
		AlarmLevels:     req.GetAlarmLevels(),
		AlarmStatuses:   types.SliceTo(req.GetAlarmStatuses(), func(item api.AlertStatus) vobj.AlertStatus { return vobj.AlertStatus(item) }),
		Keyword:         req.GetKeyword(),
	}
	list, err := s.alarmBiz.ListRealTimeAlarms(ctx, params)
	if err != nil {
		return nil, err
	}
	return &pb.ListAlarmReply{
		List: types.SliceTo(list, func(item *bizmodel.RealtimeAlarm) *adminapi.RealtimeAlarmItem {
			return build.NewBuilder().WithContext(ctx).WithRealTimeAlarm(item).ToAPI()
		}),
		Pagination: build.NewPageBuilder(params.Pagination).ToAPI(),
	}, nil
}
