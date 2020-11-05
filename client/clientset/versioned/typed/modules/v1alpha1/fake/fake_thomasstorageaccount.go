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
	v1alpha1 "kubeform.dev/kubeform/apis/modules/v1alpha1"
)

// FakeThomasStorageAccounts implements ThomasStorageAccountInterface
type FakeThomasStorageAccounts struct {
	Fake *FakeModulesV1alpha1
	ns   string
}

var thomasstorageaccountsResource = schema.GroupVersionResource{Group: "modules.kubeform.com", Version: "v1alpha1", Resource: "thomasstorageaccounts"}

var thomasstorageaccountsKind = schema.GroupVersionKind{Group: "modules.kubeform.com", Version: "v1alpha1", Kind: "ThomasStorageAccount"}

// Get takes name of the thomasStorageAccount, and returns the corresponding thomasStorageAccount object, and an error if there is any.
func (c *FakeThomasStorageAccounts) Get(name string, options v1.GetOptions) (result *v1alpha1.ThomasStorageAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(thomasstorageaccountsResource, c.ns, name), &v1alpha1.ThomasStorageAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ThomasStorageAccount), err
}

// List takes label and field selectors, and returns the list of ThomasStorageAccounts that match those selectors.
func (c *FakeThomasStorageAccounts) List(opts v1.ListOptions) (result *v1alpha1.ThomasStorageAccountList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(thomasstorageaccountsResource, thomasstorageaccountsKind, c.ns, opts), &v1alpha1.ThomasStorageAccountList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ThomasStorageAccountList{ListMeta: obj.(*v1alpha1.ThomasStorageAccountList).ListMeta}
	for _, item := range obj.(*v1alpha1.ThomasStorageAccountList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested thomasStorageAccounts.
func (c *FakeThomasStorageAccounts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(thomasstorageaccountsResource, c.ns, opts))

}

// Create takes the representation of a thomasStorageAccount and creates it.  Returns the server's representation of the thomasStorageAccount, and an error, if there is any.
func (c *FakeThomasStorageAccounts) Create(thomasStorageAccount *v1alpha1.ThomasStorageAccount) (result *v1alpha1.ThomasStorageAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(thomasstorageaccountsResource, c.ns, thomasStorageAccount), &v1alpha1.ThomasStorageAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ThomasStorageAccount), err
}

// Update takes the representation of a thomasStorageAccount and updates it. Returns the server's representation of the thomasStorageAccount, and an error, if there is any.
func (c *FakeThomasStorageAccounts) Update(thomasStorageAccount *v1alpha1.ThomasStorageAccount) (result *v1alpha1.ThomasStorageAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(thomasstorageaccountsResource, c.ns, thomasStorageAccount), &v1alpha1.ThomasStorageAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ThomasStorageAccount), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeThomasStorageAccounts) UpdateStatus(thomasStorageAccount *v1alpha1.ThomasStorageAccount) (*v1alpha1.ThomasStorageAccount, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(thomasstorageaccountsResource, "status", c.ns, thomasStorageAccount), &v1alpha1.ThomasStorageAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ThomasStorageAccount), err
}

// Delete takes name of the thomasStorageAccount and deletes it. Returns an error if one occurs.
func (c *FakeThomasStorageAccounts) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(thomasstorageaccountsResource, c.ns, name), &v1alpha1.ThomasStorageAccount{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeThomasStorageAccounts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(thomasstorageaccountsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ThomasStorageAccountList{})
	return err
}

// Patch applies the patch and returns the patched thomasStorageAccount.
func (c *FakeThomasStorageAccounts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ThomasStorageAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(thomasstorageaccountsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ThomasStorageAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ThomasStorageAccount), err
}
