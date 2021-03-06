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

// FakeMarketplaceAgreements implements MarketplaceAgreementInterface
type FakeMarketplaceAgreements struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var marketplaceagreementsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "marketplaceagreements"}

var marketplaceagreementsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "MarketplaceAgreement"}

// Get takes name of the marketplaceAgreement, and returns the corresponding marketplaceAgreement object, and an error if there is any.
func (c *FakeMarketplaceAgreements) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MarketplaceAgreement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(marketplaceagreementsResource, c.ns, name), &v1alpha1.MarketplaceAgreement{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MarketplaceAgreement), err
}

// List takes label and field selectors, and returns the list of MarketplaceAgreements that match those selectors.
func (c *FakeMarketplaceAgreements) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MarketplaceAgreementList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(marketplaceagreementsResource, marketplaceagreementsKind, c.ns, opts), &v1alpha1.MarketplaceAgreementList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MarketplaceAgreementList{ListMeta: obj.(*v1alpha1.MarketplaceAgreementList).ListMeta}
	for _, item := range obj.(*v1alpha1.MarketplaceAgreementList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested marketplaceAgreements.
func (c *FakeMarketplaceAgreements) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(marketplaceagreementsResource, c.ns, opts))

}

// Create takes the representation of a marketplaceAgreement and creates it.  Returns the server's representation of the marketplaceAgreement, and an error, if there is any.
func (c *FakeMarketplaceAgreements) Create(ctx context.Context, marketplaceAgreement *v1alpha1.MarketplaceAgreement, opts v1.CreateOptions) (result *v1alpha1.MarketplaceAgreement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(marketplaceagreementsResource, c.ns, marketplaceAgreement), &v1alpha1.MarketplaceAgreement{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MarketplaceAgreement), err
}

// Update takes the representation of a marketplaceAgreement and updates it. Returns the server's representation of the marketplaceAgreement, and an error, if there is any.
func (c *FakeMarketplaceAgreements) Update(ctx context.Context, marketplaceAgreement *v1alpha1.MarketplaceAgreement, opts v1.UpdateOptions) (result *v1alpha1.MarketplaceAgreement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(marketplaceagreementsResource, c.ns, marketplaceAgreement), &v1alpha1.MarketplaceAgreement{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MarketplaceAgreement), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMarketplaceAgreements) UpdateStatus(ctx context.Context, marketplaceAgreement *v1alpha1.MarketplaceAgreement, opts v1.UpdateOptions) (*v1alpha1.MarketplaceAgreement, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(marketplaceagreementsResource, "status", c.ns, marketplaceAgreement), &v1alpha1.MarketplaceAgreement{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MarketplaceAgreement), err
}

// Delete takes name of the marketplaceAgreement and deletes it. Returns an error if one occurs.
func (c *FakeMarketplaceAgreements) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(marketplaceagreementsResource, c.ns, name), &v1alpha1.MarketplaceAgreement{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMarketplaceAgreements) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(marketplaceagreementsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MarketplaceAgreementList{})
	return err
}

// Patch applies the patch and returns the patched marketplaceAgreement.
func (c *FakeMarketplaceAgreements) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MarketplaceAgreement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(marketplaceagreementsResource, c.ns, name, pt, data, subresources...), &v1alpha1.MarketplaceAgreement{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MarketplaceAgreement), err
}
