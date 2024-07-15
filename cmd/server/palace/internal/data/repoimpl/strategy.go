package repoimpl

import (
	"context"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/repository"
	"github.com/aide-family/moon/cmd/server/palace/internal/data"
	"github.com/aide-family/moon/pkg/palace/model"
	"github.com/aide-family/moon/pkg/palace/model/bizmodel"
	"github.com/aide-family/moon/pkg/palace/model/bizmodel/bizquery"
	"github.com/aide-family/moon/pkg/palace/model/query"
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

func (s *strategyRepositoryImpl) UpdateStatus(ctx context.Context, params *bo.UpdateStrategyStatusParams) error {
	bizDB, err := s.data.GetBizGormDB(params.TeamID)
	if err != nil {
		return err
	}
	queryWrapper := bizquery.Use(bizDB)
	_, err = queryWrapper.WithContext(ctx).Strategy.Where(queryWrapper.Strategy.ID.In(params.Ids...)).Update(queryWrapper.Strategy.Status, params.Status)
	if err != nil {
		return err
	}
	return nil
}

func (s *strategyRepositoryImpl) DeleteByID(ctx context.Context, params *bo.DelStrategyParams) error {
	bizDB, err := s.data.GetBizGormDB(params.TeamID)
	if err != nil {
		return err
	}
	queryWrapper := bizquery.Use(bizDB)
	_, err = queryWrapper.Strategy.WithContext(ctx).Where(queryWrapper.Strategy.ID.Eq(params.ID)).Delete()
	if !types.IsNil(err) {
		return err
	}
	_, err = queryWrapper.StrategyLevel.WithContext(ctx).Where(queryWrapper.StrategyLevel.StrategyID.Eq(params.ID)).Delete()
	if !types.IsNil(err) {
		return err
	}
	return nil
}

func (s *strategyRepositoryImpl) CreateStrategy(ctx context.Context, params *bo.CreateStrategyParams) (*bizmodel.Strategy, error) {
	bizDB, err := s.data.GetBizGormDB(params.TeamID)
	templateId := params.TemplateId

	if !types.IsNil(err) {
		return nil, err
	}

	mainDb := s.data.GetMainDB(ctx).WithContext(ctx)

	strategyTemplate, err := query.Use(mainDb).StrategyTemplate.Where(query.StrategyTemplate.ID.Eq(templateId)).Preload(field.Associations).First()
	if !types.IsNil(err) {
		return nil, err
	}

	strategyModel := &bizmodel.Strategy{
		Name:                   params.Name,
		StrategyTemplateID:     templateId,
		StrategyTemplateSource: vobj.StrategyTemplateSource(params.SourceType),
		Alert:                  strategyTemplate.Alert,
		Expr:                   strategyTemplate.Expr,
		Labels:                 strategyTemplate.Labels,
		Annotations:            strategyTemplate.Annotations,
		Remark:                 params.Remark,
		Step:                   params.Step,
		Datasource: types.SliceToWithFilter(params.DatasourceIds, func(datasourceId uint32) (*bizmodel.Datasource, bool) {
			if datasourceId <= 0 {
				return nil, false
			}
			return &bizmodel.Datasource{
				AllFieldModel: model.AllFieldModel{ID: datasourceId},
			}, true
		}),
		Categories: types.SliceToWithFilter(strategyTemplate.Categories, func(dict *model.SysDict) (*bizmodel.SysDict, bool) {
			if dict.ID <= 0 {
				return nil, false
			}
			return &bizmodel.SysDict{
				AllFieldModel: model.AllFieldModel{ID: dict.ID},
			}, true
		}),
	}

	err = bizquery.Use(bizDB).Transaction(func(tx *bizquery.Query) error {
		if err := tx.Strategy.WithContext(ctx).Create(strategyModel); !types.IsNil(err) {
			return err
		}
		return nil
	})
	if !types.IsNil(err) {
		return nil, err
	}

	// Creating  Strategy levels
	strategyLevels := createStrategyLevelParamsToModel(ctx, params.StrategyLevel, strategyModel.ID)
	if err := bizquery.Use(bizDB).StrategyLevel.WithContext(ctx).Create(strategyLevels...); !types.IsNil(err) {
		return nil, err
	}
	return strategyModel, nil
}

func (s *strategyRepositoryImpl) UpdateByID(ctx context.Context, params *bo.UpdateStrategyParams) error {
	bizDB, err := s.data.GetBizGormDB(params.TeamID)
	if err != nil {
		return err
	}
	updateParam := params.UpdateParam

	queryWrapper := bizquery.Use(bizDB)

	datasourceIds := types.SliceToWithFilter(updateParam.DatasourceIds, func(datasourceId uint32) (*bizmodel.Datasource, bool) {
		if datasourceId <= 0 {
			return nil, false
		}
		return &bizmodel.Datasource{
			AllFieldModel: model.AllFieldModel{ID: datasourceId},
		}, true
	})

	err = queryWrapper.Strategy.Datasource.
		Model(&bizmodel.Strategy{AllFieldModel: model.AllFieldModel{ID: params.ID}}).Replace(datasourceIds...)

	if !types.IsNil(err) {
		return err
	}

	strategyTemplate, err := queryWrapper.StrategyTemplate.Where(query.StrategyTemplate.ID.Eq(params.UpdateParam.TemplateId)).Preload(field.Associations).First()

	if strategyTemplate != nil {
		categories := types.SliceToWithFilter(strategyTemplate.Categories, func(dict *bizmodel.SysDict) (*bizmodel.SysDict, bool) {
			if dict.ID <= 0 {
				return nil, false
			}
			return &bizmodel.SysDict{
				AllFieldModel: model.AllFieldModel{ID: dict.ID},
			}, true
		})

		err = queryWrapper.Strategy.Categories.Model(&bizmodel.Strategy{AllFieldModel: model.AllFieldModel{ID: params.ID}}).Replace(categories...)
		if !types.IsNil(err) {
			return err
		}
	}
	// 删除策略等级数据
	_, err = queryWrapper.StrategyLevel.WithContext(ctx).Where(queryWrapper.StrategyLevel.StrategyID.Eq(params.ID)).Delete()
	if !types.IsNil(err) {
		return err
	}
	// Creating  Strategy levels
	strategyLevels := createStrategyLevelParamsToModel(ctx, updateParam.StrategyLevel, params.ID)
	err = bizquery.Use(bizDB).StrategyLevel.WithContext(ctx).Create(strategyLevels...)
	if !types.IsNil(err) {
		return err
	}

	return queryWrapper.Transaction(func(tx *bizquery.Query) error {
		// 更新策略
		_, err = tx.Strategy.WithContext(ctx).Where(queryWrapper.Strategy.ID.Eq(params.ID)).UpdateSimple(
			queryWrapper.Strategy.Name.Value(updateParam.Name),
			queryWrapper.Strategy.Name.Value(updateParam.Name),
			queryWrapper.Strategy.Step.Value(updateParam.Step),
			queryWrapper.Strategy.Remark.Value(updateParam.Remark),
		)
		if !types.IsNil(err) {
			return err
		}
		return nil
	})
}

func (s *strategyRepositoryImpl) GetByID(ctx context.Context, params *bo.GetStrategyDetailParams) (*bizmodel.Strategy, error) {
	bizDB, err := s.data.GetBizGormDB(params.TeamID)
	if err != nil {
		return nil, err
	}
	bizWrapper := bizquery.Use(bizDB).Strategy.WithContext(ctx)
	return bizWrapper.Where(bizquery.Use(bizDB).Strategy.ID.Eq(params.ID)).Preload(field.Associations).First()
}

func (s *strategyRepositoryImpl) FindByPage(ctx context.Context, params *bo.QueryStrategyListParams) ([]*bizmodel.Strategy, error) {
	bizDB, err := s.data.GetBizGormDB(params.TeamID)
	if err != nil {
		return nil, err
	}
	strategyWrapper := bizquery.Use(bizDB).Strategy.WithContext(ctx)

	var wheres []gen.Condition
	if !types.TextIsNull(params.Alert) {
		wheres = append(wheres, bizquery.Strategy.Alert.Like(params.Alert))
	}
	if !params.Status.IsUnknown() {
		wheres = append(wheres, bizquery.Strategy.Status.Eq(params.Status.GetValue()))
	}

	if !types.TextIsNull(params.Keyword) {
		strategyWrapper = strategyWrapper.Or(bizquery.Use(bizDB).Strategy.Name.Like(params.Keyword))
		strategyWrapper = strategyWrapper.Or(bizquery.Use(bizDB).Strategy.Remark.Like(params.Keyword))
		strategyWrapper = strategyWrapper.Or(bizquery.Use(bizDB).Strategy.Alert.Like(params.Keyword))

		dictWrapper := query.Use(s.data.GetMainDB(ctx)).SysDict.WithContext(ctx)

		dictWrapper = dictWrapper.Or(query.SysDict.Name.Like(params.Keyword))
		dictWrapper = dictWrapper.Or(query.SysDict.Value.Like(params.Keyword))
		dictWrapper = dictWrapper.Or(query.SysDict.Remark.Like(params.Keyword))

		sysDicts, err := dictWrapper.Find()
		if err != nil {
			return nil, err
		}

		categoriesIds := types.SliceTo(sysDicts, func(item *model.SysDict) uint32 {
			return item.ID
		})

		var strategyTemplateIds []uint32
		strategyTemplateCategories := query.Use(s.data.GetMainDB(ctx)).StrategyTemplateCategories.WithContext(ctx)
		_ = strategyTemplateCategories.Where(query.StrategyTemplateCategories.SysDictID.In(categoriesIds...)).
			Select(query.StrategyTemplateCategories.StrategyTemplateID).
			Scan(&strategyTemplateIds)
		if len(strategyTemplateIds) > 0 {
			strategyWrapper = strategyWrapper.Or(bizquery.Use(bizDB).Strategy.StrategyTemplateID.In(strategyTemplateIds...))
		}
	}

	strategyWrapper.Where(wheres...).Preload(field.Associations)

	if err := types.WithPageQuery[bizquery.IStrategyDo](strategyWrapper, params.Page); err != nil {
		return nil, err
	}

	return strategyWrapper.Order(bizquery.Use(bizDB).Strategy.ID.Desc()).Find()
}

func createStrategyLevelParamsToModel(ctx context.Context, params []*bo.CreateStrategyLevel, strategyId uint32) []*bizmodel.StrategyLevel {
	strategyLevel := types.SliceTo(params, func(item *bo.CreateStrategyLevel) *bizmodel.StrategyLevel {
		templateLevel := &bizmodel.StrategyLevel{
			StrategyID:  strategyId,
			Duration:    item.Duration,
			Count:       item.Count,
			SustainType: item.SustainType,
			Interval:    item.Interval,
			Condition:   item.Condition,
			Threshold:   item.Threshold,
			LevelID:     item.LevelID,
			Status:      item.Status,
		}
		templateLevel.WithContext(ctx)
		return templateLevel
	})
	return strategyLevel
}
