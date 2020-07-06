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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeApiGatewayClientCertificates implements ApiGatewayClientCertificateInterface
type FakeApiGatewayClientCertificates struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var apigatewayclientcertificatesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "apigatewayclientcertificates"}

var apigatewayclientcertificatesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "ApiGatewayClientCertificate"}

// Get takes name of the apiGatewayClientCertificate, and returns the corresponding apiGatewayClientCertificate object, and an error if there is any.
func (c *FakeApiGatewayClientCertificates) Get(name string, options v1.GetOptions) (result *v1alpha1.ApiGatewayClientCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apigatewayclientcertificatesResource, c.ns, name), &v1alpha1.ApiGatewayClientCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayClientCertificate), err
}

// List takes label and field selectors, and returns the list of ApiGatewayClientCertificates that match those selectors.
func (c *FakeApiGatewayClientCertificates) List(opts v1.ListOptions) (result *v1alpha1.ApiGatewayClientCertificateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apigatewayclientcertificatesResource, apigatewayclientcertificatesKind, c.ns, opts), &v1alpha1.ApiGatewayClientCertificateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApiGatewayClientCertificateList{ListMeta: obj.(*v1alpha1.ApiGatewayClientCertificateList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApiGatewayClientCertificateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apiGatewayClientCertificates.
func (c *FakeApiGatewayClientCertificates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apigatewayclientcertificatesResource, c.ns, opts))

}

// Create takes the representation of a apiGatewayClientCertificate and creates it.  Returns the server's representation of the apiGatewayClientCertificate, and an error, if there is any.
func (c *FakeApiGatewayClientCertificates) Create(apiGatewayClientCertificate *v1alpha1.ApiGatewayClientCertificate) (result *v1alpha1.ApiGatewayClientCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apigatewayclientcertificatesResource, c.ns, apiGatewayClientCertificate), &v1alpha1.ApiGatewayClientCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayClientCertificate), err
}

// Update takes the representation of a apiGatewayClientCertificate and updates it. Returns the server's representation of the apiGatewayClientCertificate, and an error, if there is any.
func (c *FakeApiGatewayClientCertificates) Update(apiGatewayClientCertificate *v1alpha1.ApiGatewayClientCertificate) (result *v1alpha1.ApiGatewayClientCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apigatewayclientcertificatesResource, c.ns, apiGatewayClientCertificate), &v1alpha1.ApiGatewayClientCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayClientCertificate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApiGatewayClientCertificates) UpdateStatus(apiGatewayClientCertificate *v1alpha1.ApiGatewayClientCertificate) (*v1alpha1.ApiGatewayClientCertificate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apigatewayclientcertificatesResource, "status", c.ns, apiGatewayClientCertificate), &v1alpha1.ApiGatewayClientCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayClientCertificate), err
}

// Delete takes name of the apiGatewayClientCertificate and deletes it. Returns an error if one occurs.
func (c *FakeApiGatewayClientCertificates) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(apigatewayclientcertificatesResource, c.ns, name), &v1alpha1.ApiGatewayClientCertificate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApiGatewayClientCertificates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apigatewayclientcertificatesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApiGatewayClientCertificateList{})
	return err
}

// Patch applies the patch and returns the patched apiGatewayClientCertificate.
func (c *FakeApiGatewayClientCertificates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiGatewayClientCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apigatewayclientcertificatesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApiGatewayClientCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayClientCertificate), err
}
