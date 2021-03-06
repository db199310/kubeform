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

// NetworkDdosProtectionPlanLister helps list NetworkDdosProtectionPlans.
type NetworkDdosProtectionPlanLister interface {
	// List lists all NetworkDdosProtectionPlans in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.NetworkDdosProtectionPlan, err error)
	// NetworkDdosProtectionPlans returns an object that can list and get NetworkDdosProtectionPlans.
	NetworkDdosProtectionPlans(namespace string) NetworkDdosProtectionPlanNamespaceLister
	NetworkDdosProtectionPlanListerExpansion
}

// networkDdosProtectionPlanLister implements the NetworkDdosProtectionPlanLister interface.
type networkDdosProtectionPlanLister struct {
	indexer cache.Indexer
}

// NewNetworkDdosProtectionPlanLister returns a new NetworkDdosProtectionPlanLister.
func NewNetworkDdosProtectionPlanLister(indexer cache.Indexer) NetworkDdosProtectionPlanLister {
	return &networkDdosProtectionPlanLister{indexer: indexer}
}

// List lists all NetworkDdosProtectionPlans in the indexer.
func (s *networkDdosProtectionPlanLister) List(selector labels.Selector) (ret []*v1alpha1.NetworkDdosProtectionPlan, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetworkDdosProtectionPlan))
	})
	return ret, err
}

// NetworkDdosProtectionPlans returns an object that can list and get NetworkDdosProtectionPlans.
func (s *networkDdosProtectionPlanLister) NetworkDdosProtectionPlans(namespace string) NetworkDdosProtectionPlanNamespaceLister {
	return networkDdosProtectionPlanNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NetworkDdosProtectionPlanNamespaceLister helps list and get NetworkDdosProtectionPlans.
type NetworkDdosProtectionPlanNamespaceLister interface {
	// List lists all NetworkDdosProtectionPlans in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.NetworkDdosProtectionPlan, err error)
	// Get retrieves the NetworkDdosProtectionPlan from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.NetworkDdosProtectionPlan, error)
	NetworkDdosProtectionPlanNamespaceListerExpansion
}

// networkDdosProtectionPlanNamespaceLister implements the NetworkDdosProtectionPlanNamespaceLister
// interface.
type networkDdosProtectionPlanNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NetworkDdosProtectionPlans in the indexer for a given namespace.
func (s networkDdosProtectionPlanNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.NetworkDdosProtectionPlan, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetworkDdosProtectionPlan))
	})
	return ret, err
}

// Get retrieves the NetworkDdosProtectionPlan from the indexer for a given namespace and name.
func (s networkDdosProtectionPlanNamespaceLister) Get(name string) (*v1alpha1.NetworkDdosProtectionPlan, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("networkddosprotectionplan"), name)
	}
	return obj.(*v1alpha1.NetworkDdosProtectionPlan), nil
}
