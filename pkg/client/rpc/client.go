package rpc

import (
	"context"
	"fmt"
	"github.com/aide-family/moon/api"
	"github.com/aide-family/moon/pkg/runtime"
	"github.com/aide-family/moon/pkg/runtime/client"
	"google.golang.org/grpc"
	"strings"
)

// rpc implementation of client.Client
type rpc struct {
	cc     grpc.ClientConnInterface
	scheme *runtime.Scheme
}

func New(cc grpc.ClientConnInterface, scheme *runtime.Scheme) client.Client {
	return newClient(cc, scheme)
}

func newClient(cc grpc.ClientConnInterface, scheme *runtime.Scheme) *rpc {
	return &rpc{
		cc:     cc,
		scheme: scheme,
	}
}

func (x *rpc) Get(ctx context.Context, key string, object client.Object) error {
	cOpts := []grpc.CallOption{grpc.StaticMethod()}
	kind, err := x.scheme.ObjectKind(object)
	if err != nil {
		return err
	}
	rpcFullMethodName := fmt.Sprintf("/api.%s/Get", kind)
	args := api.GetOptions{Name: key}
	return x.cc.Invoke(ctx, rpcFullMethodName, &args, object, cOpts...)
}

func (x *rpc) List(ctx context.Context, list client.ObjectList, listOption *api.ListOptions) error {
	cOpts := []grpc.CallOption{grpc.StaticMethod()}
	kind, err := x.scheme.ObjectKind(list)
	if err != nil {
		return err
	}
	if !strings.HasSuffix(kind, "List") {
		return fmt.Errorf("non-list type %T (kind %q) passed as output", list, kind)
	}
	kind = kind[:len(kind)-4]
	rpcFullMethodName := fmt.Sprintf("/api.%s/List", kind)
	return x.cc.Invoke(ctx, rpcFullMethodName, listOption, list, cOpts...)
}

func (x *rpc) Create(ctx context.Context, object client.Object) error {
	cOpts := []grpc.CallOption{grpc.StaticMethod()}
	kind, err := x.scheme.ObjectKind(object)
	if err != nil {
		return err
	}
	rpcFullMethodName := fmt.Sprintf("/api.%s/Create", kind)
	return x.cc.Invoke(ctx, rpcFullMethodName, object, object, cOpts...)
}

func (x *rpc) Update(ctx context.Context, object client.Object) error {
	cOpts := []grpc.CallOption{grpc.StaticMethod()}
	kind, err := x.scheme.ObjectKind(object)
	if err != nil {
		return err
	}
	rpcFullMethodName := fmt.Sprintf("/api.%s/Update", kind)
	return x.cc.Invoke(ctx, rpcFullMethodName, object, object, cOpts...)

}

func (x *rpc) Delete(ctx context.Context, object client.Object) error {
	cOpts := []grpc.CallOption{grpc.StaticMethod()}
	kind, err := x.scheme.ObjectKind(object)
	if err != nil {
		return err
	}
	rpcFullMethodName := fmt.Sprintf("/api.%s/Delete", kind)
	args := api.DeleteOptions{Name: object.GetName()}
	return x.cc.Invoke(ctx, rpcFullMethodName, &args, object, cOpts...)

}

func (x *rpc) Scheme() *runtime.Scheme {
	return x.scheme
}
