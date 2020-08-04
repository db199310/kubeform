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

// AppServiceHybridConnectionLister helps list AppServiceHybridConnections.
type AppServiceHybridConnectionLister interface {
	// List lists all AppServiceHybridConnections in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AppServiceHybridConnection, err error)
	// AppServiceHybridConnections returns an object that can list and get AppServiceHybridConnections.
	AppServiceHybridConnections(namespace string) AppServiceHybridConnectionNamespaceLister
	AppServiceHybridConnectionListerExpansion
}

// appServiceHybridConnectionLister implements the AppServiceHybridConnectionLister interface.
type appServiceHybridConnectionLister struct {
	indexer cache.Indexer
}

// NewAppServiceHybridConnectionLister returns a new AppServiceHybridConnectionLister.
func NewAppServiceHybridConnectionLister(indexer cache.Indexer) AppServiceHybridConnectionLister {
	return &appServiceHybridConnectionLister{indexer: indexer}
}

// List lists all AppServiceHybridConnections in the indexer.
func (s *appServiceHybridConnectionLister) List(selector labels.Selector) (ret []*v1alpha1.AppServiceHybridConnection, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AppServiceHybridConnection))
	})
	return ret, err
}

// AppServiceHybridConnections returns an object that can list and get AppServiceHybridConnections.
func (s *appServiceHybridConnectionLister) AppServiceHybridConnections(namespace string) AppServiceHybridConnectionNamespaceLister {
	return appServiceHybridConnectionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AppServiceHybridConnectionNamespaceLister helps list and get AppServiceHybridConnections.
type AppServiceHybridConnectionNamespaceLister interface {
	// List lists all AppServiceHybridConnections in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AppServiceHybridConnection, err error)
	// Get retrieves the AppServiceHybridConnection from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AppServiceHybridConnection, error)
	AppServiceHybridConnectionNamespaceListerExpansion
}

// appServiceHybridConnectionNamespaceLister implements the AppServiceHybridConnectionNamespaceLister
// interface.
type appServiceHybridConnectionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AppServiceHybridConnections in the indexer for a given namespace.
func (s appServiceHybridConnectionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AppServiceHybridConnection, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AppServiceHybridConnection))
	})
	return ret, err
}

// Get retrieves the AppServiceHybridConnection from the indexer for a given namespace and name.
func (s appServiceHybridConnectionNamespaceLister) Get(name string) (*v1alpha1.AppServiceHybridConnection, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("appservicehybridconnection"), name)
	}
	return obj.(*v1alpha1.AppServiceHybridConnection), nil
}
