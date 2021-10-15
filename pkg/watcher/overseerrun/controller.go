package overseerrun

import (
	"context"
	"time"

	"github.com/quanxiang-cloud/implant/pkg/watcher/reconciler"
	"github.com/quanxiang-cloud/overseer/pkg/client/clientset"
	"github.com/quanxiang-cloud/overseer/pkg/client/informers"
	overseerInformers "github.com/quanxiang-cloud/overseer/pkg/client/informers/overseer"
)

func NewController(ctx context.Context, client clientset.Interface, defaultResync time.Duration, opts ...reconciler.Options) *reconciler.Impl {
	return NewControllerWithConfig(ctx, client, defaultResync, opts...)
}

func NewControllerWithConfig(ctx context.Context, client clientset.Interface, defaultResync time.Duration, opts ...reconciler.Options) *reconciler.Impl {
	informer := overseerInformers.New(informers.NewSharedInformerFactory(client, defaultResync), nil).
		V1alpha1().OverseerRun().
		Informer()
	return reconciler.NewControllerWithConfig(ctx, informer, opts...)
}
