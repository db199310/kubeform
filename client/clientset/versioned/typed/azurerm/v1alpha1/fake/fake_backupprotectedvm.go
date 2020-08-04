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

// FakeBackupProtectedVms implements BackupProtectedVmInterface
type FakeBackupProtectedVms struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var backupprotectedvmsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "backupprotectedvms"}

var backupprotectedvmsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "BackupProtectedVm"}

// Get takes name of the backupProtectedVm, and returns the corresponding backupProtectedVm object, and an error if there is any.
func (c *FakeBackupProtectedVms) Get(name string, options v1.GetOptions) (result *v1alpha1.BackupProtectedVm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(backupprotectedvmsResource, c.ns, name), &v1alpha1.BackupProtectedVm{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupProtectedVm), err
}

// List takes label and field selectors, and returns the list of BackupProtectedVms that match those selectors.
func (c *FakeBackupProtectedVms) List(opts v1.ListOptions) (result *v1alpha1.BackupProtectedVmList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(backupprotectedvmsResource, backupprotectedvmsKind, c.ns, opts), &v1alpha1.BackupProtectedVmList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BackupProtectedVmList{ListMeta: obj.(*v1alpha1.BackupProtectedVmList).ListMeta}
	for _, item := range obj.(*v1alpha1.BackupProtectedVmList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested backupProtectedVms.
func (c *FakeBackupProtectedVms) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(backupprotectedvmsResource, c.ns, opts))

}

// Create takes the representation of a backupProtectedVm and creates it.  Returns the server's representation of the backupProtectedVm, and an error, if there is any.
func (c *FakeBackupProtectedVms) Create(backupProtectedVm *v1alpha1.BackupProtectedVm) (result *v1alpha1.BackupProtectedVm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(backupprotectedvmsResource, c.ns, backupProtectedVm), &v1alpha1.BackupProtectedVm{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupProtectedVm), err
}

// Update takes the representation of a backupProtectedVm and updates it. Returns the server's representation of the backupProtectedVm, and an error, if there is any.
func (c *FakeBackupProtectedVms) Update(backupProtectedVm *v1alpha1.BackupProtectedVm) (result *v1alpha1.BackupProtectedVm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(backupprotectedvmsResource, c.ns, backupProtectedVm), &v1alpha1.BackupProtectedVm{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupProtectedVm), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBackupProtectedVms) UpdateStatus(backupProtectedVm *v1alpha1.BackupProtectedVm) (*v1alpha1.BackupProtectedVm, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(backupprotectedvmsResource, "status", c.ns, backupProtectedVm), &v1alpha1.BackupProtectedVm{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupProtectedVm), err
}

// Delete takes name of the backupProtectedVm and deletes it. Returns an error if one occurs.
func (c *FakeBackupProtectedVms) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(backupprotectedvmsResource, c.ns, name), &v1alpha1.BackupProtectedVm{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBackupProtectedVms) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(backupprotectedvmsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.BackupProtectedVmList{})
	return err
}

// Patch applies the patch and returns the patched backupProtectedVm.
func (c *FakeBackupProtectedVms) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BackupProtectedVm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(backupprotectedvmsResource, c.ns, name, pt, data, subresources...), &v1alpha1.BackupProtectedVm{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupProtectedVm), err
}
