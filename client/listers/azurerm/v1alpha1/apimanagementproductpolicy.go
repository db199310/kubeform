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

// ApiManagementProductPolicyLister helps list ApiManagementProductPolicies.
type ApiManagementProductPolicyLister interface {
	// List lists all ApiManagementProductPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ApiManagementProductPolicy, err error)
	// ApiManagementProductPolicies returns an object that can list and get ApiManagementProductPolicies.
	ApiManagementProductPolicies(namespace string) ApiManagementProductPolicyNamespaceLister
	ApiManagementProductPolicyListerExpansion
}

// apiManagementProductPolicyLister implements the ApiManagementProductPolicyLister interface.
type apiManagementProductPolicyLister struct {
	indexer cache.Indexer
}

// NewApiManagementProductPolicyLister returns a new ApiManagementProductPolicyLister.
func NewApiManagementProductPolicyLister(indexer cache.Indexer) ApiManagementProductPolicyLister {
	return &apiManagementProductPolicyLister{indexer: indexer}
}

// List lists all ApiManagementProductPolicies in the indexer.
func (s *apiManagementProductPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.ApiManagementProductPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiManagementProductPolicy))
	})
	return ret, err
}

// ApiManagementProductPolicies returns an object that can list and get ApiManagementProductPolicies.
func (s *apiManagementProductPolicyLister) ApiManagementProductPolicies(namespace string) ApiManagementProductPolicyNamespaceLister {
	return apiManagementProductPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApiManagementProductPolicyNamespaceLister helps list and get ApiManagementProductPolicies.
type ApiManagementProductPolicyNamespaceLister interface {
	// List lists all ApiManagementProductPolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ApiManagementProductPolicy, err error)
	// Get retrieves the ApiManagementProductPolicy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ApiManagementProductPolicy, error)
	ApiManagementProductPolicyNamespaceListerExpansion
}

// apiManagementProductPolicyNamespaceLister implements the ApiManagementProductPolicyNamespaceLister
// interface.
type apiManagementProductPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApiManagementProductPolicies in the indexer for a given namespace.
func (s apiManagementProductPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ApiManagementProductPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiManagementProductPolicy))
	})
	return ret, err
}

// Get retrieves the ApiManagementProductPolicy from the indexer for a given namespace and name.
func (s apiManagementProductPolicyNamespaceLister) Get(name string) (*v1alpha1.ApiManagementProductPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("apimanagementproductpolicy"), name)
	}
	return obj.(*v1alpha1.ApiManagementProductPolicy), nil
}
