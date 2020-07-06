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

// FakeWafregionalRules implements WafregionalRuleInterface
type FakeWafregionalRules struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var wafregionalrulesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "wafregionalrules"}

var wafregionalrulesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "WafregionalRule"}

// Get takes name of the wafregionalRule, and returns the corresponding wafregionalRule object, and an error if there is any.
func (c *FakeWafregionalRules) Get(name string, options v1.GetOptions) (result *v1alpha1.WafregionalRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(wafregionalrulesResource, c.ns, name), &v1alpha1.WafregionalRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafregionalRule), err
}

// List takes label and field selectors, and returns the list of WafregionalRules that match those selectors.
func (c *FakeWafregionalRules) List(opts v1.ListOptions) (result *v1alpha1.WafregionalRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(wafregionalrulesResource, wafregionalrulesKind, c.ns, opts), &v1alpha1.WafregionalRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.WafregionalRuleList{ListMeta: obj.(*v1alpha1.WafregionalRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.WafregionalRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested wafregionalRules.
func (c *FakeWafregionalRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(wafregionalrulesResource, c.ns, opts))

}

// Create takes the representation of a wafregionalRule and creates it.  Returns the server's representation of the wafregionalRule, and an error, if there is any.
func (c *FakeWafregionalRules) Create(wafregionalRule *v1alpha1.WafregionalRule) (result *v1alpha1.WafregionalRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(wafregionalrulesResource, c.ns, wafregionalRule), &v1alpha1.WafregionalRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafregionalRule), err
}

// Update takes the representation of a wafregionalRule and updates it. Returns the server's representation of the wafregionalRule, and an error, if there is any.
func (c *FakeWafregionalRules) Update(wafregionalRule *v1alpha1.WafregionalRule) (result *v1alpha1.WafregionalRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(wafregionalrulesResource, c.ns, wafregionalRule), &v1alpha1.WafregionalRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafregionalRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeWafregionalRules) UpdateStatus(wafregionalRule *v1alpha1.WafregionalRule) (*v1alpha1.WafregionalRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(wafregionalrulesResource, "status", c.ns, wafregionalRule), &v1alpha1.WafregionalRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafregionalRule), err
}

// Delete takes name of the wafregionalRule and deletes it. Returns an error if one occurs.
func (c *FakeWafregionalRules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(wafregionalrulesResource, c.ns, name), &v1alpha1.WafregionalRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWafregionalRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(wafregionalrulesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.WafregionalRuleList{})
	return err
}

// Patch applies the patch and returns the patched wafregionalRule.
func (c *FakeWafregionalRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WafregionalRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(wafregionalrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.WafregionalRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafregionalRule), err
}
