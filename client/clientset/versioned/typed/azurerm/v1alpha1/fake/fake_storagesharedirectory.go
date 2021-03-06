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

// FakeStorageShareDirectories implements StorageShareDirectoryInterface
type FakeStorageShareDirectories struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var storagesharedirectoriesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "storagesharedirectories"}

var storagesharedirectoriesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "StorageShareDirectory"}

// Get takes name of the storageShareDirectory, and returns the corresponding storageShareDirectory object, and an error if there is any.
func (c *FakeStorageShareDirectories) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.StorageShareDirectory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(storagesharedirectoriesResource, c.ns, name), &v1alpha1.StorageShareDirectory{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageShareDirectory), err
}

// List takes label and field selectors, and returns the list of StorageShareDirectories that match those selectors.
func (c *FakeStorageShareDirectories) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.StorageShareDirectoryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(storagesharedirectoriesResource, storagesharedirectoriesKind, c.ns, opts), &v1alpha1.StorageShareDirectoryList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StorageShareDirectoryList{ListMeta: obj.(*v1alpha1.StorageShareDirectoryList).ListMeta}
	for _, item := range obj.(*v1alpha1.StorageShareDirectoryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested storageShareDirectories.
func (c *FakeStorageShareDirectories) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(storagesharedirectoriesResource, c.ns, opts))

}

// Create takes the representation of a storageShareDirectory and creates it.  Returns the server's representation of the storageShareDirectory, and an error, if there is any.
func (c *FakeStorageShareDirectories) Create(ctx context.Context, storageShareDirectory *v1alpha1.StorageShareDirectory, opts v1.CreateOptions) (result *v1alpha1.StorageShareDirectory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(storagesharedirectoriesResource, c.ns, storageShareDirectory), &v1alpha1.StorageShareDirectory{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageShareDirectory), err
}

// Update takes the representation of a storageShareDirectory and updates it. Returns the server's representation of the storageShareDirectory, and an error, if there is any.
func (c *FakeStorageShareDirectories) Update(ctx context.Context, storageShareDirectory *v1alpha1.StorageShareDirectory, opts v1.UpdateOptions) (result *v1alpha1.StorageShareDirectory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(storagesharedirectoriesResource, c.ns, storageShareDirectory), &v1alpha1.StorageShareDirectory{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageShareDirectory), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStorageShareDirectories) UpdateStatus(ctx context.Context, storageShareDirectory *v1alpha1.StorageShareDirectory, opts v1.UpdateOptions) (*v1alpha1.StorageShareDirectory, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(storagesharedirectoriesResource, "status", c.ns, storageShareDirectory), &v1alpha1.StorageShareDirectory{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageShareDirectory), err
}

// Delete takes name of the storageShareDirectory and deletes it. Returns an error if one occurs.
func (c *FakeStorageShareDirectories) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(storagesharedirectoriesResource, c.ns, name), &v1alpha1.StorageShareDirectory{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStorageShareDirectories) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(storagesharedirectoriesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.StorageShareDirectoryList{})
	return err
}

// Patch applies the patch and returns the patched storageShareDirectory.
func (c *FakeStorageShareDirectories) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.StorageShareDirectory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(storagesharedirectoriesResource, c.ns, name, pt, data, subresources...), &v1alpha1.StorageShareDirectory{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageShareDirectory), err
}
