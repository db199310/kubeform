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

// FakeDataLakeStores implements DataLakeStoreInterface
type FakeDataLakeStores struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var datalakestoresResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "datalakestores"}

var datalakestoresKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "DataLakeStore"}

// Get takes name of the dataLakeStore, and returns the corresponding dataLakeStore object, and an error if there is any.
func (c *FakeDataLakeStores) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DataLakeStore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(datalakestoresResource, c.ns, name), &v1alpha1.DataLakeStore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataLakeStore), err
}

// List takes label and field selectors, and returns the list of DataLakeStores that match those selectors.
func (c *FakeDataLakeStores) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DataLakeStoreList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(datalakestoresResource, datalakestoresKind, c.ns, opts), &v1alpha1.DataLakeStoreList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DataLakeStoreList{ListMeta: obj.(*v1alpha1.DataLakeStoreList).ListMeta}
	for _, item := range obj.(*v1alpha1.DataLakeStoreList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dataLakeStores.
func (c *FakeDataLakeStores) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(datalakestoresResource, c.ns, opts))

}

// Create takes the representation of a dataLakeStore and creates it.  Returns the server's representation of the dataLakeStore, and an error, if there is any.
func (c *FakeDataLakeStores) Create(ctx context.Context, dataLakeStore *v1alpha1.DataLakeStore, opts v1.CreateOptions) (result *v1alpha1.DataLakeStore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(datalakestoresResource, c.ns, dataLakeStore), &v1alpha1.DataLakeStore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataLakeStore), err
}

// Update takes the representation of a dataLakeStore and updates it. Returns the server's representation of the dataLakeStore, and an error, if there is any.
func (c *FakeDataLakeStores) Update(ctx context.Context, dataLakeStore *v1alpha1.DataLakeStore, opts v1.UpdateOptions) (result *v1alpha1.DataLakeStore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(datalakestoresResource, c.ns, dataLakeStore), &v1alpha1.DataLakeStore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataLakeStore), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDataLakeStores) UpdateStatus(ctx context.Context, dataLakeStore *v1alpha1.DataLakeStore, opts v1.UpdateOptions) (*v1alpha1.DataLakeStore, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(datalakestoresResource, "status", c.ns, dataLakeStore), &v1alpha1.DataLakeStore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataLakeStore), err
}

// Delete takes name of the dataLakeStore and deletes it. Returns an error if one occurs.
func (c *FakeDataLakeStores) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(datalakestoresResource, c.ns, name), &v1alpha1.DataLakeStore{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDataLakeStores) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(datalakestoresResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DataLakeStoreList{})
	return err
}

// Patch applies the patch and returns the patched dataLakeStore.
func (c *FakeDataLakeStores) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DataLakeStore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(datalakestoresResource, c.ns, name, pt, data, subresources...), &v1alpha1.DataLakeStore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataLakeStore), err
}
