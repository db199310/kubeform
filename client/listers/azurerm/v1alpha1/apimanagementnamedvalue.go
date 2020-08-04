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

// ApiManagementNamedValueLister helps list ApiManagementNamedValues.
type ApiManagementNamedValueLister interface {
	// List lists all ApiManagementNamedValues in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ApiManagementNamedValue, err error)
	// ApiManagementNamedValues returns an object that can list and get ApiManagementNamedValues.
	ApiManagementNamedValues(namespace string) ApiManagementNamedValueNamespaceLister
	ApiManagementNamedValueListerExpansion
}

// apiManagementNamedValueLister implements the ApiManagementNamedValueLister interface.
type apiManagementNamedValueLister struct {
	indexer cache.Indexer
}

// NewApiManagementNamedValueLister returns a new ApiManagementNamedValueLister.
func NewApiManagementNamedValueLister(indexer cache.Indexer) ApiManagementNamedValueLister {
	return &apiManagementNamedValueLister{indexer: indexer}
}

// List lists all ApiManagementNamedValues in the indexer.
func (s *apiManagementNamedValueLister) List(selector labels.Selector) (ret []*v1alpha1.ApiManagementNamedValue, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiManagementNamedValue))
	})
	return ret, err
}

// ApiManagementNamedValues returns an object that can list and get ApiManagementNamedValues.
func (s *apiManagementNamedValueLister) ApiManagementNamedValues(namespace string) ApiManagementNamedValueNamespaceLister {
	return apiManagementNamedValueNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApiManagementNamedValueNamespaceLister helps list and get ApiManagementNamedValues.
type ApiManagementNamedValueNamespaceLister interface {
	// List lists all ApiManagementNamedValues in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ApiManagementNamedValue, err error)
	// Get retrieves the ApiManagementNamedValue from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ApiManagementNamedValue, error)
	ApiManagementNamedValueNamespaceListerExpansion
}

// apiManagementNamedValueNamespaceLister implements the ApiManagementNamedValueNamespaceLister
// interface.
type apiManagementNamedValueNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApiManagementNamedValues in the indexer for a given namespace.
func (s apiManagementNamedValueNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ApiManagementNamedValue, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiManagementNamedValue))
	})
	return ret, err
}

// Get retrieves the ApiManagementNamedValue from the indexer for a given namespace and name.
func (s apiManagementNamedValueNamespaceLister) Get(name string) (*v1alpha1.ApiManagementNamedValue, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("apimanagementnamedvalue"), name)
	}
	return obj.(*v1alpha1.ApiManagementNamedValue), nil
}
