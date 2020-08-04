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

// LogAnalyticsDatasourceWindowsPerformanceCounterLister helps list LogAnalyticsDatasourceWindowsPerformanceCounters.
type LogAnalyticsDatasourceWindowsPerformanceCounterLister interface {
	// List lists all LogAnalyticsDatasourceWindowsPerformanceCounters in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.LogAnalyticsDatasourceWindowsPerformanceCounter, err error)
	// LogAnalyticsDatasourceWindowsPerformanceCounters returns an object that can list and get LogAnalyticsDatasourceWindowsPerformanceCounters.
	LogAnalyticsDatasourceWindowsPerformanceCounters(namespace string) LogAnalyticsDatasourceWindowsPerformanceCounterNamespaceLister
	LogAnalyticsDatasourceWindowsPerformanceCounterListerExpansion
}

// logAnalyticsDatasourceWindowsPerformanceCounterLister implements the LogAnalyticsDatasourceWindowsPerformanceCounterLister interface.
type logAnalyticsDatasourceWindowsPerformanceCounterLister struct {
	indexer cache.Indexer
}

// NewLogAnalyticsDatasourceWindowsPerformanceCounterLister returns a new LogAnalyticsDatasourceWindowsPerformanceCounterLister.
func NewLogAnalyticsDatasourceWindowsPerformanceCounterLister(indexer cache.Indexer) LogAnalyticsDatasourceWindowsPerformanceCounterLister {
	return &logAnalyticsDatasourceWindowsPerformanceCounterLister{indexer: indexer}
}

// List lists all LogAnalyticsDatasourceWindowsPerformanceCounters in the indexer.
func (s *logAnalyticsDatasourceWindowsPerformanceCounterLister) List(selector labels.Selector) (ret []*v1alpha1.LogAnalyticsDatasourceWindowsPerformanceCounter, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LogAnalyticsDatasourceWindowsPerformanceCounter))
	})
	return ret, err
}

// LogAnalyticsDatasourceWindowsPerformanceCounters returns an object that can list and get LogAnalyticsDatasourceWindowsPerformanceCounters.
func (s *logAnalyticsDatasourceWindowsPerformanceCounterLister) LogAnalyticsDatasourceWindowsPerformanceCounters(namespace string) LogAnalyticsDatasourceWindowsPerformanceCounterNamespaceLister {
	return logAnalyticsDatasourceWindowsPerformanceCounterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LogAnalyticsDatasourceWindowsPerformanceCounterNamespaceLister helps list and get LogAnalyticsDatasourceWindowsPerformanceCounters.
type LogAnalyticsDatasourceWindowsPerformanceCounterNamespaceLister interface {
	// List lists all LogAnalyticsDatasourceWindowsPerformanceCounters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.LogAnalyticsDatasourceWindowsPerformanceCounter, err error)
	// Get retrieves the LogAnalyticsDatasourceWindowsPerformanceCounter from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.LogAnalyticsDatasourceWindowsPerformanceCounter, error)
	LogAnalyticsDatasourceWindowsPerformanceCounterNamespaceListerExpansion
}

// logAnalyticsDatasourceWindowsPerformanceCounterNamespaceLister implements the LogAnalyticsDatasourceWindowsPerformanceCounterNamespaceLister
// interface.
type logAnalyticsDatasourceWindowsPerformanceCounterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LogAnalyticsDatasourceWindowsPerformanceCounters in the indexer for a given namespace.
func (s logAnalyticsDatasourceWindowsPerformanceCounterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LogAnalyticsDatasourceWindowsPerformanceCounter, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LogAnalyticsDatasourceWindowsPerformanceCounter))
	})
	return ret, err
}

// Get retrieves the LogAnalyticsDatasourceWindowsPerformanceCounter from the indexer for a given namespace and name.
func (s logAnalyticsDatasourceWindowsPerformanceCounterNamespaceLister) Get(name string) (*v1alpha1.LogAnalyticsDatasourceWindowsPerformanceCounter, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("loganalyticsdatasourcewindowsperformancecounter"), name)
	}
	return obj.(*v1alpha1.LogAnalyticsDatasourceWindowsPerformanceCounter), nil
}
