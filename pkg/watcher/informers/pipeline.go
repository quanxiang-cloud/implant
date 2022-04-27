package informers

import (
	"context"
	"time"

	"github.com/quanxiang-cloud/implant/pkg/watcher/reconciler"
	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	"github.com/tektoncd/pipeline/pkg/client/informers/externalversions"
	"github.com/tektoncd/pipeline/pkg/client/informers/externalversions/pipeline"
)

func NewPipelineControllerWithConfig(ctx context.Context, client Client, namespace string, defaultResync time.Duration, opts ...reconciler.Options) *reconciler.Impl {
	informer := pipeline.New(externalversions.NewSharedInformerFactory(client.(versioned.Interface), defaultResync), namespace, nil).
		V1beta1().
		PipelineRuns().
		Informer()
	return reconciler.NewControllerWithConfig(ctx, informer, opts...)
}
