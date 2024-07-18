package build

import (
	"context"
	strategyapi "github.com/aide-family/moon/api/admin/strategy"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-family/moon/pkg/vobj"

	"github.com/aide-family/moon/api"
	"github.com/aide-family/moon/api/admin"
	"github.com/aide-family/moon/pkg/palace/model/bizmodel"
	"github.com/aide-family/moon/pkg/util/types"
)

type (
	StrategyBuilder interface {
		ToApi(ctx context.Context) *admin.Strategy
		ToCreateStrategyBO(teamID uint32) *bo.CreateStrategyParams
		ToUpdateStrategyBO(teamID uint32) *bo.UpdateStrategyParams
	}

	strategyBuilder struct {
		*bizmodel.Strategy
		CreateStrategy *strategyapi.CreateStrategyRequest
		UpdateStrategy *strategyapi.UpdateStrategyRequest
		ctx            context.Context
	}
)

// ToApi 转换为API层数据
func (b *strategyBuilder) ToApi(ctx context.Context) *admin.Strategy {
	if types.IsNil(b) || types.IsNil(b.Strategy) {
		return nil
	}
	strategyLevels := types.SliceToWithFilter(b.StrategyLevel, func(level *bizmodel.StrategyLevel) (*admin.StrategyLevel, bool) {
		return NewStrategyLevelBuilder(level).ToApi(ctx), true
	})

	return &admin.Strategy{
		Name:        b.Name,
		Id:          b.ID,
		Expr:        b.Expr,
		Labels:      b.Labels.Map(),
		Annotations: b.Annotations,
		Datasource: types.SliceTo(b.Datasource, func(datasource *bizmodel.Datasource) *admin.Datasource {
			return NewBuilder().WithContext(ctx).WithDoDatasource(datasource).ToApi()
		}),
		StrategyTemplateId: b.StrategyTemplateID,
		Levels:             strategyLevels,
		Status:             api.Status(b.Status),
		Step:               b.Step,
		SourceType:         api.TemplateSourceType(b.StrategyTemplateSource),
	}
}

func (b *strategyBuilder) ToCreateStrategyBO(teamID uint32) *bo.CreateStrategyParams {
	strategyLevels := make([]*bo.CreateStrategyLevel, 0, len(b.CreateStrategy.GetStrategyLevel()))
	for _, strategyLevel := range b.CreateStrategy.GetStrategyLevel() {
		strategyLevels = append(strategyLevels, &bo.CreateStrategyLevel{
			StrategyTemplateID: b.CreateStrategy.TemplateId,
			Count:              strategyLevel.GetCount(),
			Duration:           types.NewDuration(strategyLevel.GetDuration()),
			SustainType:        vobj.Sustain(strategyLevel.SustainType),
			Interval:           types.NewDuration(strategyLevel.GetInterval()),
			Condition:          vobj.Condition(strategyLevel.GetCondition()),
			Threshold:          strategyLevel.GetThreshold(),
			Status:             vobj.Status(strategyLevel.GetStatus()),
			LevelID:            strategyLevel.GetLevelId(),
		})
	}
	return &bo.CreateStrategyParams{
		TeamID:        teamID,
		TemplateId:    b.CreateStrategy.GetTemplateId(),
		GroupId:       b.CreateStrategy.GetGroupId(),
		Name:          b.CreateStrategy.GetName(),
		Remark:        b.CreateStrategy.GetRemark(),
		Status:        vobj.Status(b.CreateStrategy.GetStatus()),
		Step:          b.CreateStrategy.GetStep(),
		SourceType:    vobj.TemplateSourceType(b.CreateStrategy.GetSourceType()),
		DatasourceIds: b.CreateStrategy.GetDatasourceIds(),
		StrategyLevel: strategyLevels,
	}
}

func (b *strategyBuilder) ToUpdateStrategyBO(teamID uint32) *bo.UpdateStrategyParams {

	strategyLevels := make([]*bo.CreateStrategyLevel, 0, len(b.UpdateStrategy.GetData().GetStrategyLevel()))
	for _, strategyLevel := range b.UpdateStrategy.GetData().GetStrategyLevel() {
		strategyLevels = append(strategyLevels, &bo.CreateStrategyLevel{
			StrategyTemplateID: b.UpdateStrategy.GetData().TemplateId,
			Count:              strategyLevel.GetCount(),
			Duration:           types.NewDuration(strategyLevel.GetDuration()),
			SustainType:        vobj.Sustain(strategyLevel.SustainType),
			Interval:           types.NewDuration(strategyLevel.GetInterval()),
			Condition:          vobj.Condition(strategyLevel.GetCondition()),
			Threshold:          strategyLevel.GetThreshold(),
			Status:             vobj.Status(strategyLevel.GetStatus()),
			LevelID:            strategyLevel.GetLevelId(),
		})
	}
	return &bo.UpdateStrategyParams{
		TeamID: teamID,
		ID:     b.UpdateStrategy.GetId(),
		UpdateParam: bo.CreateStrategyParams{
			TemplateId:    b.UpdateStrategy.GetData().GetTemplateId(),
			GroupId:       b.UpdateStrategy.GetData().GetGroupId(),
			Name:          b.UpdateStrategy.GetData().GetName(),
			Remark:        b.UpdateStrategy.GetData().GetRemark(),
			Status:        vobj.Status(b.UpdateStrategy.GetData().GetStatus()),
			Step:          b.UpdateStrategy.GetData().GetStep(),
			SourceType:    vobj.TemplateSourceType(b.UpdateStrategy.GetData().GetSourceType()),
			DatasourceIds: b.UpdateStrategy.GetData().GetDatasourceIds(),
			StrategyLevel: strategyLevels,
		},
	}
}
