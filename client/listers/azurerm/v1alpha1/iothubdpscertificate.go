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

// IothubDpsCertificateLister helps list IothubDpsCertificates.
type IothubDpsCertificateLister interface {
	// List lists all IothubDpsCertificates in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.IothubDpsCertificate, err error)
	// IothubDpsCertificates returns an object that can list and get IothubDpsCertificates.
	IothubDpsCertificates(namespace string) IothubDpsCertificateNamespaceLister
	IothubDpsCertificateListerExpansion
}

// iothubDpsCertificateLister implements the IothubDpsCertificateLister interface.
type iothubDpsCertificateLister struct {
	indexer cache.Indexer
}

// NewIothubDpsCertificateLister returns a new IothubDpsCertificateLister.
func NewIothubDpsCertificateLister(indexer cache.Indexer) IothubDpsCertificateLister {
	return &iothubDpsCertificateLister{indexer: indexer}
}

// List lists all IothubDpsCertificates in the indexer.
func (s *iothubDpsCertificateLister) List(selector labels.Selector) (ret []*v1alpha1.IothubDpsCertificate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IothubDpsCertificate))
	})
	return ret, err
}

// IothubDpsCertificates returns an object that can list and get IothubDpsCertificates.
func (s *iothubDpsCertificateLister) IothubDpsCertificates(namespace string) IothubDpsCertificateNamespaceLister {
	return iothubDpsCertificateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IothubDpsCertificateNamespaceLister helps list and get IothubDpsCertificates.
type IothubDpsCertificateNamespaceLister interface {
	// List lists all IothubDpsCertificates in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.IothubDpsCertificate, err error)
	// Get retrieves the IothubDpsCertificate from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.IothubDpsCertificate, error)
	IothubDpsCertificateNamespaceListerExpansion
}

// iothubDpsCertificateNamespaceLister implements the IothubDpsCertificateNamespaceLister
// interface.
type iothubDpsCertificateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IothubDpsCertificates in the indexer for a given namespace.
func (s iothubDpsCertificateNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IothubDpsCertificate, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IothubDpsCertificate))
	})
	return ret, err
}

// Get retrieves the IothubDpsCertificate from the indexer for a given namespace and name.
func (s iothubDpsCertificateNamespaceLister) Get(name string) (*v1alpha1.IothubDpsCertificate, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("iothubdpscertificate"), name)
	}
	return obj.(*v1alpha1.IothubDpsCertificate), nil
}
