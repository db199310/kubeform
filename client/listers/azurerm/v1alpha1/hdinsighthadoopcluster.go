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

// HdinsightHadoopClusterLister helps list HdinsightHadoopClusters.
type HdinsightHadoopClusterLister interface {
	// List lists all HdinsightHadoopClusters in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.HdinsightHadoopCluster, err error)
	// HdinsightHadoopClusters returns an object that can list and get HdinsightHadoopClusters.
	HdinsightHadoopClusters(namespace string) HdinsightHadoopClusterNamespaceLister
	HdinsightHadoopClusterListerExpansion
}

// hdinsightHadoopClusterLister implements the HdinsightHadoopClusterLister interface.
type hdinsightHadoopClusterLister struct {
	indexer cache.Indexer
}

// NewHdinsightHadoopClusterLister returns a new HdinsightHadoopClusterLister.
func NewHdinsightHadoopClusterLister(indexer cache.Indexer) HdinsightHadoopClusterLister {
	return &hdinsightHadoopClusterLister{indexer: indexer}
}

// List lists all HdinsightHadoopClusters in the indexer.
func (s *hdinsightHadoopClusterLister) List(selector labels.Selector) (ret []*v1alpha1.HdinsightHadoopCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HdinsightHadoopCluster))
	})
	return ret, err
}

// HdinsightHadoopClusters returns an object that can list and get HdinsightHadoopClusters.
func (s *hdinsightHadoopClusterLister) HdinsightHadoopClusters(namespace string) HdinsightHadoopClusterNamespaceLister {
	return hdinsightHadoopClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HdinsightHadoopClusterNamespaceLister helps list and get HdinsightHadoopClusters.
type HdinsightHadoopClusterNamespaceLister interface {
	// List lists all HdinsightHadoopClusters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.HdinsightHadoopCluster, err error)
	// Get retrieves the HdinsightHadoopCluster from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.HdinsightHadoopCluster, error)
	HdinsightHadoopClusterNamespaceListerExpansion
}

// hdinsightHadoopClusterNamespaceLister implements the HdinsightHadoopClusterNamespaceLister
// interface.
type hdinsightHadoopClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all HdinsightHadoopClusters in the indexer for a given namespace.
func (s hdinsightHadoopClusterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.HdinsightHadoopCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HdinsightHadoopCluster))
	})
	return ret, err
}

// Get retrieves the HdinsightHadoopCluster from the indexer for a given namespace and name.
func (s hdinsightHadoopClusterNamespaceLister) Get(name string) (*v1alpha1.HdinsightHadoopCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("hdinsighthadoopcluster"), name)
	}
	return obj.(*v1alpha1.HdinsightHadoopCluster), nil
}
