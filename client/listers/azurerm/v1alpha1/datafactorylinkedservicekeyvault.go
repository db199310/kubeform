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

// DataFactoryLinkedServiceKeyVaultLister helps list DataFactoryLinkedServiceKeyVaults.
type DataFactoryLinkedServiceKeyVaultLister interface {
	// List lists all DataFactoryLinkedServiceKeyVaults in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DataFactoryLinkedServiceKeyVault, err error)
	// DataFactoryLinkedServiceKeyVaults returns an object that can list and get DataFactoryLinkedServiceKeyVaults.
	DataFactoryLinkedServiceKeyVaults(namespace string) DataFactoryLinkedServiceKeyVaultNamespaceLister
	DataFactoryLinkedServiceKeyVaultListerExpansion
}

// dataFactoryLinkedServiceKeyVaultLister implements the DataFactoryLinkedServiceKeyVaultLister interface.
type dataFactoryLinkedServiceKeyVaultLister struct {
	indexer cache.Indexer
}

// NewDataFactoryLinkedServiceKeyVaultLister returns a new DataFactoryLinkedServiceKeyVaultLister.
func NewDataFactoryLinkedServiceKeyVaultLister(indexer cache.Indexer) DataFactoryLinkedServiceKeyVaultLister {
	return &dataFactoryLinkedServiceKeyVaultLister{indexer: indexer}
}

// List lists all DataFactoryLinkedServiceKeyVaults in the indexer.
func (s *dataFactoryLinkedServiceKeyVaultLister) List(selector labels.Selector) (ret []*v1alpha1.DataFactoryLinkedServiceKeyVault, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DataFactoryLinkedServiceKeyVault))
	})
	return ret, err
}

// DataFactoryLinkedServiceKeyVaults returns an object that can list and get DataFactoryLinkedServiceKeyVaults.
func (s *dataFactoryLinkedServiceKeyVaultLister) DataFactoryLinkedServiceKeyVaults(namespace string) DataFactoryLinkedServiceKeyVaultNamespaceLister {
	return dataFactoryLinkedServiceKeyVaultNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DataFactoryLinkedServiceKeyVaultNamespaceLister helps list and get DataFactoryLinkedServiceKeyVaults.
type DataFactoryLinkedServiceKeyVaultNamespaceLister interface {
	// List lists all DataFactoryLinkedServiceKeyVaults in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DataFactoryLinkedServiceKeyVault, err error)
	// Get retrieves the DataFactoryLinkedServiceKeyVault from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DataFactoryLinkedServiceKeyVault, error)
	DataFactoryLinkedServiceKeyVaultNamespaceListerExpansion
}

// dataFactoryLinkedServiceKeyVaultNamespaceLister implements the DataFactoryLinkedServiceKeyVaultNamespaceLister
// interface.
type dataFactoryLinkedServiceKeyVaultNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DataFactoryLinkedServiceKeyVaults in the indexer for a given namespace.
func (s dataFactoryLinkedServiceKeyVaultNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DataFactoryLinkedServiceKeyVault, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DataFactoryLinkedServiceKeyVault))
	})
	return ret, err
}

// Get retrieves the DataFactoryLinkedServiceKeyVault from the indexer for a given namespace and name.
func (s dataFactoryLinkedServiceKeyVaultNamespaceLister) Get(name string) (*v1alpha1.DataFactoryLinkedServiceKeyVault, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("datafactorylinkedservicekeyvault"), name)
	}
	return obj.(*v1alpha1.DataFactoryLinkedServiceKeyVault), nil
}