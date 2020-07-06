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

// FakeApplicationInsightses implements ApplicationInsightsInterface
type FakeApplicationInsightses struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var applicationinsightsesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "applicationinsightses"}

var applicationinsightsesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "ApplicationInsights"}

// Get takes name of the applicationInsights, and returns the corresponding applicationInsights object, and an error if there is any.
func (c *FakeApplicationInsightses) Get(name string, options v1.GetOptions) (result *v1alpha1.ApplicationInsights, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(applicationinsightsesResource, c.ns, name), &v1alpha1.ApplicationInsights{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationInsights), err
}

// List takes label and field selectors, and returns the list of ApplicationInsightses that match those selectors.
func (c *FakeApplicationInsightses) List(opts v1.ListOptions) (result *v1alpha1.ApplicationInsightsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(applicationinsightsesResource, applicationinsightsesKind, c.ns, opts), &v1alpha1.ApplicationInsightsList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApplicationInsightsList{ListMeta: obj.(*v1alpha1.ApplicationInsightsList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApplicationInsightsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested applicationInsightses.
func (c *FakeApplicationInsightses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(applicationinsightsesResource, c.ns, opts))

}

// Create takes the representation of a applicationInsights and creates it.  Returns the server's representation of the applicationInsights, and an error, if there is any.
func (c *FakeApplicationInsightses) Create(applicationInsights *v1alpha1.ApplicationInsights) (result *v1alpha1.ApplicationInsights, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(applicationinsightsesResource, c.ns, applicationInsights), &v1alpha1.ApplicationInsights{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationInsights), err
}

// Update takes the representation of a applicationInsights and updates it. Returns the server's representation of the applicationInsights, and an error, if there is any.
func (c *FakeApplicationInsightses) Update(applicationInsights *v1alpha1.ApplicationInsights) (result *v1alpha1.ApplicationInsights, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(applicationinsightsesResource, c.ns, applicationInsights), &v1alpha1.ApplicationInsights{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationInsights), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApplicationInsightses) UpdateStatus(applicationInsights *v1alpha1.ApplicationInsights) (*v1alpha1.ApplicationInsights, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(applicationinsightsesResource, "status", c.ns, applicationInsights), &v1alpha1.ApplicationInsights{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationInsights), err
}

// Delete takes name of the applicationInsights and deletes it. Returns an error if one occurs.
func (c *FakeApplicationInsightses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(applicationinsightsesResource, c.ns, name), &v1alpha1.ApplicationInsights{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApplicationInsightses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(applicationinsightsesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApplicationInsightsList{})
	return err
}

// Patch applies the patch and returns the patched applicationInsights.
func (c *FakeApplicationInsightses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApplicationInsights, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(applicationinsightsesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApplicationInsights{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationInsights), err
}
