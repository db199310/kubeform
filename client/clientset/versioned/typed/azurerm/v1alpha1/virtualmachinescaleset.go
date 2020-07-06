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

// VirtualMachineScaleSetsGetter has a method to return a VirtualMachineScaleSetInterface.
// A group's client should implement this interface.
type VirtualMachineScaleSetsGetter interface {
	VirtualMachineScaleSets(namespace string) VirtualMachineScaleSetInterface
}

// VirtualMachineScaleSetInterface has methods to work with VirtualMachineScaleSet resources.
type VirtualMachineScaleSetInterface interface {
	Create(*v1alpha1.VirtualMachineScaleSet) (*v1alpha1.VirtualMachineScaleSet, error)
	Update(*v1alpha1.VirtualMachineScaleSet) (*v1alpha1.VirtualMachineScaleSet, error)
	UpdateStatus(*v1alpha1.VirtualMachineScaleSet) (*v1alpha1.VirtualMachineScaleSet, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.VirtualMachineScaleSet, error)
	List(opts v1.ListOptions) (*v1alpha1.VirtualMachineScaleSetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VirtualMachineScaleSet, err error)
	VirtualMachineScaleSetExpansion
}

// virtualMachineScaleSets implements VirtualMachineScaleSetInterface
type virtualMachineScaleSets struct {
	client rest.Interface
	ns     string
}

// newVirtualMachineScaleSets returns a VirtualMachineScaleSets
func newVirtualMachineScaleSets(c *AzurermV1alpha1Client, namespace string) *virtualMachineScaleSets {
	return &virtualMachineScaleSets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the virtualMachineScaleSet, and returns the corresponding virtualMachineScaleSet object, and an error if there is any.
func (c *virtualMachineScaleSets) Get(name string, options v1.GetOptions) (result *v1alpha1.VirtualMachineScaleSet, err error) {
	result = &v1alpha1.VirtualMachineScaleSet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinescalesets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VirtualMachineScaleSets that match those selectors.
func (c *virtualMachineScaleSets) List(opts v1.ListOptions) (result *v1alpha1.VirtualMachineScaleSetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VirtualMachineScaleSetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinescalesets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested virtualMachineScaleSets.
func (c *virtualMachineScaleSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinescalesets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a virtualMachineScaleSet and creates it.  Returns the server's representation of the virtualMachineScaleSet, and an error, if there is any.
func (c *virtualMachineScaleSets) Create(virtualMachineScaleSet *v1alpha1.VirtualMachineScaleSet) (result *v1alpha1.VirtualMachineScaleSet, err error) {
	result = &v1alpha1.VirtualMachineScaleSet{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtualmachinescalesets").
		Body(virtualMachineScaleSet).
		Do().
		Into(result)
	return
}

// Update takes the representation of a virtualMachineScaleSet and updates it. Returns the server's representation of the virtualMachineScaleSet, and an error, if there is any.
func (c *virtualMachineScaleSets) Update(virtualMachineScaleSet *v1alpha1.VirtualMachineScaleSet) (result *v1alpha1.VirtualMachineScaleSet, err error) {
	result = &v1alpha1.VirtualMachineScaleSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachinescalesets").
		Name(virtualMachineScaleSet.Name).
		Body(virtualMachineScaleSet).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *virtualMachineScaleSets) UpdateStatus(virtualMachineScaleSet *v1alpha1.VirtualMachineScaleSet) (result *v1alpha1.VirtualMachineScaleSet, err error) {
	result = &v1alpha1.VirtualMachineScaleSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachinescalesets").
		Name(virtualMachineScaleSet.Name).
		SubResource("status").
		Body(virtualMachineScaleSet).
		Do().
		Into(result)
	return
}

// Delete takes name of the virtualMachineScaleSet and deletes it. Returns an error if one occurs.
func (c *virtualMachineScaleSets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachinescalesets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *virtualMachineScaleSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachinescalesets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched virtualMachineScaleSet.
func (c *virtualMachineScaleSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VirtualMachineScaleSet, err error) {
	result = &v1alpha1.VirtualMachineScaleSet{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtualmachinescalesets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
