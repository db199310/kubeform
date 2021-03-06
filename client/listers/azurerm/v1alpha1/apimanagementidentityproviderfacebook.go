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

// ApiManagementIdentityProviderFacebookLister helps list ApiManagementIdentityProviderFacebooks.
type ApiManagementIdentityProviderFacebookLister interface {
	// List lists all ApiManagementIdentityProviderFacebooks in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ApiManagementIdentityProviderFacebook, err error)
	// ApiManagementIdentityProviderFacebooks returns an object that can list and get ApiManagementIdentityProviderFacebooks.
	ApiManagementIdentityProviderFacebooks(namespace string) ApiManagementIdentityProviderFacebookNamespaceLister
	ApiManagementIdentityProviderFacebookListerExpansion
}

// apiManagementIdentityProviderFacebookLister implements the ApiManagementIdentityProviderFacebookLister interface.
type apiManagementIdentityProviderFacebookLister struct {
	indexer cache.Indexer
}

// NewApiManagementIdentityProviderFacebookLister returns a new ApiManagementIdentityProviderFacebookLister.
func NewApiManagementIdentityProviderFacebookLister(indexer cache.Indexer) ApiManagementIdentityProviderFacebookLister {
	return &apiManagementIdentityProviderFacebookLister{indexer: indexer}
}

// List lists all ApiManagementIdentityProviderFacebooks in the indexer.
func (s *apiManagementIdentityProviderFacebookLister) List(selector labels.Selector) (ret []*v1alpha1.ApiManagementIdentityProviderFacebook, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiManagementIdentityProviderFacebook))
	})
	return ret, err
}

// ApiManagementIdentityProviderFacebooks returns an object that can list and get ApiManagementIdentityProviderFacebooks.
func (s *apiManagementIdentityProviderFacebookLister) ApiManagementIdentityProviderFacebooks(namespace string) ApiManagementIdentityProviderFacebookNamespaceLister {
	return apiManagementIdentityProviderFacebookNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApiManagementIdentityProviderFacebookNamespaceLister helps list and get ApiManagementIdentityProviderFacebooks.
type ApiManagementIdentityProviderFacebookNamespaceLister interface {
	// List lists all ApiManagementIdentityProviderFacebooks in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ApiManagementIdentityProviderFacebook, err error)
	// Get retrieves the ApiManagementIdentityProviderFacebook from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ApiManagementIdentityProviderFacebook, error)
	ApiManagementIdentityProviderFacebookNamespaceListerExpansion
}

// apiManagementIdentityProviderFacebookNamespaceLister implements the ApiManagementIdentityProviderFacebookNamespaceLister
// interface.
type apiManagementIdentityProviderFacebookNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApiManagementIdentityProviderFacebooks in the indexer for a given namespace.
func (s apiManagementIdentityProviderFacebookNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ApiManagementIdentityProviderFacebook, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiManagementIdentityProviderFacebook))
	})
	return ret, err
}

// Get retrieves the ApiManagementIdentityProviderFacebook from the indexer for a given namespace and name.
func (s apiManagementIdentityProviderFacebookNamespaceLister) Get(name string) (*v1alpha1.ApiManagementIdentityProviderFacebook, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("apimanagementidentityproviderfacebook"), name)
	}
	return obj.(*v1alpha1.ApiManagementIdentityProviderFacebook), nil
}
