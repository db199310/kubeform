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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakeSentinelAlertRuleScheduleds implements SentinelAlertRuleScheduledInterface
type FakeSentinelAlertRuleScheduleds struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var sentinelalertrulescheduledsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "sentinelalertrulescheduleds"}

var sentinelalertrulescheduledsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "SentinelAlertRuleScheduled"}

// Get takes name of the sentinelAlertRuleScheduled, and returns the corresponding sentinelAlertRuleScheduled object, and an error if there is any.
func (c *FakeSentinelAlertRuleScheduleds) Get(name string, options v1.GetOptions) (result *v1alpha1.SentinelAlertRuleScheduled, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sentinelalertrulescheduledsResource, c.ns, name), &v1alpha1.SentinelAlertRuleScheduled{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SentinelAlertRuleScheduled), err
}

// List takes label and field selectors, and returns the list of SentinelAlertRuleScheduleds that match those selectors.
func (c *FakeSentinelAlertRuleScheduleds) List(opts v1.ListOptions) (result *v1alpha1.SentinelAlertRuleScheduledList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sentinelalertrulescheduledsResource, sentinelalertrulescheduledsKind, c.ns, opts), &v1alpha1.SentinelAlertRuleScheduledList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SentinelAlertRuleScheduledList{ListMeta: obj.(*v1alpha1.SentinelAlertRuleScheduledList).ListMeta}
	for _, item := range obj.(*v1alpha1.SentinelAlertRuleScheduledList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sentinelAlertRuleScheduleds.
func (c *FakeSentinelAlertRuleScheduleds) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sentinelalertrulescheduledsResource, c.ns, opts))

}

// Create takes the representation of a sentinelAlertRuleScheduled and creates it.  Returns the server's representation of the sentinelAlertRuleScheduled, and an error, if there is any.
func (c *FakeSentinelAlertRuleScheduleds) Create(sentinelAlertRuleScheduled *v1alpha1.SentinelAlertRuleScheduled) (result *v1alpha1.SentinelAlertRuleScheduled, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sentinelalertrulescheduledsResource, c.ns, sentinelAlertRuleScheduled), &v1alpha1.SentinelAlertRuleScheduled{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SentinelAlertRuleScheduled), err
}

// Update takes the representation of a sentinelAlertRuleScheduled and updates it. Returns the server's representation of the sentinelAlertRuleScheduled, and an error, if there is any.
func (c *FakeSentinelAlertRuleScheduleds) Update(sentinelAlertRuleScheduled *v1alpha1.SentinelAlertRuleScheduled) (result *v1alpha1.SentinelAlertRuleScheduled, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sentinelalertrulescheduledsResource, c.ns, sentinelAlertRuleScheduled), &v1alpha1.SentinelAlertRuleScheduled{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SentinelAlertRuleScheduled), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSentinelAlertRuleScheduleds) UpdateStatus(sentinelAlertRuleScheduled *v1alpha1.SentinelAlertRuleScheduled) (*v1alpha1.SentinelAlertRuleScheduled, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sentinelalertrulescheduledsResource, "status", c.ns, sentinelAlertRuleScheduled), &v1alpha1.SentinelAlertRuleScheduled{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SentinelAlertRuleScheduled), err
}

// Delete takes name of the sentinelAlertRuleScheduled and deletes it. Returns an error if one occurs.
func (c *FakeSentinelAlertRuleScheduleds) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sentinelalertrulescheduledsResource, c.ns, name), &v1alpha1.SentinelAlertRuleScheduled{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSentinelAlertRuleScheduleds) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sentinelalertrulescheduledsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SentinelAlertRuleScheduledList{})
	return err
}

// Patch applies the patch and returns the patched sentinelAlertRuleScheduled.
func (c *FakeSentinelAlertRuleScheduleds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SentinelAlertRuleScheduled, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sentinelalertrulescheduledsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SentinelAlertRuleScheduled{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SentinelAlertRuleScheduled), err
}
