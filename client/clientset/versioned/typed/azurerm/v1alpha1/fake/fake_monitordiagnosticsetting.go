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

// FakeMonitorDiagnosticSettings implements MonitorDiagnosticSettingInterface
type FakeMonitorDiagnosticSettings struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var monitordiagnosticsettingsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "monitordiagnosticsettings"}

var monitordiagnosticsettingsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "MonitorDiagnosticSetting"}

// Get takes name of the monitorDiagnosticSetting, and returns the corresponding monitorDiagnosticSetting object, and an error if there is any.
func (c *FakeMonitorDiagnosticSettings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MonitorDiagnosticSetting, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(monitordiagnosticsettingsResource, c.ns, name), &v1alpha1.MonitorDiagnosticSetting{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitorDiagnosticSetting), err
}

// List takes label and field selectors, and returns the list of MonitorDiagnosticSettings that match those selectors.
func (c *FakeMonitorDiagnosticSettings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MonitorDiagnosticSettingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(monitordiagnosticsettingsResource, monitordiagnosticsettingsKind, c.ns, opts), &v1alpha1.MonitorDiagnosticSettingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MonitorDiagnosticSettingList{ListMeta: obj.(*v1alpha1.MonitorDiagnosticSettingList).ListMeta}
	for _, item := range obj.(*v1alpha1.MonitorDiagnosticSettingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested monitorDiagnosticSettings.
func (c *FakeMonitorDiagnosticSettings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(monitordiagnosticsettingsResource, c.ns, opts))

}

// Create takes the representation of a monitorDiagnosticSetting and creates it.  Returns the server's representation of the monitorDiagnosticSetting, and an error, if there is any.
func (c *FakeMonitorDiagnosticSettings) Create(ctx context.Context, monitorDiagnosticSetting *v1alpha1.MonitorDiagnosticSetting, opts v1.CreateOptions) (result *v1alpha1.MonitorDiagnosticSetting, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(monitordiagnosticsettingsResource, c.ns, monitorDiagnosticSetting), &v1alpha1.MonitorDiagnosticSetting{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitorDiagnosticSetting), err
}

// Update takes the representation of a monitorDiagnosticSetting and updates it. Returns the server's representation of the monitorDiagnosticSetting, and an error, if there is any.
func (c *FakeMonitorDiagnosticSettings) Update(ctx context.Context, monitorDiagnosticSetting *v1alpha1.MonitorDiagnosticSetting, opts v1.UpdateOptions) (result *v1alpha1.MonitorDiagnosticSetting, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(monitordiagnosticsettingsResource, c.ns, monitorDiagnosticSetting), &v1alpha1.MonitorDiagnosticSetting{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitorDiagnosticSetting), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMonitorDiagnosticSettings) UpdateStatus(ctx context.Context, monitorDiagnosticSetting *v1alpha1.MonitorDiagnosticSetting, opts v1.UpdateOptions) (*v1alpha1.MonitorDiagnosticSetting, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(monitordiagnosticsettingsResource, "status", c.ns, monitorDiagnosticSetting), &v1alpha1.MonitorDiagnosticSetting{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitorDiagnosticSetting), err
}

// Delete takes name of the monitorDiagnosticSetting and deletes it. Returns an error if one occurs.
func (c *FakeMonitorDiagnosticSettings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(monitordiagnosticsettingsResource, c.ns, name), &v1alpha1.MonitorDiagnosticSetting{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMonitorDiagnosticSettings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(monitordiagnosticsettingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MonitorDiagnosticSettingList{})
	return err
}

// Patch applies the patch and returns the patched monitorDiagnosticSetting.
func (c *FakeMonitorDiagnosticSettings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MonitorDiagnosticSetting, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(monitordiagnosticsettingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.MonitorDiagnosticSetting{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitorDiagnosticSetting), err
}
