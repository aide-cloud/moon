package bizmodel

import (
	"encoding/json"

	"github.com/aide-family/moon/pkg/palace/model"
	"github.com/aide-family/moon/pkg/vobj"
)

const tableNameAlarmGroup = "alarm_group"

// AlarmGroup 告警组
type AlarmGroup struct {
	model.AllFieldModel
	Name    string           `gorm:"column:name;type:varchar(64);not null;uniqueIndex:idx__name,priority:1;comment:告警组名称"`
	Status  vobj.Status      `gorm:"column:status;type:tinyint;not null;default:1;comment:启用状态1:启用;2禁用"`
	Remark  string           `gorm:"column:remark;type:varchar(255);not null;comment:描述信息"`
	Members []*SysTeamMember `gorm:"many2many:alarm_group_members" json:"members"`
}

// UnmarshalBinary redis存储实现
func (c *AlarmGroup) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}

// MarshalBinary redis存储实现
func (c *AlarmGroup) MarshalBinary() (data []byte, err error) {
	return json.Marshal(c)
}

// TableName Strategy's table name
func (*AlarmGroup) TableName() string {
	return tableNameAlarmGroup
}
