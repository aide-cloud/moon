package repository

import (
	"context"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-family/moon/pkg/helper/model/palace"
	"github.com/aide-family/moon/pkg/vobj"
)

type (
	Dict interface {
		// Create a dict
		Create(ctx context.Context, user *bo.CreateDictParams) (*palace.SysDict, error)

		// BatchCreate batch create dict
		BatchCreate(ctx context.Context, users []*bo.CreateDictParams) error

		// GetByID get dict by id
		GetByID(ctx context.Context, id uint32) (*palace.SysDict, error)

		//FindByPage
		FindByPage(ctx context.Context, page *bo.QueryDictListParams) ([]*palace.SysDict, error)

		//DeleteByID
		DeleteByID(ctx context.Context, id uint32) error

		//UpdateStatusByIds
		UpdateStatusByIds(ctx context.Context, status vobj.Status, ids ...uint32) error

		// UpdateByID update dict by id
		UpdateByID(ctx context.Context, dict *bo.UpdateDictParams) error
	}
)
