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
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha2 "kubeform.dev/kubeform/apis/modules/v1alpha2"
)

// SDPAzFnv1Lister helps list SDPAzFnv1s.
type SDPAzFnv1Lister interface {
	// List lists all SDPAzFnv1s in the indexer.
	List(selector labels.Selector) (ret []*v1alpha2.SDPAzFnv1, err error)
	// SDPAzFnv1s returns an object that can list and get SDPAzFnv1s.
	SDPAzFnv1s(namespace string) SDPAzFnv1NamespaceLister
	SDPAzFnv1ListerExpansion
}

// sDPAzFnv1Lister implements the SDPAzFnv1Lister interface.
type sDPAzFnv1Lister struct {
	indexer cache.Indexer
}

// NewSDPAzFnv1Lister returns a new SDPAzFnv1Lister.
func NewSDPAzFnv1Lister(indexer cache.Indexer) SDPAzFnv1Lister {
	return &sDPAzFnv1Lister{indexer: indexer}
}

// List lists all SDPAzFnv1s in the indexer.
func (s *sDPAzFnv1Lister) List(selector labels.Selector) (ret []*v1alpha2.SDPAzFnv1, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.SDPAzFnv1))
	})
	return ret, err
}

// SDPAzFnv1s returns an object that can list and get SDPAzFnv1s.
func (s *sDPAzFnv1Lister) SDPAzFnv1s(namespace string) SDPAzFnv1NamespaceLister {
	return sDPAzFnv1NamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SDPAzFnv1NamespaceLister helps list and get SDPAzFnv1s.
type SDPAzFnv1NamespaceLister interface {
	// List lists all SDPAzFnv1s in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha2.SDPAzFnv1, err error)
	// Get retrieves the SDPAzFnv1 from the indexer for a given namespace and name.
	Get(name string) (*v1alpha2.SDPAzFnv1, error)
	SDPAzFnv1NamespaceListerExpansion
}

// sDPAzFnv1NamespaceLister implements the SDPAzFnv1NamespaceLister
// interface.
type sDPAzFnv1NamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SDPAzFnv1s in the indexer for a given namespace.
func (s sDPAzFnv1NamespaceLister) List(selector labels.Selector) (ret []*v1alpha2.SDPAzFnv1, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.SDPAzFnv1))
	})
	return ret, err
}

// Get retrieves the SDPAzFnv1 from the indexer for a given namespace and name.
func (s sDPAzFnv1NamespaceLister) Get(name string) (*v1alpha2.SDPAzFnv1, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha2.Resource("sdpazfnv1"), name)
	}
	return obj.(*v1alpha2.SDPAzFnv1), nil
}
