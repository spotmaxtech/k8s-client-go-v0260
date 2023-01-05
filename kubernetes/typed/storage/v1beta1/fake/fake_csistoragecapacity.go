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

	v1beta1 "github.com/spotmaxtech/k8s-api-v0260/storage/v1beta1"
	v1 "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/apis/meta/v1"
	labels "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/labels"
	schema "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/runtime/schema"
	types "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/types"
	watch "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/watch"
	storagev1beta1 "github.com/spotmaxtech/k8s-client-go-v0260/applyconfigurations/storage/v1beta1"
	testing "github.com/spotmaxtech/k8s-client-go-v0260/testing"
)

// FakeCSIStorageCapacities implements CSIStorageCapacityInterface
type FakeCSIStorageCapacities struct {
	Fake *FakeStorageV1beta1
	ns   string
}

var csistoragecapacitiesResource = schema.GroupVersionResource{Group: "storage.k8s.io", Version: "v1beta1", Resource: "csistoragecapacities"}

var csistoragecapacitiesKind = schema.GroupVersionKind{Group: "storage.k8s.io", Version: "v1beta1", Kind: "CSIStorageCapacity"}

// Get takes name of the cSIStorageCapacity, and returns the corresponding cSIStorageCapacity object, and an error if there is any.
func (c *FakeCSIStorageCapacities) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.CSIStorageCapacity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(csistoragecapacitiesResource, c.ns, name), &v1beta1.CSIStorageCapacity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CSIStorageCapacity), err
}

// List takes label and field selectors, and returns the list of CSIStorageCapacities that match those selectors.
func (c *FakeCSIStorageCapacities) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.CSIStorageCapacityList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(csistoragecapacitiesResource, csistoragecapacitiesKind, c.ns, opts), &v1beta1.CSIStorageCapacityList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.CSIStorageCapacityList{ListMeta: obj.(*v1beta1.CSIStorageCapacityList).ListMeta}
	for _, item := range obj.(*v1beta1.CSIStorageCapacityList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cSIStorageCapacities.
func (c *FakeCSIStorageCapacities) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(csistoragecapacitiesResource, c.ns, opts))

}

// Create takes the representation of a cSIStorageCapacity and creates it.  Returns the server's representation of the cSIStorageCapacity, and an error, if there is any.
func (c *FakeCSIStorageCapacities) Create(ctx context.Context, cSIStorageCapacity *v1beta1.CSIStorageCapacity, opts v1.CreateOptions) (result *v1beta1.CSIStorageCapacity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(csistoragecapacitiesResource, c.ns, cSIStorageCapacity), &v1beta1.CSIStorageCapacity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CSIStorageCapacity), err
}

// Update takes the representation of a cSIStorageCapacity and updates it. Returns the server's representation of the cSIStorageCapacity, and an error, if there is any.
func (c *FakeCSIStorageCapacities) Update(ctx context.Context, cSIStorageCapacity *v1beta1.CSIStorageCapacity, opts v1.UpdateOptions) (result *v1beta1.CSIStorageCapacity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(csistoragecapacitiesResource, c.ns, cSIStorageCapacity), &v1beta1.CSIStorageCapacity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CSIStorageCapacity), err
}

// Delete takes name of the cSIStorageCapacity and deletes it. Returns an error if one occurs.
func (c *FakeCSIStorageCapacities) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(csistoragecapacitiesResource, c.ns, name, opts), &v1beta1.CSIStorageCapacity{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCSIStorageCapacities) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(csistoragecapacitiesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.CSIStorageCapacityList{})
	return err
}

// Patch applies the patch and returns the patched cSIStorageCapacity.
func (c *FakeCSIStorageCapacities) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.CSIStorageCapacity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(csistoragecapacitiesResource, c.ns, name, pt, data, subresources...), &v1beta1.CSIStorageCapacity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CSIStorageCapacity), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied cSIStorageCapacity.
func (c *FakeCSIStorageCapacities) Apply(ctx context.Context, cSIStorageCapacity *storagev1beta1.CSIStorageCapacityApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.CSIStorageCapacity, err error) {
	if cSIStorageCapacity == nil {
		return nil, fmt.Errorf("cSIStorageCapacity provided to Apply must not be nil")
	}
	data, err := json.Marshal(cSIStorageCapacity)
	if err != nil {
		return nil, err
	}
	name := cSIStorageCapacity.Name
	if name == nil {
		return nil, fmt.Errorf("cSIStorageCapacity.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(csistoragecapacitiesResource, c.ns, *name, types.ApplyPatchType, data), &v1beta1.CSIStorageCapacity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CSIStorageCapacity), err
}
