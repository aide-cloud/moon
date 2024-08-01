package build

import (
	"context"

	"github.com/aide-family/moon/api"
	adminapi "github.com/aide-family/moon/api/admin"
	"github.com/aide-family/moon/pkg/palace/model/bizmodel"
	"github.com/aide-family/moon/pkg/util/types"
)

type (
	RealtimeAlarmBuilder interface {
		ToAPI() *adminapi.RealtimeAlarmItem
	}

	realtimeAlarmBuilder struct {
		alarm *bizmodel.RealtimeAlarm
		ctx   context.Context
	}
)

func (r *realtimeAlarmBuilder) ToAPI() *adminapi.RealtimeAlarmItem {
	if types.IsNil(r) || types.IsNil(r.alarm) {
		return nil
	}
	detail := r.alarm
	return &adminapi.RealtimeAlarmItem{
		Id:           detail.ID,
		StartsAt:     types.NewTimeByUnix(detail.StartsAt).String(),
		EndsAt:       types.NewTimeByUnix(detail.EndsAt).String(),
		Status:       api.AlertStatus(detail.Status),
		Level:        NewBuilder().WithDict(detail.Level).ToAPISelect(),
		LevelID:      detail.LevelID,
		StrategyID:   detail.StrategyID,
		Strategy:     NewBuilder().WithContext(r.ctx).WithAPIStrategy(detail.Strategy).ToAPI(),
		Summary:      detail.Summary,
		Description:  detail.Description,
		Expr:         detail.Expr,
		DatasourceID: detail.DatasourceID,
		Datasource:   NewBuilder().WithDoDatasource(detail.Datasource).ToAPI(),
		Fingerprint:  detail.Fingerprint,
	}
}

func NewRealtimeAlarmBuilder(ctx context.Context, alarm *bizmodel.RealtimeAlarm) RealtimeAlarmBuilder {
	return &realtimeAlarmBuilder{alarm: alarm, ctx: ctx}
}
