/*
Copyright 2022 The Kubernetes Authors.

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

package discovery

import (
	"testing"

	apidiscovery "github.com/spotmaxtech/k8s-api-v0260/apidiscovery/v2beta1"
	metav1 "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/apis/meta/v1"
	"github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/runtime/schema"
	"github.com/stretchr/testify/assert"
)

func TestSplitGroupsAndResources(t *testing.T) {
	tests := []struct {
		name                string
		agg                 apidiscovery.APIGroupDiscoveryList
		expectedGroups      metav1.APIGroupList
		expectedGVResources map[schema.GroupVersion]*metav1.APIResourceList
	}{
		{
			name: "Aggregated discovery: core/v1 group and pod resource",
			agg: apidiscovery.APIGroupDiscoveryList{
				Items: []apidiscovery.APIGroupDiscovery{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name: "",
						},
						Versions: []apidiscovery.APIVersionDiscovery{
							{
								Version: "v1",
								Resources: []apidiscovery.APIResourceDiscovery{
									{
										Resource: "pods",
										ResponseKind: &metav1.GroupVersionKind{
											Group:   "",
											Version: "v1",
											Kind:    "Pod",
										},
										Scope: apidiscovery.ScopeNamespace,
									},
								},
							},
						},
					},
				},
			},
			expectedGroups: metav1.APIGroupList{
				Groups: []metav1.APIGroup{
					{
						Name: "",
						Versions: []metav1.GroupVersionForDiscovery{
							{
								GroupVersion: "v1",
								Version:      "v1",
							},
						},
						PreferredVersion: metav1.GroupVersionForDiscovery{
							GroupVersion: "v1",
							Version:      "v1",
						},
					},
				},
			},
			expectedGVResources: map[schema.GroupVersion]*metav1.APIResourceList{
				{Group: "", Version: "v1"}: {
					GroupVersion: "v1",
					APIResources: []metav1.APIResource{
						{
							Name:       "pods",
							Namespaced: true,
							Group:      "",
							Version:    "v1",
							Kind:       "Pod",
						},
					},
				},
			},
		},
		{
			name: "Aggregated discovery: 1 group/1 resources at /api, 1 group/2 versions/1 resources at /apis",
			agg: apidiscovery.APIGroupDiscoveryList{
				Items: []apidiscovery.APIGroupDiscovery{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name: "apps",
						},
						Versions: []apidiscovery.APIVersionDiscovery{
							{
								Version: "v2",
								Resources: []apidiscovery.APIResourceDiscovery{
									{
										Resource: "deployments",
										ResponseKind: &metav1.GroupVersionKind{
											Group:   "apps",
											Version: "v2",
											Kind:    "Deployment",
										},
										Scope: apidiscovery.ScopeNamespace,
									},
								},
							},
							{
								Version: "v1",
								Resources: []apidiscovery.APIResourceDiscovery{
									{
										Resource: "deployments",
										ResponseKind: &metav1.GroupVersionKind{
											Group:   "apps",
											Version: "v1",
											Kind:    "Deployment",
										},
										Scope: apidiscovery.ScopeNamespace,
									},
								},
							},
						},
					},
				},
			},
			expectedGroups: metav1.APIGroupList{
				Groups: []metav1.APIGroup{
					{
						Name: "apps",
						Versions: []metav1.GroupVersionForDiscovery{
							{
								GroupVersion: "apps/v2",
								Version:      "v2",
							},
							{
								GroupVersion: "apps/v1",
								Version:      "v1",
							},
						},
						PreferredVersion: metav1.GroupVersionForDiscovery{
							GroupVersion: "apps/v2",
							Version:      "v2",
						},
					},
				},
			},
			expectedGVResources: map[schema.GroupVersion]*metav1.APIResourceList{
				{Group: "apps", Version: "v1"}: {
					GroupVersion: "apps/v1",
					APIResources: []metav1.APIResource{
						{
							Name:       "deployments",
							Namespaced: true,
							Group:      "apps",
							Version:    "v1",
							Kind:       "Deployment",
						},
					},
				},
				{Group: "apps", Version: "v2"}: {
					GroupVersion: "apps/v2",
					APIResources: []metav1.APIResource{
						{
							Name:       "deployments",
							Namespaced: true,
							Group:      "apps",
							Version:    "v2",
							Kind:       "Deployment",
						},
					},
				},
			},
		},
		{
			name: "Aggregated discovery: 1 group/2 resources at /api, 1 group/2 resources at /apis",
			agg: apidiscovery.APIGroupDiscoveryList{
				Items: []apidiscovery.APIGroupDiscovery{
					{
						Versions: []apidiscovery.APIVersionDiscovery{
							{
								Version: "v1",
								Resources: []apidiscovery.APIResourceDiscovery{
									{
										Resource: "pods",
										ResponseKind: &metav1.GroupVersionKind{
											Group:   "",
											Version: "v1",
											Kind:    "Pod",
										},
										Scope: apidiscovery.ScopeNamespace,
									},
									{
										Resource: "services",
										ResponseKind: &metav1.GroupVersionKind{
											Group:   "",
											Version: "v1",
											Kind:    "Service",
										},
										Scope: apidiscovery.ScopeNamespace,
									},
								},
							},
						},
					},
					{
						ObjectMeta: metav1.ObjectMeta{
							Name: "apps",
						},
						Versions: []apidiscovery.APIVersionDiscovery{
							{
								Version: "v1",
								Resources: []apidiscovery.APIResourceDiscovery{
									{
										Resource: "deployments",
										ResponseKind: &metav1.GroupVersionKind{
											Group:   "apps",
											Version: "v1",
											Kind:    "Deployment",
										},
										Scope: apidiscovery.ScopeNamespace,
									},
									{
										Resource: "statefulsets",
										ResponseKind: &metav1.GroupVersionKind{
											Group:   "apps",
											Version: "v1",
											Kind:    "StatefulSet",
										},
										Scope: apidiscovery.ScopeNamespace,
									},
								},
							},
						},
					},
				},
			},
			expectedGroups: metav1.APIGroupList{
				Groups: []metav1.APIGroup{
					{
						Name: "",
						Versions: []metav1.GroupVersionForDiscovery{
							{
								GroupVersion: "v1",
								Version:      "v1",
							},
						},
						PreferredVersion: metav1.GroupVersionForDiscovery{
							GroupVersion: "v1",
							Version:      "v1",
						},
					},
					{
						Name: "apps",
						Versions: []metav1.GroupVersionForDiscovery{
							{
								GroupVersion: "apps/v1",
								Version:      "v1",
							},
						},
						PreferredVersion: metav1.GroupVersionForDiscovery{
							GroupVersion: "apps/v1",
							Version:      "v1",
						},
					},
				},
			},
			expectedGVResources: map[schema.GroupVersion]*metav1.APIResourceList{
				{Group: "", Version: "v1"}: {
					GroupVersion: "v1",
					APIResources: []metav1.APIResource{
						{
							Name:       "pods",
							Namespaced: true,
							Group:      "",
							Version:    "v1",
							Kind:       "Pod",
						},
						{
							Name:       "services",
							Namespaced: true,
							Group:      "",
							Version:    "v1",
							Kind:       "Service",
						},
					},
				},
				{Group: "apps", Version: "v1"}: {
					GroupVersion: "apps/v1",
					APIResources: []metav1.APIResource{
						{
							Name:       "deployments",
							Namespaced: true,
							Group:      "apps",
							Version:    "v1",
							Kind:       "Deployment",
						},
						{
							Name:       "statefulsets",
							Namespaced: true,
							Group:      "apps",
							Version:    "v1",
							Kind:       "StatefulSet",
						},
					},
				},
			},
		},
		{
			name: "Aggregated discovery: multiple groups with cluster-scoped resources",
			agg: apidiscovery.APIGroupDiscoveryList{
				Items: []apidiscovery.APIGroupDiscovery{
					{
						Versions: []apidiscovery.APIVersionDiscovery{
							{
								Version: "v1",
								Resources: []apidiscovery.APIResourceDiscovery{
									{
										Resource: "pods",
										ResponseKind: &metav1.GroupVersionKind{
											Group:   "",
											Version: "v1",
											Kind:    "Pod",
										},
										Scope: apidiscovery.ScopeNamespace,
									},
									{
										Resource: "namespaces",
										ResponseKind: &metav1.GroupVersionKind{
											Group:   "",
											Version: "v1",
											Kind:    "Namespace",
										},
										Scope: apidiscovery.ScopeCluster,
									},
								},
							},
						},
					},
					{
						ObjectMeta: metav1.ObjectMeta{
							Name: "rbac.authorization.k8s.io",
						},
						Versions: []apidiscovery.APIVersionDiscovery{
							{
								Version: "v1",
								Resources: []apidiscovery.APIResourceDiscovery{
									{
										Resource: "roles",
										ResponseKind: &metav1.GroupVersionKind{
											Group:   "rbac.authorization.k8s.io",
											Version: "v1",
											Kind:    "Role",
										},
										Scope: apidiscovery.ScopeCluster,
									},
									{
										Resource: "clusterroles",
										ResponseKind: &metav1.GroupVersionKind{
											Group:   "rbac.authorization.k8s.io",
											Version: "v1",
											Kind:    "ClusterRole",
										},
										Scope: apidiscovery.ScopeCluster,
									},
								},
							},
						},
					},
				},
			},
			expectedGroups: metav1.APIGroupList{
				Groups: []metav1.APIGroup{
					{
						Name: "",
						Versions: []metav1.GroupVersionForDiscovery{
							{
								GroupVersion: "v1",
								Version:      "v1",
							},
						},
						PreferredVersion: metav1.GroupVersionForDiscovery{
							GroupVersion: "v1",
							Version:      "v1",
						},
					},
					{
						Name: "rbac.authorization.k8s.io",
						Versions: []metav1.GroupVersionForDiscovery{
							{
								GroupVersion: "rbac.authorization.k8s.io/v1",
								Version:      "v1",
							},
						},
						PreferredVersion: metav1.GroupVersionForDiscovery{
							GroupVersion: "rbac.authorization.k8s.io/v1",
							Version:      "v1",
						},
					},
				},
			},
			expectedGVResources: map[schema.GroupVersion]*metav1.APIResourceList{
				{Group: "", Version: "v1"}: {
					GroupVersion: "v1",
					APIResources: []metav1.APIResource{
						{
							Name:       "pods",
							Namespaced: true,
							Group:      "",
							Version:    "v1",
							Kind:       "Pod",
						},
						{
							Name:       "namespaces",
							Namespaced: false,
							Group:      "",
							Version:    "v1",
							Kind:       "Namespace",
						},
					},
				},
				{Group: "rbac.authorization.k8s.io", Version: "v1"}: {
					GroupVersion: "rbac.authorization.k8s.io/v1",
					APIResources: []metav1.APIResource{
						{
							Name:       "roles",
							Namespaced: false,
							Group:      "rbac.authorization.k8s.io",
							Version:    "v1",
							Kind:       "Role",
						},
						{
							Name:       "clusterroles",
							Namespaced: false,
							Group:      "rbac.authorization.k8s.io",
							Version:    "v1",
							Kind:       "ClusterRole",
						},
					},
				},
			},
		},
		{
			name: "Aggregated discovery with single subresource",
			agg: apidiscovery.APIGroupDiscoveryList{
				Items: []apidiscovery.APIGroupDiscovery{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name: "apps",
						},
						Versions: []apidiscovery.APIVersionDiscovery{
							{
								Version: "v1",
								Resources: []apidiscovery.APIResourceDiscovery{
									{
										Resource: "deployments",
										ResponseKind: &metav1.GroupVersionKind{
											Group:   "apps",
											Version: "v1",
											Kind:    "Deployment",
										},
										Scope:            apidiscovery.ScopeNamespace,
										SingularResource: "deployment",
										ShortNames:       []string{"deploy"},
										Verbs:            []string{"parentverb1", "parentverb2", "parentverb3", "parentverb4"},
										Categories:       []string{"all", "testcategory"},
										Subresources: []apidiscovery.APISubresourceDiscovery{
											{
												Subresource: "scale",
												ResponseKind: &metav1.GroupVersionKind{
													Group:   "apps",
													Version: "v1",
													Kind:    "Deployment",
												},
												Verbs: []string{"get", "patch", "update"},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedGroups: metav1.APIGroupList{
				Groups: []metav1.APIGroup{
					{
						Name: "apps",
						Versions: []metav1.GroupVersionForDiscovery{
							{
								GroupVersion: "apps/v1",
								Version:      "v1",
							},
						},
						PreferredVersion: metav1.GroupVersionForDiscovery{
							GroupVersion: "apps/v1",
							Version:      "v1",
						},
					},
				},
			},
			expectedGVResources: map[schema.GroupVersion]*metav1.APIResourceList{
				{Group: "apps", Version: "v1"}: {
					GroupVersion: "apps/v1",
					APIResources: []metav1.APIResource{
						{
							Name:         "deployments",
							SingularName: "deployment",
							Namespaced:   true,
							Group:        "apps",
							Version:      "v1",
							Kind:         "Deployment",
							Verbs:        []string{"parentverb1", "parentverb2", "parentverb3", "parentverb4"},
							ShortNames:   []string{"deploy"},
							Categories:   []string{"all", "testcategory"},
						},
						{
							Name:         "deployments/scale",
							SingularName: "deployment",
							Namespaced:   true,
							Group:        "apps",
							Version:      "v1",
							Kind:         "Deployment",
							Verbs:        []string{"get", "patch", "update"},
						},
					},
				},
			},
		},
		{
			name: "Aggregated discovery with multiple subresources",
			agg: apidiscovery.APIGroupDiscoveryList{
				Items: []apidiscovery.APIGroupDiscovery{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name: "apps",
						},
						Versions: []apidiscovery.APIVersionDiscovery{
							{
								Version: "v1",
								Resources: []apidiscovery.APIResourceDiscovery{
									{
										Resource: "deployments",
										ResponseKind: &metav1.GroupVersionKind{
											Group:   "apps",
											Version: "v1",
											Kind:    "Deployment",
										},
										Scope:            apidiscovery.ScopeNamespace,
										SingularResource: "deployment",
										Subresources: []apidiscovery.APISubresourceDiscovery{
											{
												Subresource: "scale",
												ResponseKind: &metav1.GroupVersionKind{
													Group:   "apps",
													Version: "v1",
													Kind:    "Deployment",
												},
												Verbs: []string{"get", "patch", "update"},
											},
											{
												Subresource: "status",
												ResponseKind: &metav1.GroupVersionKind{
													Group:   "apps",
													Version: "v1",
													Kind:    "Deployment",
												},
												Verbs: []string{"get", "patch", "update"},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedGroups: metav1.APIGroupList{
				Groups: []metav1.APIGroup{
					{
						Name: "apps",
						Versions: []metav1.GroupVersionForDiscovery{
							{
								GroupVersion: "apps/v1",
								Version:      "v1",
							},
						},
						PreferredVersion: metav1.GroupVersionForDiscovery{
							GroupVersion: "apps/v1",
							Version:      "v1",
						},
					},
				},
			},
			expectedGVResources: map[schema.GroupVersion]*metav1.APIResourceList{
				{Group: "apps", Version: "v1"}: {
					GroupVersion: "apps/v1",
					APIResources: []metav1.APIResource{
						{
							Name:         "deployments",
							SingularName: "deployment",
							Namespaced:   true,
							Group:        "apps",
							Version:      "v1",
							Kind:         "Deployment",
						},
						{
							Name:         "deployments/scale",
							SingularName: "deployment",
							Namespaced:   true,
							Group:        "apps",
							Version:      "v1",
							Kind:         "Deployment",
							Verbs:        []string{"get", "patch", "update"},
						},
						{
							Name:         "deployments/status",
							SingularName: "deployment",
							Namespaced:   true,
							Group:        "apps",
							Version:      "v1",
							Kind:         "Deployment",
							Verbs:        []string{"get", "patch", "update"},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		apiGroups, resourcesByGV := SplitGroupsAndResources(test.agg)
		assert.Equal(t, test.expectedGroups, *apiGroups)
		assert.Equal(t, test.expectedGVResources, resourcesByGV)
	}
}
