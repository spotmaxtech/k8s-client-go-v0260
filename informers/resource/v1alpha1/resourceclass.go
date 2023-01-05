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

// ResourceClassInformer provides access to a shared informer and lister for
// ResourceClasses.
type ResourceClassInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ResourceClassLister
}

type resourceClassInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewResourceClassInformer constructs a new informer for ResourceClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewResourceClassInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredResourceClassInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredResourceClassInformer constructs a new informer for ResourceClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredResourceClassInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResourceV1alpha1().ResourceClasses().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResourceV1alpha1().ResourceClasses().Watch(context.TODO(), options)
			},
		},
		&resourcev1alpha1.ResourceClass{},
		resyncPeriod,
		indexers,
	)
}

func (f *resourceClassInformer) defaultInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredResourceClassInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *resourceClassInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&resourcev1alpha1.ResourceClass{}, f.defaultInformer)
}

func (f *resourceClassInformer) Lister() v1alpha1.ResourceClassLister {
	return v1alpha1.NewResourceClassLister(f.Informer().GetIndexer())
}
