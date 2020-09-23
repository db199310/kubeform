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

// FakeBotWebApps implements BotWebAppInterface
type FakeBotWebApps struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var botwebappsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "botwebapps"}

var botwebappsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "BotWebApp"}

// Get takes name of the botWebApp, and returns the corresponding botWebApp object, and an error if there is any.
func (c *FakeBotWebApps) Get(name string, options v1.GetOptions) (result *v1alpha1.BotWebApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(botwebappsResource, c.ns, name), &v1alpha1.BotWebApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BotWebApp), err
}

// List takes label and field selectors, and returns the list of BotWebApps that match those selectors.
func (c *FakeBotWebApps) List(opts v1.ListOptions) (result *v1alpha1.BotWebAppList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(botwebappsResource, botwebappsKind, c.ns, opts), &v1alpha1.BotWebAppList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BotWebAppList{ListMeta: obj.(*v1alpha1.BotWebAppList).ListMeta}
	for _, item := range obj.(*v1alpha1.BotWebAppList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested botWebApps.
func (c *FakeBotWebApps) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(botwebappsResource, c.ns, opts))

}

// Create takes the representation of a botWebApp and creates it.  Returns the server's representation of the botWebApp, and an error, if there is any.
func (c *FakeBotWebApps) Create(botWebApp *v1alpha1.BotWebApp) (result *v1alpha1.BotWebApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(botwebappsResource, c.ns, botWebApp), &v1alpha1.BotWebApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BotWebApp), err
}

// Update takes the representation of a botWebApp and updates it. Returns the server's representation of the botWebApp, and an error, if there is any.
func (c *FakeBotWebApps) Update(botWebApp *v1alpha1.BotWebApp) (result *v1alpha1.BotWebApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(botwebappsResource, c.ns, botWebApp), &v1alpha1.BotWebApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BotWebApp), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBotWebApps) UpdateStatus(botWebApp *v1alpha1.BotWebApp) (*v1alpha1.BotWebApp, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(botwebappsResource, "status", c.ns, botWebApp), &v1alpha1.BotWebApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BotWebApp), err
}

// Delete takes name of the botWebApp and deletes it. Returns an error if one occurs.
func (c *FakeBotWebApps) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(botwebappsResource, c.ns, name), &v1alpha1.BotWebApp{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBotWebApps) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(botwebappsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.BotWebAppList{})
	return err
}

// Patch applies the patch and returns the patched botWebApp.
func (c *FakeBotWebApps) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BotWebApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(botwebappsResource, c.ns, name, pt, data, subresources...), &v1alpha1.BotWebApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BotWebApp), err
}