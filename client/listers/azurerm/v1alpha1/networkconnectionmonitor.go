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

// NetworkConnectionMonitorLister helps list NetworkConnectionMonitors.
type NetworkConnectionMonitorLister interface {
	// List lists all NetworkConnectionMonitors in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.NetworkConnectionMonitor, err error)
	// NetworkConnectionMonitors returns an object that can list and get NetworkConnectionMonitors.
	NetworkConnectionMonitors(namespace string) NetworkConnectionMonitorNamespaceLister
	NetworkConnectionMonitorListerExpansion
}

// networkConnectionMonitorLister implements the NetworkConnectionMonitorLister interface.
type networkConnectionMonitorLister struct {
	indexer cache.Indexer
}

// NewNetworkConnectionMonitorLister returns a new NetworkConnectionMonitorLister.
func NewNetworkConnectionMonitorLister(indexer cache.Indexer) NetworkConnectionMonitorLister {
	return &networkConnectionMonitorLister{indexer: indexer}
}

// List lists all NetworkConnectionMonitors in the indexer.
func (s *networkConnectionMonitorLister) List(selector labels.Selector) (ret []*v1alpha1.NetworkConnectionMonitor, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetworkConnectionMonitor))
	})
	return ret, err
}

// NetworkConnectionMonitors returns an object that can list and get NetworkConnectionMonitors.
func (s *networkConnectionMonitorLister) NetworkConnectionMonitors(namespace string) NetworkConnectionMonitorNamespaceLister {
	return networkConnectionMonitorNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NetworkConnectionMonitorNamespaceLister helps list and get NetworkConnectionMonitors.
type NetworkConnectionMonitorNamespaceLister interface {
	// List lists all NetworkConnectionMonitors in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.NetworkConnectionMonitor, err error)
	// Get retrieves the NetworkConnectionMonitor from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.NetworkConnectionMonitor, error)
	NetworkConnectionMonitorNamespaceListerExpansion
}

// networkConnectionMonitorNamespaceLister implements the NetworkConnectionMonitorNamespaceLister
// interface.
type networkConnectionMonitorNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NetworkConnectionMonitors in the indexer for a given namespace.
func (s networkConnectionMonitorNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.NetworkConnectionMonitor, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetworkConnectionMonitor))
	})
	return ret, err
}

// Get retrieves the NetworkConnectionMonitor from the indexer for a given namespace and name.
func (s networkConnectionMonitorNamespaceLister) Get(name string) (*v1alpha1.NetworkConnectionMonitor, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("networkconnectionmonitor"), name)
	}
	return obj.(*v1alpha1.NetworkConnectionMonitor), nil
}
