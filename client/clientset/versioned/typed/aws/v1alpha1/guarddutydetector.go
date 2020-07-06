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

// GuarddutyDetectorsGetter has a method to return a GuarddutyDetectorInterface.
// A group's client should implement this interface.
type GuarddutyDetectorsGetter interface {
	GuarddutyDetectors(namespace string) GuarddutyDetectorInterface
}

// GuarddutyDetectorInterface has methods to work with GuarddutyDetector resources.
type GuarddutyDetectorInterface interface {
	Create(*v1alpha1.GuarddutyDetector) (*v1alpha1.GuarddutyDetector, error)
	Update(*v1alpha1.GuarddutyDetector) (*v1alpha1.GuarddutyDetector, error)
	UpdateStatus(*v1alpha1.GuarddutyDetector) (*v1alpha1.GuarddutyDetector, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GuarddutyDetector, error)
	List(opts v1.ListOptions) (*v1alpha1.GuarddutyDetectorList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GuarddutyDetector, err error)
	GuarddutyDetectorExpansion
}

// guarddutyDetectors implements GuarddutyDetectorInterface
type guarddutyDetectors struct {
	client rest.Interface
	ns     string
}

// newGuarddutyDetectors returns a GuarddutyDetectors
func newGuarddutyDetectors(c *AwsV1alpha1Client, namespace string) *guarddutyDetectors {
	return &guarddutyDetectors{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the guarddutyDetector, and returns the corresponding guarddutyDetector object, and an error if there is any.
func (c *guarddutyDetectors) Get(name string, options v1.GetOptions) (result *v1alpha1.GuarddutyDetector, err error) {
	result = &v1alpha1.GuarddutyDetector{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("guarddutydetectors").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GuarddutyDetectors that match those selectors.
func (c *guarddutyDetectors) List(opts v1.ListOptions) (result *v1alpha1.GuarddutyDetectorList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GuarddutyDetectorList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("guarddutydetectors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested guarddutyDetectors.
func (c *guarddutyDetectors) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("guarddutydetectors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a guarddutyDetector and creates it.  Returns the server's representation of the guarddutyDetector, and an error, if there is any.
func (c *guarddutyDetectors) Create(guarddutyDetector *v1alpha1.GuarddutyDetector) (result *v1alpha1.GuarddutyDetector, err error) {
	result = &v1alpha1.GuarddutyDetector{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("guarddutydetectors").
		Body(guarddutyDetector).
		Do().
		Into(result)
	return
}

// Update takes the representation of a guarddutyDetector and updates it. Returns the server's representation of the guarddutyDetector, and an error, if there is any.
func (c *guarddutyDetectors) Update(guarddutyDetector *v1alpha1.GuarddutyDetector) (result *v1alpha1.GuarddutyDetector, err error) {
	result = &v1alpha1.GuarddutyDetector{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("guarddutydetectors").
		Name(guarddutyDetector.Name).
		Body(guarddutyDetector).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *guarddutyDetectors) UpdateStatus(guarddutyDetector *v1alpha1.GuarddutyDetector) (result *v1alpha1.GuarddutyDetector, err error) {
	result = &v1alpha1.GuarddutyDetector{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("guarddutydetectors").
		Name(guarddutyDetector.Name).
		SubResource("status").
		Body(guarddutyDetector).
		Do().
		Into(result)
	return
}

// Delete takes name of the guarddutyDetector and deletes it. Returns an error if one occurs.
func (c *guarddutyDetectors) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("guarddutydetectors").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *guarddutyDetectors) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("guarddutydetectors").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched guarddutyDetector.
func (c *guarddutyDetectors) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GuarddutyDetector, err error) {
	result = &v1alpha1.GuarddutyDetector{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("guarddutydetectors").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
