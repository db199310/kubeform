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

// FakeApiManagementAPIOperationPolicies implements ApiManagementAPIOperationPolicyInterface
type FakeApiManagementAPIOperationPolicies struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var apimanagementapioperationpoliciesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "apimanagementapioperationpolicies"}

var apimanagementapioperationpoliciesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "ApiManagementAPIOperationPolicy"}

// Get takes name of the apiManagementAPIOperationPolicy, and returns the corresponding apiManagementAPIOperationPolicy object, and an error if there is any.
func (c *FakeApiManagementAPIOperationPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ApiManagementAPIOperationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apimanagementapioperationpoliciesResource, c.ns, name), &v1alpha1.ApiManagementAPIOperationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementAPIOperationPolicy), err
}

// List takes label and field selectors, and returns the list of ApiManagementAPIOperationPolicies that match those selectors.
func (c *FakeApiManagementAPIOperationPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ApiManagementAPIOperationPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apimanagementapioperationpoliciesResource, apimanagementapioperationpoliciesKind, c.ns, opts), &v1alpha1.ApiManagementAPIOperationPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApiManagementAPIOperationPolicyList{ListMeta: obj.(*v1alpha1.ApiManagementAPIOperationPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApiManagementAPIOperationPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apiManagementAPIOperationPolicies.
func (c *FakeApiManagementAPIOperationPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apimanagementapioperationpoliciesResource, c.ns, opts))

}

// Create takes the representation of a apiManagementAPIOperationPolicy and creates it.  Returns the server's representation of the apiManagementAPIOperationPolicy, and an error, if there is any.
func (c *FakeApiManagementAPIOperationPolicies) Create(ctx context.Context, apiManagementAPIOperationPolicy *v1alpha1.ApiManagementAPIOperationPolicy, opts v1.CreateOptions) (result *v1alpha1.ApiManagementAPIOperationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apimanagementapioperationpoliciesResource, c.ns, apiManagementAPIOperationPolicy), &v1alpha1.ApiManagementAPIOperationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementAPIOperationPolicy), err
}

// Update takes the representation of a apiManagementAPIOperationPolicy and updates it. Returns the server's representation of the apiManagementAPIOperationPolicy, and an error, if there is any.
func (c *FakeApiManagementAPIOperationPolicies) Update(ctx context.Context, apiManagementAPIOperationPolicy *v1alpha1.ApiManagementAPIOperationPolicy, opts v1.UpdateOptions) (result *v1alpha1.ApiManagementAPIOperationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apimanagementapioperationpoliciesResource, c.ns, apiManagementAPIOperationPolicy), &v1alpha1.ApiManagementAPIOperationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementAPIOperationPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApiManagementAPIOperationPolicies) UpdateStatus(ctx context.Context, apiManagementAPIOperationPolicy *v1alpha1.ApiManagementAPIOperationPolicy, opts v1.UpdateOptions) (*v1alpha1.ApiManagementAPIOperationPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apimanagementapioperationpoliciesResource, "status", c.ns, apiManagementAPIOperationPolicy), &v1alpha1.ApiManagementAPIOperationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementAPIOperationPolicy), err
}

// Delete takes name of the apiManagementAPIOperationPolicy and deletes it. Returns an error if one occurs.
func (c *FakeApiManagementAPIOperationPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(apimanagementapioperationpoliciesResource, c.ns, name), &v1alpha1.ApiManagementAPIOperationPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApiManagementAPIOperationPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apimanagementapioperationpoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApiManagementAPIOperationPolicyList{})
	return err
}

// Patch applies the patch and returns the patched apiManagementAPIOperationPolicy.
func (c *FakeApiManagementAPIOperationPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApiManagementAPIOperationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apimanagementapioperationpoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApiManagementAPIOperationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementAPIOperationPolicy), err
}
