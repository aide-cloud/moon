// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v3.19.4
// source: houyi/metadata/metric.proto

package metadata

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationMetricQuery = "/api.houyi.metadata.Metric/Query"
const OperationMetricSyncMetadata = "/api.houyi.metadata.Metric/SyncMetadata"

type MetricHTTPServer interface {
	// Query 查询
	Query(context.Context, *QueryRequest) (*QueryReply, error)
	// SyncMetadata 同步元数据
	SyncMetadata(context.Context, *SyncMetadataRequest) (*SyncMetadataReply, error)
}

func RegisterMetricHTTPServer(s *http.Server, srv MetricHTTPServer) {
	r := s.Route("/")
	r.POST("/metric/sync/metadata", _Metric_SyncMetadata0_HTTP_Handler(srv))
	r.POST("/metric/query", _Metric_Query0_HTTP_Handler(srv))
}

func _Metric_SyncMetadata0_HTTP_Handler(srv MetricHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SyncMetadataRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMetricSyncMetadata)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SyncMetadata(ctx, req.(*SyncMetadataRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SyncMetadataReply)
		return ctx.Result(200, reply)
	}
}

func _Metric_Query0_HTTP_Handler(srv MetricHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in QueryRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMetricQuery)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Query(ctx, req.(*QueryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*QueryReply)
		return ctx.Result(200, reply)
	}
}

type MetricHTTPClient interface {
	Query(ctx context.Context, req *QueryRequest, opts ...http.CallOption) (rsp *QueryReply, err error)
	SyncMetadata(ctx context.Context, req *SyncMetadataRequest, opts ...http.CallOption) (rsp *SyncMetadataReply, err error)
}

type MetricHTTPClientImpl struct {
	cc *http.Client
}

func NewMetricHTTPClient(client *http.Client) MetricHTTPClient {
	return &MetricHTTPClientImpl{client}
}

func (c *MetricHTTPClientImpl) Query(ctx context.Context, in *QueryRequest, opts ...http.CallOption) (*QueryReply, error) {
	var out QueryReply
	pattern := "/metric/query"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMetricQuery))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *MetricHTTPClientImpl) SyncMetadata(ctx context.Context, in *SyncMetadataRequest, opts ...http.CallOption) (*SyncMetadataReply, error) {
	var out SyncMetadataReply
	pattern := "/metric/sync/metadata"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMetricSyncMetadata))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}