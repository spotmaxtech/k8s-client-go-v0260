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

package v1

import (
	"context"
	time "time"

	corev1 "github.com/spotmaxtech/k8s-api-v0260/core/v1"
	metav1 "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/apis/meta/v1"
	runtime "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/runtime"
	watch "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/watch"
	internalinterfaces "github.com/spotmaxtech/k8s-client-go-v0260/informers/internalinterfaces"
	kubernetes "github.com/spotmaxtech/k8s-client-go-v0260/kubernetes"
	v1 "github.com/spotmaxtech/k8s-client-go-v0260/listers/core/v1"
	cache "github.com/spotmaxtech/k8s-client-go-v0260/tools/cache"
)

// ServiceAccountInformer provides access to a shared informer and lister for
// ServiceAccounts.
type ServiceAccountInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ServiceAccountLister
}

type serviceAccountInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewServiceAccountInformer constructs a new informer for ServiceAccount type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewServiceAccountInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredServiceAccountInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredServiceAccountInformer constructs a new informer for ServiceAccount type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredServiceAccountInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().ServiceAccounts(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().ServiceAccounts(namespace).Watch(context.TODO(), options)
			},
		},
		&corev1.ServiceAccount{},
		resyncPeriod,
		indexers,
	)
}

func (f *serviceAccountInformer) defaultInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredServiceAccountInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *serviceAccountInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&corev1.ServiceAccount{}, f.defaultInformer)
}

func (f *serviceAccountInformer) Lister() v1.ServiceAccountLister {
	return v1.NewServiceAccountLister(f.Informer().GetIndexer())
}
