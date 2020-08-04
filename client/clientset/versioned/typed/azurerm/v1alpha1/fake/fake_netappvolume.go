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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakeNetappVolumes implements NetappVolumeInterface
type FakeNetappVolumes struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var netappvolumesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "netappvolumes"}

var netappvolumesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "NetappVolume"}

// Get takes name of the netappVolume, and returns the corresponding netappVolume object, and an error if there is any.
func (c *FakeNetappVolumes) Get(name string, options v1.GetOptions) (result *v1alpha1.NetappVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(netappvolumesResource, c.ns, name), &v1alpha1.NetappVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetappVolume), err
}

// List takes label and field selectors, and returns the list of NetappVolumes that match those selectors.
func (c *FakeNetappVolumes) List(opts v1.ListOptions) (result *v1alpha1.NetappVolumeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(netappvolumesResource, netappvolumesKind, c.ns, opts), &v1alpha1.NetappVolumeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NetappVolumeList{ListMeta: obj.(*v1alpha1.NetappVolumeList).ListMeta}
	for _, item := range obj.(*v1alpha1.NetappVolumeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested netappVolumes.
func (c *FakeNetappVolumes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(netappvolumesResource, c.ns, opts))

}

// Create takes the representation of a netappVolume and creates it.  Returns the server's representation of the netappVolume, and an error, if there is any.
func (c *FakeNetappVolumes) Create(netappVolume *v1alpha1.NetappVolume) (result *v1alpha1.NetappVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(netappvolumesResource, c.ns, netappVolume), &v1alpha1.NetappVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetappVolume), err
}

// Update takes the representation of a netappVolume and updates it. Returns the server's representation of the netappVolume, and an error, if there is any.
func (c *FakeNetappVolumes) Update(netappVolume *v1alpha1.NetappVolume) (result *v1alpha1.NetappVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(netappvolumesResource, c.ns, netappVolume), &v1alpha1.NetappVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetappVolume), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetappVolumes) UpdateStatus(netappVolume *v1alpha1.NetappVolume) (*v1alpha1.NetappVolume, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(netappvolumesResource, "status", c.ns, netappVolume), &v1alpha1.NetappVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetappVolume), err
}

// Delete takes name of the netappVolume and deletes it. Returns an error if one occurs.
func (c *FakeNetappVolumes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(netappvolumesResource, c.ns, name), &v1alpha1.NetappVolume{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetappVolumes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(netappvolumesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.NetappVolumeList{})
	return err
}

// Patch applies the patch and returns the patched netappVolume.
func (c *FakeNetappVolumes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NetappVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(netappvolumesResource, c.ns, name, pt, data, subresources...), &v1alpha1.NetappVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetappVolume), err
}
