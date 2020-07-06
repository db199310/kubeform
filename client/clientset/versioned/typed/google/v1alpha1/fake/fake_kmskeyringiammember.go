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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakeKmsKeyRingIamMembers implements KmsKeyRingIamMemberInterface
type FakeKmsKeyRingIamMembers struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var kmskeyringiammembersResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "kmskeyringiammembers"}

var kmskeyringiammembersKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "KmsKeyRingIamMember"}

// Get takes name of the kmsKeyRingIamMember, and returns the corresponding kmsKeyRingIamMember object, and an error if there is any.
func (c *FakeKmsKeyRingIamMembers) Get(name string, options v1.GetOptions) (result *v1alpha1.KmsKeyRingIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kmskeyringiammembersResource, c.ns, name), &v1alpha1.KmsKeyRingIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KmsKeyRingIamMember), err
}

// List takes label and field selectors, and returns the list of KmsKeyRingIamMembers that match those selectors.
func (c *FakeKmsKeyRingIamMembers) List(opts v1.ListOptions) (result *v1alpha1.KmsKeyRingIamMemberList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kmskeyringiammembersResource, kmskeyringiammembersKind, c.ns, opts), &v1alpha1.KmsKeyRingIamMemberList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KmsKeyRingIamMemberList{ListMeta: obj.(*v1alpha1.KmsKeyRingIamMemberList).ListMeta}
	for _, item := range obj.(*v1alpha1.KmsKeyRingIamMemberList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kmsKeyRingIamMembers.
func (c *FakeKmsKeyRingIamMembers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kmskeyringiammembersResource, c.ns, opts))

}

// Create takes the representation of a kmsKeyRingIamMember and creates it.  Returns the server's representation of the kmsKeyRingIamMember, and an error, if there is any.
func (c *FakeKmsKeyRingIamMembers) Create(kmsKeyRingIamMember *v1alpha1.KmsKeyRingIamMember) (result *v1alpha1.KmsKeyRingIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kmskeyringiammembersResource, c.ns, kmsKeyRingIamMember), &v1alpha1.KmsKeyRingIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KmsKeyRingIamMember), err
}

// Update takes the representation of a kmsKeyRingIamMember and updates it. Returns the server's representation of the kmsKeyRingIamMember, and an error, if there is any.
func (c *FakeKmsKeyRingIamMembers) Update(kmsKeyRingIamMember *v1alpha1.KmsKeyRingIamMember) (result *v1alpha1.KmsKeyRingIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kmskeyringiammembersResource, c.ns, kmsKeyRingIamMember), &v1alpha1.KmsKeyRingIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KmsKeyRingIamMember), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKmsKeyRingIamMembers) UpdateStatus(kmsKeyRingIamMember *v1alpha1.KmsKeyRingIamMember) (*v1alpha1.KmsKeyRingIamMember, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(kmskeyringiammembersResource, "status", c.ns, kmsKeyRingIamMember), &v1alpha1.KmsKeyRingIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KmsKeyRingIamMember), err
}

// Delete takes name of the kmsKeyRingIamMember and deletes it. Returns an error if one occurs.
func (c *FakeKmsKeyRingIamMembers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(kmskeyringiammembersResource, c.ns, name), &v1alpha1.KmsKeyRingIamMember{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKmsKeyRingIamMembers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kmskeyringiammembersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.KmsKeyRingIamMemberList{})
	return err
}

// Patch applies the patch and returns the patched kmsKeyRingIamMember.
func (c *FakeKmsKeyRingIamMembers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KmsKeyRingIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kmskeyringiammembersResource, c.ns, name, pt, data, subresources...), &v1alpha1.KmsKeyRingIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KmsKeyRingIamMember), err
}
