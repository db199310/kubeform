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

// LbNATRuleLister helps list LbNATRules.
type LbNATRuleLister interface {
	// List lists all LbNATRules in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.LbNATRule, err error)
	// LbNATRules returns an object that can list and get LbNATRules.
	LbNATRules(namespace string) LbNATRuleNamespaceLister
	LbNATRuleListerExpansion
}

// lbNATRuleLister implements the LbNATRuleLister interface.
type lbNATRuleLister struct {
	indexer cache.Indexer
}

// NewLbNATRuleLister returns a new LbNATRuleLister.
func NewLbNATRuleLister(indexer cache.Indexer) LbNATRuleLister {
	return &lbNATRuleLister{indexer: indexer}
}

// List lists all LbNATRules in the indexer.
func (s *lbNATRuleLister) List(selector labels.Selector) (ret []*v1alpha1.LbNATRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LbNATRule))
	})
	return ret, err
}

// LbNATRules returns an object that can list and get LbNATRules.
func (s *lbNATRuleLister) LbNATRules(namespace string) LbNATRuleNamespaceLister {
	return lbNATRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LbNATRuleNamespaceLister helps list and get LbNATRules.
type LbNATRuleNamespaceLister interface {
	// List lists all LbNATRules in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.LbNATRule, err error)
	// Get retrieves the LbNATRule from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.LbNATRule, error)
	LbNATRuleNamespaceListerExpansion
}

// lbNATRuleNamespaceLister implements the LbNATRuleNamespaceLister
// interface.
type lbNATRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LbNATRules in the indexer for a given namespace.
func (s lbNATRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LbNATRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LbNATRule))
	})
	return ret, err
}

// Get retrieves the LbNATRule from the indexer for a given namespace and name.
func (s lbNATRuleNamespaceLister) Get(name string) (*v1alpha1.LbNATRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("lbnatrule"), name)
	}
	return obj.(*v1alpha1.LbNATRule), nil
}
