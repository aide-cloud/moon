package build

import (
	"context"

	"github.com/aide-family/moon/api"
	"github.com/aide-family/moon/api/admin"
	strategyapi "github.com/aide-family/moon/api/admin/strategy"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-family/moon/pkg/palace/model/bizmodel"
	"github.com/aide-family/moon/pkg/util/types"
	"github.com/aide-family/moon/pkg/vobj"
)

type ToStrategyApi interface {
	ToStrategyApi(ctx context.Context) *admin.Strategy
}

type ToCreateStrategyBO interface {
	ToCreateStrategyBO(teamID uint32) *bo.CreateStrategyParams
}

type ToUpdateStrategyBO interface {
	ToUpdateStrategyBO(teamID uint32) *bo.UpdateStrategyParams
}

type strategyBuilder struct {
	Strategy              *bizmodel.Strategy
	CreateStrategyRequest *strategyapi.CreateStrategyRequest
	UpdateStrategyRequest *strategyapi.UpdateStrategyRequest
}

func NewStrategyBuilder(strategy *bizmodel.Strategy) ToStrategyApi {
	return &strategyBuilder{
		Strategy: strategy,
	}
}

// ToStrategyApi 转换为API层数据
func (b *strategyBuilder) ToStrategyApi(ctx context.Context) *admin.Strategy {
	if types.IsNil(b) || types.IsNil(b.Strategy) {
		return nil
	}
	strategyLevels := types.SliceToWithFilter(b.Strategy.StrategyLevel, func(level *bizmodel.StrategyLevel) (*admin.StrategyLevel, bool) {
		return NewStrategyLevelBuilder(level).ToApi(ctx), true
	})

	return &admin.Strategy{
		Name:        b.Strategy.Name,
		Id:          b.Strategy.ID,
		Expr:        b.Strategy.Expr,
		Labels:      b.Strategy.Labels.Map(),
		Annotations: b.Strategy.Annotations,
		Datasource: types.SliceTo(b.Strategy.Datasource, func(datasource *bizmodel.Datasource) *admin.Datasource {
			return NewDatasourceBuilder(datasource).ToApi(ctx)
		}),
		StrategyTemplateId: b.Strategy.StrategyTemplateID,
		Levels:             strategyLevels,
		Status:             api.Status(b.Strategy.Status),
		Step:               b.Strategy.Step,
		SourceType:         api.TemplateSourceType(b.Strategy.StrategyTemplateSource),
	}
}

func NewCreateRequestBuilder(create *strategyapi.CreateStrategyRequest) ToCreateStrategyBO {
	return &strategyBuilder{
		CreateStrategyRequest: create,
	}
}

func (b *strategyBuilder) ToCreateStrategyBO(teamID uint32) *bo.CreateStrategyParams {
	strategyLevels := make([]*bo.CreateStrategyLevel, 0, len(b.CreateStrategyRequest.GetStrategyLevel()))
	for _, strategyLevel := range b.CreateStrategyRequest.GetStrategyLevel() {
		strategyLevels = append(strategyLevels, &bo.CreateStrategyLevel{
			StrategyTemplateID: b.CreateStrategyRequest.TemplateId,
			Count:              strategyLevel.GetCount(),
			Duration:           types.NewDuration(strategyLevel.GetDuration()),
			SustainType:        vobj.Sustain(strategyLevel.SustainType),
			Interval:           types.NewDuration(strategyLevel.GetInterval()),
			Condition:          vobj.Condition(strategyLevel.GetCondition()),
			Threshold:          strategyLevel.GetThreshold(),
			Status:             vobj.Status(strategyLevel.Status),
			LevelID:            strategyLevel.GetLevelId(),
		})
	}

	params := &bo.CreateStrategyParams{
		TeamID:        teamID,
		TemplateId:    b.CreateStrategyRequest.GetTemplateId(),
		GroupId:       b.CreateStrategyRequest.GetGroupId(),
		Name:          b.CreateStrategyRequest.GetName(),
		Remark:        b.CreateStrategyRequest.GetRemark(),
		Status:        vobj.Status(b.CreateStrategyRequest.GetStatus()),
		Step:          b.CreateStrategyRequest.GetStep(),
		SourceType:    vobj.TemplateSourceType(b.CreateStrategyRequest.GetSourceType()),
		DatasourceIds: b.CreateStrategyRequest.GetDatasourceIds(),
		StrategyLevel: strategyLevels,
	}
	return params
}

func NewUpdateRequestBuilder(update *strategyapi.UpdateStrategyRequest) ToUpdateStrategyBO {
	return &strategyBuilder{
		UpdateStrategyRequest: update,
	}
}

func (b *strategyBuilder) ToUpdateStrategyBO(teamID uint32) *bo.UpdateStrategyParams {
	strategyLevels := make([]*bo.CreateStrategyLevel, 0, len(b.UpdateStrategyRequest.GetData().GetStrategyLevel()))
	for _, strategyLevel := range b.UpdateStrategyRequest.GetData().GetStrategyLevel() {
		strategyLevels = append(strategyLevels, &bo.CreateStrategyLevel{
			StrategyTemplateID: b.UpdateStrategyRequest.GetData().TemplateId,
			Count:              strategyLevel.GetCount(),
			Duration:           types.NewDuration(strategyLevel.GetDuration()),
			SustainType:        vobj.Sustain(strategyLevel.SustainType),
			Interval:           types.NewDuration(strategyLevel.GetInterval()),
			Condition:          vobj.Condition(strategyLevel.GetCondition()),
			Threshold:          strategyLevel.GetThreshold(),
			Status:             vobj.Status(strategyLevel.Status),
			LevelID:            strategyLevel.GetLevelId(),
		})
	}

	return &bo.UpdateStrategyParams{
		TeamID: teamID,
		ID:     b.UpdateStrategyRequest.GetId(),
		UpdateParam: bo.CreateStrategyParams{
			TemplateId:    b.UpdateStrategyRequest.GetData().GetTemplateId(),
			GroupId:       b.UpdateStrategyRequest.GetData().GetGroupId(),
			Name:          b.UpdateStrategyRequest.GetData().GetName(),
			Remark:        b.UpdateStrategyRequest.GetData().GetRemark(),
			Status:        vobj.Status(b.UpdateStrategyRequest.GetData().GetStatus()),
			Step:          b.UpdateStrategyRequest.GetData().GetStep(),
			SourceType:    vobj.TemplateSourceType(b.UpdateStrategyRequest.GetData().GetSourceType()),
			DatasourceIds: b.UpdateStrategyRequest.GetData().GetDatasourceIds(),
			StrategyLevel: strategyLevels,
		},
	}
}
