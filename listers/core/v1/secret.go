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
	v1 "github.com/spotmaxtech/k8s-api-v0260/core/v1"
	"github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/api/errors"
	"github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/labels"
	"github.com/spotmaxtech/k8s-client-go-v0260/tools/cache"
)

// SecretLister helps list Secrets.
// All objects returned here must be treated as read-only.
type SecretLister interface {
	// List lists all Secrets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Secret, err error)
	// Secrets returns an object that can list and get Secrets.
	Secrets(namespace string) SecretNamespaceLister
	SecretListerExpansion
}

// secretLister implements the SecretLister interface.
type secretLister struct {
	indexer cache.Indexer
}

// NewSecretLister returns a new SecretLister.
func NewSecretLister(indexer cache.Indexer) SecretLister {
	return &secretLister{indexer: indexer}
}

// List lists all Secrets in the indexer.
func (s *secretLister) List(selector labels.Selector) (ret []*v1.Secret, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Secret))
	})
	return ret, err
}

// Secrets returns an object that can list and get Secrets.
func (s *secretLister) Secrets(namespace string) SecretNamespaceLister {
	return secretNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SecretNamespaceLister helps list and get Secrets.
// All objects returned here must be treated as read-only.
type SecretNamespaceLister interface {
	// List lists all Secrets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Secret, err error)
	// Get retrieves the Secret from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Secret, error)
	SecretNamespaceListerExpansion
}

// secretNamespaceLister implements the SecretNamespaceLister
// interface.
type secretNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Secrets in the indexer for a given namespace.
func (s secretNamespaceLister) List(selector labels.Selector) (ret []*v1.Secret, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Secret))
	})
	return ret, err
}

// Get retrieves the Secret from the indexer for a given namespace and name.
func (s secretNamespaceLister) Get(name string) (*v1.Secret, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("secret"), name)
	}
	return obj.(*v1.Secret), nil
}
