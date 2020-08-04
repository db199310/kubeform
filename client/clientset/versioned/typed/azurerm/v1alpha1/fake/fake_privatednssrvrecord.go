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

// FakePrivateDNSSrvRecords implements PrivateDNSSrvRecordInterface
type FakePrivateDNSSrvRecords struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var privatednssrvrecordsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "privatednssrvrecords"}

var privatednssrvrecordsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "PrivateDNSSrvRecord"}

// Get takes name of the privateDNSSrvRecord, and returns the corresponding privateDNSSrvRecord object, and an error if there is any.
func (c *FakePrivateDNSSrvRecords) Get(name string, options v1.GetOptions) (result *v1alpha1.PrivateDNSSrvRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(privatednssrvrecordsResource, c.ns, name), &v1alpha1.PrivateDNSSrvRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrivateDNSSrvRecord), err
}

// List takes label and field selectors, and returns the list of PrivateDNSSrvRecords that match those selectors.
func (c *FakePrivateDNSSrvRecords) List(opts v1.ListOptions) (result *v1alpha1.PrivateDNSSrvRecordList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(privatednssrvrecordsResource, privatednssrvrecordsKind, c.ns, opts), &v1alpha1.PrivateDNSSrvRecordList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PrivateDNSSrvRecordList{ListMeta: obj.(*v1alpha1.PrivateDNSSrvRecordList).ListMeta}
	for _, item := range obj.(*v1alpha1.PrivateDNSSrvRecordList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested privateDNSSrvRecords.
func (c *FakePrivateDNSSrvRecords) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(privatednssrvrecordsResource, c.ns, opts))

}

// Create takes the representation of a privateDNSSrvRecord and creates it.  Returns the server's representation of the privateDNSSrvRecord, and an error, if there is any.
func (c *FakePrivateDNSSrvRecords) Create(privateDNSSrvRecord *v1alpha1.PrivateDNSSrvRecord) (result *v1alpha1.PrivateDNSSrvRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(privatednssrvrecordsResource, c.ns, privateDNSSrvRecord), &v1alpha1.PrivateDNSSrvRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrivateDNSSrvRecord), err
}

// Update takes the representation of a privateDNSSrvRecord and updates it. Returns the server's representation of the privateDNSSrvRecord, and an error, if there is any.
func (c *FakePrivateDNSSrvRecords) Update(privateDNSSrvRecord *v1alpha1.PrivateDNSSrvRecord) (result *v1alpha1.PrivateDNSSrvRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(privatednssrvrecordsResource, c.ns, privateDNSSrvRecord), &v1alpha1.PrivateDNSSrvRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrivateDNSSrvRecord), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePrivateDNSSrvRecords) UpdateStatus(privateDNSSrvRecord *v1alpha1.PrivateDNSSrvRecord) (*v1alpha1.PrivateDNSSrvRecord, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(privatednssrvrecordsResource, "status", c.ns, privateDNSSrvRecord), &v1alpha1.PrivateDNSSrvRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrivateDNSSrvRecord), err
}

// Delete takes name of the privateDNSSrvRecord and deletes it. Returns an error if one occurs.
func (c *FakePrivateDNSSrvRecords) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(privatednssrvrecordsResource, c.ns, name), &v1alpha1.PrivateDNSSrvRecord{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePrivateDNSSrvRecords) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(privatednssrvrecordsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.PrivateDNSSrvRecordList{})
	return err
}

// Patch applies the patch and returns the patched privateDNSSrvRecord.
func (c *FakePrivateDNSSrvRecords) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PrivateDNSSrvRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(privatednssrvrecordsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PrivateDNSSrvRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrivateDNSSrvRecord), err
}
