package bizmodel

import (
	"encoding/json"

	"github.com/aide-family/moon/pkg/palace/model"
	"github.com/aide-family/moon/pkg/vobj"
)

const TableNameStrategy = "strategies"

// Strategy mapped from table <Strategy>
type Strategy struct {
	model.AllFieldModel
	// 模板ID, 用于标记是否从模板创建而来
	StrategyTemplateID uint32 `gorm:"column:strategy_template_id;type:int unsigned;not null;comment:策略模板ID" json:"strategy_template_id"`
	// 策略模板来源（系统、团队）
	StrategyTemplateSource vobj.StrategyTemplateSource `gorm:"column:strategy_template_source;type:tinyint;not null;comment:策略模板来源（系统、团队）" json:"strategy_template_source"`
	Name                   string                      `gorm:"column:alert;type:varchar(64);not null;comment:模板名称" json:"name"`
	Alert                  string                      `gorm:"column:alert;type:varchar(64);not null;comment:策略模板名称" json:"alert"`
	Expr                   string                      `gorm:"column:expr;type:text;not null;comment:告警表达式" json:"expr"`
	Labels                 *vobj.Labels                `gorm:"column:labels;type:JSON;not null;comment:标签" json:"labels"`
	Annotations            vobj.Annotations            `gorm:"column:annotations;type:JSON;not null;comment:注解" json:"annotations"`
	Remark                 string                      `gorm:"column:remark;type:varchar(255);not null;comment:备注" json:"remark"`
	Status                 vobj.Status                 `gorm:"column:status;type:int;not null;comment:策略状态" json:"status"`
	// 采样率
	Step uint32 `gorm:"column:step;type:int;not null;comment:采样率" json:"step"`

	Datasource []*Datasource `gorm:"many2many:strategy_datasource;" json:"datasource"`
	// 策略类型
	Categories []*SysDict `gorm:"many2many:strategy_categories"`
	// 条件
	Condition string `gorm:"column:condition;type:varchar(2);not null;comment:条件" json:"condition"`
	// 阈值
	Threshold float64 `gorm:"column:threshold;type:text;not null;comment:阈值" json:"threshold"`
}

// String json string
func (c *Strategy) String() string {
	bs, _ := json.Marshal(c)
	return string(bs)
}

func (c *Strategy) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}

func (c *Strategy) MarshalBinary() (data []byte, err error) {
	return json.Marshal(c)
}

// TableName Strategy's table name
func (*Strategy) TableName() string {
	return TableNameStrategy
}
