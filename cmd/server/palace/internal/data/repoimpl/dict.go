package repoimpl

import (
	"context"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/repository"
	"github.com/aide-family/moon/cmd/server/palace/internal/data"
	"github.com/aide-family/moon/pkg/helper/model/palace"
	palacequery "github.com/aide-family/moon/pkg/helper/model/palace/query"
	"github.com/aide-family/moon/pkg/helper/model/query"
	"github.com/aide-family/moon/pkg/types"
	"github.com/aide-family/moon/pkg/vobj"
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

func (l *dictRepositoryImpl) UpdateStatusByIds(ctx context.Context, status vobj.Status, ids ...uint32) error {
	_, err := palacequery.Use(l.data.GetMainDB(ctx)).WithContext(ctx).SysDict.Where(palacequery.SysDict.ID.In(ids...)).Update(palacequery.SysDict.Status, status)
	return err
}

func (l *dictRepositoryImpl) DeleteByID(ctx context.Context, id uint32) error {
	_, err := palacequery.Use(l.data.GetMainDB(ctx)).WithContext(ctx).SysDict.Where(palacequery.SysDict.ID.Eq(id)).Delete()
	return err
}

func (l *dictRepositoryImpl) Create(ctx context.Context, dict *bo.CreateDictParams) (*palace.SysDict, error) {
	dictModel := createDictParamsToModel(dict)
	if err := dictModel.Create(ctx, l.data.GetMainDB(ctx)); !types.IsNil(err) {
		return nil, err
	}
	return dictModel, nil
}

func (l *dictRepositoryImpl) FindByPage(ctx context.Context, params *bo.QueryDictListParams) ([]*palace.SysDict, error) {
	queryWrapper := palacequery.Use(l.data.GetMainDB(ctx)).SysDict.WithContext(ctx)

	var wheres []gen.Condition
	if !params.Status.IsUnknown() {
		wheres = append(wheres, palacequery.SysDict.Status.Eq(params.Status.GetValue()))
	}

	if !params.DictType.IsUnknown() {
		wheres = append(wheres, query.SysUser.Gender.Eq(params.DictType.GetValue()))
	}

	if !types.TextIsNull(params.Keyword) {
		queryWrapper = queryWrapper.Or(
			palacequery.SysDict.Name.Like(params.Keyword),
			palacequery.SysDict.Value.Like(params.Keyword),
			palacequery.SysDict.Remark.Like(params.Keyword),
		)
	}

	queryWrapper = queryWrapper.Where(wheres...)

	if !types.IsNil(params) {
		page := params.Page
		total, err := queryWrapper.Count()
		if !types.IsNil(err) {
			return nil, err
		}
		params.Page.SetTotal(int(total))
		pageNum, pageSize := page.GetPageNum(), page.GetPageSize()
		if pageNum <= 1 {
			queryWrapper = queryWrapper.Limit(pageSize)
		} else {
			queryWrapper = queryWrapper.Offset((pageNum - 1) * pageSize).Limit(pageSize)
		}
	}
	return queryWrapper.Order(palacequery.SysDict.ID.Desc()).Find()
}

func (l *dictRepositoryImpl) BatchCreate(ctx context.Context, users []*bo.CreateDictParams) error {

	dictModels := types.SliceToWithFilter(users, func(item *bo.CreateDictParams) (*palace.SysDict, bool) {
		if types.IsNil(item) || types.TextIsNull(item.Name) {
			return nil, false
		}
		return createDictParamsToModel(item), true
	})
	return palacequery.Use(l.data.GetMainDB(ctx)).WithContext(ctx).SysDict.CreateInBatches(dictModels, 10)
}

func (l *dictRepositoryImpl) GetByID(ctx context.Context, id uint32) (*palace.SysDict, error) {
	return palacequery.Use(l.data.GetMainDB(ctx)).SysDict.WithContext(ctx).Where(query.SysUser.ID.Eq(id)).First()
}

func (l *dictRepositoryImpl) UpdateByID(ctx context.Context, dict *bo.UpdateDictParams) error {
	updateParam := dict.UpdateParam
	_, err := palacequery.Use(l.data.GetMainDB(ctx)).SysDict.WithContext(ctx).Where(query.SysUser.ID.Eq(dict.ID)).UpdateSimple(
		palacequery.SysDict.Name.Value(updateParam.Name),
		palacequery.SysDict.Value.Value(updateParam.Value),
		palacequery.SysDict.CssClass.Value(updateParam.CssClass),
		palacequery.SysDict.ColorType.Value(updateParam.ColorType),
		palacequery.SysDict.Remark.Value(updateParam.Remark),
		palacequery.SysDict.ImageUrl.Value(updateParam.ImageUrl),
		palacequery.SysDict.Icon.Value(updateParam.Icon),
	)
	return err

}

// createDictParamsToModel create dict params to model
func createDictParamsToModel(dict *bo.CreateDictParams) *palace.SysDict {
	if types.IsNil(dict) {
		return nil
	}
	return &palace.SysDict{
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
}
