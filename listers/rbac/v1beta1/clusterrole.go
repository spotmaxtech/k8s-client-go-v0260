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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/spotmaxtech/k8s-api-v0260/rbac/v1beta1"
	"github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/api/errors"
	"github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/labels"
	"github.com/spotmaxtech/k8s-client-go-v0260/tools/cache"
)

// ClusterRoleLister helps list ClusterRoles.
// All objects returned here must be treated as read-only.
type ClusterRoleLister interface {
	// List lists all ClusterRoles in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.ClusterRole, err error)
	// Get retrieves the ClusterRole from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.ClusterRole, error)
	ClusterRoleListerExpansion
}

// clusterRoleLister implements the ClusterRoleLister interface.
type clusterRoleLister struct {
	indexer cache.Indexer
}

// NewClusterRoleLister returns a new ClusterRoleLister.
func NewClusterRoleLister(indexer cache.Indexer) ClusterRoleLister {
	return &clusterRoleLister{indexer: indexer}
}

// List lists all ClusterRoles in the indexer.
func (s *clusterRoleLister) List(selector labels.Selector) (ret []*v1beta1.ClusterRole, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ClusterRole))
	})
	return ret, err
}

// Get retrieves the ClusterRole from the index for a given name.
func (s *clusterRoleLister) Get(name string) (*v1beta1.ClusterRole, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("clusterrole"), name)
	}
	return obj.(*v1beta1.ClusterRole), nil
}
