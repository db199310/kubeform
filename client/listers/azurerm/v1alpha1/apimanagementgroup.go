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

// ApiManagementGroupLister helps list ApiManagementGroups.
type ApiManagementGroupLister interface {
	// List lists all ApiManagementGroups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ApiManagementGroup, err error)
	// ApiManagementGroups returns an object that can list and get ApiManagementGroups.
	ApiManagementGroups(namespace string) ApiManagementGroupNamespaceLister
	ApiManagementGroupListerExpansion
}

// apiManagementGroupLister implements the ApiManagementGroupLister interface.
type apiManagementGroupLister struct {
	indexer cache.Indexer
}

// NewApiManagementGroupLister returns a new ApiManagementGroupLister.
func NewApiManagementGroupLister(indexer cache.Indexer) ApiManagementGroupLister {
	return &apiManagementGroupLister{indexer: indexer}
}

// List lists all ApiManagementGroups in the indexer.
func (s *apiManagementGroupLister) List(selector labels.Selector) (ret []*v1alpha1.ApiManagementGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiManagementGroup))
	})
	return ret, err
}

// ApiManagementGroups returns an object that can list and get ApiManagementGroups.
func (s *apiManagementGroupLister) ApiManagementGroups(namespace string) ApiManagementGroupNamespaceLister {
	return apiManagementGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApiManagementGroupNamespaceLister helps list and get ApiManagementGroups.
type ApiManagementGroupNamespaceLister interface {
	// List lists all ApiManagementGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ApiManagementGroup, err error)
	// Get retrieves the ApiManagementGroup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ApiManagementGroup, error)
	ApiManagementGroupNamespaceListerExpansion
}

// apiManagementGroupNamespaceLister implements the ApiManagementGroupNamespaceLister
// interface.
type apiManagementGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApiManagementGroups in the indexer for a given namespace.
func (s apiManagementGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ApiManagementGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiManagementGroup))
	})
	return ret, err
}

// Get retrieves the ApiManagementGroup from the indexer for a given namespace and name.
func (s apiManagementGroupNamespaceLister) Get(name string) (*v1alpha1.ApiManagementGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("apimanagementgroup"), name)
	}
	return obj.(*v1alpha1.ApiManagementGroup), nil
}
