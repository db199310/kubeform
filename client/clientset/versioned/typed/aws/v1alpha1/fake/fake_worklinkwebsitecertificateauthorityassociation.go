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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeWorklinkWebsiteCertificateAuthorityAssociations implements WorklinkWebsiteCertificateAuthorityAssociationInterface
type FakeWorklinkWebsiteCertificateAuthorityAssociations struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var worklinkwebsitecertificateauthorityassociationsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "worklinkwebsitecertificateauthorityassociations"}

var worklinkwebsitecertificateauthorityassociationsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "WorklinkWebsiteCertificateAuthorityAssociation"}

// Get takes name of the worklinkWebsiteCertificateAuthorityAssociation, and returns the corresponding worklinkWebsiteCertificateAuthorityAssociation object, and an error if there is any.
func (c *FakeWorklinkWebsiteCertificateAuthorityAssociations) Get(name string, options v1.GetOptions) (result *v1alpha1.WorklinkWebsiteCertificateAuthorityAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(worklinkwebsitecertificateauthorityassociationsResource, c.ns, name), &v1alpha1.WorklinkWebsiteCertificateAuthorityAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorklinkWebsiteCertificateAuthorityAssociation), err
}

// List takes label and field selectors, and returns the list of WorklinkWebsiteCertificateAuthorityAssociations that match those selectors.
func (c *FakeWorklinkWebsiteCertificateAuthorityAssociations) List(opts v1.ListOptions) (result *v1alpha1.WorklinkWebsiteCertificateAuthorityAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(worklinkwebsitecertificateauthorityassociationsResource, worklinkwebsitecertificateauthorityassociationsKind, c.ns, opts), &v1alpha1.WorklinkWebsiteCertificateAuthorityAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.WorklinkWebsiteCertificateAuthorityAssociationList{ListMeta: obj.(*v1alpha1.WorklinkWebsiteCertificateAuthorityAssociationList).ListMeta}
	for _, item := range obj.(*v1alpha1.WorklinkWebsiteCertificateAuthorityAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested worklinkWebsiteCertificateAuthorityAssociations.
func (c *FakeWorklinkWebsiteCertificateAuthorityAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(worklinkwebsitecertificateauthorityassociationsResource, c.ns, opts))

}

// Create takes the representation of a worklinkWebsiteCertificateAuthorityAssociation and creates it.  Returns the server's representation of the worklinkWebsiteCertificateAuthorityAssociation, and an error, if there is any.
func (c *FakeWorklinkWebsiteCertificateAuthorityAssociations) Create(worklinkWebsiteCertificateAuthorityAssociation *v1alpha1.WorklinkWebsiteCertificateAuthorityAssociation) (result *v1alpha1.WorklinkWebsiteCertificateAuthorityAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(worklinkwebsitecertificateauthorityassociationsResource, c.ns, worklinkWebsiteCertificateAuthorityAssociation), &v1alpha1.WorklinkWebsiteCertificateAuthorityAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorklinkWebsiteCertificateAuthorityAssociation), err
}

// Update takes the representation of a worklinkWebsiteCertificateAuthorityAssociation and updates it. Returns the server's representation of the worklinkWebsiteCertificateAuthorityAssociation, and an error, if there is any.
func (c *FakeWorklinkWebsiteCertificateAuthorityAssociations) Update(worklinkWebsiteCertificateAuthorityAssociation *v1alpha1.WorklinkWebsiteCertificateAuthorityAssociation) (result *v1alpha1.WorklinkWebsiteCertificateAuthorityAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(worklinkwebsitecertificateauthorityassociationsResource, c.ns, worklinkWebsiteCertificateAuthorityAssociation), &v1alpha1.WorklinkWebsiteCertificateAuthorityAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorklinkWebsiteCertificateAuthorityAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeWorklinkWebsiteCertificateAuthorityAssociations) UpdateStatus(worklinkWebsiteCertificateAuthorityAssociation *v1alpha1.WorklinkWebsiteCertificateAuthorityAssociation) (*v1alpha1.WorklinkWebsiteCertificateAuthorityAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(worklinkwebsitecertificateauthorityassociationsResource, "status", c.ns, worklinkWebsiteCertificateAuthorityAssociation), &v1alpha1.WorklinkWebsiteCertificateAuthorityAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorklinkWebsiteCertificateAuthorityAssociation), err
}

// Delete takes name of the worklinkWebsiteCertificateAuthorityAssociation and deletes it. Returns an error if one occurs.
func (c *FakeWorklinkWebsiteCertificateAuthorityAssociations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(worklinkwebsitecertificateauthorityassociationsResource, c.ns, name), &v1alpha1.WorklinkWebsiteCertificateAuthorityAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWorklinkWebsiteCertificateAuthorityAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(worklinkwebsitecertificateauthorityassociationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.WorklinkWebsiteCertificateAuthorityAssociationList{})
	return err
}

// Patch applies the patch and returns the patched worklinkWebsiteCertificateAuthorityAssociation.
func (c *FakeWorklinkWebsiteCertificateAuthorityAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WorklinkWebsiteCertificateAuthorityAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(worklinkwebsitecertificateauthorityassociationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.WorklinkWebsiteCertificateAuthorityAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorklinkWebsiteCertificateAuthorityAssociation), err
}
