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
	v2beta2 "github.com/spotmaxtech/k8s-client-go-v0260/kubernetes/typed/autoscaling/v2beta2"
	rest "github.com/spotmaxtech/k8s-client-go-v0260/rest"
	testing "github.com/spotmaxtech/k8s-client-go-v0260/testing"
)

type FakeAutoscalingV2beta2 struct {
	*testing.Fake
}

func (c *FakeAutoscalingV2beta2) HorizontalPodAutoscalers(namespace string) v2beta2.HorizontalPodAutoscalerInterface {
	return &FakeHorizontalPodAutoscalers{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeAutoscalingV2beta2) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
