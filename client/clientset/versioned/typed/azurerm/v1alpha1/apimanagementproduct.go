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
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// ApiManagementProductsGetter has a method to return a ApiManagementProductInterface.
// A group's client should implement this interface.
type ApiManagementProductsGetter interface {
	ApiManagementProducts(namespace string) ApiManagementProductInterface
}

// ApiManagementProductInterface has methods to work with ApiManagementProduct resources.
type ApiManagementProductInterface interface {
	Create(ctx context.Context, apiManagementProduct *v1alpha1.ApiManagementProduct, opts v1.CreateOptions) (*v1alpha1.ApiManagementProduct, error)
	Update(ctx context.Context, apiManagementProduct *v1alpha1.ApiManagementProduct, opts v1.UpdateOptions) (*v1alpha1.ApiManagementProduct, error)
	UpdateStatus(ctx context.Context, apiManagementProduct *v1alpha1.ApiManagementProduct, opts v1.UpdateOptions) (*v1alpha1.ApiManagementProduct, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ApiManagementProduct, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ApiManagementProductList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApiManagementProduct, err error)
	ApiManagementProductExpansion
}

// apiManagementProducts implements ApiManagementProductInterface
type apiManagementProducts struct {
	client rest.Interface
	ns     string
}

// newApiManagementProducts returns a ApiManagementProducts
func newApiManagementProducts(c *AzurermV1alpha1Client, namespace string) *apiManagementProducts {
	return &apiManagementProducts{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the apiManagementProduct, and returns the corresponding apiManagementProduct object, and an error if there is any.
func (c *apiManagementProducts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ApiManagementProduct, err error) {
	result = &v1alpha1.ApiManagementProduct{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apimanagementproducts").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApiManagementProducts that match those selectors.
func (c *apiManagementProducts) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ApiManagementProductList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ApiManagementProductList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apimanagementproducts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested apiManagementProducts.
func (c *apiManagementProducts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("apimanagementproducts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a apiManagementProduct and creates it.  Returns the server's representation of the apiManagementProduct, and an error, if there is any.
func (c *apiManagementProducts) Create(ctx context.Context, apiManagementProduct *v1alpha1.ApiManagementProduct, opts v1.CreateOptions) (result *v1alpha1.ApiManagementProduct, err error) {
	result = &v1alpha1.ApiManagementProduct{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("apimanagementproducts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(apiManagementProduct).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a apiManagementProduct and updates it. Returns the server's representation of the apiManagementProduct, and an error, if there is any.
func (c *apiManagementProducts) Update(ctx context.Context, apiManagementProduct *v1alpha1.ApiManagementProduct, opts v1.UpdateOptions) (result *v1alpha1.ApiManagementProduct, err error) {
	result = &v1alpha1.ApiManagementProduct{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apimanagementproducts").
		Name(apiManagementProduct.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(apiManagementProduct).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *apiManagementProducts) UpdateStatus(ctx context.Context, apiManagementProduct *v1alpha1.ApiManagementProduct, opts v1.UpdateOptions) (result *v1alpha1.ApiManagementProduct, err error) {
	result = &v1alpha1.ApiManagementProduct{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apimanagementproducts").
		Name(apiManagementProduct.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(apiManagementProduct).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the apiManagementProduct and deletes it. Returns an error if one occurs.
func (c *apiManagementProducts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apimanagementproducts").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *apiManagementProducts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apimanagementproducts").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched apiManagementProduct.
func (c *apiManagementProducts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApiManagementProduct, err error) {
	result = &v1alpha1.ApiManagementProduct{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("apimanagementproducts").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
