package repoimpl

import (
	"context"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/repository"
	"github.com/aide-family/moon/cmd/server/palace/internal/data"
	"github.com/aide-family/moon/pkg/palace/model/bizmodel"
	"github.com/aide-family/moon/pkg/util/types"
	"github.com/aide-family/moon/pkg/vobj"
)

func NewStrategyRepository(data *data.Data) repository.Strategy {
	return &strategyRepositoryImpl{
		data: data,
	}
}

type strategyRepositoryImpl struct {
	data *data.Data
}

func (s *strategyRepositoryImpl) DeleteByID(ctx context.Context, id uint32) error {

	return nil
}

func (s *strategyRepositoryImpl) UpdateStrategy(ctx context.Context, updateParams *bo.UpdateStrategyParams) error {

	return nil
}

func (s *strategyRepositoryImpl) CreateStrategy(ctx context.Context, params *bo.CreateStrategyParams) (*bizmodel.Strategy, error) {
	bizDB, err := s.data.GetBizGormDB(params.TeamID)
	if !types.IsNil(err) {
		return nil, err
	}
	bizDB.WithContext(ctx).Create(&bizmodel.Strategy{})
	return nil, nil
}

func (s *strategyRepositoryImpl) GetStrategy(ctx context.Context, id int64) (*bizmodel.Strategy, error) {
	return nil, nil
}

func (s *strategyRepositoryImpl) UpdateByID(ctx context.Context, params *bo.UpdateStrategyParams) error {
	return nil
}

func (s *strategyRepositoryImpl) UpdateStatus(ctx context.Context, status vobj.Status, ids ...uint32) error {
	return nil
}

func (s *strategyRepositoryImpl) GetByID(ctx context.Context, id uint32) (*bizmodel.Strategy, error) {

	return nil, nil
}

func (s *strategyRepositoryImpl) FindByPage(ctx context.Context, params *bo.QueryStrategyListParams) (*bizmodel.Strategy, error) {

	return nil, nil
}
