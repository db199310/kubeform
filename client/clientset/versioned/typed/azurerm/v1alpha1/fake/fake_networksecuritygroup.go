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

// FakeNetworkSecurityGroups implements NetworkSecurityGroupInterface
type FakeNetworkSecurityGroups struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var networksecuritygroupsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "networksecuritygroups"}

var networksecuritygroupsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "NetworkSecurityGroup"}

// Get takes name of the networkSecurityGroup, and returns the corresponding networkSecurityGroup object, and an error if there is any.
func (c *FakeNetworkSecurityGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NetworkSecurityGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networksecuritygroupsResource, c.ns, name), &v1alpha1.NetworkSecurityGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkSecurityGroup), err
}

// List takes label and field selectors, and returns the list of NetworkSecurityGroups that match those selectors.
func (c *FakeNetworkSecurityGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NetworkSecurityGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networksecuritygroupsResource, networksecuritygroupsKind, c.ns, opts), &v1alpha1.NetworkSecurityGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NetworkSecurityGroupList{ListMeta: obj.(*v1alpha1.NetworkSecurityGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.NetworkSecurityGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkSecurityGroups.
func (c *FakeNetworkSecurityGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networksecuritygroupsResource, c.ns, opts))

}

// Create takes the representation of a networkSecurityGroup and creates it.  Returns the server's representation of the networkSecurityGroup, and an error, if there is any.
func (c *FakeNetworkSecurityGroups) Create(ctx context.Context, networkSecurityGroup *v1alpha1.NetworkSecurityGroup, opts v1.CreateOptions) (result *v1alpha1.NetworkSecurityGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networksecuritygroupsResource, c.ns, networkSecurityGroup), &v1alpha1.NetworkSecurityGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkSecurityGroup), err
}

// Update takes the representation of a networkSecurityGroup and updates it. Returns the server's representation of the networkSecurityGroup, and an error, if there is any.
func (c *FakeNetworkSecurityGroups) Update(ctx context.Context, networkSecurityGroup *v1alpha1.NetworkSecurityGroup, opts v1.UpdateOptions) (result *v1alpha1.NetworkSecurityGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networksecuritygroupsResource, c.ns, networkSecurityGroup), &v1alpha1.NetworkSecurityGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkSecurityGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkSecurityGroups) UpdateStatus(ctx context.Context, networkSecurityGroup *v1alpha1.NetworkSecurityGroup, opts v1.UpdateOptions) (*v1alpha1.NetworkSecurityGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(networksecuritygroupsResource, "status", c.ns, networkSecurityGroup), &v1alpha1.NetworkSecurityGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkSecurityGroup), err
}

// Delete takes name of the networkSecurityGroup and deletes it. Returns an error if one occurs.
func (c *FakeNetworkSecurityGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(networksecuritygroupsResource, c.ns, name), &v1alpha1.NetworkSecurityGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkSecurityGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networksecuritygroupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NetworkSecurityGroupList{})
	return err
}

// Patch applies the patch and returns the patched networkSecurityGroup.
func (c *FakeNetworkSecurityGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NetworkSecurityGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networksecuritygroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.NetworkSecurityGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkSecurityGroup), err
}
