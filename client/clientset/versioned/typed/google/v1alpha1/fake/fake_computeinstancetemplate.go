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

// FakeComputeInstanceTemplates implements ComputeInstanceTemplateInterface
type FakeComputeInstanceTemplates struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var computeinstancetemplatesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "computeinstancetemplates"}

var computeinstancetemplatesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "ComputeInstanceTemplate"}

// Get takes name of the computeInstanceTemplate, and returns the corresponding computeInstanceTemplate object, and an error if there is any.
func (c *FakeComputeInstanceTemplates) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeInstanceTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computeinstancetemplatesResource, c.ns, name), &v1alpha1.ComputeInstanceTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeInstanceTemplate), err
}

// List takes label and field selectors, and returns the list of ComputeInstanceTemplates that match those selectors.
func (c *FakeComputeInstanceTemplates) List(opts v1.ListOptions) (result *v1alpha1.ComputeInstanceTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computeinstancetemplatesResource, computeinstancetemplatesKind, c.ns, opts), &v1alpha1.ComputeInstanceTemplateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComputeInstanceTemplateList{ListMeta: obj.(*v1alpha1.ComputeInstanceTemplateList).ListMeta}
	for _, item := range obj.(*v1alpha1.ComputeInstanceTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeInstanceTemplates.
func (c *FakeComputeInstanceTemplates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computeinstancetemplatesResource, c.ns, opts))

}

// Create takes the representation of a computeInstanceTemplate and creates it.  Returns the server's representation of the computeInstanceTemplate, and an error, if there is any.
func (c *FakeComputeInstanceTemplates) Create(computeInstanceTemplate *v1alpha1.ComputeInstanceTemplate) (result *v1alpha1.ComputeInstanceTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computeinstancetemplatesResource, c.ns, computeInstanceTemplate), &v1alpha1.ComputeInstanceTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeInstanceTemplate), err
}

// Update takes the representation of a computeInstanceTemplate and updates it. Returns the server's representation of the computeInstanceTemplate, and an error, if there is any.
func (c *FakeComputeInstanceTemplates) Update(computeInstanceTemplate *v1alpha1.ComputeInstanceTemplate) (result *v1alpha1.ComputeInstanceTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computeinstancetemplatesResource, c.ns, computeInstanceTemplate), &v1alpha1.ComputeInstanceTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeInstanceTemplate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeInstanceTemplates) UpdateStatus(computeInstanceTemplate *v1alpha1.ComputeInstanceTemplate) (*v1alpha1.ComputeInstanceTemplate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computeinstancetemplatesResource, "status", c.ns, computeInstanceTemplate), &v1alpha1.ComputeInstanceTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeInstanceTemplate), err
}

// Delete takes name of the computeInstanceTemplate and deletes it. Returns an error if one occurs.
func (c *FakeComputeInstanceTemplates) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(computeinstancetemplatesResource, c.ns, name), &v1alpha1.ComputeInstanceTemplate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeInstanceTemplates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computeinstancetemplatesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComputeInstanceTemplateList{})
	return err
}

// Patch applies the patch and returns the patched computeInstanceTemplate.
func (c *FakeComputeInstanceTemplates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeInstanceTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computeinstancetemplatesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ComputeInstanceTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeInstanceTemplate), err
}
