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

package v1alpha1

import (
	v1 "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/apis/meta/v1"
)

// ValidationApplyConfiguration represents an declarative configuration of the Validation type for use
// with apply.
type ValidationApplyConfiguration struct {
	Expression *string          `json:"expression,omitempty"`
	Message    *string          `json:"message,omitempty"`
	Reason     *v1.StatusReason `json:"reason,omitempty"`
}

// ValidationApplyConfiguration constructs an declarative configuration of the Validation type for use with
// apply.
func Validation() *ValidationApplyConfiguration {
	return &ValidationApplyConfiguration{}
}

// WithExpression sets the Expression field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Expression field is set to the value of the last call.
func (b *ValidationApplyConfiguration) WithExpression(value string) *ValidationApplyConfiguration {
	b.Expression = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *ValidationApplyConfiguration) WithMessage(value string) *ValidationApplyConfiguration {
	b.Message = &value
	return b
}

// WithReason sets the Reason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reason field is set to the value of the last call.
func (b *ValidationApplyConfiguration) WithReason(value v1.StatusReason) *ValidationApplyConfiguration {
	b.Reason = &value
	return b
}
