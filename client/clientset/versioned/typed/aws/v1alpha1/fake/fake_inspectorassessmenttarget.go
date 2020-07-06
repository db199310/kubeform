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

// FakeInspectorAssessmentTargets implements InspectorAssessmentTargetInterface
type FakeInspectorAssessmentTargets struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var inspectorassessmenttargetsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "inspectorassessmenttargets"}

var inspectorassessmenttargetsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "InspectorAssessmentTarget"}

// Get takes name of the inspectorAssessmentTarget, and returns the corresponding inspectorAssessmentTarget object, and an error if there is any.
func (c *FakeInspectorAssessmentTargets) Get(name string, options v1.GetOptions) (result *v1alpha1.InspectorAssessmentTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(inspectorassessmenttargetsResource, c.ns, name), &v1alpha1.InspectorAssessmentTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InspectorAssessmentTarget), err
}

// List takes label and field selectors, and returns the list of InspectorAssessmentTargets that match those selectors.
func (c *FakeInspectorAssessmentTargets) List(opts v1.ListOptions) (result *v1alpha1.InspectorAssessmentTargetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(inspectorassessmenttargetsResource, inspectorassessmenttargetsKind, c.ns, opts), &v1alpha1.InspectorAssessmentTargetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.InspectorAssessmentTargetList{ListMeta: obj.(*v1alpha1.InspectorAssessmentTargetList).ListMeta}
	for _, item := range obj.(*v1alpha1.InspectorAssessmentTargetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested inspectorAssessmentTargets.
func (c *FakeInspectorAssessmentTargets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(inspectorassessmenttargetsResource, c.ns, opts))

}

// Create takes the representation of a inspectorAssessmentTarget and creates it.  Returns the server's representation of the inspectorAssessmentTarget, and an error, if there is any.
func (c *FakeInspectorAssessmentTargets) Create(inspectorAssessmentTarget *v1alpha1.InspectorAssessmentTarget) (result *v1alpha1.InspectorAssessmentTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(inspectorassessmenttargetsResource, c.ns, inspectorAssessmentTarget), &v1alpha1.InspectorAssessmentTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InspectorAssessmentTarget), err
}

// Update takes the representation of a inspectorAssessmentTarget and updates it. Returns the server's representation of the inspectorAssessmentTarget, and an error, if there is any.
func (c *FakeInspectorAssessmentTargets) Update(inspectorAssessmentTarget *v1alpha1.InspectorAssessmentTarget) (result *v1alpha1.InspectorAssessmentTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(inspectorassessmenttargetsResource, c.ns, inspectorAssessmentTarget), &v1alpha1.InspectorAssessmentTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InspectorAssessmentTarget), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeInspectorAssessmentTargets) UpdateStatus(inspectorAssessmentTarget *v1alpha1.InspectorAssessmentTarget) (*v1alpha1.InspectorAssessmentTarget, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(inspectorassessmenttargetsResource, "status", c.ns, inspectorAssessmentTarget), &v1alpha1.InspectorAssessmentTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InspectorAssessmentTarget), err
}

// Delete takes name of the inspectorAssessmentTarget and deletes it. Returns an error if one occurs.
func (c *FakeInspectorAssessmentTargets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(inspectorassessmenttargetsResource, c.ns, name), &v1alpha1.InspectorAssessmentTarget{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeInspectorAssessmentTargets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(inspectorassessmenttargetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.InspectorAssessmentTargetList{})
	return err
}

// Patch applies the patch and returns the patched inspectorAssessmentTarget.
func (c *FakeInspectorAssessmentTargets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.InspectorAssessmentTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(inspectorassessmenttargetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.InspectorAssessmentTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InspectorAssessmentTarget), err
}
