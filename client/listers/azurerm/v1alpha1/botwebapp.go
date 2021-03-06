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

// BotWebAppLister helps list BotWebApps.
type BotWebAppLister interface {
	// List lists all BotWebApps in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.BotWebApp, err error)
	// BotWebApps returns an object that can list and get BotWebApps.
	BotWebApps(namespace string) BotWebAppNamespaceLister
	BotWebAppListerExpansion
}

// botWebAppLister implements the BotWebAppLister interface.
type botWebAppLister struct {
	indexer cache.Indexer
}

// NewBotWebAppLister returns a new BotWebAppLister.
func NewBotWebAppLister(indexer cache.Indexer) BotWebAppLister {
	return &botWebAppLister{indexer: indexer}
}

// List lists all BotWebApps in the indexer.
func (s *botWebAppLister) List(selector labels.Selector) (ret []*v1alpha1.BotWebApp, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BotWebApp))
	})
	return ret, err
}

// BotWebApps returns an object that can list and get BotWebApps.
func (s *botWebAppLister) BotWebApps(namespace string) BotWebAppNamespaceLister {
	return botWebAppNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BotWebAppNamespaceLister helps list and get BotWebApps.
type BotWebAppNamespaceLister interface {
	// List lists all BotWebApps in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.BotWebApp, err error)
	// Get retrieves the BotWebApp from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.BotWebApp, error)
	BotWebAppNamespaceListerExpansion
}

// botWebAppNamespaceLister implements the BotWebAppNamespaceLister
// interface.
type botWebAppNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BotWebApps in the indexer for a given namespace.
func (s botWebAppNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BotWebApp, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BotWebApp))
	})
	return ret, err
}

// Get retrieves the BotWebApp from the indexer for a given namespace and name.
func (s botWebAppNamespaceLister) Get(name string) (*v1alpha1.BotWebApp, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("botwebapp"), name)
	}
	return obj.(*v1alpha1.BotWebApp), nil
}
