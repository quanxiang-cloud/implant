package overseerrun

import (
	"context"
	"time"

	"github.com/quanxiang-cloud/implant/pkg/watcher/reconciler"
	"github.com/quanxiang-cloud/overseer/pkg/client/clientset/versioned"
	"github.com/quanxiang-cloud/overseer/pkg/client/informers/externalversions"
	"github.com/quanxiang-cloud/overseer/pkg/client/informers/externalversions/overseer"
)

func NewController(ctx context.Context, client versioned.Interface, namespace string, defaultResync time.Duration, opts ...reconciler.Options) *reconciler.Impl {
	return NewControllerWithConfig(ctx, client, namespace, defaultResync, opts...)
}

func NewControllerWithConfig(ctx context.Context, client versioned.Interface, namespace string, defaultResync time.Duration, opts ...reconciler.Options) *reconciler.Impl {
	informer := overseer.New(externalversions.NewSharedInformerFactory(client, defaultResync), namespace, nil).
		V1alpha1().
		Overseers().
		Informer()
	return reconciler.NewControllerWithConfig(ctx, informer, opts...)
}
