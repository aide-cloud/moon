package rpc

import (
	"context"
	"fmt"
	"github.com/aide-family/moon/api"
	"github.com/aide-family/moon/pkg/runtime"
	"github.com/aide-family/moon/pkg/runtime/client"
	"github.com/aide-family/moon/pkg/runtime/watch"
	"google.golang.org/grpc"
	"reflect"
)

type watchClient struct {
	*rpc
}

func WithWatchClient(cc grpc.ClientConnInterface, scheme *runtime.Scheme) client.WatchClient {
	cli := newClient(cc, scheme)
	return &watchClient{rpc: cli}
}

func (x *watchClient) Watch(ctx context.Context, obj runtime.Object, opts *api.ListOptions) (watch.Interface, error) {
	cOpts := []grpc.CallOption{grpc.StaticMethod()}
	kind, err := x.scheme.ObjectKind(obj)
	if err != nil {
		return nil, err
	}
	rpcStreamFullMethodName := fmt.Sprintf("/api.%s/Stream", kind)
	// TODO:How does Stream work ?
	stream, err := x.cc.NewStream(ctx, nil, rpcStreamFullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	if err = stream.SendMsg(opts); err != nil {
		return nil, err
	}
	if err = stream.CloseSend(); err != nil {
		return nil, err
	}
	g := NewGenericWatchClient(obj, stream)
	go g.Run(ctx)

	return g, nil
}

var _ watch.Interface = &genericWatchClient{}

type genericWatchClient struct {
	ctx        context.Context
	cancel     context.CancelFunc
	objectType any
	stream     grpc.ClientStream
	resCh      chan watch.Event
}

func NewGenericWatchClient(object any, stream grpc.ClientStream) *genericWatchClient {
	return &genericWatchClient{
		objectType: object,
		stream:     stream,
		resCh:      make(chan watch.Event),
	}
}

func (g *genericWatchClient) Stop() {
	g.cancel()
}

func (g *genericWatchClient) ResultChan() <-chan watch.Event {
	return g.resCh
}

func (g *genericWatchClient) Run(ctx context.Context) {
	g.ctx, g.cancel = context.WithCancel(ctx)
	defer g.cancel()
	for {
		select {
		case <-g.ctx.Done():
			return
		default:
			res := watch.Event{}
			object := reflect.New(reflect.TypeOf(g.objectType)).Interface()
			if err := g.stream.RecvMsg(object); err != nil {
				res.Type = watch.Error
				g.resCh <- res
				continue
			}
			res.Type = watch.Added
			res.Object = object.(runtime.Object)
			g.resCh <- res
		}
	}
}
