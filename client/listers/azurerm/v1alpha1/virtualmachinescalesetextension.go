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

// VirtualMachineScaleSetExtensionLister helps list VirtualMachineScaleSetExtensions.
type VirtualMachineScaleSetExtensionLister interface {
	// List lists all VirtualMachineScaleSetExtensions in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.VirtualMachineScaleSetExtension, err error)
	// VirtualMachineScaleSetExtensions returns an object that can list and get VirtualMachineScaleSetExtensions.
	VirtualMachineScaleSetExtensions(namespace string) VirtualMachineScaleSetExtensionNamespaceLister
	VirtualMachineScaleSetExtensionListerExpansion
}

// virtualMachineScaleSetExtensionLister implements the VirtualMachineScaleSetExtensionLister interface.
type virtualMachineScaleSetExtensionLister struct {
	indexer cache.Indexer
}

// NewVirtualMachineScaleSetExtensionLister returns a new VirtualMachineScaleSetExtensionLister.
func NewVirtualMachineScaleSetExtensionLister(indexer cache.Indexer) VirtualMachineScaleSetExtensionLister {
	return &virtualMachineScaleSetExtensionLister{indexer: indexer}
}

// List lists all VirtualMachineScaleSetExtensions in the indexer.
func (s *virtualMachineScaleSetExtensionLister) List(selector labels.Selector) (ret []*v1alpha1.VirtualMachineScaleSetExtension, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VirtualMachineScaleSetExtension))
	})
	return ret, err
}

// VirtualMachineScaleSetExtensions returns an object that can list and get VirtualMachineScaleSetExtensions.
func (s *virtualMachineScaleSetExtensionLister) VirtualMachineScaleSetExtensions(namespace string) VirtualMachineScaleSetExtensionNamespaceLister {
	return virtualMachineScaleSetExtensionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VirtualMachineScaleSetExtensionNamespaceLister helps list and get VirtualMachineScaleSetExtensions.
type VirtualMachineScaleSetExtensionNamespaceLister interface {
	// List lists all VirtualMachineScaleSetExtensions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.VirtualMachineScaleSetExtension, err error)
	// Get retrieves the VirtualMachineScaleSetExtension from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.VirtualMachineScaleSetExtension, error)
	VirtualMachineScaleSetExtensionNamespaceListerExpansion
}

// virtualMachineScaleSetExtensionNamespaceLister implements the VirtualMachineScaleSetExtensionNamespaceLister
// interface.
type virtualMachineScaleSetExtensionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VirtualMachineScaleSetExtensions in the indexer for a given namespace.
func (s virtualMachineScaleSetExtensionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.VirtualMachineScaleSetExtension, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VirtualMachineScaleSetExtension))
	})
	return ret, err
}

// Get retrieves the VirtualMachineScaleSetExtension from the indexer for a given namespace and name.
func (s virtualMachineScaleSetExtensionNamespaceLister) Get(name string) (*v1alpha1.VirtualMachineScaleSetExtension, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("virtualmachinescalesetextension"), name)
	}
	return obj.(*v1alpha1.VirtualMachineScaleSetExtension), nil
}