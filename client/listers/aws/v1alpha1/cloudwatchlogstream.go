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

// CloudwatchLogStreamLister helps list CloudwatchLogStreams.
type CloudwatchLogStreamLister interface {
	// List lists all CloudwatchLogStreams in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.CloudwatchLogStream, err error)
	// CloudwatchLogStreams returns an object that can list and get CloudwatchLogStreams.
	CloudwatchLogStreams(namespace string) CloudwatchLogStreamNamespaceLister
	CloudwatchLogStreamListerExpansion
}

// cloudwatchLogStreamLister implements the CloudwatchLogStreamLister interface.
type cloudwatchLogStreamLister struct {
	indexer cache.Indexer
}

// NewCloudwatchLogStreamLister returns a new CloudwatchLogStreamLister.
func NewCloudwatchLogStreamLister(indexer cache.Indexer) CloudwatchLogStreamLister {
	return &cloudwatchLogStreamLister{indexer: indexer}
}

// List lists all CloudwatchLogStreams in the indexer.
func (s *cloudwatchLogStreamLister) List(selector labels.Selector) (ret []*v1alpha1.CloudwatchLogStream, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CloudwatchLogStream))
	})
	return ret, err
}

// CloudwatchLogStreams returns an object that can list and get CloudwatchLogStreams.
func (s *cloudwatchLogStreamLister) CloudwatchLogStreams(namespace string) CloudwatchLogStreamNamespaceLister {
	return cloudwatchLogStreamNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CloudwatchLogStreamNamespaceLister helps list and get CloudwatchLogStreams.
type CloudwatchLogStreamNamespaceLister interface {
	// List lists all CloudwatchLogStreams in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.CloudwatchLogStream, err error)
	// Get retrieves the CloudwatchLogStream from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.CloudwatchLogStream, error)
	CloudwatchLogStreamNamespaceListerExpansion
}

// cloudwatchLogStreamNamespaceLister implements the CloudwatchLogStreamNamespaceLister
// interface.
type cloudwatchLogStreamNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CloudwatchLogStreams in the indexer for a given namespace.
func (s cloudwatchLogStreamNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CloudwatchLogStream, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CloudwatchLogStream))
	})
	return ret, err
}

// Get retrieves the CloudwatchLogStream from the indexer for a given namespace and name.
func (s cloudwatchLogStreamNamespaceLister) Get(name string) (*v1alpha1.CloudwatchLogStream, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cloudwatchlogstream"), name)
	}
	return obj.(*v1alpha1.CloudwatchLogStream), nil
}
