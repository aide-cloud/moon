package repoimpl

import (
	"context"

	"github.com/aide-family/moon/api/merr"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/repository"
	"github.com/aide-family/moon/cmd/server/palace/internal/data"
	"github.com/aide-family/moon/pkg/helper/middleware"
	"github.com/aide-family/moon/pkg/palace/imodel"
	"github.com/aide-family/moon/pkg/palace/model"
	"github.com/aide-family/moon/pkg/palace/model/bizmodel"
	"github.com/aide-family/moon/pkg/palace/model/bizmodel/bizquery"
	"github.com/aide-family/moon/pkg/palace/model/query"
	"github.com/aide-family/moon/pkg/util/types"
	"gorm.io/gen/field"

	"gorm.io/gen"
)

func NewDictRepository(data *data.Data) repository.Dict {
	return &dictRepositoryImpl{
		data: data,
	}
}

type dictRepositoryImpl struct {
	data *data.Data
}

func (l *dictRepositoryImpl) UpdateStatusByIds(ctx context.Context, params *bo.UpdateDictStatusParams) error {
	ids := params.IDs
	status := params.Status
	if middleware.GetSourceType(ctx).IsTeam() {
		bizDB, err := getBizDB(ctx, l.data)
		if !types.IsNil(err) {
			return err
		}
		if _, err = bizDB.SysDict.WithContext(ctx).Where(bizDB.SysDict.ID.In(ids...)).Update(bizDB.SysDict.Status, params.Status); !types.IsNil(err) {
			return err
		}
	}
	if _, err := query.Use(l.data.GetMainDB(ctx)).WithContext(ctx).SysDict.Where(query.SysDict.ID.In(ids...)).Update(query.SysDict.Status, status); !types.IsNil(err) {
		return err
	}
	return nil
}

func (l *dictRepositoryImpl) DeleteByID(ctx context.Context, id uint32) error {
	if middleware.GetSourceType(ctx).IsTeam() {
		bizDB, err := getBizDB(ctx, l.data)
		if !types.IsNil(err) {
			return err
		}
		_, err = bizDB.SysDict.Where(bizDB.SysDict.ID.Eq(id)).Delete()
		if !types.IsNil(err) {
			return err
		}
	}
	if _, err := query.Use(l.data.GetMainDB(ctx)).WithContext(ctx).SysDict.Where(query.SysDict.ID.Eq(id)).Delete(); !types.IsNil(err) {
		return err
	}
	return nil
}

func (l *dictRepositoryImpl) Create(ctx context.Context, dict *bo.CreateDictParams) (imodel.IDict, error) {
	if middleware.GetSourceType(ctx).IsTeam() {
		// Team  creation
		return l.createBizDictModel(ctx, dict)
	}
	// system creation
	dictModel := createDictParamsToModel(ctx, dict)
	if types.IsNil(dictModel) {
		return nil, merr.ErrorI18nDictCreateParamCannotEmpty(ctx)
	}
	q := query.Use(l.data.GetMainDB(ctx)).WithContext(ctx).SysDict
	if err := q.Create(dictModel); !types.IsNil(err) {
		return nil, err
	}
	return dictModel, nil
}

func (l *dictRepositoryImpl) FindByPage(ctx context.Context, params *bo.QueryDictListParams) ([]imodel.IDict, error) {
	if middleware.GetSourceType(ctx).IsTeam() {
		return l.listBizDictModel(ctx, params)
	}
	return l.listDictModel(ctx, params)
}

func (l *dictRepositoryImpl) GetByID(ctx context.Context, id uint32) (imodel.IDict, error) {
	if middleware.GetSourceType(ctx).IsTeam() {
		bizDB, err := getBizDB(ctx, l.data)
		if !types.IsNil(err) {
			return nil, err
		}
		bizWrapper := bizDB.SysDict.WithContext(ctx)
		return bizWrapper.Where(bizDB.SysDict.ID.Eq(id)).Preload(field.Associations).First()
	}
	return query.Use(l.data.GetMainDB(ctx)).SysDict.WithContext(ctx).Where(query.SysDict.ID.Eq(id)).First()
}

func (l *dictRepositoryImpl) UpdateByID(ctx context.Context, dict *bo.UpdateDictParams) error {
	if middleware.GetSourceType(ctx).IsTeam() {
		return l.updateBizDictModel(ctx, dict)
	} else {
		return l.updateDictModel(ctx, dict)
	}
}

func (l *dictRepositoryImpl) listDictModel(ctx context.Context, params *bo.QueryDictListParams) ([]imodel.IDict, error) {
	dict := query.Use(l.data.GetMainDB(ctx)).SysDict
	queryWrapper := dict.WithContext(ctx)

	var wheres []gen.Condition
	if !params.Status.IsUnknown() {
		wheres = append(wheres, query.SysDict.Status.Eq(params.Status.GetValue()))
	}

	if !params.DictType.IsUnknown() {
		wheres = append(wheres, query.SysDict.DictType.Eq(params.DictType.GetValue()))
	}

	if !types.TextIsNull(params.Keyword) {
		queryWrapper = queryWrapper.Or(
			query.SysDict.Name.Like(params.Keyword),
			query.SysDict.Value.Like(params.Keyword),
			query.SysDict.Remark.Like(params.Keyword),
		)
	}
	queryWrapper = queryWrapper.Where(wheres...)
	if err := types.WithPageQuery[query.ISysDictDo](queryWrapper, params.Page); err != nil {
		return nil, err
	}
	dbDicts, err := queryWrapper.Order(query.SysDict.ID.Desc()).Find()
	if !types.IsNil(err) {
		return nil, err
	}
	dicts := types.SliceTo(dbDicts, func(dict *model.SysDict) imodel.IDict {
		return dict
	})
	return dicts, nil
}

