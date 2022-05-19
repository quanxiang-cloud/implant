package reconciler

import (
	"context"

	"k8s.io/client-go/tools/cache"
)

func NewControllerWithConfig(ctx context.Context, informer cache.SharedIndexInformer, opts ...Options) *Impl {
	impl := NewImpl(opts...)

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    impl.AddFunc,
		UpdateFunc: impl.UpdateFunc,
		// DeleteFunc: impl.DeleteFunc,
	})

	impl.informer = informer

	return impl
}
