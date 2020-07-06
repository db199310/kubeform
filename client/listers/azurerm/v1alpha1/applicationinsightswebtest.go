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

// ApplicationInsightsWebTestLister helps list ApplicationInsightsWebTests.
type ApplicationInsightsWebTestLister interface {
	// List lists all ApplicationInsightsWebTests in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ApplicationInsightsWebTest, err error)
	// ApplicationInsightsWebTests returns an object that can list and get ApplicationInsightsWebTests.
	ApplicationInsightsWebTests(namespace string) ApplicationInsightsWebTestNamespaceLister
	ApplicationInsightsWebTestListerExpansion
}

// applicationInsightsWebTestLister implements the ApplicationInsightsWebTestLister interface.
type applicationInsightsWebTestLister struct {
	indexer cache.Indexer
}

// NewApplicationInsightsWebTestLister returns a new ApplicationInsightsWebTestLister.
func NewApplicationInsightsWebTestLister(indexer cache.Indexer) ApplicationInsightsWebTestLister {
	return &applicationInsightsWebTestLister{indexer: indexer}
}

// List lists all ApplicationInsightsWebTests in the indexer.
func (s *applicationInsightsWebTestLister) List(selector labels.Selector) (ret []*v1alpha1.ApplicationInsightsWebTest, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApplicationInsightsWebTest))
	})
	return ret, err
}

// ApplicationInsightsWebTests returns an object that can list and get ApplicationInsightsWebTests.
func (s *applicationInsightsWebTestLister) ApplicationInsightsWebTests(namespace string) ApplicationInsightsWebTestNamespaceLister {
	return applicationInsightsWebTestNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApplicationInsightsWebTestNamespaceLister helps list and get ApplicationInsightsWebTests.
type ApplicationInsightsWebTestNamespaceLister interface {
	// List lists all ApplicationInsightsWebTests in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ApplicationInsightsWebTest, err error)
	// Get retrieves the ApplicationInsightsWebTest from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ApplicationInsightsWebTest, error)
	ApplicationInsightsWebTestNamespaceListerExpansion
}

// applicationInsightsWebTestNamespaceLister implements the ApplicationInsightsWebTestNamespaceLister
// interface.
type applicationInsightsWebTestNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApplicationInsightsWebTests in the indexer for a given namespace.
func (s applicationInsightsWebTestNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ApplicationInsightsWebTest, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApplicationInsightsWebTest))
	})
	return ret, err
}

// Get retrieves the ApplicationInsightsWebTest from the indexer for a given namespace and name.
func (s applicationInsightsWebTestNamespaceLister) Get(name string) (*v1alpha1.ApplicationInsightsWebTest, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("applicationinsightswebtest"), name)
	}
	return obj.(*v1alpha1.ApplicationInsightsWebTest), nil
}
