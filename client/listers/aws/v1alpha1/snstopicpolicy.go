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

// SnsTopicPolicyLister helps list SnsTopicPolicies.
type SnsTopicPolicyLister interface {
	// List lists all SnsTopicPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SnsTopicPolicy, err error)
	// SnsTopicPolicies returns an object that can list and get SnsTopicPolicies.
	SnsTopicPolicies(namespace string) SnsTopicPolicyNamespaceLister
	SnsTopicPolicyListerExpansion
}

// snsTopicPolicyLister implements the SnsTopicPolicyLister interface.
type snsTopicPolicyLister struct {
	indexer cache.Indexer
}

// NewSnsTopicPolicyLister returns a new SnsTopicPolicyLister.
func NewSnsTopicPolicyLister(indexer cache.Indexer) SnsTopicPolicyLister {
	return &snsTopicPolicyLister{indexer: indexer}
}

// List lists all SnsTopicPolicies in the indexer.
func (s *snsTopicPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.SnsTopicPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SnsTopicPolicy))
	})
	return ret, err
}

// SnsTopicPolicies returns an object that can list and get SnsTopicPolicies.
func (s *snsTopicPolicyLister) SnsTopicPolicies(namespace string) SnsTopicPolicyNamespaceLister {
	return snsTopicPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SnsTopicPolicyNamespaceLister helps list and get SnsTopicPolicies.
type SnsTopicPolicyNamespaceLister interface {
	// List lists all SnsTopicPolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SnsTopicPolicy, err error)
	// Get retrieves the SnsTopicPolicy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SnsTopicPolicy, error)
	SnsTopicPolicyNamespaceListerExpansion
}

// snsTopicPolicyNamespaceLister implements the SnsTopicPolicyNamespaceLister
// interface.
type snsTopicPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SnsTopicPolicies in the indexer for a given namespace.
func (s snsTopicPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SnsTopicPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SnsTopicPolicy))
	})
	return ret, err
}

// Get retrieves the SnsTopicPolicy from the indexer for a given namespace and name.
func (s snsTopicPolicyNamespaceLister) Get(name string) (*v1alpha1.SnsTopicPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("snstopicpolicy"), name)
	}
	return obj.(*v1alpha1.SnsTopicPolicy), nil
}
