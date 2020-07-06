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

// MariadbDatabasesGetter has a method to return a MariadbDatabaseInterface.
// A group's client should implement this interface.
type MariadbDatabasesGetter interface {
	MariadbDatabases(namespace string) MariadbDatabaseInterface
}

// MariadbDatabaseInterface has methods to work with MariadbDatabase resources.
type MariadbDatabaseInterface interface {
	Create(*v1alpha1.MariadbDatabase) (*v1alpha1.MariadbDatabase, error)
	Update(*v1alpha1.MariadbDatabase) (*v1alpha1.MariadbDatabase, error)
	UpdateStatus(*v1alpha1.MariadbDatabase) (*v1alpha1.MariadbDatabase, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.MariadbDatabase, error)
	List(opts v1.ListOptions) (*v1alpha1.MariadbDatabaseList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MariadbDatabase, err error)
	MariadbDatabaseExpansion
}

// mariadbDatabases implements MariadbDatabaseInterface
type mariadbDatabases struct {
	client rest.Interface
	ns     string
}

// newMariadbDatabases returns a MariadbDatabases
func newMariadbDatabases(c *AzurermV1alpha1Client, namespace string) *mariadbDatabases {
	return &mariadbDatabases{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the mariadbDatabase, and returns the corresponding mariadbDatabase object, and an error if there is any.
func (c *mariadbDatabases) Get(name string, options v1.GetOptions) (result *v1alpha1.MariadbDatabase, err error) {
	result = &v1alpha1.MariadbDatabase{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mariadbdatabases").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MariadbDatabases that match those selectors.
func (c *mariadbDatabases) List(opts v1.ListOptions) (result *v1alpha1.MariadbDatabaseList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.MariadbDatabaseList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mariadbdatabases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested mariadbDatabases.
func (c *mariadbDatabases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mariadbdatabases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a mariadbDatabase and creates it.  Returns the server's representation of the mariadbDatabase, and an error, if there is any.
func (c *mariadbDatabases) Create(mariadbDatabase *v1alpha1.MariadbDatabase) (result *v1alpha1.MariadbDatabase, err error) {
	result = &v1alpha1.MariadbDatabase{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mariadbdatabases").
		Body(mariadbDatabase).
		Do().
		Into(result)
	return
}

// Update takes the representation of a mariadbDatabase and updates it. Returns the server's representation of the mariadbDatabase, and an error, if there is any.
func (c *mariadbDatabases) Update(mariadbDatabase *v1alpha1.MariadbDatabase) (result *v1alpha1.MariadbDatabase, err error) {
	result = &v1alpha1.MariadbDatabase{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mariadbdatabases").
		Name(mariadbDatabase.Name).
		Body(mariadbDatabase).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *mariadbDatabases) UpdateStatus(mariadbDatabase *v1alpha1.MariadbDatabase) (result *v1alpha1.MariadbDatabase, err error) {
	result = &v1alpha1.MariadbDatabase{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mariadbdatabases").
		Name(mariadbDatabase.Name).
		SubResource("status").
		Body(mariadbDatabase).
		Do().
		Into(result)
	return
}

// Delete takes name of the mariadbDatabase and deletes it. Returns an error if one occurs.
func (c *mariadbDatabases) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mariadbdatabases").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *mariadbDatabases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mariadbdatabases").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched mariadbDatabase.
func (c *mariadbDatabases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MariadbDatabase, err error) {
	result = &v1alpha1.MariadbDatabase{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mariadbdatabases").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
