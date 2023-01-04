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

	certificatesv1 "k8s.io/api/certificates/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "github.com/spotmaxtech/k8s-client-go-v0260/informers/internalinterfaces"
	kubernetes "github.com/spotmaxtech/k8s-client-go-v0260/kubernetes"
	v1 "github.com/spotmaxtech/k8s-client-go-v0260/listers/certificates/v1"
	cache "github.com/spotmaxtech/k8s-client-go-v0260/tools/cache"
)

// CertificateSigningRequestInformer provides access to a shared informer and lister for
// CertificateSigningRequests.
type CertificateSigningRequestInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.CertificateSigningRequestLister
}

type certificateSigningRequestInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewCertificateSigningRequestInformer constructs a new informer for CertificateSigningRequest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCertificateSigningRequestInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCertificateSigningRequestInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredCertificateSigningRequestInformer constructs a new informer for CertificateSigningRequest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCertificateSigningRequestInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CertificatesV1().CertificateSigningRequests().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CertificatesV1().CertificateSigningRequests().Watch(context.TODO(), options)
			},
		},
		&certificatesv1.CertificateSigningRequest{},
		resyncPeriod,
		indexers,
	)
}

func (f *certificateSigningRequestInformer) defaultInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCertificateSigningRequestInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *certificateSigningRequestInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&certificatesv1.CertificateSigningRequest{}, f.defaultInformer)
}

func (f *certificateSigningRequestInformer) Lister() v1.CertificateSigningRequestLister {
	return v1.NewCertificateSigningRequestLister(f.Informer().GetIndexer())
}
