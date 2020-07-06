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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeIamUserPolicies implements IamUserPolicyInterface
type FakeIamUserPolicies struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var iamuserpoliciesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "iamuserpolicies"}

var iamuserpoliciesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "IamUserPolicy"}

// Get takes name of the iamUserPolicy, and returns the corresponding iamUserPolicy object, and an error if there is any.
func (c *FakeIamUserPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.IamUserPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(iamuserpoliciesResource, c.ns, name), &v1alpha1.IamUserPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamUserPolicy), err
}

// List takes label and field selectors, and returns the list of IamUserPolicies that match those selectors.
func (c *FakeIamUserPolicies) List(opts v1.ListOptions) (result *v1alpha1.IamUserPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(iamuserpoliciesResource, iamuserpoliciesKind, c.ns, opts), &v1alpha1.IamUserPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IamUserPolicyList{ListMeta: obj.(*v1alpha1.IamUserPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.IamUserPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iamUserPolicies.
func (c *FakeIamUserPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(iamuserpoliciesResource, c.ns, opts))

}

// Create takes the representation of a iamUserPolicy and creates it.  Returns the server's representation of the iamUserPolicy, and an error, if there is any.
func (c *FakeIamUserPolicies) Create(iamUserPolicy *v1alpha1.IamUserPolicy) (result *v1alpha1.IamUserPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(iamuserpoliciesResource, c.ns, iamUserPolicy), &v1alpha1.IamUserPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamUserPolicy), err
}

// Update takes the representation of a iamUserPolicy and updates it. Returns the server's representation of the iamUserPolicy, and an error, if there is any.
func (c *FakeIamUserPolicies) Update(iamUserPolicy *v1alpha1.IamUserPolicy) (result *v1alpha1.IamUserPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(iamuserpoliciesResource, c.ns, iamUserPolicy), &v1alpha1.IamUserPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamUserPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIamUserPolicies) UpdateStatus(iamUserPolicy *v1alpha1.IamUserPolicy) (*v1alpha1.IamUserPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(iamuserpoliciesResource, "status", c.ns, iamUserPolicy), &v1alpha1.IamUserPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamUserPolicy), err
}

// Delete takes name of the iamUserPolicy and deletes it. Returns an error if one occurs.
func (c *FakeIamUserPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(iamuserpoliciesResource, c.ns, name), &v1alpha1.IamUserPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIamUserPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(iamuserpoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.IamUserPolicyList{})
	return err
}

// Patch applies the patch and returns the patched iamUserPolicy.
func (c *FakeIamUserPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IamUserPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(iamuserpoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.IamUserPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamUserPolicy), err
}
