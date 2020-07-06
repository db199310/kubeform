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

// EbsVolumesGetter has a method to return a EbsVolumeInterface.
// A group's client should implement this interface.
type EbsVolumesGetter interface {
	EbsVolumes(namespace string) EbsVolumeInterface
}

// EbsVolumeInterface has methods to work with EbsVolume resources.
type EbsVolumeInterface interface {
	Create(*v1alpha1.EbsVolume) (*v1alpha1.EbsVolume, error)
	Update(*v1alpha1.EbsVolume) (*v1alpha1.EbsVolume, error)
	UpdateStatus(*v1alpha1.EbsVolume) (*v1alpha1.EbsVolume, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.EbsVolume, error)
	List(opts v1.ListOptions) (*v1alpha1.EbsVolumeList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EbsVolume, err error)
	EbsVolumeExpansion
}

// ebsVolumes implements EbsVolumeInterface
type ebsVolumes struct {
	client rest.Interface
	ns     string
}

// newEbsVolumes returns a EbsVolumes
func newEbsVolumes(c *AwsV1alpha1Client, namespace string) *ebsVolumes {
	return &ebsVolumes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the ebsVolume, and returns the corresponding ebsVolume object, and an error if there is any.
func (c *ebsVolumes) Get(name string, options v1.GetOptions) (result *v1alpha1.EbsVolume, err error) {
	result = &v1alpha1.EbsVolume{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ebsvolumes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EbsVolumes that match those selectors.
func (c *ebsVolumes) List(opts v1.ListOptions) (result *v1alpha1.EbsVolumeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.EbsVolumeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ebsvolumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ebsVolumes.
func (c *ebsVolumes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("ebsvolumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a ebsVolume and creates it.  Returns the server's representation of the ebsVolume, and an error, if there is any.
func (c *ebsVolumes) Create(ebsVolume *v1alpha1.EbsVolume) (result *v1alpha1.EbsVolume, err error) {
	result = &v1alpha1.EbsVolume{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("ebsvolumes").
		Body(ebsVolume).
		Do().
		Into(result)
	return
}

// Update takes the representation of a ebsVolume and updates it. Returns the server's representation of the ebsVolume, and an error, if there is any.
func (c *ebsVolumes) Update(ebsVolume *v1alpha1.EbsVolume) (result *v1alpha1.EbsVolume, err error) {
	result = &v1alpha1.EbsVolume{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ebsvolumes").
		Name(ebsVolume.Name).
		Body(ebsVolume).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *ebsVolumes) UpdateStatus(ebsVolume *v1alpha1.EbsVolume) (result *v1alpha1.EbsVolume, err error) {
	result = &v1alpha1.EbsVolume{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ebsvolumes").
		Name(ebsVolume.Name).
		SubResource("status").
		Body(ebsVolume).
		Do().
		Into(result)
	return
}

// Delete takes name of the ebsVolume and deletes it. Returns an error if one occurs.
func (c *ebsVolumes) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ebsvolumes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ebsVolumes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ebsvolumes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched ebsVolume.
func (c *ebsVolumes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EbsVolume, err error) {
	result = &v1alpha1.EbsVolume{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("ebsvolumes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
