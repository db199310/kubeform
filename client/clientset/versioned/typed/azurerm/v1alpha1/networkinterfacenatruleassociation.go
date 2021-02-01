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

// NetworkInterfaceNATRuleAssociationsGetter has a method to return a NetworkInterfaceNATRuleAssociationInterface.
// A group's client should implement this interface.
type NetworkInterfaceNATRuleAssociationsGetter interface {
	NetworkInterfaceNATRuleAssociations(namespace string) NetworkInterfaceNATRuleAssociationInterface
}

// NetworkInterfaceNATRuleAssociationInterface has methods to work with NetworkInterfaceNATRuleAssociation resources.
type NetworkInterfaceNATRuleAssociationInterface interface {
	Create(ctx context.Context, networkInterfaceNATRuleAssociation *v1alpha1.NetworkInterfaceNATRuleAssociation, opts v1.CreateOptions) (*v1alpha1.NetworkInterfaceNATRuleAssociation, error)
	Update(ctx context.Context, networkInterfaceNATRuleAssociation *v1alpha1.NetworkInterfaceNATRuleAssociation, opts v1.UpdateOptions) (*v1alpha1.NetworkInterfaceNATRuleAssociation, error)
	UpdateStatus(ctx context.Context, networkInterfaceNATRuleAssociation *v1alpha1.NetworkInterfaceNATRuleAssociation, opts v1.UpdateOptions) (*v1alpha1.NetworkInterfaceNATRuleAssociation, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.NetworkInterfaceNATRuleAssociation, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.NetworkInterfaceNATRuleAssociationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NetworkInterfaceNATRuleAssociation, err error)
	NetworkInterfaceNATRuleAssociationExpansion
}

// networkInterfaceNATRuleAssociations implements NetworkInterfaceNATRuleAssociationInterface
type networkInterfaceNATRuleAssociations struct {
	client rest.Interface
	ns     string
}

// newNetworkInterfaceNATRuleAssociations returns a NetworkInterfaceNATRuleAssociations
func newNetworkInterfaceNATRuleAssociations(c *AzurermV1alpha1Client, namespace string) *networkInterfaceNATRuleAssociations {
	return &networkInterfaceNATRuleAssociations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the networkInterfaceNATRuleAssociation, and returns the corresponding networkInterfaceNATRuleAssociation object, and an error if there is any.
func (c *networkInterfaceNATRuleAssociations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NetworkInterfaceNATRuleAssociation, err error) {
	result = &v1alpha1.NetworkInterfaceNATRuleAssociation{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkinterfacenatruleassociations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NetworkInterfaceNATRuleAssociations that match those selectors.
func (c *networkInterfaceNATRuleAssociations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NetworkInterfaceNATRuleAssociationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NetworkInterfaceNATRuleAssociationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkinterfacenatruleassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested networkInterfaceNATRuleAssociations.
func (c *networkInterfaceNATRuleAssociations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("networkinterfacenatruleassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a networkInterfaceNATRuleAssociation and creates it.  Returns the server's representation of the networkInterfaceNATRuleAssociation, and an error, if there is any.
func (c *networkInterfaceNATRuleAssociations) Create(ctx context.Context, networkInterfaceNATRuleAssociation *v1alpha1.NetworkInterfaceNATRuleAssociation, opts v1.CreateOptions) (result *v1alpha1.NetworkInterfaceNATRuleAssociation, err error) {
	result = &v1alpha1.NetworkInterfaceNATRuleAssociation{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("networkinterfacenatruleassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(networkInterfaceNATRuleAssociation).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a networkInterfaceNATRuleAssociation and updates it. Returns the server's representation of the networkInterfaceNATRuleAssociation, and an error, if there is any.
func (c *networkInterfaceNATRuleAssociations) Update(ctx context.Context, networkInterfaceNATRuleAssociation *v1alpha1.NetworkInterfaceNATRuleAssociation, opts v1.UpdateOptions) (result *v1alpha1.NetworkInterfaceNATRuleAssociation, err error) {
	result = &v1alpha1.NetworkInterfaceNATRuleAssociation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("networkinterfacenatruleassociations").
		Name(networkInterfaceNATRuleAssociation.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(networkInterfaceNATRuleAssociation).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *networkInterfaceNATRuleAssociations) UpdateStatus(ctx context.Context, networkInterfaceNATRuleAssociation *v1alpha1.NetworkInterfaceNATRuleAssociation, opts v1.UpdateOptions) (result *v1alpha1.NetworkInterfaceNATRuleAssociation, err error) {
	result = &v1alpha1.NetworkInterfaceNATRuleAssociation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("networkinterfacenatruleassociations").
		Name(networkInterfaceNATRuleAssociation.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(networkInterfaceNATRuleAssociation).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the networkInterfaceNATRuleAssociation and deletes it. Returns an error if one occurs.
func (c *networkInterfaceNATRuleAssociations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networkinterfacenatruleassociations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *networkInterfaceNATRuleAssociations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networkinterfacenatruleassociations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched networkInterfaceNATRuleAssociation.
func (c *networkInterfaceNATRuleAssociations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NetworkInterfaceNATRuleAssociation, err error) {
	result = &v1alpha1.NetworkInterfaceNATRuleAssociation{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("networkinterfacenatruleassociations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
