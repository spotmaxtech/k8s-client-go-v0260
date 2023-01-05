/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v2beta1

import (
	v2beta1 "github.com/spotmaxtech/k8s-api-v0260/autoscaling/v2beta1"
	v1 "github.com/spotmaxtech/k8s-api-v0260/core/v1"
	metav1 "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/apis/meta/v1"
)

// HorizontalPodAutoscalerConditionApplyConfiguration represents an declarative configuration of the HorizontalPodAutoscalerCondition type for use
// with apply.
type HorizontalPodAutoscalerConditionApplyConfiguration struct {
	Type               *v2beta1.HorizontalPodAutoscalerConditionType `json:"type,omitempty"`
	Status             *v1.ConditionStatus                           `json:"status,omitempty"`
	LastTransitionTime *metav1.Time                                  `json:"lastTransitionTime,omitempty"`
	Reason             *string                                       `json:"reason,omitempty"`
	Message            *string                                       `json:"message,omitempty"`
}

// HorizontalPodAutoscalerConditionApplyConfiguration constructs an declarative configuration of the HorizontalPodAutoscalerCondition type for use with
// apply.
func HorizontalPodAutoscalerCondition() *HorizontalPodAutoscalerConditionApplyConfiguration {
	return &HorizontalPodAutoscalerConditionApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *HorizontalPodAutoscalerConditionApplyConfiguration) WithType(value v2beta1.HorizontalPodAutoscalerConditionType) *HorizontalPodAutoscalerConditionApplyConfiguration {
	b.Type = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *HorizontalPodAutoscalerConditionApplyConfiguration) WithStatus(value v1.ConditionStatus) *HorizontalPodAutoscalerConditionApplyConfiguration {
	b.Status = &value
	return b
}

// WithLastTransitionTime sets the LastTransitionTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastTransitionTime field is set to the value of the last call.
func (b *HorizontalPodAutoscalerConditionApplyConfiguration) WithLastTransitionTime(value metav1.Time) *HorizontalPodAutoscalerConditionApplyConfiguration {
	b.LastTransitionTime = &value
	return b
}

// WithReason sets the Reason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reason field is set to the value of the last call.
func (b *HorizontalPodAutoscalerConditionApplyConfiguration) WithReason(value string) *HorizontalPodAutoscalerConditionApplyConfiguration {
	b.Reason = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *HorizontalPodAutoscalerConditionApplyConfiguration) WithMessage(value string) *HorizontalPodAutoscalerConditionApplyConfiguration {
	b.Message = &value
	return b
}
