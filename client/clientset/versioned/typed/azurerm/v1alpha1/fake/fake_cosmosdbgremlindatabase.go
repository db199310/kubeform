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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakeCosmosdbGremlinDatabases implements CosmosdbGremlinDatabaseInterface
type FakeCosmosdbGremlinDatabases struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var cosmosdbgremlindatabasesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "cosmosdbgremlindatabases"}

var cosmosdbgremlindatabasesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "CosmosdbGremlinDatabase"}

// Get takes name of the cosmosdbGremlinDatabase, and returns the corresponding cosmosdbGremlinDatabase object, and an error if there is any.
func (c *FakeCosmosdbGremlinDatabases) Get(name string, options v1.GetOptions) (result *v1alpha1.CosmosdbGremlinDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cosmosdbgremlindatabasesResource, c.ns, name), &v1alpha1.CosmosdbGremlinDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CosmosdbGremlinDatabase), err
}

// List takes label and field selectors, and returns the list of CosmosdbGremlinDatabases that match those selectors.
func (c *FakeCosmosdbGremlinDatabases) List(opts v1.ListOptions) (result *v1alpha1.CosmosdbGremlinDatabaseList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cosmosdbgremlindatabasesResource, cosmosdbgremlindatabasesKind, c.ns, opts), &v1alpha1.CosmosdbGremlinDatabaseList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CosmosdbGremlinDatabaseList{ListMeta: obj.(*v1alpha1.CosmosdbGremlinDatabaseList).ListMeta}
	for _, item := range obj.(*v1alpha1.CosmosdbGremlinDatabaseList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cosmosdbGremlinDatabases.
func (c *FakeCosmosdbGremlinDatabases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cosmosdbgremlindatabasesResource, c.ns, opts))

}

// Create takes the representation of a cosmosdbGremlinDatabase and creates it.  Returns the server's representation of the cosmosdbGremlinDatabase, and an error, if there is any.
func (c *FakeCosmosdbGremlinDatabases) Create(cosmosdbGremlinDatabase *v1alpha1.CosmosdbGremlinDatabase) (result *v1alpha1.CosmosdbGremlinDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cosmosdbgremlindatabasesResource, c.ns, cosmosdbGremlinDatabase), &v1alpha1.CosmosdbGremlinDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CosmosdbGremlinDatabase), err
}

// Update takes the representation of a cosmosdbGremlinDatabase and updates it. Returns the server's representation of the cosmosdbGremlinDatabase, and an error, if there is any.
func (c *FakeCosmosdbGremlinDatabases) Update(cosmosdbGremlinDatabase *v1alpha1.CosmosdbGremlinDatabase) (result *v1alpha1.CosmosdbGremlinDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cosmosdbgremlindatabasesResource, c.ns, cosmosdbGremlinDatabase), &v1alpha1.CosmosdbGremlinDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CosmosdbGremlinDatabase), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCosmosdbGremlinDatabases) UpdateStatus(cosmosdbGremlinDatabase *v1alpha1.CosmosdbGremlinDatabase) (*v1alpha1.CosmosdbGremlinDatabase, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cosmosdbgremlindatabasesResource, "status", c.ns, cosmosdbGremlinDatabase), &v1alpha1.CosmosdbGremlinDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CosmosdbGremlinDatabase), err
}

// Delete takes name of the cosmosdbGremlinDatabase and deletes it. Returns an error if one occurs.
func (c *FakeCosmosdbGremlinDatabases) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cosmosdbgremlindatabasesResource, c.ns, name), &v1alpha1.CosmosdbGremlinDatabase{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCosmosdbGremlinDatabases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cosmosdbgremlindatabasesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CosmosdbGremlinDatabaseList{})
	return err
}

// Patch applies the patch and returns the patched cosmosdbGremlinDatabase.
func (c *FakeCosmosdbGremlinDatabases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CosmosdbGremlinDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cosmosdbgremlindatabasesResource, c.ns, name, pt, data, subresources...), &v1alpha1.CosmosdbGremlinDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CosmosdbGremlinDatabase), err
}
