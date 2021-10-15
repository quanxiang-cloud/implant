package postman

import (
	"context"
	"fmt"

	pb "github.com/quanxiang-cloud/implant/pkg/proto/v1alpha1"
	"github.com/quanxiang-cloud/implant/pkg/watcher/reconciler"
	"google.golang.org/grpc"
)

func New(ctx context.Context, target string) (*Sender, error) {
	client, err := connet(ctx, target)
	if err != nil {
		return nil, err
	}

	return &Sender{
		client: client,
	}, nil
}

func connet(ctx context.Context, target string) (pb.ImplantClient, error) {
	conn, err := grpc.DialContext(ctx, target,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)

	if err != nil {
		return nil, err
	}

	return pb.NewImplantClient(conn), nil
}

type Sender struct {
	client pb.ImplantClient
}

func (s *Sender) SendFN(e chan<- error) func(obj interface{}) {
	return func(obj interface{}) {
		ro, ok := obj.(reconciler.Object)
		if !ok {
			e <- fmt.Errorf("unknown obj type")
		}

		objectmeta, status := serializeObj(*ro.Summary)

		var err error

		ctx := context.Background()
		switch ro.Method {
		case reconciler.ADD:
			_, err = s.client.Create(ctx, &pb.CreateReq{
				ObjectMeta: &objectmeta,
				State:      &status,
			})
		case reconciler.UPDATE:
			_, err = s.client.Update(ctx, &pb.UpdateReq{
				ObjectMeta: &objectmeta,
				State:      &status,
			})
		case reconciler.DELETE:
			_, err = s.client.Delete(ctx, &pb.DeleteReq{
				ObjectMeta: &objectmeta,
				State:      &status,
			})
		}

		if err != nil {
			e <- err
		}
	}

}

func serializeObj(sm reconciler.StatusSummary) (pb.ObjectMeta, pb.Status) {
	objectMeta := pb.ObjectMeta{
		Name:           sm.Name,
		Namepsace:      sm.Namespace,
		Uid:            string(sm.UID),
		CreationTime:   sm.CreationTimestamp.String(),
		Labels:         sm.Labels,
		OwnerReference: make([]*pb.OwnerReference, 0, len(sm.OwnerReferences)),
	}

	for _, of := range sm.OwnerReferences {
		objectMeta.OwnerReference = append(objectMeta.OwnerReference,
			&pb.OwnerReference{
				Name: of.Name,
				Uid:  string(of.UID),
			})
	}

	condition := &pb.Condition{
		Status:             string(sm.Status.Condition.Status),
		Reason:             sm.Status.Condition.Reason,
		Message:            sm.Status.Condition.Message,
		ResourceRef:        make(map[string]*pb.StepCondition, len(sm.Status.Condition.ResourceRef)),
		LastTransitionTime: sm.Status.Condition.LastTransitionTime.String(),
	}

	for key, sc := range sm.Status.Condition.ResourceRef {
		condition.ResourceRef[key] = &pb.StepCondition{
			State:   string(sc.State),
			Reason:  sc.Reason,
			Message: sc.Message,
		}
	}

	return objectMeta, pb.Status{
		Conditions: condition,
	}
}
