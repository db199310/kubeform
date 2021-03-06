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
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakeAutomationJobSchedules implements AutomationJobScheduleInterface
type FakeAutomationJobSchedules struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var automationjobschedulesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "automationjobschedules"}

var automationjobschedulesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AutomationJobSchedule"}

// Get takes name of the automationJobSchedule, and returns the corresponding automationJobSchedule object, and an error if there is any.
func (c *FakeAutomationJobSchedules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AutomationJobSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(automationjobschedulesResource, c.ns, name), &v1alpha1.AutomationJobSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutomationJobSchedule), err
}

// List takes label and field selectors, and returns the list of AutomationJobSchedules that match those selectors.
func (c *FakeAutomationJobSchedules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AutomationJobScheduleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(automationjobschedulesResource, automationjobschedulesKind, c.ns, opts), &v1alpha1.AutomationJobScheduleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AutomationJobScheduleList{ListMeta: obj.(*v1alpha1.AutomationJobScheduleList).ListMeta}
	for _, item := range obj.(*v1alpha1.AutomationJobScheduleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested automationJobSchedules.
func (c *FakeAutomationJobSchedules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(automationjobschedulesResource, c.ns, opts))

}

// Create takes the representation of a automationJobSchedule and creates it.  Returns the server's representation of the automationJobSchedule, and an error, if there is any.
func (c *FakeAutomationJobSchedules) Create(ctx context.Context, automationJobSchedule *v1alpha1.AutomationJobSchedule, opts v1.CreateOptions) (result *v1alpha1.AutomationJobSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(automationjobschedulesResource, c.ns, automationJobSchedule), &v1alpha1.AutomationJobSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutomationJobSchedule), err
}

// Update takes the representation of a automationJobSchedule and updates it. Returns the server's representation of the automationJobSchedule, and an error, if there is any.
func (c *FakeAutomationJobSchedules) Update(ctx context.Context, automationJobSchedule *v1alpha1.AutomationJobSchedule, opts v1.UpdateOptions) (result *v1alpha1.AutomationJobSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(automationjobschedulesResource, c.ns, automationJobSchedule), &v1alpha1.AutomationJobSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutomationJobSchedule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAutomationJobSchedules) UpdateStatus(ctx context.Context, automationJobSchedule *v1alpha1.AutomationJobSchedule, opts v1.UpdateOptions) (*v1alpha1.AutomationJobSchedule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(automationjobschedulesResource, "status", c.ns, automationJobSchedule), &v1alpha1.AutomationJobSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutomationJobSchedule), err
}

// Delete takes name of the automationJobSchedule and deletes it. Returns an error if one occurs.
func (c *FakeAutomationJobSchedules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(automationjobschedulesResource, c.ns, name), &v1alpha1.AutomationJobSchedule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAutomationJobSchedules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(automationjobschedulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AutomationJobScheduleList{})
	return err
}

// Patch applies the patch and returns the patched automationJobSchedule.
func (c *FakeAutomationJobSchedules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AutomationJobSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(automationjobschedulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AutomationJobSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutomationJobSchedule), err
}
