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

// VpcPeeringConnectionsGetter has a method to return a VpcPeeringConnectionInterface.
// A group's client should implement this interface.
type VpcPeeringConnectionsGetter interface {
	VpcPeeringConnections(namespace string) VpcPeeringConnectionInterface
}

// VpcPeeringConnectionInterface has methods to work with VpcPeeringConnection resources.
type VpcPeeringConnectionInterface interface {
	Create(*v1alpha1.VpcPeeringConnection) (*v1alpha1.VpcPeeringConnection, error)
	Update(*v1alpha1.VpcPeeringConnection) (*v1alpha1.VpcPeeringConnection, error)
	UpdateStatus(*v1alpha1.VpcPeeringConnection) (*v1alpha1.VpcPeeringConnection, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.VpcPeeringConnection, error)
	List(opts v1.ListOptions) (*v1alpha1.VpcPeeringConnectionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VpcPeeringConnection, err error)
	VpcPeeringConnectionExpansion
}

// vpcPeeringConnections implements VpcPeeringConnectionInterface
type vpcPeeringConnections struct {
	client rest.Interface
	ns     string
}

// newVpcPeeringConnections returns a VpcPeeringConnections
func newVpcPeeringConnections(c *AwsV1alpha1Client, namespace string) *vpcPeeringConnections {
	return &vpcPeeringConnections{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the vpcPeeringConnection, and returns the corresponding vpcPeeringConnection object, and an error if there is any.
func (c *vpcPeeringConnections) Get(name string, options v1.GetOptions) (result *v1alpha1.VpcPeeringConnection, err error) {
	result = &v1alpha1.VpcPeeringConnection{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vpcpeeringconnections").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VpcPeeringConnections that match those selectors.
func (c *vpcPeeringConnections) List(opts v1.ListOptions) (result *v1alpha1.VpcPeeringConnectionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VpcPeeringConnectionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vpcpeeringconnections").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested vpcPeeringConnections.
func (c *vpcPeeringConnections) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("vpcpeeringconnections").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a vpcPeeringConnection and creates it.  Returns the server's representation of the vpcPeeringConnection, and an error, if there is any.
func (c *vpcPeeringConnections) Create(vpcPeeringConnection *v1alpha1.VpcPeeringConnection) (result *v1alpha1.VpcPeeringConnection, err error) {
	result = &v1alpha1.VpcPeeringConnection{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("vpcpeeringconnections").
		Body(vpcPeeringConnection).
		Do().
		Into(result)
	return
}

// Update takes the representation of a vpcPeeringConnection and updates it. Returns the server's representation of the vpcPeeringConnection, and an error, if there is any.
func (c *vpcPeeringConnections) Update(vpcPeeringConnection *v1alpha1.VpcPeeringConnection) (result *v1alpha1.VpcPeeringConnection, err error) {
	result = &v1alpha1.VpcPeeringConnection{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vpcpeeringconnections").
		Name(vpcPeeringConnection.Name).
		Body(vpcPeeringConnection).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *vpcPeeringConnections) UpdateStatus(vpcPeeringConnection *v1alpha1.VpcPeeringConnection) (result *v1alpha1.VpcPeeringConnection, err error) {
	result = &v1alpha1.VpcPeeringConnection{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vpcpeeringconnections").
		Name(vpcPeeringConnection.Name).
		SubResource("status").
		Body(vpcPeeringConnection).
		Do().
		Into(result)
	return
}

// Delete takes name of the vpcPeeringConnection and deletes it. Returns an error if one occurs.
func (c *vpcPeeringConnections) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vpcpeeringconnections").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *vpcPeeringConnections) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vpcpeeringconnections").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched vpcPeeringConnection.
func (c *vpcPeeringConnections) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VpcPeeringConnection, err error) {
	result = &v1alpha1.VpcPeeringConnection{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("vpcpeeringconnections").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
