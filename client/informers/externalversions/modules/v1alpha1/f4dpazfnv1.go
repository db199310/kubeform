/*
Copyright The Kubeform Authors.

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
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	modulesv1alpha1 "kubeform.dev/kubeform/apis/modules/v1alpha1"
	versioned "kubeform.dev/kubeform/client/clientset/versioned"
	internalinterfaces "kubeform.dev/kubeform/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/kubeform/client/listers/modules/v1alpha1"
)

// F4dpAzFnv1Informer provides access to a shared informer and lister for
// F4dpAzFnv1s.
type F4dpAzFnv1Informer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.F4dpAzFnv1Lister
}

type f4dpAzFnv1Informer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewF4dpAzFnv1Informer constructs a new informer for F4dpAzFnv1 type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewF4dpAzFnv1Informer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredF4dpAzFnv1Informer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredF4dpAzFnv1Informer constructs a new informer for F4dpAzFnv1 type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredF4dpAzFnv1Informer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ModulesV1alpha1().F4dpAzFnv1s(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ModulesV1alpha1().F4dpAzFnv1s(namespace).Watch(options)
			},
		},
		&modulesv1alpha1.F4dpAzFnv1{},
		resyncPeriod,
		indexers,
	)
}

func (f *f4dpAzFnv1Informer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredF4dpAzFnv1Informer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *f4dpAzFnv1Informer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&modulesv1alpha1.F4dpAzFnv1{}, f.defaultInformer)
}

func (f *f4dpAzFnv1Informer) Lister() v1alpha1.F4dpAzFnv1Lister {
	return v1alpha1.NewF4dpAzFnv1Lister(f.Informer().GetIndexer())
}