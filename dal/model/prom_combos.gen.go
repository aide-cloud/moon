// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNamePromCombo = "prom_combos"

// PromCombo mapped from table <prom_combos>
type PromCombo struct {
	ID        int32          `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`
	Name      string         `gorm:"column:name;type:varchar(64);not null;comment:套餐名称" json:"name"`       // 套餐名称
	Remark    string         `gorm:"column:remark;type:varchar(2048);not null;comment:套餐说明" json:"remark"` // 套餐说明
	CreatedAt time.Time      `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`
	Rules     []*PromRule    `gorm:"many2many:prom_combo_strategies" json:"rules"`
}

// TableName PromCombo's table name
func (*PromCombo) TableName() string {
	return TableNamePromCombo
}
