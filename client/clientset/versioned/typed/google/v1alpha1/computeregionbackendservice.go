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

// ComputeRegionBackendServicesGetter has a method to return a ComputeRegionBackendServiceInterface.
// A group's client should implement this interface.
type ComputeRegionBackendServicesGetter interface {
	ComputeRegionBackendServices(namespace string) ComputeRegionBackendServiceInterface
}

// ComputeRegionBackendServiceInterface has methods to work with ComputeRegionBackendService resources.
type ComputeRegionBackendServiceInterface interface {
	Create(*v1alpha1.ComputeRegionBackendService) (*v1alpha1.ComputeRegionBackendService, error)
	Update(*v1alpha1.ComputeRegionBackendService) (*v1alpha1.ComputeRegionBackendService, error)
	UpdateStatus(*v1alpha1.ComputeRegionBackendService) (*v1alpha1.ComputeRegionBackendService, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ComputeRegionBackendService, error)
	List(opts v1.ListOptions) (*v1alpha1.ComputeRegionBackendServiceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeRegionBackendService, err error)
	ComputeRegionBackendServiceExpansion
}

// computeRegionBackendServices implements ComputeRegionBackendServiceInterface
type computeRegionBackendServices struct {
	client rest.Interface
	ns     string
}

// newComputeRegionBackendServices returns a ComputeRegionBackendServices
func newComputeRegionBackendServices(c *GoogleV1alpha1Client, namespace string) *computeRegionBackendServices {
	return &computeRegionBackendServices{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the computeRegionBackendService, and returns the corresponding computeRegionBackendService object, and an error if there is any.
func (c *computeRegionBackendServices) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeRegionBackendService, err error) {
	result = &v1alpha1.ComputeRegionBackendService{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computeregionbackendservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComputeRegionBackendServices that match those selectors.
func (c *computeRegionBackendServices) List(opts v1.ListOptions) (result *v1alpha1.ComputeRegionBackendServiceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ComputeRegionBackendServiceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computeregionbackendservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested computeRegionBackendServices.
func (c *computeRegionBackendServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("computeregionbackendservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a computeRegionBackendService and creates it.  Returns the server's representation of the computeRegionBackendService, and an error, if there is any.
func (c *computeRegionBackendServices) Create(computeRegionBackendService *v1alpha1.ComputeRegionBackendService) (result *v1alpha1.ComputeRegionBackendService, err error) {
	result = &v1alpha1.ComputeRegionBackendService{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("computeregionbackendservices").
		Body(computeRegionBackendService).
		Do().
		Into(result)
	return
}

// Update takes the representation of a computeRegionBackendService and updates it. Returns the server's representation of the computeRegionBackendService, and an error, if there is any.
func (c *computeRegionBackendServices) Update(computeRegionBackendService *v1alpha1.ComputeRegionBackendService) (result *v1alpha1.ComputeRegionBackendService, err error) {
	result = &v1alpha1.ComputeRegionBackendService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computeregionbackendservices").
		Name(computeRegionBackendService.Name).
		Body(computeRegionBackendService).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *computeRegionBackendServices) UpdateStatus(computeRegionBackendService *v1alpha1.ComputeRegionBackendService) (result *v1alpha1.ComputeRegionBackendService, err error) {
	result = &v1alpha1.ComputeRegionBackendService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computeregionbackendservices").
		Name(computeRegionBackendService.Name).
		SubResource("status").
		Body(computeRegionBackendService).
		Do().
		Into(result)
	return
}

// Delete takes name of the computeRegionBackendService and deletes it. Returns an error if one occurs.
func (c *computeRegionBackendServices) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computeregionbackendservices").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *computeRegionBackendServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computeregionbackendservices").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched computeRegionBackendService.
func (c *computeRegionBackendServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeRegionBackendService, err error) {
	result = &v1alpha1.ComputeRegionBackendService{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("computeregionbackendservices").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
