package biz

import (
	"context"
	pb "github.com/aide-family/moon/api/admin/dict"
	"github.com/aide-family/moon/api/merr"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/repository"
	"github.com/aide-family/moon/pkg/helper/model/palace"
	"github.com/aide-family/moon/pkg/util/types"
)

func NewDictBiz(dictRepo repository.Dict) *DictBiz {
	return &DictBiz{
		dictRepo: dictRepo,
	}
}

type DictBiz struct {
	dictRepo repository.Dict
}

// CreateDict 创建字典
func (b *DictBiz) CreateDict(ctx context.Context, dictParam *bo.CreateDictParams) (*palace.SysDict, error) {
	dictDo, err := b.dictRepo.Create(ctx, dictParam)
	if !types.IsNil(err) {
		return nil, merr.ErrorI18nSystemErr(ctx).WithCause(err)
	}
	return dictDo, nil
}

// UpdateDict 更新字典
func (b *DictBiz) UpdateDict(ctx context.Context, updateParam *bo.UpdateDictParams) (*pb.UpdateDictReply, error) {
	if err := b.dictRepo.UpdateByID(ctx, updateParam); !types.IsNil(err) {
		return nil, err
	}
	return &pb.UpdateDictReply{}, nil
}

// ListDict 列表字典
func (b *DictBiz) ListDict(ctx context.Context, listParam *bo.QueryDictListParams) ([]*palace.SysDict, error) {
	dictDos, err := b.dictRepo.FindByPage(ctx, listParam)
	if !types.IsNil(err) {
		return nil, merr.ErrorI18nSystemErr(ctx).WithCause(err)
	}
	return dictDos, nil

}
