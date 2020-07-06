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

// FakeAlbs implements AlbInterface
type FakeAlbs struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var albsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "albs"}

var albsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "Alb"}

// Get takes name of the alb, and returns the corresponding alb object, and an error if there is any.
func (c *FakeAlbs) Get(name string, options v1.GetOptions) (result *v1alpha1.Alb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(albsResource, c.ns, name), &v1alpha1.Alb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Alb), err
}

// List takes label and field selectors, and returns the list of Albs that match those selectors.
func (c *FakeAlbs) List(opts v1.ListOptions) (result *v1alpha1.AlbList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(albsResource, albsKind, c.ns, opts), &v1alpha1.AlbList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AlbList{ListMeta: obj.(*v1alpha1.AlbList).ListMeta}
	for _, item := range obj.(*v1alpha1.AlbList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested albs.
func (c *FakeAlbs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(albsResource, c.ns, opts))

}

// Create takes the representation of a alb and creates it.  Returns the server's representation of the alb, and an error, if there is any.
func (c *FakeAlbs) Create(alb *v1alpha1.Alb) (result *v1alpha1.Alb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(albsResource, c.ns, alb), &v1alpha1.Alb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Alb), err
}

// Update takes the representation of a alb and updates it. Returns the server's representation of the alb, and an error, if there is any.
func (c *FakeAlbs) Update(alb *v1alpha1.Alb) (result *v1alpha1.Alb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(albsResource, c.ns, alb), &v1alpha1.Alb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Alb), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAlbs) UpdateStatus(alb *v1alpha1.Alb) (*v1alpha1.Alb, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(albsResource, "status", c.ns, alb), &v1alpha1.Alb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Alb), err
}

// Delete takes name of the alb and deletes it. Returns an error if one occurs.
func (c *FakeAlbs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(albsResource, c.ns, name), &v1alpha1.Alb{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAlbs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(albsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AlbList{})
	return err
}

// Patch applies the patch and returns the patched alb.
func (c *FakeAlbs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Alb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(albsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Alb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Alb), err
}
