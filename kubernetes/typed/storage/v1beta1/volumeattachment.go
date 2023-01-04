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

	v1beta1 "k8s.io/api/storage/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	storagev1beta1 "github.com/spotmaxtech/k8s-client-go-v0260/applyconfigurations/storage/v1beta1"
	scheme "github.com/spotmaxtech/k8s-client-go-v0260/kubernetes/scheme"
	rest "github.com/spotmaxtech/k8s-client-go-v0260/rest"
)

// VolumeAttachmentsGetter has a method to return a VolumeAttachmentInterface.
// A group's client should implement this interface.
type VolumeAttachmentsGetter interface {
	VolumeAttachments() VolumeAttachmentInterface
}

// VolumeAttachmentInterface has methods to work with VolumeAttachment resources.
type VolumeAttachmentInterface interface {
	Create(ctx context.Context, volumeAttachment *v1beta1.VolumeAttachment, opts v1.CreateOptions) (*v1beta1.VolumeAttachment, error)
	Update(ctx context.Context, volumeAttachment *v1beta1.VolumeAttachment, opts v1.UpdateOptions) (*v1beta1.VolumeAttachment, error)
	UpdateStatus(ctx context.Context, volumeAttachment *v1beta1.VolumeAttachment, opts v1.UpdateOptions) (*v1beta1.VolumeAttachment, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.VolumeAttachment, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.VolumeAttachmentList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VolumeAttachment, err error)
	Apply(ctx context.Context, volumeAttachment *storagev1beta1.VolumeAttachmentApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.VolumeAttachment, err error)
	ApplyStatus(ctx context.Context, volumeAttachment *storagev1beta1.VolumeAttachmentApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.VolumeAttachment, err error)
	VolumeAttachmentExpansion
}

// volumeAttachments implements VolumeAttachmentInterface
type volumeAttachments struct {
	client rest.Interface
}

// newVolumeAttachments returns a VolumeAttachments
func newVolumeAttachments(c *StorageV1beta1Client) *volumeAttachments {
	return &volumeAttachments{
		client: c.RESTClient(),
	}
}

// Get takes name of the volumeAttachment, and returns the corresponding volumeAttachment object, and an error if there is any.
func (c *volumeAttachments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VolumeAttachment, err error) {
	result = &v1beta1.VolumeAttachment{}
	err = c.client.Get().
		Resource("volumeattachments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VolumeAttachments that match those selectors.
func (c *volumeAttachments) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VolumeAttachmentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.VolumeAttachmentList{}
	err = c.client.Get().
		Resource("volumeattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested volumeAttachments.
func (c *volumeAttachments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("volumeattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a volumeAttachment and creates it.  Returns the server's representation of the volumeAttachment, and an error, if there is any.
func (c *volumeAttachments) Create(ctx context.Context, volumeAttachment *v1beta1.VolumeAttachment, opts v1.CreateOptions) (result *v1beta1.VolumeAttachment, err error) {
	result = &v1beta1.VolumeAttachment{}
	err = c.client.Post().
		Resource("volumeattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumeAttachment).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a volumeAttachment and updates it. Returns the server's representation of the volumeAttachment, and an error, if there is any.
func (c *volumeAttachments) Update(ctx context.Context, volumeAttachment *v1beta1.VolumeAttachment, opts v1.UpdateOptions) (result *v1beta1.VolumeAttachment, err error) {
	result = &v1beta1.VolumeAttachment{}
	err = c.client.Put().
		Resource("volumeattachments").
		Name(volumeAttachment.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumeAttachment).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *volumeAttachments) UpdateStatus(ctx context.Context, volumeAttachment *v1beta1.VolumeAttachment, opts v1.UpdateOptions) (result *v1beta1.VolumeAttachment, err error) {
	result = &v1beta1.VolumeAttachment{}
	err = c.client.Put().
		Resource("volumeattachments").
		Name(volumeAttachment.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumeAttachment).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the volumeAttachment and deletes it. Returns an error if one occurs.
func (c *volumeAttachments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("volumeattachments").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *volumeAttachments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("volumeattachments").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched volumeAttachment.
func (c *volumeAttachments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VolumeAttachment, err error) {
	result = &v1beta1.VolumeAttachment{}
	err = c.client.Patch(pt).
		Resource("volumeattachments").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied volumeAttachment.
func (c *volumeAttachments) Apply(ctx context.Context, volumeAttachment *storagev1beta1.VolumeAttachmentApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.VolumeAttachment, err error) {
	if volumeAttachment == nil {
		return nil, fmt.Errorf("volumeAttachment provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(volumeAttachment)
	if err != nil {
		return nil, err
	}
	name := volumeAttachment.Name
	if name == nil {
		return nil, fmt.Errorf("volumeAttachment.Name must be provided to Apply")
	}
	result = &v1beta1.VolumeAttachment{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("volumeattachments").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *volumeAttachments) ApplyStatus(ctx context.Context, volumeAttachment *storagev1beta1.VolumeAttachmentApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.VolumeAttachment, err error) {
	if volumeAttachment == nil {
		return nil, fmt.Errorf("volumeAttachment provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(volumeAttachment)
	if err != nil {
		return nil, err
	}

	name := volumeAttachment.Name
	if name == nil {
		return nil, fmt.Errorf("volumeAttachment.Name must be provided to Apply")
	}

	result = &v1beta1.VolumeAttachment{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("volumeattachments").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
