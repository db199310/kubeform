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

// MacieMemberAccountAssociationLister helps list MacieMemberAccountAssociations.
type MacieMemberAccountAssociationLister interface {
	// List lists all MacieMemberAccountAssociations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.MacieMemberAccountAssociation, err error)
	// MacieMemberAccountAssociations returns an object that can list and get MacieMemberAccountAssociations.
	MacieMemberAccountAssociations(namespace string) MacieMemberAccountAssociationNamespaceLister
	MacieMemberAccountAssociationListerExpansion
}

// macieMemberAccountAssociationLister implements the MacieMemberAccountAssociationLister interface.
type macieMemberAccountAssociationLister struct {
	indexer cache.Indexer
}

// NewMacieMemberAccountAssociationLister returns a new MacieMemberAccountAssociationLister.
func NewMacieMemberAccountAssociationLister(indexer cache.Indexer) MacieMemberAccountAssociationLister {
	return &macieMemberAccountAssociationLister{indexer: indexer}
}

// List lists all MacieMemberAccountAssociations in the indexer.
func (s *macieMemberAccountAssociationLister) List(selector labels.Selector) (ret []*v1alpha1.MacieMemberAccountAssociation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MacieMemberAccountAssociation))
	})
	return ret, err
}

// MacieMemberAccountAssociations returns an object that can list and get MacieMemberAccountAssociations.
func (s *macieMemberAccountAssociationLister) MacieMemberAccountAssociations(namespace string) MacieMemberAccountAssociationNamespaceLister {
	return macieMemberAccountAssociationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MacieMemberAccountAssociationNamespaceLister helps list and get MacieMemberAccountAssociations.
type MacieMemberAccountAssociationNamespaceLister interface {
	// List lists all MacieMemberAccountAssociations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.MacieMemberAccountAssociation, err error)
	// Get retrieves the MacieMemberAccountAssociation from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.MacieMemberAccountAssociation, error)
	MacieMemberAccountAssociationNamespaceListerExpansion
}

// macieMemberAccountAssociationNamespaceLister implements the MacieMemberAccountAssociationNamespaceLister
// interface.
type macieMemberAccountAssociationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MacieMemberAccountAssociations in the indexer for a given namespace.
func (s macieMemberAccountAssociationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MacieMemberAccountAssociation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MacieMemberAccountAssociation))
	})
	return ret, err
}

// Get retrieves the MacieMemberAccountAssociation from the indexer for a given namespace and name.
func (s macieMemberAccountAssociationNamespaceLister) Get(name string) (*v1alpha1.MacieMemberAccountAssociation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("maciememberaccountassociation"), name)
	}
	return obj.(*v1alpha1.MacieMemberAccountAssociation), nil
}
