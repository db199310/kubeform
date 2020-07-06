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

// SesActiveReceiptRuleSetsGetter has a method to return a SesActiveReceiptRuleSetInterface.
// A group's client should implement this interface.
type SesActiveReceiptRuleSetsGetter interface {
	SesActiveReceiptRuleSets(namespace string) SesActiveReceiptRuleSetInterface
}

// SesActiveReceiptRuleSetInterface has methods to work with SesActiveReceiptRuleSet resources.
type SesActiveReceiptRuleSetInterface interface {
	Create(*v1alpha1.SesActiveReceiptRuleSet) (*v1alpha1.SesActiveReceiptRuleSet, error)
	Update(*v1alpha1.SesActiveReceiptRuleSet) (*v1alpha1.SesActiveReceiptRuleSet, error)
	UpdateStatus(*v1alpha1.SesActiveReceiptRuleSet) (*v1alpha1.SesActiveReceiptRuleSet, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SesActiveReceiptRuleSet, error)
	List(opts v1.ListOptions) (*v1alpha1.SesActiveReceiptRuleSetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SesActiveReceiptRuleSet, err error)
	SesActiveReceiptRuleSetExpansion
}

// sesActiveReceiptRuleSets implements SesActiveReceiptRuleSetInterface
type sesActiveReceiptRuleSets struct {
	client rest.Interface
	ns     string
}

// newSesActiveReceiptRuleSets returns a SesActiveReceiptRuleSets
func newSesActiveReceiptRuleSets(c *AwsV1alpha1Client, namespace string) *sesActiveReceiptRuleSets {
	return &sesActiveReceiptRuleSets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sesActiveReceiptRuleSet, and returns the corresponding sesActiveReceiptRuleSet object, and an error if there is any.
func (c *sesActiveReceiptRuleSets) Get(name string, options v1.GetOptions) (result *v1alpha1.SesActiveReceiptRuleSet, err error) {
	result = &v1alpha1.SesActiveReceiptRuleSet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sesactivereceiptrulesets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SesActiveReceiptRuleSets that match those selectors.
func (c *sesActiveReceiptRuleSets) List(opts v1.ListOptions) (result *v1alpha1.SesActiveReceiptRuleSetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SesActiveReceiptRuleSetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sesactivereceiptrulesets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sesActiveReceiptRuleSets.
func (c *sesActiveReceiptRuleSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sesactivereceiptrulesets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a sesActiveReceiptRuleSet and creates it.  Returns the server's representation of the sesActiveReceiptRuleSet, and an error, if there is any.
func (c *sesActiveReceiptRuleSets) Create(sesActiveReceiptRuleSet *v1alpha1.SesActiveReceiptRuleSet) (result *v1alpha1.SesActiveReceiptRuleSet, err error) {
	result = &v1alpha1.SesActiveReceiptRuleSet{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sesactivereceiptrulesets").
		Body(sesActiveReceiptRuleSet).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sesActiveReceiptRuleSet and updates it. Returns the server's representation of the sesActiveReceiptRuleSet, and an error, if there is any.
func (c *sesActiveReceiptRuleSets) Update(sesActiveReceiptRuleSet *v1alpha1.SesActiveReceiptRuleSet) (result *v1alpha1.SesActiveReceiptRuleSet, err error) {
	result = &v1alpha1.SesActiveReceiptRuleSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sesactivereceiptrulesets").
		Name(sesActiveReceiptRuleSet.Name).
		Body(sesActiveReceiptRuleSet).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *sesActiveReceiptRuleSets) UpdateStatus(sesActiveReceiptRuleSet *v1alpha1.SesActiveReceiptRuleSet) (result *v1alpha1.SesActiveReceiptRuleSet, err error) {
	result = &v1alpha1.SesActiveReceiptRuleSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sesactivereceiptrulesets").
		Name(sesActiveReceiptRuleSet.Name).
		SubResource("status").
		Body(sesActiveReceiptRuleSet).
		Do().
		Into(result)
	return
}

// Delete takes name of the sesActiveReceiptRuleSet and deletes it. Returns an error if one occurs.
func (c *sesActiveReceiptRuleSets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sesactivereceiptrulesets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sesActiveReceiptRuleSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sesactivereceiptrulesets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sesActiveReceiptRuleSet.
func (c *sesActiveReceiptRuleSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SesActiveReceiptRuleSet, err error) {
	result = &v1alpha1.SesActiveReceiptRuleSet{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sesactivereceiptrulesets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
