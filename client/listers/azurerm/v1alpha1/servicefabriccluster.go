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

// ServiceFabricClusterLister helps list ServiceFabricClusters.
type ServiceFabricClusterLister interface {
	// List lists all ServiceFabricClusters in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ServiceFabricCluster, err error)
	// ServiceFabricClusters returns an object that can list and get ServiceFabricClusters.
	ServiceFabricClusters(namespace string) ServiceFabricClusterNamespaceLister
	ServiceFabricClusterListerExpansion
}

// serviceFabricClusterLister implements the ServiceFabricClusterLister interface.
type serviceFabricClusterLister struct {
	indexer cache.Indexer
}

// NewServiceFabricClusterLister returns a new ServiceFabricClusterLister.
func NewServiceFabricClusterLister(indexer cache.Indexer) ServiceFabricClusterLister {
	return &serviceFabricClusterLister{indexer: indexer}
}

// List lists all ServiceFabricClusters in the indexer.
func (s *serviceFabricClusterLister) List(selector labels.Selector) (ret []*v1alpha1.ServiceFabricCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServiceFabricCluster))
	})
	return ret, err
}

// ServiceFabricClusters returns an object that can list and get ServiceFabricClusters.
func (s *serviceFabricClusterLister) ServiceFabricClusters(namespace string) ServiceFabricClusterNamespaceLister {
	return serviceFabricClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServiceFabricClusterNamespaceLister helps list and get ServiceFabricClusters.
type ServiceFabricClusterNamespaceLister interface {
	// List lists all ServiceFabricClusters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ServiceFabricCluster, err error)
	// Get retrieves the ServiceFabricCluster from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ServiceFabricCluster, error)
	ServiceFabricClusterNamespaceListerExpansion
}

// serviceFabricClusterNamespaceLister implements the ServiceFabricClusterNamespaceLister
// interface.
type serviceFabricClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ServiceFabricClusters in the indexer for a given namespace.
func (s serviceFabricClusterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ServiceFabricCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServiceFabricCluster))
	})
	return ret, err
}

// Get retrieves the ServiceFabricCluster from the indexer for a given namespace and name.
func (s serviceFabricClusterNamespaceLister) Get(name string) (*v1alpha1.ServiceFabricCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("servicefabriccluster"), name)
	}
	return obj.(*v1alpha1.ServiceFabricCluster), nil
}
