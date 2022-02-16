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
	for _, ref := range sm.Status.VersatileStatus {
		conditions := make([]*v1alpha1.Condition, 0, len(ref.Conditions))
		for _, elem := range ref.Conditions {
			conditions = append(conditions, &v1alpha1.Condition{
				Status:             string(elem.Status),
				LastTransitionTime: elem.LastTransitionTime.String(),
				Message:            elem.Message,
				Reason:             elem.Reason,
			})
		}

		resourceRef[ref.Ref] = &v1alpha1.RefCondition{
			Name:       ref.Ref,
			Conditions: conditions,
		}
	}
	pipelineRuns := &v1alpha1.PipelineRun{
		Status: string(sm.Status.PipelineRuns.Status),
	}

	originTaskRun := sm.Status.PipelineRuns.TaskRuns
	taskruns := make([]*v1alpha1.TaskRun, 0, len(originTaskRun))
	for _, taskRun := range originTaskRun {
		steps := make([]*v1alpha1.Step, 0, len(taskRun.Steps))
		for _, step := range taskRun.Steps {
			steps = append(steps, &v1alpha1.Step{
				Name:           step.Name,
				StartTime:      step.StartTime.String(),
				CompletionTime: step.CompletionTime.String(),
			})
		}

		taskruns = append(taskruns, &v1alpha1.TaskRun{
			TaskName:       taskRun.Name,
			Status:         string(taskRun.Status),
			StartTime:      taskRun.StartTime.String(),
			CompletionTime: taskRun.CompletionTime.String(),
			Step:           steps,
		})
	}

	pipelineRuns.TaskRun = append(pipelineRuns.TaskRun, taskruns...)
	return objectMeta, pb.Status{
		Status:         string(sm.Status.Status),
		ResourceRef:    resourceRef,
		CompletionTime: sm.Status.CompletionTime.String(),
		PipelineRun:    pipelineRuns,
	}
}
