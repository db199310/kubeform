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

// SDPAzsbv1Lister helps list SDPAzsbv1s.
type SDPAzsbv1Lister interface {
	// List lists all SDPAzsbv1s in the indexer.
	List(selector labels.Selector) (ret []*v1alpha2.SDPAzsbv1, err error)
	// SDPAzsbv1s returns an object that can list and get SDPAzsbv1s.
	SDPAzsbv1s(namespace string) SDPAzsbv1NamespaceLister
	SDPAzsbv1ListerExpansion
}

// sDPAzsbv1Lister implements the SDPAzsbv1Lister interface.
type sDPAzsbv1Lister struct {
	indexer cache.Indexer
}

// NewSDPAzsbv1Lister returns a new SDPAzsbv1Lister.
func NewSDPAzsbv1Lister(indexer cache.Indexer) SDPAzsbv1Lister {
	return &sDPAzsbv1Lister{indexer: indexer}
}

// List lists all SDPAzsbv1s in the indexer.
func (s *sDPAzsbv1Lister) List(selector labels.Selector) (ret []*v1alpha2.SDPAzsbv1, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.SDPAzsbv1))
	})
	return ret, err
}

// SDPAzsbv1s returns an object that can list and get SDPAzsbv1s.
func (s *sDPAzsbv1Lister) SDPAzsbv1s(namespace string) SDPAzsbv1NamespaceLister {
	return sDPAzsbv1NamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SDPAzsbv1NamespaceLister helps list and get SDPAzsbv1s.
type SDPAzsbv1NamespaceLister interface {
	// List lists all SDPAzsbv1s in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha2.SDPAzsbv1, err error)
	// Get retrieves the SDPAzsbv1 from the indexer for a given namespace and name.
	Get(name string) (*v1alpha2.SDPAzsbv1, error)
	SDPAzsbv1NamespaceListerExpansion
}

// sDPAzsbv1NamespaceLister implements the SDPAzsbv1NamespaceLister
// interface.
type sDPAzsbv1NamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SDPAzsbv1s in the indexer for a given namespace.
func (s sDPAzsbv1NamespaceLister) List(selector labels.Selector) (ret []*v1alpha2.SDPAzsbv1, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.SDPAzsbv1))
	})
	return ret, err
}

// Get retrieves the SDPAzsbv1 from the indexer for a given namespace and name.
func (s sDPAzsbv1NamespaceLister) Get(name string) (*v1alpha2.SDPAzsbv1, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha2.Resource("sdpazsbv1"), name)
	}
	return obj.(*v1alpha2.SDPAzsbv1), nil
}
