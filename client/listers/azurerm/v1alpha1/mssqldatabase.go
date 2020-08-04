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

// MssqlDatabaseLister helps list MssqlDatabases.
type MssqlDatabaseLister interface {
	// List lists all MssqlDatabases in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.MssqlDatabase, err error)
	// MssqlDatabases returns an object that can list and get MssqlDatabases.
	MssqlDatabases(namespace string) MssqlDatabaseNamespaceLister
	MssqlDatabaseListerExpansion
}

// mssqlDatabaseLister implements the MssqlDatabaseLister interface.
type mssqlDatabaseLister struct {
	indexer cache.Indexer
}

// NewMssqlDatabaseLister returns a new MssqlDatabaseLister.
func NewMssqlDatabaseLister(indexer cache.Indexer) MssqlDatabaseLister {
	return &mssqlDatabaseLister{indexer: indexer}
}

// List lists all MssqlDatabases in the indexer.
func (s *mssqlDatabaseLister) List(selector labels.Selector) (ret []*v1alpha1.MssqlDatabase, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MssqlDatabase))
	})
	return ret, err
}

// MssqlDatabases returns an object that can list and get MssqlDatabases.
func (s *mssqlDatabaseLister) MssqlDatabases(namespace string) MssqlDatabaseNamespaceLister {
	return mssqlDatabaseNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MssqlDatabaseNamespaceLister helps list and get MssqlDatabases.
type MssqlDatabaseNamespaceLister interface {
	// List lists all MssqlDatabases in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.MssqlDatabase, err error)
	// Get retrieves the MssqlDatabase from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.MssqlDatabase, error)
	MssqlDatabaseNamespaceListerExpansion
}

// mssqlDatabaseNamespaceLister implements the MssqlDatabaseNamespaceLister
// interface.
type mssqlDatabaseNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MssqlDatabases in the indexer for a given namespace.
func (s mssqlDatabaseNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MssqlDatabase, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MssqlDatabase))
	})
	return ret, err
}

// Get retrieves the MssqlDatabase from the indexer for a given namespace and name.
func (s mssqlDatabaseNamespaceLister) Get(name string) (*v1alpha1.MssqlDatabase, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("mssqldatabase"), name)
	}
	return obj.(*v1alpha1.MssqlDatabase), nil
}
