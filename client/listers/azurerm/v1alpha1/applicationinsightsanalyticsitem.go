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

// ApplicationInsightsAnalyticsItemLister helps list ApplicationInsightsAnalyticsItems.
type ApplicationInsightsAnalyticsItemLister interface {
	// List lists all ApplicationInsightsAnalyticsItems in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ApplicationInsightsAnalyticsItem, err error)
	// ApplicationInsightsAnalyticsItems returns an object that can list and get ApplicationInsightsAnalyticsItems.
	ApplicationInsightsAnalyticsItems(namespace string) ApplicationInsightsAnalyticsItemNamespaceLister
	ApplicationInsightsAnalyticsItemListerExpansion
}

// applicationInsightsAnalyticsItemLister implements the ApplicationInsightsAnalyticsItemLister interface.
type applicationInsightsAnalyticsItemLister struct {
	indexer cache.Indexer
}

// NewApplicationInsightsAnalyticsItemLister returns a new ApplicationInsightsAnalyticsItemLister.
func NewApplicationInsightsAnalyticsItemLister(indexer cache.Indexer) ApplicationInsightsAnalyticsItemLister {
	return &applicationInsightsAnalyticsItemLister{indexer: indexer}
}

// List lists all ApplicationInsightsAnalyticsItems in the indexer.
func (s *applicationInsightsAnalyticsItemLister) List(selector labels.Selector) (ret []*v1alpha1.ApplicationInsightsAnalyticsItem, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApplicationInsightsAnalyticsItem))
	})
	return ret, err
}

// ApplicationInsightsAnalyticsItems returns an object that can list and get ApplicationInsightsAnalyticsItems.
func (s *applicationInsightsAnalyticsItemLister) ApplicationInsightsAnalyticsItems(namespace string) ApplicationInsightsAnalyticsItemNamespaceLister {
	return applicationInsightsAnalyticsItemNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApplicationInsightsAnalyticsItemNamespaceLister helps list and get ApplicationInsightsAnalyticsItems.
type ApplicationInsightsAnalyticsItemNamespaceLister interface {
	// List lists all ApplicationInsightsAnalyticsItems in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ApplicationInsightsAnalyticsItem, err error)
	// Get retrieves the ApplicationInsightsAnalyticsItem from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ApplicationInsightsAnalyticsItem, error)
	ApplicationInsightsAnalyticsItemNamespaceListerExpansion
}

// applicationInsightsAnalyticsItemNamespaceLister implements the ApplicationInsightsAnalyticsItemNamespaceLister
// interface.
type applicationInsightsAnalyticsItemNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApplicationInsightsAnalyticsItems in the indexer for a given namespace.
func (s applicationInsightsAnalyticsItemNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ApplicationInsightsAnalyticsItem, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApplicationInsightsAnalyticsItem))
	})
	return ret, err
}

// Get retrieves the ApplicationInsightsAnalyticsItem from the indexer for a given namespace and name.
func (s applicationInsightsAnalyticsItemNamespaceLister) Get(name string) (*v1alpha1.ApplicationInsightsAnalyticsItem, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("applicationinsightsanalyticsitem"), name)
	}
	return obj.(*v1alpha1.ApplicationInsightsAnalyticsItem), nil
}
