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

// FakeNetworkInterfaceApplicationGatewayBackendAddressPoolAssociations implements NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationInterface
type FakeNetworkInterfaceApplicationGatewayBackendAddressPoolAssociations struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var networkinterfaceapplicationgatewaybackendaddresspoolassociationsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "networkinterfaceapplicationgatewaybackendaddresspoolassociations"}

var networkinterfaceapplicationgatewaybackendaddresspoolassociationsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation"}

// Get takes name of the networkInterfaceApplicationGatewayBackendAddressPoolAssociation, and returns the corresponding networkInterfaceApplicationGatewayBackendAddressPoolAssociation object, and an error if there is any.
func (c *FakeNetworkInterfaceApplicationGatewayBackendAddressPoolAssociations) Get(name string, options v1.GetOptions) (result *v1alpha1.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkinterfaceapplicationgatewaybackendaddresspoolassociationsResource, c.ns, name), &v1alpha1.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation), err
}

// List takes label and field selectors, and returns the list of NetworkInterfaceApplicationGatewayBackendAddressPoolAssociations that match those selectors.
func (c *FakeNetworkInterfaceApplicationGatewayBackendAddressPoolAssociations) List(opts v1.ListOptions) (result *v1alpha1.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkinterfaceapplicationgatewaybackendaddresspoolassociationsResource, networkinterfaceapplicationgatewaybackendaddresspoolassociationsKind, c.ns, opts), &v1alpha1.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{ListMeta: obj.(*v1alpha1.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList).ListMeta}
	for _, item := range obj.(*v1alpha1.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkInterfaceApplicationGatewayBackendAddressPoolAssociations.
func (c *FakeNetworkInterfaceApplicationGatewayBackendAddressPoolAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkinterfaceapplicationgatewaybackendaddresspoolassociationsResource, c.ns, opts))

}

// Create takes the representation of a networkInterfaceApplicationGatewayBackendAddressPoolAssociation and creates it.  Returns the server's representation of the networkInterfaceApplicationGatewayBackendAddressPoolAssociation, and an error, if there is any.
func (c *FakeNetworkInterfaceApplicationGatewayBackendAddressPoolAssociations) Create(networkInterfaceApplicationGatewayBackendAddressPoolAssociation *v1alpha1.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation) (result *v1alpha1.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkinterfaceapplicationgatewaybackendaddresspoolassociationsResource, c.ns, networkInterfaceApplicationGatewayBackendAddressPoolAssociation), &v1alpha1.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation), err
}

// Update takes the representation of a networkInterfaceApplicationGatewayBackendAddressPoolAssociation and updates it. Returns the server's representation of the networkInterfaceApplicationGatewayBackendAddressPoolAssociation, and an error, if there is any.
func (c *FakeNetworkInterfaceApplicationGatewayBackendAddressPoolAssociations) Update(networkInterfaceApplicationGatewayBackendAddressPoolAssociation *v1alpha1.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation) (result *v1alpha1.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkinterfaceapplicationgatewaybackendaddresspoolassociationsResource, c.ns, networkInterfaceApplicationGatewayBackendAddressPoolAssociation), &v1alpha1.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkInterfaceApplicationGatewayBackendAddressPoolAssociations) UpdateStatus(networkInterfaceApplicationGatewayBackendAddressPoolAssociation *v1alpha1.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation) (*v1alpha1.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(networkinterfaceapplicationgatewaybackendaddresspoolassociationsResource, "status", c.ns, networkInterfaceApplicationGatewayBackendAddressPoolAssociation), &v1alpha1.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation), err
}

// Delete takes name of the networkInterfaceApplicationGatewayBackendAddressPoolAssociation and deletes it. Returns an error if one occurs.
func (c *FakeNetworkInterfaceApplicationGatewayBackendAddressPoolAssociations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(networkinterfaceapplicationgatewaybackendaddresspoolassociationsResource, c.ns, name), &v1alpha1.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkInterfaceApplicationGatewayBackendAddressPoolAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networkinterfaceapplicationgatewaybackendaddresspoolassociationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{})
	return err
}

// Patch applies the patch and returns the patched networkInterfaceApplicationGatewayBackendAddressPoolAssociation.
func (c *FakeNetworkInterfaceApplicationGatewayBackendAddressPoolAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkinterfaceapplicationgatewaybackendaddresspoolassociationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation), err
}
