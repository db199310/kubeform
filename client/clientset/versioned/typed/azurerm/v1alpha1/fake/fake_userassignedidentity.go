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

// FakeUserAssignedIdentities implements UserAssignedIdentityInterface
type FakeUserAssignedIdentities struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var userassignedidentitiesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "userassignedidentities"}

var userassignedidentitiesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "UserAssignedIdentity"}

// Get takes name of the userAssignedIdentity, and returns the corresponding userAssignedIdentity object, and an error if there is any.
func (c *FakeUserAssignedIdentities) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.UserAssignedIdentity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(userassignedidentitiesResource, c.ns, name), &v1alpha1.UserAssignedIdentity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserAssignedIdentity), err
}

// List takes label and field selectors, and returns the list of UserAssignedIdentities that match those selectors.
func (c *FakeUserAssignedIdentities) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.UserAssignedIdentityList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(userassignedidentitiesResource, userassignedidentitiesKind, c.ns, opts), &v1alpha1.UserAssignedIdentityList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.UserAssignedIdentityList{ListMeta: obj.(*v1alpha1.UserAssignedIdentityList).ListMeta}
	for _, item := range obj.(*v1alpha1.UserAssignedIdentityList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested userAssignedIdentities.
func (c *FakeUserAssignedIdentities) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(userassignedidentitiesResource, c.ns, opts))

}

// Create takes the representation of a userAssignedIdentity and creates it.  Returns the server's representation of the userAssignedIdentity, and an error, if there is any.
func (c *FakeUserAssignedIdentities) Create(ctx context.Context, userAssignedIdentity *v1alpha1.UserAssignedIdentity, opts v1.CreateOptions) (result *v1alpha1.UserAssignedIdentity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(userassignedidentitiesResource, c.ns, userAssignedIdentity), &v1alpha1.UserAssignedIdentity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserAssignedIdentity), err
}

// Update takes the representation of a userAssignedIdentity and updates it. Returns the server's representation of the userAssignedIdentity, and an error, if there is any.
func (c *FakeUserAssignedIdentities) Update(ctx context.Context, userAssignedIdentity *v1alpha1.UserAssignedIdentity, opts v1.UpdateOptions) (result *v1alpha1.UserAssignedIdentity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(userassignedidentitiesResource, c.ns, userAssignedIdentity), &v1alpha1.UserAssignedIdentity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserAssignedIdentity), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeUserAssignedIdentities) UpdateStatus(ctx context.Context, userAssignedIdentity *v1alpha1.UserAssignedIdentity, opts v1.UpdateOptions) (*v1alpha1.UserAssignedIdentity, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(userassignedidentitiesResource, "status", c.ns, userAssignedIdentity), &v1alpha1.UserAssignedIdentity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserAssignedIdentity), err
}

// Delete takes name of the userAssignedIdentity and deletes it. Returns an error if one occurs.
func (c *FakeUserAssignedIdentities) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(userassignedidentitiesResource, c.ns, name), &v1alpha1.UserAssignedIdentity{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeUserAssignedIdentities) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(userassignedidentitiesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.UserAssignedIdentityList{})
	return err
}

// Patch applies the patch and returns the patched userAssignedIdentity.
func (c *FakeUserAssignedIdentities) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.UserAssignedIdentity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(userassignedidentitiesResource, c.ns, name, pt, data, subresources...), &v1alpha1.UserAssignedIdentity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserAssignedIdentity), err
}
