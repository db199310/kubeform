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

// ServicebusQueueAuthorizationRuleLister helps list ServicebusQueueAuthorizationRules.
type ServicebusQueueAuthorizationRuleLister interface {
	// List lists all ServicebusQueueAuthorizationRules in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ServicebusQueueAuthorizationRule, err error)
	// ServicebusQueueAuthorizationRules returns an object that can list and get ServicebusQueueAuthorizationRules.
	ServicebusQueueAuthorizationRules(namespace string) ServicebusQueueAuthorizationRuleNamespaceLister
	ServicebusQueueAuthorizationRuleListerExpansion
}

// servicebusQueueAuthorizationRuleLister implements the ServicebusQueueAuthorizationRuleLister interface.
type servicebusQueueAuthorizationRuleLister struct {
	indexer cache.Indexer
}

// NewServicebusQueueAuthorizationRuleLister returns a new ServicebusQueueAuthorizationRuleLister.
func NewServicebusQueueAuthorizationRuleLister(indexer cache.Indexer) ServicebusQueueAuthorizationRuleLister {
	return &servicebusQueueAuthorizationRuleLister{indexer: indexer}
}

// List lists all ServicebusQueueAuthorizationRules in the indexer.
func (s *servicebusQueueAuthorizationRuleLister) List(selector labels.Selector) (ret []*v1alpha1.ServicebusQueueAuthorizationRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServicebusQueueAuthorizationRule))
	})
	return ret, err
}

// ServicebusQueueAuthorizationRules returns an object that can list and get ServicebusQueueAuthorizationRules.
func (s *servicebusQueueAuthorizationRuleLister) ServicebusQueueAuthorizationRules(namespace string) ServicebusQueueAuthorizationRuleNamespaceLister {
	return servicebusQueueAuthorizationRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServicebusQueueAuthorizationRuleNamespaceLister helps list and get ServicebusQueueAuthorizationRules.
type ServicebusQueueAuthorizationRuleNamespaceLister interface {
	// List lists all ServicebusQueueAuthorizationRules in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ServicebusQueueAuthorizationRule, err error)
	// Get retrieves the ServicebusQueueAuthorizationRule from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ServicebusQueueAuthorizationRule, error)
	ServicebusQueueAuthorizationRuleNamespaceListerExpansion
}

// servicebusQueueAuthorizationRuleNamespaceLister implements the ServicebusQueueAuthorizationRuleNamespaceLister
// interface.
type servicebusQueueAuthorizationRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ServicebusQueueAuthorizationRules in the indexer for a given namespace.
func (s servicebusQueueAuthorizationRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ServicebusQueueAuthorizationRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServicebusQueueAuthorizationRule))
	})
	return ret, err
}

// Get retrieves the ServicebusQueueAuthorizationRule from the indexer for a given namespace and name.
func (s servicebusQueueAuthorizationRuleNamespaceLister) Get(name string) (*v1alpha1.ServicebusQueueAuthorizationRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("servicebusqueueauthorizationrule"), name)
	}
	return obj.(*v1alpha1.ServicebusQueueAuthorizationRule), nil
}
