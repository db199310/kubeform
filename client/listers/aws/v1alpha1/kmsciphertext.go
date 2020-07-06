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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// KmsCiphertextLister helps list KmsCiphertexts.
type KmsCiphertextLister interface {
	// List lists all KmsCiphertexts in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.KmsCiphertext, err error)
	// KmsCiphertexts returns an object that can list and get KmsCiphertexts.
	KmsCiphertexts(namespace string) KmsCiphertextNamespaceLister
	KmsCiphertextListerExpansion
}

// kmsCiphertextLister implements the KmsCiphertextLister interface.
type kmsCiphertextLister struct {
	indexer cache.Indexer
}

// NewKmsCiphertextLister returns a new KmsCiphertextLister.
func NewKmsCiphertextLister(indexer cache.Indexer) KmsCiphertextLister {
	return &kmsCiphertextLister{indexer: indexer}
}

// List lists all KmsCiphertexts in the indexer.
func (s *kmsCiphertextLister) List(selector labels.Selector) (ret []*v1alpha1.KmsCiphertext, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KmsCiphertext))
	})
	return ret, err
}

// KmsCiphertexts returns an object that can list and get KmsCiphertexts.
func (s *kmsCiphertextLister) KmsCiphertexts(namespace string) KmsCiphertextNamespaceLister {
	return kmsCiphertextNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// KmsCiphertextNamespaceLister helps list and get KmsCiphertexts.
type KmsCiphertextNamespaceLister interface {
	// List lists all KmsCiphertexts in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.KmsCiphertext, err error)
	// Get retrieves the KmsCiphertext from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.KmsCiphertext, error)
	KmsCiphertextNamespaceListerExpansion
}

// kmsCiphertextNamespaceLister implements the KmsCiphertextNamespaceLister
// interface.
type kmsCiphertextNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all KmsCiphertexts in the indexer for a given namespace.
func (s kmsCiphertextNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.KmsCiphertext, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KmsCiphertext))
	})
	return ret, err
}

// Get retrieves the KmsCiphertext from the indexer for a given namespace and name.
func (s kmsCiphertextNamespaceLister) Get(name string) (*v1alpha1.KmsCiphertext, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("kmsciphertext"), name)
	}
	return obj.(*v1alpha1.KmsCiphertext), nil
}
