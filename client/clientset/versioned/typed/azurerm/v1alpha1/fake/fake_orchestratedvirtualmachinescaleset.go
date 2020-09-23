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

// FakeOrchestratedVirtualMachineScaleSets implements OrchestratedVirtualMachineScaleSetInterface
type FakeOrchestratedVirtualMachineScaleSets struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var orchestratedvirtualmachinescalesetsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "orchestratedvirtualmachinescalesets"}

var orchestratedvirtualmachinescalesetsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "OrchestratedVirtualMachineScaleSet"}

// Get takes name of the orchestratedVirtualMachineScaleSet, and returns the corresponding orchestratedVirtualMachineScaleSet object, and an error if there is any.
func (c *FakeOrchestratedVirtualMachineScaleSets) Get(name string, options v1.GetOptions) (result *v1alpha1.OrchestratedVirtualMachineScaleSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(orchestratedvirtualmachinescalesetsResource, c.ns, name), &v1alpha1.OrchestratedVirtualMachineScaleSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrchestratedVirtualMachineScaleSet), err
}

// List takes label and field selectors, and returns the list of OrchestratedVirtualMachineScaleSets that match those selectors.
func (c *FakeOrchestratedVirtualMachineScaleSets) List(opts v1.ListOptions) (result *v1alpha1.OrchestratedVirtualMachineScaleSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(orchestratedvirtualmachinescalesetsResource, orchestratedvirtualmachinescalesetsKind, c.ns, opts), &v1alpha1.OrchestratedVirtualMachineScaleSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OrchestratedVirtualMachineScaleSetList{ListMeta: obj.(*v1alpha1.OrchestratedVirtualMachineScaleSetList).ListMeta}
	for _, item := range obj.(*v1alpha1.OrchestratedVirtualMachineScaleSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested orchestratedVirtualMachineScaleSets.
func (c *FakeOrchestratedVirtualMachineScaleSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(orchestratedvirtualmachinescalesetsResource, c.ns, opts))

}

// Create takes the representation of a orchestratedVirtualMachineScaleSet and creates it.  Returns the server's representation of the orchestratedVirtualMachineScaleSet, and an error, if there is any.
func (c *FakeOrchestratedVirtualMachineScaleSets) Create(orchestratedVirtualMachineScaleSet *v1alpha1.OrchestratedVirtualMachineScaleSet) (result *v1alpha1.OrchestratedVirtualMachineScaleSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(orchestratedvirtualmachinescalesetsResource, c.ns, orchestratedVirtualMachineScaleSet), &v1alpha1.OrchestratedVirtualMachineScaleSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrchestratedVirtualMachineScaleSet), err
}

// Update takes the representation of a orchestratedVirtualMachineScaleSet and updates it. Returns the server's representation of the orchestratedVirtualMachineScaleSet, and an error, if there is any.
func (c *FakeOrchestratedVirtualMachineScaleSets) Update(orchestratedVirtualMachineScaleSet *v1alpha1.OrchestratedVirtualMachineScaleSet) (result *v1alpha1.OrchestratedVirtualMachineScaleSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(orchestratedvirtualmachinescalesetsResource, c.ns, orchestratedVirtualMachineScaleSet), &v1alpha1.OrchestratedVirtualMachineScaleSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrchestratedVirtualMachineScaleSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOrchestratedVirtualMachineScaleSets) UpdateStatus(orchestratedVirtualMachineScaleSet *v1alpha1.OrchestratedVirtualMachineScaleSet) (*v1alpha1.OrchestratedVirtualMachineScaleSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(orchestratedvirtualmachinescalesetsResource, "status", c.ns, orchestratedVirtualMachineScaleSet), &v1alpha1.OrchestratedVirtualMachineScaleSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrchestratedVirtualMachineScaleSet), err
}

// Delete takes name of the orchestratedVirtualMachineScaleSet and deletes it. Returns an error if one occurs.
func (c *FakeOrchestratedVirtualMachineScaleSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(orchestratedvirtualmachinescalesetsResource, c.ns, name), &v1alpha1.OrchestratedVirtualMachineScaleSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOrchestratedVirtualMachineScaleSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(orchestratedvirtualmachinescalesetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.OrchestratedVirtualMachineScaleSetList{})
	return err
}

// Patch applies the patch and returns the patched orchestratedVirtualMachineScaleSet.
func (c *FakeOrchestratedVirtualMachineScaleSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OrchestratedVirtualMachineScaleSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(orchestratedvirtualmachinescalesetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.OrchestratedVirtualMachineScaleSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrchestratedVirtualMachineScaleSet), err
}