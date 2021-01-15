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

package v1alpha2

import (
	v1alpha2 "kubeform.dev/kubeform/apis/modules/v1alpha2"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SDPAzSqlv1Lister helps list SDPAzSqlv1s.
type SDPAzSqlv1Lister interface {
	// List lists all SDPAzSqlv1s in the indexer.
	List(selector labels.Selector) (ret []*v1alpha2.SDPAzSqlv1, err error)
	// SDPAzSqlv1s returns an object that can list and get SDPAzSqlv1s.
	SDPAzSqlv1s(namespace string) SDPAzSqlv1NamespaceLister
	SDPAzSqlv1ListerExpansion
}

// sDPAzSqlv1Lister implements the SDPAzSqlv1Lister interface.
type sDPAzSqlv1Lister struct {
	indexer cache.Indexer
}

// NewSDPAzSqlv1Lister returns a new SDPAzSqlv1Lister.
func NewSDPAzSqlv1Lister(indexer cache.Indexer) SDPAzSqlv1Lister {
	return &sDPAzSqlv1Lister{indexer: indexer}
}

// List lists all SDPAzSqlv1s in the indexer.
func (s *sDPAzSqlv1Lister) List(selector labels.Selector) (ret []*v1alpha2.SDPAzSqlv1, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.SDPAzSqlv1))
	})
	return ret, err
}

// SDPAzSqlv1s returns an object that can list and get SDPAzSqlv1s.
func (s *sDPAzSqlv1Lister) SDPAzSqlv1s(namespace string) SDPAzSqlv1NamespaceLister {
	return sDPAzSqlv1NamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SDPAzSqlv1NamespaceLister helps list and get SDPAzSqlv1s.
type SDPAzSqlv1NamespaceLister interface {
	// List lists all SDPAzSqlv1s in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha2.SDPAzSqlv1, err error)
	// Get retrieves the SDPAzSqlv1 from the indexer for a given namespace and name.
	Get(name string) (*v1alpha2.SDPAzSqlv1, error)
	SDPAzSqlv1NamespaceListerExpansion
}

// sDPAzSqlv1NamespaceLister implements the SDPAzSqlv1NamespaceLister
// interface.
type sDPAzSqlv1NamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SDPAzSqlv1s in the indexer for a given namespace.
func (s sDPAzSqlv1NamespaceLister) List(selector labels.Selector) (ret []*v1alpha2.SDPAzSqlv1, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.SDPAzSqlv1))
	})
	return ret, err
}

// Get retrieves the SDPAzSqlv1 from the indexer for a given namespace and name.
func (s sDPAzSqlv1NamespaceLister) Get(name string) (*v1alpha2.SDPAzSqlv1, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha2.Resource("sdpazsqlv1"), name)
	}
	return obj.(*v1alpha2.SDPAzSqlv1), nil
}
