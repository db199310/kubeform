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

// IamPolicyAttachmentLister helps list IamPolicyAttachments.
type IamPolicyAttachmentLister interface {
	// List lists all IamPolicyAttachments in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.IamPolicyAttachment, err error)
	// IamPolicyAttachments returns an object that can list and get IamPolicyAttachments.
	IamPolicyAttachments(namespace string) IamPolicyAttachmentNamespaceLister
	IamPolicyAttachmentListerExpansion
}

// iamPolicyAttachmentLister implements the IamPolicyAttachmentLister interface.
type iamPolicyAttachmentLister struct {
	indexer cache.Indexer
}

// NewIamPolicyAttachmentLister returns a new IamPolicyAttachmentLister.
func NewIamPolicyAttachmentLister(indexer cache.Indexer) IamPolicyAttachmentLister {
	return &iamPolicyAttachmentLister{indexer: indexer}
}

// List lists all IamPolicyAttachments in the indexer.
func (s *iamPolicyAttachmentLister) List(selector labels.Selector) (ret []*v1alpha1.IamPolicyAttachment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IamPolicyAttachment))
	})
	return ret, err
}

// IamPolicyAttachments returns an object that can list and get IamPolicyAttachments.
func (s *iamPolicyAttachmentLister) IamPolicyAttachments(namespace string) IamPolicyAttachmentNamespaceLister {
	return iamPolicyAttachmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IamPolicyAttachmentNamespaceLister helps list and get IamPolicyAttachments.
type IamPolicyAttachmentNamespaceLister interface {
	// List lists all IamPolicyAttachments in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.IamPolicyAttachment, err error)
	// Get retrieves the IamPolicyAttachment from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.IamPolicyAttachment, error)
	IamPolicyAttachmentNamespaceListerExpansion
}

// iamPolicyAttachmentNamespaceLister implements the IamPolicyAttachmentNamespaceLister
// interface.
type iamPolicyAttachmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IamPolicyAttachments in the indexer for a given namespace.
func (s iamPolicyAttachmentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IamPolicyAttachment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IamPolicyAttachment))
	})
	return ret, err
}

// Get retrieves the IamPolicyAttachment from the indexer for a given namespace and name.
func (s iamPolicyAttachmentNamespaceLister) Get(name string) (*v1alpha1.IamPolicyAttachment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("iampolicyattachment"), name)
	}
	return obj.(*v1alpha1.IamPolicyAttachment), nil
}
