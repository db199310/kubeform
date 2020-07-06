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

// FakeWafregionalIpsets implements WafregionalIpsetInterface
type FakeWafregionalIpsets struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var wafregionalipsetsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "wafregionalipsets"}

var wafregionalipsetsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "WafregionalIpset"}

// Get takes name of the wafregionalIpset, and returns the corresponding wafregionalIpset object, and an error if there is any.
func (c *FakeWafregionalIpsets) Get(name string, options v1.GetOptions) (result *v1alpha1.WafregionalIpset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(wafregionalipsetsResource, c.ns, name), &v1alpha1.WafregionalIpset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafregionalIpset), err
}

// List takes label and field selectors, and returns the list of WafregionalIpsets that match those selectors.
func (c *FakeWafregionalIpsets) List(opts v1.ListOptions) (result *v1alpha1.WafregionalIpsetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(wafregionalipsetsResource, wafregionalipsetsKind, c.ns, opts), &v1alpha1.WafregionalIpsetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.WafregionalIpsetList{ListMeta: obj.(*v1alpha1.WafregionalIpsetList).ListMeta}
	for _, item := range obj.(*v1alpha1.WafregionalIpsetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested wafregionalIpsets.
func (c *FakeWafregionalIpsets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(wafregionalipsetsResource, c.ns, opts))

}

// Create takes the representation of a wafregionalIpset and creates it.  Returns the server's representation of the wafregionalIpset, and an error, if there is any.
func (c *FakeWafregionalIpsets) Create(wafregionalIpset *v1alpha1.WafregionalIpset) (result *v1alpha1.WafregionalIpset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(wafregionalipsetsResource, c.ns, wafregionalIpset), &v1alpha1.WafregionalIpset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafregionalIpset), err
}

// Update takes the representation of a wafregionalIpset and updates it. Returns the server's representation of the wafregionalIpset, and an error, if there is any.
func (c *FakeWafregionalIpsets) Update(wafregionalIpset *v1alpha1.WafregionalIpset) (result *v1alpha1.WafregionalIpset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(wafregionalipsetsResource, c.ns, wafregionalIpset), &v1alpha1.WafregionalIpset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafregionalIpset), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeWafregionalIpsets) UpdateStatus(wafregionalIpset *v1alpha1.WafregionalIpset) (*v1alpha1.WafregionalIpset, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(wafregionalipsetsResource, "status", c.ns, wafregionalIpset), &v1alpha1.WafregionalIpset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafregionalIpset), err
}

// Delete takes name of the wafregionalIpset and deletes it. Returns an error if one occurs.
func (c *FakeWafregionalIpsets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(wafregionalipsetsResource, c.ns, name), &v1alpha1.WafregionalIpset{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWafregionalIpsets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(wafregionalipsetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.WafregionalIpsetList{})
	return err
}

// Patch applies the patch and returns the patched wafregionalIpset.
func (c *FakeWafregionalIpsets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WafregionalIpset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(wafregionalipsetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.WafregionalIpset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafregionalIpset), err
}
