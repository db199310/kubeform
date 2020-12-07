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

// FunctionAppsGetter has a method to return a FunctionAppInterface.
// A group's client should implement this interface.
type FunctionAppsGetter interface {
	FunctionApps(namespace string) FunctionAppInterface
}

// FunctionAppInterface has methods to work with FunctionApp resources.
type FunctionAppInterface interface {
	Create(ctx context.Context, functionApp *v1alpha1.FunctionApp, opts v1.CreateOptions) (*v1alpha1.FunctionApp, error)
	Update(ctx context.Context, functionApp *v1alpha1.FunctionApp, opts v1.UpdateOptions) (*v1alpha1.FunctionApp, error)
	UpdateStatus(ctx context.Context, functionApp *v1alpha1.FunctionApp, opts v1.UpdateOptions) (*v1alpha1.FunctionApp, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.FunctionApp, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.FunctionAppList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FunctionApp, err error)
	FunctionAppExpansion
}

// functionApps implements FunctionAppInterface
type functionApps struct {
	client rest.Interface
	ns     string
}

// newFunctionApps returns a FunctionApps
func newFunctionApps(c *AzurermV1alpha1Client, namespace string) *functionApps {
	return &functionApps{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the functionApp, and returns the corresponding functionApp object, and an error if there is any.
func (c *functionApps) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FunctionApp, err error) {
	result = &v1alpha1.FunctionApp{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("functionapps").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FunctionApps that match those selectors.
func (c *functionApps) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FunctionAppList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.FunctionAppList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("functionapps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested functionApps.
func (c *functionApps) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("functionapps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a functionApp and creates it.  Returns the server's representation of the functionApp, and an error, if there is any.
func (c *functionApps) Create(ctx context.Context, functionApp *v1alpha1.FunctionApp, opts v1.CreateOptions) (result *v1alpha1.FunctionApp, err error) {
	result = &v1alpha1.FunctionApp{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("functionapps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(functionApp).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a functionApp and updates it. Returns the server's representation of the functionApp, and an error, if there is any.
func (c *functionApps) Update(ctx context.Context, functionApp *v1alpha1.FunctionApp, opts v1.UpdateOptions) (result *v1alpha1.FunctionApp, err error) {
	result = &v1alpha1.FunctionApp{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("functionapps").
		Name(functionApp.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(functionApp).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *functionApps) UpdateStatus(ctx context.Context, functionApp *v1alpha1.FunctionApp, opts v1.UpdateOptions) (result *v1alpha1.FunctionApp, err error) {
	result = &v1alpha1.FunctionApp{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("functionapps").
		Name(functionApp.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(functionApp).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the functionApp and deletes it. Returns an error if one occurs.
func (c *functionApps) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("functionapps").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *functionApps) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("functionapps").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched functionApp.
func (c *functionApps) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FunctionApp, err error) {
	result = &v1alpha1.FunctionApp{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("functionapps").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
