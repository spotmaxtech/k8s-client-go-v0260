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

package v1

import (
	v1 "github.com/spotmaxtech/k8s-api-v0260/core/v1"
)

// TolerationApplyConfiguration represents an declarative configuration of the Toleration type for use
// with apply.
type TolerationApplyConfiguration struct {
	Key               *string                `json:"key,omitempty"`
	Operator          *v1.TolerationOperator `json:"operator,omitempty"`
	Value             *string                `json:"value,omitempty"`
	Effect            *v1.TaintEffect        `json:"effect,omitempty"`
	TolerationSeconds *int64                 `json:"tolerationSeconds,omitempty"`
}

// TolerationApplyConfiguration constructs an declarative configuration of the Toleration type for use with
// apply.
func Toleration() *TolerationApplyConfiguration {
	return &TolerationApplyConfiguration{}
}

// WithKey sets the Key field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Key field is set to the value of the last call.
func (b *TolerationApplyConfiguration) WithKey(value string) *TolerationApplyConfiguration {
	b.Key = &value
	return b
}

// WithOperator sets the Operator field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Operator field is set to the value of the last call.
func (b *TolerationApplyConfiguration) WithOperator(value v1.TolerationOperator) *TolerationApplyConfiguration {
	b.Operator = &value
	return b
}

// WithValue sets the Value field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Value field is set to the value of the last call.
func (b *TolerationApplyConfiguration) WithValue(value string) *TolerationApplyConfiguration {
	b.Value = &value
	return b
}

// WithEffect sets the Effect field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Effect field is set to the value of the last call.
func (b *TolerationApplyConfiguration) WithEffect(value v1.TaintEffect) *TolerationApplyConfiguration {
	b.Effect = &value
	return b
}

// WithTolerationSeconds sets the TolerationSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TolerationSeconds field is set to the value of the last call.
func (b *TolerationApplyConfiguration) WithTolerationSeconds(value int64) *TolerationApplyConfiguration {
	b.TolerationSeconds = &value
	return b
}
