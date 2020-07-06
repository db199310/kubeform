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

package v1alpha1

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// DnsCnameRecordsGetter has a method to return a DnsCnameRecordInterface.
// A group's client should implement this interface.
type DnsCnameRecordsGetter interface {
	DnsCnameRecords(namespace string) DnsCnameRecordInterface
}

// DnsCnameRecordInterface has methods to work with DnsCnameRecord resources.
type DnsCnameRecordInterface interface {
	Create(*v1alpha1.DnsCnameRecord) (*v1alpha1.DnsCnameRecord, error)
	Update(*v1alpha1.DnsCnameRecord) (*v1alpha1.DnsCnameRecord, error)
	UpdateStatus(*v1alpha1.DnsCnameRecord) (*v1alpha1.DnsCnameRecord, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DnsCnameRecord, error)
	List(opts v1.ListOptions) (*v1alpha1.DnsCnameRecordList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DnsCnameRecord, err error)
	DnsCnameRecordExpansion
}

// dnsCnameRecords implements DnsCnameRecordInterface
type dnsCnameRecords struct {
	client rest.Interface
	ns     string
}

// newDnsCnameRecords returns a DnsCnameRecords
func newDnsCnameRecords(c *AzurermV1alpha1Client, namespace string) *dnsCnameRecords {
	return &dnsCnameRecords{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dnsCnameRecord, and returns the corresponding dnsCnameRecord object, and an error if there is any.
func (c *dnsCnameRecords) Get(name string, options v1.GetOptions) (result *v1alpha1.DnsCnameRecord, err error) {
	result = &v1alpha1.DnsCnameRecord{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dnscnamerecords").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DnsCnameRecords that match those selectors.
func (c *dnsCnameRecords) List(opts v1.ListOptions) (result *v1alpha1.DnsCnameRecordList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DnsCnameRecordList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dnscnamerecords").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dnsCnameRecords.
func (c *dnsCnameRecords) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dnscnamerecords").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a dnsCnameRecord and creates it.  Returns the server's representation of the dnsCnameRecord, and an error, if there is any.
func (c *dnsCnameRecords) Create(dnsCnameRecord *v1alpha1.DnsCnameRecord) (result *v1alpha1.DnsCnameRecord, err error) {
	result = &v1alpha1.DnsCnameRecord{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dnscnamerecords").
		Body(dnsCnameRecord).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dnsCnameRecord and updates it. Returns the server's representation of the dnsCnameRecord, and an error, if there is any.
func (c *dnsCnameRecords) Update(dnsCnameRecord *v1alpha1.DnsCnameRecord) (result *v1alpha1.DnsCnameRecord, err error) {
	result = &v1alpha1.DnsCnameRecord{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dnscnamerecords").
		Name(dnsCnameRecord.Name).
		Body(dnsCnameRecord).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *dnsCnameRecords) UpdateStatus(dnsCnameRecord *v1alpha1.DnsCnameRecord) (result *v1alpha1.DnsCnameRecord, err error) {
	result = &v1alpha1.DnsCnameRecord{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dnscnamerecords").
		Name(dnsCnameRecord.Name).
		SubResource("status").
		Body(dnsCnameRecord).
		Do().
		Into(result)
	return
}

// Delete takes name of the dnsCnameRecord and deletes it. Returns an error if one occurs.
func (c *dnsCnameRecords) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dnscnamerecords").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dnsCnameRecords) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dnscnamerecords").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dnsCnameRecord.
func (c *dnsCnameRecords) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DnsCnameRecord, err error) {
	result = &v1alpha1.DnsCnameRecord{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dnscnamerecords").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
