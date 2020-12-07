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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// KustoEventhubDataConnectionLister helps list KustoEventhubDataConnections.
type KustoEventhubDataConnectionLister interface {
	// List lists all KustoEventhubDataConnections in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.KustoEventhubDataConnection, err error)
	// KustoEventhubDataConnections returns an object that can list and get KustoEventhubDataConnections.
	KustoEventhubDataConnections(namespace string) KustoEventhubDataConnectionNamespaceLister
	KustoEventhubDataConnectionListerExpansion
}

// kustoEventhubDataConnectionLister implements the KustoEventhubDataConnectionLister interface.
type kustoEventhubDataConnectionLister struct {
	indexer cache.Indexer
}

// NewKustoEventhubDataConnectionLister returns a new KustoEventhubDataConnectionLister.
func NewKustoEventhubDataConnectionLister(indexer cache.Indexer) KustoEventhubDataConnectionLister {
	return &kustoEventhubDataConnectionLister{indexer: indexer}
}

// List lists all KustoEventhubDataConnections in the indexer.
func (s *kustoEventhubDataConnectionLister) List(selector labels.Selector) (ret []*v1alpha1.KustoEventhubDataConnection, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KustoEventhubDataConnection))
	})
	return ret, err
}

// KustoEventhubDataConnections returns an object that can list and get KustoEventhubDataConnections.
func (s *kustoEventhubDataConnectionLister) KustoEventhubDataConnections(namespace string) KustoEventhubDataConnectionNamespaceLister {
	return kustoEventhubDataConnectionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// KustoEventhubDataConnectionNamespaceLister helps list and get KustoEventhubDataConnections.
type KustoEventhubDataConnectionNamespaceLister interface {
	// List lists all KustoEventhubDataConnections in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.KustoEventhubDataConnection, err error)
	// Get retrieves the KustoEventhubDataConnection from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.KustoEventhubDataConnection, error)
	KustoEventhubDataConnectionNamespaceListerExpansion
}

// kustoEventhubDataConnectionNamespaceLister implements the KustoEventhubDataConnectionNamespaceLister
// interface.
type kustoEventhubDataConnectionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all KustoEventhubDataConnections in the indexer for a given namespace.
func (s kustoEventhubDataConnectionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.KustoEventhubDataConnection, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KustoEventhubDataConnection))
	})
	return ret, err
}

// Get retrieves the KustoEventhubDataConnection from the indexer for a given namespace and name.
func (s kustoEventhubDataConnectionNamespaceLister) Get(name string) (*v1alpha1.KustoEventhubDataConnection, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("kustoeventhubdataconnection"), name)
	}
	return obj.(*v1alpha1.KustoEventhubDataConnection), nil
}
