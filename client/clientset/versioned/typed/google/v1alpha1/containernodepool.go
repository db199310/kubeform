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

// ContainerNodePoolsGetter has a method to return a ContainerNodePoolInterface.
// A group's client should implement this interface.
type ContainerNodePoolsGetter interface {
	ContainerNodePools(namespace string) ContainerNodePoolInterface
}

// ContainerNodePoolInterface has methods to work with ContainerNodePool resources.
type ContainerNodePoolInterface interface {
	Create(*v1alpha1.ContainerNodePool) (*v1alpha1.ContainerNodePool, error)
	Update(*v1alpha1.ContainerNodePool) (*v1alpha1.ContainerNodePool, error)
	UpdateStatus(*v1alpha1.ContainerNodePool) (*v1alpha1.ContainerNodePool, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ContainerNodePool, error)
	List(opts v1.ListOptions) (*v1alpha1.ContainerNodePoolList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ContainerNodePool, err error)
	ContainerNodePoolExpansion
}

// containerNodePools implements ContainerNodePoolInterface
type containerNodePools struct {
	client rest.Interface
	ns     string
}

// newContainerNodePools returns a ContainerNodePools
func newContainerNodePools(c *GoogleV1alpha1Client, namespace string) *containerNodePools {
	return &containerNodePools{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the containerNodePool, and returns the corresponding containerNodePool object, and an error if there is any.
func (c *containerNodePools) Get(name string, options v1.GetOptions) (result *v1alpha1.ContainerNodePool, err error) {
	result = &v1alpha1.ContainerNodePool{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("containernodepools").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ContainerNodePools that match those selectors.
func (c *containerNodePools) List(opts v1.ListOptions) (result *v1alpha1.ContainerNodePoolList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ContainerNodePoolList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("containernodepools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested containerNodePools.
func (c *containerNodePools) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("containernodepools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a containerNodePool and creates it.  Returns the server's representation of the containerNodePool, and an error, if there is any.
func (c *containerNodePools) Create(containerNodePool *v1alpha1.ContainerNodePool) (result *v1alpha1.ContainerNodePool, err error) {
	result = &v1alpha1.ContainerNodePool{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("containernodepools").
		Body(containerNodePool).
		Do().
		Into(result)
	return
}

// Update takes the representation of a containerNodePool and updates it. Returns the server's representation of the containerNodePool, and an error, if there is any.
func (c *containerNodePools) Update(containerNodePool *v1alpha1.ContainerNodePool) (result *v1alpha1.ContainerNodePool, err error) {
	result = &v1alpha1.ContainerNodePool{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("containernodepools").
		Name(containerNodePool.Name).
		Body(containerNodePool).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *containerNodePools) UpdateStatus(containerNodePool *v1alpha1.ContainerNodePool) (result *v1alpha1.ContainerNodePool, err error) {
	result = &v1alpha1.ContainerNodePool{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("containernodepools").
		Name(containerNodePool.Name).
		SubResource("status").
		Body(containerNodePool).
		Do().
		Into(result)
	return
}

// Delete takes name of the containerNodePool and deletes it. Returns an error if one occurs.
func (c *containerNodePools) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("containernodepools").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *containerNodePools) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("containernodepools").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched containerNodePool.
func (c *containerNodePools) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ContainerNodePool, err error) {
	result = &v1alpha1.ContainerNodePool{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("containernodepools").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
