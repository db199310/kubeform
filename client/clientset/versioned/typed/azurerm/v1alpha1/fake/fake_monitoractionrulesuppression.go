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

// FakeMonitorActionRuleSuppressions implements MonitorActionRuleSuppressionInterface
type FakeMonitorActionRuleSuppressions struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var monitoractionrulesuppressionsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "monitoractionrulesuppressions"}

var monitoractionrulesuppressionsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "MonitorActionRuleSuppression"}

// Get takes name of the monitorActionRuleSuppression, and returns the corresponding monitorActionRuleSuppression object, and an error if there is any.
func (c *FakeMonitorActionRuleSuppressions) Get(name string, options v1.GetOptions) (result *v1alpha1.MonitorActionRuleSuppression, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(monitoractionrulesuppressionsResource, c.ns, name), &v1alpha1.MonitorActionRuleSuppression{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitorActionRuleSuppression), err
}

// List takes label and field selectors, and returns the list of MonitorActionRuleSuppressions that match those selectors.
func (c *FakeMonitorActionRuleSuppressions) List(opts v1.ListOptions) (result *v1alpha1.MonitorActionRuleSuppressionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(monitoractionrulesuppressionsResource, monitoractionrulesuppressionsKind, c.ns, opts), &v1alpha1.MonitorActionRuleSuppressionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MonitorActionRuleSuppressionList{ListMeta: obj.(*v1alpha1.MonitorActionRuleSuppressionList).ListMeta}
	for _, item := range obj.(*v1alpha1.MonitorActionRuleSuppressionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested monitorActionRuleSuppressions.
func (c *FakeMonitorActionRuleSuppressions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(monitoractionrulesuppressionsResource, c.ns, opts))

}

// Create takes the representation of a monitorActionRuleSuppression and creates it.  Returns the server's representation of the monitorActionRuleSuppression, and an error, if there is any.
func (c *FakeMonitorActionRuleSuppressions) Create(monitorActionRuleSuppression *v1alpha1.MonitorActionRuleSuppression) (result *v1alpha1.MonitorActionRuleSuppression, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(monitoractionrulesuppressionsResource, c.ns, monitorActionRuleSuppression), &v1alpha1.MonitorActionRuleSuppression{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitorActionRuleSuppression), err
}

// Update takes the representation of a monitorActionRuleSuppression and updates it. Returns the server's representation of the monitorActionRuleSuppression, and an error, if there is any.
func (c *FakeMonitorActionRuleSuppressions) Update(monitorActionRuleSuppression *v1alpha1.MonitorActionRuleSuppression) (result *v1alpha1.MonitorActionRuleSuppression, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(monitoractionrulesuppressionsResource, c.ns, monitorActionRuleSuppression), &v1alpha1.MonitorActionRuleSuppression{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitorActionRuleSuppression), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMonitorActionRuleSuppressions) UpdateStatus(monitorActionRuleSuppression *v1alpha1.MonitorActionRuleSuppression) (*v1alpha1.MonitorActionRuleSuppression, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(monitoractionrulesuppressionsResource, "status", c.ns, monitorActionRuleSuppression), &v1alpha1.MonitorActionRuleSuppression{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitorActionRuleSuppression), err
}

// Delete takes name of the monitorActionRuleSuppression and deletes it. Returns an error if one occurs.
func (c *FakeMonitorActionRuleSuppressions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(monitoractionrulesuppressionsResource, c.ns, name), &v1alpha1.MonitorActionRuleSuppression{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMonitorActionRuleSuppressions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(monitoractionrulesuppressionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.MonitorActionRuleSuppressionList{})
	return err
}

// Patch applies the patch and returns the patched monitorActionRuleSuppression.
func (c *FakeMonitorActionRuleSuppressions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MonitorActionRuleSuppression, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(monitoractionrulesuppressionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.MonitorActionRuleSuppression{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitorActionRuleSuppression), err
}