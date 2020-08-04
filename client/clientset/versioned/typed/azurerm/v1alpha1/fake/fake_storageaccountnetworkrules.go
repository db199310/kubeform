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

// FakeStorageAccountNetworkRuleses implements StorageAccountNetworkRulesInterface
type FakeStorageAccountNetworkRuleses struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var storageaccountnetworkrulesesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "storageaccountnetworkruleses"}

var storageaccountnetworkrulesesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "StorageAccountNetworkRules"}

// Get takes name of the storageAccountNetworkRules, and returns the corresponding storageAccountNetworkRules object, and an error if there is any.
func (c *FakeStorageAccountNetworkRuleses) Get(name string, options v1.GetOptions) (result *v1alpha1.StorageAccountNetworkRules, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(storageaccountnetworkrulesesResource, c.ns, name), &v1alpha1.StorageAccountNetworkRules{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageAccountNetworkRules), err
}

// List takes label and field selectors, and returns the list of StorageAccountNetworkRuleses that match those selectors.
func (c *FakeStorageAccountNetworkRuleses) List(opts v1.ListOptions) (result *v1alpha1.StorageAccountNetworkRulesList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(storageaccountnetworkrulesesResource, storageaccountnetworkrulesesKind, c.ns, opts), &v1alpha1.StorageAccountNetworkRulesList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StorageAccountNetworkRulesList{ListMeta: obj.(*v1alpha1.StorageAccountNetworkRulesList).ListMeta}
	for _, item := range obj.(*v1alpha1.StorageAccountNetworkRulesList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested storageAccountNetworkRuleses.
func (c *FakeStorageAccountNetworkRuleses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(storageaccountnetworkrulesesResource, c.ns, opts))

}

// Create takes the representation of a storageAccountNetworkRules and creates it.  Returns the server's representation of the storageAccountNetworkRules, and an error, if there is any.
func (c *FakeStorageAccountNetworkRuleses) Create(storageAccountNetworkRules *v1alpha1.StorageAccountNetworkRules) (result *v1alpha1.StorageAccountNetworkRules, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(storageaccountnetworkrulesesResource, c.ns, storageAccountNetworkRules), &v1alpha1.StorageAccountNetworkRules{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageAccountNetworkRules), err
}

// Update takes the representation of a storageAccountNetworkRules and updates it. Returns the server's representation of the storageAccountNetworkRules, and an error, if there is any.
func (c *FakeStorageAccountNetworkRuleses) Update(storageAccountNetworkRules *v1alpha1.StorageAccountNetworkRules) (result *v1alpha1.StorageAccountNetworkRules, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(storageaccountnetworkrulesesResource, c.ns, storageAccountNetworkRules), &v1alpha1.StorageAccountNetworkRules{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageAccountNetworkRules), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStorageAccountNetworkRuleses) UpdateStatus(storageAccountNetworkRules *v1alpha1.StorageAccountNetworkRules) (*v1alpha1.StorageAccountNetworkRules, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(storageaccountnetworkrulesesResource, "status", c.ns, storageAccountNetworkRules), &v1alpha1.StorageAccountNetworkRules{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageAccountNetworkRules), err
}

// Delete takes name of the storageAccountNetworkRules and deletes it. Returns an error if one occurs.
func (c *FakeStorageAccountNetworkRuleses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(storageaccountnetworkrulesesResource, c.ns, name), &v1alpha1.StorageAccountNetworkRules{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStorageAccountNetworkRuleses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(storageaccountnetworkrulesesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.StorageAccountNetworkRulesList{})
	return err
}

// Patch applies the patch and returns the patched storageAccountNetworkRules.
func (c *FakeStorageAccountNetworkRuleses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StorageAccountNetworkRules, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(storageaccountnetworkrulesesResource, c.ns, name, pt, data, subresources...), &v1alpha1.StorageAccountNetworkRules{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageAccountNetworkRules), err
}
