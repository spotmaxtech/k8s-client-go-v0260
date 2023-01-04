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

package v1beta1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1beta1 "k8s.io/api/extensions/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	extensionsv1beta1 "github.com/spotmaxtech/k8s-client-go-v0260/applyconfigurations/extensions/v1beta1"
	scheme "github.com/spotmaxtech/k8s-client-go-v0260/kubernetes/scheme"
	rest "github.com/spotmaxtech/k8s-client-go-v0260/rest"
)

// NetworkPoliciesGetter has a method to return a NetworkPolicyInterface.
// A group's client should implement this interface.
type NetworkPoliciesGetter interface {
	NetworkPolicies(namespace string) NetworkPolicyInterface
}

// NetworkPolicyInterface has methods to work with NetworkPolicy resources.
type NetworkPolicyInterface interface {
	Create(ctx context.Context, networkPolicy *v1beta1.NetworkPolicy, opts v1.CreateOptions) (*v1beta1.NetworkPolicy, error)
	Update(ctx context.Context, networkPolicy *v1beta1.NetworkPolicy, opts v1.UpdateOptions) (*v1beta1.NetworkPolicy, error)
	UpdateStatus(ctx context.Context, networkPolicy *v1beta1.NetworkPolicy, opts v1.UpdateOptions) (*v1beta1.NetworkPolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.NetworkPolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.NetworkPolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.NetworkPolicy, err error)
	Apply(ctx context.Context, networkPolicy *extensionsv1beta1.NetworkPolicyApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.NetworkPolicy, err error)
	ApplyStatus(ctx context.Context, networkPolicy *extensionsv1beta1.NetworkPolicyApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.NetworkPolicy, err error)
	NetworkPolicyExpansion
}

// networkPolicies implements NetworkPolicyInterface
type networkPolicies struct {
	client rest.Interface
	ns     string
}

// newNetworkPolicies returns a NetworkPolicies
func newNetworkPolicies(c *ExtensionsV1beta1Client, namespace string) *networkPolicies {
	return &networkPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the networkPolicy, and returns the corresponding networkPolicy object, and an error if there is any.
func (c *networkPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.NetworkPolicy, err error) {
	result = &v1beta1.NetworkPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NetworkPolicies that match those selectors.
func (c *networkPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.NetworkPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.NetworkPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested networkPolicies.
func (c *networkPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("networkpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a networkPolicy and creates it.  Returns the server's representation of the networkPolicy, and an error, if there is any.
func (c *networkPolicies) Create(ctx context.Context, networkPolicy *v1beta1.NetworkPolicy, opts v1.CreateOptions) (result *v1beta1.NetworkPolicy, err error) {
	result = &v1beta1.NetworkPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("networkpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(networkPolicy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a networkPolicy and updates it. Returns the server's representation of the networkPolicy, and an error, if there is any.
func (c *networkPolicies) Update(ctx context.Context, networkPolicy *v1beta1.NetworkPolicy, opts v1.UpdateOptions) (result *v1beta1.NetworkPolicy, err error) {
	result = &v1beta1.NetworkPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("networkpolicies").
		Name(networkPolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(networkPolicy).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *networkPolicies) UpdateStatus(ctx context.Context, networkPolicy *v1beta1.NetworkPolicy, opts v1.UpdateOptions) (result *v1beta1.NetworkPolicy, err error) {
	result = &v1beta1.NetworkPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("networkpolicies").
		Name(networkPolicy.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(networkPolicy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the networkPolicy and deletes it. Returns an error if one occurs.
func (c *networkPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networkpolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *networkPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networkpolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched networkPolicy.
func (c *networkPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.NetworkPolicy, err error) {
	result = &v1beta1.NetworkPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("networkpolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied networkPolicy.
func (c *networkPolicies) Apply(ctx context.Context, networkPolicy *extensionsv1beta1.NetworkPolicyApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.NetworkPolicy, err error) {
	if networkPolicy == nil {
		return nil, fmt.Errorf("networkPolicy provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(networkPolicy)
	if err != nil {
		return nil, err
	}
	name := networkPolicy.Name
	if name == nil {
		return nil, fmt.Errorf("networkPolicy.Name must be provided to Apply")
	}
	result = &v1beta1.NetworkPolicy{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("networkpolicies").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *networkPolicies) ApplyStatus(ctx context.Context, networkPolicy *extensionsv1beta1.NetworkPolicyApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.NetworkPolicy, err error) {
	if networkPolicy == nil {
		return nil, fmt.Errorf("networkPolicy provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(networkPolicy)
	if err != nil {
		return nil, err
	}

	name := networkPolicy.Name
	if name == nil {
		return nil, fmt.Errorf("networkPolicy.Name must be provided to Apply")
	}

	result = &v1beta1.NetworkPolicy{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("networkpolicies").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
