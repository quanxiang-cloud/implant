package reconciler

import (
	"context"

	"github.com/golang/groupcache/lru"
	hs "github.com/mitchellh/hashstructure/v2"
	v1alpha1 "github.com/quanxiang-cloud/overseer/pkg/apis/overseer/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog"
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

func WithCache(ctx context.Context, maxEntries int) Options {
	return func(i *Impl) {
		i.cache = lru.New(maxEntries)
	}
}

type Impl struct {
	Name string

	informer cache.SharedIndexInformer
	cache    *lru.Cache
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

func (i *Impl) shouldDiscard(obj interface{}) bool {
	if i.cache == nil {
		return false
	}
	hash, err := hs.Hash(obj, hs.FormatV2, nil)
	if err != nil {
		klog.Error(err, obj)
		return false
	}
	_, ok := i.cache.Get(hash)
	if ok {
		return ok
	}

	i.cache.Add(hash, struct{}{})
	return false
}

func (i *Impl) post(method string, obj interface{}) {
	if i.shouldDiscard(obj) {
		return
	}
	sm, ok := i.getState(obj)
	if ok {
		i.queue.Send(Object{
			Method:  method,
			Summary: sm,
		})
	}
}

func (i *Impl) getState(obj interface{}) (*StatusSummary, bool) {
	osr, ok := obj.(*v1alpha1.Overseer)
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
	Status v1alpha1.OverseerStatus
}
