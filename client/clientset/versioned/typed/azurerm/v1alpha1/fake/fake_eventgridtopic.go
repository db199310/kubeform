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

// FakeEventgridTopics implements EventgridTopicInterface
type FakeEventgridTopics struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var eventgridtopicsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "eventgridtopics"}

var eventgridtopicsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "EventgridTopic"}

// Get takes name of the eventgridTopic, and returns the corresponding eventgridTopic object, and an error if there is any.
func (c *FakeEventgridTopics) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.EventgridTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(eventgridtopicsResource, c.ns, name), &v1alpha1.EventgridTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventgridTopic), err
}

// List takes label and field selectors, and returns the list of EventgridTopics that match those selectors.
func (c *FakeEventgridTopics) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EventgridTopicList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(eventgridtopicsResource, eventgridtopicsKind, c.ns, opts), &v1alpha1.EventgridTopicList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EventgridTopicList{ListMeta: obj.(*v1alpha1.EventgridTopicList).ListMeta}
	for _, item := range obj.(*v1alpha1.EventgridTopicList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested eventgridTopics.
func (c *FakeEventgridTopics) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(eventgridtopicsResource, c.ns, opts))

}

// Create takes the representation of a eventgridTopic and creates it.  Returns the server's representation of the eventgridTopic, and an error, if there is any.
func (c *FakeEventgridTopics) Create(ctx context.Context, eventgridTopic *v1alpha1.EventgridTopic, opts v1.CreateOptions) (result *v1alpha1.EventgridTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(eventgridtopicsResource, c.ns, eventgridTopic), &v1alpha1.EventgridTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventgridTopic), err
}

// Update takes the representation of a eventgridTopic and updates it. Returns the server's representation of the eventgridTopic, and an error, if there is any.
func (c *FakeEventgridTopics) Update(ctx context.Context, eventgridTopic *v1alpha1.EventgridTopic, opts v1.UpdateOptions) (result *v1alpha1.EventgridTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(eventgridtopicsResource, c.ns, eventgridTopic), &v1alpha1.EventgridTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventgridTopic), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEventgridTopics) UpdateStatus(ctx context.Context, eventgridTopic *v1alpha1.EventgridTopic, opts v1.UpdateOptions) (*v1alpha1.EventgridTopic, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(eventgridtopicsResource, "status", c.ns, eventgridTopic), &v1alpha1.EventgridTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventgridTopic), err
}

// Delete takes name of the eventgridTopic and deletes it. Returns an error if one occurs.
func (c *FakeEventgridTopics) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(eventgridtopicsResource, c.ns, name), &v1alpha1.EventgridTopic{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEventgridTopics) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(eventgridtopicsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.EventgridTopicList{})
	return err
}

// Patch applies the patch and returns the patched eventgridTopic.
func (c *FakeEventgridTopics) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EventgridTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(eventgridtopicsResource, c.ns, name, pt, data, subresources...), &v1alpha1.EventgridTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventgridTopic), err
}
