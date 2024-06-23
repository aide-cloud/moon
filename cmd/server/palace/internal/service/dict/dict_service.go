package dict

import (
	"context"
	"github.com/aide-family/moon/api/admin"
	pb "github.com/aide-family/moon/api/admin/dict"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-family/moon/cmd/server/palace/internal/data"
	"github.com/aide-family/moon/cmd/server/palace/internal/service/build"
	"github.com/aide-family/moon/pkg/helper/model/palace"
	"github.com/aide-family/moon/pkg/util/types"
	"github.com/aide-family/moon/pkg/vobj"
)

type Service struct {
	pb.UnimplementedDictServer

	dictBiz *biz.DictBiz
}

func NewDictService(dictBiz *biz.DictBiz) *Service {
	return &Service{
		dictBiz: dictBiz,
	}
}

type dictRepositoryImpl struct {
	data *data.Data
}

func (s *Service) CreateDict(ctx context.Context, req *pb.CreateDictRequest) (*pb.CreateDictReply, error) {
	createParams := bo.CreateDictParams{
		Name:         req.Name,
		Value:        req.Value,
		DictType:     vobj.DictType(req.GetDictType()),
		ColorType:    req.ColorType,
		CssClass:     req.CssClass,
		Icon:         req.Icon,
		ImageUrl:     req.ImageUrl,
		Status:       vobj.Status(req.GetStatus()),
		Remark:       req.Remark,
		LanguageCode: req.LanguageCode,
	}

	_, err := s.dictBiz.CreateDict(ctx, &createParams)
	if err != nil {
		return nil, err
	}
	if !types.IsNil(err) {
		return nil, err
	}
	return &pb.CreateDictReply{}, nil
}

func (s *Service) UpdateDict(ctx context.Context, req *pb.UpdateDictRequest) (*pb.UpdateDictReply, error) {
	data := req.GetData()
	createParams := bo.CreateDictParams{
		Name:         data.Name,
		Value:        data.Value,
		DictType:     vobj.DictType(data.GetDictType()),
		ColorType:    data.ColorType,
		CssClass:     data.CssClass,
		Icon:         data.Icon,
		ImageUrl:     data.ImageUrl,
		Status:       vobj.Status(data.GetStatus()),
		Remark:       data.Remark,
		LanguageCode: data.LanguageCode,
	}

	updateParams := bo.UpdateDictParams{
		ID:          req.Id,
		UpdateParam: createParams,
	}
	if _, err := s.dictBiz.UpdateDict(ctx, &updateParams); !types.IsNil(err) {
		return nil, err
	}
	return &pb.UpdateDictReply{}, nil
}

func (s *Service) ListDict(ctx context.Context, req *pb.GetDictSelectListRequest) (*pb.ListDictReply, error) {

	queryParams := &bo.QueryDictListParams{
		Keyword:  req.GetKeyword(),
		Page:     types.NewPagination(req.GetPagination()),
		Status:   vobj.Status(req.GetStatus()),
		DictType: vobj.DictType(req.GetDictType()),
	}

	dictPage, err := s.dictBiz.ListDict(ctx, queryParams)
	if !types.IsNil(err) {
		return nil, err
	}
	return &pb.ListDictReply{
		Pagination: build.NewPageBuilder(queryParams.Page).ToApi(),
		List: types.SliceTo(dictPage, func(dict *palace.SysDict) *admin.Dict {
			return build.NewDictBuild(dict).ToApi()
		}),
	}, nil
}

func (s *Service) BatchUpdateDictStatus(context.Context, *pb.BatchUpdateDictStatusRequest) (*pb.BatchUpdateDictStatusReply, error) {

	return &pb.BatchUpdateDictStatusReply{}, nil
}

func (s *Service) DeleteDict(context.Context, *pb.DeleteDictRequest) (*pb.DeleteDictReply, error) {

	return &pb.DeleteDictReply{}, nil
}

func (s *Service) GetDict(ctx context.Context, req *pb.GetDictRequest) (*pb.GetDictReply, error) {

	dictDO, err := s.dictBiz.GetDict(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	resDict := build.NewDictBuild(dictDO).ToApi()
	return &pb.GetDictReply{
		Dict: resDict,
	}, nil
}
