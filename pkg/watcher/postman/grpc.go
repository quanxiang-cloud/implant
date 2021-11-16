package postman

import (
	"context"
	"fmt"

	"github.com/quanxiang-cloud/implant/pkg/proto/v1alpha1"
	pb "github.com/quanxiang-cloud/implant/pkg/proto/v1alpha1"
	"github.com/quanxiang-cloud/implant/pkg/watcher/reconciler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	rs "google.golang.org/grpc/status"
	"k8s.io/klog"
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

		code := rs.Code(err)
		switch code {
		case codes.OK:
			return
		case codes.Unknown:
			klog.Error(err)
		default:
			klog.Error(err)
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

	resourceRef := make(map[string]*v1alpha1.RefCondition)
	for key, ref := range sm.Status.ResourceRef {
		conditions := make([]*v1alpha1.Condition, 0, len(ref.Conditions))
		for _, elem := range ref.Conditions {
			conditions = append(conditions, &v1alpha1.Condition{
				Status:             string(elem.Status),
				LastTransitionTime: elem.LastTransitionTime.String(),
				Message:            elem.Message,
				Reason:             elem.Reason,
			})
		}

		resourceRef[key] = &v1alpha1.RefCondition{
			Name:       ref.Name,
			Conditions: conditions,
		}
	}

	return objectMeta, pb.Status{
		Phase:              sm.Status.Phase.Sting(),
		Status:             string(sm.Status.Status),
		Reason:             sm.Status.Reason,
		Message:            sm.Status.Message,
		ResourceRef:        resourceRef,
		LastTransitionTime: sm.Status.LastTransitionTime.String(),
	}
}
