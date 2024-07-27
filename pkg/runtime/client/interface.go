package client

import (
	"context"
	"github.com/aide-family/moon/api"
	"github.com/aide-family/moon/pkg/runtime"
	"github.com/aide-family/moon/pkg/runtime/meta"
	"github.com/aide-family/moon/pkg/runtime/watch"
)

type Object interface {
	meta.Object
	runtime.Object
}

type ObjectList interface {
	runtime.Object
}

type Reader interface {
	Get(ctx context.Context, key string, object Object) error
	List(ctx context.Context, object ObjectList, option *api.ListOptions) error
}

type Writer interface {
	Create(ctx context.Context, object Object) error
	Update(ctx context.Context, object Object) error
	Delete(ctx context.Context, object Object) error
}

type Client interface {
	Reader
	Writer
	Scheme() *runtime.Scheme
}

type WatchClient interface {
	Client
	Watch(ctx context.Context, obj runtime.Object, opts *api.ListOptions) (watch.Interface, error)
}
