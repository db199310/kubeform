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

// MariadbDatabaseLister helps list MariadbDatabases.
type MariadbDatabaseLister interface {
	// List lists all MariadbDatabases in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.MariadbDatabase, err error)
	// MariadbDatabases returns an object that can list and get MariadbDatabases.
	MariadbDatabases(namespace string) MariadbDatabaseNamespaceLister
	MariadbDatabaseListerExpansion
}

// mariadbDatabaseLister implements the MariadbDatabaseLister interface.
type mariadbDatabaseLister struct {
	indexer cache.Indexer
}

// NewMariadbDatabaseLister returns a new MariadbDatabaseLister.
func NewMariadbDatabaseLister(indexer cache.Indexer) MariadbDatabaseLister {
	return &mariadbDatabaseLister{indexer: indexer}
}

// List lists all MariadbDatabases in the indexer.
func (s *mariadbDatabaseLister) List(selector labels.Selector) (ret []*v1alpha1.MariadbDatabase, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MariadbDatabase))
	})
	return ret, err
}

// MariadbDatabases returns an object that can list and get MariadbDatabases.
func (s *mariadbDatabaseLister) MariadbDatabases(namespace string) MariadbDatabaseNamespaceLister {
	return mariadbDatabaseNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MariadbDatabaseNamespaceLister helps list and get MariadbDatabases.
type MariadbDatabaseNamespaceLister interface {
	// List lists all MariadbDatabases in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.MariadbDatabase, err error)
	// Get retrieves the MariadbDatabase from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.MariadbDatabase, error)
	MariadbDatabaseNamespaceListerExpansion
}

// mariadbDatabaseNamespaceLister implements the MariadbDatabaseNamespaceLister
// interface.
type mariadbDatabaseNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MariadbDatabases in the indexer for a given namespace.
func (s mariadbDatabaseNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MariadbDatabase, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MariadbDatabase))
	})
	return ret, err
}

// Get retrieves the MariadbDatabase from the indexer for a given namespace and name.
func (s mariadbDatabaseNamespaceLister) Get(name string) (*v1alpha1.MariadbDatabase, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("mariadbdatabase"), name)
	}
	return obj.(*v1alpha1.MariadbDatabase), nil
}
