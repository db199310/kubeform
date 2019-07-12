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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// AwsDatapipelinePipelinesGetter has a method to return a AwsDatapipelinePipelineInterface.
// A group's client should implement this interface.
type AwsDatapipelinePipelinesGetter interface {
	AwsDatapipelinePipelines() AwsDatapipelinePipelineInterface
}

// AwsDatapipelinePipelineInterface has methods to work with AwsDatapipelinePipeline resources.
type AwsDatapipelinePipelineInterface interface {
	Create(*v1alpha1.AwsDatapipelinePipeline) (*v1alpha1.AwsDatapipelinePipeline, error)
	Update(*v1alpha1.AwsDatapipelinePipeline) (*v1alpha1.AwsDatapipelinePipeline, error)
	UpdateStatus(*v1alpha1.AwsDatapipelinePipeline) (*v1alpha1.AwsDatapipelinePipeline, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AwsDatapipelinePipeline, error)
	List(opts v1.ListOptions) (*v1alpha1.AwsDatapipelinePipelineList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDatapipelinePipeline, err error)
	AwsDatapipelinePipelineExpansion
}

// awsDatapipelinePipelines implements AwsDatapipelinePipelineInterface
type awsDatapipelinePipelines struct {
	client rest.Interface
}

// newAwsDatapipelinePipelines returns a AwsDatapipelinePipelines
func newAwsDatapipelinePipelines(c *AwsV1alpha1Client) *awsDatapipelinePipelines {
	return &awsDatapipelinePipelines{
		client: c.RESTClient(),
	}
}

// Get takes name of the awsDatapipelinePipeline, and returns the corresponding awsDatapipelinePipeline object, and an error if there is any.
func (c *awsDatapipelinePipelines) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsDatapipelinePipeline, err error) {
	result = &v1alpha1.AwsDatapipelinePipeline{}
	err = c.client.Get().
		Resource("awsdatapipelinepipelines").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsDatapipelinePipelines that match those selectors.
func (c *awsDatapipelinePipelines) List(opts v1.ListOptions) (result *v1alpha1.AwsDatapipelinePipelineList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AwsDatapipelinePipelineList{}
	err = c.client.Get().
		Resource("awsdatapipelinepipelines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsDatapipelinePipelines.
func (c *awsDatapipelinePipelines) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("awsdatapipelinepipelines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a awsDatapipelinePipeline and creates it.  Returns the server's representation of the awsDatapipelinePipeline, and an error, if there is any.
func (c *awsDatapipelinePipelines) Create(awsDatapipelinePipeline *v1alpha1.AwsDatapipelinePipeline) (result *v1alpha1.AwsDatapipelinePipeline, err error) {
	result = &v1alpha1.AwsDatapipelinePipeline{}
	err = c.client.Post().
		Resource("awsdatapipelinepipelines").
		Body(awsDatapipelinePipeline).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsDatapipelinePipeline and updates it. Returns the server's representation of the awsDatapipelinePipeline, and an error, if there is any.
func (c *awsDatapipelinePipelines) Update(awsDatapipelinePipeline *v1alpha1.AwsDatapipelinePipeline) (result *v1alpha1.AwsDatapipelinePipeline, err error) {
	result = &v1alpha1.AwsDatapipelinePipeline{}
	err = c.client.Put().
		Resource("awsdatapipelinepipelines").
		Name(awsDatapipelinePipeline.Name).
		Body(awsDatapipelinePipeline).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *awsDatapipelinePipelines) UpdateStatus(awsDatapipelinePipeline *v1alpha1.AwsDatapipelinePipeline) (result *v1alpha1.AwsDatapipelinePipeline, err error) {
	result = &v1alpha1.AwsDatapipelinePipeline{}
	err = c.client.Put().
		Resource("awsdatapipelinepipelines").
		Name(awsDatapipelinePipeline.Name).
		SubResource("status").
		Body(awsDatapipelinePipeline).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsDatapipelinePipeline and deletes it. Returns an error if one occurs.
func (c *awsDatapipelinePipelines) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("awsdatapipelinepipelines").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsDatapipelinePipelines) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("awsdatapipelinepipelines").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsDatapipelinePipeline.
func (c *awsDatapipelinePipelines) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDatapipelinePipeline, err error) {
	result = &v1alpha1.AwsDatapipelinePipeline{}
	err = c.client.Patch(pt).
		Resource("awsdatapipelinepipelines").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
