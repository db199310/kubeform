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

// FakeCosmosdbTables implements CosmosdbTableInterface
type FakeCosmosdbTables struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var cosmosdbtablesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "cosmosdbtables"}

var cosmosdbtablesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "CosmosdbTable"}

// Get takes name of the cosmosdbTable, and returns the corresponding cosmosdbTable object, and an error if there is any.
func (c *FakeCosmosdbTables) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CosmosdbTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cosmosdbtablesResource, c.ns, name), &v1alpha1.CosmosdbTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CosmosdbTable), err
}

// List takes label and field selectors, and returns the list of CosmosdbTables that match those selectors.
func (c *FakeCosmosdbTables) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CosmosdbTableList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cosmosdbtablesResource, cosmosdbtablesKind, c.ns, opts), &v1alpha1.CosmosdbTableList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CosmosdbTableList{ListMeta: obj.(*v1alpha1.CosmosdbTableList).ListMeta}
	for _, item := range obj.(*v1alpha1.CosmosdbTableList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cosmosdbTables.
func (c *FakeCosmosdbTables) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cosmosdbtablesResource, c.ns, opts))

}

// Create takes the representation of a cosmosdbTable and creates it.  Returns the server's representation of the cosmosdbTable, and an error, if there is any.
func (c *FakeCosmosdbTables) Create(ctx context.Context, cosmosdbTable *v1alpha1.CosmosdbTable, opts v1.CreateOptions) (result *v1alpha1.CosmosdbTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cosmosdbtablesResource, c.ns, cosmosdbTable), &v1alpha1.CosmosdbTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CosmosdbTable), err
}

// Update takes the representation of a cosmosdbTable and updates it. Returns the server's representation of the cosmosdbTable, and an error, if there is any.
func (c *FakeCosmosdbTables) Update(ctx context.Context, cosmosdbTable *v1alpha1.CosmosdbTable, opts v1.UpdateOptions) (result *v1alpha1.CosmosdbTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cosmosdbtablesResource, c.ns, cosmosdbTable), &v1alpha1.CosmosdbTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CosmosdbTable), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCosmosdbTables) UpdateStatus(ctx context.Context, cosmosdbTable *v1alpha1.CosmosdbTable, opts v1.UpdateOptions) (*v1alpha1.CosmosdbTable, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cosmosdbtablesResource, "status", c.ns, cosmosdbTable), &v1alpha1.CosmosdbTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CosmosdbTable), err
}

// Delete takes name of the cosmosdbTable and deletes it. Returns an error if one occurs.
func (c *FakeCosmosdbTables) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cosmosdbtablesResource, c.ns, name), &v1alpha1.CosmosdbTable{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCosmosdbTables) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cosmosdbtablesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CosmosdbTableList{})
	return err
}

// Patch applies the patch and returns the patched cosmosdbTable.
func (c *FakeCosmosdbTables) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CosmosdbTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cosmosdbtablesResource, c.ns, name, pt, data, subresources...), &v1alpha1.CosmosdbTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CosmosdbTable), err
}
