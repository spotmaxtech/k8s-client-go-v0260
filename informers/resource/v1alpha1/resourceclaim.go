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

	resourcev1alpha1 "github.com/spotmaxtech/k8s-api-v0260/resource/v1alpha1"
	v1 "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/apis/meta/v1"
	runtime "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/runtime"
	watch "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/watch"
	internalinterfaces "github.com/spotmaxtech/k8s-client-go-v0260/informers/internalinterfaces"
	kubernetes "github.com/spotmaxtech/k8s-client-go-v0260/kubernetes"
	v1alpha1 "github.com/spotmaxtech/k8s-client-go-v0260/listers/resource/v1alpha1"
	cache "github.com/spotmaxtech/k8s-client-go-v0260/tools/cache"
)

// ResourceClaimInformer provides access to a shared informer and lister for
// ResourceClaims.
type ResourceClaimInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ResourceClaimLister
}

type resourceClaimInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewResourceClaimInformer constructs a new informer for ResourceClaim type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewResourceClaimInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredResourceClaimInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredResourceClaimInformer constructs a new informer for ResourceClaim type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredResourceClaimInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResourceV1alpha1().ResourceClaims(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResourceV1alpha1().ResourceClaims(namespace).Watch(context.TODO(), options)
			},
		},
		&resourcev1alpha1.ResourceClaim{},
		resyncPeriod,
		indexers,
	)
}

func (f *resourceClaimInformer) defaultInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredResourceClaimInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *resourceClaimInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&resourcev1alpha1.ResourceClaim{}, f.defaultInformer)
}

func (f *resourceClaimInformer) Lister() v1alpha1.ResourceClaimLister {
	return v1alpha1.NewResourceClaimLister(f.Informer().GetIndexer())
}
