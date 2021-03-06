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

// SiteRecoveryReplicationPolicyLister helps list SiteRecoveryReplicationPolicies.
type SiteRecoveryReplicationPolicyLister interface {
	// List lists all SiteRecoveryReplicationPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SiteRecoveryReplicationPolicy, err error)
	// SiteRecoveryReplicationPolicies returns an object that can list and get SiteRecoveryReplicationPolicies.
	SiteRecoveryReplicationPolicies(namespace string) SiteRecoveryReplicationPolicyNamespaceLister
	SiteRecoveryReplicationPolicyListerExpansion
}

// siteRecoveryReplicationPolicyLister implements the SiteRecoveryReplicationPolicyLister interface.
type siteRecoveryReplicationPolicyLister struct {
	indexer cache.Indexer
}

// NewSiteRecoveryReplicationPolicyLister returns a new SiteRecoveryReplicationPolicyLister.
func NewSiteRecoveryReplicationPolicyLister(indexer cache.Indexer) SiteRecoveryReplicationPolicyLister {
	return &siteRecoveryReplicationPolicyLister{indexer: indexer}
}

// List lists all SiteRecoveryReplicationPolicies in the indexer.
func (s *siteRecoveryReplicationPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.SiteRecoveryReplicationPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SiteRecoveryReplicationPolicy))
	})
	return ret, err
}

// SiteRecoveryReplicationPolicies returns an object that can list and get SiteRecoveryReplicationPolicies.
func (s *siteRecoveryReplicationPolicyLister) SiteRecoveryReplicationPolicies(namespace string) SiteRecoveryReplicationPolicyNamespaceLister {
	return siteRecoveryReplicationPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SiteRecoveryReplicationPolicyNamespaceLister helps list and get SiteRecoveryReplicationPolicies.
type SiteRecoveryReplicationPolicyNamespaceLister interface {
	// List lists all SiteRecoveryReplicationPolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SiteRecoveryReplicationPolicy, err error)
	// Get retrieves the SiteRecoveryReplicationPolicy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SiteRecoveryReplicationPolicy, error)
	SiteRecoveryReplicationPolicyNamespaceListerExpansion
}

// siteRecoveryReplicationPolicyNamespaceLister implements the SiteRecoveryReplicationPolicyNamespaceLister
// interface.
type siteRecoveryReplicationPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SiteRecoveryReplicationPolicies in the indexer for a given namespace.
func (s siteRecoveryReplicationPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SiteRecoveryReplicationPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SiteRecoveryReplicationPolicy))
	})
	return ret, err
}

// Get retrieves the SiteRecoveryReplicationPolicy from the indexer for a given namespace and name.
func (s siteRecoveryReplicationPolicyNamespaceLister) Get(name string) (*v1alpha1.SiteRecoveryReplicationPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("siterecoveryreplicationpolicy"), name)
	}
	return obj.(*v1alpha1.SiteRecoveryReplicationPolicy), nil
}
