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

// FakeEventgridEventSubscriptions implements EventgridEventSubscriptionInterface
type FakeEventgridEventSubscriptions struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var eventgrideventsubscriptionsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "eventgrideventsubscriptions"}

var eventgrideventsubscriptionsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "EventgridEventSubscription"}

// Get takes name of the eventgridEventSubscription, and returns the corresponding eventgridEventSubscription object, and an error if there is any.
func (c *FakeEventgridEventSubscriptions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.EventgridEventSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(eventgrideventsubscriptionsResource, c.ns, name), &v1alpha1.EventgridEventSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventgridEventSubscription), err
}

// List takes label and field selectors, and returns the list of EventgridEventSubscriptions that match those selectors.
func (c *FakeEventgridEventSubscriptions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EventgridEventSubscriptionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(eventgrideventsubscriptionsResource, eventgrideventsubscriptionsKind, c.ns, opts), &v1alpha1.EventgridEventSubscriptionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EventgridEventSubscriptionList{ListMeta: obj.(*v1alpha1.EventgridEventSubscriptionList).ListMeta}
	for _, item := range obj.(*v1alpha1.EventgridEventSubscriptionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested eventgridEventSubscriptions.
func (c *FakeEventgridEventSubscriptions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(eventgrideventsubscriptionsResource, c.ns, opts))

}

// Create takes the representation of a eventgridEventSubscription and creates it.  Returns the server's representation of the eventgridEventSubscription, and an error, if there is any.
func (c *FakeEventgridEventSubscriptions) Create(ctx context.Context, eventgridEventSubscription *v1alpha1.EventgridEventSubscription, opts v1.CreateOptions) (result *v1alpha1.EventgridEventSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(eventgrideventsubscriptionsResource, c.ns, eventgridEventSubscription), &v1alpha1.EventgridEventSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventgridEventSubscription), err
}

// Update takes the representation of a eventgridEventSubscription and updates it. Returns the server's representation of the eventgridEventSubscription, and an error, if there is any.
func (c *FakeEventgridEventSubscriptions) Update(ctx context.Context, eventgridEventSubscription *v1alpha1.EventgridEventSubscription, opts v1.UpdateOptions) (result *v1alpha1.EventgridEventSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(eventgrideventsubscriptionsResource, c.ns, eventgridEventSubscription), &v1alpha1.EventgridEventSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventgridEventSubscription), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEventgridEventSubscriptions) UpdateStatus(ctx context.Context, eventgridEventSubscription *v1alpha1.EventgridEventSubscription, opts v1.UpdateOptions) (*v1alpha1.EventgridEventSubscription, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(eventgrideventsubscriptionsResource, "status", c.ns, eventgridEventSubscription), &v1alpha1.EventgridEventSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventgridEventSubscription), err
}

// Delete takes name of the eventgridEventSubscription and deletes it. Returns an error if one occurs.
func (c *FakeEventgridEventSubscriptions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(eventgrideventsubscriptionsResource, c.ns, name), &v1alpha1.EventgridEventSubscription{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEventgridEventSubscriptions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(eventgrideventsubscriptionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.EventgridEventSubscriptionList{})
	return err
}

// Patch applies the patch and returns the patched eventgridEventSubscription.
func (c *FakeEventgridEventSubscriptions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EventgridEventSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(eventgrideventsubscriptionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.EventgridEventSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventgridEventSubscription), err
}
