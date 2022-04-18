package openfunction

import (
	"context"
	"time"

	"github.com/openfunction/pkg/client/clientset/versioned"
	"github.com/quanxiang-cloud/implant/pkg/client/informers/externalversions"
	"github.com/quanxiang-cloud/implant/pkg/client/informers/externalversions/core"
	"github.com/quanxiang-cloud/implant/pkg/watcher/reconciler"
)

func NewController(ctx context.Context, client versioned.Interface, namespace string, defaultResync time.Duration, opts ...reconciler.Options) *reconciler.Impl {
	return NewControllerWithConfig(ctx, client, namespace, defaultResync, opts...)
}

func NewControllerWithConfig(ctx context.Context, client versioned.Interface, namespace string, defaultResync time.Duration, opts ...reconciler.Options) *reconciler.Impl {
	informer := core.New(externalversions.NewSharedInformerFactory(client, defaultResync), namespace, nil).
		V1beta1().
		Functions().
		Informer()
	return reconciler.NewControllerWithConfig(ctx, informer, opts...)
}
