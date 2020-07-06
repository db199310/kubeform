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

// FakeSpannerInstanceIamMembers implements SpannerInstanceIamMemberInterface
type FakeSpannerInstanceIamMembers struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var spannerinstanceiammembersResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "spannerinstanceiammembers"}

var spannerinstanceiammembersKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "SpannerInstanceIamMember"}

// Get takes name of the spannerInstanceIamMember, and returns the corresponding spannerInstanceIamMember object, and an error if there is any.
func (c *FakeSpannerInstanceIamMembers) Get(name string, options v1.GetOptions) (result *v1alpha1.SpannerInstanceIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(spannerinstanceiammembersResource, c.ns, name), &v1alpha1.SpannerInstanceIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SpannerInstanceIamMember), err
}

// List takes label and field selectors, and returns the list of SpannerInstanceIamMembers that match those selectors.
func (c *FakeSpannerInstanceIamMembers) List(opts v1.ListOptions) (result *v1alpha1.SpannerInstanceIamMemberList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(spannerinstanceiammembersResource, spannerinstanceiammembersKind, c.ns, opts), &v1alpha1.SpannerInstanceIamMemberList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SpannerInstanceIamMemberList{ListMeta: obj.(*v1alpha1.SpannerInstanceIamMemberList).ListMeta}
	for _, item := range obj.(*v1alpha1.SpannerInstanceIamMemberList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested spannerInstanceIamMembers.
func (c *FakeSpannerInstanceIamMembers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(spannerinstanceiammembersResource, c.ns, opts))

}

// Create takes the representation of a spannerInstanceIamMember and creates it.  Returns the server's representation of the spannerInstanceIamMember, and an error, if there is any.
func (c *FakeSpannerInstanceIamMembers) Create(spannerInstanceIamMember *v1alpha1.SpannerInstanceIamMember) (result *v1alpha1.SpannerInstanceIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(spannerinstanceiammembersResource, c.ns, spannerInstanceIamMember), &v1alpha1.SpannerInstanceIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SpannerInstanceIamMember), err
}

// Update takes the representation of a spannerInstanceIamMember and updates it. Returns the server's representation of the spannerInstanceIamMember, and an error, if there is any.
func (c *FakeSpannerInstanceIamMembers) Update(spannerInstanceIamMember *v1alpha1.SpannerInstanceIamMember) (result *v1alpha1.SpannerInstanceIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(spannerinstanceiammembersResource, c.ns, spannerInstanceIamMember), &v1alpha1.SpannerInstanceIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SpannerInstanceIamMember), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSpannerInstanceIamMembers) UpdateStatus(spannerInstanceIamMember *v1alpha1.SpannerInstanceIamMember) (*v1alpha1.SpannerInstanceIamMember, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(spannerinstanceiammembersResource, "status", c.ns, spannerInstanceIamMember), &v1alpha1.SpannerInstanceIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SpannerInstanceIamMember), err
}

// Delete takes name of the spannerInstanceIamMember and deletes it. Returns an error if one occurs.
func (c *FakeSpannerInstanceIamMembers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(spannerinstanceiammembersResource, c.ns, name), &v1alpha1.SpannerInstanceIamMember{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSpannerInstanceIamMembers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(spannerinstanceiammembersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SpannerInstanceIamMemberList{})
	return err
}

// Patch applies the patch and returns the patched spannerInstanceIamMember.
func (c *FakeSpannerInstanceIamMembers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SpannerInstanceIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(spannerinstanceiammembersResource, c.ns, name, pt, data, subresources...), &v1alpha1.SpannerInstanceIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SpannerInstanceIamMember), err
}
