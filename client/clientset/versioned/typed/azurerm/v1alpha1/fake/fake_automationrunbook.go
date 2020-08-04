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

// FakeAutomationRunbooks implements AutomationRunbookInterface
type FakeAutomationRunbooks struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var automationrunbooksResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "automationrunbooks"}

var automationrunbooksKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AutomationRunbook"}

// Get takes name of the automationRunbook, and returns the corresponding automationRunbook object, and an error if there is any.
func (c *FakeAutomationRunbooks) Get(name string, options v1.GetOptions) (result *v1alpha1.AutomationRunbook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(automationrunbooksResource, c.ns, name), &v1alpha1.AutomationRunbook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutomationRunbook), err
}

// List takes label and field selectors, and returns the list of AutomationRunbooks that match those selectors.
func (c *FakeAutomationRunbooks) List(opts v1.ListOptions) (result *v1alpha1.AutomationRunbookList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(automationrunbooksResource, automationrunbooksKind, c.ns, opts), &v1alpha1.AutomationRunbookList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AutomationRunbookList{ListMeta: obj.(*v1alpha1.AutomationRunbookList).ListMeta}
	for _, item := range obj.(*v1alpha1.AutomationRunbookList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested automationRunbooks.
func (c *FakeAutomationRunbooks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(automationrunbooksResource, c.ns, opts))

}

// Create takes the representation of a automationRunbook and creates it.  Returns the server's representation of the automationRunbook, and an error, if there is any.
func (c *FakeAutomationRunbooks) Create(automationRunbook *v1alpha1.AutomationRunbook) (result *v1alpha1.AutomationRunbook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(automationrunbooksResource, c.ns, automationRunbook), &v1alpha1.AutomationRunbook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutomationRunbook), err
}

// Update takes the representation of a automationRunbook and updates it. Returns the server's representation of the automationRunbook, and an error, if there is any.
func (c *FakeAutomationRunbooks) Update(automationRunbook *v1alpha1.AutomationRunbook) (result *v1alpha1.AutomationRunbook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(automationrunbooksResource, c.ns, automationRunbook), &v1alpha1.AutomationRunbook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutomationRunbook), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAutomationRunbooks) UpdateStatus(automationRunbook *v1alpha1.AutomationRunbook) (*v1alpha1.AutomationRunbook, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(automationrunbooksResource, "status", c.ns, automationRunbook), &v1alpha1.AutomationRunbook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutomationRunbook), err
}

// Delete takes name of the automationRunbook and deletes it. Returns an error if one occurs.
func (c *FakeAutomationRunbooks) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(automationrunbooksResource, c.ns, name), &v1alpha1.AutomationRunbook{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAutomationRunbooks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(automationrunbooksResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AutomationRunbookList{})
	return err
}

// Patch applies the patch and returns the patched automationRunbook.
func (c *FakeAutomationRunbooks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AutomationRunbook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(automationrunbooksResource, c.ns, name, pt, data, subresources...), &v1alpha1.AutomationRunbook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutomationRunbook), err
}
