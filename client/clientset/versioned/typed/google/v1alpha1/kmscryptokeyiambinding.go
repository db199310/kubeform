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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// KmsCryptoKeyIamBindingsGetter has a method to return a KmsCryptoKeyIamBindingInterface.
// A group's client should implement this interface.
type KmsCryptoKeyIamBindingsGetter interface {
	KmsCryptoKeyIamBindings(namespace string) KmsCryptoKeyIamBindingInterface
}

// KmsCryptoKeyIamBindingInterface has methods to work with KmsCryptoKeyIamBinding resources.
type KmsCryptoKeyIamBindingInterface interface {
	Create(*v1alpha1.KmsCryptoKeyIamBinding) (*v1alpha1.KmsCryptoKeyIamBinding, error)
	Update(*v1alpha1.KmsCryptoKeyIamBinding) (*v1alpha1.KmsCryptoKeyIamBinding, error)
	UpdateStatus(*v1alpha1.KmsCryptoKeyIamBinding) (*v1alpha1.KmsCryptoKeyIamBinding, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.KmsCryptoKeyIamBinding, error)
	List(opts v1.ListOptions) (*v1alpha1.KmsCryptoKeyIamBindingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KmsCryptoKeyIamBinding, err error)
	KmsCryptoKeyIamBindingExpansion
}

// kmsCryptoKeyIamBindings implements KmsCryptoKeyIamBindingInterface
type kmsCryptoKeyIamBindings struct {
	client rest.Interface
	ns     string
}

// newKmsCryptoKeyIamBindings returns a KmsCryptoKeyIamBindings
func newKmsCryptoKeyIamBindings(c *GoogleV1alpha1Client, namespace string) *kmsCryptoKeyIamBindings {
	return &kmsCryptoKeyIamBindings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kmsCryptoKeyIamBinding, and returns the corresponding kmsCryptoKeyIamBinding object, and an error if there is any.
func (c *kmsCryptoKeyIamBindings) Get(name string, options v1.GetOptions) (result *v1alpha1.KmsCryptoKeyIamBinding, err error) {
	result = &v1alpha1.KmsCryptoKeyIamBinding{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kmscryptokeyiambindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KmsCryptoKeyIamBindings that match those selectors.
func (c *kmsCryptoKeyIamBindings) List(opts v1.ListOptions) (result *v1alpha1.KmsCryptoKeyIamBindingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.KmsCryptoKeyIamBindingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kmscryptokeyiambindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kmsCryptoKeyIamBindings.
func (c *kmsCryptoKeyIamBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kmscryptokeyiambindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a kmsCryptoKeyIamBinding and creates it.  Returns the server's representation of the kmsCryptoKeyIamBinding, and an error, if there is any.
func (c *kmsCryptoKeyIamBindings) Create(kmsCryptoKeyIamBinding *v1alpha1.KmsCryptoKeyIamBinding) (result *v1alpha1.KmsCryptoKeyIamBinding, err error) {
	result = &v1alpha1.KmsCryptoKeyIamBinding{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kmscryptokeyiambindings").
		Body(kmsCryptoKeyIamBinding).
		Do().
		Into(result)
	return
}

// Update takes the representation of a kmsCryptoKeyIamBinding and updates it. Returns the server's representation of the kmsCryptoKeyIamBinding, and an error, if there is any.
func (c *kmsCryptoKeyIamBindings) Update(kmsCryptoKeyIamBinding *v1alpha1.KmsCryptoKeyIamBinding) (result *v1alpha1.KmsCryptoKeyIamBinding, err error) {
	result = &v1alpha1.KmsCryptoKeyIamBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kmscryptokeyiambindings").
		Name(kmsCryptoKeyIamBinding.Name).
		Body(kmsCryptoKeyIamBinding).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *kmsCryptoKeyIamBindings) UpdateStatus(kmsCryptoKeyIamBinding *v1alpha1.KmsCryptoKeyIamBinding) (result *v1alpha1.KmsCryptoKeyIamBinding, err error) {
	result = &v1alpha1.KmsCryptoKeyIamBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kmscryptokeyiambindings").
		Name(kmsCryptoKeyIamBinding.Name).
		SubResource("status").
		Body(kmsCryptoKeyIamBinding).
		Do().
		Into(result)
	return
}

// Delete takes name of the kmsCryptoKeyIamBinding and deletes it. Returns an error if one occurs.
func (c *kmsCryptoKeyIamBindings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kmscryptokeyiambindings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kmsCryptoKeyIamBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kmscryptokeyiambindings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched kmsCryptoKeyIamBinding.
func (c *kmsCryptoKeyIamBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KmsCryptoKeyIamBinding, err error) {
	result = &v1alpha1.KmsCryptoKeyIamBinding{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kmscryptokeyiambindings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
