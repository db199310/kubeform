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

// FakePrivateLinkServices implements PrivateLinkServiceInterface
type FakePrivateLinkServices struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var privatelinkservicesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "privatelinkservices"}

var privatelinkservicesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "PrivateLinkService"}

// Get takes name of the privateLinkService, and returns the corresponding privateLinkService object, and an error if there is any.
func (c *FakePrivateLinkServices) Get(name string, options v1.GetOptions) (result *v1alpha1.PrivateLinkService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(privatelinkservicesResource, c.ns, name), &v1alpha1.PrivateLinkService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrivateLinkService), err
}

// List takes label and field selectors, and returns the list of PrivateLinkServices that match those selectors.
func (c *FakePrivateLinkServices) List(opts v1.ListOptions) (result *v1alpha1.PrivateLinkServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(privatelinkservicesResource, privatelinkservicesKind, c.ns, opts), &v1alpha1.PrivateLinkServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PrivateLinkServiceList{ListMeta: obj.(*v1alpha1.PrivateLinkServiceList).ListMeta}
	for _, item := range obj.(*v1alpha1.PrivateLinkServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested privateLinkServices.
func (c *FakePrivateLinkServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(privatelinkservicesResource, c.ns, opts))

}

// Create takes the representation of a privateLinkService and creates it.  Returns the server's representation of the privateLinkService, and an error, if there is any.
func (c *FakePrivateLinkServices) Create(privateLinkService *v1alpha1.PrivateLinkService) (result *v1alpha1.PrivateLinkService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(privatelinkservicesResource, c.ns, privateLinkService), &v1alpha1.PrivateLinkService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrivateLinkService), err
}

// Update takes the representation of a privateLinkService and updates it. Returns the server's representation of the privateLinkService, and an error, if there is any.
func (c *FakePrivateLinkServices) Update(privateLinkService *v1alpha1.PrivateLinkService) (result *v1alpha1.PrivateLinkService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(privatelinkservicesResource, c.ns, privateLinkService), &v1alpha1.PrivateLinkService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrivateLinkService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePrivateLinkServices) UpdateStatus(privateLinkService *v1alpha1.PrivateLinkService) (*v1alpha1.PrivateLinkService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(privatelinkservicesResource, "status", c.ns, privateLinkService), &v1alpha1.PrivateLinkService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrivateLinkService), err
}

// Delete takes name of the privateLinkService and deletes it. Returns an error if one occurs.
func (c *FakePrivateLinkServices) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(privatelinkservicesResource, c.ns, name), &v1alpha1.PrivateLinkService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePrivateLinkServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(privatelinkservicesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.PrivateLinkServiceList{})
	return err
}

// Patch applies the patch and returns the patched privateLinkService.
func (c *FakePrivateLinkServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PrivateLinkService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(privatelinkservicesResource, c.ns, name, pt, data, subresources...), &v1alpha1.PrivateLinkService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrivateLinkService), err
}
