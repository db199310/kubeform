/*
Copyright 2019 The Kubeform Authors.

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

// ApplicationInsightsApiKeysGetter has a method to return a ApplicationInsightsApiKeyInterface.
// A group's client should implement this interface.
type ApplicationInsightsApiKeysGetter interface {
	ApplicationInsightsApiKeys() ApplicationInsightsApiKeyInterface
}

// ApplicationInsightsApiKeyInterface has methods to work with ApplicationInsightsApiKey resources.
type ApplicationInsightsApiKeyInterface interface {
	Create(*v1alpha1.ApplicationInsightsApiKey) (*v1alpha1.ApplicationInsightsApiKey, error)
	Update(*v1alpha1.ApplicationInsightsApiKey) (*v1alpha1.ApplicationInsightsApiKey, error)
	UpdateStatus(*v1alpha1.ApplicationInsightsApiKey) (*v1alpha1.ApplicationInsightsApiKey, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ApplicationInsightsApiKey, error)
	List(opts v1.ListOptions) (*v1alpha1.ApplicationInsightsApiKeyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApplicationInsightsApiKey, err error)
	ApplicationInsightsApiKeyExpansion
}

// applicationInsightsApiKeys implements ApplicationInsightsApiKeyInterface
type applicationInsightsApiKeys struct {
	client rest.Interface
}

// newApplicationInsightsApiKeys returns a ApplicationInsightsApiKeys
func newApplicationInsightsApiKeys(c *AzurermV1alpha1Client) *applicationInsightsApiKeys {
	return &applicationInsightsApiKeys{
		client: c.RESTClient(),
	}
}

// Get takes name of the applicationInsightsApiKey, and returns the corresponding applicationInsightsApiKey object, and an error if there is any.
func (c *applicationInsightsApiKeys) Get(name string, options v1.GetOptions) (result *v1alpha1.ApplicationInsightsApiKey, err error) {
	result = &v1alpha1.ApplicationInsightsApiKey{}
	err = c.client.Get().
		Resource("applicationinsightsapikeys").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApplicationInsightsApiKeys that match those selectors.
func (c *applicationInsightsApiKeys) List(opts v1.ListOptions) (result *v1alpha1.ApplicationInsightsApiKeyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ApplicationInsightsApiKeyList{}
	err = c.client.Get().
		Resource("applicationinsightsapikeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested applicationInsightsApiKeys.
func (c *applicationInsightsApiKeys) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("applicationinsightsapikeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a applicationInsightsApiKey and creates it.  Returns the server's representation of the applicationInsightsApiKey, and an error, if there is any.
func (c *applicationInsightsApiKeys) Create(applicationInsightsApiKey *v1alpha1.ApplicationInsightsApiKey) (result *v1alpha1.ApplicationInsightsApiKey, err error) {
	result = &v1alpha1.ApplicationInsightsApiKey{}
	err = c.client.Post().
		Resource("applicationinsightsapikeys").
		Body(applicationInsightsApiKey).
		Do().
		Into(result)
	return
}

// Update takes the representation of a applicationInsightsApiKey and updates it. Returns the server's representation of the applicationInsightsApiKey, and an error, if there is any.
func (c *applicationInsightsApiKeys) Update(applicationInsightsApiKey *v1alpha1.ApplicationInsightsApiKey) (result *v1alpha1.ApplicationInsightsApiKey, err error) {
	result = &v1alpha1.ApplicationInsightsApiKey{}
	err = c.client.Put().
		Resource("applicationinsightsapikeys").
		Name(applicationInsightsApiKey.Name).
		Body(applicationInsightsApiKey).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *applicationInsightsApiKeys) UpdateStatus(applicationInsightsApiKey *v1alpha1.ApplicationInsightsApiKey) (result *v1alpha1.ApplicationInsightsApiKey, err error) {
	result = &v1alpha1.ApplicationInsightsApiKey{}
	err = c.client.Put().
		Resource("applicationinsightsapikeys").
		Name(applicationInsightsApiKey.Name).
		SubResource("status").
		Body(applicationInsightsApiKey).
		Do().
		Into(result)
	return
}

// Delete takes name of the applicationInsightsApiKey and deletes it. Returns an error if one occurs.
func (c *applicationInsightsApiKeys) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("applicationinsightsapikeys").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *applicationInsightsApiKeys) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("applicationinsightsapikeys").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched applicationInsightsApiKey.
func (c *applicationInsightsApiKeys) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApplicationInsightsApiKey, err error) {
	result = &v1alpha1.ApplicationInsightsApiKey{}
	err = c.client.Patch(pt).
		Resource("applicationinsightsapikeys").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}