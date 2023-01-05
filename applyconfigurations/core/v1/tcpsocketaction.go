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
	intstr "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/util/intstr"
)

// TCPSocketActionApplyConfiguration represents an declarative configuration of the TCPSocketAction type for use
// with apply.
type TCPSocketActionApplyConfiguration struct {
	Port *intstr.IntOrString `json:"port,omitempty"`
	Host *string             `json:"host,omitempty"`
}

// TCPSocketActionApplyConfiguration constructs an declarative configuration of the TCPSocketAction type for use with
// apply.
func TCPSocketAction() *TCPSocketActionApplyConfiguration {
	return &TCPSocketActionApplyConfiguration{}
}

// WithPort sets the Port field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Port field is set to the value of the last call.
func (b *TCPSocketActionApplyConfiguration) WithPort(value intstr.IntOrString) *TCPSocketActionApplyConfiguration {
	b.Port = &value
	return b
}

// WithHost sets the Host field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Host field is set to the value of the last call.
func (b *TCPSocketActionApplyConfiguration) WithHost(value string) *TCPSocketActionApplyConfiguration {
	b.Host = &value
	return b
}
