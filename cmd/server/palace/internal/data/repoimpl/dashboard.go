package repoimpl

import (
	"context"

	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/repository"
	"github.com/aide-family/moon/cmd/server/palace/internal/data"
	"github.com/aide-family/moon/pkg/palace/model/bizmodel"
)

// NewDashboardRepository 创建仪表盘操作实现
func NewDashboardRepository(data *data.Data) repository.Dashboard {
	return &dashboardRepositoryImpl{data: data}
}

type dashboardRepositoryImpl struct {
	data *data.Data
}

func (d *dashboardRepositoryImpl) AddDashboard(ctx context.Context, req *bo.AddDashboardParams) error {
	//TODO implement me
	panic("implement me")
}

func (d *dashboardRepositoryImpl) DeleteDashboard(ctx context.Context, req *bo.DeleteDashboardParams) error {
	//TODO implement me
	panic("implement me")
}

func (d *dashboardRepositoryImpl) UpdateDashboard(ctx context.Context, req *bo.UpdateDashboardParams) error {
	//TODO implement me
	panic("implement me")
}

func (d *dashboardRepositoryImpl) GetDashboard(ctx context.Context, id uint32) (*bizmodel.Dashboard, error) {
	//TODO implement me
	panic("implement me")
}

func (d *dashboardRepositoryImpl) ListDashboard(ctx context.Context, params *bo.ListDashboardParams) ([]*bizmodel.Dashboard, error) {
	//TODO implement me
	panic("implement me")
}
