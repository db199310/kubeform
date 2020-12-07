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
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// MysqlDatabasesGetter has a method to return a MysqlDatabaseInterface.
// A group's client should implement this interface.
type MysqlDatabasesGetter interface {
	MysqlDatabases(namespace string) MysqlDatabaseInterface
}

// MysqlDatabaseInterface has methods to work with MysqlDatabase resources.
type MysqlDatabaseInterface interface {
	Create(ctx context.Context, mysqlDatabase *v1alpha1.MysqlDatabase, opts v1.CreateOptions) (*v1alpha1.MysqlDatabase, error)
	Update(ctx context.Context, mysqlDatabase *v1alpha1.MysqlDatabase, opts v1.UpdateOptions) (*v1alpha1.MysqlDatabase, error)
	UpdateStatus(ctx context.Context, mysqlDatabase *v1alpha1.MysqlDatabase, opts v1.UpdateOptions) (*v1alpha1.MysqlDatabase, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.MysqlDatabase, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.MysqlDatabaseList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MysqlDatabase, err error)
	MysqlDatabaseExpansion
}

// mysqlDatabases implements MysqlDatabaseInterface
type mysqlDatabases struct {
	client rest.Interface
	ns     string
}

// newMysqlDatabases returns a MysqlDatabases
func newMysqlDatabases(c *AzurermV1alpha1Client, namespace string) *mysqlDatabases {
	return &mysqlDatabases{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the mysqlDatabase, and returns the corresponding mysqlDatabase object, and an error if there is any.
func (c *mysqlDatabases) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MysqlDatabase, err error) {
	result = &v1alpha1.MysqlDatabase{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mysqldatabases").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MysqlDatabases that match those selectors.
func (c *mysqlDatabases) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MysqlDatabaseList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.MysqlDatabaseList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mysqldatabases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested mysqlDatabases.
func (c *mysqlDatabases) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mysqldatabases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a mysqlDatabase and creates it.  Returns the server's representation of the mysqlDatabase, and an error, if there is any.
func (c *mysqlDatabases) Create(ctx context.Context, mysqlDatabase *v1alpha1.MysqlDatabase, opts v1.CreateOptions) (result *v1alpha1.MysqlDatabase, err error) {
	result = &v1alpha1.MysqlDatabase{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mysqldatabases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(mysqlDatabase).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a mysqlDatabase and updates it. Returns the server's representation of the mysqlDatabase, and an error, if there is any.
func (c *mysqlDatabases) Update(ctx context.Context, mysqlDatabase *v1alpha1.MysqlDatabase, opts v1.UpdateOptions) (result *v1alpha1.MysqlDatabase, err error) {
	result = &v1alpha1.MysqlDatabase{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mysqldatabases").
		Name(mysqlDatabase.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(mysqlDatabase).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *mysqlDatabases) UpdateStatus(ctx context.Context, mysqlDatabase *v1alpha1.MysqlDatabase, opts v1.UpdateOptions) (result *v1alpha1.MysqlDatabase, err error) {
	result = &v1alpha1.MysqlDatabase{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mysqldatabases").
		Name(mysqlDatabase.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(mysqlDatabase).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the mysqlDatabase and deletes it. Returns an error if one occurs.
func (c *mysqlDatabases) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mysqldatabases").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *mysqlDatabases) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mysqldatabases").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched mysqlDatabase.
func (c *mysqlDatabases) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MysqlDatabase, err error) {
	result = &v1alpha1.MysqlDatabase{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mysqldatabases").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
