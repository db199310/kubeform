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

// FakeWafWebACLs implements WafWebACLInterface
type FakeWafWebACLs struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var wafwebaclsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "wafwebacls"}

var wafwebaclsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "WafWebACL"}

// Get takes name of the wafWebACL, and returns the corresponding wafWebACL object, and an error if there is any.
func (c *FakeWafWebACLs) Get(name string, options v1.GetOptions) (result *v1alpha1.WafWebACL, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(wafwebaclsResource, c.ns, name), &v1alpha1.WafWebACL{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafWebACL), err
}

// List takes label and field selectors, and returns the list of WafWebACLs that match those selectors.
func (c *FakeWafWebACLs) List(opts v1.ListOptions) (result *v1alpha1.WafWebACLList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(wafwebaclsResource, wafwebaclsKind, c.ns, opts), &v1alpha1.WafWebACLList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.WafWebACLList{ListMeta: obj.(*v1alpha1.WafWebACLList).ListMeta}
	for _, item := range obj.(*v1alpha1.WafWebACLList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested wafWebACLs.
func (c *FakeWafWebACLs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(wafwebaclsResource, c.ns, opts))

}

// Create takes the representation of a wafWebACL and creates it.  Returns the server's representation of the wafWebACL, and an error, if there is any.
func (c *FakeWafWebACLs) Create(wafWebACL *v1alpha1.WafWebACL) (result *v1alpha1.WafWebACL, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(wafwebaclsResource, c.ns, wafWebACL), &v1alpha1.WafWebACL{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafWebACL), err
}

// Update takes the representation of a wafWebACL and updates it. Returns the server's representation of the wafWebACL, and an error, if there is any.
func (c *FakeWafWebACLs) Update(wafWebACL *v1alpha1.WafWebACL) (result *v1alpha1.WafWebACL, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(wafwebaclsResource, c.ns, wafWebACL), &v1alpha1.WafWebACL{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafWebACL), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeWafWebACLs) UpdateStatus(wafWebACL *v1alpha1.WafWebACL) (*v1alpha1.WafWebACL, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(wafwebaclsResource, "status", c.ns, wafWebACL), &v1alpha1.WafWebACL{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafWebACL), err
}

// Delete takes name of the wafWebACL and deletes it. Returns an error if one occurs.
func (c *FakeWafWebACLs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(wafwebaclsResource, c.ns, name), &v1alpha1.WafWebACL{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWafWebACLs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(wafwebaclsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.WafWebACLList{})
	return err
}

// Patch applies the patch and returns the patched wafWebACL.
func (c *FakeWafWebACLs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WafWebACL, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(wafwebaclsResource, c.ns, name, pt, data, subresources...), &v1alpha1.WafWebACL{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafWebACL), err
}
