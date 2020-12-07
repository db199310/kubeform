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

// FakeIothubEndpointStorageContainers implements IothubEndpointStorageContainerInterface
type FakeIothubEndpointStorageContainers struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var iothubendpointstoragecontainersResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "iothubendpointstoragecontainers"}

var iothubendpointstoragecontainersKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "IothubEndpointStorageContainer"}

// Get takes name of the iothubEndpointStorageContainer, and returns the corresponding iothubEndpointStorageContainer object, and an error if there is any.
func (c *FakeIothubEndpointStorageContainers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.IothubEndpointStorageContainer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(iothubendpointstoragecontainersResource, c.ns, name), &v1alpha1.IothubEndpointStorageContainer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IothubEndpointStorageContainer), err
}

// List takes label and field selectors, and returns the list of IothubEndpointStorageContainers that match those selectors.
func (c *FakeIothubEndpointStorageContainers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.IothubEndpointStorageContainerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(iothubendpointstoragecontainersResource, iothubendpointstoragecontainersKind, c.ns, opts), &v1alpha1.IothubEndpointStorageContainerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IothubEndpointStorageContainerList{ListMeta: obj.(*v1alpha1.IothubEndpointStorageContainerList).ListMeta}
	for _, item := range obj.(*v1alpha1.IothubEndpointStorageContainerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iothubEndpointStorageContainers.
func (c *FakeIothubEndpointStorageContainers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(iothubendpointstoragecontainersResource, c.ns, opts))

}

// Create takes the representation of a iothubEndpointStorageContainer and creates it.  Returns the server's representation of the iothubEndpointStorageContainer, and an error, if there is any.
func (c *FakeIothubEndpointStorageContainers) Create(ctx context.Context, iothubEndpointStorageContainer *v1alpha1.IothubEndpointStorageContainer, opts v1.CreateOptions) (result *v1alpha1.IothubEndpointStorageContainer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(iothubendpointstoragecontainersResource, c.ns, iothubEndpointStorageContainer), &v1alpha1.IothubEndpointStorageContainer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IothubEndpointStorageContainer), err
}

// Update takes the representation of a iothubEndpointStorageContainer and updates it. Returns the server's representation of the iothubEndpointStorageContainer, and an error, if there is any.
func (c *FakeIothubEndpointStorageContainers) Update(ctx context.Context, iothubEndpointStorageContainer *v1alpha1.IothubEndpointStorageContainer, opts v1.UpdateOptions) (result *v1alpha1.IothubEndpointStorageContainer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(iothubendpointstoragecontainersResource, c.ns, iothubEndpointStorageContainer), &v1alpha1.IothubEndpointStorageContainer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IothubEndpointStorageContainer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIothubEndpointStorageContainers) UpdateStatus(ctx context.Context, iothubEndpointStorageContainer *v1alpha1.IothubEndpointStorageContainer, opts v1.UpdateOptions) (*v1alpha1.IothubEndpointStorageContainer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(iothubendpointstoragecontainersResource, "status", c.ns, iothubEndpointStorageContainer), &v1alpha1.IothubEndpointStorageContainer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IothubEndpointStorageContainer), err
}

// Delete takes name of the iothubEndpointStorageContainer and deletes it. Returns an error if one occurs.
func (c *FakeIothubEndpointStorageContainers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(iothubendpointstoragecontainersResource, c.ns, name), &v1alpha1.IothubEndpointStorageContainer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIothubEndpointStorageContainers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(iothubendpointstoragecontainersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.IothubEndpointStorageContainerList{})
	return err
}

// Patch applies the patch and returns the patched iothubEndpointStorageContainer.
func (c *FakeIothubEndpointStorageContainers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.IothubEndpointStorageContainer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(iothubendpointstoragecontainersResource, c.ns, name, pt, data, subresources...), &v1alpha1.IothubEndpointStorageContainer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IothubEndpointStorageContainer), err
}
