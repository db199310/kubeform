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

// CosmosdbMongoCollectionLister helps list CosmosdbMongoCollections.
type CosmosdbMongoCollectionLister interface {
	// List lists all CosmosdbMongoCollections in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.CosmosdbMongoCollection, err error)
	// CosmosdbMongoCollections returns an object that can list and get CosmosdbMongoCollections.
	CosmosdbMongoCollections(namespace string) CosmosdbMongoCollectionNamespaceLister
	CosmosdbMongoCollectionListerExpansion
}

// cosmosdbMongoCollectionLister implements the CosmosdbMongoCollectionLister interface.
type cosmosdbMongoCollectionLister struct {
	indexer cache.Indexer
}

// NewCosmosdbMongoCollectionLister returns a new CosmosdbMongoCollectionLister.
func NewCosmosdbMongoCollectionLister(indexer cache.Indexer) CosmosdbMongoCollectionLister {
	return &cosmosdbMongoCollectionLister{indexer: indexer}
}

// List lists all CosmosdbMongoCollections in the indexer.
func (s *cosmosdbMongoCollectionLister) List(selector labels.Selector) (ret []*v1alpha1.CosmosdbMongoCollection, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CosmosdbMongoCollection))
	})
	return ret, err
}

// CosmosdbMongoCollections returns an object that can list and get CosmosdbMongoCollections.
func (s *cosmosdbMongoCollectionLister) CosmosdbMongoCollections(namespace string) CosmosdbMongoCollectionNamespaceLister {
	return cosmosdbMongoCollectionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CosmosdbMongoCollectionNamespaceLister helps list and get CosmosdbMongoCollections.
type CosmosdbMongoCollectionNamespaceLister interface {
	// List lists all CosmosdbMongoCollections in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.CosmosdbMongoCollection, err error)
	// Get retrieves the CosmosdbMongoCollection from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.CosmosdbMongoCollection, error)
	CosmosdbMongoCollectionNamespaceListerExpansion
}

// cosmosdbMongoCollectionNamespaceLister implements the CosmosdbMongoCollectionNamespaceLister
// interface.
type cosmosdbMongoCollectionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CosmosdbMongoCollections in the indexer for a given namespace.
func (s cosmosdbMongoCollectionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CosmosdbMongoCollection, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CosmosdbMongoCollection))
	})
	return ret, err
}

// Get retrieves the CosmosdbMongoCollection from the indexer for a given namespace and name.
func (s cosmosdbMongoCollectionNamespaceLister) Get(name string) (*v1alpha1.CosmosdbMongoCollection, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cosmosdbmongocollection"), name)
	}
	return obj.(*v1alpha1.CosmosdbMongoCollection), nil
}
