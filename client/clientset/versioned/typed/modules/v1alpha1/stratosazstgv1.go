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

// StratosAzStgv1sGetter has a method to return a StratosAzStgv1Interface.
// A group's client should implement this interface.
type StratosAzStgv1sGetter interface {
	StratosAzStgv1s(namespace string) StratosAzStgv1Interface
}

// StratosAzStgv1Interface has methods to work with StratosAzStgv1 resources.
type StratosAzStgv1Interface interface {
	Create(*v1alpha1.StratosAzStgv1) (*v1alpha1.StratosAzStgv1, error)
	Update(*v1alpha1.StratosAzStgv1) (*v1alpha1.StratosAzStgv1, error)
	UpdateStatus(*v1alpha1.StratosAzStgv1) (*v1alpha1.StratosAzStgv1, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.StratosAzStgv1, error)
	List(opts v1.ListOptions) (*v1alpha1.StratosAzStgv1List, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StratosAzStgv1, err error)
	StratosAzStgv1Expansion
}

// stratosAzStgv1s implements StratosAzStgv1Interface
type stratosAzStgv1s struct {
	client rest.Interface
	ns     string
}

// newStratosAzStgv1s returns a StratosAzStgv1s
func newStratosAzStgv1s(c *ModulesV1alpha1Client, namespace string) *stratosAzStgv1s {
	return &stratosAzStgv1s{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the stratosAzStgv1, and returns the corresponding stratosAzStgv1 object, and an error if there is any.
func (c *stratosAzStgv1s) Get(name string, options v1.GetOptions) (result *v1alpha1.StratosAzStgv1, err error) {
	result = &v1alpha1.StratosAzStgv1{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("stratosazstgv1s").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of StratosAzStgv1s that match those selectors.
func (c *stratosAzStgv1s) List(opts v1.ListOptions) (result *v1alpha1.StratosAzStgv1List, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.StratosAzStgv1List{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("stratosazstgv1s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested stratosAzStgv1s.
func (c *stratosAzStgv1s) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("stratosazstgv1s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a stratosAzStgv1 and creates it.  Returns the server's representation of the stratosAzStgv1, and an error, if there is any.
func (c *stratosAzStgv1s) Create(stratosAzStgv1 *v1alpha1.StratosAzStgv1) (result *v1alpha1.StratosAzStgv1, err error) {
	result = &v1alpha1.StratosAzStgv1{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("stratosazstgv1s").
		Body(stratosAzStgv1).
		Do().
		Into(result)
	return
}

// Update takes the representation of a stratosAzStgv1 and updates it. Returns the server's representation of the stratosAzStgv1, and an error, if there is any.
func (c *stratosAzStgv1s) Update(stratosAzStgv1 *v1alpha1.StratosAzStgv1) (result *v1alpha1.StratosAzStgv1, err error) {
	result = &v1alpha1.StratosAzStgv1{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("stratosazstgv1s").
		Name(stratosAzStgv1.Name).
		Body(stratosAzStgv1).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *stratosAzStgv1s) UpdateStatus(stratosAzStgv1 *v1alpha1.StratosAzStgv1) (result *v1alpha1.StratosAzStgv1, err error) {
	result = &v1alpha1.StratosAzStgv1{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("stratosazstgv1s").
		Name(stratosAzStgv1.Name).
		SubResource("status").
		Body(stratosAzStgv1).
		Do().
		Into(result)
	return
}

// Delete takes name of the stratosAzStgv1 and deletes it. Returns an error if one occurs.
func (c *stratosAzStgv1s) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("stratosazstgv1s").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *stratosAzStgv1s) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("stratosazstgv1s").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched stratosAzStgv1.
func (c *stratosAzStgv1s) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StratosAzStgv1, err error) {
	result = &v1alpha1.StratosAzStgv1{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("stratosazstgv1s").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
