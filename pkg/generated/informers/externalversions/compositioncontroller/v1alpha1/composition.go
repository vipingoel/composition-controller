/*
Copyright 2021 The Kubernetes sample-controller Authors.

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
	compositioncontrollerv1alpha1 "composition-controller/pkg/apis/compositioncontroller/v1alpha1"
	versioned "composition-controller/pkg/generated/clientset/versioned"
	internalinterfaces "composition-controller/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "composition-controller/pkg/generated/listers/compositioncontroller/v1alpha1"
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CompositionInformer provides access to a shared informer and lister for
// Compositions.
type CompositionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CompositionLister
}

type compositionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCompositionInformer constructs a new informer for Composition type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCompositionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCompositionInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCompositionInformer constructs a new informer for Composition type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCompositionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CrdV1alpha1().Compositions(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CrdV1alpha1().Compositions(namespace).Watch(context.TODO(), options)
			},
		},
		&compositioncontrollerv1alpha1.Composition{},
		resyncPeriod,
		indexers,
	)
}

func (f *compositionInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCompositionInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *compositionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&compositioncontrollerv1alpha1.Composition{}, f.defaultInformer)
}

func (f *compositionInformer) Lister() v1alpha1.CompositionLister {
	return v1alpha1.NewCompositionLister(f.Informer().GetIndexer())
}
