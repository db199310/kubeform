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

// SDPAzsbv1sGetter has a method to return a SDPAzsbv1Interface.
// A group's client should implement this interface.
type SDPAzsbv1sGetter interface {
	SDPAzsbv1s(namespace string) SDPAzsbv1Interface
}

// SDPAzsbv1Interface has methods to work with SDPAzsbv1 resources.
type SDPAzsbv1Interface interface {
	Create(*v1alpha2.SDPAzsbv1) (*v1alpha2.SDPAzsbv1, error)
	Update(*v1alpha2.SDPAzsbv1) (*v1alpha2.SDPAzsbv1, error)
	UpdateStatus(*v1alpha2.SDPAzsbv1) (*v1alpha2.SDPAzsbv1, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha2.SDPAzsbv1, error)
	List(opts v1.ListOptions) (*v1alpha2.SDPAzsbv1List, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.SDPAzsbv1, err error)
	SDPAzsbv1Expansion
}

// sDPAzsbv1s implements SDPAzsbv1Interface
type sDPAzsbv1s struct {
	client rest.Interface
	ns     string
}

// newSDPAzsbv1s returns a SDPAzsbv1s
func newSDPAzsbv1s(c *ModulesV1alpha2Client, namespace string) *sDPAzsbv1s {
	return &sDPAzsbv1s{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sDPAzsbv1, and returns the corresponding sDPAzsbv1 object, and an error if there is any.
func (c *sDPAzsbv1s) Get(name string, options v1.GetOptions) (result *v1alpha2.SDPAzsbv1, err error) {
	result = &v1alpha2.SDPAzsbv1{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sdpazsbv1s").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SDPAzsbv1s that match those selectors.
func (c *sDPAzsbv1s) List(opts v1.ListOptions) (result *v1alpha2.SDPAzsbv1List, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.SDPAzsbv1List{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sdpazsbv1s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sDPAzsbv1s.
func (c *sDPAzsbv1s) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sdpazsbv1s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a sDPAzsbv1 and creates it.  Returns the server's representation of the sDPAzsbv1, and an error, if there is any.
func (c *sDPAzsbv1s) Create(sDPAzsbv1 *v1alpha2.SDPAzsbv1) (result *v1alpha2.SDPAzsbv1, err error) {
	result = &v1alpha2.SDPAzsbv1{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sdpazsbv1s").
		Body(sDPAzsbv1).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sDPAzsbv1 and updates it. Returns the server's representation of the sDPAzsbv1, and an error, if there is any.
func (c *sDPAzsbv1s) Update(sDPAzsbv1 *v1alpha2.SDPAzsbv1) (result *v1alpha2.SDPAzsbv1, err error) {
	result = &v1alpha2.SDPAzsbv1{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sdpazsbv1s").
		Name(sDPAzsbv1.Name).
		Body(sDPAzsbv1).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *sDPAzsbv1s) UpdateStatus(sDPAzsbv1 *v1alpha2.SDPAzsbv1) (result *v1alpha2.SDPAzsbv1, err error) {
	result = &v1alpha2.SDPAzsbv1{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sdpazsbv1s").
		Name(sDPAzsbv1.Name).
		SubResource("status").
		Body(sDPAzsbv1).
		Do().
		Into(result)
	return
}

// Delete takes name of the sDPAzsbv1 and deletes it. Returns an error if one occurs.
func (c *sDPAzsbv1s) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sdpazsbv1s").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sDPAzsbv1s) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sdpazsbv1s").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sDPAzsbv1.
func (c *sDPAzsbv1s) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.SDPAzsbv1, err error) {
	result = &v1alpha2.SDPAzsbv1{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sdpazsbv1s").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
