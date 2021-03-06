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

// AppServiceLister helps list AppServices.
type AppServiceLister interface {
	// List lists all AppServices in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AppService, err error)
	// AppServices returns an object that can list and get AppServices.
	AppServices(namespace string) AppServiceNamespaceLister
	AppServiceListerExpansion
}

// appServiceLister implements the AppServiceLister interface.
type appServiceLister struct {
	indexer cache.Indexer
}

// NewAppServiceLister returns a new AppServiceLister.
func NewAppServiceLister(indexer cache.Indexer) AppServiceLister {
	return &appServiceLister{indexer: indexer}
}

// List lists all AppServices in the indexer.
func (s *appServiceLister) List(selector labels.Selector) (ret []*v1alpha1.AppService, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AppService))
	})
	return ret, err
}

// AppServices returns an object that can list and get AppServices.
func (s *appServiceLister) AppServices(namespace string) AppServiceNamespaceLister {
	return appServiceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AppServiceNamespaceLister helps list and get AppServices.
type AppServiceNamespaceLister interface {
	// List lists all AppServices in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AppService, err error)
	// Get retrieves the AppService from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AppService, error)
	AppServiceNamespaceListerExpansion
}

// appServiceNamespaceLister implements the AppServiceNamespaceLister
// interface.
type appServiceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AppServices in the indexer for a given namespace.
func (s appServiceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AppService, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AppService))
	})
	return ret, err
}

// Get retrieves the AppService from the indexer for a given namespace and name.
func (s appServiceNamespaceLister) Get(name string) (*v1alpha1.AppService, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("appservice"), name)
	}
	return obj.(*v1alpha1.AppService), nil
}
