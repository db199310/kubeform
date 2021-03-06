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

// FakeDataFactoryDatasetSQLServerTables implements DataFactoryDatasetSQLServerTableInterface
type FakeDataFactoryDatasetSQLServerTables struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var datafactorydatasetsqlservertablesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "datafactorydatasetsqlservertables"}

var datafactorydatasetsqlservertablesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "DataFactoryDatasetSQLServerTable"}

// Get takes name of the dataFactoryDatasetSQLServerTable, and returns the corresponding dataFactoryDatasetSQLServerTable object, and an error if there is any.
func (c *FakeDataFactoryDatasetSQLServerTables) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DataFactoryDatasetSQLServerTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(datafactorydatasetsqlservertablesResource, c.ns, name), &v1alpha1.DataFactoryDatasetSQLServerTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataFactoryDatasetSQLServerTable), err
}

// List takes label and field selectors, and returns the list of DataFactoryDatasetSQLServerTables that match those selectors.
func (c *FakeDataFactoryDatasetSQLServerTables) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DataFactoryDatasetSQLServerTableList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(datafactorydatasetsqlservertablesResource, datafactorydatasetsqlservertablesKind, c.ns, opts), &v1alpha1.DataFactoryDatasetSQLServerTableList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DataFactoryDatasetSQLServerTableList{ListMeta: obj.(*v1alpha1.DataFactoryDatasetSQLServerTableList).ListMeta}
	for _, item := range obj.(*v1alpha1.DataFactoryDatasetSQLServerTableList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dataFactoryDatasetSQLServerTables.
func (c *FakeDataFactoryDatasetSQLServerTables) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(datafactorydatasetsqlservertablesResource, c.ns, opts))

}

// Create takes the representation of a dataFactoryDatasetSQLServerTable and creates it.  Returns the server's representation of the dataFactoryDatasetSQLServerTable, and an error, if there is any.
func (c *FakeDataFactoryDatasetSQLServerTables) Create(ctx context.Context, dataFactoryDatasetSQLServerTable *v1alpha1.DataFactoryDatasetSQLServerTable, opts v1.CreateOptions) (result *v1alpha1.DataFactoryDatasetSQLServerTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(datafactorydatasetsqlservertablesResource, c.ns, dataFactoryDatasetSQLServerTable), &v1alpha1.DataFactoryDatasetSQLServerTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataFactoryDatasetSQLServerTable), err
}

// Update takes the representation of a dataFactoryDatasetSQLServerTable and updates it. Returns the server's representation of the dataFactoryDatasetSQLServerTable, and an error, if there is any.
func (c *FakeDataFactoryDatasetSQLServerTables) Update(ctx context.Context, dataFactoryDatasetSQLServerTable *v1alpha1.DataFactoryDatasetSQLServerTable, opts v1.UpdateOptions) (result *v1alpha1.DataFactoryDatasetSQLServerTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(datafactorydatasetsqlservertablesResource, c.ns, dataFactoryDatasetSQLServerTable), &v1alpha1.DataFactoryDatasetSQLServerTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataFactoryDatasetSQLServerTable), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDataFactoryDatasetSQLServerTables) UpdateStatus(ctx context.Context, dataFactoryDatasetSQLServerTable *v1alpha1.DataFactoryDatasetSQLServerTable, opts v1.UpdateOptions) (*v1alpha1.DataFactoryDatasetSQLServerTable, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(datafactorydatasetsqlservertablesResource, "status", c.ns, dataFactoryDatasetSQLServerTable), &v1alpha1.DataFactoryDatasetSQLServerTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataFactoryDatasetSQLServerTable), err
}

// Delete takes name of the dataFactoryDatasetSQLServerTable and deletes it. Returns an error if one occurs.
func (c *FakeDataFactoryDatasetSQLServerTables) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(datafactorydatasetsqlservertablesResource, c.ns, name), &v1alpha1.DataFactoryDatasetSQLServerTable{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDataFactoryDatasetSQLServerTables) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(datafactorydatasetsqlservertablesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DataFactoryDatasetSQLServerTableList{})
	return err
}

// Patch applies the patch and returns the patched dataFactoryDatasetSQLServerTable.
func (c *FakeDataFactoryDatasetSQLServerTables) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DataFactoryDatasetSQLServerTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(datafactorydatasetsqlservertablesResource, c.ns, name, pt, data, subresources...), &v1alpha1.DataFactoryDatasetSQLServerTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataFactoryDatasetSQLServerTable), err
}
