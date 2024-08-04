package alarm

import (
	"context"

	alarmyapi "github.com/aide-family/moon/api/admin/alarm"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz"
	"github.com/aide-family/moon/cmd/server/palace/internal/service/build"
	"github.com/aide-family/moon/pkg/util/types"
)

// AlarmGroupService 告警管理服务
type AlarmGroupService struct {
	alarmyapi.UnimplementedAlarmServer
	alarmGroupBiz *biz.AlarmGroupBiz
}

// NewAlarmService 创建告警管理服务
func NewAlarmService(alarmGroupBiz *biz.AlarmGroupBiz) *AlarmGroupService {
	return &AlarmGroupService{
		alarmGroupBiz: alarmGroupBiz,
	}
}

// CreateAlarmGroup 创建告警组
func (s *AlarmGroupService) CreateAlarmGroup(ctx context.Context, req *alarmyapi.CreateAlarmGroupRequest) (*alarmyapi.CreateAlarmGroupReply, error) {
	param := build.NewBuilder().WithContext(ctx).AlarmModule().WithAPICreateAlarmGroupRequest(req).ToBo()
	if _, err := s.alarmGroupBiz.CreateAlarmGroup(ctx, param); !types.IsNil(err) {
		return nil, err
	}
	return &alarmyapi.CreateAlarmGroupReply{}, nil
}

// DeleteAlarmGroup 删除告警组
func (s *AlarmGroupService) DeleteAlarmGroup(ctx context.Context, req *alarmyapi.DeleteAlarmGroupRequest) (*alarmyapi.DeleteAlarmGroupReply, error) {
	if err := s.alarmGroupBiz.DeleteAlarmGroup(ctx, req.GetId()); !types.IsNil(err) {
		return nil, err
	}
	return &alarmyapi.DeleteAlarmGroupReply{}, nil
}

// ListAlarmGroup 获取告警组列表
func (s *AlarmGroupService) ListAlarmGroup(ctx context.Context, req *alarmyapi.ListAlarmGroupRequest) (*alarmyapi.ListAlarmGroupReply, error) {
	param := build.NewBuilder().WithContext(ctx).AlarmModule().WithAPIQueryAlarmGroupListRequest(req).ToBo()
	alarmGroups, err := s.alarmGroupBiz.ListPage(ctx, param)
	if !types.IsNil(err) {
		return nil, err
	}
	return &alarmyapi.ListAlarmGroupReply{
		Pagination: build.NewPageBuilder(param.Page).ToAPI(),
		List: build.NewBuilder().
			WithContext(ctx).
			AlarmModule().
			WithDosAlarmGroup(alarmGroups).
			ToAPIs(),
	}, nil
}

// GetAlarmGroup 获取告警组详细信息
func (s *AlarmGroupService) GetAlarmGroup(ctx context.Context, req *alarmyapi.GetAlarmGroupRequest) (*alarmyapi.GetAlarmGroupReply, error) {
	detail, err := s.alarmGroupBiz.GetAlarmGroupDetail(ctx, req.GetId())
	if !types.IsNil(err) {
		return nil, err
	}
	return &alarmyapi.GetAlarmGroupReply{Detail: build.NewBuilder().
		WithContext(ctx).AlarmModule().
		WithDoAlarmGroup(detail).ToAPI(),
	}, nil
}

// UpdateAlarmGroup 更新告警组信息
func (s *AlarmGroupService) UpdateAlarmGroup(ctx context.Context, req *alarmyapi.UpdateAlarmGroupRequest) (*alarmyapi.UpdateAlarmGroupReply, error) {
	param := build.NewBuilder().WithContext(ctx).AlarmModule().WithAPIUpdateAlarmGroupRequest(req).ToBo()
	err := s.alarmGroupBiz.UpdateAlarmGroup(ctx, param)
	if !types.IsNil(err) {
		return nil, err
	}
	return &alarmyapi.UpdateAlarmGroupReply{}, nil
}

// UpdateAlarmGroupStatus 更新告警组状态
func (s *AlarmGroupService) UpdateAlarmGroupStatus(ctx context.Context, req *alarmyapi.UpdateAlarmGroupStatusRequest) (*alarmyapi.UpdateAlarmGroupStatusReply, error) {
	param := build.NewBuilder().WithContext(ctx).AlarmModule().WithAPIUpdateStatusAlarmGroupRequest(req).ToBo()
	err := s.alarmGroupBiz.UpdateStatus(ctx, param)
	if !types.IsNil(err) {
		return nil, err
	}
	return &alarmyapi.UpdateAlarmGroupStatusReply{}, nil
}
