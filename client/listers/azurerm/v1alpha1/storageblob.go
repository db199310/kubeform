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

// StorageBlobLister helps list StorageBlobs.
type StorageBlobLister interface {
	// List lists all StorageBlobs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.StorageBlob, err error)
	// StorageBlobs returns an object that can list and get StorageBlobs.
	StorageBlobs(namespace string) StorageBlobNamespaceLister
	StorageBlobListerExpansion
}

// storageBlobLister implements the StorageBlobLister interface.
type storageBlobLister struct {
	indexer cache.Indexer
}

// NewStorageBlobLister returns a new StorageBlobLister.
func NewStorageBlobLister(indexer cache.Indexer) StorageBlobLister {
	return &storageBlobLister{indexer: indexer}
}

// List lists all StorageBlobs in the indexer.
func (s *storageBlobLister) List(selector labels.Selector) (ret []*v1alpha1.StorageBlob, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StorageBlob))
	})
	return ret, err
}

// StorageBlobs returns an object that can list and get StorageBlobs.
func (s *storageBlobLister) StorageBlobs(namespace string) StorageBlobNamespaceLister {
	return storageBlobNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StorageBlobNamespaceLister helps list and get StorageBlobs.
type StorageBlobNamespaceLister interface {
	// List lists all StorageBlobs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.StorageBlob, err error)
	// Get retrieves the StorageBlob from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.StorageBlob, error)
	StorageBlobNamespaceListerExpansion
}

// storageBlobNamespaceLister implements the StorageBlobNamespaceLister
// interface.
type storageBlobNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StorageBlobs in the indexer for a given namespace.
func (s storageBlobNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.StorageBlob, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StorageBlob))
	})
	return ret, err
}

// Get retrieves the StorageBlob from the indexer for a given namespace and name.
func (s storageBlobNamespaceLister) Get(name string) (*v1alpha1.StorageBlob, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("storageblob"), name)
	}
	return obj.(*v1alpha1.StorageBlob), nil
}
