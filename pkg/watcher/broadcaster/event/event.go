package event

import (
	fnV1beta1 "github.com/openfunction/apis/core/v1beta1"
	prV1beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Message struct {
	*FnMessage
	*PrMessage
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

type Data struct {
	*FnStatusSummary
	*PRStatusSummary
}

type FnStatusSummary struct {
	metav1.ObjectMeta
	Status fnV1beta1.FunctionStatus
}

type PRStatusSummary struct {
	metav1.ObjectMeta
	Status prV1beta1.PipelineRunStatus
}
