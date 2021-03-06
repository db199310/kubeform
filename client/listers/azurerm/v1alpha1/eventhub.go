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

// EventhubLister helps list Eventhubs.
type EventhubLister interface {
	// List lists all Eventhubs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Eventhub, err error)
	// Eventhubs returns an object that can list and get Eventhubs.
	Eventhubs(namespace string) EventhubNamespaceLister
	EventhubListerExpansion
}

// eventhubLister implements the EventhubLister interface.
type eventhubLister struct {
	indexer cache.Indexer
}

// NewEventhubLister returns a new EventhubLister.
func NewEventhubLister(indexer cache.Indexer) EventhubLister {
	return &eventhubLister{indexer: indexer}
}

// List lists all Eventhubs in the indexer.
func (s *eventhubLister) List(selector labels.Selector) (ret []*v1alpha1.Eventhub, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Eventhub))
	})
	return ret, err
}

// Eventhubs returns an object that can list and get Eventhubs.
func (s *eventhubLister) Eventhubs(namespace string) EventhubNamespaceLister {
	return eventhubNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EventhubNamespaceLister helps list and get Eventhubs.
type EventhubNamespaceLister interface {
	// List lists all Eventhubs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Eventhub, err error)
	// Get retrieves the Eventhub from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Eventhub, error)
	EventhubNamespaceListerExpansion
}

// eventhubNamespaceLister implements the EventhubNamespaceLister
// interface.
type eventhubNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Eventhubs in the indexer for a given namespace.
func (s eventhubNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Eventhub, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Eventhub))
	})
	return ret, err
}

// Get retrieves the Eventhub from the indexer for a given namespace and name.
func (s eventhubNamespaceLister) Get(name string) (*v1alpha1.Eventhub, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("eventhub"), name)
	}
	return obj.(*v1alpha1.Eventhub), nil
}
