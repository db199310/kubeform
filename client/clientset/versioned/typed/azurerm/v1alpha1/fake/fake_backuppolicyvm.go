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

// FakeBackupPolicyVms implements BackupPolicyVmInterface
type FakeBackupPolicyVms struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var backuppolicyvmsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "backuppolicyvms"}

var backuppolicyvmsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "BackupPolicyVm"}

// Get takes name of the backupPolicyVm, and returns the corresponding backupPolicyVm object, and an error if there is any.
func (c *FakeBackupPolicyVms) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BackupPolicyVm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(backuppolicyvmsResource, c.ns, name), &v1alpha1.BackupPolicyVm{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupPolicyVm), err
}

// List takes label and field selectors, and returns the list of BackupPolicyVms that match those selectors.
func (c *FakeBackupPolicyVms) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BackupPolicyVmList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(backuppolicyvmsResource, backuppolicyvmsKind, c.ns, opts), &v1alpha1.BackupPolicyVmList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BackupPolicyVmList{ListMeta: obj.(*v1alpha1.BackupPolicyVmList).ListMeta}
	for _, item := range obj.(*v1alpha1.BackupPolicyVmList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested backupPolicyVms.
func (c *FakeBackupPolicyVms) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(backuppolicyvmsResource, c.ns, opts))

}

// Create takes the representation of a backupPolicyVm and creates it.  Returns the server's representation of the backupPolicyVm, and an error, if there is any.
func (c *FakeBackupPolicyVms) Create(ctx context.Context, backupPolicyVm *v1alpha1.BackupPolicyVm, opts v1.CreateOptions) (result *v1alpha1.BackupPolicyVm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(backuppolicyvmsResource, c.ns, backupPolicyVm), &v1alpha1.BackupPolicyVm{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupPolicyVm), err
}

// Update takes the representation of a backupPolicyVm and updates it. Returns the server's representation of the backupPolicyVm, and an error, if there is any.
func (c *FakeBackupPolicyVms) Update(ctx context.Context, backupPolicyVm *v1alpha1.BackupPolicyVm, opts v1.UpdateOptions) (result *v1alpha1.BackupPolicyVm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(backuppolicyvmsResource, c.ns, backupPolicyVm), &v1alpha1.BackupPolicyVm{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupPolicyVm), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBackupPolicyVms) UpdateStatus(ctx context.Context, backupPolicyVm *v1alpha1.BackupPolicyVm, opts v1.UpdateOptions) (*v1alpha1.BackupPolicyVm, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(backuppolicyvmsResource, "status", c.ns, backupPolicyVm), &v1alpha1.BackupPolicyVm{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupPolicyVm), err
}

// Delete takes name of the backupPolicyVm and deletes it. Returns an error if one occurs.
func (c *FakeBackupPolicyVms) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(backuppolicyvmsResource, c.ns, name), &v1alpha1.BackupPolicyVm{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBackupPolicyVms) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(backuppolicyvmsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BackupPolicyVmList{})
	return err
}

// Patch applies the patch and returns the patched backupPolicyVm.
func (c *FakeBackupPolicyVms) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BackupPolicyVm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(backuppolicyvmsResource, c.ns, name, pt, data, subresources...), &v1alpha1.BackupPolicyVm{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupPolicyVm), err
}
