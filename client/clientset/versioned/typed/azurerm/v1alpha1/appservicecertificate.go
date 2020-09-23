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

// AppServiceCertificatesGetter has a method to return a AppServiceCertificateInterface.
// A group's client should implement this interface.
type AppServiceCertificatesGetter interface {
	AppServiceCertificates(namespace string) AppServiceCertificateInterface
}

// AppServiceCertificateInterface has methods to work with AppServiceCertificate resources.
type AppServiceCertificateInterface interface {
	Create(*v1alpha1.AppServiceCertificate) (*v1alpha1.AppServiceCertificate, error)
	Update(*v1alpha1.AppServiceCertificate) (*v1alpha1.AppServiceCertificate, error)
	UpdateStatus(*v1alpha1.AppServiceCertificate) (*v1alpha1.AppServiceCertificate, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AppServiceCertificate, error)
	List(opts v1.ListOptions) (*v1alpha1.AppServiceCertificateList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AppServiceCertificate, err error)
	AppServiceCertificateExpansion
}

// appServiceCertificates implements AppServiceCertificateInterface
type appServiceCertificates struct {
	client rest.Interface
	ns     string
}

// newAppServiceCertificates returns a AppServiceCertificates
func newAppServiceCertificates(c *AzurermV1alpha1Client, namespace string) *appServiceCertificates {
	return &appServiceCertificates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the appServiceCertificate, and returns the corresponding appServiceCertificate object, and an error if there is any.
func (c *appServiceCertificates) Get(name string, options v1.GetOptions) (result *v1alpha1.AppServiceCertificate, err error) {
	result = &v1alpha1.AppServiceCertificate{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("appservicecertificates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AppServiceCertificates that match those selectors.
func (c *appServiceCertificates) List(opts v1.ListOptions) (result *v1alpha1.AppServiceCertificateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AppServiceCertificateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("appservicecertificates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested appServiceCertificates.
func (c *appServiceCertificates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("appservicecertificates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a appServiceCertificate and creates it.  Returns the server's representation of the appServiceCertificate, and an error, if there is any.
func (c *appServiceCertificates) Create(appServiceCertificate *v1alpha1.AppServiceCertificate) (result *v1alpha1.AppServiceCertificate, err error) {
	result = &v1alpha1.AppServiceCertificate{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("appservicecertificates").
		Body(appServiceCertificate).
		Do().
		Into(result)
	return
}

// Update takes the representation of a appServiceCertificate and updates it. Returns the server's representation of the appServiceCertificate, and an error, if there is any.
func (c *appServiceCertificates) Update(appServiceCertificate *v1alpha1.AppServiceCertificate) (result *v1alpha1.AppServiceCertificate, err error) {
	result = &v1alpha1.AppServiceCertificate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("appservicecertificates").
		Name(appServiceCertificate.Name).
		Body(appServiceCertificate).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *appServiceCertificates) UpdateStatus(appServiceCertificate *v1alpha1.AppServiceCertificate) (result *v1alpha1.AppServiceCertificate, err error) {
	result = &v1alpha1.AppServiceCertificate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("appservicecertificates").
		Name(appServiceCertificate.Name).
		SubResource("status").
		Body(appServiceCertificate).
		Do().
		Into(result)
	return
}

// Delete takes name of the appServiceCertificate and deletes it. Returns an error if one occurs.
func (c *appServiceCertificates) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("appservicecertificates").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *appServiceCertificates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("appservicecertificates").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched appServiceCertificate.
func (c *appServiceCertificates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AppServiceCertificate, err error) {
	result = &v1alpha1.AppServiceCertificate{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("appservicecertificates").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}