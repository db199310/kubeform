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

// FakeWebApplicationFirewallPolicies implements WebApplicationFirewallPolicyInterface
type FakeWebApplicationFirewallPolicies struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var webapplicationfirewallpoliciesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "webapplicationfirewallpolicies"}

var webapplicationfirewallpoliciesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "WebApplicationFirewallPolicy"}

// Get takes name of the webApplicationFirewallPolicy, and returns the corresponding webApplicationFirewallPolicy object, and an error if there is any.
func (c *FakeWebApplicationFirewallPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.WebApplicationFirewallPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(webapplicationfirewallpoliciesResource, c.ns, name), &v1alpha1.WebApplicationFirewallPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WebApplicationFirewallPolicy), err
}

// List takes label and field selectors, and returns the list of WebApplicationFirewallPolicies that match those selectors.
func (c *FakeWebApplicationFirewallPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.WebApplicationFirewallPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(webapplicationfirewallpoliciesResource, webapplicationfirewallpoliciesKind, c.ns, opts), &v1alpha1.WebApplicationFirewallPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.WebApplicationFirewallPolicyList{ListMeta: obj.(*v1alpha1.WebApplicationFirewallPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.WebApplicationFirewallPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested webApplicationFirewallPolicies.
func (c *FakeWebApplicationFirewallPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(webapplicationfirewallpoliciesResource, c.ns, opts))

}

// Create takes the representation of a webApplicationFirewallPolicy and creates it.  Returns the server's representation of the webApplicationFirewallPolicy, and an error, if there is any.
func (c *FakeWebApplicationFirewallPolicies) Create(ctx context.Context, webApplicationFirewallPolicy *v1alpha1.WebApplicationFirewallPolicy, opts v1.CreateOptions) (result *v1alpha1.WebApplicationFirewallPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(webapplicationfirewallpoliciesResource, c.ns, webApplicationFirewallPolicy), &v1alpha1.WebApplicationFirewallPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WebApplicationFirewallPolicy), err
}

// Update takes the representation of a webApplicationFirewallPolicy and updates it. Returns the server's representation of the webApplicationFirewallPolicy, and an error, if there is any.
func (c *FakeWebApplicationFirewallPolicies) Update(ctx context.Context, webApplicationFirewallPolicy *v1alpha1.WebApplicationFirewallPolicy, opts v1.UpdateOptions) (result *v1alpha1.WebApplicationFirewallPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(webapplicationfirewallpoliciesResource, c.ns, webApplicationFirewallPolicy), &v1alpha1.WebApplicationFirewallPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WebApplicationFirewallPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeWebApplicationFirewallPolicies) UpdateStatus(ctx context.Context, webApplicationFirewallPolicy *v1alpha1.WebApplicationFirewallPolicy, opts v1.UpdateOptions) (*v1alpha1.WebApplicationFirewallPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(webapplicationfirewallpoliciesResource, "status", c.ns, webApplicationFirewallPolicy), &v1alpha1.WebApplicationFirewallPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WebApplicationFirewallPolicy), err
}

// Delete takes name of the webApplicationFirewallPolicy and deletes it. Returns an error if one occurs.
func (c *FakeWebApplicationFirewallPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(webapplicationfirewallpoliciesResource, c.ns, name), &v1alpha1.WebApplicationFirewallPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWebApplicationFirewallPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(webapplicationfirewallpoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.WebApplicationFirewallPolicyList{})
	return err
}

// Patch applies the patch and returns the patched webApplicationFirewallPolicy.
func (c *FakeWebApplicationFirewallPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.WebApplicationFirewallPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(webapplicationfirewallpoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.WebApplicationFirewallPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WebApplicationFirewallPolicy), err
}
