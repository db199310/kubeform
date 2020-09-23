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

// VirtualHubConnectionLister helps list VirtualHubConnections.
type VirtualHubConnectionLister interface {
	// List lists all VirtualHubConnections in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.VirtualHubConnection, err error)
	// VirtualHubConnections returns an object that can list and get VirtualHubConnections.
	VirtualHubConnections(namespace string) VirtualHubConnectionNamespaceLister
	VirtualHubConnectionListerExpansion
}

// virtualHubConnectionLister implements the VirtualHubConnectionLister interface.
type virtualHubConnectionLister struct {
	indexer cache.Indexer
}

// NewVirtualHubConnectionLister returns a new VirtualHubConnectionLister.
func NewVirtualHubConnectionLister(indexer cache.Indexer) VirtualHubConnectionLister {
	return &virtualHubConnectionLister{indexer: indexer}
}

// List lists all VirtualHubConnections in the indexer.
func (s *virtualHubConnectionLister) List(selector labels.Selector) (ret []*v1alpha1.VirtualHubConnection, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VirtualHubConnection))
	})
	return ret, err
}

// VirtualHubConnections returns an object that can list and get VirtualHubConnections.
func (s *virtualHubConnectionLister) VirtualHubConnections(namespace string) VirtualHubConnectionNamespaceLister {
	return virtualHubConnectionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VirtualHubConnectionNamespaceLister helps list and get VirtualHubConnections.
type VirtualHubConnectionNamespaceLister interface {
	// List lists all VirtualHubConnections in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.VirtualHubConnection, err error)
	// Get retrieves the VirtualHubConnection from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.VirtualHubConnection, error)
	VirtualHubConnectionNamespaceListerExpansion
}

// virtualHubConnectionNamespaceLister implements the VirtualHubConnectionNamespaceLister
// interface.
type virtualHubConnectionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VirtualHubConnections in the indexer for a given namespace.
func (s virtualHubConnectionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.VirtualHubConnection, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VirtualHubConnection))
	})
	return ret, err
}

// Get retrieves the VirtualHubConnection from the indexer for a given namespace and name.
func (s virtualHubConnectionNamespaceLister) Get(name string) (*v1alpha1.VirtualHubConnection, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("virtualhubconnection"), name)
	}
	return obj.(*v1alpha1.VirtualHubConnection), nil
}