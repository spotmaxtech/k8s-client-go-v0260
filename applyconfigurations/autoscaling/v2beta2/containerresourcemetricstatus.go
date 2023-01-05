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

package v2beta2

import (
	v1 "github.com/spotmaxtech/k8s-api-v0260/core/v1"
)

// ContainerResourceMetricStatusApplyConfiguration represents an declarative configuration of the ContainerResourceMetricStatus type for use
// with apply.
type ContainerResourceMetricStatusApplyConfiguration struct {
	Name      *v1.ResourceName                     `json:"name,omitempty"`
	Current   *MetricValueStatusApplyConfiguration `json:"current,omitempty"`
	Container *string                              `json:"container,omitempty"`
}

// ContainerResourceMetricStatusApplyConfiguration constructs an declarative configuration of the ContainerResourceMetricStatus type for use with
// apply.
func ContainerResourceMetricStatus() *ContainerResourceMetricStatusApplyConfiguration {
	return &ContainerResourceMetricStatusApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ContainerResourceMetricStatusApplyConfiguration) WithName(value v1.ResourceName) *ContainerResourceMetricStatusApplyConfiguration {
	b.Name = &value
	return b
}

// WithCurrent sets the Current field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Current field is set to the value of the last call.
func (b *ContainerResourceMetricStatusApplyConfiguration) WithCurrent(value *MetricValueStatusApplyConfiguration) *ContainerResourceMetricStatusApplyConfiguration {
	b.Current = value
	return b
}

// WithContainer sets the Container field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Container field is set to the value of the last call.
func (b *ContainerResourceMetricStatusApplyConfiguration) WithContainer(value string) *ContainerResourceMetricStatusApplyConfiguration {
	b.Container = &value
	return b
}
