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

// PublicIPPrefixLister helps list PublicIPPrefixes.
type PublicIPPrefixLister interface {
	// List lists all PublicIPPrefixes in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.PublicIPPrefix, err error)
	// PublicIPPrefixes returns an object that can list and get PublicIPPrefixes.
	PublicIPPrefixes(namespace string) PublicIPPrefixNamespaceLister
	PublicIPPrefixListerExpansion
}

// publicIPPrefixLister implements the PublicIPPrefixLister interface.
type publicIPPrefixLister struct {
	indexer cache.Indexer
}

// NewPublicIPPrefixLister returns a new PublicIPPrefixLister.
func NewPublicIPPrefixLister(indexer cache.Indexer) PublicIPPrefixLister {
	return &publicIPPrefixLister{indexer: indexer}
}

// List lists all PublicIPPrefixes in the indexer.
func (s *publicIPPrefixLister) List(selector labels.Selector) (ret []*v1alpha1.PublicIPPrefix, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PublicIPPrefix))
	})
	return ret, err
}

// PublicIPPrefixes returns an object that can list and get PublicIPPrefixes.
func (s *publicIPPrefixLister) PublicIPPrefixes(namespace string) PublicIPPrefixNamespaceLister {
	return publicIPPrefixNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PublicIPPrefixNamespaceLister helps list and get PublicIPPrefixes.
type PublicIPPrefixNamespaceLister interface {
	// List lists all PublicIPPrefixes in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.PublicIPPrefix, err error)
	// Get retrieves the PublicIPPrefix from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.PublicIPPrefix, error)
	PublicIPPrefixNamespaceListerExpansion
}

// publicIPPrefixNamespaceLister implements the PublicIPPrefixNamespaceLister
// interface.
type publicIPPrefixNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PublicIPPrefixes in the indexer for a given namespace.
func (s publicIPPrefixNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PublicIPPrefix, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PublicIPPrefix))
	})
	return ret, err
}

// Get retrieves the PublicIPPrefix from the indexer for a given namespace and name.
func (s publicIPPrefixNamespaceLister) Get(name string) (*v1alpha1.PublicIPPrefix, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("publicipprefix"), name)
	}
	return obj.(*v1alpha1.PublicIPPrefix), nil
}
