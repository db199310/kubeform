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

// LbsGetter has a method to return a LbInterface.
// A group's client should implement this interface.
type LbsGetter interface {
	Lbs(namespace string) LbInterface
}

// LbInterface has methods to work with Lb resources.
type LbInterface interface {
	Create(*v1alpha1.Lb) (*v1alpha1.Lb, error)
	Update(*v1alpha1.Lb) (*v1alpha1.Lb, error)
	UpdateStatus(*v1alpha1.Lb) (*v1alpha1.Lb, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Lb, error)
	List(opts v1.ListOptions) (*v1alpha1.LbList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Lb, err error)
	LbExpansion
}

// lbs implements LbInterface
type lbs struct {
	client rest.Interface
	ns     string
}

// newLbs returns a Lbs
func newLbs(c *AzurermV1alpha1Client, namespace string) *lbs {
	return &lbs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the lb, and returns the corresponding lb object, and an error if there is any.
func (c *lbs) Get(name string, options v1.GetOptions) (result *v1alpha1.Lb, err error) {
	result = &v1alpha1.Lb{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("lbs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Lbs that match those selectors.
func (c *lbs) List(opts v1.ListOptions) (result *v1alpha1.LbList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.LbList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("lbs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested lbs.
func (c *lbs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("lbs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a lb and creates it.  Returns the server's representation of the lb, and an error, if there is any.
func (c *lbs) Create(lb *v1alpha1.Lb) (result *v1alpha1.Lb, err error) {
	result = &v1alpha1.Lb{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("lbs").
		Body(lb).
		Do().
		Into(result)
	return
}

// Update takes the representation of a lb and updates it. Returns the server's representation of the lb, and an error, if there is any.
func (c *lbs) Update(lb *v1alpha1.Lb) (result *v1alpha1.Lb, err error) {
	result = &v1alpha1.Lb{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("lbs").
		Name(lb.Name).
		Body(lb).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *lbs) UpdateStatus(lb *v1alpha1.Lb) (result *v1alpha1.Lb, err error) {
	result = &v1alpha1.Lb{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("lbs").
		Name(lb.Name).
		SubResource("status").
		Body(lb).
		Do().
		Into(result)
	return
}

// Delete takes name of the lb and deletes it. Returns an error if one occurs.
func (c *lbs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("lbs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *lbs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("lbs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched lb.
func (c *lbs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Lb, err error) {
	result = &v1alpha1.Lb{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("lbs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
