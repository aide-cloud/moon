package manager

import (
	"context"
	"github.com/aide-family/moon/api"
	"github.com/aide-family/moon/pkg/client/rpc"
	"github.com/aide-family/moon/pkg/runtime"
	"github.com/aide-family/moon/pkg/runtime/watch"
)

// ListFunc knows how to list resources
type ListFunc func(options *api.ListOptions) (runtime.Object, error)

// WatchFunc knows how to watch resources
type WatchFunc func(options *api.ListOptions) (watch.Interface, error)

type ListAndWatch struct {
	ListFunc  ListFunc
	WatchFunc WatchFunc
}

func NewListWatchWithClient(cli *rpc.RESTClient, kind string, scheme *runtime.Scheme, option *api.ListOptions) *ListAndWatch {
	listFunc := func(options *api.ListOptions) (runtime.Object, error) {
		obj, err := scheme.New(kind + "List")
		if err != nil {
			return nil, err
		}
		err = cli.List().
			API("api").
			Path("").
			Service(kind).
			Params(options).
			Into(context.TODO(), obj)
		if err != nil {
			return nil, err
		}
		return obj, err
	}
	watchFunc := func(options *api.ListOptions) (watch.Interface, error) {
		obj, err := scheme.New(kind)
		if err != nil {
			return nil, err
		}
		return cli.Watch().
			API("api").
			Path("").
			Service(kind).
			Params(options).
			Watch(context.TODO(), obj)
	}
	return &ListAndWatch{
		ListFunc:  listFunc,
		WatchFunc: watchFunc,
	}
}

// List a set of place resources
func (lw *ListAndWatch) List(options *api.ListOptions) (runtime.Object, error) {
	// ListWatch is used in Reflector, which already supports pagination.
	// Don't paginate here to avoid duplication.
	return lw.ListFunc(options)
}

// Watch a set of place resources
func (lw *ListAndWatch) Watch(options *api.ListOptions) (watch.Interface, error) {
	return lw.WatchFunc(options)
}
