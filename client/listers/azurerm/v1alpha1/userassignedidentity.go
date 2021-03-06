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

// UserAssignedIdentityLister helps list UserAssignedIdentities.
type UserAssignedIdentityLister interface {
	// List lists all UserAssignedIdentities in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.UserAssignedIdentity, err error)
	// UserAssignedIdentities returns an object that can list and get UserAssignedIdentities.
	UserAssignedIdentities(namespace string) UserAssignedIdentityNamespaceLister
	UserAssignedIdentityListerExpansion
}

// userAssignedIdentityLister implements the UserAssignedIdentityLister interface.
type userAssignedIdentityLister struct {
	indexer cache.Indexer
}

// NewUserAssignedIdentityLister returns a new UserAssignedIdentityLister.
func NewUserAssignedIdentityLister(indexer cache.Indexer) UserAssignedIdentityLister {
	return &userAssignedIdentityLister{indexer: indexer}
}

// List lists all UserAssignedIdentities in the indexer.
func (s *userAssignedIdentityLister) List(selector labels.Selector) (ret []*v1alpha1.UserAssignedIdentity, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.UserAssignedIdentity))
	})
	return ret, err
}

// UserAssignedIdentities returns an object that can list and get UserAssignedIdentities.
func (s *userAssignedIdentityLister) UserAssignedIdentities(namespace string) UserAssignedIdentityNamespaceLister {
	return userAssignedIdentityNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// UserAssignedIdentityNamespaceLister helps list and get UserAssignedIdentities.
type UserAssignedIdentityNamespaceLister interface {
	// List lists all UserAssignedIdentities in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.UserAssignedIdentity, err error)
	// Get retrieves the UserAssignedIdentity from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.UserAssignedIdentity, error)
	UserAssignedIdentityNamespaceListerExpansion
}

// userAssignedIdentityNamespaceLister implements the UserAssignedIdentityNamespaceLister
// interface.
type userAssignedIdentityNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all UserAssignedIdentities in the indexer for a given namespace.
func (s userAssignedIdentityNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.UserAssignedIdentity, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.UserAssignedIdentity))
	})
	return ret, err
}

// Get retrieves the UserAssignedIdentity from the indexer for a given namespace and name.
func (s userAssignedIdentityNamespaceLister) Get(name string) (*v1alpha1.UserAssignedIdentity, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("userassignedidentity"), name)
	}
	return obj.(*v1alpha1.UserAssignedIdentity), nil
}
