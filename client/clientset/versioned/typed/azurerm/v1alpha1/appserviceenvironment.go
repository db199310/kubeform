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

// AppServiceEnvironmentsGetter has a method to return a AppServiceEnvironmentInterface.
// A group's client should implement this interface.
type AppServiceEnvironmentsGetter interface {
	AppServiceEnvironments(namespace string) AppServiceEnvironmentInterface
}

// AppServiceEnvironmentInterface has methods to work with AppServiceEnvironment resources.
type AppServiceEnvironmentInterface interface {
	Create(*v1alpha1.AppServiceEnvironment) (*v1alpha1.AppServiceEnvironment, error)
	Update(*v1alpha1.AppServiceEnvironment) (*v1alpha1.AppServiceEnvironment, error)
	UpdateStatus(*v1alpha1.AppServiceEnvironment) (*v1alpha1.AppServiceEnvironment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AppServiceEnvironment, error)
	List(opts v1.ListOptions) (*v1alpha1.AppServiceEnvironmentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AppServiceEnvironment, err error)
	AppServiceEnvironmentExpansion
}

// appServiceEnvironments implements AppServiceEnvironmentInterface
type appServiceEnvironments struct {
	client rest.Interface
	ns     string
}

// newAppServiceEnvironments returns a AppServiceEnvironments
func newAppServiceEnvironments(c *AzurermV1alpha1Client, namespace string) *appServiceEnvironments {
	return &appServiceEnvironments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the appServiceEnvironment, and returns the corresponding appServiceEnvironment object, and an error if there is any.
func (c *appServiceEnvironments) Get(name string, options v1.GetOptions) (result *v1alpha1.AppServiceEnvironment, err error) {
	result = &v1alpha1.AppServiceEnvironment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("appserviceenvironments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AppServiceEnvironments that match those selectors.
func (c *appServiceEnvironments) List(opts v1.ListOptions) (result *v1alpha1.AppServiceEnvironmentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AppServiceEnvironmentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("appserviceenvironments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested appServiceEnvironments.
func (c *appServiceEnvironments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("appserviceenvironments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a appServiceEnvironment and creates it.  Returns the server's representation of the appServiceEnvironment, and an error, if there is any.
func (c *appServiceEnvironments) Create(appServiceEnvironment *v1alpha1.AppServiceEnvironment) (result *v1alpha1.AppServiceEnvironment, err error) {
	result = &v1alpha1.AppServiceEnvironment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("appserviceenvironments").
		Body(appServiceEnvironment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a appServiceEnvironment and updates it. Returns the server's representation of the appServiceEnvironment, and an error, if there is any.
func (c *appServiceEnvironments) Update(appServiceEnvironment *v1alpha1.AppServiceEnvironment) (result *v1alpha1.AppServiceEnvironment, err error) {
	result = &v1alpha1.AppServiceEnvironment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("appserviceenvironments").
		Name(appServiceEnvironment.Name).
		Body(appServiceEnvironment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *appServiceEnvironments) UpdateStatus(appServiceEnvironment *v1alpha1.AppServiceEnvironment) (result *v1alpha1.AppServiceEnvironment, err error) {
	result = &v1alpha1.AppServiceEnvironment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("appserviceenvironments").
		Name(appServiceEnvironment.Name).
		SubResource("status").
		Body(appServiceEnvironment).
		Do().
		Into(result)
	return
}

// Delete takes name of the appServiceEnvironment and deletes it. Returns an error if one occurs.
func (c *appServiceEnvironments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("appserviceenvironments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *appServiceEnvironments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("appserviceenvironments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched appServiceEnvironment.
func (c *appServiceEnvironments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AppServiceEnvironment, err error) {
	result = &v1alpha1.AppServiceEnvironment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("appserviceenvironments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
