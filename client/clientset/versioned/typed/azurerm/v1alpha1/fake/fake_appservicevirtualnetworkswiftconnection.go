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

// FakeAppServiceVirtualNetworkSwiftConnections implements AppServiceVirtualNetworkSwiftConnectionInterface
type FakeAppServiceVirtualNetworkSwiftConnections struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var appservicevirtualnetworkswiftconnectionsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "appservicevirtualnetworkswiftconnections"}

var appservicevirtualnetworkswiftconnectionsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AppServiceVirtualNetworkSwiftConnection"}

// Get takes name of the appServiceVirtualNetworkSwiftConnection, and returns the corresponding appServiceVirtualNetworkSwiftConnection object, and an error if there is any.
func (c *FakeAppServiceVirtualNetworkSwiftConnections) Get(name string, options v1.GetOptions) (result *v1alpha1.AppServiceVirtualNetworkSwiftConnection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(appservicevirtualnetworkswiftconnectionsResource, c.ns, name), &v1alpha1.AppServiceVirtualNetworkSwiftConnection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppServiceVirtualNetworkSwiftConnection), err
}

// List takes label and field selectors, and returns the list of AppServiceVirtualNetworkSwiftConnections that match those selectors.
func (c *FakeAppServiceVirtualNetworkSwiftConnections) List(opts v1.ListOptions) (result *v1alpha1.AppServiceVirtualNetworkSwiftConnectionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(appservicevirtualnetworkswiftconnectionsResource, appservicevirtualnetworkswiftconnectionsKind, c.ns, opts), &v1alpha1.AppServiceVirtualNetworkSwiftConnectionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AppServiceVirtualNetworkSwiftConnectionList{ListMeta: obj.(*v1alpha1.AppServiceVirtualNetworkSwiftConnectionList).ListMeta}
	for _, item := range obj.(*v1alpha1.AppServiceVirtualNetworkSwiftConnectionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested appServiceVirtualNetworkSwiftConnections.
func (c *FakeAppServiceVirtualNetworkSwiftConnections) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(appservicevirtualnetworkswiftconnectionsResource, c.ns, opts))

}

// Create takes the representation of a appServiceVirtualNetworkSwiftConnection and creates it.  Returns the server's representation of the appServiceVirtualNetworkSwiftConnection, and an error, if there is any.
func (c *FakeAppServiceVirtualNetworkSwiftConnections) Create(appServiceVirtualNetworkSwiftConnection *v1alpha1.AppServiceVirtualNetworkSwiftConnection) (result *v1alpha1.AppServiceVirtualNetworkSwiftConnection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(appservicevirtualnetworkswiftconnectionsResource, c.ns, appServiceVirtualNetworkSwiftConnection), &v1alpha1.AppServiceVirtualNetworkSwiftConnection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppServiceVirtualNetworkSwiftConnection), err
}

// Update takes the representation of a appServiceVirtualNetworkSwiftConnection and updates it. Returns the server's representation of the appServiceVirtualNetworkSwiftConnection, and an error, if there is any.
func (c *FakeAppServiceVirtualNetworkSwiftConnections) Update(appServiceVirtualNetworkSwiftConnection *v1alpha1.AppServiceVirtualNetworkSwiftConnection) (result *v1alpha1.AppServiceVirtualNetworkSwiftConnection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(appservicevirtualnetworkswiftconnectionsResource, c.ns, appServiceVirtualNetworkSwiftConnection), &v1alpha1.AppServiceVirtualNetworkSwiftConnection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppServiceVirtualNetworkSwiftConnection), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAppServiceVirtualNetworkSwiftConnections) UpdateStatus(appServiceVirtualNetworkSwiftConnection *v1alpha1.AppServiceVirtualNetworkSwiftConnection) (*v1alpha1.AppServiceVirtualNetworkSwiftConnection, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(appservicevirtualnetworkswiftconnectionsResource, "status", c.ns, appServiceVirtualNetworkSwiftConnection), &v1alpha1.AppServiceVirtualNetworkSwiftConnection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppServiceVirtualNetworkSwiftConnection), err
}

// Delete takes name of the appServiceVirtualNetworkSwiftConnection and deletes it. Returns an error if one occurs.
func (c *FakeAppServiceVirtualNetworkSwiftConnections) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(appservicevirtualnetworkswiftconnectionsResource, c.ns, name), &v1alpha1.AppServiceVirtualNetworkSwiftConnection{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAppServiceVirtualNetworkSwiftConnections) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(appservicevirtualnetworkswiftconnectionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AppServiceVirtualNetworkSwiftConnectionList{})
	return err
}

// Patch applies the patch and returns the patched appServiceVirtualNetworkSwiftConnection.
func (c *FakeAppServiceVirtualNetworkSwiftConnections) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AppServiceVirtualNetworkSwiftConnection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(appservicevirtualnetworkswiftconnectionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AppServiceVirtualNetworkSwiftConnection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AppServiceVirtualNetworkSwiftConnection), err
}
