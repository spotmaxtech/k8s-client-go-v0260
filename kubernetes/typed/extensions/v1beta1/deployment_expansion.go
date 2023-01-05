/*
Copyright 2016 The Kubernetes Authors.

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

package v1beta1

import (
	"context"

	"github.com/spotmaxtech/k8s-api-v0260/extensions/v1beta1"
	metav1 "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/apis/meta/v1"
	scheme "github.com/spotmaxtech/k8s-client-go-v0260/kubernetes/scheme"
)

// The DeploymentExpansion interface allows manually adding extra methods to the DeploymentInterface.
type DeploymentExpansion interface {
	Rollback(context.Context, *v1beta1.DeploymentRollback, metav1.CreateOptions) error
}

// Rollback applied the provided DeploymentRollback to the named deployment in the current namespace.
func (c *deployments) Rollback(ctx context.Context, deploymentRollback *v1beta1.DeploymentRollback, opts metav1.CreateOptions) error {
	return c.client.Post().Namespace(c.ns).Resource("deployments").Name(deploymentRollback.Name).VersionedParams(&opts, scheme.ParameterCodec).SubResource("rollback").Body(deploymentRollback).Do(ctx).Error()
}
