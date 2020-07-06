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

// FakeSesDomainMailFroms implements SesDomainMailFromInterface
type FakeSesDomainMailFroms struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var sesdomainmailfromsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "sesdomainmailfroms"}

var sesdomainmailfromsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "SesDomainMailFrom"}

// Get takes name of the sesDomainMailFrom, and returns the corresponding sesDomainMailFrom object, and an error if there is any.
func (c *FakeSesDomainMailFroms) Get(name string, options v1.GetOptions) (result *v1alpha1.SesDomainMailFrom, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sesdomainmailfromsResource, c.ns, name), &v1alpha1.SesDomainMailFrom{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SesDomainMailFrom), err
}

// List takes label and field selectors, and returns the list of SesDomainMailFroms that match those selectors.
func (c *FakeSesDomainMailFroms) List(opts v1.ListOptions) (result *v1alpha1.SesDomainMailFromList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sesdomainmailfromsResource, sesdomainmailfromsKind, c.ns, opts), &v1alpha1.SesDomainMailFromList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SesDomainMailFromList{ListMeta: obj.(*v1alpha1.SesDomainMailFromList).ListMeta}
	for _, item := range obj.(*v1alpha1.SesDomainMailFromList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sesDomainMailFroms.
func (c *FakeSesDomainMailFroms) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sesdomainmailfromsResource, c.ns, opts))

}

// Create takes the representation of a sesDomainMailFrom and creates it.  Returns the server's representation of the sesDomainMailFrom, and an error, if there is any.
func (c *FakeSesDomainMailFroms) Create(sesDomainMailFrom *v1alpha1.SesDomainMailFrom) (result *v1alpha1.SesDomainMailFrom, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sesdomainmailfromsResource, c.ns, sesDomainMailFrom), &v1alpha1.SesDomainMailFrom{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SesDomainMailFrom), err
}

// Update takes the representation of a sesDomainMailFrom and updates it. Returns the server's representation of the sesDomainMailFrom, and an error, if there is any.
func (c *FakeSesDomainMailFroms) Update(sesDomainMailFrom *v1alpha1.SesDomainMailFrom) (result *v1alpha1.SesDomainMailFrom, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sesdomainmailfromsResource, c.ns, sesDomainMailFrom), &v1alpha1.SesDomainMailFrom{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SesDomainMailFrom), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSesDomainMailFroms) UpdateStatus(sesDomainMailFrom *v1alpha1.SesDomainMailFrom) (*v1alpha1.SesDomainMailFrom, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sesdomainmailfromsResource, "status", c.ns, sesDomainMailFrom), &v1alpha1.SesDomainMailFrom{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SesDomainMailFrom), err
}

// Delete takes name of the sesDomainMailFrom and deletes it. Returns an error if one occurs.
func (c *FakeSesDomainMailFroms) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sesdomainmailfromsResource, c.ns, name), &v1alpha1.SesDomainMailFrom{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSesDomainMailFroms) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sesdomainmailfromsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SesDomainMailFromList{})
	return err
}

// Patch applies the patch and returns the patched sesDomainMailFrom.
func (c *FakeSesDomainMailFroms) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SesDomainMailFrom, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sesdomainmailfromsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SesDomainMailFrom{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SesDomainMailFrom), err
}