func (l *dictRepositoryImpl) listBizDictModel(ctx context.Context, params *bo.QueryDictListParams) ([]imodel.IDict, error) {
	bizDB, err := getBizDB(ctx, l.data)
	if !types.IsNil(err) {
		return nil, err
	}
	bizWrapper := bizDB.SysDict.WithContext(ctx)

	var wheres []gen.Condition

	if !params.Status.IsUnknown() {
		wheres = append(wheres, bizDB.SysDict.Status.Eq(params.Status.GetValue()))
	}
	if !types.TextIsNull(params.Keyword) {
		bizWrapper = bizWrapper.Or(bizDB.SysDict.Name.Like(params.Keyword))
		bizWrapper = bizWrapper.Or(bizDB.SysDict.Value.Like(params.Keyword))
		bizWrapper = bizWrapper.Or(bizDB.SysDict.Remark.Like(params.Keyword))
	}

	bizWrapper = bizWrapper.Where(wheres...).Preload(field.Associations)

	if err := types.WithPageQuery[bizquery.ISysDictDo](bizWrapper, params.Page); err != nil {
		return nil, err
	}
	sysDicts, err := bizWrapper.Order(bizDB.SysDict.ID.Desc()).Find()
	if !types.IsNil(err) {
		return nil, err
	}
	dicts := types.SliceTo(sysDicts, func(dict *bizmodel.SysDict) imodel.IDict {
		return dict
	})
	return dicts, nil
}

// createBizDictModel create team dict model
func (l *dictRepositoryImpl) createBizDictModel(ctx context.Context, dict *bo.CreateDictParams) (*bizmodel.SysDict, error) {
	bizDB, err := getBizDB(ctx, l.data)
	if !types.IsNil(err) {
		return nil, err
	}
	dictBizModel := createBizDictParamsToModel(ctx, dict)
	if err := bizDB.SysDict.WithContext(ctx).Create(dictBizModel); !types.IsNil(err) {
		return nil, err
	}
	return dictBizModel, nil
}

func (l *dictRepositoryImpl) updateDictModel(ctx context.Context, params *bo.UpdateDictParams) error {
	id := params.ID
	updateParam := params.UpdateParam
	_, err := query.Use(l.data.GetMainDB(ctx)).SysDict.WithContext(ctx).Where(query.SysDict.ID.Eq(id)).UpdateSimple(
		query.SysDict.Name.Value(updateParam.Name),
		query.SysDict.Value.Value(updateParam.Value),
		query.SysDict.CssClass.Value(updateParam.CssClass),
		query.SysDict.ColorType.Value(updateParam.ColorType),
		query.SysDict.Remark.Value(updateParam.Remark),
		query.SysDict.ImageUrl.Value(updateParam.ImageUrl),
		query.SysDict.Icon.Value(updateParam.Icon),
	)
	return err
}

func (l *dictRepositoryImpl) updateBizDictModel(ctx context.Context, params *bo.UpdateDictParams) error {
	bizDB, err := getBizDB(ctx, l.data)
	if !types.IsNil(err) {
		return err
	}
	updateParam := params.UpdateParam
	id := params.ID
	_, err = bizDB.SysDict.Where(bizDB.SysDict.ID.Eq(id)).UpdateSimple(
		bizDB.SysDict.Name.Value(updateParam.Name),
		bizDB.SysDict.Remark.Value(updateParam.Remark),
		bizDB.SysDict.Value.Value(updateParam.Value),
		bizDB.SysDict.CssClass.Value(updateParam.CssClass),
		bizDB.SysDict.ColorType.Value(updateParam.ColorType),
		bizDB.SysDict.ImageUrl.Value(updateParam.ImageUrl),
		bizDB.SysDict.Icon.Value(updateParam.Icon),
	)
	return err
}

func createBizDictParamsToModel(ctx context.Context, dict *bo.CreateDictParams) *bizmodel.SysDict {
	if types.IsNil(dict) {
		return nil
	}
	modelDict := &bizmodel.SysDict{
		Name:         dict.Name,
		Value:        dict.Value,
		DictType:     dict.DictType,
		ColorType:    dict.ColorType,
		CssClass:     dict.CssClass,
		Status:       dict.Status,
		Remark:       dict.Remark,
		Icon:         dict.Icon,
		ImageUrl:     dict.ImageUrl,
		LanguageCode: dict.LanguageCode,
	}
	modelDict.WithContext(ctx)
	return modelDict
}

// createDictParamsToModel create dict params to model
func createDictParamsToModel(ctx context.Context, dict *bo.CreateDictParams) *model.SysDict {
	if types.IsNil(dict) {
		return nil
	}
	dictModel := &model.SysDict{
		Name:         dict.Name,
		Value:        dict.Value,
		DictType:     dict.DictType,
		ColorType:    dict.ColorType,
		CssClass:     dict.CssClass,
		Status:       dict.Status,
		Remark:       dict.Remark,
		Icon:         dict.Icon,
		ImageUrl:     dict.ImageUrl,
		LanguageCode: dict.LanguageCode,
	}
	dictModel.WithContext(ctx)
	return dictModel
}
