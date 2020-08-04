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

// VirtualMachineScaleSetExtensionsGetter has a method to return a VirtualMachineScaleSetExtensionInterface.
// A group's client should implement this interface.
type VirtualMachineScaleSetExtensionsGetter interface {
	VirtualMachineScaleSetExtensions(namespace string) VirtualMachineScaleSetExtensionInterface
}

// VirtualMachineScaleSetExtensionInterface has methods to work with VirtualMachineScaleSetExtension resources.
type VirtualMachineScaleSetExtensionInterface interface {
	Create(*v1alpha1.VirtualMachineScaleSetExtension) (*v1alpha1.VirtualMachineScaleSetExtension, error)
	Update(*v1alpha1.VirtualMachineScaleSetExtension) (*v1alpha1.VirtualMachineScaleSetExtension, error)
	UpdateStatus(*v1alpha1.VirtualMachineScaleSetExtension) (*v1alpha1.VirtualMachineScaleSetExtension, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.VirtualMachineScaleSetExtension, error)
	List(opts v1.ListOptions) (*v1alpha1.VirtualMachineScaleSetExtensionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VirtualMachineScaleSetExtension, err error)
	VirtualMachineScaleSetExtensionExpansion
}

// virtualMachineScaleSetExtensions implements VirtualMachineScaleSetExtensionInterface
type virtualMachineScaleSetExtensions struct {
	client rest.Interface
	ns     string
}

// newVirtualMachineScaleSetExtensions returns a VirtualMachineScaleSetExtensions
func newVirtualMachineScaleSetExtensions(c *AzurermV1alpha1Client, namespace string) *virtualMachineScaleSetExtensions {
	return &virtualMachineScaleSetExtensions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the virtualMachineScaleSetExtension, and returns the corresponding virtualMachineScaleSetExtension object, and an error if there is any.
func (c *virtualMachineScaleSetExtensions) Get(name string, options v1.GetOptions) (result *v1alpha1.VirtualMachineScaleSetExtension, err error) {
	result = &v1alpha1.VirtualMachineScaleSetExtension{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinescalesetextensions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VirtualMachineScaleSetExtensions that match those selectors.
func (c *virtualMachineScaleSetExtensions) List(opts v1.ListOptions) (result *v1alpha1.VirtualMachineScaleSetExtensionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VirtualMachineScaleSetExtensionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinescalesetextensions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested virtualMachineScaleSetExtensions.
func (c *virtualMachineScaleSetExtensions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinescalesetextensions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a virtualMachineScaleSetExtension and creates it.  Returns the server's representation of the virtualMachineScaleSetExtension, and an error, if there is any.
func (c *virtualMachineScaleSetExtensions) Create(virtualMachineScaleSetExtension *v1alpha1.VirtualMachineScaleSetExtension) (result *v1alpha1.VirtualMachineScaleSetExtension, err error) {
	result = &v1alpha1.VirtualMachineScaleSetExtension{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtualmachinescalesetextensions").
		Body(virtualMachineScaleSetExtension).
		Do().
		Into(result)
	return
}

// Update takes the representation of a virtualMachineScaleSetExtension and updates it. Returns the server's representation of the virtualMachineScaleSetExtension, and an error, if there is any.
func (c *virtualMachineScaleSetExtensions) Update(virtualMachineScaleSetExtension *v1alpha1.VirtualMachineScaleSetExtension) (result *v1alpha1.VirtualMachineScaleSetExtension, err error) {
	result = &v1alpha1.VirtualMachineScaleSetExtension{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachinescalesetextensions").
		Name(virtualMachineScaleSetExtension.Name).
		Body(virtualMachineScaleSetExtension).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *virtualMachineScaleSetExtensions) UpdateStatus(virtualMachineScaleSetExtension *v1alpha1.VirtualMachineScaleSetExtension) (result *v1alpha1.VirtualMachineScaleSetExtension, err error) {
	result = &v1alpha1.VirtualMachineScaleSetExtension{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachinescalesetextensions").
		Name(virtualMachineScaleSetExtension.Name).
		SubResource("status").
		Body(virtualMachineScaleSetExtension).
		Do().
		Into(result)
	return
}

// Delete takes name of the virtualMachineScaleSetExtension and deletes it. Returns an error if one occurs.
func (c *virtualMachineScaleSetExtensions) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachinescalesetextensions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *virtualMachineScaleSetExtensions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachinescalesetextensions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched virtualMachineScaleSetExtension.
func (c *virtualMachineScaleSetExtensions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VirtualMachineScaleSetExtension, err error) {
	result = &v1alpha1.VirtualMachineScaleSetExtension{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtualmachinescalesetextensions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
