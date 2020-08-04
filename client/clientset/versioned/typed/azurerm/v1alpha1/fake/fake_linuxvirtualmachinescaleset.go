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

// FakeLinuxVirtualMachineScaleSets implements LinuxVirtualMachineScaleSetInterface
type FakeLinuxVirtualMachineScaleSets struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var linuxvirtualmachinescalesetsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "linuxvirtualmachinescalesets"}

var linuxvirtualmachinescalesetsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "LinuxVirtualMachineScaleSet"}

// Get takes name of the linuxVirtualMachineScaleSet, and returns the corresponding linuxVirtualMachineScaleSet object, and an error if there is any.
func (c *FakeLinuxVirtualMachineScaleSets) Get(name string, options v1.GetOptions) (result *v1alpha1.LinuxVirtualMachineScaleSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(linuxvirtualmachinescalesetsResource, c.ns, name), &v1alpha1.LinuxVirtualMachineScaleSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LinuxVirtualMachineScaleSet), err
}

// List takes label and field selectors, and returns the list of LinuxVirtualMachineScaleSets that match those selectors.
func (c *FakeLinuxVirtualMachineScaleSets) List(opts v1.ListOptions) (result *v1alpha1.LinuxVirtualMachineScaleSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(linuxvirtualmachinescalesetsResource, linuxvirtualmachinescalesetsKind, c.ns, opts), &v1alpha1.LinuxVirtualMachineScaleSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LinuxVirtualMachineScaleSetList{ListMeta: obj.(*v1alpha1.LinuxVirtualMachineScaleSetList).ListMeta}
	for _, item := range obj.(*v1alpha1.LinuxVirtualMachineScaleSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested linuxVirtualMachineScaleSets.
func (c *FakeLinuxVirtualMachineScaleSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(linuxvirtualmachinescalesetsResource, c.ns, opts))

}

// Create takes the representation of a linuxVirtualMachineScaleSet and creates it.  Returns the server's representation of the linuxVirtualMachineScaleSet, and an error, if there is any.
func (c *FakeLinuxVirtualMachineScaleSets) Create(linuxVirtualMachineScaleSet *v1alpha1.LinuxVirtualMachineScaleSet) (result *v1alpha1.LinuxVirtualMachineScaleSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(linuxvirtualmachinescalesetsResource, c.ns, linuxVirtualMachineScaleSet), &v1alpha1.LinuxVirtualMachineScaleSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LinuxVirtualMachineScaleSet), err
}

// Update takes the representation of a linuxVirtualMachineScaleSet and updates it. Returns the server's representation of the linuxVirtualMachineScaleSet, and an error, if there is any.
func (c *FakeLinuxVirtualMachineScaleSets) Update(linuxVirtualMachineScaleSet *v1alpha1.LinuxVirtualMachineScaleSet) (result *v1alpha1.LinuxVirtualMachineScaleSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(linuxvirtualmachinescalesetsResource, c.ns, linuxVirtualMachineScaleSet), &v1alpha1.LinuxVirtualMachineScaleSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LinuxVirtualMachineScaleSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLinuxVirtualMachineScaleSets) UpdateStatus(linuxVirtualMachineScaleSet *v1alpha1.LinuxVirtualMachineScaleSet) (*v1alpha1.LinuxVirtualMachineScaleSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(linuxvirtualmachinescalesetsResource, "status", c.ns, linuxVirtualMachineScaleSet), &v1alpha1.LinuxVirtualMachineScaleSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LinuxVirtualMachineScaleSet), err
}

// Delete takes name of the linuxVirtualMachineScaleSet and deletes it. Returns an error if one occurs.
func (c *FakeLinuxVirtualMachineScaleSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(linuxvirtualmachinescalesetsResource, c.ns, name), &v1alpha1.LinuxVirtualMachineScaleSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLinuxVirtualMachineScaleSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(linuxvirtualmachinescalesetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.LinuxVirtualMachineScaleSetList{})
	return err
}

// Patch applies the patch and returns the patched linuxVirtualMachineScaleSet.
func (c *FakeLinuxVirtualMachineScaleSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LinuxVirtualMachineScaleSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(linuxvirtualmachinescalesetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.LinuxVirtualMachineScaleSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LinuxVirtualMachineScaleSet), err
}
