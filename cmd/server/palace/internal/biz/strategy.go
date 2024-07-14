package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"

	"github.com/aide-family/moon/api/merr"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/repository"
	"github.com/aide-family/moon/pkg/helper/middleware"
	"github.com/aide-family/moon/pkg/palace/model/bizmodel"
	"github.com/aide-family/moon/pkg/util/types"
)

func NewStrategyBiz(dictRepo repository.Strategy) *StrategyBiz {
	return &StrategyBiz{
		strategyRepo: dictRepo,
	}
}

type StrategyBiz struct {
	strategyRepo repository.Strategy
}

// GetStrategy 获取策略
func (b *StrategyBiz) GetStrategy(ctx context.Context, param *bo.GetStrategyDetailParams) (*bizmodel.Strategy, error) {
	claims, ok := middleware.ParseJwtClaims(ctx)
	if !ok {
		return nil, merr.ErrorI18nUnLoginErr(ctx)
	}
	teamId := claims.Team
	param.TeamID = teamId
	strategy, err := b.strategyRepo.GetByID(ctx, param)
	if !types.IsNil(err) {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, merr.ErrorI18nStrategyNotFoundErr(ctx)
		}
		return nil, merr.ErrorI18nSystemErr(ctx).WithCause(err)
	}
	return strategy, nil
}

// CreateStrategy 创建策略
func (b *StrategyBiz) CreateStrategy(ctx context.Context, param *bo.CreateStrategyParams) (*bizmodel.Strategy, error) {
	_, err := b.strategyRepo.CreateStrategy(ctx, param)
	if err != nil {
		return nil, err
	}
	if !types.IsNil(err) {
		return nil, merr.ErrorI18nSystemErr(ctx).WithCause(err)
	}
	return nil, nil
}

// UpdateByID 更新策略
func (b *StrategyBiz) UpdateByID(ctx context.Context, param *bo.UpdateStrategyParams) error {
	err := b.strategyRepo.UpdateByID(ctx, param)
	if !types.IsNil(err) {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return merr.ErrorI18nStrategyNotFoundErr(ctx)
		}
		return merr.ErrorI18nSystemErr(ctx).WithCause(err)
	}
	return nil
}

// UpdateStatus 更新策略状态
func (b *StrategyBiz) UpdateStatus(ctx context.Context, param *bo.UpdateStrategyStatusParams) error {
	err := b.strategyRepo.UpdateStatus(ctx, param)
	if !types.IsNil(err) {
		return merr.ErrorI18nSystemErr(ctx).WithCause(err)
	}
	return nil
}

// DeleteByID 删除策略
func (b *StrategyBiz) DeleteByID(ctx context.Context, param *bo.DelStrategyParams) error {
	err := b.strategyRepo.DeleteByID(ctx, param)
	if !types.IsNil(err) {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return merr.ErrorI18nStrategyNotFoundErr(ctx)
		}
		return merr.ErrorI18nSystemErr(ctx).WithCause(err)
	}
	return nil
}

// StrategyPage 获取策略分页
func (b *StrategyBiz) StrategyPage(ctx context.Context, param *bo.QueryStrategyListParams) ([]*bizmodel.Strategy, error) {
	strategies, err := b.strategyRepo.FindByPage(ctx, param)
	if !types.IsNil(err) {
		return nil, merr.ErrorI18nSystemErr(ctx).WithCause(err)
	}
	return strategies, nil
}
