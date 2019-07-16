/*
Copyright 2019 The Kubeform Authors.

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
	v1alpha1 "kubeform.dev/kubeform/apis/linode/v1alpha1"
)

// FakeRdnses implements RdnsInterface
type FakeRdnses struct {
	Fake *FakeLinodeV1alpha1
}

var rdnsesResource = schema.GroupVersionResource{Group: "linode.kubeform.com", Version: "v1alpha1", Resource: "rdnses"}

var rdnsesKind = schema.GroupVersionKind{Group: "linode.kubeform.com", Version: "v1alpha1", Kind: "Rdns"}

// Get takes name of the rdns, and returns the corresponding rdns object, and an error if there is any.
func (c *FakeRdnses) Get(name string, options v1.GetOptions) (result *v1alpha1.Rdns, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(rdnsesResource, name), &v1alpha1.Rdns{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Rdns), err
}

// List takes label and field selectors, and returns the list of Rdnses that match those selectors.
func (c *FakeRdnses) List(opts v1.ListOptions) (result *v1alpha1.RdnsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(rdnsesResource, rdnsesKind, opts), &v1alpha1.RdnsList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RdnsList{ListMeta: obj.(*v1alpha1.RdnsList).ListMeta}
	for _, item := range obj.(*v1alpha1.RdnsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested rdnses.
func (c *FakeRdnses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(rdnsesResource, opts))
}

// Create takes the representation of a rdns and creates it.  Returns the server's representation of the rdns, and an error, if there is any.
func (c *FakeRdnses) Create(rdns *v1alpha1.Rdns) (result *v1alpha1.Rdns, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(rdnsesResource, rdns), &v1alpha1.Rdns{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Rdns), err
}

// Update takes the representation of a rdns and updates it. Returns the server's representation of the rdns, and an error, if there is any.
func (c *FakeRdnses) Update(rdns *v1alpha1.Rdns) (result *v1alpha1.Rdns, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(rdnsesResource, rdns), &v1alpha1.Rdns{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Rdns), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRdnses) UpdateStatus(rdns *v1alpha1.Rdns) (*v1alpha1.Rdns, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(rdnsesResource, "status", rdns), &v1alpha1.Rdns{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Rdns), err
}

// Delete takes name of the rdns and deletes it. Returns an error if one occurs.
func (c *FakeRdnses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(rdnsesResource, name), &v1alpha1.Rdns{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRdnses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(rdnsesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.RdnsList{})
	return err
}

// Patch applies the patch and returns the patched rdns.
func (c *FakeRdnses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Rdns, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(rdnsesResource, name, pt, data, subresources...), &v1alpha1.Rdns{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Rdns), err
}