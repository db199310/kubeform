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

package v1alpha2

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha2 "kubeform.dev/kubeform/apis/modules/v1alpha2"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// ThomasStorageAccountsGetter has a method to return a ThomasStorageAccountInterface.
// A group's client should implement this interface.
type ThomasStorageAccountsGetter interface {
	ThomasStorageAccounts(namespace string) ThomasStorageAccountInterface
}

// ThomasStorageAccountInterface has methods to work with ThomasStorageAccount resources.
type ThomasStorageAccountInterface interface {
	Create(*v1alpha2.ThomasStorageAccount) (*v1alpha2.ThomasStorageAccount, error)
	Update(*v1alpha2.ThomasStorageAccount) (*v1alpha2.ThomasStorageAccount, error)
	UpdateStatus(*v1alpha2.ThomasStorageAccount) (*v1alpha2.ThomasStorageAccount, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha2.ThomasStorageAccount, error)
	List(opts v1.ListOptions) (*v1alpha2.ThomasStorageAccountList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.ThomasStorageAccount, err error)
	ThomasStorageAccountExpansion
}

// thomasStorageAccounts implements ThomasStorageAccountInterface
type thomasStorageAccounts struct {
	client rest.Interface
	ns     string
}

// newThomasStorageAccounts returns a ThomasStorageAccounts
func newThomasStorageAccounts(c *ModulesV1alpha2Client, namespace string) *thomasStorageAccounts {
	return &thomasStorageAccounts{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the thomasStorageAccount, and returns the corresponding thomasStorageAccount object, and an error if there is any.
func (c *thomasStorageAccounts) Get(name string, options v1.GetOptions) (result *v1alpha2.ThomasStorageAccount, err error) {
	result = &v1alpha2.ThomasStorageAccount{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("thomasstorageaccounts").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ThomasStorageAccounts that match those selectors.
func (c *thomasStorageAccounts) List(opts v1.ListOptions) (result *v1alpha2.ThomasStorageAccountList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.ThomasStorageAccountList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("thomasstorageaccounts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested thomasStorageAccounts.
func (c *thomasStorageAccounts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("thomasstorageaccounts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a thomasStorageAccount and creates it.  Returns the server's representation of the thomasStorageAccount, and an error, if there is any.
func (c *thomasStorageAccounts) Create(thomasStorageAccount *v1alpha2.ThomasStorageAccount) (result *v1alpha2.ThomasStorageAccount, err error) {
	result = &v1alpha2.ThomasStorageAccount{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("thomasstorageaccounts").
		Body(thomasStorageAccount).
		Do().
		Into(result)
	return
}

// Update takes the representation of a thomasStorageAccount and updates it. Returns the server's representation of the thomasStorageAccount, and an error, if there is any.
func (c *thomasStorageAccounts) Update(thomasStorageAccount *v1alpha2.ThomasStorageAccount) (result *v1alpha2.ThomasStorageAccount, err error) {
	result = &v1alpha2.ThomasStorageAccount{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("thomasstorageaccounts").
		Name(thomasStorageAccount.Name).
		Body(thomasStorageAccount).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *thomasStorageAccounts) UpdateStatus(thomasStorageAccount *v1alpha2.ThomasStorageAccount) (result *v1alpha2.ThomasStorageAccount, err error) {
	result = &v1alpha2.ThomasStorageAccount{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("thomasstorageaccounts").
		Name(thomasStorageAccount.Name).
		SubResource("status").
		Body(thomasStorageAccount).
		Do().
		Into(result)
	return
}

// Delete takes name of the thomasStorageAccount and deletes it. Returns an error if one occurs.
func (c *thomasStorageAccounts) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("thomasstorageaccounts").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *thomasStorageAccounts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("thomasstorageaccounts").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched thomasStorageAccount.
func (c *thomasStorageAccounts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.ThomasStorageAccount, err error) {
	result = &v1alpha2.ThomasStorageAccount{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("thomasstorageaccounts").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
