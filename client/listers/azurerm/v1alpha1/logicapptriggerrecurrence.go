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

// LogicAppTriggerRecurrenceLister helps list LogicAppTriggerRecurrences.
type LogicAppTriggerRecurrenceLister interface {
	// List lists all LogicAppTriggerRecurrences in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.LogicAppTriggerRecurrence, err error)
	// LogicAppTriggerRecurrences returns an object that can list and get LogicAppTriggerRecurrences.
	LogicAppTriggerRecurrences(namespace string) LogicAppTriggerRecurrenceNamespaceLister
	LogicAppTriggerRecurrenceListerExpansion
}

// logicAppTriggerRecurrenceLister implements the LogicAppTriggerRecurrenceLister interface.
type logicAppTriggerRecurrenceLister struct {
	indexer cache.Indexer
}

// NewLogicAppTriggerRecurrenceLister returns a new LogicAppTriggerRecurrenceLister.
func NewLogicAppTriggerRecurrenceLister(indexer cache.Indexer) LogicAppTriggerRecurrenceLister {
	return &logicAppTriggerRecurrenceLister{indexer: indexer}
}

// List lists all LogicAppTriggerRecurrences in the indexer.
func (s *logicAppTriggerRecurrenceLister) List(selector labels.Selector) (ret []*v1alpha1.LogicAppTriggerRecurrence, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LogicAppTriggerRecurrence))
	})
	return ret, err
}

// LogicAppTriggerRecurrences returns an object that can list and get LogicAppTriggerRecurrences.
func (s *logicAppTriggerRecurrenceLister) LogicAppTriggerRecurrences(namespace string) LogicAppTriggerRecurrenceNamespaceLister {
	return logicAppTriggerRecurrenceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LogicAppTriggerRecurrenceNamespaceLister helps list and get LogicAppTriggerRecurrences.
type LogicAppTriggerRecurrenceNamespaceLister interface {
	// List lists all LogicAppTriggerRecurrences in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.LogicAppTriggerRecurrence, err error)
	// Get retrieves the LogicAppTriggerRecurrence from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.LogicAppTriggerRecurrence, error)
	LogicAppTriggerRecurrenceNamespaceListerExpansion
}

// logicAppTriggerRecurrenceNamespaceLister implements the LogicAppTriggerRecurrenceNamespaceLister
// interface.
type logicAppTriggerRecurrenceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LogicAppTriggerRecurrences in the indexer for a given namespace.
func (s logicAppTriggerRecurrenceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LogicAppTriggerRecurrence, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LogicAppTriggerRecurrence))
	})
	return ret, err
}

// Get retrieves the LogicAppTriggerRecurrence from the indexer for a given namespace and name.
func (s logicAppTriggerRecurrenceNamespaceLister) Get(name string) (*v1alpha1.LogicAppTriggerRecurrence, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("logicapptriggerrecurrence"), name)
	}
	return obj.(*v1alpha1.LogicAppTriggerRecurrence), nil
}
