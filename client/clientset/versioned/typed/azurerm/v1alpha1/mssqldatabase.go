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

// MssqlDatabasesGetter has a method to return a MssqlDatabaseInterface.
// A group's client should implement this interface.
type MssqlDatabasesGetter interface {
	MssqlDatabases(namespace string) MssqlDatabaseInterface
}

// MssqlDatabaseInterface has methods to work with MssqlDatabase resources.
type MssqlDatabaseInterface interface {
	Create(*v1alpha1.MssqlDatabase) (*v1alpha1.MssqlDatabase, error)
	Update(*v1alpha1.MssqlDatabase) (*v1alpha1.MssqlDatabase, error)
	UpdateStatus(*v1alpha1.MssqlDatabase) (*v1alpha1.MssqlDatabase, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.MssqlDatabase, error)
	List(opts v1.ListOptions) (*v1alpha1.MssqlDatabaseList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MssqlDatabase, err error)
	MssqlDatabaseExpansion
}

// mssqlDatabases implements MssqlDatabaseInterface
type mssqlDatabases struct {
	client rest.Interface
	ns     string
}

// newMssqlDatabases returns a MssqlDatabases
func newMssqlDatabases(c *AzurermV1alpha1Client, namespace string) *mssqlDatabases {
	return &mssqlDatabases{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the mssqlDatabase, and returns the corresponding mssqlDatabase object, and an error if there is any.
func (c *mssqlDatabases) Get(name string, options v1.GetOptions) (result *v1alpha1.MssqlDatabase, err error) {
	result = &v1alpha1.MssqlDatabase{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mssqldatabases").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MssqlDatabases that match those selectors.
func (c *mssqlDatabases) List(opts v1.ListOptions) (result *v1alpha1.MssqlDatabaseList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.MssqlDatabaseList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mssqldatabases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested mssqlDatabases.
func (c *mssqlDatabases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mssqldatabases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a mssqlDatabase and creates it.  Returns the server's representation of the mssqlDatabase, and an error, if there is any.
func (c *mssqlDatabases) Create(mssqlDatabase *v1alpha1.MssqlDatabase) (result *v1alpha1.MssqlDatabase, err error) {
	result = &v1alpha1.MssqlDatabase{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mssqldatabases").
		Body(mssqlDatabase).
		Do().
		Into(result)
	return
}

// Update takes the representation of a mssqlDatabase and updates it. Returns the server's representation of the mssqlDatabase, and an error, if there is any.
func (c *mssqlDatabases) Update(mssqlDatabase *v1alpha1.MssqlDatabase) (result *v1alpha1.MssqlDatabase, err error) {
	result = &v1alpha1.MssqlDatabase{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mssqldatabases").
		Name(mssqlDatabase.Name).
		Body(mssqlDatabase).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *mssqlDatabases) UpdateStatus(mssqlDatabase *v1alpha1.MssqlDatabase) (result *v1alpha1.MssqlDatabase, err error) {
	result = &v1alpha1.MssqlDatabase{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mssqldatabases").
		Name(mssqlDatabase.Name).
		SubResource("status").
		Body(mssqlDatabase).
		Do().
		Into(result)
	return
}

// Delete takes name of the mssqlDatabase and deletes it. Returns an error if one occurs.
func (c *mssqlDatabases) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mssqldatabases").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *mssqlDatabases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mssqldatabases").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched mssqlDatabase.
func (c *mssqlDatabases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MssqlDatabase, err error) {
	result = &v1alpha1.MssqlDatabase{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mssqldatabases").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}