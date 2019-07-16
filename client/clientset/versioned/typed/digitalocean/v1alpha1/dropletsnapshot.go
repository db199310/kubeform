/*
Copyright 2019 The Kubeform Authors.

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
	v1alpha1 "kubeform.dev/kubeform/apis/digitalocean/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// DropletSnapshotsGetter has a method to return a DropletSnapshotInterface.
// A group's client should implement this interface.
type DropletSnapshotsGetter interface {
	DropletSnapshots() DropletSnapshotInterface
}

// DropletSnapshotInterface has methods to work with DropletSnapshot resources.
type DropletSnapshotInterface interface {
	Create(*v1alpha1.DropletSnapshot) (*v1alpha1.DropletSnapshot, error)
	Update(*v1alpha1.DropletSnapshot) (*v1alpha1.DropletSnapshot, error)
	UpdateStatus(*v1alpha1.DropletSnapshot) (*v1alpha1.DropletSnapshot, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DropletSnapshot, error)
	List(opts v1.ListOptions) (*v1alpha1.DropletSnapshotList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DropletSnapshot, err error)
	DropletSnapshotExpansion
}

// dropletSnapshots implements DropletSnapshotInterface
type dropletSnapshots struct {
	client rest.Interface
}

// newDropletSnapshots returns a DropletSnapshots
func newDropletSnapshots(c *DigitaloceanV1alpha1Client) *dropletSnapshots {
	return &dropletSnapshots{
		client: c.RESTClient(),
	}
}

// Get takes name of the dropletSnapshot, and returns the corresponding dropletSnapshot object, and an error if there is any.
func (c *dropletSnapshots) Get(name string, options v1.GetOptions) (result *v1alpha1.DropletSnapshot, err error) {
	result = &v1alpha1.DropletSnapshot{}
	err = c.client.Get().
		Resource("dropletsnapshots").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DropletSnapshots that match those selectors.
func (c *dropletSnapshots) List(opts v1.ListOptions) (result *v1alpha1.DropletSnapshotList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DropletSnapshotList{}
	err = c.client.Get().
		Resource("dropletsnapshots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dropletSnapshots.
func (c *dropletSnapshots) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("dropletsnapshots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a dropletSnapshot and creates it.  Returns the server's representation of the dropletSnapshot, and an error, if there is any.
func (c *dropletSnapshots) Create(dropletSnapshot *v1alpha1.DropletSnapshot) (result *v1alpha1.DropletSnapshot, err error) {
	result = &v1alpha1.DropletSnapshot{}
	err = c.client.Post().
		Resource("dropletsnapshots").
		Body(dropletSnapshot).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dropletSnapshot and updates it. Returns the server's representation of the dropletSnapshot, and an error, if there is any.
func (c *dropletSnapshots) Update(dropletSnapshot *v1alpha1.DropletSnapshot) (result *v1alpha1.DropletSnapshot, err error) {
	result = &v1alpha1.DropletSnapshot{}
	err = c.client.Put().
		Resource("dropletsnapshots").
		Name(dropletSnapshot.Name).
		Body(dropletSnapshot).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *dropletSnapshots) UpdateStatus(dropletSnapshot *v1alpha1.DropletSnapshot) (result *v1alpha1.DropletSnapshot, err error) {
	result = &v1alpha1.DropletSnapshot{}
	err = c.client.Put().
		Resource("dropletsnapshots").
		Name(dropletSnapshot.Name).
		SubResource("status").
		Body(dropletSnapshot).
		Do().
		Into(result)
	return
}

// Delete takes name of the dropletSnapshot and deletes it. Returns an error if one occurs.
func (c *dropletSnapshots) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("dropletsnapshots").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dropletSnapshots) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("dropletsnapshots").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dropletSnapshot.
func (c *dropletSnapshots) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DropletSnapshot, err error) {
	result = &v1alpha1.DropletSnapshot{}
	err = c.client.Patch(pt).
		Resource("dropletsnapshots").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}