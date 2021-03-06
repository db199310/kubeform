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

// BotChannelMsTeamsLister helps list BotChannelMsTeamses.
type BotChannelMsTeamsLister interface {
	// List lists all BotChannelMsTeamses in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.BotChannelMsTeams, err error)
	// BotChannelMsTeamses returns an object that can list and get BotChannelMsTeamses.
	BotChannelMsTeamses(namespace string) BotChannelMsTeamsNamespaceLister
	BotChannelMsTeamsListerExpansion
}

// botChannelMsTeamsLister implements the BotChannelMsTeamsLister interface.
type botChannelMsTeamsLister struct {
	indexer cache.Indexer
}

// NewBotChannelMsTeamsLister returns a new BotChannelMsTeamsLister.
func NewBotChannelMsTeamsLister(indexer cache.Indexer) BotChannelMsTeamsLister {
	return &botChannelMsTeamsLister{indexer: indexer}
}

// List lists all BotChannelMsTeamses in the indexer.
func (s *botChannelMsTeamsLister) List(selector labels.Selector) (ret []*v1alpha1.BotChannelMsTeams, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BotChannelMsTeams))
	})
	return ret, err
}

// BotChannelMsTeamses returns an object that can list and get BotChannelMsTeamses.
func (s *botChannelMsTeamsLister) BotChannelMsTeamses(namespace string) BotChannelMsTeamsNamespaceLister {
	return botChannelMsTeamsNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BotChannelMsTeamsNamespaceLister helps list and get BotChannelMsTeamses.
type BotChannelMsTeamsNamespaceLister interface {
	// List lists all BotChannelMsTeamses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.BotChannelMsTeams, err error)
	// Get retrieves the BotChannelMsTeams from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.BotChannelMsTeams, error)
	BotChannelMsTeamsNamespaceListerExpansion
}

// botChannelMsTeamsNamespaceLister implements the BotChannelMsTeamsNamespaceLister
// interface.
type botChannelMsTeamsNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BotChannelMsTeamses in the indexer for a given namespace.
func (s botChannelMsTeamsNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BotChannelMsTeams, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BotChannelMsTeams))
	})
	return ret, err
}

// Get retrieves the BotChannelMsTeams from the indexer for a given namespace and name.
func (s botChannelMsTeamsNamespaceLister) Get(name string) (*v1alpha1.BotChannelMsTeams, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("botchannelmsteams"), name)
	}
	return obj.(*v1alpha1.BotChannelMsTeams), nil
}
