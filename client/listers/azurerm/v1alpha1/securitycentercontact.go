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

// SecurityCenterContactLister helps list SecurityCenterContacts.
type SecurityCenterContactLister interface {
	// List lists all SecurityCenterContacts in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SecurityCenterContact, err error)
	// SecurityCenterContacts returns an object that can list and get SecurityCenterContacts.
	SecurityCenterContacts(namespace string) SecurityCenterContactNamespaceLister
	SecurityCenterContactListerExpansion
}

// securityCenterContactLister implements the SecurityCenterContactLister interface.
type securityCenterContactLister struct {
	indexer cache.Indexer
}

// NewSecurityCenterContactLister returns a new SecurityCenterContactLister.
func NewSecurityCenterContactLister(indexer cache.Indexer) SecurityCenterContactLister {
	return &securityCenterContactLister{indexer: indexer}
}

// List lists all SecurityCenterContacts in the indexer.
func (s *securityCenterContactLister) List(selector labels.Selector) (ret []*v1alpha1.SecurityCenterContact, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SecurityCenterContact))
	})
	return ret, err
}

// SecurityCenterContacts returns an object that can list and get SecurityCenterContacts.
func (s *securityCenterContactLister) SecurityCenterContacts(namespace string) SecurityCenterContactNamespaceLister {
	return securityCenterContactNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SecurityCenterContactNamespaceLister helps list and get SecurityCenterContacts.
type SecurityCenterContactNamespaceLister interface {
	// List lists all SecurityCenterContacts in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SecurityCenterContact, err error)
	// Get retrieves the SecurityCenterContact from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SecurityCenterContact, error)
	SecurityCenterContactNamespaceListerExpansion
}

// securityCenterContactNamespaceLister implements the SecurityCenterContactNamespaceLister
// interface.
type securityCenterContactNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SecurityCenterContacts in the indexer for a given namespace.
func (s securityCenterContactNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SecurityCenterContact, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SecurityCenterContact))
	})
	return ret, err
}

// Get retrieves the SecurityCenterContact from the indexer for a given namespace and name.
func (s securityCenterContactNamespaceLister) Get(name string) (*v1alpha1.SecurityCenterContact, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("securitycentercontact"), name)
	}
	return obj.(*v1alpha1.SecurityCenterContact), nil
}
