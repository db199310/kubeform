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

// VpcDHCPOptionsAssociationLister helps list VpcDHCPOptionsAssociations.
type VpcDHCPOptionsAssociationLister interface {
	// List lists all VpcDHCPOptionsAssociations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.VpcDHCPOptionsAssociation, err error)
	// VpcDHCPOptionsAssociations returns an object that can list and get VpcDHCPOptionsAssociations.
	VpcDHCPOptionsAssociations(namespace string) VpcDHCPOptionsAssociationNamespaceLister
	VpcDHCPOptionsAssociationListerExpansion
}

// vpcDHCPOptionsAssociationLister implements the VpcDHCPOptionsAssociationLister interface.
type vpcDHCPOptionsAssociationLister struct {
	indexer cache.Indexer
}

// NewVpcDHCPOptionsAssociationLister returns a new VpcDHCPOptionsAssociationLister.
func NewVpcDHCPOptionsAssociationLister(indexer cache.Indexer) VpcDHCPOptionsAssociationLister {
	return &vpcDHCPOptionsAssociationLister{indexer: indexer}
}

// List lists all VpcDHCPOptionsAssociations in the indexer.
func (s *vpcDHCPOptionsAssociationLister) List(selector labels.Selector) (ret []*v1alpha1.VpcDHCPOptionsAssociation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VpcDHCPOptionsAssociation))
	})
	return ret, err
}

// VpcDHCPOptionsAssociations returns an object that can list and get VpcDHCPOptionsAssociations.
func (s *vpcDHCPOptionsAssociationLister) VpcDHCPOptionsAssociations(namespace string) VpcDHCPOptionsAssociationNamespaceLister {
	return vpcDHCPOptionsAssociationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VpcDHCPOptionsAssociationNamespaceLister helps list and get VpcDHCPOptionsAssociations.
type VpcDHCPOptionsAssociationNamespaceLister interface {
	// List lists all VpcDHCPOptionsAssociations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.VpcDHCPOptionsAssociation, err error)
	// Get retrieves the VpcDHCPOptionsAssociation from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.VpcDHCPOptionsAssociation, error)
	VpcDHCPOptionsAssociationNamespaceListerExpansion
}

// vpcDHCPOptionsAssociationNamespaceLister implements the VpcDHCPOptionsAssociationNamespaceLister
// interface.
type vpcDHCPOptionsAssociationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VpcDHCPOptionsAssociations in the indexer for a given namespace.
func (s vpcDHCPOptionsAssociationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.VpcDHCPOptionsAssociation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VpcDHCPOptionsAssociation))
	})
	return ret, err
}

// Get retrieves the VpcDHCPOptionsAssociation from the indexer for a given namespace and name.
func (s vpcDHCPOptionsAssociationNamespaceLister) Get(name string) (*v1alpha1.VpcDHCPOptionsAssociation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("vpcdhcpoptionsassociation"), name)
	}
	return obj.(*v1alpha1.VpcDHCPOptionsAssociation), nil
}
