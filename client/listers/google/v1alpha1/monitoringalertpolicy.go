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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// MonitoringAlertPolicyLister helps list MonitoringAlertPolicies.
type MonitoringAlertPolicyLister interface {
	// List lists all MonitoringAlertPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.MonitoringAlertPolicy, err error)
	// MonitoringAlertPolicies returns an object that can list and get MonitoringAlertPolicies.
	MonitoringAlertPolicies(namespace string) MonitoringAlertPolicyNamespaceLister
	MonitoringAlertPolicyListerExpansion
}

// monitoringAlertPolicyLister implements the MonitoringAlertPolicyLister interface.
type monitoringAlertPolicyLister struct {
	indexer cache.Indexer
}

// NewMonitoringAlertPolicyLister returns a new MonitoringAlertPolicyLister.
func NewMonitoringAlertPolicyLister(indexer cache.Indexer) MonitoringAlertPolicyLister {
	return &monitoringAlertPolicyLister{indexer: indexer}
}

// List lists all MonitoringAlertPolicies in the indexer.
func (s *monitoringAlertPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.MonitoringAlertPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MonitoringAlertPolicy))
	})
	return ret, err
}

// MonitoringAlertPolicies returns an object that can list and get MonitoringAlertPolicies.
func (s *monitoringAlertPolicyLister) MonitoringAlertPolicies(namespace string) MonitoringAlertPolicyNamespaceLister {
	return monitoringAlertPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MonitoringAlertPolicyNamespaceLister helps list and get MonitoringAlertPolicies.
type MonitoringAlertPolicyNamespaceLister interface {
	// List lists all MonitoringAlertPolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.MonitoringAlertPolicy, err error)
	// Get retrieves the MonitoringAlertPolicy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.MonitoringAlertPolicy, error)
	MonitoringAlertPolicyNamespaceListerExpansion
}

// monitoringAlertPolicyNamespaceLister implements the MonitoringAlertPolicyNamespaceLister
// interface.
type monitoringAlertPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MonitoringAlertPolicies in the indexer for a given namespace.
func (s monitoringAlertPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MonitoringAlertPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MonitoringAlertPolicy))
	})
	return ret, err
}

// Get retrieves the MonitoringAlertPolicy from the indexer for a given namespace and name.
func (s monitoringAlertPolicyNamespaceLister) Get(name string) (*v1alpha1.MonitoringAlertPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("monitoringalertpolicy"), name)
	}
	return obj.(*v1alpha1.MonitoringAlertPolicy), nil
}
