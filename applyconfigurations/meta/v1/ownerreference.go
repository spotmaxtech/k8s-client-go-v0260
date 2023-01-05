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
	types "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/types"
)

// OwnerReferenceApplyConfiguration represents an declarative configuration of the OwnerReference type for use
// with apply.
type OwnerReferenceApplyConfiguration struct {
	APIVersion         *string    `json:"apiVersion,omitempty"`
	Kind               *string    `json:"kind,omitempty"`
	Name               *string    `json:"name,omitempty"`
	UID                *types.UID `json:"uid,omitempty"`
	Controller         *bool      `json:"controller,omitempty"`
	BlockOwnerDeletion *bool      `json:"blockOwnerDeletion,omitempty"`
}

// OwnerReferenceApplyConfiguration constructs an declarative configuration of the OwnerReference type for use with
// apply.
func OwnerReference() *OwnerReferenceApplyConfiguration {
	return &OwnerReferenceApplyConfiguration{}
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *OwnerReferenceApplyConfiguration) WithAPIVersion(value string) *OwnerReferenceApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *OwnerReferenceApplyConfiguration) WithKind(value string) *OwnerReferenceApplyConfiguration {
	b.Kind = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *OwnerReferenceApplyConfiguration) WithName(value string) *OwnerReferenceApplyConfiguration {
	b.Name = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *OwnerReferenceApplyConfiguration) WithUID(value types.UID) *OwnerReferenceApplyConfiguration {
	b.UID = &value
	return b
}

// WithController sets the Controller field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Controller field is set to the value of the last call.
func (b *OwnerReferenceApplyConfiguration) WithController(value bool) *OwnerReferenceApplyConfiguration {
	b.Controller = &value
	return b
}

// WithBlockOwnerDeletion sets the BlockOwnerDeletion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BlockOwnerDeletion field is set to the value of the last call.
func (b *OwnerReferenceApplyConfiguration) WithBlockOwnerDeletion(value bool) *OwnerReferenceApplyConfiguration {
	b.BlockOwnerDeletion = &value
	return b
}
