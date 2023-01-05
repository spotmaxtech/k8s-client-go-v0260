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

	corev1 "github.com/spotmaxtech/k8s-api-v0260/core/v1"
	v1 "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/apis/meta/v1"
	labels "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/labels"
	schema "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/runtime/schema"
	types "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/types"
	watch "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/watch"
	applyconfigurationscorev1 "github.com/spotmaxtech/k8s-client-go-v0260/applyconfigurations/core/v1"
	testing "github.com/spotmaxtech/k8s-client-go-v0260/testing"
)

// FakeEndpoints implements EndpointsInterface
type FakeEndpoints struct {
	Fake *FakeCoreV1
	ns   string
}

var endpointsResource = schema.GroupVersionResource{Group: "", Version: "v1", Resource: "endpoints"}

var endpointsKind = schema.GroupVersionKind{Group: "", Version: "v1", Kind: "Endpoints"}

// Get takes name of the endpoints, and returns the corresponding endpoints object, and an error if there is any.
func (c *FakeEndpoints) Get(ctx context.Context, name string, options v1.GetOptions) (result *corev1.Endpoints, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(endpointsResource, c.ns, name), &corev1.Endpoints{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Endpoints), err
}

// List takes label and field selectors, and returns the list of Endpoints that match those selectors.
func (c *FakeEndpoints) List(ctx context.Context, opts v1.ListOptions) (result *corev1.EndpointsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(endpointsResource, endpointsKind, c.ns, opts), &corev1.EndpointsList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.EndpointsList{ListMeta: obj.(*corev1.EndpointsList).ListMeta}
	for _, item := range obj.(*corev1.EndpointsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested endpoints.
func (c *FakeEndpoints) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(endpointsResource, c.ns, opts))

}

// Create takes the representation of a endpoints and creates it.  Returns the server's representation of the endpoints, and an error, if there is any.
func (c *FakeEndpoints) Create(ctx context.Context, endpoints *corev1.Endpoints, opts v1.CreateOptions) (result *corev1.Endpoints, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(endpointsResource, c.ns, endpoints), &corev1.Endpoints{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Endpoints), err
}

// Update takes the representation of a endpoints and updates it. Returns the server's representation of the endpoints, and an error, if there is any.
func (c *FakeEndpoints) Update(ctx context.Context, endpoints *corev1.Endpoints, opts v1.UpdateOptions) (result *corev1.Endpoints, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(endpointsResource, c.ns, endpoints), &corev1.Endpoints{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Endpoints), err
}

// Delete takes name of the endpoints and deletes it. Returns an error if one occurs.
func (c *FakeEndpoints) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(endpointsResource, c.ns, name, opts), &corev1.Endpoints{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEndpoints) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(endpointsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &corev1.EndpointsList{})
	return err
}

// Patch applies the patch and returns the patched endpoints.
func (c *FakeEndpoints) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *corev1.Endpoints, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(endpointsResource, c.ns, name, pt, data, subresources...), &corev1.Endpoints{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Endpoints), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied endpoints.
func (c *FakeEndpoints) Apply(ctx context.Context, endpoints *applyconfigurationscorev1.EndpointsApplyConfiguration, opts v1.ApplyOptions) (result *corev1.Endpoints, err error) {
	if endpoints == nil {
		return nil, fmt.Errorf("endpoints provided to Apply must not be nil")
	}
	data, err := json.Marshal(endpoints)
	if err != nil {
		return nil, err
	}
	name := endpoints.Name
	if name == nil {
		return nil, fmt.Errorf("endpoints.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(endpointsResource, c.ns, *name, types.ApplyPatchType, data), &corev1.Endpoints{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Endpoints), err
}
