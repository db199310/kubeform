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

// FakeMssqlVirtualMachines implements MssqlVirtualMachineInterface
type FakeMssqlVirtualMachines struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var mssqlvirtualmachinesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "mssqlvirtualmachines"}

var mssqlvirtualmachinesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "MssqlVirtualMachine"}

// Get takes name of the mssqlVirtualMachine, and returns the corresponding mssqlVirtualMachine object, and an error if there is any.
func (c *FakeMssqlVirtualMachines) Get(name string, options v1.GetOptions) (result *v1alpha1.MssqlVirtualMachine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mssqlvirtualmachinesResource, c.ns, name), &v1alpha1.MssqlVirtualMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MssqlVirtualMachine), err
}

// List takes label and field selectors, and returns the list of MssqlVirtualMachines that match those selectors.
func (c *FakeMssqlVirtualMachines) List(opts v1.ListOptions) (result *v1alpha1.MssqlVirtualMachineList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mssqlvirtualmachinesResource, mssqlvirtualmachinesKind, c.ns, opts), &v1alpha1.MssqlVirtualMachineList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MssqlVirtualMachineList{ListMeta: obj.(*v1alpha1.MssqlVirtualMachineList).ListMeta}
	for _, item := range obj.(*v1alpha1.MssqlVirtualMachineList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mssqlVirtualMachines.
func (c *FakeMssqlVirtualMachines) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mssqlvirtualmachinesResource, c.ns, opts))

}

// Create takes the representation of a mssqlVirtualMachine and creates it.  Returns the server's representation of the mssqlVirtualMachine, and an error, if there is any.
func (c *FakeMssqlVirtualMachines) Create(mssqlVirtualMachine *v1alpha1.MssqlVirtualMachine) (result *v1alpha1.MssqlVirtualMachine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mssqlvirtualmachinesResource, c.ns, mssqlVirtualMachine), &v1alpha1.MssqlVirtualMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MssqlVirtualMachine), err
}

// Update takes the representation of a mssqlVirtualMachine and updates it. Returns the server's representation of the mssqlVirtualMachine, and an error, if there is any.
func (c *FakeMssqlVirtualMachines) Update(mssqlVirtualMachine *v1alpha1.MssqlVirtualMachine) (result *v1alpha1.MssqlVirtualMachine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mssqlvirtualmachinesResource, c.ns, mssqlVirtualMachine), &v1alpha1.MssqlVirtualMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MssqlVirtualMachine), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMssqlVirtualMachines) UpdateStatus(mssqlVirtualMachine *v1alpha1.MssqlVirtualMachine) (*v1alpha1.MssqlVirtualMachine, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(mssqlvirtualmachinesResource, "status", c.ns, mssqlVirtualMachine), &v1alpha1.MssqlVirtualMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MssqlVirtualMachine), err
}

// Delete takes name of the mssqlVirtualMachine and deletes it. Returns an error if one occurs.
func (c *FakeMssqlVirtualMachines) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(mssqlvirtualmachinesResource, c.ns, name), &v1alpha1.MssqlVirtualMachine{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMssqlVirtualMachines) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mssqlvirtualmachinesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.MssqlVirtualMachineList{})
	return err
}

// Patch applies the patch and returns the patched mssqlVirtualMachine.
func (c *FakeMssqlVirtualMachines) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MssqlVirtualMachine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mssqlvirtualmachinesResource, c.ns, name, pt, data, subresources...), &v1alpha1.MssqlVirtualMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MssqlVirtualMachine), err
}
