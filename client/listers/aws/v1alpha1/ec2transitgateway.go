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

// Ec2TransitGatewayLister helps list Ec2TransitGateways.
type Ec2TransitGatewayLister interface {
	// List lists all Ec2TransitGateways in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Ec2TransitGateway, err error)
	// Ec2TransitGateways returns an object that can list and get Ec2TransitGateways.
	Ec2TransitGateways(namespace string) Ec2TransitGatewayNamespaceLister
	Ec2TransitGatewayListerExpansion
}

// ec2TransitGatewayLister implements the Ec2TransitGatewayLister interface.
type ec2TransitGatewayLister struct {
	indexer cache.Indexer
}

// NewEc2TransitGatewayLister returns a new Ec2TransitGatewayLister.
func NewEc2TransitGatewayLister(indexer cache.Indexer) Ec2TransitGatewayLister {
	return &ec2TransitGatewayLister{indexer: indexer}
}

// List lists all Ec2TransitGateways in the indexer.
func (s *ec2TransitGatewayLister) List(selector labels.Selector) (ret []*v1alpha1.Ec2TransitGateway, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Ec2TransitGateway))
	})
	return ret, err
}

// Ec2TransitGateways returns an object that can list and get Ec2TransitGateways.
func (s *ec2TransitGatewayLister) Ec2TransitGateways(namespace string) Ec2TransitGatewayNamespaceLister {
	return ec2TransitGatewayNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// Ec2TransitGatewayNamespaceLister helps list and get Ec2TransitGateways.
type Ec2TransitGatewayNamespaceLister interface {
	// List lists all Ec2TransitGateways in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Ec2TransitGateway, err error)
	// Get retrieves the Ec2TransitGateway from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Ec2TransitGateway, error)
	Ec2TransitGatewayNamespaceListerExpansion
}

// ec2TransitGatewayNamespaceLister implements the Ec2TransitGatewayNamespaceLister
// interface.
type ec2TransitGatewayNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Ec2TransitGateways in the indexer for a given namespace.
func (s ec2TransitGatewayNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Ec2TransitGateway, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Ec2TransitGateway))
	})
	return ret, err
}

// Get retrieves the Ec2TransitGateway from the indexer for a given namespace and name.
func (s ec2TransitGatewayNamespaceLister) Get(name string) (*v1alpha1.Ec2TransitGateway, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("ec2transitgateway"), name)
	}
	return obj.(*v1alpha1.Ec2TransitGateway), nil
}
