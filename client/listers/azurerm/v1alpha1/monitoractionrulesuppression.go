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

// MonitorActionRuleSuppressionLister helps list MonitorActionRuleSuppressions.
type MonitorActionRuleSuppressionLister interface {
	// List lists all MonitorActionRuleSuppressions in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.MonitorActionRuleSuppression, err error)
	// MonitorActionRuleSuppressions returns an object that can list and get MonitorActionRuleSuppressions.
	MonitorActionRuleSuppressions(namespace string) MonitorActionRuleSuppressionNamespaceLister
	MonitorActionRuleSuppressionListerExpansion
}

// monitorActionRuleSuppressionLister implements the MonitorActionRuleSuppressionLister interface.
type monitorActionRuleSuppressionLister struct {
	indexer cache.Indexer
}

// NewMonitorActionRuleSuppressionLister returns a new MonitorActionRuleSuppressionLister.
func NewMonitorActionRuleSuppressionLister(indexer cache.Indexer) MonitorActionRuleSuppressionLister {
	return &monitorActionRuleSuppressionLister{indexer: indexer}
}

// List lists all MonitorActionRuleSuppressions in the indexer.
func (s *monitorActionRuleSuppressionLister) List(selector labels.Selector) (ret []*v1alpha1.MonitorActionRuleSuppression, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MonitorActionRuleSuppression))
	})
	return ret, err
}

// MonitorActionRuleSuppressions returns an object that can list and get MonitorActionRuleSuppressions.
func (s *monitorActionRuleSuppressionLister) MonitorActionRuleSuppressions(namespace string) MonitorActionRuleSuppressionNamespaceLister {
	return monitorActionRuleSuppressionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MonitorActionRuleSuppressionNamespaceLister helps list and get MonitorActionRuleSuppressions.
type MonitorActionRuleSuppressionNamespaceLister interface {
	// List lists all MonitorActionRuleSuppressions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.MonitorActionRuleSuppression, err error)
	// Get retrieves the MonitorActionRuleSuppression from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.MonitorActionRuleSuppression, error)
	MonitorActionRuleSuppressionNamespaceListerExpansion
}

// monitorActionRuleSuppressionNamespaceLister implements the MonitorActionRuleSuppressionNamespaceLister
// interface.
type monitorActionRuleSuppressionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MonitorActionRuleSuppressions in the indexer for a given namespace.
func (s monitorActionRuleSuppressionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MonitorActionRuleSuppression, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MonitorActionRuleSuppression))
	})
	return ret, err
}

// Get retrieves the MonitorActionRuleSuppression from the indexer for a given namespace and name.
func (s monitorActionRuleSuppressionNamespaceLister) Get(name string) (*v1alpha1.MonitorActionRuleSuppression, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("monitoractionrulesuppression"), name)
	}
	return obj.(*v1alpha1.MonitorActionRuleSuppression), nil
}
