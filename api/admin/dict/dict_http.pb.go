// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v5.27.1
// source: admin/dict/dict.proto

package dict

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

const OperationDictBatchUpdateDictStatus = "/api.admin.dict.Dict/BatchUpdateDictStatus"
const OperationDictCreateDict = "/api.admin.dict.Dict/CreateDict"
const OperationDictDeleteDict = "/api.admin.dict.Dict/DeleteDict"
const OperationDictGetDict = "/api.admin.dict.Dict/GetDict"
const OperationDictListDict = "/api.admin.dict.Dict/ListDict"
const OperationDictUpdateDict = "/api.admin.dict.Dict/UpdateDict"

type DictHTTPServer interface {
	// BatchUpdateDictStatus 批量修改字典状态
	BatchUpdateDictStatus(context.Context, *BatchUpdateDictStatusRequest) (*BatchUpdateDictStatusReply, error)
	CreateDict(context.Context, *CreateDictRequest) (*CreateDictReply, error)
	DeleteDict(context.Context, *DeleteDictRequest) (*DeleteDictReply, error)
	GetDict(context.Context, *GetDictRequest) (*GetDictReply, error)
	// ListDict 字典列表
	ListDict(context.Context, *GetDictSelectListRequest) (*ListDictReply, error)
	// UpdateDict 更新用户
	UpdateDict(context.Context, *UpdateDictRequest) (*UpdateDictReply, error)
}

func RegisterDictHTTPServer(s *http.Server, srv DictHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/dict/create", _Dict_CreateDict0_HTTP_Handler(srv))
	r.PUT("/v1/dict/update/{id}", _Dict_UpdateDict0_HTTP_Handler(srv))
	r.POST("/v1/dict/list", _Dict_ListDict0_HTTP_Handler(srv))
	r.PUT("/v1/dict/status", _Dict_BatchUpdateDictStatus0_HTTP_Handler(srv))
	r.DELETE("/v1/dict/delete/{id}", _Dict_DeleteDict0_HTTP_Handler(srv))
	r.GET("/v1/dict/get/{id}", _Dict_GetDict0_HTTP_Handler(srv))
}

func _Dict_CreateDict0_HTTP_Handler(srv DictHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateDictRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictCreateDict)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateDict(ctx, req.(*CreateDictRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateDictReply)
		return ctx.Result(200, reply)
	}
}

func _Dict_UpdateDict0_HTTP_Handler(srv DictHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateDictRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictUpdateDict)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateDict(ctx, req.(*UpdateDictRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateDictReply)
		return ctx.Result(200, reply)
	}
}

func _Dict_ListDict0_HTTP_Handler(srv DictHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetDictSelectListRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictListDict)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListDict(ctx, req.(*GetDictSelectListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListDictReply)
		return ctx.Result(200, reply)
	}
}

func _Dict_BatchUpdateDictStatus0_HTTP_Handler(srv DictHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BatchUpdateDictStatusRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictBatchUpdateDictStatus)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BatchUpdateDictStatus(ctx, req.(*BatchUpdateDictStatusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BatchUpdateDictStatusReply)
		return ctx.Result(200, reply)
	}
}

func _Dict_DeleteDict0_HTTP_Handler(srv DictHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteDictRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictDeleteDict)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteDict(ctx, req.(*DeleteDictRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteDictReply)
		return ctx.Result(200, reply)
	}
}

func _Dict_GetDict0_HTTP_Handler(srv DictHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetDictRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictGetDict)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetDict(ctx, req.(*GetDictRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetDictReply)
		return ctx.Result(200, reply)
	}
}

type DictHTTPClient interface {
	BatchUpdateDictStatus(ctx context.Context, req *BatchUpdateDictStatusRequest, opts ...http.CallOption) (rsp *BatchUpdateDictStatusReply, err error)
	CreateDict(ctx context.Context, req *CreateDictRequest, opts ...http.CallOption) (rsp *CreateDictReply, err error)
	DeleteDict(ctx context.Context, req *DeleteDictRequest, opts ...http.CallOption) (rsp *DeleteDictReply, err error)
	GetDict(ctx context.Context, req *GetDictRequest, opts ...http.CallOption) (rsp *GetDictReply, err error)
	ListDict(ctx context.Context, req *GetDictSelectListRequest, opts ...http.CallOption) (rsp *ListDictReply, err error)
	UpdateDict(ctx context.Context, req *UpdateDictRequest, opts ...http.CallOption) (rsp *UpdateDictReply, err error)
}

type DictHTTPClientImpl struct {
	cc *http.Client
}

func NewDictHTTPClient(client *http.Client) DictHTTPClient {
	return &DictHTTPClientImpl{client}
}

func (c *DictHTTPClientImpl) BatchUpdateDictStatus(ctx context.Context, in *BatchUpdateDictStatusRequest, opts ...http.CallOption) (*BatchUpdateDictStatusReply, error) {
	var out BatchUpdateDictStatusReply
	pattern := "/v1/dict/status"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDictBatchUpdateDictStatus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *DictHTTPClientImpl) CreateDict(ctx context.Context, in *CreateDictRequest, opts ...http.CallOption) (*CreateDictReply, error) {
	var out CreateDictReply
	pattern := "/v1/dict/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDictCreateDict))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *DictHTTPClientImpl) DeleteDict(ctx context.Context, in *DeleteDictRequest, opts ...http.CallOption) (*DeleteDictReply, error) {
	var out DeleteDictReply
	pattern := "/v1/dict/delete/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDictDeleteDict))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *DictHTTPClientImpl) GetDict(ctx context.Context, in *GetDictRequest, opts ...http.CallOption) (*GetDictReply, error) {
	var out GetDictReply
	pattern := "/v1/dict/get/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDictGetDict))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *DictHTTPClientImpl) ListDict(ctx context.Context, in *GetDictSelectListRequest, opts ...http.CallOption) (*ListDictReply, error) {
	var out ListDictReply
	pattern := "/v1/dict/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDictListDict))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *DictHTTPClientImpl) UpdateDict(ctx context.Context, in *UpdateDictRequest, opts ...http.CallOption) (*UpdateDictReply, error) {
	var out UpdateDictReply
	pattern := "/v1/dict/update/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDictUpdateDict))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
