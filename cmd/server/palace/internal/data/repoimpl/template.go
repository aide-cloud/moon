package repoimpl

import (
	"context"

	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/repository"
	"github.com/aide-family/moon/cmd/server/palace/internal/data"
	"github.com/aide-family/moon/pkg/palace/model"
	"github.com/aide-family/moon/pkg/palace/model/query"
	"github.com/aide-family/moon/pkg/util/types"
	"github.com/aide-family/moon/pkg/vobj"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func NewTemplateRepository(data *data.Data) repository.Template {
	return &templateRepositoryImpl{data: data}
}

type templateRepositoryImpl struct {
	data *data.Data
}

func (l *templateRepositoryImpl) CreateTemplateStrategy(ctx context.Context, createParam *bo.CreateTemplateStrategyParams) error {
	return query.Use(l.data.GetMainDB(ctx)).Transaction(func(tx *query.Query) error {
		templateStrategy := createTemplateStrategy(createParam)
		templateStrategy.WithContext(ctx)
		// 创建主数据
		if err := tx.StrategyTemplate.WithContext(ctx).Create(templateStrategy); err != nil {
			return err
		}
		StrategyTemplateID := templateStrategy.ID
		strategyLevelTemplates := types.SliceTo(createParam.StrategyLevelTemplates, func(item *model.StrategyLevelTemplate) *model.StrategyLevelTemplate {
			item.StrategyTemplateID = StrategyTemplateID
			return item
		})
		// 创建子数据
		if err := tx.StrategyLevelTemplate.WithContext(ctx).Create(strategyLevelTemplates...); err != nil {
			return err
		}
		return nil
	})
}

func (l *templateRepositoryImpl) UpdateTemplateStrategy(ctx context.Context, updateParam *bo.UpdateTemplateStrategyParams) error {
	return query.Use(l.data.GetMainDB(ctx)).Transaction(func(tx *query.Query) error {
		// 删除全部关联模板等级数据
		if _, err := tx.StrategyLevelTemplate.WithContext(ctx).Where(query.StrategyLevelTemplate.StrategyID.Eq(updateParam.ID)).Delete(); err != nil {
			return err
		}

		strategyLevelTemplates := types.SliceTo(updateParam.Data.StrategyLevelTemplates, func(item *model.StrategyLevelTemplate) *model.StrategyLevelTemplate {
			return &model.StrategyLevelTemplate{
				StrategyTemplateID: updateParam.ID,
				Duration:           item.Duration,
				LevelID:            item.LevelID,
				Count:              item.Count,
				SustainType:        item.SustainType,
				Interval:           item.Interval,
				Condition:          item.Condition,
				Threshold:          item.Threshold,
				Status:             item.Status,
			}
		})

		if err := tx.StrategyLevelTemplate.WithContext(ctx).Clauses(clause.OnConflict{UpdateAll: true}).Create(strategyLevelTemplates...); err != nil {
			return err
		}
		_, err := tx.StrategyTemplate.WithContext(ctx).
			Where(query.StrategyTemplate.ID.Eq(updateParam.ID)).
			UpdateSimple(
				query.StrategyTemplate.Expr.Value(updateParam.Data.Expr),
				query.StrategyTemplate.Remark.Value(updateParam.Data.Remark),
				query.StrategyTemplate.Labels.Value(updateParam.Data.Labels),
				query.StrategyTemplate.Annotations.Value(updateParam.Data.Annotations),
				query.StrategyTemplate.Alert.Value(updateParam.Data.Alert),
				query.StrategyTemplate.Status.Value(updateParam.Data.Status.GetValue()),
			)
		return err
	})
}

func (l *templateRepositoryImpl) DeleteTemplateStrategy(ctx context.Context, id uint32) error {
	return query.Use(l.data.GetMainDB(ctx)).Transaction(func(tx *query.Query) error {
		// 删除关联数据
		if _, err := tx.StrategyLevelTemplate.WithContext(ctx).Where(query.StrategyLevelTemplate.StrategyID.Eq(id)).Delete(); err != nil {
			return err
		}
		// 删除策略
		_, err := query.Use(l.data.GetMainDB(ctx)).WithContext(ctx).
			StrategyTemplate.
			Where(query.StrategyTemplate.ID.Eq(id)).
			Delete()
		return err
	})
}

func (l *templateRepositoryImpl) GetTemplateStrategy(ctx context.Context, id uint32) (*model.StrategyTemplate, error) {
	q := query.Use(l.data.GetMainDB(ctx)).WithContext(ctx).StrategyTemplate
	return q.Preload(field.Associations).
		Where(query.StrategyTemplate.ID.Eq(id)).
		First()
}

func (l *templateRepositoryImpl) ListTemplateStrategy(ctx context.Context, params *bo.QueryTemplateStrategyListParams) ([]*model.StrategyTemplate, error) {
	queryWrapper := query.Use(l.data.GetMainDB(ctx)).StrategyTemplate.WithContext(ctx)

	var wheres []gen.Condition
	if !types.TextIsNull(params.Alert) {
		wheres = append(wheres, query.StrategyTemplate.Alert.Like(params.Alert))
	}
	if !params.Status.IsUnknown() {
		wheres = append(wheres, query.StrategyTemplate.Status.Eq(params.Status.GetValue()))
	}

	if !types.TextIsNull(params.Keyword) {
		queryWrapper = queryWrapper.Or(query.StrategyTemplate.Alert.Like(params.Keyword))
		queryWrapper = queryWrapper.Or(query.StrategyTemplate.Expr.Like(params.Keyword))
		queryWrapper = queryWrapper.Or(query.StrategyTemplate.Remark.Like(params.Keyword))
	}

	queryWrapper = queryWrapper.Where(wheres...)
	if err := types.WithPageQuery[query.IStrategyTemplateDo](queryWrapper, params.Page); err != nil {
		return nil, err
	}
	return queryWrapper.Order(query.StrategyTemplate.ID).Find()
}

func (l *templateRepositoryImpl) UpdateTemplateStrategyStatus(ctx context.Context, status vobj.Status, ids ...uint32) error {
	_, err := query.Use(l.data.GetMainDB(ctx)).WithContext(ctx).
		StrategyTemplate.
		Where(query.StrategyTemplate.ID.In(ids...)).
		UpdateSimple(query.StrategyTemplate.Status.Value(status.GetValue()))
	return err
}

func createTemplateStrategy(createParam *bo.CreateTemplateStrategyParams) *model.StrategyTemplate {
	return &model.StrategyTemplate{
		Alert:       createParam.Alert,
		Expr:        createParam.Expr,
		Status:      createParam.Status,
		Remark:      createParam.Remark,
		Labels:      createParam.Labels,
		Annotations: createParam.Annotations,
	}
}

func transitionStrategyLevelTemplate(params []*bo.CreateStrategyLevelTemplate, StrategyID uint32) []*model.StrategyLevelTemplate {
	return types.SliceToWithFilter(params, func(item *bo.CreateStrategyLevelTemplate) (*model.StrategyLevelTemplate, bool) {
		if types.IsNil(item) || types.TextIsNull(item.Condition) {
			return nil, false
		}
		return &model.StrategyLevelTemplate{
			StrategyTemplateID: StrategyID,
			Duration:           item.Duration,
			Count:              item.Count,
			SustainType:        item.SustainType,
			Interval:           item.Interval,
			Condition:          item.Condition,
			Threshold:          item.Threshold,
			LevelID:            item.LevelID,
			Status:             item.Status,
		}, true
	})
}
