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

// FakeStoragegatewayCachedIscsiVolumes implements StoragegatewayCachedIscsiVolumeInterface
type FakeStoragegatewayCachedIscsiVolumes struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var storagegatewaycachediscsivolumesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "storagegatewaycachediscsivolumes"}

var storagegatewaycachediscsivolumesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "StoragegatewayCachedIscsiVolume"}

// Get takes name of the storagegatewayCachedIscsiVolume, and returns the corresponding storagegatewayCachedIscsiVolume object, and an error if there is any.
func (c *FakeStoragegatewayCachedIscsiVolumes) Get(name string, options v1.GetOptions) (result *v1alpha1.StoragegatewayCachedIscsiVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(storagegatewaycachediscsivolumesResource, c.ns, name), &v1alpha1.StoragegatewayCachedIscsiVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StoragegatewayCachedIscsiVolume), err
}

// List takes label and field selectors, and returns the list of StoragegatewayCachedIscsiVolumes that match those selectors.
func (c *FakeStoragegatewayCachedIscsiVolumes) List(opts v1.ListOptions) (result *v1alpha1.StoragegatewayCachedIscsiVolumeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(storagegatewaycachediscsivolumesResource, storagegatewaycachediscsivolumesKind, c.ns, opts), &v1alpha1.StoragegatewayCachedIscsiVolumeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StoragegatewayCachedIscsiVolumeList{ListMeta: obj.(*v1alpha1.StoragegatewayCachedIscsiVolumeList).ListMeta}
	for _, item := range obj.(*v1alpha1.StoragegatewayCachedIscsiVolumeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested storagegatewayCachedIscsiVolumes.
func (c *FakeStoragegatewayCachedIscsiVolumes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(storagegatewaycachediscsivolumesResource, c.ns, opts))

}

// Create takes the representation of a storagegatewayCachedIscsiVolume and creates it.  Returns the server's representation of the storagegatewayCachedIscsiVolume, and an error, if there is any.
func (c *FakeStoragegatewayCachedIscsiVolumes) Create(storagegatewayCachedIscsiVolume *v1alpha1.StoragegatewayCachedIscsiVolume) (result *v1alpha1.StoragegatewayCachedIscsiVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(storagegatewaycachediscsivolumesResource, c.ns, storagegatewayCachedIscsiVolume), &v1alpha1.StoragegatewayCachedIscsiVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StoragegatewayCachedIscsiVolume), err
}

// Update takes the representation of a storagegatewayCachedIscsiVolume and updates it. Returns the server's representation of the storagegatewayCachedIscsiVolume, and an error, if there is any.
func (c *FakeStoragegatewayCachedIscsiVolumes) Update(storagegatewayCachedIscsiVolume *v1alpha1.StoragegatewayCachedIscsiVolume) (result *v1alpha1.StoragegatewayCachedIscsiVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(storagegatewaycachediscsivolumesResource, c.ns, storagegatewayCachedIscsiVolume), &v1alpha1.StoragegatewayCachedIscsiVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StoragegatewayCachedIscsiVolume), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStoragegatewayCachedIscsiVolumes) UpdateStatus(storagegatewayCachedIscsiVolume *v1alpha1.StoragegatewayCachedIscsiVolume) (*v1alpha1.StoragegatewayCachedIscsiVolume, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(storagegatewaycachediscsivolumesResource, "status", c.ns, storagegatewayCachedIscsiVolume), &v1alpha1.StoragegatewayCachedIscsiVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StoragegatewayCachedIscsiVolume), err
}

// Delete takes name of the storagegatewayCachedIscsiVolume and deletes it. Returns an error if one occurs.
func (c *FakeStoragegatewayCachedIscsiVolumes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(storagegatewaycachediscsivolumesResource, c.ns, name), &v1alpha1.StoragegatewayCachedIscsiVolume{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStoragegatewayCachedIscsiVolumes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(storagegatewaycachediscsivolumesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.StoragegatewayCachedIscsiVolumeList{})
	return err
}

// Patch applies the patch and returns the patched storagegatewayCachedIscsiVolume.
func (c *FakeStoragegatewayCachedIscsiVolumes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StoragegatewayCachedIscsiVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(storagegatewaycachediscsivolumesResource, c.ns, name, pt, data, subresources...), &v1alpha1.StoragegatewayCachedIscsiVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StoragegatewayCachedIscsiVolume), err
}
