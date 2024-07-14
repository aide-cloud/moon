package strategy

import (
	"context"

	"github.com/aide-family/moon/api/admin"
	"github.com/aide-family/moon/api/merr"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-family/moon/cmd/server/palace/internal/service/build"
	"github.com/aide-family/moon/pkg/helper/middleware"
	"github.com/aide-family/moon/pkg/palace/model/bizmodel"
	"github.com/aide-family/moon/pkg/util/types"
	"github.com/aide-family/moon/pkg/vobj"

	strategyapi "github.com/aide-family/moon/api/admin/strategy"
)

type Service struct {
	strategyapi.UnimplementedStrategyServer
	templateBiz *biz.TemplateBiz
	strategy    *biz.StrategyBiz
}

func NewStrategyService(templateBiz *biz.TemplateBiz, strategy *biz.StrategyBiz) *Service {
	return &Service{
		templateBiz: templateBiz,
		strategy:    strategy,
	}
}

func (s *Service) CreateStrategyGroup(ctx context.Context, req *strategyapi.CreateStrategyGroupRequest) (*strategyapi.CreateStrategyGroupReply, error) {
	return &strategyapi.CreateStrategyGroupReply{}, nil
}

func (s *Service) DeleteStrategyGroup(ctx context.Context, req *strategyapi.DeleteStrategyGroupRequest) (*strategyapi.DeleteStrategyGroupReply, error) {
	return &strategyapi.DeleteStrategyGroupReply{}, nil
}

func (s *Service) ListStrategyGroup(ctx context.Context, req *strategyapi.ListStrategyGroupRequest) (*strategyapi.ListStrategyGroupReply, error) {
	return &strategyapi.ListStrategyGroupReply{}, nil
}

func (s *Service) GetStrategyGroup(ctx context.Context, req *strategyapi.GetStrategyGroupRequest) (*strategyapi.GetStrategyGroupReply, error) {
	return &strategyapi.GetStrategyGroupReply{}, nil
}

func (s *Service) UpdateStrategyGroup(ctx context.Context, req *strategyapi.UpdateStrategyGroupRequest) (*strategyapi.UpdateStrategyGroupReply, error) {
	return &strategyapi.UpdateStrategyGroupReply{}, nil
}

func (s *Service) CreateStrategy(ctx context.Context, req *strategyapi.CreateStrategyRequest) (*strategyapi.CreateStrategyReply, error) {
	claims, ok := middleware.ParseJwtClaims(ctx)
	if !ok {
		return nil, merr.ErrorI18nUnLoginErr(ctx)
	}
	param := &bo.CreateStrategyParams{
		TeamID:        claims.GetTeam(),
		TemplateId:    req.GetTemplateId(),
		GroupId:       req.GetGroupId(),
		Name:          req.GetName(),
		Remark:        req.GetRemark(),
		Status:        vobj.Status(req.GetStatus()),
		Step:          req.GetStep(),
		SourceType:    vobj.TemplateSourceType(req.GetSourceType()),
		DatasourceIds: req.GetDatasourceIds(),
		Threshold:     req.GetThreshold(),
		Condition:     req.GetCondition(),
	}
	_, err := s.strategy.CreateStrategy(ctx, param)
	if err != nil {
		return nil, err
	}
	return &strategyapi.CreateStrategyReply{}, nil
}

func (s *Service) UpdateStrategy(ctx context.Context, req *strategyapi.UpdateStrategyRequest) (*strategyapi.UpdateStrategyReply, error) {
	claims, ok := middleware.ParseJwtClaims(ctx)
	if !ok {
		return nil, merr.ErrorI18nUnLoginErr(ctx)
	}
	param := &bo.UpdateStrategyParams{
		TeamID: claims.GetTeam(),
		ID:     req.GetId(),
		UpdateParam: bo.CreateStrategyParams{
			TemplateId:    req.GetData().GetTemplateId(),
			GroupId:       req.GetData().GetGroupId(),
			Name:          req.GetData().GetName(),
			Remark:        req.GetData().GetRemark(),
			Status:        vobj.Status(req.GetData().GetStatus()),
			Step:          req.GetData().GetStep(),
			SourceType:    vobj.TemplateSourceType(req.GetData().GetSourceType()),
			DatasourceIds: req.GetData().GetDatasourceIds(),
			Threshold:     req.GetData().GetThreshold(),
			Condition:     req.GetData().GetCondition(),
		},
	}
	err := s.strategy.UpdateByID(ctx, param)
	if !types.IsNil(err) {
		return nil, err
	}
	return &strategyapi.UpdateStrategyReply{}, nil
}

func (s *Service) DeleteStrategy(ctx context.Context, req *strategyapi.DeleteStrategyRequest) (*strategyapi.DeleteStrategyReply, error) {
	claims, ok := middleware.ParseJwtClaims(ctx)
	if !ok {
		return nil, merr.ErrorI18nUnLoginErr(ctx)
	}
	param := &bo.DelStrategyParams{
		TeamID: claims.GetTeam(),
		ID:     req.GetId(),
	}
	err := s.strategy.DeleteByID(ctx, param)
	if !types.IsNil(err) {
		return nil, err
	}
	return &strategyapi.DeleteStrategyReply{}, nil
}

func (s *Service) GetStrategy(ctx context.Context, req *strategyapi.GetStrategyRequest) (*strategyapi.GetStrategyReply, error) {
	claims, ok := middleware.ParseJwtClaims(ctx)
	if !ok {
		return nil, merr.ErrorI18nUnLoginErr(ctx)
	}
	param := &bo.GetStrategyDetailParams{
		TeamID: claims.GetTeam(),
		ID:     req.GetId(),
	}
	strategy, err := s.strategy.GetStrategy(ctx, param)
	if err != nil {
		return nil, err
	}
	return &strategyapi.GetStrategyReply{
		Detail: build.NewStrategyBuilder(strategy).ToApi(ctx),
	}, nil
}

func (s *Service) ListStrategy(ctx context.Context, req *strategyapi.ListStrategyRequest) (*strategyapi.ListStrategyReply, error) {
	claims, ok := middleware.ParseJwtClaims(ctx)
	if !ok {
		return nil, merr.ErrorI18nUnLoginErr(ctx)
	}
	params := &bo.QueryStrategyListParams{
		TeamID:     claims.GetTeam(),
		Page:       types.NewPagination(req.GetPagination()),
		Status:     vobj.Status(req.GetStatus()),
		Keyword:    req.GetKeyword(),
		SourceType: vobj.TemplateSourceType(req.GetDatasourceType()),
	}
	strategies, err := s.strategy.StrategyPage(ctx, params)
	if err != nil {
		return nil, err
	}
	return &strategyapi.ListStrategyReply{
		Pagination: build.NewPageBuilder(params.Page).ToApi(),
		List: types.SliceTo(strategies, func(str *bizmodel.Strategy) *admin.Strategy {
			return build.NewStrategyBuilder(str).ToApi(ctx)
		}),
	}, nil
}
