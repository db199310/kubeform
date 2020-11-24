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
	v1alpha1 "kubeform.dev/kubeform/apis/modules/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// SDPAzdbv1sGetter has a method to return a SDPAzdbv1Interface.
// A group's client should implement this interface.
type SDPAzdbv1sGetter interface {
	SDPAzdbv1s(namespace string) SDPAzdbv1Interface
}

// SDPAzdbv1Interface has methods to work with SDPAzdbv1 resources.
type SDPAzdbv1Interface interface {
	Create(*v1alpha1.SDPAzdbv1) (*v1alpha1.SDPAzdbv1, error)
	Update(*v1alpha1.SDPAzdbv1) (*v1alpha1.SDPAzdbv1, error)
	UpdateStatus(*v1alpha1.SDPAzdbv1) (*v1alpha1.SDPAzdbv1, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SDPAzdbv1, error)
	List(opts v1.ListOptions) (*v1alpha1.SDPAzdbv1List, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SDPAzdbv1, err error)
	SDPAzdbv1Expansion
}

// sDPAzdbv1s implements SDPAzdbv1Interface
type sDPAzdbv1s struct {
	client rest.Interface
	ns     string
}

// newSDPAzdbv1s returns a SDPAzdbv1s
func newSDPAzdbv1s(c *ModulesV1alpha1Client, namespace string) *sDPAzdbv1s {
	return &sDPAzdbv1s{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sDPAzdbv1, and returns the corresponding sDPAzdbv1 object, and an error if there is any.
func (c *sDPAzdbv1s) Get(name string, options v1.GetOptions) (result *v1alpha1.SDPAzdbv1, err error) {
	result = &v1alpha1.SDPAzdbv1{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sdpazdbv1s").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SDPAzdbv1s that match those selectors.
func (c *sDPAzdbv1s) List(opts v1.ListOptions) (result *v1alpha1.SDPAzdbv1List, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SDPAzdbv1List{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sdpazdbv1s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sDPAzdbv1s.
func (c *sDPAzdbv1s) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sdpazdbv1s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a sDPAzdbv1 and creates it.  Returns the server's representation of the sDPAzdbv1, and an error, if there is any.
func (c *sDPAzdbv1s) Create(sDPAzdbv1 *v1alpha1.SDPAzdbv1) (result *v1alpha1.SDPAzdbv1, err error) {
	result = &v1alpha1.SDPAzdbv1{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sdpazdbv1s").
		Body(sDPAzdbv1).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sDPAzdbv1 and updates it. Returns the server's representation of the sDPAzdbv1, and an error, if there is any.
func (c *sDPAzdbv1s) Update(sDPAzdbv1 *v1alpha1.SDPAzdbv1) (result *v1alpha1.SDPAzdbv1, err error) {
	result = &v1alpha1.SDPAzdbv1{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sdpazdbv1s").
		Name(sDPAzdbv1.Name).
		Body(sDPAzdbv1).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *sDPAzdbv1s) UpdateStatus(sDPAzdbv1 *v1alpha1.SDPAzdbv1) (result *v1alpha1.SDPAzdbv1, err error) {
	result = &v1alpha1.SDPAzdbv1{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sdpazdbv1s").
		Name(sDPAzdbv1.Name).
		SubResource("status").
		Body(sDPAzdbv1).
		Do().
		Into(result)
	return
}

// Delete takes name of the sDPAzdbv1 and deletes it. Returns an error if one occurs.
func (c *sDPAzdbv1s) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sdpazdbv1s").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sDPAzdbv1s) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sdpazdbv1s").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sDPAzdbv1.
func (c *sDPAzdbv1s) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SDPAzdbv1, err error) {
	result = &v1alpha1.SDPAzdbv1{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sdpazdbv1s").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
