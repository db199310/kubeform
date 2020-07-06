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

// FakeStorageBucketIamPolicies implements StorageBucketIamPolicyInterface
type FakeStorageBucketIamPolicies struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var storagebucketiampoliciesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "storagebucketiampolicies"}

var storagebucketiampoliciesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "StorageBucketIamPolicy"}

// Get takes name of the storageBucketIamPolicy, and returns the corresponding storageBucketIamPolicy object, and an error if there is any.
func (c *FakeStorageBucketIamPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.StorageBucketIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(storagebucketiampoliciesResource, c.ns, name), &v1alpha1.StorageBucketIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageBucketIamPolicy), err
}

// List takes label and field selectors, and returns the list of StorageBucketIamPolicies that match those selectors.
func (c *FakeStorageBucketIamPolicies) List(opts v1.ListOptions) (result *v1alpha1.StorageBucketIamPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(storagebucketiampoliciesResource, storagebucketiampoliciesKind, c.ns, opts), &v1alpha1.StorageBucketIamPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StorageBucketIamPolicyList{ListMeta: obj.(*v1alpha1.StorageBucketIamPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.StorageBucketIamPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested storageBucketIamPolicies.
func (c *FakeStorageBucketIamPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(storagebucketiampoliciesResource, c.ns, opts))

}

// Create takes the representation of a storageBucketIamPolicy and creates it.  Returns the server's representation of the storageBucketIamPolicy, and an error, if there is any.
func (c *FakeStorageBucketIamPolicies) Create(storageBucketIamPolicy *v1alpha1.StorageBucketIamPolicy) (result *v1alpha1.StorageBucketIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(storagebucketiampoliciesResource, c.ns, storageBucketIamPolicy), &v1alpha1.StorageBucketIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageBucketIamPolicy), err
}

// Update takes the representation of a storageBucketIamPolicy and updates it. Returns the server's representation of the storageBucketIamPolicy, and an error, if there is any.
func (c *FakeStorageBucketIamPolicies) Update(storageBucketIamPolicy *v1alpha1.StorageBucketIamPolicy) (result *v1alpha1.StorageBucketIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(storagebucketiampoliciesResource, c.ns, storageBucketIamPolicy), &v1alpha1.StorageBucketIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageBucketIamPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStorageBucketIamPolicies) UpdateStatus(storageBucketIamPolicy *v1alpha1.StorageBucketIamPolicy) (*v1alpha1.StorageBucketIamPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(storagebucketiampoliciesResource, "status", c.ns, storageBucketIamPolicy), &v1alpha1.StorageBucketIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageBucketIamPolicy), err
}

// Delete takes name of the storageBucketIamPolicy and deletes it. Returns an error if one occurs.
func (c *FakeStorageBucketIamPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(storagebucketiampoliciesResource, c.ns, name), &v1alpha1.StorageBucketIamPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStorageBucketIamPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(storagebucketiampoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.StorageBucketIamPolicyList{})
	return err
}

// Patch applies the patch and returns the patched storageBucketIamPolicy.
func (c *FakeStorageBucketIamPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StorageBucketIamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(storagebucketiampoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.StorageBucketIamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageBucketIamPolicy), err
}
