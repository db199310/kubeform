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

// SDPAzStgv1Lister helps list SDPAzStgv1s.
type SDPAzStgv1Lister interface {
	// List lists all SDPAzStgv1s in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SDPAzStgv1, err error)
	// SDPAzStgv1s returns an object that can list and get SDPAzStgv1s.
	SDPAzStgv1s(namespace string) SDPAzStgv1NamespaceLister
	SDPAzStgv1ListerExpansion
}

// sDPAzStgv1Lister implements the SDPAzStgv1Lister interface.
type sDPAzStgv1Lister struct {
	indexer cache.Indexer
}

// NewSDPAzStgv1Lister returns a new SDPAzStgv1Lister.
func NewSDPAzStgv1Lister(indexer cache.Indexer) SDPAzStgv1Lister {
	return &sDPAzStgv1Lister{indexer: indexer}
}

// List lists all SDPAzStgv1s in the indexer.
func (s *sDPAzStgv1Lister) List(selector labels.Selector) (ret []*v1alpha1.SDPAzStgv1, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SDPAzStgv1))
	})
	return ret, err
}

// SDPAzStgv1s returns an object that can list and get SDPAzStgv1s.
func (s *sDPAzStgv1Lister) SDPAzStgv1s(namespace string) SDPAzStgv1NamespaceLister {
	return sDPAzStgv1NamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SDPAzStgv1NamespaceLister helps list and get SDPAzStgv1s.
type SDPAzStgv1NamespaceLister interface {
	// List lists all SDPAzStgv1s in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SDPAzStgv1, err error)
	// Get retrieves the SDPAzStgv1 from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SDPAzStgv1, error)
	SDPAzStgv1NamespaceListerExpansion
}

// sDPAzStgv1NamespaceLister implements the SDPAzStgv1NamespaceLister
// interface.
type sDPAzStgv1NamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SDPAzStgv1s in the indexer for a given namespace.
func (s sDPAzStgv1NamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SDPAzStgv1, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SDPAzStgv1))
	})
	return ret, err
}

// Get retrieves the SDPAzStgv1 from the indexer for a given namespace and name.
func (s sDPAzStgv1NamespaceLister) Get(name string) (*v1alpha1.SDPAzStgv1, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("sdpazstgv1"), name)
	}
	return obj.(*v1alpha1.SDPAzStgv1), nil
}
