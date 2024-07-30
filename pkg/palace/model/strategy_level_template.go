package model

import (
	"encoding/json"

	"github.com/aide-family/moon/pkg/util/types"
	"github.com/aide-family/moon/pkg/vobj"
)

const tableNameStrategyLevelTemplates = "strategy_level_templates"

// StrategyLevelTemplate 策略等级明细模型
type StrategyLevelTemplate struct {
	AllFieldModel
	// 所属策略模板
	StrategyTemplateID uint32            `gorm:"column:strategy_template_id;type:int unsigned;not null;comment:策略模板ID" json:"strategy_template_id"`
	StrategyTemplate   *StrategyTemplate `gorm:"foreignKey:StrategyTemplateID" json:"strategy_template"`

	// 持续时间
	Duration *types.Duration `gorm:"column:duration;type:bigint(20);not null;comment:告警持续时间" json:"duration"`
	// 持续次数
	Count uint32 `gorm:"column:count;type:int unsigned;not null;comment:持续次数" json:"count"`
	// 持续事件类型
	SustainType vobj.Sustain `gorm:"column:sustain_type;type:int(11);not null;comment:持续类型" json:"sustain_type"`
	// 条件
	Condition vobj.Condition `gorm:"column:condition;type:int;not null;comment:条件" json:"condition"`
	// 阈值
	Threshold float64 `gorm:"column:threshold;type:text;not null;comment:阈值" json:"threshold"`
	// 告警等级 对应sys_dict字典id
	LevelID uint32   `gorm:"column:level_id;type:int unsigned;not null;comment:告警等级" json:"level_id"`
	Level   *SysDict `gorm:"foreignKey:LevelID" json:"level"`

	// 状态
	Status vobj.Status `gorm:"column:status;type:int;not null;comment:策略状态" json:"status"`
}

// String json string
func (c *StrategyLevelTemplate) String() string {
	bs, _ := json.Marshal(c)
	return string(bs)
}

// UnmarshalBinary StrategyLevelTemplate's UnmarshalBinary
func (c *StrategyLevelTemplate) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}

// MarshalBinary StrategyLevelTemplate's MarshalBinary
func (c *StrategyLevelTemplate) MarshalBinary() (data []byte, err error) {
	return json.Marshal(c)
}

// TableName StrategyLevelTemplate's table name
func (*StrategyLevelTemplate) TableName() string {
	return tableNameStrategyLevelTemplates
}
