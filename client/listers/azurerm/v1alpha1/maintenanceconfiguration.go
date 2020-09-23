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

// MaintenanceConfigurationLister helps list MaintenanceConfigurations.
type MaintenanceConfigurationLister interface {
	// List lists all MaintenanceConfigurations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.MaintenanceConfiguration, err error)
	// MaintenanceConfigurations returns an object that can list and get MaintenanceConfigurations.
	MaintenanceConfigurations(namespace string) MaintenanceConfigurationNamespaceLister
	MaintenanceConfigurationListerExpansion
}

// maintenanceConfigurationLister implements the MaintenanceConfigurationLister interface.
type maintenanceConfigurationLister struct {
	indexer cache.Indexer
}

// NewMaintenanceConfigurationLister returns a new MaintenanceConfigurationLister.
func NewMaintenanceConfigurationLister(indexer cache.Indexer) MaintenanceConfigurationLister {
	return &maintenanceConfigurationLister{indexer: indexer}
}

// List lists all MaintenanceConfigurations in the indexer.
func (s *maintenanceConfigurationLister) List(selector labels.Selector) (ret []*v1alpha1.MaintenanceConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MaintenanceConfiguration))
	})
	return ret, err
}

// MaintenanceConfigurations returns an object that can list and get MaintenanceConfigurations.
func (s *maintenanceConfigurationLister) MaintenanceConfigurations(namespace string) MaintenanceConfigurationNamespaceLister {
	return maintenanceConfigurationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MaintenanceConfigurationNamespaceLister helps list and get MaintenanceConfigurations.
type MaintenanceConfigurationNamespaceLister interface {
	// List lists all MaintenanceConfigurations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.MaintenanceConfiguration, err error)
	// Get retrieves the MaintenanceConfiguration from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.MaintenanceConfiguration, error)
	MaintenanceConfigurationNamespaceListerExpansion
}

// maintenanceConfigurationNamespaceLister implements the MaintenanceConfigurationNamespaceLister
// interface.
type maintenanceConfigurationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MaintenanceConfigurations in the indexer for a given namespace.
func (s maintenanceConfigurationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MaintenanceConfiguration, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MaintenanceConfiguration))
	})
	return ret, err
}

// Get retrieves the MaintenanceConfiguration from the indexer for a given namespace and name.
func (s maintenanceConfigurationNamespaceLister) Get(name string) (*v1alpha1.MaintenanceConfiguration, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("maintenanceconfiguration"), name)
	}
	return obj.(*v1alpha1.MaintenanceConfiguration), nil
}