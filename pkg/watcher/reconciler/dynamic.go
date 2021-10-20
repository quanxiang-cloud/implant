package reconciler

import (
	"context"

	v1alpha1 "github.com/quanxiang-cloud/overseer/pkg/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
)

func NewImpl(opts ...Options) *Impl {
	impl := &Impl{
		queue: newQueue(),
	}

	for _, opt := range opts {
		opt(impl)
	}

	return impl
}

type Options func(*Impl)

func WithConsumer(ctx context.Context, fn func(interface{})) Options {
	return func(i *Impl) {
		go i.queue.Consumer(ctx, fn)
	}
}

type Impl struct {
	Name string

	informer cache.SharedIndexInformer
	queue    *queue
}

func (i *Impl) AddFunc(obj interface{}) {
	i.post(ADD, obj)
}
func (i *Impl) UpdateFunc(oldObj, newObj interface{}) {
	i.post(UPDATE, newObj)
}
func (i *Impl) DeleteFunc(obj interface{}) {
	i.post(DELETE, obj)
}

func (i *Impl) post(method string, obj interface{}) {
	sm, ok := i.getState(obj)
	if ok {
		i.queue.Send(Object{
			Method:  method,
			Summary: sm,
		})
	}
}

func (i *Impl) getState(obj interface{}) (*StatusSummary, bool) {
	osr, ok := obj.(*v1alpha1.OverseerRun)
	if !ok {
		return nil, false
	}

	return &StatusSummary{
		ObjectMeta: osr.ObjectMeta,
		Status:     osr.Status,
	}, true
}

func (i *Impl) Run(stopCh <-chan struct{}) {
	i.informer.Run(stopCh)
}

const (
	ADD    = "ADD"
	UPDATE = "UPDATE"
	DELETE = "DELETE"
)

type Object struct {
	Method  string
	Summary *StatusSummary
}

type StatusSummary struct {
	metav1.ObjectMeta
	Status v1alpha1.OverseerRunStatus
}
