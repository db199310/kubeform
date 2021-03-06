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

// EventhubNamespace_Lister helps list EventhubNamespace_s.
type EventhubNamespace_Lister interface {
	// List lists all EventhubNamespace_s in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.EventhubNamespace_, err error)
	// EventhubNamespace_s returns an object that can list and get EventhubNamespace_s.
	EventhubNamespace_s(namespace string) EventhubNamespace_NamespaceLister
	EventhubNamespace_ListerExpansion
}

// eventhubNamespace_Lister implements the EventhubNamespace_Lister interface.
type eventhubNamespace_Lister struct {
	indexer cache.Indexer
}

// NewEventhubNamespace_Lister returns a new EventhubNamespace_Lister.
func NewEventhubNamespace_Lister(indexer cache.Indexer) EventhubNamespace_Lister {
	return &eventhubNamespace_Lister{indexer: indexer}
}

// List lists all EventhubNamespace_s in the indexer.
func (s *eventhubNamespace_Lister) List(selector labels.Selector) (ret []*v1alpha1.EventhubNamespace_, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EventhubNamespace_))
	})
	return ret, err
}

// EventhubNamespace_s returns an object that can list and get EventhubNamespace_s.
func (s *eventhubNamespace_Lister) EventhubNamespace_s(namespace string) EventhubNamespace_NamespaceLister {
	return eventhubNamespace_NamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EventhubNamespace_NamespaceLister helps list and get EventhubNamespace_s.
type EventhubNamespace_NamespaceLister interface {
	// List lists all EventhubNamespace_s in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.EventhubNamespace_, err error)
	// Get retrieves the EventhubNamespace_ from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.EventhubNamespace_, error)
	EventhubNamespace_NamespaceListerExpansion
}

// eventhubNamespace_NamespaceLister implements the EventhubNamespace_NamespaceLister
// interface.
type eventhubNamespace_NamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EventhubNamespace_s in the indexer for a given namespace.
func (s eventhubNamespace_NamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EventhubNamespace_, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EventhubNamespace_))
	})
	return ret, err
}

// Get retrieves the EventhubNamespace_ from the indexer for a given namespace and name.
func (s eventhubNamespace_NamespaceLister) Get(name string) (*v1alpha1.EventhubNamespace_, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("eventhubnamespace_"), name)
	}
	return obj.(*v1alpha1.EventhubNamespace_), nil
}
