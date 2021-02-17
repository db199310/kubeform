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

// PrivateLinkServiceLister helps list PrivateLinkServices.
type PrivateLinkServiceLister interface {
	// List lists all PrivateLinkServices in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.PrivateLinkService, err error)
	// PrivateLinkServices returns an object that can list and get PrivateLinkServices.
	PrivateLinkServices(namespace string) PrivateLinkServiceNamespaceLister
	PrivateLinkServiceListerExpansion
}

// privateLinkServiceLister implements the PrivateLinkServiceLister interface.
type privateLinkServiceLister struct {
	indexer cache.Indexer
}

// NewPrivateLinkServiceLister returns a new PrivateLinkServiceLister.
func NewPrivateLinkServiceLister(indexer cache.Indexer) PrivateLinkServiceLister {
	return &privateLinkServiceLister{indexer: indexer}
}

// List lists all PrivateLinkServices in the indexer.
func (s *privateLinkServiceLister) List(selector labels.Selector) (ret []*v1alpha1.PrivateLinkService, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PrivateLinkService))
	})
	return ret, err
}

// PrivateLinkServices returns an object that can list and get PrivateLinkServices.
func (s *privateLinkServiceLister) PrivateLinkServices(namespace string) PrivateLinkServiceNamespaceLister {
	return privateLinkServiceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PrivateLinkServiceNamespaceLister helps list and get PrivateLinkServices.
type PrivateLinkServiceNamespaceLister interface {
	// List lists all PrivateLinkServices in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.PrivateLinkService, err error)
	// Get retrieves the PrivateLinkService from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.PrivateLinkService, error)
	PrivateLinkServiceNamespaceListerExpansion
}

// privateLinkServiceNamespaceLister implements the PrivateLinkServiceNamespaceLister
// interface.
type privateLinkServiceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PrivateLinkServices in the indexer for a given namespace.
func (s privateLinkServiceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PrivateLinkService, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PrivateLinkService))
	})
	return ret, err
}

// Get retrieves the PrivateLinkService from the indexer for a given namespace and name.
func (s privateLinkServiceNamespaceLister) Get(name string) (*v1alpha1.PrivateLinkService, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("privatelinkservice"), name)
	}
	return obj.(*v1alpha1.PrivateLinkService), nil
}
