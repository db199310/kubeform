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

// FakeSiteRecoveryReplicationPolicies implements SiteRecoveryReplicationPolicyInterface
type FakeSiteRecoveryReplicationPolicies struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var siterecoveryreplicationpoliciesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "siterecoveryreplicationpolicies"}

var siterecoveryreplicationpoliciesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "SiteRecoveryReplicationPolicy"}

// Get takes name of the siteRecoveryReplicationPolicy, and returns the corresponding siteRecoveryReplicationPolicy object, and an error if there is any.
func (c *FakeSiteRecoveryReplicationPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.SiteRecoveryReplicationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(siterecoveryreplicationpoliciesResource, c.ns, name), &v1alpha1.SiteRecoveryReplicationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SiteRecoveryReplicationPolicy), err
}

// List takes label and field selectors, and returns the list of SiteRecoveryReplicationPolicies that match those selectors.
func (c *FakeSiteRecoveryReplicationPolicies) List(opts v1.ListOptions) (result *v1alpha1.SiteRecoveryReplicationPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(siterecoveryreplicationpoliciesResource, siterecoveryreplicationpoliciesKind, c.ns, opts), &v1alpha1.SiteRecoveryReplicationPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SiteRecoveryReplicationPolicyList{ListMeta: obj.(*v1alpha1.SiteRecoveryReplicationPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.SiteRecoveryReplicationPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested siteRecoveryReplicationPolicies.
func (c *FakeSiteRecoveryReplicationPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(siterecoveryreplicationpoliciesResource, c.ns, opts))

}

// Create takes the representation of a siteRecoveryReplicationPolicy and creates it.  Returns the server's representation of the siteRecoveryReplicationPolicy, and an error, if there is any.
func (c *FakeSiteRecoveryReplicationPolicies) Create(siteRecoveryReplicationPolicy *v1alpha1.SiteRecoveryReplicationPolicy) (result *v1alpha1.SiteRecoveryReplicationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(siterecoveryreplicationpoliciesResource, c.ns, siteRecoveryReplicationPolicy), &v1alpha1.SiteRecoveryReplicationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SiteRecoveryReplicationPolicy), err
}

// Update takes the representation of a siteRecoveryReplicationPolicy and updates it. Returns the server's representation of the siteRecoveryReplicationPolicy, and an error, if there is any.
func (c *FakeSiteRecoveryReplicationPolicies) Update(siteRecoveryReplicationPolicy *v1alpha1.SiteRecoveryReplicationPolicy) (result *v1alpha1.SiteRecoveryReplicationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(siterecoveryreplicationpoliciesResource, c.ns, siteRecoveryReplicationPolicy), &v1alpha1.SiteRecoveryReplicationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SiteRecoveryReplicationPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSiteRecoveryReplicationPolicies) UpdateStatus(siteRecoveryReplicationPolicy *v1alpha1.SiteRecoveryReplicationPolicy) (*v1alpha1.SiteRecoveryReplicationPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(siterecoveryreplicationpoliciesResource, "status", c.ns, siteRecoveryReplicationPolicy), &v1alpha1.SiteRecoveryReplicationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SiteRecoveryReplicationPolicy), err
}

// Delete takes name of the siteRecoveryReplicationPolicy and deletes it. Returns an error if one occurs.
func (c *FakeSiteRecoveryReplicationPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(siterecoveryreplicationpoliciesResource, c.ns, name), &v1alpha1.SiteRecoveryReplicationPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSiteRecoveryReplicationPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(siterecoveryreplicationpoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SiteRecoveryReplicationPolicyList{})
	return err
}

// Patch applies the patch and returns the patched siteRecoveryReplicationPolicy.
func (c *FakeSiteRecoveryReplicationPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SiteRecoveryReplicationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(siterecoveryreplicationpoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.SiteRecoveryReplicationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SiteRecoveryReplicationPolicy), err
}