package rpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

type Request struct {
	c *RESTClient

	callOpts []grpc.CallOption

	method  string
	api     string
	path    string
	service string

	params, out any

	err error
}

func NewRequest(c *RESTClient) *Request {
	return &Request{c: c, callOpts: make([]grpc.CallOption, 0)}
}

func (x *Request) Method(method string) *Request {
	if method == "" {
		x.err = fmt.Errorf("method is empty")
		return x
	}
	x.method = method
	return x
}

func (x *Request) API(api string) *Request {
	x.api = api
	return x
}

func (x *Request) Path(path string) *Request {
	x.path = path
	return x
}

func (x *Request) Service(service string) *Request {
	if x.err != nil {
		return x
	}
	if len(service) == 0 {
		x.err = fmt.Errorf("service is empty")
		return x
	}
	x.service = service
	return x
}

func (x *Request) Params(params any) *Request {
	if params == nil {
		return x
	}
	x.params = params
	return x
}

func (x *Request) Into(ctx context.Context, out any) error {
	if x.err != nil {
		return x.err
	}
	if x.out == nil {
		return fmt.Errorf("out is nil")
	}
	x.out = out
	x.do(ctx)
	return x.err
}

func (x *Request) do(ctx context.Context) {
	rpcMethodName := x.buildRPCMethodName()
	if len(x.callOpts) == 0 {
		x.callOpts = append(x.callOpts, grpc.StaticMethod())
	}
	x.err = x.c.client.Invoke(ctx, rpcMethodName, x.params, x.out, x.callOpts...)
}

func (x *Request) RPCMethodName() string {
	return x.buildRPCMethodName()
}

func (x *Request) buildRPCMethodName() string {
	if x.api == "" {
		x.api = "api"
	}
	if x.path != "" {
		return fmt.Sprintf("/%s.%s.%s/%s", x.api, x.path, x.service, x.method)
	}
	return fmt.Sprintf("/%s.%s/%s", x.api, x.service, x.method)
}
