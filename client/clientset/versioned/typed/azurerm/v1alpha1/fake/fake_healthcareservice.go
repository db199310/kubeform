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

// FakeHealthcareServices implements HealthcareServiceInterface
type FakeHealthcareServices struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var healthcareservicesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "healthcareservices"}

var healthcareservicesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "HealthcareService"}

// Get takes name of the healthcareService, and returns the corresponding healthcareService object, and an error if there is any.
func (c *FakeHealthcareServices) Get(name string, options v1.GetOptions) (result *v1alpha1.HealthcareService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(healthcareservicesResource, c.ns, name), &v1alpha1.HealthcareService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HealthcareService), err
}

// List takes label and field selectors, and returns the list of HealthcareServices that match those selectors.
func (c *FakeHealthcareServices) List(opts v1.ListOptions) (result *v1alpha1.HealthcareServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(healthcareservicesResource, healthcareservicesKind, c.ns, opts), &v1alpha1.HealthcareServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HealthcareServiceList{ListMeta: obj.(*v1alpha1.HealthcareServiceList).ListMeta}
	for _, item := range obj.(*v1alpha1.HealthcareServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested healthcareServices.
func (c *FakeHealthcareServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(healthcareservicesResource, c.ns, opts))

}

// Create takes the representation of a healthcareService and creates it.  Returns the server's representation of the healthcareService, and an error, if there is any.
func (c *FakeHealthcareServices) Create(healthcareService *v1alpha1.HealthcareService) (result *v1alpha1.HealthcareService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(healthcareservicesResource, c.ns, healthcareService), &v1alpha1.HealthcareService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HealthcareService), err
}

// Update takes the representation of a healthcareService and updates it. Returns the server's representation of the healthcareService, and an error, if there is any.
func (c *FakeHealthcareServices) Update(healthcareService *v1alpha1.HealthcareService) (result *v1alpha1.HealthcareService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(healthcareservicesResource, c.ns, healthcareService), &v1alpha1.HealthcareService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HealthcareService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHealthcareServices) UpdateStatus(healthcareService *v1alpha1.HealthcareService) (*v1alpha1.HealthcareService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(healthcareservicesResource, "status", c.ns, healthcareService), &v1alpha1.HealthcareService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HealthcareService), err
}

// Delete takes name of the healthcareService and deletes it. Returns an error if one occurs.
func (c *FakeHealthcareServices) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(healthcareservicesResource, c.ns, name), &v1alpha1.HealthcareService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHealthcareServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(healthcareservicesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.HealthcareServiceList{})
	return err
}

// Patch applies the patch and returns the patched healthcareService.
func (c *FakeHealthcareServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.HealthcareService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(healthcareservicesResource, c.ns, name, pt, data, subresources...), &v1alpha1.HealthcareService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HealthcareService), err
}
