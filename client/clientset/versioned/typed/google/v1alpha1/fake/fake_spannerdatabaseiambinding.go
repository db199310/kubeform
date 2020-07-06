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

// FakeSpannerDatabaseIamBindings implements SpannerDatabaseIamBindingInterface
type FakeSpannerDatabaseIamBindings struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var spannerdatabaseiambindingsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "spannerdatabaseiambindings"}

var spannerdatabaseiambindingsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "SpannerDatabaseIamBinding"}

// Get takes name of the spannerDatabaseIamBinding, and returns the corresponding spannerDatabaseIamBinding object, and an error if there is any.
func (c *FakeSpannerDatabaseIamBindings) Get(name string, options v1.GetOptions) (result *v1alpha1.SpannerDatabaseIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(spannerdatabaseiambindingsResource, c.ns, name), &v1alpha1.SpannerDatabaseIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SpannerDatabaseIamBinding), err
}

// List takes label and field selectors, and returns the list of SpannerDatabaseIamBindings that match those selectors.
func (c *FakeSpannerDatabaseIamBindings) List(opts v1.ListOptions) (result *v1alpha1.SpannerDatabaseIamBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(spannerdatabaseiambindingsResource, spannerdatabaseiambindingsKind, c.ns, opts), &v1alpha1.SpannerDatabaseIamBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SpannerDatabaseIamBindingList{ListMeta: obj.(*v1alpha1.SpannerDatabaseIamBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.SpannerDatabaseIamBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested spannerDatabaseIamBindings.
func (c *FakeSpannerDatabaseIamBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(spannerdatabaseiambindingsResource, c.ns, opts))

}

// Create takes the representation of a spannerDatabaseIamBinding and creates it.  Returns the server's representation of the spannerDatabaseIamBinding, and an error, if there is any.
func (c *FakeSpannerDatabaseIamBindings) Create(spannerDatabaseIamBinding *v1alpha1.SpannerDatabaseIamBinding) (result *v1alpha1.SpannerDatabaseIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(spannerdatabaseiambindingsResource, c.ns, spannerDatabaseIamBinding), &v1alpha1.SpannerDatabaseIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SpannerDatabaseIamBinding), err
}

// Update takes the representation of a spannerDatabaseIamBinding and updates it. Returns the server's representation of the spannerDatabaseIamBinding, and an error, if there is any.
func (c *FakeSpannerDatabaseIamBindings) Update(spannerDatabaseIamBinding *v1alpha1.SpannerDatabaseIamBinding) (result *v1alpha1.SpannerDatabaseIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(spannerdatabaseiambindingsResource, c.ns, spannerDatabaseIamBinding), &v1alpha1.SpannerDatabaseIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SpannerDatabaseIamBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSpannerDatabaseIamBindings) UpdateStatus(spannerDatabaseIamBinding *v1alpha1.SpannerDatabaseIamBinding) (*v1alpha1.SpannerDatabaseIamBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(spannerdatabaseiambindingsResource, "status", c.ns, spannerDatabaseIamBinding), &v1alpha1.SpannerDatabaseIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SpannerDatabaseIamBinding), err
}

// Delete takes name of the spannerDatabaseIamBinding and deletes it. Returns an error if one occurs.
func (c *FakeSpannerDatabaseIamBindings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(spannerdatabaseiambindingsResource, c.ns, name), &v1alpha1.SpannerDatabaseIamBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSpannerDatabaseIamBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(spannerdatabaseiambindingsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SpannerDatabaseIamBindingList{})
	return err
}

// Patch applies the patch and returns the patched spannerDatabaseIamBinding.
func (c *FakeSpannerDatabaseIamBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SpannerDatabaseIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(spannerdatabaseiambindingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SpannerDatabaseIamBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SpannerDatabaseIamBinding), err
}
