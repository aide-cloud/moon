package repository

import (
	"context"

	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-family/moon/pkg/palace/model"
)

type (
	Dict interface {
		// Create 创建字典
		Create(ctx context.Context, user *bo.CreateDictParams) (model.IDict, error)

		// BatchCreate 批量创建字典
		BatchCreate(ctx context.Context, users []*bo.CreateDictParams) error

		// GetByID 通过id 获取字典详情
		GetByID(ctx context.Context, id uint32) (model.IDict, error)

		// FindByPage 分页查询字典列表
		FindByPage(ctx context.Context, page *bo.QueryDictListParams) ([]model.IDict, error)

		// DeleteByID 通过ID删除字典
		DeleteByID(ctx context.Context, params *bo.DeleteDictParams) error

		// UpdateStatusByIds 通过ID列表批量更新字典状态
		UpdateStatusByIds(ctx context.Context, updateParams *bo.UpdateDictStatusParams) error

		// UpdateByID 通过ID更新字典数据
		UpdateByID(ctx context.Context, dict *bo.UpdateDictParams) error
	}
)
