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

// FakeCognitiveAccounts implements CognitiveAccountInterface
type FakeCognitiveAccounts struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var cognitiveaccountsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "cognitiveaccounts"}

var cognitiveaccountsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "CognitiveAccount"}

// Get takes name of the cognitiveAccount, and returns the corresponding cognitiveAccount object, and an error if there is any.
func (c *FakeCognitiveAccounts) Get(name string, options v1.GetOptions) (result *v1alpha1.CognitiveAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cognitiveaccountsResource, c.ns, name), &v1alpha1.CognitiveAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CognitiveAccount), err
}

// List takes label and field selectors, and returns the list of CognitiveAccounts that match those selectors.
func (c *FakeCognitiveAccounts) List(opts v1.ListOptions) (result *v1alpha1.CognitiveAccountList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cognitiveaccountsResource, cognitiveaccountsKind, c.ns, opts), &v1alpha1.CognitiveAccountList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CognitiveAccountList{ListMeta: obj.(*v1alpha1.CognitiveAccountList).ListMeta}
	for _, item := range obj.(*v1alpha1.CognitiveAccountList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cognitiveAccounts.
func (c *FakeCognitiveAccounts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cognitiveaccountsResource, c.ns, opts))

}

// Create takes the representation of a cognitiveAccount and creates it.  Returns the server's representation of the cognitiveAccount, and an error, if there is any.
func (c *FakeCognitiveAccounts) Create(cognitiveAccount *v1alpha1.CognitiveAccount) (result *v1alpha1.CognitiveAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cognitiveaccountsResource, c.ns, cognitiveAccount), &v1alpha1.CognitiveAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CognitiveAccount), err
}

// Update takes the representation of a cognitiveAccount and updates it. Returns the server's representation of the cognitiveAccount, and an error, if there is any.
func (c *FakeCognitiveAccounts) Update(cognitiveAccount *v1alpha1.CognitiveAccount) (result *v1alpha1.CognitiveAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cognitiveaccountsResource, c.ns, cognitiveAccount), &v1alpha1.CognitiveAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CognitiveAccount), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCognitiveAccounts) UpdateStatus(cognitiveAccount *v1alpha1.CognitiveAccount) (*v1alpha1.CognitiveAccount, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cognitiveaccountsResource, "status", c.ns, cognitiveAccount), &v1alpha1.CognitiveAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CognitiveAccount), err
}

// Delete takes name of the cognitiveAccount and deletes it. Returns an error if one occurs.
func (c *FakeCognitiveAccounts) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cognitiveaccountsResource, c.ns, name), &v1alpha1.CognitiveAccount{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCognitiveAccounts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cognitiveaccountsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CognitiveAccountList{})
	return err
}

// Patch applies the patch and returns the patched cognitiveAccount.
func (c *FakeCognitiveAccounts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CognitiveAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cognitiveaccountsResource, c.ns, name, pt, data, subresources...), &v1alpha1.CognitiveAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CognitiveAccount), err
}
