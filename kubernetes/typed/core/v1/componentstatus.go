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

package v1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1 "github.com/spotmaxtech/k8s-api-v0260/core/v1"
	metav1 "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/apis/meta/v1"
	types "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/types"
	watch "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/watch"
	corev1 "github.com/spotmaxtech/k8s-client-go-v0260/applyconfigurations/core/v1"
	scheme "github.com/spotmaxtech/k8s-client-go-v0260/kubernetes/scheme"
	rest "github.com/spotmaxtech/k8s-client-go-v0260/rest"
)

// ComponentStatusesGetter has a method to return a ComponentStatusInterface.
// A group's client should implement this interface.
type ComponentStatusesGetter interface {
	ComponentStatuses() ComponentStatusInterface
}

// ComponentStatusInterface has methods to work with ComponentStatus resources.
type ComponentStatusInterface interface {
	Create(ctx context.Context, componentStatus *v1.ComponentStatus, opts metav1.CreateOptions) (*v1.ComponentStatus, error)
	Update(ctx context.Context, componentStatus *v1.ComponentStatus, opts metav1.UpdateOptions) (*v1.ComponentStatus, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ComponentStatus, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ComponentStatusList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ComponentStatus, err error)
	Apply(ctx context.Context, componentStatus *corev1.ComponentStatusApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ComponentStatus, err error)
	ComponentStatusExpansion
}

// componentStatuses implements ComponentStatusInterface
type componentStatuses struct {
	client rest.Interface
}

// newComponentStatuses returns a ComponentStatuses
func newComponentStatuses(c *CoreV1Client) *componentStatuses {
	return &componentStatuses{
		client: c.RESTClient(),
	}
}

// Get takes name of the componentStatus, and returns the corresponding componentStatus object, and an error if there is any.
func (c *componentStatuses) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ComponentStatus, err error) {
	result = &v1.ComponentStatus{}
	err = c.client.Get().
		Resource("componentstatuses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComponentStatuses that match those selectors.
func (c *componentStatuses) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ComponentStatusList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ComponentStatusList{}
	err = c.client.Get().
		Resource("componentstatuses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested componentStatuses.
func (c *componentStatuses) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("componentstatuses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a componentStatus and creates it.  Returns the server's representation of the componentStatus, and an error, if there is any.
func (c *componentStatuses) Create(ctx context.Context, componentStatus *v1.ComponentStatus, opts metav1.CreateOptions) (result *v1.ComponentStatus, err error) {
	result = &v1.ComponentStatus{}
	err = c.client.Post().
		Resource("componentstatuses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(componentStatus).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a componentStatus and updates it. Returns the server's representation of the componentStatus, and an error, if there is any.
func (c *componentStatuses) Update(ctx context.Context, componentStatus *v1.ComponentStatus, opts metav1.UpdateOptions) (result *v1.ComponentStatus, err error) {
	result = &v1.ComponentStatus{}
	err = c.client.Put().
		Resource("componentstatuses").
		Name(componentStatus.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(componentStatus).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the componentStatus and deletes it. Returns an error if one occurs.
func (c *componentStatuses) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("componentstatuses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *componentStatuses) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("componentstatuses").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched componentStatus.
func (c *componentStatuses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ComponentStatus, err error) {
	result = &v1.ComponentStatus{}
	err = c.client.Patch(pt).
		Resource("componentstatuses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied componentStatus.
func (c *componentStatuses) Apply(ctx context.Context, componentStatus *corev1.ComponentStatusApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ComponentStatus, err error) {
	if componentStatus == nil {
		return nil, fmt.Errorf("componentStatus provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(componentStatus)
	if err != nil {
		return nil, err
	}
	name := componentStatus.Name
	if name == nil {
		return nil, fmt.Errorf("componentStatus.Name must be provided to Apply")
	}
	result = &v1.ComponentStatus{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("componentstatuses").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
