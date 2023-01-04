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

package v1beta2

import (
	"context"
	time "time"

	flowcontrolv1beta2 "k8s.io/api/flowcontrol/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "github.com/spotmaxtech/k8s-client-go-v0260/informers/internalinterfaces"
	kubernetes "github.com/spotmaxtech/k8s-client-go-v0260/kubernetes"
	v1beta2 "github.com/spotmaxtech/k8s-client-go-v0260/listers/flowcontrol/v1beta2"
	cache "github.com/spotmaxtech/k8s-client-go-v0260/tools/cache"
)

// FlowSchemaInformer provides access to a shared informer and lister for
// FlowSchemas.
type FlowSchemaInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta2.FlowSchemaLister
}

type flowSchemaInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewFlowSchemaInformer constructs a new informer for FlowSchema type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFlowSchemaInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFlowSchemaInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredFlowSchemaInformer constructs a new informer for FlowSchema type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFlowSchemaInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FlowcontrolV1beta2().FlowSchemas().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FlowcontrolV1beta2().FlowSchemas().Watch(context.TODO(), options)
			},
		},
		&flowcontrolv1beta2.FlowSchema{},
		resyncPeriod,
		indexers,
	)
}

func (f *flowSchemaInformer) defaultInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFlowSchemaInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *flowSchemaInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&flowcontrolv1beta2.FlowSchema{}, f.defaultInformer)
}

func (f *flowSchemaInformer) Lister() v1beta2.FlowSchemaLister {
	return v1beta2.NewFlowSchemaLister(f.Informer().GetIndexer())
}
