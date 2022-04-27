package watcher

import (
	"context"

	"github.com/quanxiang-cloud/cabin/tailormade/client"
	"github.com/quanxiang-cloud/implant/pkg/watcher/informers"
	"github.com/quanxiang-cloud/implant/pkg/watcher/postman"
	"github.com/quanxiang-cloud/implant/pkg/watcher/reconciler"
	"k8s.io/klog/v2"
)

type Watcher struct {
	CTX context.Context
	informers.Oper
}

type Client informers.Client

func NewWatcherWithOper(ctx context.Context, o informers.Oper) *Watcher {
	return &Watcher{
		CTX: ctx,
		Oper: informers.Oper{
			Namespace:     o.Namespace,
			DefaultResync: o.DefaultResync,
			Options:       o.Options,
		},
	}
}

func (w *Watcher) Opts(opts ...reconciler.Options) *Watcher {
	w.Oper.Opts(w.CTX, opts...)
	return w
}

func (w *Watcher) Concurrency(n int, cacheMaxEntries int) *Watcher {
	for i := 0; i < n; i++ {
		w.Oper.Opts(w.CTX, reconciler.WithCache(w.CTX, cacheMaxEntries))
	}
	return w
}

func (w *Watcher) Sender(c *client.Config, url string, errChan chan error) *Watcher {
	worker := postman.New(w.CTX, c, url)
	w.Oper.Opts(w.CTX, reconciler.WithConsumer(
		w.CTX,
		worker.Send(errChan),
	))
	return w
}

func (w *Watcher) Run(client Client) error {
	return w.Oper.Run(w.CTX, client)
}

func (w *Watcher) RunOrDie(client Client) {
	err := w.Oper.Run(w.CTX, client)
	if err != nil {
		klog.Error(err)
		panic(err)
	}
}
