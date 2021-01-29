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

// SDPAzappserviceplanv1Lister helps list SDPAzappserviceplanv1s.
type SDPAzappserviceplanv1Lister interface {
	// List lists all SDPAzappserviceplanv1s in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SDPAzappserviceplanv1, err error)
	// SDPAzappserviceplanv1s returns an object that can list and get SDPAzappserviceplanv1s.
	SDPAzappserviceplanv1s(namespace string) SDPAzappserviceplanv1NamespaceLister
	SDPAzappserviceplanv1ListerExpansion
}

// sDPAzappserviceplanv1Lister implements the SDPAzappserviceplanv1Lister interface.
type sDPAzappserviceplanv1Lister struct {
	indexer cache.Indexer
}

// NewSDPAzappserviceplanv1Lister returns a new SDPAzappserviceplanv1Lister.
func NewSDPAzappserviceplanv1Lister(indexer cache.Indexer) SDPAzappserviceplanv1Lister {
	return &sDPAzappserviceplanv1Lister{indexer: indexer}
}

// List lists all SDPAzappserviceplanv1s in the indexer.
func (s *sDPAzappserviceplanv1Lister) List(selector labels.Selector) (ret []*v1alpha1.SDPAzappserviceplanv1, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SDPAzappserviceplanv1))
	})
	return ret, err
}

// SDPAzappserviceplanv1s returns an object that can list and get SDPAzappserviceplanv1s.
func (s *sDPAzappserviceplanv1Lister) SDPAzappserviceplanv1s(namespace string) SDPAzappserviceplanv1NamespaceLister {
	return sDPAzappserviceplanv1NamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SDPAzappserviceplanv1NamespaceLister helps list and get SDPAzappserviceplanv1s.
type SDPAzappserviceplanv1NamespaceLister interface {
	// List lists all SDPAzappserviceplanv1s in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SDPAzappserviceplanv1, err error)
	// Get retrieves the SDPAzappserviceplanv1 from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SDPAzappserviceplanv1, error)
	SDPAzappserviceplanv1NamespaceListerExpansion
}

// sDPAzappserviceplanv1NamespaceLister implements the SDPAzappserviceplanv1NamespaceLister
// interface.
type sDPAzappserviceplanv1NamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SDPAzappserviceplanv1s in the indexer for a given namespace.
func (s sDPAzappserviceplanv1NamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SDPAzappserviceplanv1, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SDPAzappserviceplanv1))
	})
	return ret, err
}

// Get retrieves the SDPAzappserviceplanv1 from the indexer for a given namespace and name.
func (s sDPAzappserviceplanv1NamespaceLister) Get(name string) (*v1alpha1.SDPAzappserviceplanv1, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("sdpazappserviceplanv1"), name)
	}
	return obj.(*v1alpha1.SDPAzappserviceplanv1), nil
}
