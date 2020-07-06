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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// ComputeSSLCertificatesGetter has a method to return a ComputeSSLCertificateInterface.
// A group's client should implement this interface.
type ComputeSSLCertificatesGetter interface {
	ComputeSSLCertificates(namespace string) ComputeSSLCertificateInterface
}

// ComputeSSLCertificateInterface has methods to work with ComputeSSLCertificate resources.
type ComputeSSLCertificateInterface interface {
	Create(*v1alpha1.ComputeSSLCertificate) (*v1alpha1.ComputeSSLCertificate, error)
	Update(*v1alpha1.ComputeSSLCertificate) (*v1alpha1.ComputeSSLCertificate, error)
	UpdateStatus(*v1alpha1.ComputeSSLCertificate) (*v1alpha1.ComputeSSLCertificate, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ComputeSSLCertificate, error)
	List(opts v1.ListOptions) (*v1alpha1.ComputeSSLCertificateList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeSSLCertificate, err error)
	ComputeSSLCertificateExpansion
}

// computeSSLCertificates implements ComputeSSLCertificateInterface
type computeSSLCertificates struct {
	client rest.Interface
	ns     string
}

// newComputeSSLCertificates returns a ComputeSSLCertificates
func newComputeSSLCertificates(c *GoogleV1alpha1Client, namespace string) *computeSSLCertificates {
	return &computeSSLCertificates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the computeSSLCertificate, and returns the corresponding computeSSLCertificate object, and an error if there is any.
func (c *computeSSLCertificates) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeSSLCertificate, err error) {
	result = &v1alpha1.ComputeSSLCertificate{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computesslcertificates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComputeSSLCertificates that match those selectors.
func (c *computeSSLCertificates) List(opts v1.ListOptions) (result *v1alpha1.ComputeSSLCertificateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ComputeSSLCertificateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computesslcertificates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested computeSSLCertificates.
func (c *computeSSLCertificates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("computesslcertificates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a computeSSLCertificate and creates it.  Returns the server's representation of the computeSSLCertificate, and an error, if there is any.
func (c *computeSSLCertificates) Create(computeSSLCertificate *v1alpha1.ComputeSSLCertificate) (result *v1alpha1.ComputeSSLCertificate, err error) {
	result = &v1alpha1.ComputeSSLCertificate{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("computesslcertificates").
		Body(computeSSLCertificate).
		Do().
		Into(result)
	return
}

// Update takes the representation of a computeSSLCertificate and updates it. Returns the server's representation of the computeSSLCertificate, and an error, if there is any.
func (c *computeSSLCertificates) Update(computeSSLCertificate *v1alpha1.ComputeSSLCertificate) (result *v1alpha1.ComputeSSLCertificate, err error) {
	result = &v1alpha1.ComputeSSLCertificate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computesslcertificates").
		Name(computeSSLCertificate.Name).
		Body(computeSSLCertificate).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *computeSSLCertificates) UpdateStatus(computeSSLCertificate *v1alpha1.ComputeSSLCertificate) (result *v1alpha1.ComputeSSLCertificate, err error) {
	result = &v1alpha1.ComputeSSLCertificate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computesslcertificates").
		Name(computeSSLCertificate.Name).
		SubResource("status").
		Body(computeSSLCertificate).
		Do().
		Into(result)
	return
}

// Delete takes name of the computeSSLCertificate and deletes it. Returns an error if one occurs.
func (c *computeSSLCertificates) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computesslcertificates").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *computeSSLCertificates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computesslcertificates").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched computeSSLCertificate.
func (c *computeSSLCertificates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeSSLCertificate, err error) {
	result = &v1alpha1.ComputeSSLCertificate{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("computesslcertificates").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
