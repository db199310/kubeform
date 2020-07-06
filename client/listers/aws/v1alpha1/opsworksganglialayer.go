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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// OpsworksGangliaLayerLister helps list OpsworksGangliaLayers.
type OpsworksGangliaLayerLister interface {
	// List lists all OpsworksGangliaLayers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.OpsworksGangliaLayer, err error)
	// OpsworksGangliaLayers returns an object that can list and get OpsworksGangliaLayers.
	OpsworksGangliaLayers(namespace string) OpsworksGangliaLayerNamespaceLister
	OpsworksGangliaLayerListerExpansion
}

// opsworksGangliaLayerLister implements the OpsworksGangliaLayerLister interface.
type opsworksGangliaLayerLister struct {
	indexer cache.Indexer
}

// NewOpsworksGangliaLayerLister returns a new OpsworksGangliaLayerLister.
func NewOpsworksGangliaLayerLister(indexer cache.Indexer) OpsworksGangliaLayerLister {
	return &opsworksGangliaLayerLister{indexer: indexer}
}

// List lists all OpsworksGangliaLayers in the indexer.
func (s *opsworksGangliaLayerLister) List(selector labels.Selector) (ret []*v1alpha1.OpsworksGangliaLayer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OpsworksGangliaLayer))
	})
	return ret, err
}

// OpsworksGangliaLayers returns an object that can list and get OpsworksGangliaLayers.
func (s *opsworksGangliaLayerLister) OpsworksGangliaLayers(namespace string) OpsworksGangliaLayerNamespaceLister {
	return opsworksGangliaLayerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// OpsworksGangliaLayerNamespaceLister helps list and get OpsworksGangliaLayers.
type OpsworksGangliaLayerNamespaceLister interface {
	// List lists all OpsworksGangliaLayers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.OpsworksGangliaLayer, err error)
	// Get retrieves the OpsworksGangliaLayer from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.OpsworksGangliaLayer, error)
	OpsworksGangliaLayerNamespaceListerExpansion
}

// opsworksGangliaLayerNamespaceLister implements the OpsworksGangliaLayerNamespaceLister
// interface.
type opsworksGangliaLayerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all OpsworksGangliaLayers in the indexer for a given namespace.
func (s opsworksGangliaLayerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.OpsworksGangliaLayer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OpsworksGangliaLayer))
	})
	return ret, err
}

// Get retrieves the OpsworksGangliaLayer from the indexer for a given namespace and name.
func (s opsworksGangliaLayerNamespaceLister) Get(name string) (*v1alpha1.OpsworksGangliaLayer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("opsworksganglialayer"), name)
	}
	return obj.(*v1alpha1.OpsworksGangliaLayer), nil
}
