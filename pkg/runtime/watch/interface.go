package watch

import "github.com/aide-family/moon/pkg/runtime"

type Interface interface {
	Stop()

	ResultChan() <-chan Event
}

type Event struct {
	Type EventType

	Object runtime.Object
}

type EventType string

const (
	Added    EventType = "ADDED"
	Modified EventType = "MODIFIED"
	Deleted  EventType = "DELETED"
	Bookmark EventType = "BOOKMARK"
	Error    EventType = "ERROR"
)
