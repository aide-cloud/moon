package bo

import (
	"github.com/aide-family/moon/pkg/util/types"
	"github.com/aide-family/moon/pkg/vobj"
)

type (

	// CreateAlarmGroupParams 创建告警组请求参数
	CreateAlarmGroupParams struct {
		// 告警组名称
		Name string `json:"name,omitempty"`
		// 告警组说明信息
		Remark string `json:"remark,omitempty"`
		// 告警组状态
		Status vobj.Status `json:"status,omitempty"`
		// 告警分组通知人
		NoticeUsers []*CreateNoticeUserParams `json:"noticeUsers,omitempty"`
	}

	// CreateNoticeUserParams 创建通知人参数
	CreateNoticeUserParams struct {
		// 用户id
		UserId uint32
		// 通知方式
		NotifyType vobj.NotifyType
	}

	// UpdateAlarmGroupStatusParams 更新告警组状态请求参数
	UpdateAlarmGroupStatusParams struct {
		IDs    []uint32 `json:"ids"`
		Status vobj.Status
	}

	// UpdateAlarmGroupParams 更新告警组请求参数
	UpdateAlarmGroupParams struct {
		ID          uint32 `json:"id"`
		UpdateParam *CreateAlarmGroupParams
	}

	// QueryAlarmGroupListParams 查询告警组列表请求参数
	QueryAlarmGroupListParams struct {
		Keyword string `json:"keyword"`
		Page    types.Pagination
		Name    string
		Status  vobj.Status
	}

	// GetRealTimeAlarmParams 获取实时告警参数
	GetRealTimeAlarmParams struct {
		// 告警ID
		RealtimeAlarmID uint32
		// 告警指纹
		Fingerprint string
	}

	// GetRealTimeAlarmsParams 获取实时告警列表参数
	GetRealTimeAlarmsParams struct {
		// 分页参数
		Pagination types.Pagination
		// 告警时间范围
		EventAtStart int64
		EventAtEnd   int64
		// 告警恢复时间
		ResolvedAtStart int64
		ResolvedAtEnd   int64
		// 告警级别
		AlarmLevels []uint32
		// 告警状态
		AlarmStatuses []vobj.AlertStatus
		// 关键字
		Keyword string
		// 告警页面
		AlarmPageID uint32
		// 我的告警
		MyAlarm bool
	}
)
