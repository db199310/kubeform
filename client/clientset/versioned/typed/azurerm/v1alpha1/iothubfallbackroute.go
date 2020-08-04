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

// IothubFallbackRoutesGetter has a method to return a IothubFallbackRouteInterface.
// A group's client should implement this interface.
type IothubFallbackRoutesGetter interface {
	IothubFallbackRoutes(namespace string) IothubFallbackRouteInterface
}

// IothubFallbackRouteInterface has methods to work with IothubFallbackRoute resources.
type IothubFallbackRouteInterface interface {
	Create(*v1alpha1.IothubFallbackRoute) (*v1alpha1.IothubFallbackRoute, error)
	Update(*v1alpha1.IothubFallbackRoute) (*v1alpha1.IothubFallbackRoute, error)
	UpdateStatus(*v1alpha1.IothubFallbackRoute) (*v1alpha1.IothubFallbackRoute, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.IothubFallbackRoute, error)
	List(opts v1.ListOptions) (*v1alpha1.IothubFallbackRouteList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IothubFallbackRoute, err error)
	IothubFallbackRouteExpansion
}

// iothubFallbackRoutes implements IothubFallbackRouteInterface
type iothubFallbackRoutes struct {
	client rest.Interface
	ns     string
}

// newIothubFallbackRoutes returns a IothubFallbackRoutes
func newIothubFallbackRoutes(c *AzurermV1alpha1Client, namespace string) *iothubFallbackRoutes {
	return &iothubFallbackRoutes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the iothubFallbackRoute, and returns the corresponding iothubFallbackRoute object, and an error if there is any.
func (c *iothubFallbackRoutes) Get(name string, options v1.GetOptions) (result *v1alpha1.IothubFallbackRoute, err error) {
	result = &v1alpha1.IothubFallbackRoute{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iothubfallbackroutes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IothubFallbackRoutes that match those selectors.
func (c *iothubFallbackRoutes) List(opts v1.ListOptions) (result *v1alpha1.IothubFallbackRouteList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.IothubFallbackRouteList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iothubfallbackroutes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested iothubFallbackRoutes.
func (c *iothubFallbackRoutes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("iothubfallbackroutes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a iothubFallbackRoute and creates it.  Returns the server's representation of the iothubFallbackRoute, and an error, if there is any.
func (c *iothubFallbackRoutes) Create(iothubFallbackRoute *v1alpha1.IothubFallbackRoute) (result *v1alpha1.IothubFallbackRoute, err error) {
	result = &v1alpha1.IothubFallbackRoute{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("iothubfallbackroutes").
		Body(iothubFallbackRoute).
		Do().
		Into(result)
	return
}

// Update takes the representation of a iothubFallbackRoute and updates it. Returns the server's representation of the iothubFallbackRoute, and an error, if there is any.
func (c *iothubFallbackRoutes) Update(iothubFallbackRoute *v1alpha1.IothubFallbackRoute) (result *v1alpha1.IothubFallbackRoute, err error) {
	result = &v1alpha1.IothubFallbackRoute{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iothubfallbackroutes").
		Name(iothubFallbackRoute.Name).
		Body(iothubFallbackRoute).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *iothubFallbackRoutes) UpdateStatus(iothubFallbackRoute *v1alpha1.IothubFallbackRoute) (result *v1alpha1.IothubFallbackRoute, err error) {
	result = &v1alpha1.IothubFallbackRoute{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iothubfallbackroutes").
		Name(iothubFallbackRoute.Name).
		SubResource("status").
		Body(iothubFallbackRoute).
		Do().
		Into(result)
	return
}

// Delete takes name of the iothubFallbackRoute and deletes it. Returns an error if one occurs.
func (c *iothubFallbackRoutes) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iothubfallbackroutes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *iothubFallbackRoutes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iothubfallbackroutes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched iothubFallbackRoute.
func (c *iothubFallbackRoutes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IothubFallbackRoute, err error) {
	result = &v1alpha1.IothubFallbackRoute{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("iothubfallbackroutes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
