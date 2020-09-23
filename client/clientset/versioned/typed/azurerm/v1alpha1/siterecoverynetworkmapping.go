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

// SiteRecoveryNetworkMappingsGetter has a method to return a SiteRecoveryNetworkMappingInterface.
// A group's client should implement this interface.
type SiteRecoveryNetworkMappingsGetter interface {
	SiteRecoveryNetworkMappings(namespace string) SiteRecoveryNetworkMappingInterface
}

// SiteRecoveryNetworkMappingInterface has methods to work with SiteRecoveryNetworkMapping resources.
type SiteRecoveryNetworkMappingInterface interface {
	Create(*v1alpha1.SiteRecoveryNetworkMapping) (*v1alpha1.SiteRecoveryNetworkMapping, error)
	Update(*v1alpha1.SiteRecoveryNetworkMapping) (*v1alpha1.SiteRecoveryNetworkMapping, error)
	UpdateStatus(*v1alpha1.SiteRecoveryNetworkMapping) (*v1alpha1.SiteRecoveryNetworkMapping, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SiteRecoveryNetworkMapping, error)
	List(opts v1.ListOptions) (*v1alpha1.SiteRecoveryNetworkMappingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SiteRecoveryNetworkMapping, err error)
	SiteRecoveryNetworkMappingExpansion
}

// siteRecoveryNetworkMappings implements SiteRecoveryNetworkMappingInterface
type siteRecoveryNetworkMappings struct {
	client rest.Interface
	ns     string
}

// newSiteRecoveryNetworkMappings returns a SiteRecoveryNetworkMappings
func newSiteRecoveryNetworkMappings(c *AzurermV1alpha1Client, namespace string) *siteRecoveryNetworkMappings {
	return &siteRecoveryNetworkMappings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the siteRecoveryNetworkMapping, and returns the corresponding siteRecoveryNetworkMapping object, and an error if there is any.
func (c *siteRecoveryNetworkMappings) Get(name string, options v1.GetOptions) (result *v1alpha1.SiteRecoveryNetworkMapping, err error) {
	result = &v1alpha1.SiteRecoveryNetworkMapping{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("siterecoverynetworkmappings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SiteRecoveryNetworkMappings that match those selectors.
func (c *siteRecoveryNetworkMappings) List(opts v1.ListOptions) (result *v1alpha1.SiteRecoveryNetworkMappingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SiteRecoveryNetworkMappingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("siterecoverynetworkmappings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested siteRecoveryNetworkMappings.
func (c *siteRecoveryNetworkMappings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("siterecoverynetworkmappings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a siteRecoveryNetworkMapping and creates it.  Returns the server's representation of the siteRecoveryNetworkMapping, and an error, if there is any.
func (c *siteRecoveryNetworkMappings) Create(siteRecoveryNetworkMapping *v1alpha1.SiteRecoveryNetworkMapping) (result *v1alpha1.SiteRecoveryNetworkMapping, err error) {
	result = &v1alpha1.SiteRecoveryNetworkMapping{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("siterecoverynetworkmappings").
		Body(siteRecoveryNetworkMapping).
		Do().
		Into(result)
	return
}

// Update takes the representation of a siteRecoveryNetworkMapping and updates it. Returns the server's representation of the siteRecoveryNetworkMapping, and an error, if there is any.
func (c *siteRecoveryNetworkMappings) Update(siteRecoveryNetworkMapping *v1alpha1.SiteRecoveryNetworkMapping) (result *v1alpha1.SiteRecoveryNetworkMapping, err error) {
	result = &v1alpha1.SiteRecoveryNetworkMapping{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("siterecoverynetworkmappings").
		Name(siteRecoveryNetworkMapping.Name).
		Body(siteRecoveryNetworkMapping).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *siteRecoveryNetworkMappings) UpdateStatus(siteRecoveryNetworkMapping *v1alpha1.SiteRecoveryNetworkMapping) (result *v1alpha1.SiteRecoveryNetworkMapping, err error) {
	result = &v1alpha1.SiteRecoveryNetworkMapping{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("siterecoverynetworkmappings").
		Name(siteRecoveryNetworkMapping.Name).
		SubResource("status").
		Body(siteRecoveryNetworkMapping).
		Do().
		Into(result)
	return
}

// Delete takes name of the siteRecoveryNetworkMapping and deletes it. Returns an error if one occurs.
func (c *siteRecoveryNetworkMappings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("siterecoverynetworkmappings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *siteRecoveryNetworkMappings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("siterecoverynetworkmappings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched siteRecoveryNetworkMapping.
func (c *siteRecoveryNetworkMappings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SiteRecoveryNetworkMapping, err error) {
	result = &v1alpha1.SiteRecoveryNetworkMapping{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("siterecoverynetworkmappings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}