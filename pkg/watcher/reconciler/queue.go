package reconciler

import "context"

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

func (q *queue) Consumer(ctx context.Context, fn func(interface{})) {
	for {
		select {
		case item, ok := <-q.Chan:
			if !ok {
				return
			}
			fn(item)
		case <-ctx.Done():
		}
	}
}
