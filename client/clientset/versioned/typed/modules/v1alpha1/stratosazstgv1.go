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
	v1alpha1 "kubeform.dev/kubeform/apis/modules/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// StratosAzStgv1sGetter has a method to return a StratosAzStgv1Interface.
// A group's client should implement this interface.
type StratosAzStgv1sGetter interface {
	StratosAzStgv1s(namespace string) StratosAzStgv1Interface
}

// StratosAzStgv1Interface has methods to work with StratosAzStgv1 resources.
type StratosAzStgv1Interface interface {
	Create(ctx context.Context, stratosAzStgv1 *v1alpha1.StratosAzStgv1, opts v1.CreateOptions) (*v1alpha1.StratosAzStgv1, error)
	Update(ctx context.Context, stratosAzStgv1 *v1alpha1.StratosAzStgv1, opts v1.UpdateOptions) (*v1alpha1.StratosAzStgv1, error)
	UpdateStatus(ctx context.Context, stratosAzStgv1 *v1alpha1.StratosAzStgv1, opts v1.UpdateOptions) (*v1alpha1.StratosAzStgv1, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.StratosAzStgv1, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.StratosAzStgv1List, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.StratosAzStgv1, err error)
	StratosAzStgv1Expansion
}

// stratosAzStgv1s implements StratosAzStgv1Interface
type stratosAzStgv1s struct {
	client rest.Interface
	ns     string
}

// newStratosAzStgv1s returns a StratosAzStgv1s
func newStratosAzStgv1s(c *ModulesV1alpha1Client, namespace string) *stratosAzStgv1s {
	return &stratosAzStgv1s{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the stratosAzStgv1, and returns the corresponding stratosAzStgv1 object, and an error if there is any.
func (c *stratosAzStgv1s) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.StratosAzStgv1, err error) {
	result = &v1alpha1.StratosAzStgv1{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("stratosazstgv1s").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of StratosAzStgv1s that match those selectors.
func (c *stratosAzStgv1s) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.StratosAzStgv1List, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.StratosAzStgv1List{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("stratosazstgv1s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested stratosAzStgv1s.
func (c *stratosAzStgv1s) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("stratosazstgv1s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a stratosAzStgv1 and creates it.  Returns the server's representation of the stratosAzStgv1, and an error, if there is any.
func (c *stratosAzStgv1s) Create(ctx context.Context, stratosAzStgv1 *v1alpha1.StratosAzStgv1, opts v1.CreateOptions) (result *v1alpha1.StratosAzStgv1, err error) {
	result = &v1alpha1.StratosAzStgv1{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("stratosazstgv1s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(stratosAzStgv1).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a stratosAzStgv1 and updates it. Returns the server's representation of the stratosAzStgv1, and an error, if there is any.
func (c *stratosAzStgv1s) Update(ctx context.Context, stratosAzStgv1 *v1alpha1.StratosAzStgv1, opts v1.UpdateOptions) (result *v1alpha1.StratosAzStgv1, err error) {
	result = &v1alpha1.StratosAzStgv1{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("stratosazstgv1s").
		Name(stratosAzStgv1.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(stratosAzStgv1).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *stratosAzStgv1s) UpdateStatus(ctx context.Context, stratosAzStgv1 *v1alpha1.StratosAzStgv1, opts v1.UpdateOptions) (result *v1alpha1.StratosAzStgv1, err error) {
	result = &v1alpha1.StratosAzStgv1{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("stratosazstgv1s").
		Name(stratosAzStgv1.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(stratosAzStgv1).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the stratosAzStgv1 and deletes it. Returns an error if one occurs.
func (c *stratosAzStgv1s) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("stratosazstgv1s").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *stratosAzStgv1s) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("stratosazstgv1s").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched stratosAzStgv1.
func (c *stratosAzStgv1s) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.StratosAzStgv1, err error) {
	result = &v1alpha1.StratosAzStgv1{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("stratosazstgv1s").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
