package informers

import (
	"context"
	"time"

	"github.com/openfunction/pkg/client/clientset/versioned"
	"github.com/quanxiang-cloud/implant/pkg/client/informers/externalversions"
	"github.com/quanxiang-cloud/implant/pkg/client/informers/externalversions/core"
	"github.com/quanxiang-cloud/implant/pkg/watcher/reconciler"
)

func NewFnController(ctx context.Context, client versioned.Interface, namespace string, defaultResync time.Duration, opts ...reconciler.Options) *reconciler.Impl {
	return NewFnControllerWithConfig(ctx, client, namespace, defaultResync, opts...)
}

func NewFnControllerWithConfig(ctx context.Context, client versioned.Interface, namespace string, defaultResync time.Duration, opts ...reconciler.Options) *reconciler.Impl {
	informer := core.New(externalversions.NewSharedInformerFactory(client, defaultResync), namespace, nil).
		V1beta1().
		Functions().
		Informer()
	return reconciler.NewControllerWithConfig(ctx, informer, opts...)
}
