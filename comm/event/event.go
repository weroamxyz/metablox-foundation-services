package event

import (
	"github.com/asaskevich/EventBus"
	"github.com/pkg/errors"
)

const (
	WorkloadValidated Event = iota
)

var (
	ErrUnknownEvent = errors.New("Unknown Event")
	eventBus        = EventBus.New()
	eventMap        = map[Event]string{
		WorkloadValidated: "WorkloadValidated",
	}
)

type Event int

func (e Event) Topic() string {
	return eventMap[e]
}

func Subscribe(event Event, fn interface{}) error {
	if event.Topic() == "" {
		return ErrUnknownEvent
	}
	return eventBus.Subscribe(event.Topic(), fn)
}

func SubscribeAsync(event Event, fn interface{}, transactional bool) error {
	if event.Topic() == "" {
		return ErrUnknownEvent
	}
	return eventBus.SubscribeAsync(event.Topic(), fn, transactional)
}

func SubscribeOnce(event Event, fn interface{}) error {
	if event.Topic() == "" {
		return ErrUnknownEvent
	}
	return eventBus.SubscribeOnce(event.Topic(), fn)
}

func SubscribeOnceAsync(event Event, fn interface{}) error {
	if event.Topic() == "" {
		return ErrUnknownEvent
	}
	return eventBus.SubscribeOnceAsync(event.Topic(), fn)
}

func Unsubscribe(event Event, handler interface{}) error {
	if event.Topic() == "" {
		return ErrUnknownEvent
	}
	return eventBus.Unsubscribe(event.Topic(), handler)
}

func Publish(event Event, args ...interface{}) {
	eventBus.Publish(event.Topic(), args...)
}

func HasCallback(event Event) bool {
	return eventBus.HasCallback(event.Topic())
}

func WaitAsync() {
	eventBus.WaitAsync()
}
