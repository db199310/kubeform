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

// FakeDnsARecords implements DnsARecordInterface
type FakeDnsARecords struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var dnsarecordsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "dnsarecords"}

var dnsarecordsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "DnsARecord"}

// Get takes name of the dnsARecord, and returns the corresponding dnsARecord object, and an error if there is any.
func (c *FakeDnsARecords) Get(name string, options v1.GetOptions) (result *v1alpha1.DnsARecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dnsarecordsResource, c.ns, name), &v1alpha1.DnsARecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DnsARecord), err
}

// List takes label and field selectors, and returns the list of DnsARecords that match those selectors.
func (c *FakeDnsARecords) List(opts v1.ListOptions) (result *v1alpha1.DnsARecordList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dnsarecordsResource, dnsarecordsKind, c.ns, opts), &v1alpha1.DnsARecordList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DnsARecordList{ListMeta: obj.(*v1alpha1.DnsARecordList).ListMeta}
	for _, item := range obj.(*v1alpha1.DnsARecordList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dnsARecords.
func (c *FakeDnsARecords) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dnsarecordsResource, c.ns, opts))

}

// Create takes the representation of a dnsARecord and creates it.  Returns the server's representation of the dnsARecord, and an error, if there is any.
func (c *FakeDnsARecords) Create(dnsARecord *v1alpha1.DnsARecord) (result *v1alpha1.DnsARecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dnsarecordsResource, c.ns, dnsARecord), &v1alpha1.DnsARecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DnsARecord), err
}

// Update takes the representation of a dnsARecord and updates it. Returns the server's representation of the dnsARecord, and an error, if there is any.
func (c *FakeDnsARecords) Update(dnsARecord *v1alpha1.DnsARecord) (result *v1alpha1.DnsARecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dnsarecordsResource, c.ns, dnsARecord), &v1alpha1.DnsARecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DnsARecord), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDnsARecords) UpdateStatus(dnsARecord *v1alpha1.DnsARecord) (*v1alpha1.DnsARecord, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dnsarecordsResource, "status", c.ns, dnsARecord), &v1alpha1.DnsARecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DnsARecord), err
}

// Delete takes name of the dnsARecord and deletes it. Returns an error if one occurs.
func (c *FakeDnsARecords) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dnsarecordsResource, c.ns, name), &v1alpha1.DnsARecord{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDnsARecords) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dnsarecordsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DnsARecordList{})
	return err
}

// Patch applies the patch and returns the patched dnsARecord.
func (c *FakeDnsARecords) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DnsARecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dnsarecordsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DnsARecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DnsARecord), err
}
