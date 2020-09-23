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

// MachineLearningWorkspacesGetter has a method to return a MachineLearningWorkspaceInterface.
// A group's client should implement this interface.
type MachineLearningWorkspacesGetter interface {
	MachineLearningWorkspaces(namespace string) MachineLearningWorkspaceInterface
}

// MachineLearningWorkspaceInterface has methods to work with MachineLearningWorkspace resources.
type MachineLearningWorkspaceInterface interface {
	Create(*v1alpha1.MachineLearningWorkspace) (*v1alpha1.MachineLearningWorkspace, error)
	Update(*v1alpha1.MachineLearningWorkspace) (*v1alpha1.MachineLearningWorkspace, error)
	UpdateStatus(*v1alpha1.MachineLearningWorkspace) (*v1alpha1.MachineLearningWorkspace, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.MachineLearningWorkspace, error)
	List(opts v1.ListOptions) (*v1alpha1.MachineLearningWorkspaceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MachineLearningWorkspace, err error)
	MachineLearningWorkspaceExpansion
}

// machineLearningWorkspaces implements MachineLearningWorkspaceInterface
type machineLearningWorkspaces struct {
	client rest.Interface
	ns     string
}

// newMachineLearningWorkspaces returns a MachineLearningWorkspaces
func newMachineLearningWorkspaces(c *AzurermV1alpha1Client, namespace string) *machineLearningWorkspaces {
	return &machineLearningWorkspaces{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the machineLearningWorkspace, and returns the corresponding machineLearningWorkspace object, and an error if there is any.
func (c *machineLearningWorkspaces) Get(name string, options v1.GetOptions) (result *v1alpha1.MachineLearningWorkspace, err error) {
	result = &v1alpha1.MachineLearningWorkspace{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("machinelearningworkspaces").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MachineLearningWorkspaces that match those selectors.
func (c *machineLearningWorkspaces) List(opts v1.ListOptions) (result *v1alpha1.MachineLearningWorkspaceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.MachineLearningWorkspaceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("machinelearningworkspaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested machineLearningWorkspaces.
func (c *machineLearningWorkspaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("machinelearningworkspaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a machineLearningWorkspace and creates it.  Returns the server's representation of the machineLearningWorkspace, and an error, if there is any.
func (c *machineLearningWorkspaces) Create(machineLearningWorkspace *v1alpha1.MachineLearningWorkspace) (result *v1alpha1.MachineLearningWorkspace, err error) {
	result = &v1alpha1.MachineLearningWorkspace{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("machinelearningworkspaces").
		Body(machineLearningWorkspace).
		Do().
		Into(result)
	return
}

// Update takes the representation of a machineLearningWorkspace and updates it. Returns the server's representation of the machineLearningWorkspace, and an error, if there is any.
func (c *machineLearningWorkspaces) Update(machineLearningWorkspace *v1alpha1.MachineLearningWorkspace) (result *v1alpha1.MachineLearningWorkspace, err error) {
	result = &v1alpha1.MachineLearningWorkspace{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("machinelearningworkspaces").
		Name(machineLearningWorkspace.Name).
		Body(machineLearningWorkspace).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *machineLearningWorkspaces) UpdateStatus(machineLearningWorkspace *v1alpha1.MachineLearningWorkspace) (result *v1alpha1.MachineLearningWorkspace, err error) {
	result = &v1alpha1.MachineLearningWorkspace{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("machinelearningworkspaces").
		Name(machineLearningWorkspace.Name).
		SubResource("status").
		Body(machineLearningWorkspace).
		Do().
		Into(result)
	return
}

// Delete takes name of the machineLearningWorkspace and deletes it. Returns an error if one occurs.
func (c *machineLearningWorkspaces) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("machinelearningworkspaces").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *machineLearningWorkspaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("machinelearningworkspaces").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched machineLearningWorkspace.
func (c *machineLearningWorkspaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MachineLearningWorkspace, err error) {
	result = &v1alpha1.MachineLearningWorkspace{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("machinelearningworkspaces").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}