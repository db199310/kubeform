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

// EventgridTopicsGetter has a method to return a EventgridTopicInterface.
// A group's client should implement this interface.
type EventgridTopicsGetter interface {
	EventgridTopics(namespace string) EventgridTopicInterface
}

// EventgridTopicInterface has methods to work with EventgridTopic resources.
type EventgridTopicInterface interface {
	Create(*v1alpha1.EventgridTopic) (*v1alpha1.EventgridTopic, error)
	Update(*v1alpha1.EventgridTopic) (*v1alpha1.EventgridTopic, error)
	UpdateStatus(*v1alpha1.EventgridTopic) (*v1alpha1.EventgridTopic, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.EventgridTopic, error)
	List(opts v1.ListOptions) (*v1alpha1.EventgridTopicList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EventgridTopic, err error)
	EventgridTopicExpansion
}

// eventgridTopics implements EventgridTopicInterface
type eventgridTopics struct {
	client rest.Interface
	ns     string
}

// newEventgridTopics returns a EventgridTopics
func newEventgridTopics(c *AzurermV1alpha1Client, namespace string) *eventgridTopics {
	return &eventgridTopics{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the eventgridTopic, and returns the corresponding eventgridTopic object, and an error if there is any.
func (c *eventgridTopics) Get(name string, options v1.GetOptions) (result *v1alpha1.EventgridTopic, err error) {
	result = &v1alpha1.EventgridTopic{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("eventgridtopics").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EventgridTopics that match those selectors.
func (c *eventgridTopics) List(opts v1.ListOptions) (result *v1alpha1.EventgridTopicList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.EventgridTopicList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("eventgridtopics").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested eventgridTopics.
func (c *eventgridTopics) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("eventgridtopics").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a eventgridTopic and creates it.  Returns the server's representation of the eventgridTopic, and an error, if there is any.
func (c *eventgridTopics) Create(eventgridTopic *v1alpha1.EventgridTopic) (result *v1alpha1.EventgridTopic, err error) {
	result = &v1alpha1.EventgridTopic{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("eventgridtopics").
		Body(eventgridTopic).
		Do().
		Into(result)
	return
}

// Update takes the representation of a eventgridTopic and updates it. Returns the server's representation of the eventgridTopic, and an error, if there is any.
func (c *eventgridTopics) Update(eventgridTopic *v1alpha1.EventgridTopic) (result *v1alpha1.EventgridTopic, err error) {
	result = &v1alpha1.EventgridTopic{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("eventgridtopics").
		Name(eventgridTopic.Name).
		Body(eventgridTopic).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *eventgridTopics) UpdateStatus(eventgridTopic *v1alpha1.EventgridTopic) (result *v1alpha1.EventgridTopic, err error) {
	result = &v1alpha1.EventgridTopic{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("eventgridtopics").
		Name(eventgridTopic.Name).
		SubResource("status").
		Body(eventgridTopic).
		Do().
		Into(result)
	return
}

// Delete takes name of the eventgridTopic and deletes it. Returns an error if one occurs.
func (c *eventgridTopics) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("eventgridtopics").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *eventgridTopics) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("eventgridtopics").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched eventgridTopic.
func (c *eventgridTopics) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EventgridTopic, err error) {
	result = &v1alpha1.EventgridTopic{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("eventgridtopics").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
