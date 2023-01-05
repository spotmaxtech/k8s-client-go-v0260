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

// Package install installs the experimental API group, making it available as
// an option to all of the API encoding/decoding machinery.
package install

import (
	"github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/runtime"
	utilruntime "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/util/runtime"
	"github.com/spotmaxtech/k8s-client-go-v0260/pkg/apis/clientauthentication"
	"github.com/spotmaxtech/k8s-client-go-v0260/pkg/apis/clientauthentication/v1"
	"github.com/spotmaxtech/k8s-client-go-v0260/pkg/apis/clientauthentication/v1beta1"
)

// Install registers the API group and adds types to a scheme
func Install(scheme *runtime.Scheme) {
	utilruntime.Must(clientauthentication.AddToScheme(scheme))
	utilruntime.Must(v1.AddToScheme(scheme))
	utilruntime.Must(v1beta1.AddToScheme(scheme))
}
