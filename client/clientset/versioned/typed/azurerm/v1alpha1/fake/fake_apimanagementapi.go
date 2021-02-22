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

// FakeApiManagementAPIs implements ApiManagementAPIInterface
type FakeApiManagementAPIs struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var apimanagementapisResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "apimanagementapis"}

var apimanagementapisKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "ApiManagementAPI"}

// Get takes name of the apiManagementAPI, and returns the corresponding apiManagementAPI object, and an error if there is any.
func (c *FakeApiManagementAPIs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ApiManagementAPI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apimanagementapisResource, c.ns, name), &v1alpha1.ApiManagementAPI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementAPI), err
}

// List takes label and field selectors, and returns the list of ApiManagementAPIs that match those selectors.
func (c *FakeApiManagementAPIs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ApiManagementAPIList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apimanagementapisResource, apimanagementapisKind, c.ns, opts), &v1alpha1.ApiManagementAPIList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApiManagementAPIList{ListMeta: obj.(*v1alpha1.ApiManagementAPIList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApiManagementAPIList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apiManagementAPIs.
func (c *FakeApiManagementAPIs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apimanagementapisResource, c.ns, opts))

}

// Create takes the representation of a apiManagementAPI and creates it.  Returns the server's representation of the apiManagementAPI, and an error, if there is any.
func (c *FakeApiManagementAPIs) Create(ctx context.Context, apiManagementAPI *v1alpha1.ApiManagementAPI, opts v1.CreateOptions) (result *v1alpha1.ApiManagementAPI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apimanagementapisResource, c.ns, apiManagementAPI), &v1alpha1.ApiManagementAPI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementAPI), err
}

// Update takes the representation of a apiManagementAPI and updates it. Returns the server's representation of the apiManagementAPI, and an error, if there is any.
func (c *FakeApiManagementAPIs) Update(ctx context.Context, apiManagementAPI *v1alpha1.ApiManagementAPI, opts v1.UpdateOptions) (result *v1alpha1.ApiManagementAPI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apimanagementapisResource, c.ns, apiManagementAPI), &v1alpha1.ApiManagementAPI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementAPI), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApiManagementAPIs) UpdateStatus(ctx context.Context, apiManagementAPI *v1alpha1.ApiManagementAPI, opts v1.UpdateOptions) (*v1alpha1.ApiManagementAPI, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apimanagementapisResource, "status", c.ns, apiManagementAPI), &v1alpha1.ApiManagementAPI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementAPI), err
}

// Delete takes name of the apiManagementAPI and deletes it. Returns an error if one occurs.
func (c *FakeApiManagementAPIs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(apimanagementapisResource, c.ns, name), &v1alpha1.ApiManagementAPI{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApiManagementAPIs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apimanagementapisResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApiManagementAPIList{})
	return err
}

// Patch applies the patch and returns the patched apiManagementAPI.
func (c *FakeApiManagementAPIs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApiManagementAPI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apimanagementapisResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApiManagementAPI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementAPI), err
}
