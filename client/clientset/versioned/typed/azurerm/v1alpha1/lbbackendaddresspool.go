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

// LbBackendAddressPoolsGetter has a method to return a LbBackendAddressPoolInterface.
// A group's client should implement this interface.
type LbBackendAddressPoolsGetter interface {
	LbBackendAddressPools(namespace string) LbBackendAddressPoolInterface
}

// LbBackendAddressPoolInterface has methods to work with LbBackendAddressPool resources.
type LbBackendAddressPoolInterface interface {
	Create(ctx context.Context, lbBackendAddressPool *v1alpha1.LbBackendAddressPool, opts v1.CreateOptions) (*v1alpha1.LbBackendAddressPool, error)
	Update(ctx context.Context, lbBackendAddressPool *v1alpha1.LbBackendAddressPool, opts v1.UpdateOptions) (*v1alpha1.LbBackendAddressPool, error)
	UpdateStatus(ctx context.Context, lbBackendAddressPool *v1alpha1.LbBackendAddressPool, opts v1.UpdateOptions) (*v1alpha1.LbBackendAddressPool, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.LbBackendAddressPool, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.LbBackendAddressPoolList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LbBackendAddressPool, err error)
	LbBackendAddressPoolExpansion
}

// lbBackendAddressPools implements LbBackendAddressPoolInterface
type lbBackendAddressPools struct {
	client rest.Interface
	ns     string
}

// newLbBackendAddressPools returns a LbBackendAddressPools
func newLbBackendAddressPools(c *AzurermV1alpha1Client, namespace string) *lbBackendAddressPools {
	return &lbBackendAddressPools{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the lbBackendAddressPool, and returns the corresponding lbBackendAddressPool object, and an error if there is any.
func (c *lbBackendAddressPools) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LbBackendAddressPool, err error) {
	result = &v1alpha1.LbBackendAddressPool{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("lbbackendaddresspools").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LbBackendAddressPools that match those selectors.
func (c *lbBackendAddressPools) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LbBackendAddressPoolList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.LbBackendAddressPoolList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("lbbackendaddresspools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested lbBackendAddressPools.
func (c *lbBackendAddressPools) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("lbbackendaddresspools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a lbBackendAddressPool and creates it.  Returns the server's representation of the lbBackendAddressPool, and an error, if there is any.
func (c *lbBackendAddressPools) Create(ctx context.Context, lbBackendAddressPool *v1alpha1.LbBackendAddressPool, opts v1.CreateOptions) (result *v1alpha1.LbBackendAddressPool, err error) {
	result = &v1alpha1.LbBackendAddressPool{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("lbbackendaddresspools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(lbBackendAddressPool).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a lbBackendAddressPool and updates it. Returns the server's representation of the lbBackendAddressPool, and an error, if there is any.
func (c *lbBackendAddressPools) Update(ctx context.Context, lbBackendAddressPool *v1alpha1.LbBackendAddressPool, opts v1.UpdateOptions) (result *v1alpha1.LbBackendAddressPool, err error) {
	result = &v1alpha1.LbBackendAddressPool{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("lbbackendaddresspools").
		Name(lbBackendAddressPool.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(lbBackendAddressPool).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *lbBackendAddressPools) UpdateStatus(ctx context.Context, lbBackendAddressPool *v1alpha1.LbBackendAddressPool, opts v1.UpdateOptions) (result *v1alpha1.LbBackendAddressPool, err error) {
	result = &v1alpha1.LbBackendAddressPool{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("lbbackendaddresspools").
		Name(lbBackendAddressPool.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(lbBackendAddressPool).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the lbBackendAddressPool and deletes it. Returns an error if one occurs.
func (c *lbBackendAddressPools) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("lbbackendaddresspools").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *lbBackendAddressPools) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("lbbackendaddresspools").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched lbBackendAddressPool.
func (c *lbBackendAddressPools) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LbBackendAddressPool, err error) {
	result = &v1alpha1.LbBackendAddressPool{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("lbbackendaddresspools").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
