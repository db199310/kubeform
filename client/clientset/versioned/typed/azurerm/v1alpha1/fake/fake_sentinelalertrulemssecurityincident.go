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

// FakeSentinelAlertRuleMsSecurityIncidents implements SentinelAlertRuleMsSecurityIncidentInterface
type FakeSentinelAlertRuleMsSecurityIncidents struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var sentinelalertrulemssecurityincidentsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "sentinelalertrulemssecurityincidents"}

var sentinelalertrulemssecurityincidentsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "SentinelAlertRuleMsSecurityIncident"}

// Get takes name of the sentinelAlertRuleMsSecurityIncident, and returns the corresponding sentinelAlertRuleMsSecurityIncident object, and an error if there is any.
func (c *FakeSentinelAlertRuleMsSecurityIncidents) Get(name string, options v1.GetOptions) (result *v1alpha1.SentinelAlertRuleMsSecurityIncident, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sentinelalertrulemssecurityincidentsResource, c.ns, name), &v1alpha1.SentinelAlertRuleMsSecurityIncident{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SentinelAlertRuleMsSecurityIncident), err
}

// List takes label and field selectors, and returns the list of SentinelAlertRuleMsSecurityIncidents that match those selectors.
func (c *FakeSentinelAlertRuleMsSecurityIncidents) List(opts v1.ListOptions) (result *v1alpha1.SentinelAlertRuleMsSecurityIncidentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sentinelalertrulemssecurityincidentsResource, sentinelalertrulemssecurityincidentsKind, c.ns, opts), &v1alpha1.SentinelAlertRuleMsSecurityIncidentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SentinelAlertRuleMsSecurityIncidentList{ListMeta: obj.(*v1alpha1.SentinelAlertRuleMsSecurityIncidentList).ListMeta}
	for _, item := range obj.(*v1alpha1.SentinelAlertRuleMsSecurityIncidentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sentinelAlertRuleMsSecurityIncidents.
func (c *FakeSentinelAlertRuleMsSecurityIncidents) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sentinelalertrulemssecurityincidentsResource, c.ns, opts))

}

// Create takes the representation of a sentinelAlertRuleMsSecurityIncident and creates it.  Returns the server's representation of the sentinelAlertRuleMsSecurityIncident, and an error, if there is any.
func (c *FakeSentinelAlertRuleMsSecurityIncidents) Create(sentinelAlertRuleMsSecurityIncident *v1alpha1.SentinelAlertRuleMsSecurityIncident) (result *v1alpha1.SentinelAlertRuleMsSecurityIncident, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sentinelalertrulemssecurityincidentsResource, c.ns, sentinelAlertRuleMsSecurityIncident), &v1alpha1.SentinelAlertRuleMsSecurityIncident{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SentinelAlertRuleMsSecurityIncident), err
}

// Update takes the representation of a sentinelAlertRuleMsSecurityIncident and updates it. Returns the server's representation of the sentinelAlertRuleMsSecurityIncident, and an error, if there is any.
func (c *FakeSentinelAlertRuleMsSecurityIncidents) Update(sentinelAlertRuleMsSecurityIncident *v1alpha1.SentinelAlertRuleMsSecurityIncident) (result *v1alpha1.SentinelAlertRuleMsSecurityIncident, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sentinelalertrulemssecurityincidentsResource, c.ns, sentinelAlertRuleMsSecurityIncident), &v1alpha1.SentinelAlertRuleMsSecurityIncident{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SentinelAlertRuleMsSecurityIncident), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSentinelAlertRuleMsSecurityIncidents) UpdateStatus(sentinelAlertRuleMsSecurityIncident *v1alpha1.SentinelAlertRuleMsSecurityIncident) (*v1alpha1.SentinelAlertRuleMsSecurityIncident, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sentinelalertrulemssecurityincidentsResource, "status", c.ns, sentinelAlertRuleMsSecurityIncident), &v1alpha1.SentinelAlertRuleMsSecurityIncident{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SentinelAlertRuleMsSecurityIncident), err
}

// Delete takes name of the sentinelAlertRuleMsSecurityIncident and deletes it. Returns an error if one occurs.
func (c *FakeSentinelAlertRuleMsSecurityIncidents) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sentinelalertrulemssecurityincidentsResource, c.ns, name), &v1alpha1.SentinelAlertRuleMsSecurityIncident{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSentinelAlertRuleMsSecurityIncidents) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sentinelalertrulemssecurityincidentsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SentinelAlertRuleMsSecurityIncidentList{})
	return err
}

// Patch applies the patch and returns the patched sentinelAlertRuleMsSecurityIncident.
func (c *FakeSentinelAlertRuleMsSecurityIncidents) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SentinelAlertRuleMsSecurityIncident, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sentinelalertrulemssecurityincidentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SentinelAlertRuleMsSecurityIncident{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SentinelAlertRuleMsSecurityIncident), err
}
