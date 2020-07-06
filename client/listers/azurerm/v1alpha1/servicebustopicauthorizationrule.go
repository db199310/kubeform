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

// ServicebusTopicAuthorizationRuleLister helps list ServicebusTopicAuthorizationRules.
type ServicebusTopicAuthorizationRuleLister interface {
	// List lists all ServicebusTopicAuthorizationRules in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ServicebusTopicAuthorizationRule, err error)
	// ServicebusTopicAuthorizationRules returns an object that can list and get ServicebusTopicAuthorizationRules.
	ServicebusTopicAuthorizationRules(namespace string) ServicebusTopicAuthorizationRuleNamespaceLister
	ServicebusTopicAuthorizationRuleListerExpansion
}

// servicebusTopicAuthorizationRuleLister implements the ServicebusTopicAuthorizationRuleLister interface.
type servicebusTopicAuthorizationRuleLister struct {
	indexer cache.Indexer
}

// NewServicebusTopicAuthorizationRuleLister returns a new ServicebusTopicAuthorizationRuleLister.
func NewServicebusTopicAuthorizationRuleLister(indexer cache.Indexer) ServicebusTopicAuthorizationRuleLister {
	return &servicebusTopicAuthorizationRuleLister{indexer: indexer}
}

// List lists all ServicebusTopicAuthorizationRules in the indexer.
func (s *servicebusTopicAuthorizationRuleLister) List(selector labels.Selector) (ret []*v1alpha1.ServicebusTopicAuthorizationRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServicebusTopicAuthorizationRule))
	})
	return ret, err
}

// ServicebusTopicAuthorizationRules returns an object that can list and get ServicebusTopicAuthorizationRules.
func (s *servicebusTopicAuthorizationRuleLister) ServicebusTopicAuthorizationRules(namespace string) ServicebusTopicAuthorizationRuleNamespaceLister {
	return servicebusTopicAuthorizationRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServicebusTopicAuthorizationRuleNamespaceLister helps list and get ServicebusTopicAuthorizationRules.
type ServicebusTopicAuthorizationRuleNamespaceLister interface {
	// List lists all ServicebusTopicAuthorizationRules in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ServicebusTopicAuthorizationRule, err error)
	// Get retrieves the ServicebusTopicAuthorizationRule from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ServicebusTopicAuthorizationRule, error)
	ServicebusTopicAuthorizationRuleNamespaceListerExpansion
}

// servicebusTopicAuthorizationRuleNamespaceLister implements the ServicebusTopicAuthorizationRuleNamespaceLister
// interface.
type servicebusTopicAuthorizationRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ServicebusTopicAuthorizationRules in the indexer for a given namespace.
func (s servicebusTopicAuthorizationRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ServicebusTopicAuthorizationRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServicebusTopicAuthorizationRule))
	})
	return ret, err
}

// Get retrieves the ServicebusTopicAuthorizationRule from the indexer for a given namespace and name.
func (s servicebusTopicAuthorizationRuleNamespaceLister) Get(name string) (*v1alpha1.ServicebusTopicAuthorizationRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("servicebustopicauthorizationrule"), name)
	}
	return obj.(*v1alpha1.ServicebusTopicAuthorizationRule), nil
}
