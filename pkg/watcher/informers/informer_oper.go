package informers

import (
	"context"
	"fmt"
	"reflect"
	"time"

	fnClientset "github.com/openfunction/pkg/client/clientset/versioned"
	"github.com/quanxiang-cloud/implant/pkg/watcher/reconciler"
	tkClientset "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
)

type Client interface{}

var (
	informers = map[reflect.Type]f{
		reflect.TypeOf(&fnClientset.Clientset{}): NewFnControllerWithConfig,
		reflect.TypeOf(&tkClientset.Clientset{}): NewPipelineControllerWithConfig,
	}
)

type f func(ctx context.Context, client Client, namespace string, defaultResync time.Duration, opts ...reconciler.Options) *reconciler.Impl

type Oper struct {
	Options       []reconciler.Options
	Namespace     string
	DefaultResync time.Duration
}

func (w *Oper) Opts(ctx context.Context, opts ...reconciler.Options) {
	if w.Options == nil {
		w.Options = make([]reconciler.Options, 0)
	}

	w.Options = append(w.Options, opts...)
}

func (w *Oper) Run(ctx context.Context, client Client) error {
	t := reflect.TypeOf(client)
	f, ok := informers[t]
	if ok {
		impl := f(ctx, client, w.Namespace, w.DefaultResync, w.Options...)
		go impl.Run(ctx.Done())
		return nil
	}
	return fmt.Errorf("the type of (%v) is not supported", client)
}
