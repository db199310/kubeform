/*
Copyright 2019 The Kubeform Authors.

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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakeStorageDefaultObjectAcls implements StorageDefaultObjectAclInterface
type FakeStorageDefaultObjectAcls struct {
	Fake *FakeGoogleV1alpha1
}

var storagedefaultobjectaclsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "storagedefaultobjectacls"}

var storagedefaultobjectaclsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "StorageDefaultObjectAcl"}

// Get takes name of the storageDefaultObjectAcl, and returns the corresponding storageDefaultObjectAcl object, and an error if there is any.
func (c *FakeStorageDefaultObjectAcls) Get(name string, options v1.GetOptions) (result *v1alpha1.StorageDefaultObjectAcl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(storagedefaultobjectaclsResource, name), &v1alpha1.StorageDefaultObjectAcl{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageDefaultObjectAcl), err
}

// List takes label and field selectors, and returns the list of StorageDefaultObjectAcls that match those selectors.
func (c *FakeStorageDefaultObjectAcls) List(opts v1.ListOptions) (result *v1alpha1.StorageDefaultObjectAclList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(storagedefaultobjectaclsResource, storagedefaultobjectaclsKind, opts), &v1alpha1.StorageDefaultObjectAclList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StorageDefaultObjectAclList{ListMeta: obj.(*v1alpha1.StorageDefaultObjectAclList).ListMeta}
	for _, item := range obj.(*v1alpha1.StorageDefaultObjectAclList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested storageDefaultObjectAcls.
func (c *FakeStorageDefaultObjectAcls) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(storagedefaultobjectaclsResource, opts))
}

// Create takes the representation of a storageDefaultObjectAcl and creates it.  Returns the server's representation of the storageDefaultObjectAcl, and an error, if there is any.
func (c *FakeStorageDefaultObjectAcls) Create(storageDefaultObjectAcl *v1alpha1.StorageDefaultObjectAcl) (result *v1alpha1.StorageDefaultObjectAcl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(storagedefaultobjectaclsResource, storageDefaultObjectAcl), &v1alpha1.StorageDefaultObjectAcl{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageDefaultObjectAcl), err
}

// Update takes the representation of a storageDefaultObjectAcl and updates it. Returns the server's representation of the storageDefaultObjectAcl, and an error, if there is any.
func (c *FakeStorageDefaultObjectAcls) Update(storageDefaultObjectAcl *v1alpha1.StorageDefaultObjectAcl) (result *v1alpha1.StorageDefaultObjectAcl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(storagedefaultobjectaclsResource, storageDefaultObjectAcl), &v1alpha1.StorageDefaultObjectAcl{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageDefaultObjectAcl), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStorageDefaultObjectAcls) UpdateStatus(storageDefaultObjectAcl *v1alpha1.StorageDefaultObjectAcl) (*v1alpha1.StorageDefaultObjectAcl, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(storagedefaultobjectaclsResource, "status", storageDefaultObjectAcl), &v1alpha1.StorageDefaultObjectAcl{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageDefaultObjectAcl), err
}

// Delete takes name of the storageDefaultObjectAcl and deletes it. Returns an error if one occurs.
func (c *FakeStorageDefaultObjectAcls) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(storagedefaultobjectaclsResource, name), &v1alpha1.StorageDefaultObjectAcl{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStorageDefaultObjectAcls) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(storagedefaultobjectaclsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.StorageDefaultObjectAclList{})
	return err
}

// Patch applies the patch and returns the patched storageDefaultObjectAcl.
func (c *FakeStorageDefaultObjectAcls) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StorageDefaultObjectAcl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(storagedefaultobjectaclsResource, name, pt, data, subresources...), &v1alpha1.StorageDefaultObjectAcl{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageDefaultObjectAcl), err
}