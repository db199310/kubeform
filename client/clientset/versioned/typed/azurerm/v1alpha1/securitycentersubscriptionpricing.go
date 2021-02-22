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

// SecurityCenterSubscriptionPricingsGetter has a method to return a SecurityCenterSubscriptionPricingInterface.
// A group's client should implement this interface.
type SecurityCenterSubscriptionPricingsGetter interface {
	SecurityCenterSubscriptionPricings(namespace string) SecurityCenterSubscriptionPricingInterface
}

// SecurityCenterSubscriptionPricingInterface has methods to work with SecurityCenterSubscriptionPricing resources.
type SecurityCenterSubscriptionPricingInterface interface {
	Create(ctx context.Context, securityCenterSubscriptionPricing *v1alpha1.SecurityCenterSubscriptionPricing, opts v1.CreateOptions) (*v1alpha1.SecurityCenterSubscriptionPricing, error)
	Update(ctx context.Context, securityCenterSubscriptionPricing *v1alpha1.SecurityCenterSubscriptionPricing, opts v1.UpdateOptions) (*v1alpha1.SecurityCenterSubscriptionPricing, error)
	UpdateStatus(ctx context.Context, securityCenterSubscriptionPricing *v1alpha1.SecurityCenterSubscriptionPricing, opts v1.UpdateOptions) (*v1alpha1.SecurityCenterSubscriptionPricing, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.SecurityCenterSubscriptionPricing, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.SecurityCenterSubscriptionPricingList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SecurityCenterSubscriptionPricing, err error)
	SecurityCenterSubscriptionPricingExpansion
}

// securityCenterSubscriptionPricings implements SecurityCenterSubscriptionPricingInterface
type securityCenterSubscriptionPricings struct {
	client rest.Interface
	ns     string
}

// newSecurityCenterSubscriptionPricings returns a SecurityCenterSubscriptionPricings
func newSecurityCenterSubscriptionPricings(c *AzurermV1alpha1Client, namespace string) *securityCenterSubscriptionPricings {
	return &securityCenterSubscriptionPricings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the securityCenterSubscriptionPricing, and returns the corresponding securityCenterSubscriptionPricing object, and an error if there is any.
func (c *securityCenterSubscriptionPricings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SecurityCenterSubscriptionPricing, err error) {
	result = &v1alpha1.SecurityCenterSubscriptionPricing{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("securitycentersubscriptionpricings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SecurityCenterSubscriptionPricings that match those selectors.
func (c *securityCenterSubscriptionPricings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SecurityCenterSubscriptionPricingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SecurityCenterSubscriptionPricingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("securitycentersubscriptionpricings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested securityCenterSubscriptionPricings.
func (c *securityCenterSubscriptionPricings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("securitycentersubscriptionpricings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a securityCenterSubscriptionPricing and creates it.  Returns the server's representation of the securityCenterSubscriptionPricing, and an error, if there is any.
func (c *securityCenterSubscriptionPricings) Create(ctx context.Context, securityCenterSubscriptionPricing *v1alpha1.SecurityCenterSubscriptionPricing, opts v1.CreateOptions) (result *v1alpha1.SecurityCenterSubscriptionPricing, err error) {
	result = &v1alpha1.SecurityCenterSubscriptionPricing{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("securitycentersubscriptionpricings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(securityCenterSubscriptionPricing).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a securityCenterSubscriptionPricing and updates it. Returns the server's representation of the securityCenterSubscriptionPricing, and an error, if there is any.
func (c *securityCenterSubscriptionPricings) Update(ctx context.Context, securityCenterSubscriptionPricing *v1alpha1.SecurityCenterSubscriptionPricing, opts v1.UpdateOptions) (result *v1alpha1.SecurityCenterSubscriptionPricing, err error) {
	result = &v1alpha1.SecurityCenterSubscriptionPricing{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("securitycentersubscriptionpricings").
		Name(securityCenterSubscriptionPricing.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(securityCenterSubscriptionPricing).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *securityCenterSubscriptionPricings) UpdateStatus(ctx context.Context, securityCenterSubscriptionPricing *v1alpha1.SecurityCenterSubscriptionPricing, opts v1.UpdateOptions) (result *v1alpha1.SecurityCenterSubscriptionPricing, err error) {
	result = &v1alpha1.SecurityCenterSubscriptionPricing{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("securitycentersubscriptionpricings").
		Name(securityCenterSubscriptionPricing.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(securityCenterSubscriptionPricing).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the securityCenterSubscriptionPricing and deletes it. Returns an error if one occurs.
func (c *securityCenterSubscriptionPricings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("securitycentersubscriptionpricings").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *securityCenterSubscriptionPricings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("securitycentersubscriptionpricings").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched securityCenterSubscriptionPricing.
func (c *securityCenterSubscriptionPricings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SecurityCenterSubscriptionPricing, err error) {
	result = &v1alpha1.SecurityCenterSubscriptionPricing{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("securitycentersubscriptionpricings").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
