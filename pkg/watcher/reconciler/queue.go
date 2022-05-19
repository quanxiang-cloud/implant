package reconciler

import (
	"context"

	"github.com/quanxiang-cloud/implant/pkg/watcher/broadcaster/bus"
)

func newQueue() *queue {
	return &queue{
		Chan: make(chan interface{}),
	}
}

type queue struct {
	Chan chan interface{}
}

func (q *queue) Send(item interface{}) {
	q.Chan <- item
}

func (q *queue) Consumer(ctx context.Context, bus *bus.EventBus) {
	for {
		select {
		case item, ok := <-q.Chan:
			if !ok {
				return
			}
			bus.Send(ctx, item)
		case <-ctx.Done():
		}
	}
}
