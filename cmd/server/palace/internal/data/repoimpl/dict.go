package repoimpl

import (
	"context"

	"gorm.io/gen"

	"github.com/aide-family/moon/api/merr"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/repository"
	"github.com/aide-family/moon/cmd/server/palace/internal/data"
	"github.com/aide-family/moon/pkg/helper/middleware"
	"github.com/aide-family/moon/pkg/palace/model"
	"github.com/aide-family/moon/pkg/palace/model/bizmodel"
	"github.com/aide-family/moon/pkg/palace/model/bizmodel/bizquery"
	"github.com/aide-family/moon/pkg/palace/model/query"
	"github.com/aide-family/moon/pkg/util/types"
	"github.com/aide-family/moon/pkg/vobj"
	"gorm.io/gen/field"
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
	isTeam := params.SourceType.GetValue() == vobj.SourceTeam.GetValue()
	ids := params.IDs
	status := params.Status
	if isTeam {
		claims, ok := middleware.ParseJwtClaims(ctx)
		if !ok {
			return merr.ErrorI18nUnLoginErr(ctx)
		}
		bizDB, err := l.data.GetBizGormDB(claims.GetTeam())
		if !types.IsNil(err) {
			return err
		}
		bizWrapper := bizquery.Use(bizDB).SysDict.WithContext(ctx)
		if _, err = bizWrapper.Where(bizquery.Use(bizDB).SysDict.ID.In(ids...)).Update(bizquery.Use(bizDB).SysDict.Status, params.Status); !types.IsNil(err) {
			return err
		}
	} else {
		_, err := query.Use(l.data.GetMainDB(ctx)).WithContext(ctx).SysDict.Where(query.SysDict.ID.In(ids...)).Update(query.SysDict.Status, status)
		return err
	}
	return nil
}

func (l *dictRepositoryImpl) DeleteByID(ctx context.Context, params *bo.DeleteDictParams) error {
	isTeam := params.SourceType.GetValue() == vobj.SourceTeam.GetValue()
	id := params.ID
	if isTeam {
		claims, ok := middleware.ParseJwtClaims(ctx)
		if !ok {
			return merr.ErrorI18nUnLoginErr(ctx)
		}
		bizDB, err := l.data.GetBizGormDB(claims.GetTeam())
		if !types.IsNil(err) {
			return err
		}
		err = bizquery.Use(bizDB).Transaction(func(tx *bizquery.Query) error {
			_, err := tx.SysDict.WithContext(ctx).Where(bizquery.Use(bizDB).SysDict.ID.Eq(id)).Delete()
			return err
		})
		if !types.IsNil(err) {
			return err
		}
	} else {
		_, err := query.Use(l.data.GetMainDB(ctx)).WithContext(ctx).SysDict.Where(query.SysDict.ID.Eq(id)).Delete()
		return err
	}
	return nil
}

func (l *dictRepositoryImpl) Create(ctx context.Context, dict *bo.CreateDictParams) (model.IDict, error) {
	sourceType, ok := middleware.ParseSourceTypeInfo(ctx)
	if !ok {
		return nil, merr.ErrorI18nRequestSourceParsingError(ctx)
	}
	isTeam := vobj.GetSourceType(sourceType.GetSourceCode()).GetValue() == vobj.SourceTeam.GetValue()
	if isTeam {
		// Team  creation
		return createDictModel(ctx, l.data, dict)
	}
	// system creation
	dictModel := createDictParamsToModel(ctx, dict)
	q := query.Use(l.data.GetMainDB(ctx)).WithContext(ctx).SysDict
	if err := q.Create(dictModel); !types.IsNil(err) {
		return nil, err
	}
	return dictModel, nil
}

func (l *dictRepositoryImpl) FindByPage(ctx context.Context, params *bo.QueryDictListParams) ([]model.IDict, error) {
	isTeam := params.SourceType.GetValue() == vobj.SourceTeam.GetValue()
	if isTeam {
		// Team  creation
		return listBizDictModel(ctx, l.data, params)
	} else {
		return listDictModel(ctx, l.data, params)
	}
}

func (l *dictRepositoryImpl) BatchCreate(ctx context.Context, createDicts []*bo.CreateDictParams) error {
	for _, dict := range createDicts {
		_, err := l.Create(ctx, dict)
		if !types.IsNil(err) {
			return err
		}
	}
	return nil
}

func (l *dictRepositoryImpl) GetByID(ctx context.Context, id uint32, sourceType vobj.SourceType) (model.IDict, error) {

	isTeam := sourceType.GetValue() == vobj.SourceTeam.GetValue()
	if isTeam {
		claims, ok := middleware.ParseJwtClaims(ctx)
		if !ok {
			return nil, merr.ErrorI18nUnLoginErr(ctx)
		}
		bizDB, err := l.data.GetBizGormDB(claims.GetTeam())
		if !types.IsNil(err) {
			return nil, err
		}
		bizWrapper := bizquery.Use(bizDB).SysDict.WithContext(ctx)
		return bizWrapper.Where(bizquery.Use(bizDB).SysDict.ID.Eq(id)).Preload(field.Associations).First()
	}

	return query.Use(l.data.GetMainDB(ctx)).SysDict.WithContext(ctx).Where(query.SysDict.ID.Eq(id)).First()
}

