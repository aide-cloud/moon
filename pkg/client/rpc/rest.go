package rpc

import (
	"google.golang.org/grpc"
)

type Interface interface {
	Method(method string) *Request
	Create() *Request
	Get() *Request
	List() *Request
	Update() *Request
	Delete() *Request
	Watch() *Request
}

const (
	MethodCreate = "Create"
	MethodGet    = "Get"
	MethodList   = "List"
	MethodUpdate = "Update"
	MethodDelete = "Delete"
	MethodWatch  = "Watch"
)

var _ Interface = &RESTClient{}

type RESTClient struct {
	client *grpc.ClientConn
}

func NewRESTClient() (*RESTClient, error) {
	var client *grpc.ClientConn
	// TODO: build grpc client
	return &RESTClient{client: client}, nil
}

func (x *RESTClient) Create() *Request {
	return x.Method(MethodCreate)
}

func (x *RESTClient) Get() *Request {
	return x.Method(MethodGet)
}

func (x *RESTClient) List() *Request {
	return x.Method(MethodList)
}

func (x *RESTClient) Update() *Request {
	return x.Method(MethodUpdate)
}

func (x *RESTClient) Delete() *Request {
	return x.Method(MethodDelete)
}
func (x *RESTClient) Watch() *Request {
	return x.Method(MethodWatch)
}

func (x *RESTClient) Method(method string) *Request {
	return NewRequest(x).Method(method)
}
