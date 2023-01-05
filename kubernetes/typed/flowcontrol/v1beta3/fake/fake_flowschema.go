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

	v1beta3 "github.com/spotmaxtech/k8s-api-v0260/flowcontrol/v1beta3"
	v1 "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/apis/meta/v1"
	labels "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/labels"
	schema "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/runtime/schema"
	types "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/types"
	watch "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/watch"
	flowcontrolv1beta3 "github.com/spotmaxtech/k8s-client-go-v0260/applyconfigurations/flowcontrol/v1beta3"
	testing "github.com/spotmaxtech/k8s-client-go-v0260/testing"
)

// FakeFlowSchemas implements FlowSchemaInterface
type FakeFlowSchemas struct {
	Fake *FakeFlowcontrolV1beta3
}

var flowschemasResource = schema.GroupVersionResource{Group: "flowcontrol.apiserver.k8s.io", Version: "v1beta3", Resource: "flowschemas"}

var flowschemasKind = schema.GroupVersionKind{Group: "flowcontrol.apiserver.k8s.io", Version: "v1beta3", Kind: "FlowSchema"}

// Get takes name of the flowSchema, and returns the corresponding flowSchema object, and an error if there is any.
func (c *FakeFlowSchemas) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta3.FlowSchema, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(flowschemasResource, name), &v1beta3.FlowSchema{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta3.FlowSchema), err
}

// List takes label and field selectors, and returns the list of FlowSchemas that match those selectors.
func (c *FakeFlowSchemas) List(ctx context.Context, opts v1.ListOptions) (result *v1beta3.FlowSchemaList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(flowschemasResource, flowschemasKind, opts), &v1beta3.FlowSchemaList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta3.FlowSchemaList{ListMeta: obj.(*v1beta3.FlowSchemaList).ListMeta}
	for _, item := range obj.(*v1beta3.FlowSchemaList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested flowSchemas.
func (c *FakeFlowSchemas) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(flowschemasResource, opts))
}

// Create takes the representation of a flowSchema and creates it.  Returns the server's representation of the flowSchema, and an error, if there is any.
func (c *FakeFlowSchemas) Create(ctx context.Context, flowSchema *v1beta3.FlowSchema, opts v1.CreateOptions) (result *v1beta3.FlowSchema, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(flowschemasResource, flowSchema), &v1beta3.FlowSchema{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta3.FlowSchema), err
}

// Update takes the representation of a flowSchema and updates it. Returns the server's representation of the flowSchema, and an error, if there is any.
func (c *FakeFlowSchemas) Update(ctx context.Context, flowSchema *v1beta3.FlowSchema, opts v1.UpdateOptions) (result *v1beta3.FlowSchema, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(flowschemasResource, flowSchema), &v1beta3.FlowSchema{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta3.FlowSchema), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFlowSchemas) UpdateStatus(ctx context.Context, flowSchema *v1beta3.FlowSchema, opts v1.UpdateOptions) (*v1beta3.FlowSchema, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(flowschemasResource, "status", flowSchema), &v1beta3.FlowSchema{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta3.FlowSchema), err
}

// Delete takes name of the flowSchema and deletes it. Returns an error if one occurs.
func (c *FakeFlowSchemas) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(flowschemasResource, name, opts), &v1beta3.FlowSchema{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFlowSchemas) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(flowschemasResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta3.FlowSchemaList{})
	return err
}

// Patch applies the patch and returns the patched flowSchema.
func (c *FakeFlowSchemas) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta3.FlowSchema, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(flowschemasResource, name, pt, data, subresources...), &v1beta3.FlowSchema{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta3.FlowSchema), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied flowSchema.
func (c *FakeFlowSchemas) Apply(ctx context.Context, flowSchema *flowcontrolv1beta3.FlowSchemaApplyConfiguration, opts v1.ApplyOptions) (result *v1beta3.FlowSchema, err error) {
	if flowSchema == nil {
		return nil, fmt.Errorf("flowSchema provided to Apply must not be nil")
	}
	data, err := json.Marshal(flowSchema)
	if err != nil {
		return nil, err
	}
	name := flowSchema.Name
	if name == nil {
		return nil, fmt.Errorf("flowSchema.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(flowschemasResource, *name, types.ApplyPatchType, data), &v1beta3.FlowSchema{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta3.FlowSchema), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeFlowSchemas) ApplyStatus(ctx context.Context, flowSchema *flowcontrolv1beta3.FlowSchemaApplyConfiguration, opts v1.ApplyOptions) (result *v1beta3.FlowSchema, err error) {
	if flowSchema == nil {
		return nil, fmt.Errorf("flowSchema provided to Apply must not be nil")
	}
	data, err := json.Marshal(flowSchema)
	if err != nil {
		return nil, err
	}
	name := flowSchema.Name
	if name == nil {
		return nil, fmt.Errorf("flowSchema.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(flowschemasResource, *name, types.ApplyPatchType, data, "status"), &v1beta3.FlowSchema{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta3.FlowSchema), err
}
