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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	apiserverinternalv1alpha1 "k8s.io/api/apiserverinternal/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "github.com/spotmaxtech/k8s-client-go-v0260/informers/internalinterfaces"
	kubernetes "github.com/spotmaxtech/k8s-client-go-v0260/kubernetes"
	v1alpha1 "github.com/spotmaxtech/k8s-client-go-v0260/listers/apiserverinternal/v1alpha1"
	cache "github.com/spotmaxtech/k8s-client-go-v0260/tools/cache"
)

// StorageVersionInformer provides access to a shared informer and lister for
// StorageVersions.
type StorageVersionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.StorageVersionLister
}

type storageVersionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewStorageVersionInformer constructs a new informer for StorageVersion type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewStorageVersionInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredStorageVersionInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredStorageVersionInformer constructs a new informer for StorageVersion type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredStorageVersionInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.InternalV1alpha1().StorageVersions().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.InternalV1alpha1().StorageVersions().Watch(context.TODO(), options)
			},
		},
		&apiserverinternalv1alpha1.StorageVersion{},
		resyncPeriod,
		indexers,
	)
}

func (f *storageVersionInformer) defaultInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredStorageVersionInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *storageVersionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apiserverinternalv1alpha1.StorageVersion{}, f.defaultInformer)
}

func (f *storageVersionInformer) Lister() v1alpha1.StorageVersionLister {
	return v1alpha1.NewStorageVersionLister(f.Informer().GetIndexer())
}
