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

// FakeEventhubNamespace_s implements EventhubNamespace_Interface
type FakeEventhubNamespace_s struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var eventhubnamespace_sResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "eventhubnamespace_s"}

var eventhubnamespace_sKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "EventhubNamespace_"}

// Get takes name of the eventhubNamespace_, and returns the corresponding eventhubNamespace_ object, and an error if there is any.
func (c *FakeEventhubNamespace_s) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.EventhubNamespace_, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(eventhubnamespace_sResource, c.ns, name), &v1alpha1.EventhubNamespace_{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventhubNamespace_), err
}

// List takes label and field selectors, and returns the list of EventhubNamespace_s that match those selectors.
func (c *FakeEventhubNamespace_s) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EventhubNamespace_List, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(eventhubnamespace_sResource, eventhubnamespace_sKind, c.ns, opts), &v1alpha1.EventhubNamespace_List{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EventhubNamespace_List{ListMeta: obj.(*v1alpha1.EventhubNamespace_List).ListMeta}
	for _, item := range obj.(*v1alpha1.EventhubNamespace_List).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested eventhubNamespace_s.
func (c *FakeEventhubNamespace_s) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(eventhubnamespace_sResource, c.ns, opts))

}

// Create takes the representation of a eventhubNamespace_ and creates it.  Returns the server's representation of the eventhubNamespace_, and an error, if there is any.
func (c *FakeEventhubNamespace_s) Create(ctx context.Context, eventhubNamespace_ *v1alpha1.EventhubNamespace_, opts v1.CreateOptions) (result *v1alpha1.EventhubNamespace_, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(eventhubnamespace_sResource, c.ns, eventhubNamespace_), &v1alpha1.EventhubNamespace_{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventhubNamespace_), err
}

// Update takes the representation of a eventhubNamespace_ and updates it. Returns the server's representation of the eventhubNamespace_, and an error, if there is any.
func (c *FakeEventhubNamespace_s) Update(ctx context.Context, eventhubNamespace_ *v1alpha1.EventhubNamespace_, opts v1.UpdateOptions) (result *v1alpha1.EventhubNamespace_, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(eventhubnamespace_sResource, c.ns, eventhubNamespace_), &v1alpha1.EventhubNamespace_{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventhubNamespace_), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEventhubNamespace_s) UpdateStatus(ctx context.Context, eventhubNamespace_ *v1alpha1.EventhubNamespace_, opts v1.UpdateOptions) (*v1alpha1.EventhubNamespace_, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(eventhubnamespace_sResource, "status", c.ns, eventhubNamespace_), &v1alpha1.EventhubNamespace_{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventhubNamespace_), err
}

// Delete takes name of the eventhubNamespace_ and deletes it. Returns an error if one occurs.
func (c *FakeEventhubNamespace_s) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(eventhubnamespace_sResource, c.ns, name), &v1alpha1.EventhubNamespace_{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEventhubNamespace_s) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(eventhubnamespace_sResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.EventhubNamespace_List{})
	return err
}

// Patch applies the patch and returns the patched eventhubNamespace_.
func (c *FakeEventhubNamespace_s) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EventhubNamespace_, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(eventhubnamespace_sResource, c.ns, name, pt, data, subresources...), &v1alpha1.EventhubNamespace_{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventhubNamespace_), err
}
