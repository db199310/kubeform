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

// FakeManagementGroups implements ManagementGroupInterface
type FakeManagementGroups struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var managementgroupsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "managementgroups"}

var managementgroupsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "ManagementGroup"}

// Get takes name of the managementGroup, and returns the corresponding managementGroup object, and an error if there is any.
func (c *FakeManagementGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ManagementGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(managementgroupsResource, c.ns, name), &v1alpha1.ManagementGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagementGroup), err
}

// List takes label and field selectors, and returns the list of ManagementGroups that match those selectors.
func (c *FakeManagementGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ManagementGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(managementgroupsResource, managementgroupsKind, c.ns, opts), &v1alpha1.ManagementGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ManagementGroupList{ListMeta: obj.(*v1alpha1.ManagementGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.ManagementGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested managementGroups.
func (c *FakeManagementGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(managementgroupsResource, c.ns, opts))

}

// Create takes the representation of a managementGroup and creates it.  Returns the server's representation of the managementGroup, and an error, if there is any.
func (c *FakeManagementGroups) Create(ctx context.Context, managementGroup *v1alpha1.ManagementGroup, opts v1.CreateOptions) (result *v1alpha1.ManagementGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(managementgroupsResource, c.ns, managementGroup), &v1alpha1.ManagementGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagementGroup), err
}

// Update takes the representation of a managementGroup and updates it. Returns the server's representation of the managementGroup, and an error, if there is any.
func (c *FakeManagementGroups) Update(ctx context.Context, managementGroup *v1alpha1.ManagementGroup, opts v1.UpdateOptions) (result *v1alpha1.ManagementGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(managementgroupsResource, c.ns, managementGroup), &v1alpha1.ManagementGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagementGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeManagementGroups) UpdateStatus(ctx context.Context, managementGroup *v1alpha1.ManagementGroup, opts v1.UpdateOptions) (*v1alpha1.ManagementGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(managementgroupsResource, "status", c.ns, managementGroup), &v1alpha1.ManagementGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagementGroup), err
}

// Delete takes name of the managementGroup and deletes it. Returns an error if one occurs.
func (c *FakeManagementGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(managementgroupsResource, c.ns, name), &v1alpha1.ManagementGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeManagementGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(managementgroupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ManagementGroupList{})
	return err
}

// Patch applies the patch and returns the patched managementGroup.
func (c *FakeManagementGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ManagementGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(managementgroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ManagementGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagementGroup), err
}
