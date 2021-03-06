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

// StorageQueueLister helps list StorageQueues.
type StorageQueueLister interface {
	// List lists all StorageQueues in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.StorageQueue, err error)
	// StorageQueues returns an object that can list and get StorageQueues.
	StorageQueues(namespace string) StorageQueueNamespaceLister
	StorageQueueListerExpansion
}

// storageQueueLister implements the StorageQueueLister interface.
type storageQueueLister struct {
	indexer cache.Indexer
}

// NewStorageQueueLister returns a new StorageQueueLister.
func NewStorageQueueLister(indexer cache.Indexer) StorageQueueLister {
	return &storageQueueLister{indexer: indexer}
}

// List lists all StorageQueues in the indexer.
func (s *storageQueueLister) List(selector labels.Selector) (ret []*v1alpha1.StorageQueue, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StorageQueue))
	})
	return ret, err
}

// StorageQueues returns an object that can list and get StorageQueues.
func (s *storageQueueLister) StorageQueues(namespace string) StorageQueueNamespaceLister {
	return storageQueueNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StorageQueueNamespaceLister helps list and get StorageQueues.
type StorageQueueNamespaceLister interface {
	// List lists all StorageQueues in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.StorageQueue, err error)
	// Get retrieves the StorageQueue from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.StorageQueue, error)
	StorageQueueNamespaceListerExpansion
}

// storageQueueNamespaceLister implements the StorageQueueNamespaceLister
// interface.
type storageQueueNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StorageQueues in the indexer for a given namespace.
func (s storageQueueNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.StorageQueue, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StorageQueue))
	})
	return ret, err
}

// Get retrieves the StorageQueue from the indexer for a given namespace and name.
func (s storageQueueNamespaceLister) Get(name string) (*v1alpha1.StorageQueue, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("storagequeue"), name)
	}
	return obj.(*v1alpha1.StorageQueue), nil
}
