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

// EventhubNamespace_sGetter has a method to return a EventhubNamespace_Interface.
// A group's client should implement this interface.
type EventhubNamespace_sGetter interface {
	EventhubNamespace_s(namespace string) EventhubNamespace_Interface
}

// EventhubNamespace_Interface has methods to work with EventhubNamespace_ resources.
type EventhubNamespace_Interface interface {
	Create(*v1alpha1.EventhubNamespace_) (*v1alpha1.EventhubNamespace_, error)
	Update(*v1alpha1.EventhubNamespace_) (*v1alpha1.EventhubNamespace_, error)
	UpdateStatus(*v1alpha1.EventhubNamespace_) (*v1alpha1.EventhubNamespace_, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.EventhubNamespace_, error)
	List(opts v1.ListOptions) (*v1alpha1.EventhubNamespace_List, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EventhubNamespace_, err error)
	EventhubNamespace_Expansion
}

// eventhubNamespace_s implements EventhubNamespace_Interface
type eventhubNamespace_s struct {
	client rest.Interface
	ns     string
}

// newEventhubNamespace_s returns a EventhubNamespace_s
func newEventhubNamespace_s(c *AzurermV1alpha1Client, namespace string) *eventhubNamespace_s {
	return &eventhubNamespace_s{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the eventhubNamespace_, and returns the corresponding eventhubNamespace_ object, and an error if there is any.
func (c *eventhubNamespace_s) Get(name string, options v1.GetOptions) (result *v1alpha1.EventhubNamespace_, err error) {
	result = &v1alpha1.EventhubNamespace_{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("eventhubnamespace_s").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EventhubNamespace_s that match those selectors.
func (c *eventhubNamespace_s) List(opts v1.ListOptions) (result *v1alpha1.EventhubNamespace_List, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.EventhubNamespace_List{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("eventhubnamespace_s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested eventhubNamespace_s.
func (c *eventhubNamespace_s) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("eventhubnamespace_s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a eventhubNamespace_ and creates it.  Returns the server's representation of the eventhubNamespace_, and an error, if there is any.
func (c *eventhubNamespace_s) Create(eventhubNamespace_ *v1alpha1.EventhubNamespace_) (result *v1alpha1.EventhubNamespace_, err error) {
	result = &v1alpha1.EventhubNamespace_{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("eventhubnamespace_s").
		Body(eventhubNamespace_).
		Do().
		Into(result)
	return
}

// Update takes the representation of a eventhubNamespace_ and updates it. Returns the server's representation of the eventhubNamespace_, and an error, if there is any.
func (c *eventhubNamespace_s) Update(eventhubNamespace_ *v1alpha1.EventhubNamespace_) (result *v1alpha1.EventhubNamespace_, err error) {
	result = &v1alpha1.EventhubNamespace_{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("eventhubnamespace_s").
		Name(eventhubNamespace_.Name).
		Body(eventhubNamespace_).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *eventhubNamespace_s) UpdateStatus(eventhubNamespace_ *v1alpha1.EventhubNamespace_) (result *v1alpha1.EventhubNamespace_, err error) {
	result = &v1alpha1.EventhubNamespace_{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("eventhubnamespace_s").
		Name(eventhubNamespace_.Name).
		SubResource("status").
		Body(eventhubNamespace_).
		Do().
		Into(result)
	return
}

// Delete takes name of the eventhubNamespace_ and deletes it. Returns an error if one occurs.
func (c *eventhubNamespace_s) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("eventhubnamespace_s").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *eventhubNamespace_s) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("eventhubnamespace_s").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched eventhubNamespace_.
func (c *eventhubNamespace_s) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EventhubNamespace_, err error) {
	result = &v1alpha1.EventhubNamespace_{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("eventhubnamespace_s").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
