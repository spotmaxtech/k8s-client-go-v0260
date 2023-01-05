/*
Copyright 2017 The Kubernetes Authors.

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

package scale

import (
	"context"

	autoscalingapi "github.com/spotmaxtech/k8s-api-v0260/autoscaling/v1"
	metav1 "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/apis/meta/v1"
	"github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/runtime/schema"
	"github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/types"
)

// ScalesGetter can produce a ScaleInterface
type ScalesGetter interface {
	// Scales produces a ScaleInterface for a particular namespace.
	// Set namespace to the empty string for non-namespaced resources.
	Scales(namespace string) ScaleInterface
}

// ScaleInterface can fetch and update scales for
// resources in a particular namespace which implement
// the scale subresource.
type ScaleInterface interface {
	// Get fetches the scale of the given scalable resource.
	Get(ctx context.Context, resource schema.GroupResource, name string, opts metav1.GetOptions) (*autoscalingapi.Scale, error)

	// Update updates the scale of the given scalable resource.
	Update(ctx context.Context, resource schema.GroupResource, scale *autoscalingapi.Scale, opts metav1.UpdateOptions) (*autoscalingapi.Scale, error)

	// Patch patches the scale of the given scalable resource.
	Patch(ctx context.Context, gvr schema.GroupVersionResource, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions) (*autoscalingapi.Scale, error)
}
