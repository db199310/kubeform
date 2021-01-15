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

package v1alpha2

import (
	"time"

	v1alpha2 "kubeform.dev/kubeform/apis/modules/v1alpha2"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SDPAzAppv1sGetter has a method to return a SDPAzAppv1Interface.
// A group's client should implement this interface.
type SDPAzAppv1sGetter interface {
	SDPAzAppv1s(namespace string) SDPAzAppv1Interface
}

// SDPAzAppv1Interface has methods to work with SDPAzAppv1 resources.
type SDPAzAppv1Interface interface {
	Create(*v1alpha2.SDPAzAppv1) (*v1alpha2.SDPAzAppv1, error)
	Update(*v1alpha2.SDPAzAppv1) (*v1alpha2.SDPAzAppv1, error)
	UpdateStatus(*v1alpha2.SDPAzAppv1) (*v1alpha2.SDPAzAppv1, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha2.SDPAzAppv1, error)
	List(opts v1.ListOptions) (*v1alpha2.SDPAzAppv1List, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.SDPAzAppv1, err error)
	SDPAzAppv1Expansion
}

// sDPAzAppv1s implements SDPAzAppv1Interface
type sDPAzAppv1s struct {
	client rest.Interface
	ns     string
}

// newSDPAzAppv1s returns a SDPAzAppv1s
func newSDPAzAppv1s(c *ModulesV1alpha2Client, namespace string) *sDPAzAppv1s {
	return &sDPAzAppv1s{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sDPAzAppv1, and returns the corresponding sDPAzAppv1 object, and an error if there is any.
func (c *sDPAzAppv1s) Get(name string, options v1.GetOptions) (result *v1alpha2.SDPAzAppv1, err error) {
	result = &v1alpha2.SDPAzAppv1{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sdpazappv1s").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SDPAzAppv1s that match those selectors.
func (c *sDPAzAppv1s) List(opts v1.ListOptions) (result *v1alpha2.SDPAzAppv1List, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.SDPAzAppv1List{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sdpazappv1s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sDPAzAppv1s.
func (c *sDPAzAppv1s) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sdpazappv1s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a sDPAzAppv1 and creates it.  Returns the server's representation of the sDPAzAppv1, and an error, if there is any.
func (c *sDPAzAppv1s) Create(sDPAzAppv1 *v1alpha2.SDPAzAppv1) (result *v1alpha2.SDPAzAppv1, err error) {
	result = &v1alpha2.SDPAzAppv1{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sdpazappv1s").
		Body(sDPAzAppv1).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sDPAzAppv1 and updates it. Returns the server's representation of the sDPAzAppv1, and an error, if there is any.
func (c *sDPAzAppv1s) Update(sDPAzAppv1 *v1alpha2.SDPAzAppv1) (result *v1alpha2.SDPAzAppv1, err error) {
	result = &v1alpha2.SDPAzAppv1{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sdpazappv1s").
		Name(sDPAzAppv1.Name).
		Body(sDPAzAppv1).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *sDPAzAppv1s) UpdateStatus(sDPAzAppv1 *v1alpha2.SDPAzAppv1) (result *v1alpha2.SDPAzAppv1, err error) {
	result = &v1alpha2.SDPAzAppv1{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sdpazappv1s").
		Name(sDPAzAppv1.Name).
		SubResource("status").
		Body(sDPAzAppv1).
		Do().
		Into(result)
	return
}

// Delete takes name of the sDPAzAppv1 and deletes it. Returns an error if one occurs.
func (c *sDPAzAppv1s) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sdpazappv1s").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sDPAzAppv1s) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sdpazappv1s").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sDPAzAppv1.
func (c *sDPAzAppv1s) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.SDPAzAppv1, err error) {
	result = &v1alpha2.SDPAzAppv1{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sdpazappv1s").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
