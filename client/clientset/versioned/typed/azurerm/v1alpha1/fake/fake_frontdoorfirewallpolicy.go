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

// FakeFrontdoorFirewallPolicies implements FrontdoorFirewallPolicyInterface
type FakeFrontdoorFirewallPolicies struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var frontdoorfirewallpoliciesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "frontdoorfirewallpolicies"}

var frontdoorfirewallpoliciesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "FrontdoorFirewallPolicy"}

// Get takes name of the frontdoorFirewallPolicy, and returns the corresponding frontdoorFirewallPolicy object, and an error if there is any.
func (c *FakeFrontdoorFirewallPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FrontdoorFirewallPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(frontdoorfirewallpoliciesResource, c.ns, name), &v1alpha1.FrontdoorFirewallPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FrontdoorFirewallPolicy), err
}

// List takes label and field selectors, and returns the list of FrontdoorFirewallPolicies that match those selectors.
func (c *FakeFrontdoorFirewallPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FrontdoorFirewallPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(frontdoorfirewallpoliciesResource, frontdoorfirewallpoliciesKind, c.ns, opts), &v1alpha1.FrontdoorFirewallPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FrontdoorFirewallPolicyList{ListMeta: obj.(*v1alpha1.FrontdoorFirewallPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.FrontdoorFirewallPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested frontdoorFirewallPolicies.
func (c *FakeFrontdoorFirewallPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(frontdoorfirewallpoliciesResource, c.ns, opts))

}

// Create takes the representation of a frontdoorFirewallPolicy and creates it.  Returns the server's representation of the frontdoorFirewallPolicy, and an error, if there is any.
func (c *FakeFrontdoorFirewallPolicies) Create(ctx context.Context, frontdoorFirewallPolicy *v1alpha1.FrontdoorFirewallPolicy, opts v1.CreateOptions) (result *v1alpha1.FrontdoorFirewallPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(frontdoorfirewallpoliciesResource, c.ns, frontdoorFirewallPolicy), &v1alpha1.FrontdoorFirewallPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FrontdoorFirewallPolicy), err
}

// Update takes the representation of a frontdoorFirewallPolicy and updates it. Returns the server's representation of the frontdoorFirewallPolicy, and an error, if there is any.
func (c *FakeFrontdoorFirewallPolicies) Update(ctx context.Context, frontdoorFirewallPolicy *v1alpha1.FrontdoorFirewallPolicy, opts v1.UpdateOptions) (result *v1alpha1.FrontdoorFirewallPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(frontdoorfirewallpoliciesResource, c.ns, frontdoorFirewallPolicy), &v1alpha1.FrontdoorFirewallPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FrontdoorFirewallPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFrontdoorFirewallPolicies) UpdateStatus(ctx context.Context, frontdoorFirewallPolicy *v1alpha1.FrontdoorFirewallPolicy, opts v1.UpdateOptions) (*v1alpha1.FrontdoorFirewallPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(frontdoorfirewallpoliciesResource, "status", c.ns, frontdoorFirewallPolicy), &v1alpha1.FrontdoorFirewallPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FrontdoorFirewallPolicy), err
}

// Delete takes name of the frontdoorFirewallPolicy and deletes it. Returns an error if one occurs.
func (c *FakeFrontdoorFirewallPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(frontdoorfirewallpoliciesResource, c.ns, name), &v1alpha1.FrontdoorFirewallPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFrontdoorFirewallPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(frontdoorfirewallpoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FrontdoorFirewallPolicyList{})
	return err
}

// Patch applies the patch and returns the patched frontdoorFirewallPolicy.
func (c *FakeFrontdoorFirewallPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FrontdoorFirewallPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(frontdoorfirewallpoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.FrontdoorFirewallPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FrontdoorFirewallPolicy), err
}
