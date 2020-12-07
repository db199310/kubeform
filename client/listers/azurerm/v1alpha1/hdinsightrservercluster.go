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

// HdinsightRserverClusterLister helps list HdinsightRserverClusters.
type HdinsightRserverClusterLister interface {
	// List lists all HdinsightRserverClusters in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.HdinsightRserverCluster, err error)
	// HdinsightRserverClusters returns an object that can list and get HdinsightRserverClusters.
	HdinsightRserverClusters(namespace string) HdinsightRserverClusterNamespaceLister
	HdinsightRserverClusterListerExpansion
}

// hdinsightRserverClusterLister implements the HdinsightRserverClusterLister interface.
type hdinsightRserverClusterLister struct {
	indexer cache.Indexer
}

// NewHdinsightRserverClusterLister returns a new HdinsightRserverClusterLister.
func NewHdinsightRserverClusterLister(indexer cache.Indexer) HdinsightRserverClusterLister {
	return &hdinsightRserverClusterLister{indexer: indexer}
}

// List lists all HdinsightRserverClusters in the indexer.
func (s *hdinsightRserverClusterLister) List(selector labels.Selector) (ret []*v1alpha1.HdinsightRserverCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HdinsightRserverCluster))
	})
	return ret, err
}

// HdinsightRserverClusters returns an object that can list and get HdinsightRserverClusters.
func (s *hdinsightRserverClusterLister) HdinsightRserverClusters(namespace string) HdinsightRserverClusterNamespaceLister {
	return hdinsightRserverClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HdinsightRserverClusterNamespaceLister helps list and get HdinsightRserverClusters.
type HdinsightRserverClusterNamespaceLister interface {
	// List lists all HdinsightRserverClusters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.HdinsightRserverCluster, err error)
	// Get retrieves the HdinsightRserverCluster from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.HdinsightRserverCluster, error)
	HdinsightRserverClusterNamespaceListerExpansion
}

// hdinsightRserverClusterNamespaceLister implements the HdinsightRserverClusterNamespaceLister
// interface.
type hdinsightRserverClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all HdinsightRserverClusters in the indexer for a given namespace.
func (s hdinsightRserverClusterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.HdinsightRserverCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HdinsightRserverCluster))
	})
	return ret, err
}

// Get retrieves the HdinsightRserverCluster from the indexer for a given namespace and name.
func (s hdinsightRserverClusterNamespaceLister) Get(name string) (*v1alpha1.HdinsightRserverCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("hdinsightrservercluster"), name)
	}
	return obj.(*v1alpha1.HdinsightRserverCluster), nil
}
