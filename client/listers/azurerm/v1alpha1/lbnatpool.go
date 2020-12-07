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

// LbNATPoolLister helps list LbNATPools.
type LbNATPoolLister interface {
	// List lists all LbNATPools in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.LbNATPool, err error)
	// LbNATPools returns an object that can list and get LbNATPools.
	LbNATPools(namespace string) LbNATPoolNamespaceLister
	LbNATPoolListerExpansion
}

// lbNATPoolLister implements the LbNATPoolLister interface.
type lbNATPoolLister struct {
	indexer cache.Indexer
}

// NewLbNATPoolLister returns a new LbNATPoolLister.
func NewLbNATPoolLister(indexer cache.Indexer) LbNATPoolLister {
	return &lbNATPoolLister{indexer: indexer}
}

// List lists all LbNATPools in the indexer.
func (s *lbNATPoolLister) List(selector labels.Selector) (ret []*v1alpha1.LbNATPool, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LbNATPool))
	})
	return ret, err
}

// LbNATPools returns an object that can list and get LbNATPools.
func (s *lbNATPoolLister) LbNATPools(namespace string) LbNATPoolNamespaceLister {
	return lbNATPoolNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LbNATPoolNamespaceLister helps list and get LbNATPools.
type LbNATPoolNamespaceLister interface {
	// List lists all LbNATPools in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.LbNATPool, err error)
	// Get retrieves the LbNATPool from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.LbNATPool, error)
	LbNATPoolNamespaceListerExpansion
}

// lbNATPoolNamespaceLister implements the LbNATPoolNamespaceLister
// interface.
type lbNATPoolNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LbNATPools in the indexer for a given namespace.
func (s lbNATPoolNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LbNATPool, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LbNATPool))
	})
	return ret, err
}

// Get retrieves the LbNATPool from the indexer for a given namespace and name.
func (s lbNATPoolNamespaceLister) Get(name string) (*v1alpha1.LbNATPool, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("lbnatpool"), name)
	}
	return obj.(*v1alpha1.LbNATPool), nil
}
