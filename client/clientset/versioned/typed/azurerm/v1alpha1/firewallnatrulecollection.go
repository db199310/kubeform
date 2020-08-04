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

// FirewallNATRuleCollectionsGetter has a method to return a FirewallNATRuleCollectionInterface.
// A group's client should implement this interface.
type FirewallNATRuleCollectionsGetter interface {
	FirewallNATRuleCollections(namespace string) FirewallNATRuleCollectionInterface
}

// FirewallNATRuleCollectionInterface has methods to work with FirewallNATRuleCollection resources.
type FirewallNATRuleCollectionInterface interface {
	Create(*v1alpha1.FirewallNATRuleCollection) (*v1alpha1.FirewallNATRuleCollection, error)
	Update(*v1alpha1.FirewallNATRuleCollection) (*v1alpha1.FirewallNATRuleCollection, error)
	UpdateStatus(*v1alpha1.FirewallNATRuleCollection) (*v1alpha1.FirewallNATRuleCollection, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.FirewallNATRuleCollection, error)
	List(opts v1.ListOptions) (*v1alpha1.FirewallNATRuleCollectionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FirewallNATRuleCollection, err error)
	FirewallNATRuleCollectionExpansion
}

// firewallNATRuleCollections implements FirewallNATRuleCollectionInterface
type firewallNATRuleCollections struct {
	client rest.Interface
	ns     string
}

// newFirewallNATRuleCollections returns a FirewallNATRuleCollections
func newFirewallNATRuleCollections(c *AzurermV1alpha1Client, namespace string) *firewallNATRuleCollections {
	return &firewallNATRuleCollections{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the firewallNATRuleCollection, and returns the corresponding firewallNATRuleCollection object, and an error if there is any.
func (c *firewallNATRuleCollections) Get(name string, options v1.GetOptions) (result *v1alpha1.FirewallNATRuleCollection, err error) {
	result = &v1alpha1.FirewallNATRuleCollection{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("firewallnatrulecollections").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FirewallNATRuleCollections that match those selectors.
func (c *firewallNATRuleCollections) List(opts v1.ListOptions) (result *v1alpha1.FirewallNATRuleCollectionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.FirewallNATRuleCollectionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("firewallnatrulecollections").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested firewallNATRuleCollections.
func (c *firewallNATRuleCollections) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("firewallnatrulecollections").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a firewallNATRuleCollection and creates it.  Returns the server's representation of the firewallNATRuleCollection, and an error, if there is any.
func (c *firewallNATRuleCollections) Create(firewallNATRuleCollection *v1alpha1.FirewallNATRuleCollection) (result *v1alpha1.FirewallNATRuleCollection, err error) {
	result = &v1alpha1.FirewallNATRuleCollection{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("firewallnatrulecollections").
		Body(firewallNATRuleCollection).
		Do().
		Into(result)
	return
}

// Update takes the representation of a firewallNATRuleCollection and updates it. Returns the server's representation of the firewallNATRuleCollection, and an error, if there is any.
func (c *firewallNATRuleCollections) Update(firewallNATRuleCollection *v1alpha1.FirewallNATRuleCollection) (result *v1alpha1.FirewallNATRuleCollection, err error) {
	result = &v1alpha1.FirewallNATRuleCollection{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("firewallnatrulecollections").
		Name(firewallNATRuleCollection.Name).
		Body(firewallNATRuleCollection).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *firewallNATRuleCollections) UpdateStatus(firewallNATRuleCollection *v1alpha1.FirewallNATRuleCollection) (result *v1alpha1.FirewallNATRuleCollection, err error) {
	result = &v1alpha1.FirewallNATRuleCollection{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("firewallnatrulecollections").
		Name(firewallNATRuleCollection.Name).
		SubResource("status").
		Body(firewallNATRuleCollection).
		Do().
		Into(result)
	return
}

// Delete takes name of the firewallNATRuleCollection and deletes it. Returns an error if one occurs.
func (c *firewallNATRuleCollections) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("firewallnatrulecollections").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *firewallNATRuleCollections) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("firewallnatrulecollections").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched firewallNATRuleCollection.
func (c *firewallNATRuleCollections) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FirewallNATRuleCollection, err error) {
	result = &v1alpha1.FirewallNATRuleCollection{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("firewallnatrulecollections").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
