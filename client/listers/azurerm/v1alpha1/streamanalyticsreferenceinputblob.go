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

// StreamAnalyticsReferenceInputBlobLister helps list StreamAnalyticsReferenceInputBlobs.
type StreamAnalyticsReferenceInputBlobLister interface {
	// List lists all StreamAnalyticsReferenceInputBlobs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.StreamAnalyticsReferenceInputBlob, err error)
	// StreamAnalyticsReferenceInputBlobs returns an object that can list and get StreamAnalyticsReferenceInputBlobs.
	StreamAnalyticsReferenceInputBlobs(namespace string) StreamAnalyticsReferenceInputBlobNamespaceLister
	StreamAnalyticsReferenceInputBlobListerExpansion
}

// streamAnalyticsReferenceInputBlobLister implements the StreamAnalyticsReferenceInputBlobLister interface.
type streamAnalyticsReferenceInputBlobLister struct {
	indexer cache.Indexer
}

// NewStreamAnalyticsReferenceInputBlobLister returns a new StreamAnalyticsReferenceInputBlobLister.
func NewStreamAnalyticsReferenceInputBlobLister(indexer cache.Indexer) StreamAnalyticsReferenceInputBlobLister {
	return &streamAnalyticsReferenceInputBlobLister{indexer: indexer}
}

// List lists all StreamAnalyticsReferenceInputBlobs in the indexer.
func (s *streamAnalyticsReferenceInputBlobLister) List(selector labels.Selector) (ret []*v1alpha1.StreamAnalyticsReferenceInputBlob, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StreamAnalyticsReferenceInputBlob))
	})
	return ret, err
}

// StreamAnalyticsReferenceInputBlobs returns an object that can list and get StreamAnalyticsReferenceInputBlobs.
func (s *streamAnalyticsReferenceInputBlobLister) StreamAnalyticsReferenceInputBlobs(namespace string) StreamAnalyticsReferenceInputBlobNamespaceLister {
	return streamAnalyticsReferenceInputBlobNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StreamAnalyticsReferenceInputBlobNamespaceLister helps list and get StreamAnalyticsReferenceInputBlobs.
type StreamAnalyticsReferenceInputBlobNamespaceLister interface {
	// List lists all StreamAnalyticsReferenceInputBlobs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.StreamAnalyticsReferenceInputBlob, err error)
	// Get retrieves the StreamAnalyticsReferenceInputBlob from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.StreamAnalyticsReferenceInputBlob, error)
	StreamAnalyticsReferenceInputBlobNamespaceListerExpansion
}

// streamAnalyticsReferenceInputBlobNamespaceLister implements the StreamAnalyticsReferenceInputBlobNamespaceLister
// interface.
type streamAnalyticsReferenceInputBlobNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StreamAnalyticsReferenceInputBlobs in the indexer for a given namespace.
func (s streamAnalyticsReferenceInputBlobNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.StreamAnalyticsReferenceInputBlob, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StreamAnalyticsReferenceInputBlob))
	})
	return ret, err
}

// Get retrieves the StreamAnalyticsReferenceInputBlob from the indexer for a given namespace and name.
func (s streamAnalyticsReferenceInputBlobNamespaceLister) Get(name string) (*v1alpha1.StreamAnalyticsReferenceInputBlob, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("streamanalyticsreferenceinputblob"), name)
	}
	return obj.(*v1alpha1.StreamAnalyticsReferenceInputBlob), nil
}
