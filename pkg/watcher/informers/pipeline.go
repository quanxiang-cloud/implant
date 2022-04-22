package informers

import (
	"context"
	"time"

	"github.com/quanxiang-cloud/implant/pkg/watcher/reconciler"
	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	"github.com/tektoncd/pipeline/pkg/client/informers/externalversions"
	"github.com/tektoncd/pipeline/pkg/client/informers/externalversions/pipeline"
)

func NewPipelineController(ctx context.Context, client versioned.Interface, namespace string, defaultResync time.Duration, opts ...reconciler.Options) *reconciler.Impl {
	return NewPipelineControllerWithConfig(ctx, client, namespace, defaultResync, opts...)
}

func NewPipelineControllerWithConfig(ctx context.Context, client versioned.Interface, namespace string, defaultResync time.Duration, opts ...reconciler.Options) *reconciler.Impl {
	informer := pipeline.New(externalversions.NewSharedInformerFactory(client, defaultResync), namespace, nil).
		V1beta1().
		PipelineRuns().
		Informer()
	return reconciler.NewControllerWithConfig(ctx, informer, opts...)
}
