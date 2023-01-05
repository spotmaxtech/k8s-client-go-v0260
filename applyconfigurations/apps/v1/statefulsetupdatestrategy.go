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
	v1 "github.com/spotmaxtech/k8s-api-v0260/apps/v1"
)

// StatefulSetUpdateStrategyApplyConfiguration represents an declarative configuration of the StatefulSetUpdateStrategy type for use
// with apply.
type StatefulSetUpdateStrategyApplyConfiguration struct {
	Type          *v1.StatefulSetUpdateStrategyType                   `json:"type,omitempty"`
	RollingUpdate *RollingUpdateStatefulSetStrategyApplyConfiguration `json:"rollingUpdate,omitempty"`
}

// StatefulSetUpdateStrategyApplyConfiguration constructs an declarative configuration of the StatefulSetUpdateStrategy type for use with
// apply.
func StatefulSetUpdateStrategy() *StatefulSetUpdateStrategyApplyConfiguration {
	return &StatefulSetUpdateStrategyApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *StatefulSetUpdateStrategyApplyConfiguration) WithType(value v1.StatefulSetUpdateStrategyType) *StatefulSetUpdateStrategyApplyConfiguration {
	b.Type = &value
	return b
}

// WithRollingUpdate sets the RollingUpdate field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RollingUpdate field is set to the value of the last call.
func (b *StatefulSetUpdateStrategyApplyConfiguration) WithRollingUpdate(value *RollingUpdateStatefulSetStrategyApplyConfiguration) *StatefulSetUpdateStrategyApplyConfiguration {
	b.RollingUpdate = value
	return b
}
