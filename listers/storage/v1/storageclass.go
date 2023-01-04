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

package v1

import (
	v1 "k8s.io/api/storage/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"github.com/spotmaxtech/k8s-client-go-v0260/tools/cache"
)

// StorageClassLister helps list StorageClasses.
// All objects returned here must be treated as read-only.
type StorageClassLister interface {
	// List lists all StorageClasses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.StorageClass, err error)
	// Get retrieves the StorageClass from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.StorageClass, error)
	StorageClassListerExpansion
}

// storageClassLister implements the StorageClassLister interface.
type storageClassLister struct {
	indexer cache.Indexer
}

// NewStorageClassLister returns a new StorageClassLister.
func NewStorageClassLister(indexer cache.Indexer) StorageClassLister {
	return &storageClassLister{indexer: indexer}
}

// List lists all StorageClasses in the indexer.
func (s *storageClassLister) List(selector labels.Selector) (ret []*v1.StorageClass, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.StorageClass))
	})
	return ret, err
}

// Get retrieves the StorageClass from the index for a given name.
func (s *storageClassLister) Get(name string) (*v1.StorageClass, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("storageclass"), name)
	}
	return obj.(*v1.StorageClass), nil
}
