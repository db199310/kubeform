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

// KeyVaultCertificateLister helps list KeyVaultCertificates.
type KeyVaultCertificateLister interface {
	// List lists all KeyVaultCertificates in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.KeyVaultCertificate, err error)
	// KeyVaultCertificates returns an object that can list and get KeyVaultCertificates.
	KeyVaultCertificates(namespace string) KeyVaultCertificateNamespaceLister
	KeyVaultCertificateListerExpansion
}

// keyVaultCertificateLister implements the KeyVaultCertificateLister interface.
type keyVaultCertificateLister struct {
	indexer cache.Indexer
}

// NewKeyVaultCertificateLister returns a new KeyVaultCertificateLister.
func NewKeyVaultCertificateLister(indexer cache.Indexer) KeyVaultCertificateLister {
	return &keyVaultCertificateLister{indexer: indexer}
}

// List lists all KeyVaultCertificates in the indexer.
func (s *keyVaultCertificateLister) List(selector labels.Selector) (ret []*v1alpha1.KeyVaultCertificate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KeyVaultCertificate))
	})
	return ret, err
}

// KeyVaultCertificates returns an object that can list and get KeyVaultCertificates.
func (s *keyVaultCertificateLister) KeyVaultCertificates(namespace string) KeyVaultCertificateNamespaceLister {
	return keyVaultCertificateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// KeyVaultCertificateNamespaceLister helps list and get KeyVaultCertificates.
type KeyVaultCertificateNamespaceLister interface {
	// List lists all KeyVaultCertificates in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.KeyVaultCertificate, err error)
	// Get retrieves the KeyVaultCertificate from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.KeyVaultCertificate, error)
	KeyVaultCertificateNamespaceListerExpansion
}

// keyVaultCertificateNamespaceLister implements the KeyVaultCertificateNamespaceLister
// interface.
type keyVaultCertificateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all KeyVaultCertificates in the indexer for a given namespace.
func (s keyVaultCertificateNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.KeyVaultCertificate, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KeyVaultCertificate))
	})
	return ret, err
}

// Get retrieves the KeyVaultCertificate from the indexer for a given namespace and name.
func (s keyVaultCertificateNamespaceLister) Get(name string) (*v1alpha1.KeyVaultCertificate, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("keyvaultcertificate"), name)
	}
	return obj.(*v1alpha1.KeyVaultCertificate), nil
}
