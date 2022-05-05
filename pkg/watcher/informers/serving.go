package informers

import (
	"context"
	"time"

	"github.com/quanxiang-cloud/implant/pkg/watcher/reconciler"
	"knative.dev/serving/pkg/client/clientset/versioned"
	"knative.dev/serving/pkg/client/informers/externalversions"
	"knative.dev/serving/pkg/client/informers/externalversions/serving"
)

func NewServingControllerWithConfig(ctx context.Context, client Client, namespace string, defaultResync time.Duration, opts ...reconciler.Options) *reconciler.Impl {
	informer := serving.New(externalversions.NewSharedInformerFactory(client.(versioned.Interface), defaultResync), namespace, nil).
		V1().
		Services().
		Informer()
	return reconciler.NewControllerWithConfig(ctx, informer, opts...)
}
