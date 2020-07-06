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

// SqlServerLister helps list SqlServers.
type SqlServerLister interface {
	// List lists all SqlServers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SqlServer, err error)
	// SqlServers returns an object that can list and get SqlServers.
	SqlServers(namespace string) SqlServerNamespaceLister
	SqlServerListerExpansion
}

// sqlServerLister implements the SqlServerLister interface.
type sqlServerLister struct {
	indexer cache.Indexer
}

// NewSqlServerLister returns a new SqlServerLister.
func NewSqlServerLister(indexer cache.Indexer) SqlServerLister {
	return &sqlServerLister{indexer: indexer}
}

// List lists all SqlServers in the indexer.
func (s *sqlServerLister) List(selector labels.Selector) (ret []*v1alpha1.SqlServer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SqlServer))
	})
	return ret, err
}

// SqlServers returns an object that can list and get SqlServers.
func (s *sqlServerLister) SqlServers(namespace string) SqlServerNamespaceLister {
	return sqlServerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SqlServerNamespaceLister helps list and get SqlServers.
type SqlServerNamespaceLister interface {
	// List lists all SqlServers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SqlServer, err error)
	// Get retrieves the SqlServer from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SqlServer, error)
	SqlServerNamespaceListerExpansion
}

// sqlServerNamespaceLister implements the SqlServerNamespaceLister
// interface.
type sqlServerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SqlServers in the indexer for a given namespace.
func (s sqlServerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SqlServer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SqlServer))
	})
	return ret, err
}

// Get retrieves the SqlServer from the indexer for a given namespace and name.
func (s sqlServerNamespaceLister) Get(name string) (*v1alpha1.SqlServer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("sqlserver"), name)
	}
	return obj.(*v1alpha1.SqlServer), nil
}
