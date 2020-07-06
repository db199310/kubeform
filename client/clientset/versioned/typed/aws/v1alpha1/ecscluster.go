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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// EcsClustersGetter has a method to return a EcsClusterInterface.
// A group's client should implement this interface.
type EcsClustersGetter interface {
	EcsClusters(namespace string) EcsClusterInterface
}

// EcsClusterInterface has methods to work with EcsCluster resources.
type EcsClusterInterface interface {
	Create(*v1alpha1.EcsCluster) (*v1alpha1.EcsCluster, error)
	Update(*v1alpha1.EcsCluster) (*v1alpha1.EcsCluster, error)
	UpdateStatus(*v1alpha1.EcsCluster) (*v1alpha1.EcsCluster, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.EcsCluster, error)
	List(opts v1.ListOptions) (*v1alpha1.EcsClusterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EcsCluster, err error)
	EcsClusterExpansion
}

// ecsClusters implements EcsClusterInterface
type ecsClusters struct {
	client rest.Interface
	ns     string
}

// newEcsClusters returns a EcsClusters
func newEcsClusters(c *AwsV1alpha1Client, namespace string) *ecsClusters {
	return &ecsClusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the ecsCluster, and returns the corresponding ecsCluster object, and an error if there is any.
func (c *ecsClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.EcsCluster, err error) {
	result = &v1alpha1.EcsCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ecsclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EcsClusters that match those selectors.
func (c *ecsClusters) List(opts v1.ListOptions) (result *v1alpha1.EcsClusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.EcsClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ecsclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ecsClusters.
func (c *ecsClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("ecsclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a ecsCluster and creates it.  Returns the server's representation of the ecsCluster, and an error, if there is any.
func (c *ecsClusters) Create(ecsCluster *v1alpha1.EcsCluster) (result *v1alpha1.EcsCluster, err error) {
	result = &v1alpha1.EcsCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("ecsclusters").
		Body(ecsCluster).
		Do().
		Into(result)
	return
}

// Update takes the representation of a ecsCluster and updates it. Returns the server's representation of the ecsCluster, and an error, if there is any.
func (c *ecsClusters) Update(ecsCluster *v1alpha1.EcsCluster) (result *v1alpha1.EcsCluster, err error) {
	result = &v1alpha1.EcsCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ecsclusters").
		Name(ecsCluster.Name).
		Body(ecsCluster).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *ecsClusters) UpdateStatus(ecsCluster *v1alpha1.EcsCluster) (result *v1alpha1.EcsCluster, err error) {
	result = &v1alpha1.EcsCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ecsclusters").
		Name(ecsCluster.Name).
		SubResource("status").
		Body(ecsCluster).
		Do().
		Into(result)
	return
}

// Delete takes name of the ecsCluster and deletes it. Returns an error if one occurs.
func (c *ecsClusters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ecsclusters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ecsClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ecsclusters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched ecsCluster.
func (c *ecsClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EcsCluster, err error) {
	result = &v1alpha1.EcsCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("ecsclusters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
