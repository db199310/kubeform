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

// FakeServicebusSubscriptionRules implements ServicebusSubscriptionRuleInterface
type FakeServicebusSubscriptionRules struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var servicebussubscriptionrulesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "servicebussubscriptionrules"}

var servicebussubscriptionrulesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "ServicebusSubscriptionRule"}

// Get takes name of the servicebusSubscriptionRule, and returns the corresponding servicebusSubscriptionRule object, and an error if there is any.
func (c *FakeServicebusSubscriptionRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ServicebusSubscriptionRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servicebussubscriptionrulesResource, c.ns, name), &v1alpha1.ServicebusSubscriptionRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicebusSubscriptionRule), err
}

// List takes label and field selectors, and returns the list of ServicebusSubscriptionRules that match those selectors.
func (c *FakeServicebusSubscriptionRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ServicebusSubscriptionRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servicebussubscriptionrulesResource, servicebussubscriptionrulesKind, c.ns, opts), &v1alpha1.ServicebusSubscriptionRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ServicebusSubscriptionRuleList{ListMeta: obj.(*v1alpha1.ServicebusSubscriptionRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.ServicebusSubscriptionRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested servicebusSubscriptionRules.
func (c *FakeServicebusSubscriptionRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servicebussubscriptionrulesResource, c.ns, opts))

}

// Create takes the representation of a servicebusSubscriptionRule and creates it.  Returns the server's representation of the servicebusSubscriptionRule, and an error, if there is any.
func (c *FakeServicebusSubscriptionRules) Create(ctx context.Context, servicebusSubscriptionRule *v1alpha1.ServicebusSubscriptionRule, opts v1.CreateOptions) (result *v1alpha1.ServicebusSubscriptionRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servicebussubscriptionrulesResource, c.ns, servicebusSubscriptionRule), &v1alpha1.ServicebusSubscriptionRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicebusSubscriptionRule), err
}

// Update takes the representation of a servicebusSubscriptionRule and updates it. Returns the server's representation of the servicebusSubscriptionRule, and an error, if there is any.
func (c *FakeServicebusSubscriptionRules) Update(ctx context.Context, servicebusSubscriptionRule *v1alpha1.ServicebusSubscriptionRule, opts v1.UpdateOptions) (result *v1alpha1.ServicebusSubscriptionRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servicebussubscriptionrulesResource, c.ns, servicebusSubscriptionRule), &v1alpha1.ServicebusSubscriptionRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicebusSubscriptionRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServicebusSubscriptionRules) UpdateStatus(ctx context.Context, servicebusSubscriptionRule *v1alpha1.ServicebusSubscriptionRule, opts v1.UpdateOptions) (*v1alpha1.ServicebusSubscriptionRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(servicebussubscriptionrulesResource, "status", c.ns, servicebusSubscriptionRule), &v1alpha1.ServicebusSubscriptionRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicebusSubscriptionRule), err
}

// Delete takes name of the servicebusSubscriptionRule and deletes it. Returns an error if one occurs.
func (c *FakeServicebusSubscriptionRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(servicebussubscriptionrulesResource, c.ns, name), &v1alpha1.ServicebusSubscriptionRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServicebusSubscriptionRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servicebussubscriptionrulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ServicebusSubscriptionRuleList{})
	return err
}

// Patch applies the patch and returns the patched servicebusSubscriptionRule.
func (c *FakeServicebusSubscriptionRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServicebusSubscriptionRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servicebussubscriptionrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ServicebusSubscriptionRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicebusSubscriptionRule), err
}
