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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	nodev1 "k8s.io/api/node/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationsnodev1 "github.com/spotmaxtech/k8s-client-go-v0260/applyconfigurations/node/v1"
	testing "github.com/spotmaxtech/k8s-client-go-v0260/testing"
)

// FakeRuntimeClasses implements RuntimeClassInterface
type FakeRuntimeClasses struct {
	Fake *FakeNodeV1
}

var runtimeclassesResource = schema.GroupVersionResource{Group: "node.k8s.io", Version: "v1", Resource: "runtimeclasses"}

var runtimeclassesKind = schema.GroupVersionKind{Group: "node.k8s.io", Version: "v1", Kind: "RuntimeClass"}

// Get takes name of the runtimeClass, and returns the corresponding runtimeClass object, and an error if there is any.
func (c *FakeRuntimeClasses) Get(ctx context.Context, name string, options v1.GetOptions) (result *nodev1.RuntimeClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(runtimeclassesResource, name), &nodev1.RuntimeClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*nodev1.RuntimeClass), err
}

// List takes label and field selectors, and returns the list of RuntimeClasses that match those selectors.
func (c *FakeRuntimeClasses) List(ctx context.Context, opts v1.ListOptions) (result *nodev1.RuntimeClassList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(runtimeclassesResource, runtimeclassesKind, opts), &nodev1.RuntimeClassList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &nodev1.RuntimeClassList{ListMeta: obj.(*nodev1.RuntimeClassList).ListMeta}
	for _, item := range obj.(*nodev1.RuntimeClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested runtimeClasses.
func (c *FakeRuntimeClasses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(runtimeclassesResource, opts))
}

// Create takes the representation of a runtimeClass and creates it.  Returns the server's representation of the runtimeClass, and an error, if there is any.
func (c *FakeRuntimeClasses) Create(ctx context.Context, runtimeClass *nodev1.RuntimeClass, opts v1.CreateOptions) (result *nodev1.RuntimeClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(runtimeclassesResource, runtimeClass), &nodev1.RuntimeClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*nodev1.RuntimeClass), err
}

// Update takes the representation of a runtimeClass and updates it. Returns the server's representation of the runtimeClass, and an error, if there is any.
func (c *FakeRuntimeClasses) Update(ctx context.Context, runtimeClass *nodev1.RuntimeClass, opts v1.UpdateOptions) (result *nodev1.RuntimeClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(runtimeclassesResource, runtimeClass), &nodev1.RuntimeClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*nodev1.RuntimeClass), err
}

// Delete takes name of the runtimeClass and deletes it. Returns an error if one occurs.
func (c *FakeRuntimeClasses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(runtimeclassesResource, name, opts), &nodev1.RuntimeClass{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRuntimeClasses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(runtimeclassesResource, listOpts)

	_, err := c.Fake.Invokes(action, &nodev1.RuntimeClassList{})
	return err
}

// Patch applies the patch and returns the patched runtimeClass.
func (c *FakeRuntimeClasses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *nodev1.RuntimeClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(runtimeclassesResource, name, pt, data, subresources...), &nodev1.RuntimeClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*nodev1.RuntimeClass), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied runtimeClass.
func (c *FakeRuntimeClasses) Apply(ctx context.Context, runtimeClass *applyconfigurationsnodev1.RuntimeClassApplyConfiguration, opts v1.ApplyOptions) (result *nodev1.RuntimeClass, err error) {
	if runtimeClass == nil {
		return nil, fmt.Errorf("runtimeClass provided to Apply must not be nil")
	}
	data, err := json.Marshal(runtimeClass)
	if err != nil {
		return nil, err
	}
	name := runtimeClass.Name
	if name == nil {
		return nil, fmt.Errorf("runtimeClass.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(runtimeclassesResource, *name, types.ApplyPatchType, data), &nodev1.RuntimeClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*nodev1.RuntimeClass), err
}
