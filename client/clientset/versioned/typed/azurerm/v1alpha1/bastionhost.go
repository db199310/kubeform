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

// BastionHostsGetter has a method to return a BastionHostInterface.
// A group's client should implement this interface.
type BastionHostsGetter interface {
	BastionHosts(namespace string) BastionHostInterface
}

// BastionHostInterface has methods to work with BastionHost resources.
type BastionHostInterface interface {
	Create(*v1alpha1.BastionHost) (*v1alpha1.BastionHost, error)
	Update(*v1alpha1.BastionHost) (*v1alpha1.BastionHost, error)
	UpdateStatus(*v1alpha1.BastionHost) (*v1alpha1.BastionHost, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.BastionHost, error)
	List(opts v1.ListOptions) (*v1alpha1.BastionHostList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BastionHost, err error)
	BastionHostExpansion
}

// bastionHosts implements BastionHostInterface
type bastionHosts struct {
	client rest.Interface
	ns     string
}

// newBastionHosts returns a BastionHosts
func newBastionHosts(c *AzurermV1alpha1Client, namespace string) *bastionHosts {
	return &bastionHosts{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the bastionHost, and returns the corresponding bastionHost object, and an error if there is any.
func (c *bastionHosts) Get(name string, options v1.GetOptions) (result *v1alpha1.BastionHost, err error) {
	result = &v1alpha1.BastionHost{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bastionhosts").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BastionHosts that match those selectors.
func (c *bastionHosts) List(opts v1.ListOptions) (result *v1alpha1.BastionHostList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.BastionHostList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bastionhosts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested bastionHosts.
func (c *bastionHosts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("bastionhosts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a bastionHost and creates it.  Returns the server's representation of the bastionHost, and an error, if there is any.
func (c *bastionHosts) Create(bastionHost *v1alpha1.BastionHost) (result *v1alpha1.BastionHost, err error) {
	result = &v1alpha1.BastionHost{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("bastionhosts").
		Body(bastionHost).
		Do().
		Into(result)
	return
}

// Update takes the representation of a bastionHost and updates it. Returns the server's representation of the bastionHost, and an error, if there is any.
func (c *bastionHosts) Update(bastionHost *v1alpha1.BastionHost) (result *v1alpha1.BastionHost, err error) {
	result = &v1alpha1.BastionHost{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bastionhosts").
		Name(bastionHost.Name).
		Body(bastionHost).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *bastionHosts) UpdateStatus(bastionHost *v1alpha1.BastionHost) (result *v1alpha1.BastionHost, err error) {
	result = &v1alpha1.BastionHost{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bastionhosts").
		Name(bastionHost.Name).
		SubResource("status").
		Body(bastionHost).
		Do().
		Into(result)
	return
}

// Delete takes name of the bastionHost and deletes it. Returns an error if one occurs.
func (c *bastionHosts) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bastionhosts").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *bastionHosts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bastionhosts").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched bastionHost.
func (c *bastionHosts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BastionHost, err error) {
	result = &v1alpha1.BastionHost{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("bastionhosts").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
