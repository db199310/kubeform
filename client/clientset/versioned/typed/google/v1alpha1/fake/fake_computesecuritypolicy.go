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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakeComputeSecurityPolicies implements ComputeSecurityPolicyInterface
type FakeComputeSecurityPolicies struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var computesecuritypoliciesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "computesecuritypolicies"}

var computesecuritypoliciesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "ComputeSecurityPolicy"}

// Get takes name of the computeSecurityPolicy, and returns the corresponding computeSecurityPolicy object, and an error if there is any.
func (c *FakeComputeSecurityPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeSecurityPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computesecuritypoliciesResource, c.ns, name), &v1alpha1.ComputeSecurityPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeSecurityPolicy), err
}

// List takes label and field selectors, and returns the list of ComputeSecurityPolicies that match those selectors.
func (c *FakeComputeSecurityPolicies) List(opts v1.ListOptions) (result *v1alpha1.ComputeSecurityPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computesecuritypoliciesResource, computesecuritypoliciesKind, c.ns, opts), &v1alpha1.ComputeSecurityPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComputeSecurityPolicyList{ListMeta: obj.(*v1alpha1.ComputeSecurityPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.ComputeSecurityPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeSecurityPolicies.
func (c *FakeComputeSecurityPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computesecuritypoliciesResource, c.ns, opts))

}

// Create takes the representation of a computeSecurityPolicy and creates it.  Returns the server's representation of the computeSecurityPolicy, and an error, if there is any.
func (c *FakeComputeSecurityPolicies) Create(computeSecurityPolicy *v1alpha1.ComputeSecurityPolicy) (result *v1alpha1.ComputeSecurityPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computesecuritypoliciesResource, c.ns, computeSecurityPolicy), &v1alpha1.ComputeSecurityPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeSecurityPolicy), err
}

// Update takes the representation of a computeSecurityPolicy and updates it. Returns the server's representation of the computeSecurityPolicy, and an error, if there is any.
func (c *FakeComputeSecurityPolicies) Update(computeSecurityPolicy *v1alpha1.ComputeSecurityPolicy) (result *v1alpha1.ComputeSecurityPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computesecuritypoliciesResource, c.ns, computeSecurityPolicy), &v1alpha1.ComputeSecurityPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeSecurityPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeSecurityPolicies) UpdateStatus(computeSecurityPolicy *v1alpha1.ComputeSecurityPolicy) (*v1alpha1.ComputeSecurityPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computesecuritypoliciesResource, "status", c.ns, computeSecurityPolicy), &v1alpha1.ComputeSecurityPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeSecurityPolicy), err
}

// Delete takes name of the computeSecurityPolicy and deletes it. Returns an error if one occurs.
func (c *FakeComputeSecurityPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(computesecuritypoliciesResource, c.ns, name), &v1alpha1.ComputeSecurityPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeSecurityPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computesecuritypoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComputeSecurityPolicyList{})
	return err
}

// Patch applies the patch and returns the patched computeSecurityPolicy.
func (c *FakeComputeSecurityPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeSecurityPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computesecuritypoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ComputeSecurityPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeSecurityPolicy), err
}
