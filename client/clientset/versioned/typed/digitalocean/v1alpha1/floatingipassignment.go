/*
Copyright 2019 The Kubeform Authors.

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
	v1alpha1 "kubeform.dev/kubeform/apis/digitalocean/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// FloatingIpAssignmentsGetter has a method to return a FloatingIpAssignmentInterface.
// A group's client should implement this interface.
type FloatingIpAssignmentsGetter interface {
	FloatingIpAssignments() FloatingIpAssignmentInterface
}

// FloatingIpAssignmentInterface has methods to work with FloatingIpAssignment resources.
type FloatingIpAssignmentInterface interface {
	Create(*v1alpha1.FloatingIpAssignment) (*v1alpha1.FloatingIpAssignment, error)
	Update(*v1alpha1.FloatingIpAssignment) (*v1alpha1.FloatingIpAssignment, error)
	UpdateStatus(*v1alpha1.FloatingIpAssignment) (*v1alpha1.FloatingIpAssignment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.FloatingIpAssignment, error)
	List(opts v1.ListOptions) (*v1alpha1.FloatingIpAssignmentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FloatingIpAssignment, err error)
	FloatingIpAssignmentExpansion
}

// floatingIpAssignments implements FloatingIpAssignmentInterface
type floatingIpAssignments struct {
	client rest.Interface
}

// newFloatingIpAssignments returns a FloatingIpAssignments
func newFloatingIpAssignments(c *DigitaloceanV1alpha1Client) *floatingIpAssignments {
	return &floatingIpAssignments{
		client: c.RESTClient(),
	}
}

// Get takes name of the floatingIpAssignment, and returns the corresponding floatingIpAssignment object, and an error if there is any.
func (c *floatingIpAssignments) Get(name string, options v1.GetOptions) (result *v1alpha1.FloatingIpAssignment, err error) {
	result = &v1alpha1.FloatingIpAssignment{}
	err = c.client.Get().
		Resource("floatingipassignments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FloatingIpAssignments that match those selectors.
func (c *floatingIpAssignments) List(opts v1.ListOptions) (result *v1alpha1.FloatingIpAssignmentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.FloatingIpAssignmentList{}
	err = c.client.Get().
		Resource("floatingipassignments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested floatingIpAssignments.
func (c *floatingIpAssignments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("floatingipassignments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a floatingIpAssignment and creates it.  Returns the server's representation of the floatingIpAssignment, and an error, if there is any.
func (c *floatingIpAssignments) Create(floatingIpAssignment *v1alpha1.FloatingIpAssignment) (result *v1alpha1.FloatingIpAssignment, err error) {
	result = &v1alpha1.FloatingIpAssignment{}
	err = c.client.Post().
		Resource("floatingipassignments").
		Body(floatingIpAssignment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a floatingIpAssignment and updates it. Returns the server's representation of the floatingIpAssignment, and an error, if there is any.
func (c *floatingIpAssignments) Update(floatingIpAssignment *v1alpha1.FloatingIpAssignment) (result *v1alpha1.FloatingIpAssignment, err error) {
	result = &v1alpha1.FloatingIpAssignment{}
	err = c.client.Put().
		Resource("floatingipassignments").
		Name(floatingIpAssignment.Name).
		Body(floatingIpAssignment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *floatingIpAssignments) UpdateStatus(floatingIpAssignment *v1alpha1.FloatingIpAssignment) (result *v1alpha1.FloatingIpAssignment, err error) {
	result = &v1alpha1.FloatingIpAssignment{}
	err = c.client.Put().
		Resource("floatingipassignments").
		Name(floatingIpAssignment.Name).
		SubResource("status").
		Body(floatingIpAssignment).
		Do().
		Into(result)
	return
}

// Delete takes name of the floatingIpAssignment and deletes it. Returns an error if one occurs.
func (c *floatingIpAssignments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("floatingipassignments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *floatingIpAssignments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("floatingipassignments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched floatingIpAssignment.
func (c *floatingIpAssignments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FloatingIpAssignment, err error) {
	result = &v1alpha1.FloatingIpAssignment{}
	err = c.client.Patch(pt).
		Resource("floatingipassignments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}