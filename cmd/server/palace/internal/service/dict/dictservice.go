package dict

import (
	"context"

	"github.com/aide-family/moon/api/admin"
	dictapi "github.com/aide-family/moon/api/admin/dict"
	"github.com/aide-family/moon/api/merr"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-family/moon/cmd/server/palace/internal/service/build"
	"github.com/aide-family/moon/pkg/helper/middleware"
	"github.com/aide-family/moon/pkg/palace/model"
	"github.com/aide-family/moon/pkg/palace/model/bizmodel"
	"github.com/aide-family/moon/pkg/util/types"
	"github.com/aide-family/moon/pkg/vobj"
)

type Service struct {
	dictapi.UnimplementedDictServer

	dictBiz *biz.DictBiz
}

func NewDictService(dictBiz *biz.DictBiz) *Service {
	return &Service{
		dictBiz: dictBiz,
	}
}

func (s *Service) CreateDict(ctx context.Context, req *dictapi.CreateDictRequest) (*dictapi.CreateDictReply, error) {
	createParams := build.NewBuilder().WithCreateBoDict(req).ToCreateDictBO()
	_, err := s.dictBiz.CreateDict(ctx, createParams)
	if err != nil {
		return nil, err
	}
	if !types.IsNil(err) {
		return nil, err
	}
	return &dictapi.CreateDictReply{}, nil
}

func (s *Service) UpdateDict(ctx context.Context, req *dictapi.UpdateDictRequest) (*dictapi.UpdateDictReply, error) {
	updateParams := build.NewBuilder().WithUpdateBoDict(req).ToUpdateDictBO()
	sourceType, ok := middleware.ParseSourceTypeInfo(ctx)
	if !ok {
		return nil, merr.ErrorI18nRequestSourceParsingError(ctx)
	}
	updateParams.SourceType = vobj.GetSourceType(sourceType.GetSourceCode())
	if err := s.dictBiz.UpdateDict(ctx, updateParams); !types.IsNil(err) {
		return nil, err
	}
	return &dictapi.UpdateDictReply{}, nil
}

func (s *Service) ListDict(ctx context.Context, req *dictapi.GetDictSelectListRequest) (*dictapi.ListDictReply, error) {
	sourceType, ok := middleware.ParseSourceTypeInfo(ctx)
	if !ok {
		return nil, merr.ErrorI18nRequestSourceParsingError(ctx)
	}
	queryParams := &bo.QueryDictListParams{
		Keyword:    req.GetKeyword(),
		Page:       types.NewPagination(req.GetPagination()),
		Status:     vobj.Status(req.GetStatus()),
		DictType:   vobj.DictType(req.GetDictType()),
		SourceType: vobj.GetSourceType(sourceType.GetSourceCode()),
	}

	dictPage, err := s.dictBiz.ListDict(ctx, queryParams)
	if !types.IsNil(err) {
		return nil, err
	}
	return &dictapi.ListDictReply{
		Pagination: build.NewPageBuilder(queryParams.Page).ToApi(),
		List: types.SliceTo(dictPage, func(dict model.IDict) *admin.Dict {
			resDict := &admin.Dict{}
			if vobj.GetSourceType(sourceType.GetSourceCode()).GetValue() == vobj.SourceTeam.GetValue() {
				resDict = build.NewBuilder().WithApiBizDict(dict.(*bizmodel.SysDict)).ToBizApi()
			} else {
				resDict = build.NewBuilder().WithApiDict(dict.(*model.SysDict)).ToApi()
			}
			return resDict
		}),
	}, nil
}

func (s *Service) BatchUpdateDictStatus(ctx context.Context, params *dictapi.BatchUpdateDictStatusRequest) (*dictapi.BatchUpdateDictStatusReply, error) {
	updateParams := bo.UpdateDictStatusParams{
		IDs:    params.GetIds(),
		Status: vobj.Status(params.Status),
	}
	err := s.dictBiz.UpdateDictStatusByIds(ctx, &updateParams)
	if !types.IsNil(err) {
		return nil, merr.ErrorI18nSystemErr(ctx).WithCause(err)
	}
	return &dictapi.BatchUpdateDictStatusReply{}, nil
}

func (s *Service) DeleteDict(ctx context.Context, req *dictapi.DeleteDictRequest) (*dictapi.DeleteDictReply, error) {
	sourceType, ok := middleware.ParseSourceTypeInfo(ctx)
	if !ok {
		return nil, merr.ErrorI18nRequestSourceParsingError(ctx)
	}

	params := &bo.DeleteDictParams{
		ID:         req.GetId(),
		SourceType: vobj.GetSourceType(sourceType.GetSourceCode()),
	}
	err := s.dictBiz.DeleteDictById(ctx, params)
	if !types.IsNil(err) {
		return nil, merr.ErrorI18nSystemErr(ctx).WithCause(err)
	}
	return &dictapi.DeleteDictReply{}, nil
}

func (s *Service) GetDict(ctx context.Context, req *dictapi.GetDictRequest) (*dictapi.GetDictReply, error) {
	sourceType, ok := middleware.ParseSourceTypeInfo(ctx)
	if !ok {
		return nil, merr.ErrorI18nRequestSourceParsingError(ctx)
	}
	params := &bo.GetDictDetailParams{
		ID:         req.GetId(),
		SourceType: vobj.GetSourceType(sourceType.GetSourceCode()),
	}
	dictDO, err := s.dictBiz.GetDict(ctx, params)
	if !types.IsNil(err) {
		return nil, err
	}
	resDict := &admin.Dict{}
	if vobj.GetSourceType(sourceType.GetSourceCode()).GetValue() == vobj.SourceTeam.GetValue() {
		resDict = build.NewBuilder().WithApiBizDict(dictDO.(*bizmodel.SysDict)).ToBizApi()
	} else {
		resDict = build.NewBuilder().WithApiDict(dictDO.(*model.SysDict)).ToApi()
	}
	return &dictapi.GetDictReply{
		Dict: resDict,
	}, nil
}

// ListDictType 获取字典类型列表
func (s *Service) ListDictType(_ context.Context, _ *dictapi.ListDictTypeRequest) (*dictapi.ListDictTypeReply, error) {
	return &dictapi.ListDictTypeReply{
		List: build.NewDictTypeBuilder().ToApi(),
	}, nil
}