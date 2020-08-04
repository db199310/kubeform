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

// FakeDevTestGlobalVmShutdownSchedules implements DevTestGlobalVmShutdownScheduleInterface
type FakeDevTestGlobalVmShutdownSchedules struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var devtestglobalvmshutdownschedulesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "devtestglobalvmshutdownschedules"}

var devtestglobalvmshutdownschedulesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "DevTestGlobalVmShutdownSchedule"}

// Get takes name of the devTestGlobalVmShutdownSchedule, and returns the corresponding devTestGlobalVmShutdownSchedule object, and an error if there is any.
func (c *FakeDevTestGlobalVmShutdownSchedules) Get(name string, options v1.GetOptions) (result *v1alpha1.DevTestGlobalVmShutdownSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(devtestglobalvmshutdownschedulesResource, c.ns, name), &v1alpha1.DevTestGlobalVmShutdownSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevTestGlobalVmShutdownSchedule), err
}

// List takes label and field selectors, and returns the list of DevTestGlobalVmShutdownSchedules that match those selectors.
func (c *FakeDevTestGlobalVmShutdownSchedules) List(opts v1.ListOptions) (result *v1alpha1.DevTestGlobalVmShutdownScheduleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(devtestglobalvmshutdownschedulesResource, devtestglobalvmshutdownschedulesKind, c.ns, opts), &v1alpha1.DevTestGlobalVmShutdownScheduleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DevTestGlobalVmShutdownScheduleList{ListMeta: obj.(*v1alpha1.DevTestGlobalVmShutdownScheduleList).ListMeta}
	for _, item := range obj.(*v1alpha1.DevTestGlobalVmShutdownScheduleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested devTestGlobalVmShutdownSchedules.
func (c *FakeDevTestGlobalVmShutdownSchedules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(devtestglobalvmshutdownschedulesResource, c.ns, opts))

}

// Create takes the representation of a devTestGlobalVmShutdownSchedule and creates it.  Returns the server's representation of the devTestGlobalVmShutdownSchedule, and an error, if there is any.
func (c *FakeDevTestGlobalVmShutdownSchedules) Create(devTestGlobalVmShutdownSchedule *v1alpha1.DevTestGlobalVmShutdownSchedule) (result *v1alpha1.DevTestGlobalVmShutdownSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(devtestglobalvmshutdownschedulesResource, c.ns, devTestGlobalVmShutdownSchedule), &v1alpha1.DevTestGlobalVmShutdownSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevTestGlobalVmShutdownSchedule), err
}

// Update takes the representation of a devTestGlobalVmShutdownSchedule and updates it. Returns the server's representation of the devTestGlobalVmShutdownSchedule, and an error, if there is any.
func (c *FakeDevTestGlobalVmShutdownSchedules) Update(devTestGlobalVmShutdownSchedule *v1alpha1.DevTestGlobalVmShutdownSchedule) (result *v1alpha1.DevTestGlobalVmShutdownSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(devtestglobalvmshutdownschedulesResource, c.ns, devTestGlobalVmShutdownSchedule), &v1alpha1.DevTestGlobalVmShutdownSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevTestGlobalVmShutdownSchedule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDevTestGlobalVmShutdownSchedules) UpdateStatus(devTestGlobalVmShutdownSchedule *v1alpha1.DevTestGlobalVmShutdownSchedule) (*v1alpha1.DevTestGlobalVmShutdownSchedule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(devtestglobalvmshutdownschedulesResource, "status", c.ns, devTestGlobalVmShutdownSchedule), &v1alpha1.DevTestGlobalVmShutdownSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevTestGlobalVmShutdownSchedule), err
}

// Delete takes name of the devTestGlobalVmShutdownSchedule and deletes it. Returns an error if one occurs.
func (c *FakeDevTestGlobalVmShutdownSchedules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(devtestglobalvmshutdownschedulesResource, c.ns, name), &v1alpha1.DevTestGlobalVmShutdownSchedule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDevTestGlobalVmShutdownSchedules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(devtestglobalvmshutdownschedulesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DevTestGlobalVmShutdownScheduleList{})
	return err
}

// Patch applies the patch and returns the patched devTestGlobalVmShutdownSchedule.
func (c *FakeDevTestGlobalVmShutdownSchedules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DevTestGlobalVmShutdownSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(devtestglobalvmshutdownschedulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.DevTestGlobalVmShutdownSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevTestGlobalVmShutdownSchedule), err
}