func (l *dictRepositoryImpl) UpdateByID(ctx context.Context, dict *bo.UpdateDictParams) error {
	isTeam := dict.SourceType.GetValue() == vobj.SourceTeam.GetValue()
	if isTeam {
		err := updateBizDictModel(ctx, l.data, dict)
		if !types.IsNil(err) {
			return err
		}
	} else {
		err := updateDictModel(ctx, l.data, dict)
		if !types.IsNil(err) {
			return err
		}
	}
	return nil
}

func listDictModel(ctx context.Context, data *data.Data, params *bo.QueryDictListParams) ([]model.IDict, error) {
	dict := query.Use(data.GetMainDB(ctx)).SysDict
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
	dbDicts, err := queryWrapper.Order(dict.ID.Desc()).Find()
	if !types.IsNil(err) {
		return nil, err
	}
	dicts := types.SliceTo(dbDicts, func(dict *model.SysDict) model.IDict {
		return dict
	})
	return dicts, nil
}

func listBizDictModel(ctx context.Context, data *data.Data, params *bo.QueryDictListParams) ([]model.IDict, error) {
	claims, ok := middleware.ParseJwtClaims(ctx)
	if !ok {
		return nil, merr.ErrorI18nUnLoginErr(ctx)
	}

	bizDB, err := data.GetBizGormDB(claims.GetTeam())
	if !types.IsNil(err) {
		return nil, err
	}
	bizWrapper := bizquery.Use(bizDB).SysDict.WithContext(ctx)

	var wheres []gen.Condition

	if !params.Status.IsUnknown() {
		wheres = append(wheres, bizquery.Use(bizDB).SysDict.Status.Eq(params.Status.GetValue()))
	}
	if !types.TextIsNull(params.Keyword) {
		bizWrapper = bizWrapper.Or(bizquery.Use(bizDB).SysDict.Name.Like(params.Keyword))
		bizWrapper = bizWrapper.Or(bizquery.Use(bizDB).SysDict.Value.Like(params.Keyword))
		bizWrapper = bizWrapper.Or(bizquery.Use(bizDB).SysDict.Remark.Like(params.Keyword))
	}

	bizWrapper = bizWrapper.Where(wheres...).Preload(field.Associations)

	if err := types.WithPageQuery[bizquery.ISysDictDo](bizWrapper, params.Page); err != nil {
		return nil, err
	}
	sysDicts, err := bizWrapper.Order(bizquery.Use(bizDB).SysDict.ID.Desc()).Find()
	if !types.IsNil(err) {
		return nil, err
	}
	dicts := types.SliceTo(sysDicts, func(dict *bizmodel.SysDict) model.IDict {
		return dict
	})
	return dicts, nil
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

// createDictModel create team dict model
func createDictModel(ctx context.Context, data *data.Data, dict *bo.CreateDictParams) (*bizmodel.SysDict, error) {
	claims, ok := middleware.ParseJwtClaims(ctx)
	if !ok {
		return nil, merr.ErrorI18nUnLoginErr(ctx)
	}
	bizDB, err := data.GetBizGormDB(claims.GetTeam())
	if !types.IsNil(err) {
		return nil, err
	}
	dictWrapper := bizquery.Use(bizDB)

	dictBizModel := createBizDictParamsToModel(ctx, dict)
	err = dictWrapper.Transaction(func(tx *bizquery.Query) error {
		if err := tx.SysDict.WithContext(ctx).Create(dictBizModel); !types.IsNil(err) {
			return err
		}
		return err
	})
	if !types.IsNil(err) {
		return nil, err
	}
	return dictBizModel, nil
}

func updateDictModel(ctx context.Context, data *data.Data, params *bo.UpdateDictParams) error {
	id := params.ID
	updateParam := params.UpdateParam
	_, err := query.Use(data.GetMainDB(ctx)).SysDict.WithContext(ctx).Where(query.SysDict.ID.Eq(id)).UpdateSimple(
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

func updateBizDictModel(ctx context.Context, data *data.Data, params *bo.UpdateDictParams) error {
	claims, ok := middleware.ParseJwtClaims(ctx)
	if !ok {
		return merr.ErrorI18nUnLoginErr(ctx)
	}
	bizDB, err := data.GetBizGormDB(claims.GetTeam())
	if !types.IsNil(err) {
		return err
	}
	updateParam := params.UpdateParam
	id := params.ID
	dictWrapper := bizquery.Use(bizDB)
	err = dictWrapper.Transaction(func(tx *bizquery.Query) error {
		if _, err = tx.SysDict.WithContext(ctx).Where(bizquery.Use(bizDB).SysDict.ID.Eq(id)).UpdateSimple(
			dictWrapper.SysDict.Name.Value(updateParam.Name),
			dictWrapper.SysDict.Remark.Value(updateParam.Remark),
			dictWrapper.SysDict.Value.Value(updateParam.Value),
			dictWrapper.SysDict.CssClass.Value(updateParam.CssClass),
			dictWrapper.SysDict.ColorType.Value(updateParam.ColorType),
			dictWrapper.SysDict.ImageUrl.Value(updateParam.ImageUrl),
			dictWrapper.SysDict.Icon.Value(updateParam.Icon),
		); !types.IsNil(err) {
			return err
		}
		return nil
	})
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
