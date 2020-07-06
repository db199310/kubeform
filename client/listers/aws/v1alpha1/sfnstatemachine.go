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

// SfnStateMachineLister helps list SfnStateMachines.
type SfnStateMachineLister interface {
	// List lists all SfnStateMachines in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SfnStateMachine, err error)
	// SfnStateMachines returns an object that can list and get SfnStateMachines.
	SfnStateMachines(namespace string) SfnStateMachineNamespaceLister
	SfnStateMachineListerExpansion
}

// sfnStateMachineLister implements the SfnStateMachineLister interface.
type sfnStateMachineLister struct {
	indexer cache.Indexer
}

// NewSfnStateMachineLister returns a new SfnStateMachineLister.
func NewSfnStateMachineLister(indexer cache.Indexer) SfnStateMachineLister {
	return &sfnStateMachineLister{indexer: indexer}
}

// List lists all SfnStateMachines in the indexer.
func (s *sfnStateMachineLister) List(selector labels.Selector) (ret []*v1alpha1.SfnStateMachine, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SfnStateMachine))
	})
	return ret, err
}

// SfnStateMachines returns an object that can list and get SfnStateMachines.
func (s *sfnStateMachineLister) SfnStateMachines(namespace string) SfnStateMachineNamespaceLister {
	return sfnStateMachineNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SfnStateMachineNamespaceLister helps list and get SfnStateMachines.
type SfnStateMachineNamespaceLister interface {
	// List lists all SfnStateMachines in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SfnStateMachine, err error)
	// Get retrieves the SfnStateMachine from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SfnStateMachine, error)
	SfnStateMachineNamespaceListerExpansion
}

// sfnStateMachineNamespaceLister implements the SfnStateMachineNamespaceLister
// interface.
type sfnStateMachineNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SfnStateMachines in the indexer for a given namespace.
func (s sfnStateMachineNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SfnStateMachine, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SfnStateMachine))
	})
	return ret, err
}

// Get retrieves the SfnStateMachine from the indexer for a given namespace and name.
func (s sfnStateMachineNamespaceLister) Get(name string) (*v1alpha1.SfnStateMachine, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("sfnstatemachine"), name)
	}
	return obj.(*v1alpha1.SfnStateMachine), nil
}
