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

// SentinelAlertRuleScheduledLister helps list SentinelAlertRuleScheduleds.
type SentinelAlertRuleScheduledLister interface {
	// List lists all SentinelAlertRuleScheduleds in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SentinelAlertRuleScheduled, err error)
	// SentinelAlertRuleScheduleds returns an object that can list and get SentinelAlertRuleScheduleds.
	SentinelAlertRuleScheduleds(namespace string) SentinelAlertRuleScheduledNamespaceLister
	SentinelAlertRuleScheduledListerExpansion
}

// sentinelAlertRuleScheduledLister implements the SentinelAlertRuleScheduledLister interface.
type sentinelAlertRuleScheduledLister struct {
	indexer cache.Indexer
}

// NewSentinelAlertRuleScheduledLister returns a new SentinelAlertRuleScheduledLister.
func NewSentinelAlertRuleScheduledLister(indexer cache.Indexer) SentinelAlertRuleScheduledLister {
	return &sentinelAlertRuleScheduledLister{indexer: indexer}
}

// List lists all SentinelAlertRuleScheduleds in the indexer.
func (s *sentinelAlertRuleScheduledLister) List(selector labels.Selector) (ret []*v1alpha1.SentinelAlertRuleScheduled, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SentinelAlertRuleScheduled))
	})
	return ret, err
}

// SentinelAlertRuleScheduleds returns an object that can list and get SentinelAlertRuleScheduleds.
func (s *sentinelAlertRuleScheduledLister) SentinelAlertRuleScheduleds(namespace string) SentinelAlertRuleScheduledNamespaceLister {
	return sentinelAlertRuleScheduledNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SentinelAlertRuleScheduledNamespaceLister helps list and get SentinelAlertRuleScheduleds.
type SentinelAlertRuleScheduledNamespaceLister interface {
	// List lists all SentinelAlertRuleScheduleds in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SentinelAlertRuleScheduled, err error)
	// Get retrieves the SentinelAlertRuleScheduled from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SentinelAlertRuleScheduled, error)
	SentinelAlertRuleScheduledNamespaceListerExpansion
}

// sentinelAlertRuleScheduledNamespaceLister implements the SentinelAlertRuleScheduledNamespaceLister
// interface.
type sentinelAlertRuleScheduledNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SentinelAlertRuleScheduleds in the indexer for a given namespace.
func (s sentinelAlertRuleScheduledNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SentinelAlertRuleScheduled, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SentinelAlertRuleScheduled))
	})
	return ret, err
}

// Get retrieves the SentinelAlertRuleScheduled from the indexer for a given namespace and name.
func (s sentinelAlertRuleScheduledNamespaceLister) Get(name string) (*v1alpha1.SentinelAlertRuleScheduled, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("sentinelalertrulescheduled"), name)
	}
	return obj.(*v1alpha1.SentinelAlertRuleScheduled), nil
}
