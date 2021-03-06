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

// MariadbVirtualNetworkRulesGetter has a method to return a MariadbVirtualNetworkRuleInterface.
// A group's client should implement this interface.
type MariadbVirtualNetworkRulesGetter interface {
	MariadbVirtualNetworkRules(namespace string) MariadbVirtualNetworkRuleInterface
}

// MariadbVirtualNetworkRuleInterface has methods to work with MariadbVirtualNetworkRule resources.
type MariadbVirtualNetworkRuleInterface interface {
	Create(ctx context.Context, mariadbVirtualNetworkRule *v1alpha1.MariadbVirtualNetworkRule, opts v1.CreateOptions) (*v1alpha1.MariadbVirtualNetworkRule, error)
	Update(ctx context.Context, mariadbVirtualNetworkRule *v1alpha1.MariadbVirtualNetworkRule, opts v1.UpdateOptions) (*v1alpha1.MariadbVirtualNetworkRule, error)
	UpdateStatus(ctx context.Context, mariadbVirtualNetworkRule *v1alpha1.MariadbVirtualNetworkRule, opts v1.UpdateOptions) (*v1alpha1.MariadbVirtualNetworkRule, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.MariadbVirtualNetworkRule, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.MariadbVirtualNetworkRuleList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MariadbVirtualNetworkRule, err error)
	MariadbVirtualNetworkRuleExpansion
}

// mariadbVirtualNetworkRules implements MariadbVirtualNetworkRuleInterface
type mariadbVirtualNetworkRules struct {
	client rest.Interface
	ns     string
}

// newMariadbVirtualNetworkRules returns a MariadbVirtualNetworkRules
func newMariadbVirtualNetworkRules(c *AzurermV1alpha1Client, namespace string) *mariadbVirtualNetworkRules {
	return &mariadbVirtualNetworkRules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the mariadbVirtualNetworkRule, and returns the corresponding mariadbVirtualNetworkRule object, and an error if there is any.
func (c *mariadbVirtualNetworkRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MariadbVirtualNetworkRule, err error) {
	result = &v1alpha1.MariadbVirtualNetworkRule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mariadbvirtualnetworkrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MariadbVirtualNetworkRules that match those selectors.
func (c *mariadbVirtualNetworkRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MariadbVirtualNetworkRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.MariadbVirtualNetworkRuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mariadbvirtualnetworkrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested mariadbVirtualNetworkRules.
func (c *mariadbVirtualNetworkRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mariadbvirtualnetworkrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a mariadbVirtualNetworkRule and creates it.  Returns the server's representation of the mariadbVirtualNetworkRule, and an error, if there is any.
func (c *mariadbVirtualNetworkRules) Create(ctx context.Context, mariadbVirtualNetworkRule *v1alpha1.MariadbVirtualNetworkRule, opts v1.CreateOptions) (result *v1alpha1.MariadbVirtualNetworkRule, err error) {
	result = &v1alpha1.MariadbVirtualNetworkRule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mariadbvirtualnetworkrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(mariadbVirtualNetworkRule).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a mariadbVirtualNetworkRule and updates it. Returns the server's representation of the mariadbVirtualNetworkRule, and an error, if there is any.
func (c *mariadbVirtualNetworkRules) Update(ctx context.Context, mariadbVirtualNetworkRule *v1alpha1.MariadbVirtualNetworkRule, opts v1.UpdateOptions) (result *v1alpha1.MariadbVirtualNetworkRule, err error) {
	result = &v1alpha1.MariadbVirtualNetworkRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mariadbvirtualnetworkrules").
		Name(mariadbVirtualNetworkRule.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(mariadbVirtualNetworkRule).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *mariadbVirtualNetworkRules) UpdateStatus(ctx context.Context, mariadbVirtualNetworkRule *v1alpha1.MariadbVirtualNetworkRule, opts v1.UpdateOptions) (result *v1alpha1.MariadbVirtualNetworkRule, err error) {
	result = &v1alpha1.MariadbVirtualNetworkRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mariadbvirtualnetworkrules").
		Name(mariadbVirtualNetworkRule.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(mariadbVirtualNetworkRule).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the mariadbVirtualNetworkRule and deletes it. Returns an error if one occurs.
func (c *mariadbVirtualNetworkRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mariadbvirtualnetworkrules").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *mariadbVirtualNetworkRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mariadbvirtualnetworkrules").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched mariadbVirtualNetworkRule.
func (c *mariadbVirtualNetworkRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MariadbVirtualNetworkRule, err error) {
	result = &v1alpha1.MariadbVirtualNetworkRule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mariadbvirtualnetworkrules").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
