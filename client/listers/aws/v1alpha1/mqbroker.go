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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// MqBrokerLister helps list MqBrokers.
type MqBrokerLister interface {
	// List lists all MqBrokers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.MqBroker, err error)
	// MqBrokers returns an object that can list and get MqBrokers.
	MqBrokers(namespace string) MqBrokerNamespaceLister
	MqBrokerListerExpansion
}

// mqBrokerLister implements the MqBrokerLister interface.
type mqBrokerLister struct {
	indexer cache.Indexer
}

// NewMqBrokerLister returns a new MqBrokerLister.
func NewMqBrokerLister(indexer cache.Indexer) MqBrokerLister {
	return &mqBrokerLister{indexer: indexer}
}

// List lists all MqBrokers in the indexer.
func (s *mqBrokerLister) List(selector labels.Selector) (ret []*v1alpha1.MqBroker, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MqBroker))
	})
	return ret, err
}

// MqBrokers returns an object that can list and get MqBrokers.
func (s *mqBrokerLister) MqBrokers(namespace string) MqBrokerNamespaceLister {
	return mqBrokerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MqBrokerNamespaceLister helps list and get MqBrokers.
type MqBrokerNamespaceLister interface {
	// List lists all MqBrokers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.MqBroker, err error)
	// Get retrieves the MqBroker from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.MqBroker, error)
	MqBrokerNamespaceListerExpansion
}

// mqBrokerNamespaceLister implements the MqBrokerNamespaceLister
// interface.
type mqBrokerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MqBrokers in the indexer for a given namespace.
func (s mqBrokerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MqBroker, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MqBroker))
	})
	return ret, err
}

// Get retrieves the MqBroker from the indexer for a given namespace and name.
func (s mqBrokerNamespaceLister) Get(name string) (*v1alpha1.MqBroker, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("mqbroker"), name)
	}
	return obj.(*v1alpha1.MqBroker), nil
}
