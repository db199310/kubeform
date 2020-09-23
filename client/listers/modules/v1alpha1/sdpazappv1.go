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
	v1alpha1 "kubeform.dev/kubeform/apis/modules/v1alpha1"
)

// SDPAzAppv1Lister helps list SDPAzAppv1s.
type SDPAzAppv1Lister interface {
	// List lists all SDPAzAppv1s in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SDPAzAppv1, err error)
	// SDPAzAppv1s returns an object that can list and get SDPAzAppv1s.
	SDPAzAppv1s(namespace string) SDPAzAppv1NamespaceLister
	SDPAzAppv1ListerExpansion
}

// sDPAzAppv1Lister implements the SDPAzAppv1Lister interface.
type sDPAzAppv1Lister struct {
	indexer cache.Indexer
}

// NewSDPAzAppv1Lister returns a new SDPAzAppv1Lister.
func NewSDPAzAppv1Lister(indexer cache.Indexer) SDPAzAppv1Lister {
	return &sDPAzAppv1Lister{indexer: indexer}
}

// List lists all SDPAzAppv1s in the indexer.
func (s *sDPAzAppv1Lister) List(selector labels.Selector) (ret []*v1alpha1.SDPAzAppv1, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SDPAzAppv1))
	})
	return ret, err
}

// SDPAzAppv1s returns an object that can list and get SDPAzAppv1s.
func (s *sDPAzAppv1Lister) SDPAzAppv1s(namespace string) SDPAzAppv1NamespaceLister {
	return sDPAzAppv1NamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SDPAzAppv1NamespaceLister helps list and get SDPAzAppv1s.
type SDPAzAppv1NamespaceLister interface {
	// List lists all SDPAzAppv1s in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SDPAzAppv1, err error)
	// Get retrieves the SDPAzAppv1 from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SDPAzAppv1, error)
	SDPAzAppv1NamespaceListerExpansion
}

// sDPAzAppv1NamespaceLister implements the SDPAzAppv1NamespaceLister
// interface.
type sDPAzAppv1NamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SDPAzAppv1s in the indexer for a given namespace.
func (s sDPAzAppv1NamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SDPAzAppv1, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SDPAzAppv1))
	})
	return ret, err
}

// Get retrieves the SDPAzAppv1 from the indexer for a given namespace and name.
func (s sDPAzAppv1NamespaceLister) Get(name string) (*v1alpha1.SDPAzAppv1, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("sdpazappv1"), name)
	}
	return obj.(*v1alpha1.SDPAzAppv1), nil
}