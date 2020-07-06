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
	v1alpha1 "kubeform.dev/kubeform/apis/digitalocean/v1alpha1"
)

// FakeFloatingIPAssignments implements FloatingIPAssignmentInterface
type FakeFloatingIPAssignments struct {
	Fake *FakeDigitaloceanV1alpha1
	ns   string
}

var floatingipassignmentsResource = schema.GroupVersionResource{Group: "digitalocean.kubeform.com", Version: "v1alpha1", Resource: "floatingipassignments"}

var floatingipassignmentsKind = schema.GroupVersionKind{Group: "digitalocean.kubeform.com", Version: "v1alpha1", Kind: "FloatingIPAssignment"}

// Get takes name of the floatingIPAssignment, and returns the corresponding floatingIPAssignment object, and an error if there is any.
func (c *FakeFloatingIPAssignments) Get(name string, options v1.GetOptions) (result *v1alpha1.FloatingIPAssignment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(floatingipassignmentsResource, c.ns, name), &v1alpha1.FloatingIPAssignment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FloatingIPAssignment), err
}

// List takes label and field selectors, and returns the list of FloatingIPAssignments that match those selectors.
func (c *FakeFloatingIPAssignments) List(opts v1.ListOptions) (result *v1alpha1.FloatingIPAssignmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(floatingipassignmentsResource, floatingipassignmentsKind, c.ns, opts), &v1alpha1.FloatingIPAssignmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FloatingIPAssignmentList{ListMeta: obj.(*v1alpha1.FloatingIPAssignmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.FloatingIPAssignmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested floatingIPAssignments.
func (c *FakeFloatingIPAssignments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(floatingipassignmentsResource, c.ns, opts))

}

// Create takes the representation of a floatingIPAssignment and creates it.  Returns the server's representation of the floatingIPAssignment, and an error, if there is any.
func (c *FakeFloatingIPAssignments) Create(floatingIPAssignment *v1alpha1.FloatingIPAssignment) (result *v1alpha1.FloatingIPAssignment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(floatingipassignmentsResource, c.ns, floatingIPAssignment), &v1alpha1.FloatingIPAssignment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FloatingIPAssignment), err
}

// Update takes the representation of a floatingIPAssignment and updates it. Returns the server's representation of the floatingIPAssignment, and an error, if there is any.
func (c *FakeFloatingIPAssignments) Update(floatingIPAssignment *v1alpha1.FloatingIPAssignment) (result *v1alpha1.FloatingIPAssignment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(floatingipassignmentsResource, c.ns, floatingIPAssignment), &v1alpha1.FloatingIPAssignment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FloatingIPAssignment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFloatingIPAssignments) UpdateStatus(floatingIPAssignment *v1alpha1.FloatingIPAssignment) (*v1alpha1.FloatingIPAssignment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(floatingipassignmentsResource, "status", c.ns, floatingIPAssignment), &v1alpha1.FloatingIPAssignment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FloatingIPAssignment), err
}

// Delete takes name of the floatingIPAssignment and deletes it. Returns an error if one occurs.
func (c *FakeFloatingIPAssignments) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(floatingipassignmentsResource, c.ns, name), &v1alpha1.FloatingIPAssignment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFloatingIPAssignments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(floatingipassignmentsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.FloatingIPAssignmentList{})
	return err
}

// Patch applies the patch and returns the patched floatingIPAssignment.
func (c *FakeFloatingIPAssignments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FloatingIPAssignment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(floatingipassignmentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.FloatingIPAssignment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FloatingIPAssignment), err
}
