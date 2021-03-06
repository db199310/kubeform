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

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakeApiManagementCertificates implements ApiManagementCertificateInterface
type FakeApiManagementCertificates struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var apimanagementcertificatesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "apimanagementcertificates"}

var apimanagementcertificatesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "ApiManagementCertificate"}

// Get takes name of the apiManagementCertificate, and returns the corresponding apiManagementCertificate object, and an error if there is any.
func (c *FakeApiManagementCertificates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ApiManagementCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apimanagementcertificatesResource, c.ns, name), &v1alpha1.ApiManagementCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementCertificate), err
}

// List takes label and field selectors, and returns the list of ApiManagementCertificates that match those selectors.
func (c *FakeApiManagementCertificates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ApiManagementCertificateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apimanagementcertificatesResource, apimanagementcertificatesKind, c.ns, opts), &v1alpha1.ApiManagementCertificateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApiManagementCertificateList{ListMeta: obj.(*v1alpha1.ApiManagementCertificateList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApiManagementCertificateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apiManagementCertificates.
func (c *FakeApiManagementCertificates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apimanagementcertificatesResource, c.ns, opts))

}

// Create takes the representation of a apiManagementCertificate and creates it.  Returns the server's representation of the apiManagementCertificate, and an error, if there is any.
func (c *FakeApiManagementCertificates) Create(ctx context.Context, apiManagementCertificate *v1alpha1.ApiManagementCertificate, opts v1.CreateOptions) (result *v1alpha1.ApiManagementCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apimanagementcertificatesResource, c.ns, apiManagementCertificate), &v1alpha1.ApiManagementCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementCertificate), err
}

// Update takes the representation of a apiManagementCertificate and updates it. Returns the server's representation of the apiManagementCertificate, and an error, if there is any.
func (c *FakeApiManagementCertificates) Update(ctx context.Context, apiManagementCertificate *v1alpha1.ApiManagementCertificate, opts v1.UpdateOptions) (result *v1alpha1.ApiManagementCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apimanagementcertificatesResource, c.ns, apiManagementCertificate), &v1alpha1.ApiManagementCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementCertificate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApiManagementCertificates) UpdateStatus(ctx context.Context, apiManagementCertificate *v1alpha1.ApiManagementCertificate, opts v1.UpdateOptions) (*v1alpha1.ApiManagementCertificate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apimanagementcertificatesResource, "status", c.ns, apiManagementCertificate), &v1alpha1.ApiManagementCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementCertificate), err
}

// Delete takes name of the apiManagementCertificate and deletes it. Returns an error if one occurs.
func (c *FakeApiManagementCertificates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(apimanagementcertificatesResource, c.ns, name), &v1alpha1.ApiManagementCertificate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApiManagementCertificates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apimanagementcertificatesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApiManagementCertificateList{})
	return err
}

// Patch applies the patch and returns the patched apiManagementCertificate.
func (c *FakeApiManagementCertificates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApiManagementCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apimanagementcertificatesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApiManagementCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiManagementCertificate), err
}
