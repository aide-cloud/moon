package cache

import (
	"github.com/aide-family/moon/api"
	"github.com/aide-family/moon/pkg/runtime"
	"github.com/aide-family/moon/pkg/runtime/watch"
)

// ListFunc knows how to list resources
type ListFunc func(options api.ListOptions) (runtime.Object, error)

// WatchFunc knows how to watch resources
type WatchFunc func(options api.ListOptions) (watch.Interface, error)

type ListAndWatch struct {
	ListFunc  ListFunc
	WatchFunc WatchFunc
}
