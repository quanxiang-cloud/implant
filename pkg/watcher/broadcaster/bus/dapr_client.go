package bus

import (
	"context"
	"fmt"

	daprd "github.com/dapr/go-sdk/client"
	"github.com/quanxiang-cloud/implant/pkg/watcher/broadcaster/event"
	"k8s.io/klog/v2"
)

type EventBus struct {
	daprClient daprd.Client

	pubsubName string
	errChan    chan error
}

func NewDaprClient(ctx context.Context, errChan chan error, opts ...Option) (*EventBus, error) {
	client, err := daprd.NewClient()
	if err != nil {
		return nil, err
	}

	bus := &EventBus{
		daprClient: client,
	}

	for _, fn := range opts {
		if err := fn(bus); err != nil {
			return nil, err
		}
	}

	return bus, nil
}

type Option func(*EventBus) error

func WithPubsubName(pubsubName string) Option {
	return func(eb *EventBus) error {
		eb.pubsubName = pubsubName
		return nil
	}
}

func (b *EventBus) Send(ctx context.Context, obj interface{}) {
	data, ok := obj.(*event.Data)
	if !ok {
		b.errChan <- fmt.Errorf("unknown obj type")
	}
	err := b.sendMessage(ctx, data)
	if err != nil {
		klog.Error(err)
		b.errChan <- err
	}
}

func (b *EventBus) sendMessage(ctx context.Context, req *event.Data) error {
	// TODO: remove
	var topic string = "lowcode.faas"
	msg := b.serialize(req)
	// TODO: check msg
	if err := b.publish(ctx, topic, msg); err != nil {
		klog.Error(err)
		return err
	}

	return nil
}

func (b *EventBus) serialize(data *event.Data) *event.Message {
	msg := &event.Message{}
	if data.FnStatusSummary != nil {
		klog.Info("serialize funtions...")
		msg.Fn = b.serializeFn(data)
	}
	if data.PRStatusSummary != nil {
		klog.Info("serialize pipelineRun...")
		msg.Pr = b.serializePr(data)
	}
	if data.SvcStatusSummary != nil {
		klog.Info("serialize ksvc...")
		msg.Svc = b.serializeSvc(data)
	}
	return msg
}

func (b *EventBus) serializeFn(data *event.Data) *event.FnMessage {
	return &event.FnMessage{
		Name:        data.FnStatusSummary.Name,
		Topic:       data.FnStatusSummary.Namespace,
		State:       data.FnStatusSummary.Status.Build.State,
		ResourceRef: data.FnStatusSummary.Status.Build.ResourceRef,
	}
}

func (b *EventBus) serializePr(data *event.Data) *event.PrMessage {
	var state string = ""
	l := len(data.PRStatusSummary.Status.Conditions)
	if l > 0 {
		state = string(data.PRStatusSummary.Status.Conditions[l-1].Status)
	}

	return &event.PrMessage{
		Name:  data.PRStatusSummary.Name,
		Topic: data.PRStatusSummary.Namespace,
		State: state,
	}
}

func (b *EventBus) serializeSvc(data *event.Data) *event.SvcMessage {
	var state string = ""
	l := len(data.SvcStatusSummary.Status.Conditions)
	if l > 0 {
		state = string(data.SvcStatusSummary.Status.Conditions[l-1].Status)
	}
	return &event.SvcMessage{
		Name:  data.SvcStatusSummary.Name,
		Topic: data.SvcStatusSummary.Namespace,
		State: state,
	}
}

func (b *EventBus) publish(ctx context.Context, topic string, data interface{}) error {
	klog.Info("send message ", " topic ", topic)
	if err := b.daprClient.PublishEvent(ctx, b.pubsubName, topic, data); err != nil {
		klog.Error(err, "publishEvent", "topic", topic, "pubsubName", b.pubsubName)
		return err
	}
	return nil
}

func (b *EventBus) Close() error {
	b.daprClient.Close()
	return nil
}
