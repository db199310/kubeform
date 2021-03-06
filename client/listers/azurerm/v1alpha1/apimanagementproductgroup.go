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

// ApiManagementProductGroupLister helps list ApiManagementProductGroups.
type ApiManagementProductGroupLister interface {
	// List lists all ApiManagementProductGroups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ApiManagementProductGroup, err error)
	// ApiManagementProductGroups returns an object that can list and get ApiManagementProductGroups.
	ApiManagementProductGroups(namespace string) ApiManagementProductGroupNamespaceLister
	ApiManagementProductGroupListerExpansion
}

// apiManagementProductGroupLister implements the ApiManagementProductGroupLister interface.
type apiManagementProductGroupLister struct {
	indexer cache.Indexer
}

// NewApiManagementProductGroupLister returns a new ApiManagementProductGroupLister.
func NewApiManagementProductGroupLister(indexer cache.Indexer) ApiManagementProductGroupLister {
	return &apiManagementProductGroupLister{indexer: indexer}
}

// List lists all ApiManagementProductGroups in the indexer.
func (s *apiManagementProductGroupLister) List(selector labels.Selector) (ret []*v1alpha1.ApiManagementProductGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiManagementProductGroup))
	})
	return ret, err
}

// ApiManagementProductGroups returns an object that can list and get ApiManagementProductGroups.
func (s *apiManagementProductGroupLister) ApiManagementProductGroups(namespace string) ApiManagementProductGroupNamespaceLister {
	return apiManagementProductGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApiManagementProductGroupNamespaceLister helps list and get ApiManagementProductGroups.
type ApiManagementProductGroupNamespaceLister interface {
	// List lists all ApiManagementProductGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ApiManagementProductGroup, err error)
	// Get retrieves the ApiManagementProductGroup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ApiManagementProductGroup, error)
	ApiManagementProductGroupNamespaceListerExpansion
}

// apiManagementProductGroupNamespaceLister implements the ApiManagementProductGroupNamespaceLister
// interface.
type apiManagementProductGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApiManagementProductGroups in the indexer for a given namespace.
func (s apiManagementProductGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ApiManagementProductGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiManagementProductGroup))
	})
	return ret, err
}

// Get retrieves the ApiManagementProductGroup from the indexer for a given namespace and name.
func (s apiManagementProductGroupNamespaceLister) Get(name string) (*v1alpha1.ApiManagementProductGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("apimanagementproductgroup"), name)
	}
	return obj.(*v1alpha1.ApiManagementProductGroup), nil
}
