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

// CdnProfilesGetter has a method to return a CdnProfileInterface.
// A group's client should implement this interface.
type CdnProfilesGetter interface {
	CdnProfiles(namespace string) CdnProfileInterface
}

// CdnProfileInterface has methods to work with CdnProfile resources.
type CdnProfileInterface interface {
	Create(ctx context.Context, cdnProfile *v1alpha1.CdnProfile, opts v1.CreateOptions) (*v1alpha1.CdnProfile, error)
	Update(ctx context.Context, cdnProfile *v1alpha1.CdnProfile, opts v1.UpdateOptions) (*v1alpha1.CdnProfile, error)
	UpdateStatus(ctx context.Context, cdnProfile *v1alpha1.CdnProfile, opts v1.UpdateOptions) (*v1alpha1.CdnProfile, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.CdnProfile, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.CdnProfileList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CdnProfile, err error)
	CdnProfileExpansion
}

// cdnProfiles implements CdnProfileInterface
type cdnProfiles struct {
	client rest.Interface
	ns     string
}

// newCdnProfiles returns a CdnProfiles
func newCdnProfiles(c *AzurermV1alpha1Client, namespace string) *cdnProfiles {
	return &cdnProfiles{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cdnProfile, and returns the corresponding cdnProfile object, and an error if there is any.
func (c *cdnProfiles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CdnProfile, err error) {
	result = &v1alpha1.CdnProfile{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cdnprofiles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CdnProfiles that match those selectors.
func (c *cdnProfiles) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CdnProfileList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CdnProfileList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cdnprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cdnProfiles.
func (c *cdnProfiles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cdnprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a cdnProfile and creates it.  Returns the server's representation of the cdnProfile, and an error, if there is any.
func (c *cdnProfiles) Create(ctx context.Context, cdnProfile *v1alpha1.CdnProfile, opts v1.CreateOptions) (result *v1alpha1.CdnProfile, err error) {
	result = &v1alpha1.CdnProfile{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cdnprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cdnProfile).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a cdnProfile and updates it. Returns the server's representation of the cdnProfile, and an error, if there is any.
func (c *cdnProfiles) Update(ctx context.Context, cdnProfile *v1alpha1.CdnProfile, opts v1.UpdateOptions) (result *v1alpha1.CdnProfile, err error) {
	result = &v1alpha1.CdnProfile{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cdnprofiles").
		Name(cdnProfile.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cdnProfile).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *cdnProfiles) UpdateStatus(ctx context.Context, cdnProfile *v1alpha1.CdnProfile, opts v1.UpdateOptions) (result *v1alpha1.CdnProfile, err error) {
	result = &v1alpha1.CdnProfile{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cdnprofiles").
		Name(cdnProfile.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cdnProfile).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the cdnProfile and deletes it. Returns an error if one occurs.
func (c *cdnProfiles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cdnprofiles").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cdnProfiles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cdnprofiles").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched cdnProfile.
func (c *cdnProfiles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CdnProfile, err error) {
	result = &v1alpha1.CdnProfile{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cdnprofiles").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
