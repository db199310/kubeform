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

// FakeServicebusNamespaceAuthorizationRules implements ServicebusNamespaceAuthorizationRuleInterface
type FakeServicebusNamespaceAuthorizationRules struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var servicebusnamespaceauthorizationrulesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "servicebusnamespaceauthorizationrules"}

var servicebusnamespaceauthorizationrulesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "ServicebusNamespaceAuthorizationRule"}

// Get takes name of the servicebusNamespaceAuthorizationRule, and returns the corresponding servicebusNamespaceAuthorizationRule object, and an error if there is any.
func (c *FakeServicebusNamespaceAuthorizationRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ServicebusNamespaceAuthorizationRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servicebusnamespaceauthorizationrulesResource, c.ns, name), &v1alpha1.ServicebusNamespaceAuthorizationRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicebusNamespaceAuthorizationRule), err
}

// List takes label and field selectors, and returns the list of ServicebusNamespaceAuthorizationRules that match those selectors.
func (c *FakeServicebusNamespaceAuthorizationRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ServicebusNamespaceAuthorizationRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servicebusnamespaceauthorizationrulesResource, servicebusnamespaceauthorizationrulesKind, c.ns, opts), &v1alpha1.ServicebusNamespaceAuthorizationRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ServicebusNamespaceAuthorizationRuleList{ListMeta: obj.(*v1alpha1.ServicebusNamespaceAuthorizationRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.ServicebusNamespaceAuthorizationRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested servicebusNamespaceAuthorizationRules.
func (c *FakeServicebusNamespaceAuthorizationRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servicebusnamespaceauthorizationrulesResource, c.ns, opts))

}

// Create takes the representation of a servicebusNamespaceAuthorizationRule and creates it.  Returns the server's representation of the servicebusNamespaceAuthorizationRule, and an error, if there is any.
func (c *FakeServicebusNamespaceAuthorizationRules) Create(ctx context.Context, servicebusNamespaceAuthorizationRule *v1alpha1.ServicebusNamespaceAuthorizationRule, opts v1.CreateOptions) (result *v1alpha1.ServicebusNamespaceAuthorizationRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servicebusnamespaceauthorizationrulesResource, c.ns, servicebusNamespaceAuthorizationRule), &v1alpha1.ServicebusNamespaceAuthorizationRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicebusNamespaceAuthorizationRule), err
}

// Update takes the representation of a servicebusNamespaceAuthorizationRule and updates it. Returns the server's representation of the servicebusNamespaceAuthorizationRule, and an error, if there is any.
func (c *FakeServicebusNamespaceAuthorizationRules) Update(ctx context.Context, servicebusNamespaceAuthorizationRule *v1alpha1.ServicebusNamespaceAuthorizationRule, opts v1.UpdateOptions) (result *v1alpha1.ServicebusNamespaceAuthorizationRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servicebusnamespaceauthorizationrulesResource, c.ns, servicebusNamespaceAuthorizationRule), &v1alpha1.ServicebusNamespaceAuthorizationRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicebusNamespaceAuthorizationRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServicebusNamespaceAuthorizationRules) UpdateStatus(ctx context.Context, servicebusNamespaceAuthorizationRule *v1alpha1.ServicebusNamespaceAuthorizationRule, opts v1.UpdateOptions) (*v1alpha1.ServicebusNamespaceAuthorizationRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(servicebusnamespaceauthorizationrulesResource, "status", c.ns, servicebusNamespaceAuthorizationRule), &v1alpha1.ServicebusNamespaceAuthorizationRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicebusNamespaceAuthorizationRule), err
}

// Delete takes name of the servicebusNamespaceAuthorizationRule and deletes it. Returns an error if one occurs.
func (c *FakeServicebusNamespaceAuthorizationRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(servicebusnamespaceauthorizationrulesResource, c.ns, name), &v1alpha1.ServicebusNamespaceAuthorizationRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServicebusNamespaceAuthorizationRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servicebusnamespaceauthorizationrulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ServicebusNamespaceAuthorizationRuleList{})
	return err
}

// Patch applies the patch and returns the patched servicebusNamespaceAuthorizationRule.
func (c *FakeServicebusNamespaceAuthorizationRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServicebusNamespaceAuthorizationRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servicebusnamespaceauthorizationrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ServicebusNamespaceAuthorizationRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicebusNamespaceAuthorizationRule), err
}
