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

// SDPAzadfv1sGetter has a method to return a SDPAzadfv1Interface.
// A group's client should implement this interface.
type SDPAzadfv1sGetter interface {
	SDPAzadfv1s(namespace string) SDPAzadfv1Interface
}

// SDPAzadfv1Interface has methods to work with SDPAzadfv1 resources.
type SDPAzadfv1Interface interface {
	Create(ctx context.Context, sDPAzadfv1 *v1alpha1.SDPAzadfv1, opts v1.CreateOptions) (*v1alpha1.SDPAzadfv1, error)
	Update(ctx context.Context, sDPAzadfv1 *v1alpha1.SDPAzadfv1, opts v1.UpdateOptions) (*v1alpha1.SDPAzadfv1, error)
	UpdateStatus(ctx context.Context, sDPAzadfv1 *v1alpha1.SDPAzadfv1, opts v1.UpdateOptions) (*v1alpha1.SDPAzadfv1, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.SDPAzadfv1, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.SDPAzadfv1List, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SDPAzadfv1, err error)
	SDPAzadfv1Expansion
}

// sDPAzadfv1s implements SDPAzadfv1Interface
type sDPAzadfv1s struct {
	client rest.Interface
	ns     string
}

// newSDPAzadfv1s returns a SDPAzadfv1s
func newSDPAzadfv1s(c *ModulesV1alpha1Client, namespace string) *sDPAzadfv1s {
	return &sDPAzadfv1s{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sDPAzadfv1, and returns the corresponding sDPAzadfv1 object, and an error if there is any.
func (c *sDPAzadfv1s) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SDPAzadfv1, err error) {
	result = &v1alpha1.SDPAzadfv1{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sdpazadfv1s").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SDPAzadfv1s that match those selectors.
func (c *sDPAzadfv1s) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SDPAzadfv1List, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SDPAzadfv1List{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sdpazadfv1s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sDPAzadfv1s.
func (c *sDPAzadfv1s) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sdpazadfv1s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a sDPAzadfv1 and creates it.  Returns the server's representation of the sDPAzadfv1, and an error, if there is any.
func (c *sDPAzadfv1s) Create(ctx context.Context, sDPAzadfv1 *v1alpha1.SDPAzadfv1, opts v1.CreateOptions) (result *v1alpha1.SDPAzadfv1, err error) {
	result = &v1alpha1.SDPAzadfv1{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sdpazadfv1s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sDPAzadfv1).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a sDPAzadfv1 and updates it. Returns the server's representation of the sDPAzadfv1, and an error, if there is any.
func (c *sDPAzadfv1s) Update(ctx context.Context, sDPAzadfv1 *v1alpha1.SDPAzadfv1, opts v1.UpdateOptions) (result *v1alpha1.SDPAzadfv1, err error) {
	result = &v1alpha1.SDPAzadfv1{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sdpazadfv1s").
		Name(sDPAzadfv1.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sDPAzadfv1).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *sDPAzadfv1s) UpdateStatus(ctx context.Context, sDPAzadfv1 *v1alpha1.SDPAzadfv1, opts v1.UpdateOptions) (result *v1alpha1.SDPAzadfv1, err error) {
	result = &v1alpha1.SDPAzadfv1{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sdpazadfv1s").
		Name(sDPAzadfv1.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sDPAzadfv1).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the sDPAzadfv1 and deletes it. Returns an error if one occurs.
func (c *sDPAzadfv1s) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sdpazadfv1s").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sDPAzadfv1s) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sdpazadfv1s").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched sDPAzadfv1.
func (c *sDPAzadfv1s) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SDPAzadfv1, err error) {
	result = &v1alpha1.SDPAzadfv1{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sdpazadfv1s").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
