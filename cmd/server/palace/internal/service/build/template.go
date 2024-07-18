package build

import (
	"context"

	"github.com/aide-family/moon/api"
	"github.com/aide-family/moon/api/admin"
	strategyapi "github.com/aide-family/moon/api/admin/strategy"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-family/moon/cmd/server/palace/internal/data/runtimecache"
	"github.com/aide-family/moon/pkg/palace/model"
	"github.com/aide-family/moon/pkg/util/types"
	"github.com/aide-family/moon/pkg/vobj"
)

type ToTemplateApi interface {
	ToTemplateApi(ctx context.Context) *admin.StrategyTemplate
}

type ToCreateTemplateBO interface {
	ToCreateTemplateBO() *bo.CreateTemplateStrategyParams
}

type ToUpdateTemplateBO interface {
	ToUpdateTemplateBO() *bo.UpdateTemplateStrategyParams
}

type templateStrategyBuilder struct {
	*model.StrategyTemplate
	CreateStrategy *strategyapi.CreateTemplateStrategyRequest
	UpdateStrategy *strategyapi.UpdateTemplateStrategyRequest
}

func NewTemplateStrategyBuilder(templateStrategy *model.StrategyTemplate) ToTemplateApi {
	return &templateStrategyBuilder{
		StrategyTemplate: templateStrategy,
	}
}

func (b *templateStrategyBuilder) ToTemplateApi(ctx context.Context) *admin.StrategyTemplate {
	if types.IsNil(b) || types.IsNil(b.StrategyTemplate) {
		return nil
	}
	cache := runtimecache.GetRuntimeCache()
	return &admin.StrategyTemplate{
		Id:    b.ID,
		Alert: b.Alert,
		Expr:  b.Expr,
		Levels: types.SliceTo(b.StrategyLevelTemplates, func(item *model.StrategyLevelTemplate) *admin.StrategyLevelTemplate {
			return NewTemplateStrategyLevelBuilder(item).ToApi()
		}),
		Labels:      b.Labels.Map(),
		Annotations: b.Annotations,
		Status:      api.Status(b.Status),
		CreatedAt:   b.CreatedAt.String(),
		UpdatedAt:   b.UpdatedAt.String(),
		Remark:      b.Remark,
		Creator:     NewUserBuilder(cache.GetUser(ctx, b.CreatorID)).ToApi(),
		Categories: types.SliceTo(b.Categories, func(item *model.SysDict) *admin.Select {
			return NewDictBuild(item).ToApiSelect()
		}),
	}
}

func NewCreateTemplateRequestBuilder(createTemplateRequest *strategyapi.CreateTemplateStrategyRequest) ToCreateTemplateBO {
	return &templateStrategyBuilder{
		CreateStrategy: createTemplateRequest,
	}
}

func (b *templateStrategyBuilder) ToCreateTemplateBO() *bo.CreateTemplateStrategyParams {
	strategyLevelTemplates := make([]*bo.CreateStrategyLevelTemplate, 0, len(b.CreateStrategy.GetLevel()))
	for levelID, mutationStrategyLevelTemplate := range b.CreateStrategy.GetLevel() {
		strategyLevelTemplates = append(strategyLevelTemplates, &bo.CreateStrategyLevelTemplate{
			Duration:    &types.Duration{Duration: mutationStrategyLevelTemplate.Duration},
			Count:       mutationStrategyLevelTemplate.GetCount(),
			SustainType: vobj.Sustain(mutationStrategyLevelTemplate.GetSustainType()),
			Condition:   vobj.Condition(mutationStrategyLevelTemplate.GetCondition()),
			Threshold:   mutationStrategyLevelTemplate.GetThreshold(),
			LevelID:     levelID,
			Status:      vobj.StatusEnable,
		})
	}

	return &bo.CreateTemplateStrategyParams{
		Alert:                  b.CreateStrategy.GetAlert(),
		Expr:                   b.CreateStrategy.GetExpr(),
		Status:                 vobj.StatusEnable,
		Remark:                 b.CreateStrategy.GetRemark(),
		Labels:                 vobj.NewLabels(b.CreateStrategy.GetLabels()),
		Annotations:            b.CreateStrategy.GetAnnotations(),
		StrategyLevelTemplates: strategyLevelTemplates,
		CategoriesIDs:          b.CreateStrategy.GetCategoriesIds(),
	}
}

func NewUpdateTemplateRequestBuilder(updateStrategy *strategyapi.UpdateTemplateStrategyRequest) ToUpdateTemplateBO {
	return &templateStrategyBuilder{
		UpdateStrategy: updateStrategy,
	}
}

func (b *templateStrategyBuilder) ToUpdateTemplateBO() *bo.UpdateTemplateStrategyParams {
	strategyLevelTemplates := make([]*bo.CreateStrategyLevelTemplate, 0, len(b.UpdateStrategy.GetLevel()))
	for levelID, mutationStrategyLevelTemplate := range b.UpdateStrategy.GetLevel() {
		strategyLevelTemplates = append(strategyLevelTemplates, &bo.CreateStrategyLevelTemplate{
			StrategyTemplateID: b.UpdateStrategy.GetId(),
			Duration:           &types.Duration{Duration: mutationStrategyLevelTemplate.Duration},
			Count:              mutationStrategyLevelTemplate.GetCount(),
			SustainType:        vobj.Sustain(mutationStrategyLevelTemplate.GetSustainType()),
			Condition:          vobj.Condition(mutationStrategyLevelTemplate.GetCondition()),
			Threshold:          mutationStrategyLevelTemplate.GetThreshold(),
			LevelID:            levelID,
			Status:             vobj.StatusEnable,
		})
	}
	return &bo.UpdateTemplateStrategyParams{
		Data: bo.CreateTemplateStrategyParams{
			Alert:                  b.UpdateStrategy.GetAlert(),
			Expr:                   b.UpdateStrategy.GetExpr(),
			Status:                 vobj.StatusEnable,
			Remark:                 b.UpdateStrategy.GetRemark(),
			Labels:                 vobj.NewLabels(b.UpdateStrategy.GetLabels()),
			Annotations:            b.UpdateStrategy.GetAnnotations(),
			StrategyLevelTemplates: strategyLevelTemplates,
		},
		ID: b.UpdateStrategy.Id,
	}
}

type TemplateStrategyLevelBuilder struct {
	*model.StrategyLevelTemplate
}

func NewTemplateStrategyLevelBuilder(level *model.StrategyLevelTemplate) *TemplateStrategyLevelBuilder {
	return &TemplateStrategyLevelBuilder{
		StrategyLevelTemplate: level,
	}
}

func (b *TemplateStrategyLevelBuilder) ToApi() *admin.StrategyLevelTemplate {
	if types.IsNil(b) || types.IsNil(b.StrategyLevelTemplate) {
		return nil
	}
	return &admin.StrategyLevelTemplate{
		Id:          b.ID,
		Duration:    b.Duration.GetDuration(),
		Count:       b.Count,
		SustainType: api.SustainType(b.SustainType),
		Status:      api.Status(b.Status),
		LevelId:     b.LevelID,
		Level:       NewDictBuild(b.Level).ToApiSelect(),
		Threshold:   b.Threshold,
		StrategyId:  b.StrategyTemplateID,
		Condition:   api.Condition(b.Condition),
	}
}
