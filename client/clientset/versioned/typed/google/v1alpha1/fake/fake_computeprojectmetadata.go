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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakeComputeProjectMetadatas implements ComputeProjectMetadataInterface
type FakeComputeProjectMetadatas struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var computeprojectmetadatasResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "computeprojectmetadatas"}

var computeprojectmetadatasKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "ComputeProjectMetadata"}

// Get takes name of the computeProjectMetadata, and returns the corresponding computeProjectMetadata object, and an error if there is any.
func (c *FakeComputeProjectMetadatas) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeProjectMetadata, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computeprojectmetadatasResource, c.ns, name), &v1alpha1.ComputeProjectMetadata{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeProjectMetadata), err
}

// List takes label and field selectors, and returns the list of ComputeProjectMetadatas that match those selectors.
func (c *FakeComputeProjectMetadatas) List(opts v1.ListOptions) (result *v1alpha1.ComputeProjectMetadataList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computeprojectmetadatasResource, computeprojectmetadatasKind, c.ns, opts), &v1alpha1.ComputeProjectMetadataList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComputeProjectMetadataList{ListMeta: obj.(*v1alpha1.ComputeProjectMetadataList).ListMeta}
	for _, item := range obj.(*v1alpha1.ComputeProjectMetadataList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeProjectMetadatas.
func (c *FakeComputeProjectMetadatas) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computeprojectmetadatasResource, c.ns, opts))

}

// Create takes the representation of a computeProjectMetadata and creates it.  Returns the server's representation of the computeProjectMetadata, and an error, if there is any.
func (c *FakeComputeProjectMetadatas) Create(computeProjectMetadata *v1alpha1.ComputeProjectMetadata) (result *v1alpha1.ComputeProjectMetadata, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computeprojectmetadatasResource, c.ns, computeProjectMetadata), &v1alpha1.ComputeProjectMetadata{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeProjectMetadata), err
}

// Update takes the representation of a computeProjectMetadata and updates it. Returns the server's representation of the computeProjectMetadata, and an error, if there is any.
func (c *FakeComputeProjectMetadatas) Update(computeProjectMetadata *v1alpha1.ComputeProjectMetadata) (result *v1alpha1.ComputeProjectMetadata, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computeprojectmetadatasResource, c.ns, computeProjectMetadata), &v1alpha1.ComputeProjectMetadata{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeProjectMetadata), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeProjectMetadatas) UpdateStatus(computeProjectMetadata *v1alpha1.ComputeProjectMetadata) (*v1alpha1.ComputeProjectMetadata, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computeprojectmetadatasResource, "status", c.ns, computeProjectMetadata), &v1alpha1.ComputeProjectMetadata{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeProjectMetadata), err
}

// Delete takes name of the computeProjectMetadata and deletes it. Returns an error if one occurs.
func (c *FakeComputeProjectMetadatas) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(computeprojectmetadatasResource, c.ns, name), &v1alpha1.ComputeProjectMetadata{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeProjectMetadatas) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computeprojectmetadatasResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComputeProjectMetadataList{})
	return err
}

// Patch applies the patch and returns the patched computeProjectMetadata.
func (c *FakeComputeProjectMetadatas) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeProjectMetadata, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computeprojectmetadatasResource, c.ns, name, pt, data, subresources...), &v1alpha1.ComputeProjectMetadata{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeProjectMetadata), err
}
