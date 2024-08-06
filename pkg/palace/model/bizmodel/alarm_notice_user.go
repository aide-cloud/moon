package bizmodel

import (
	"encoding/json"

	"github.com/aide-family/moon/pkg/palace/model"
	"github.com/aide-family/moon/pkg/vobj"
)

const tableNameAlarmNoticeUsers = "alarm_notice_user"

type AlarmNoticeUser struct {
	model.AllFieldModel
	AlarmNoticeType vobj.NotifyType `gorm:"column:notice_type;type:int;not null;comment:通知类型;" json:"alarm_notice_type"`
	UserId          int32           `gorm:"column:user_id;type:int;not null;comment:通知人id;uniqueIndex:idx__notice__alarm_user_id__deleted_at" json:"user_id"`
	AlarmGroupID    int32           `gorm:"column:alarm_group_id;type:int;comment:告警分组id;uniqueIndex:idx__notice__alarm_group_id__deleted_at" json:"alarm_group_id"`
}

// UnmarshalBinary redis存储实现
func (c *AlarmNoticeUser) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}

// MarshalBinary redis存储实现
func (c *AlarmNoticeUser) MarshalBinary() (data []byte, err error) {
	return json.Marshal(c)
}

// TableName Strategy's table name
func (*AlarmNoticeUser) TableName() string {
	return tableNameAlarmNoticeUsers
}
