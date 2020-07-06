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

// DefaultSecurityGroupLister helps list DefaultSecurityGroups.
type DefaultSecurityGroupLister interface {
	// List lists all DefaultSecurityGroups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DefaultSecurityGroup, err error)
	// DefaultSecurityGroups returns an object that can list and get DefaultSecurityGroups.
	DefaultSecurityGroups(namespace string) DefaultSecurityGroupNamespaceLister
	DefaultSecurityGroupListerExpansion
}

// defaultSecurityGroupLister implements the DefaultSecurityGroupLister interface.
type defaultSecurityGroupLister struct {
	indexer cache.Indexer
}

// NewDefaultSecurityGroupLister returns a new DefaultSecurityGroupLister.
func NewDefaultSecurityGroupLister(indexer cache.Indexer) DefaultSecurityGroupLister {
	return &defaultSecurityGroupLister{indexer: indexer}
}

// List lists all DefaultSecurityGroups in the indexer.
func (s *defaultSecurityGroupLister) List(selector labels.Selector) (ret []*v1alpha1.DefaultSecurityGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DefaultSecurityGroup))
	})
	return ret, err
}

// DefaultSecurityGroups returns an object that can list and get DefaultSecurityGroups.
func (s *defaultSecurityGroupLister) DefaultSecurityGroups(namespace string) DefaultSecurityGroupNamespaceLister {
	return defaultSecurityGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DefaultSecurityGroupNamespaceLister helps list and get DefaultSecurityGroups.
type DefaultSecurityGroupNamespaceLister interface {
	// List lists all DefaultSecurityGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DefaultSecurityGroup, err error)
	// Get retrieves the DefaultSecurityGroup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DefaultSecurityGroup, error)
	DefaultSecurityGroupNamespaceListerExpansion
}

// defaultSecurityGroupNamespaceLister implements the DefaultSecurityGroupNamespaceLister
// interface.
type defaultSecurityGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DefaultSecurityGroups in the indexer for a given namespace.
func (s defaultSecurityGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DefaultSecurityGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DefaultSecurityGroup))
	})
	return ret, err
}

// Get retrieves the DefaultSecurityGroup from the indexer for a given namespace and name.
func (s defaultSecurityGroupNamespaceLister) Get(name string) (*v1alpha1.DefaultSecurityGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("defaultsecuritygroup"), name)
	}
	return obj.(*v1alpha1.DefaultSecurityGroup), nil
}
