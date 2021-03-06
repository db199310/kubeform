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

// FakeAdvancedThreatProtections implements AdvancedThreatProtectionInterface
type FakeAdvancedThreatProtections struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var advancedthreatprotectionsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "advancedthreatprotections"}

var advancedthreatprotectionsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "AdvancedThreatProtection"}

// Get takes name of the advancedThreatProtection, and returns the corresponding advancedThreatProtection object, and an error if there is any.
func (c *FakeAdvancedThreatProtections) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AdvancedThreatProtection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(advancedthreatprotectionsResource, c.ns, name), &v1alpha1.AdvancedThreatProtection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AdvancedThreatProtection), err
}

// List takes label and field selectors, and returns the list of AdvancedThreatProtections that match those selectors.
func (c *FakeAdvancedThreatProtections) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AdvancedThreatProtectionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(advancedthreatprotectionsResource, advancedthreatprotectionsKind, c.ns, opts), &v1alpha1.AdvancedThreatProtectionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AdvancedThreatProtectionList{ListMeta: obj.(*v1alpha1.AdvancedThreatProtectionList).ListMeta}
	for _, item := range obj.(*v1alpha1.AdvancedThreatProtectionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested advancedThreatProtections.
func (c *FakeAdvancedThreatProtections) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(advancedthreatprotectionsResource, c.ns, opts))

}

// Create takes the representation of a advancedThreatProtection and creates it.  Returns the server's representation of the advancedThreatProtection, and an error, if there is any.
func (c *FakeAdvancedThreatProtections) Create(ctx context.Context, advancedThreatProtection *v1alpha1.AdvancedThreatProtection, opts v1.CreateOptions) (result *v1alpha1.AdvancedThreatProtection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(advancedthreatprotectionsResource, c.ns, advancedThreatProtection), &v1alpha1.AdvancedThreatProtection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AdvancedThreatProtection), err
}

// Update takes the representation of a advancedThreatProtection and updates it. Returns the server's representation of the advancedThreatProtection, and an error, if there is any.
func (c *FakeAdvancedThreatProtections) Update(ctx context.Context, advancedThreatProtection *v1alpha1.AdvancedThreatProtection, opts v1.UpdateOptions) (result *v1alpha1.AdvancedThreatProtection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(advancedthreatprotectionsResource, c.ns, advancedThreatProtection), &v1alpha1.AdvancedThreatProtection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AdvancedThreatProtection), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAdvancedThreatProtections) UpdateStatus(ctx context.Context, advancedThreatProtection *v1alpha1.AdvancedThreatProtection, opts v1.UpdateOptions) (*v1alpha1.AdvancedThreatProtection, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(advancedthreatprotectionsResource, "status", c.ns, advancedThreatProtection), &v1alpha1.AdvancedThreatProtection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AdvancedThreatProtection), err
}

// Delete takes name of the advancedThreatProtection and deletes it. Returns an error if one occurs.
func (c *FakeAdvancedThreatProtections) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(advancedthreatprotectionsResource, c.ns, name), &v1alpha1.AdvancedThreatProtection{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAdvancedThreatProtections) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(advancedthreatprotectionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AdvancedThreatProtectionList{})
	return err
}

// Patch applies the patch and returns the patched advancedThreatProtection.
func (c *FakeAdvancedThreatProtections) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AdvancedThreatProtection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(advancedthreatprotectionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AdvancedThreatProtection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AdvancedThreatProtection), err
}
