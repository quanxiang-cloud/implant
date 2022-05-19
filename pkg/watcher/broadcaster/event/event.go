package event

import (
	fnV1beta1 "github.com/openfunction/apis/core/v1beta1"
	prV1beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	knV1 "knative.dev/serving/pkg/apis/serving/v1"
)

type Message struct {
	Fn  *FnMessage  `json:"fn"`
	Pr  *PrMessage  `json:"pr"`
	Svc *SvcMessage `json:"svc"`
}

type FnMessage struct {
	Name        string `json:"name,omitempty"`
	Topic       string `json:"topic,omitempty"`
	State       string `json:"state,omitempty"`
	ResourceRef string `json:"resource_ref,omitempty"`
}

type PrMessage struct {
	Name  string `json:"name,omitempty"`
	Topic string `json:"topic,omitempty"`
	State string `json:"state,omitempty"`
}

type SvcMessage struct {
	Name    string `json:"name,omitempty"`
	Topic   string `json:"topic,omitempty"`
	State   string `json:"state,omitempty"`
	Message string `json:"message,omitempty"`
}

type Data struct {
	*FnStatusSummary
	*PRStatusSummary
	*SvcStatusSummary
}

type FnStatusSummary struct {
	metav1.ObjectMeta
	Status fnV1beta1.FunctionStatus
}

type PRStatusSummary struct {
	metav1.ObjectMeta
	Status prV1beta1.PipelineRunStatus
}

type SvcStatusSummary struct {
	metav1.ObjectMeta
	Status knV1.ServiceStatus
}
