package model

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"

	"github.com/aide-family/moon/pkg/helper/middleware"
	"github.com/aide-family/moon/pkg/util/types"
	"github.com/aide-family/moon/pkg/vobj"
)

type BaseModel struct {
	ctx context.Context `gorm:"-"`

	CreatedAt types.Time            `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"`
	UpdatedAt types.Time            `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint;not null;default:0;" json:"deleted_at"`

	// 创建人
	CreatorID uint32 `gorm:"column:creator;type:int unsigned;not null;comment:创建者" json:"creator_id"`
}

type AllFieldModel struct {
	ID uint32 `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`
	BaseModel
}

// WithContext 获取上下文
func (u *BaseModel) WithContext(ctx context.Context) *BaseModel {
	u.ctx = ctx
	return u
}

func (u *BaseModel) BeforeCreate(_ *gorm.DB) (err error) {
	if u.ctx == nil {
		return
	}
	claims, ok := middleware.ParseJwtClaims(u.ctx)
	if !ok {
		return
	}
	u.CreatorID = claims.GetUser()
	return
}

func (u *BaseModel) GetContext() context.Context {
	if types.IsNil(u.ctx) {
		return context.TODO()
	}
	return u.ctx
}

type IDict interface {
	GetID() uint32
	GetName() string
	GetValue() string
	GetDictType() vobj.DictType
	GetColorType() string
	GetCssClass() string
	GetIcon() string
	GetImageUrl() string
	GetStatus() vobj.Status
	GetLanguageCode() string
	GetRemark() string
	GetCreatedAt() *types.Time
	GetUpdatedAt() *types.Time
	GetCreatorID() uint32
	GetDeletedAt() soft_delete.DeletedAt
}
