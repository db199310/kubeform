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

// AzureadServicePrincipalPasswordLister helps list AzureadServicePrincipalPasswords.
type AzureadServicePrincipalPasswordLister interface {
	// List lists all AzureadServicePrincipalPasswords in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AzureadServicePrincipalPassword, err error)
	// AzureadServicePrincipalPasswords returns an object that can list and get AzureadServicePrincipalPasswords.
	AzureadServicePrincipalPasswords(namespace string) AzureadServicePrincipalPasswordNamespaceLister
	AzureadServicePrincipalPasswordListerExpansion
}

// azureadServicePrincipalPasswordLister implements the AzureadServicePrincipalPasswordLister interface.
type azureadServicePrincipalPasswordLister struct {
	indexer cache.Indexer
}

// NewAzureadServicePrincipalPasswordLister returns a new AzureadServicePrincipalPasswordLister.
func NewAzureadServicePrincipalPasswordLister(indexer cache.Indexer) AzureadServicePrincipalPasswordLister {
	return &azureadServicePrincipalPasswordLister{indexer: indexer}
}

// List lists all AzureadServicePrincipalPasswords in the indexer.
func (s *azureadServicePrincipalPasswordLister) List(selector labels.Selector) (ret []*v1alpha1.AzureadServicePrincipalPassword, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AzureadServicePrincipalPassword))
	})
	return ret, err
}

// AzureadServicePrincipalPasswords returns an object that can list and get AzureadServicePrincipalPasswords.
func (s *azureadServicePrincipalPasswordLister) AzureadServicePrincipalPasswords(namespace string) AzureadServicePrincipalPasswordNamespaceLister {
	return azureadServicePrincipalPasswordNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AzureadServicePrincipalPasswordNamespaceLister helps list and get AzureadServicePrincipalPasswords.
type AzureadServicePrincipalPasswordNamespaceLister interface {
	// List lists all AzureadServicePrincipalPasswords in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AzureadServicePrincipalPassword, err error)
	// Get retrieves the AzureadServicePrincipalPassword from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AzureadServicePrincipalPassword, error)
	AzureadServicePrincipalPasswordNamespaceListerExpansion
}

// azureadServicePrincipalPasswordNamespaceLister implements the AzureadServicePrincipalPasswordNamespaceLister
// interface.
type azureadServicePrincipalPasswordNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AzureadServicePrincipalPasswords in the indexer for a given namespace.
func (s azureadServicePrincipalPasswordNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AzureadServicePrincipalPassword, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AzureadServicePrincipalPassword))
	})
	return ret, err
}

// Get retrieves the AzureadServicePrincipalPassword from the indexer for a given namespace and name.
func (s azureadServicePrincipalPasswordNamespaceLister) Get(name string) (*v1alpha1.AzureadServicePrincipalPassword, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("azureadserviceprincipalpassword"), name)
	}
	return obj.(*v1alpha1.AzureadServicePrincipalPassword), nil
}
