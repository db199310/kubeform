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

// DevTestSchedulesGetter has a method to return a DevTestScheduleInterface.
// A group's client should implement this interface.
type DevTestSchedulesGetter interface {
	DevTestSchedules(namespace string) DevTestScheduleInterface
}

// DevTestScheduleInterface has methods to work with DevTestSchedule resources.
type DevTestScheduleInterface interface {
	Create(*v1alpha1.DevTestSchedule) (*v1alpha1.DevTestSchedule, error)
	Update(*v1alpha1.DevTestSchedule) (*v1alpha1.DevTestSchedule, error)
	UpdateStatus(*v1alpha1.DevTestSchedule) (*v1alpha1.DevTestSchedule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DevTestSchedule, error)
	List(opts v1.ListOptions) (*v1alpha1.DevTestScheduleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DevTestSchedule, err error)
	DevTestScheduleExpansion
}

// devTestSchedules implements DevTestScheduleInterface
type devTestSchedules struct {
	client rest.Interface
	ns     string
}

// newDevTestSchedules returns a DevTestSchedules
func newDevTestSchedules(c *AzurermV1alpha1Client, namespace string) *devTestSchedules {
	return &devTestSchedules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the devTestSchedule, and returns the corresponding devTestSchedule object, and an error if there is any.
func (c *devTestSchedules) Get(name string, options v1.GetOptions) (result *v1alpha1.DevTestSchedule, err error) {
	result = &v1alpha1.DevTestSchedule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("devtestschedules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DevTestSchedules that match those selectors.
func (c *devTestSchedules) List(opts v1.ListOptions) (result *v1alpha1.DevTestScheduleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DevTestScheduleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("devtestschedules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested devTestSchedules.
func (c *devTestSchedules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("devtestschedules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a devTestSchedule and creates it.  Returns the server's representation of the devTestSchedule, and an error, if there is any.
func (c *devTestSchedules) Create(devTestSchedule *v1alpha1.DevTestSchedule) (result *v1alpha1.DevTestSchedule, err error) {
	result = &v1alpha1.DevTestSchedule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("devtestschedules").
		Body(devTestSchedule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a devTestSchedule and updates it. Returns the server's representation of the devTestSchedule, and an error, if there is any.
func (c *devTestSchedules) Update(devTestSchedule *v1alpha1.DevTestSchedule) (result *v1alpha1.DevTestSchedule, err error) {
	result = &v1alpha1.DevTestSchedule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("devtestschedules").
		Name(devTestSchedule.Name).
		Body(devTestSchedule).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *devTestSchedules) UpdateStatus(devTestSchedule *v1alpha1.DevTestSchedule) (result *v1alpha1.DevTestSchedule, err error) {
	result = &v1alpha1.DevTestSchedule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("devtestschedules").
		Name(devTestSchedule.Name).
		SubResource("status").
		Body(devTestSchedule).
		Do().
		Into(result)
	return
}

// Delete takes name of the devTestSchedule and deletes it. Returns an error if one occurs.
func (c *devTestSchedules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("devtestschedules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *devTestSchedules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("devtestschedules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched devTestSchedule.
func (c *devTestSchedules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DevTestSchedule, err error) {
	result = &v1alpha1.DevTestSchedule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("devtestschedules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}