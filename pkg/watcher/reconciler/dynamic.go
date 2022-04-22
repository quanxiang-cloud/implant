package reconciler

import (
	"context"

	"github.com/golang/groupcache/lru"
	hs "github.com/mitchellh/hashstructure/v2"
	fnV1beta1 "github.com/openfunction/apis/core/v1beta1"
	prV1beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
	klog "k8s.io/klog/v2"
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
type ParseOpt func(obj interface{}, method string) interface{}

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

func WithFunction(ctx context.Context) Options {
	return func(i *Impl) {
		i.parse = func(obj interface{}, method string) interface{} {
			fn, ok := obj.(*fnV1beta1.Function)
			if !ok {
				return obj
			}

			if fn.Status.Build == nil {
				return nil
			}

			return Object{
				Method: method,
				FnSummary: &FnStatusSummary{
					ObjectMeta: fn.ObjectMeta,
					Status:     fn.Status,
				},
			}
		}
	}
}

func WithPipelineRun(ctx context.Context) Options {
	return func(i *Impl) {
		i.parse = func(obj interface{}, method string) interface{} {
			pr, ok := obj.(*prV1beta1.PipelineRun)
			if !ok {
				return obj
			}

			return Object{
				Method: method,
				PRSummary: &PRStatusSummary{
					ObjectMeta: pr.ObjectMeta,
					Status:     pr.Status,
				},
			}
		}
	}
}

type Impl struct {
	Name string

	informer cache.SharedIndexInformer
	cache    *lru.Cache
	queue    *queue
	parse    ParseOpt
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
	if i.parse != nil {
		obj = i.parse(obj, method)
		if obj == nil {
			return
		}
	}
	i.queue.Send(obj)
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
	Method    string
	FnSummary *FnStatusSummary
	PRSummary *PRStatusSummary
}

type FnStatusSummary struct {
	metav1.ObjectMeta
	Status fnV1beta1.FunctionStatus
}

type PRStatusSummary struct {
	metav1.ObjectMeta
	Status prV1beta1.PipelineRunStatus
}
