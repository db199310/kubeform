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

// CdnProfileLister helps list CdnProfiles.
type CdnProfileLister interface {
	// List lists all CdnProfiles in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.CdnProfile, err error)
	// CdnProfiles returns an object that can list and get CdnProfiles.
	CdnProfiles(namespace string) CdnProfileNamespaceLister
	CdnProfileListerExpansion
}

// cdnProfileLister implements the CdnProfileLister interface.
type cdnProfileLister struct {
	indexer cache.Indexer
}

// NewCdnProfileLister returns a new CdnProfileLister.
func NewCdnProfileLister(indexer cache.Indexer) CdnProfileLister {
	return &cdnProfileLister{indexer: indexer}
}

// List lists all CdnProfiles in the indexer.
func (s *cdnProfileLister) List(selector labels.Selector) (ret []*v1alpha1.CdnProfile, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CdnProfile))
	})
	return ret, err
}

// CdnProfiles returns an object that can list and get CdnProfiles.
func (s *cdnProfileLister) CdnProfiles(namespace string) CdnProfileNamespaceLister {
	return cdnProfileNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CdnProfileNamespaceLister helps list and get CdnProfiles.
type CdnProfileNamespaceLister interface {
	// List lists all CdnProfiles in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.CdnProfile, err error)
	// Get retrieves the CdnProfile from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.CdnProfile, error)
	CdnProfileNamespaceListerExpansion
}

// cdnProfileNamespaceLister implements the CdnProfileNamespaceLister
// interface.
type cdnProfileNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CdnProfiles in the indexer for a given namespace.
func (s cdnProfileNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CdnProfile, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CdnProfile))
	})
	return ret, err
}

// Get retrieves the CdnProfile from the indexer for a given namespace and name.
func (s cdnProfileNamespaceLister) Get(name string) (*v1alpha1.CdnProfile, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cdnprofile"), name)
	}
	return obj.(*v1alpha1.CdnProfile), nil
}
