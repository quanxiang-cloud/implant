package postman

import (
	"context"
	"fmt"

	"github.com/quanxiang-cloud/cabin/tailormade/client"
	bc "github.com/quanxiang-cloud/implant/pkg/broadcaster/v1beta1"
	"github.com/quanxiang-cloud/implant/pkg/watcher/reconciler"
	"k8s.io/klog/v2"
)

func New(ctx context.Context, c *client.Config, target string) (*Sender, error) {
	return &Sender{
		client: bc.New(c, target),
	}, nil
}

type Sender struct {
	client *bc.Client
}

func (s *Sender) SendFN(e chan<- error) func(obj interface{}) {
	return func(obj interface{}) {
		ro, ok := obj.(reconciler.Object)
		if !ok {
			e <- fmt.Errorf("unknown obj type")
		}

		objFn := serializeObj(*ro.Summary)

		ctx := context.Background()
		_, err := s.client.Send(ctx, objFn)

		if err != nil {
			klog.Error(err)
			e <- err
		}
	}
}

func serializeObj(sm reconciler.StatusSummary) bc.Function {
	return bc.Function{
		Name:        sm.Name,
		ResourceRef: sm.Status.Build.ResourceRef,
		State:       sm.Status.Build.State,
		Topic:       "build",
	}
}
