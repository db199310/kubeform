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

// FakeStreamAnalyticsOutputServicebusTopics implements StreamAnalyticsOutputServicebusTopicInterface
type FakeStreamAnalyticsOutputServicebusTopics struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var streamanalyticsoutputservicebustopicsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "streamanalyticsoutputservicebustopics"}

var streamanalyticsoutputservicebustopicsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "StreamAnalyticsOutputServicebusTopic"}

// Get takes name of the streamAnalyticsOutputServicebusTopic, and returns the corresponding streamAnalyticsOutputServicebusTopic object, and an error if there is any.
func (c *FakeStreamAnalyticsOutputServicebusTopics) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.StreamAnalyticsOutputServicebusTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(streamanalyticsoutputservicebustopicsResource, c.ns, name), &v1alpha1.StreamAnalyticsOutputServicebusTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StreamAnalyticsOutputServicebusTopic), err
}

// List takes label and field selectors, and returns the list of StreamAnalyticsOutputServicebusTopics that match those selectors.
func (c *FakeStreamAnalyticsOutputServicebusTopics) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.StreamAnalyticsOutputServicebusTopicList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(streamanalyticsoutputservicebustopicsResource, streamanalyticsoutputservicebustopicsKind, c.ns, opts), &v1alpha1.StreamAnalyticsOutputServicebusTopicList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StreamAnalyticsOutputServicebusTopicList{ListMeta: obj.(*v1alpha1.StreamAnalyticsOutputServicebusTopicList).ListMeta}
	for _, item := range obj.(*v1alpha1.StreamAnalyticsOutputServicebusTopicList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested streamAnalyticsOutputServicebusTopics.
func (c *FakeStreamAnalyticsOutputServicebusTopics) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(streamanalyticsoutputservicebustopicsResource, c.ns, opts))

}

// Create takes the representation of a streamAnalyticsOutputServicebusTopic and creates it.  Returns the server's representation of the streamAnalyticsOutputServicebusTopic, and an error, if there is any.
func (c *FakeStreamAnalyticsOutputServicebusTopics) Create(ctx context.Context, streamAnalyticsOutputServicebusTopic *v1alpha1.StreamAnalyticsOutputServicebusTopic, opts v1.CreateOptions) (result *v1alpha1.StreamAnalyticsOutputServicebusTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(streamanalyticsoutputservicebustopicsResource, c.ns, streamAnalyticsOutputServicebusTopic), &v1alpha1.StreamAnalyticsOutputServicebusTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StreamAnalyticsOutputServicebusTopic), err
}

// Update takes the representation of a streamAnalyticsOutputServicebusTopic and updates it. Returns the server's representation of the streamAnalyticsOutputServicebusTopic, and an error, if there is any.
func (c *FakeStreamAnalyticsOutputServicebusTopics) Update(ctx context.Context, streamAnalyticsOutputServicebusTopic *v1alpha1.StreamAnalyticsOutputServicebusTopic, opts v1.UpdateOptions) (result *v1alpha1.StreamAnalyticsOutputServicebusTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(streamanalyticsoutputservicebustopicsResource, c.ns, streamAnalyticsOutputServicebusTopic), &v1alpha1.StreamAnalyticsOutputServicebusTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StreamAnalyticsOutputServicebusTopic), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStreamAnalyticsOutputServicebusTopics) UpdateStatus(ctx context.Context, streamAnalyticsOutputServicebusTopic *v1alpha1.StreamAnalyticsOutputServicebusTopic, opts v1.UpdateOptions) (*v1alpha1.StreamAnalyticsOutputServicebusTopic, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(streamanalyticsoutputservicebustopicsResource, "status", c.ns, streamAnalyticsOutputServicebusTopic), &v1alpha1.StreamAnalyticsOutputServicebusTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StreamAnalyticsOutputServicebusTopic), err
}

// Delete takes name of the streamAnalyticsOutputServicebusTopic and deletes it. Returns an error if one occurs.
func (c *FakeStreamAnalyticsOutputServicebusTopics) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(streamanalyticsoutputservicebustopicsResource, c.ns, name), &v1alpha1.StreamAnalyticsOutputServicebusTopic{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStreamAnalyticsOutputServicebusTopics) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(streamanalyticsoutputservicebustopicsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.StreamAnalyticsOutputServicebusTopicList{})
	return err
}

// Patch applies the patch and returns the patched streamAnalyticsOutputServicebusTopic.
func (c *FakeStreamAnalyticsOutputServicebusTopics) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.StreamAnalyticsOutputServicebusTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(streamanalyticsoutputservicebustopicsResource, c.ns, name, pt, data, subresources...), &v1alpha1.StreamAnalyticsOutputServicebusTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StreamAnalyticsOutputServicebusTopic), err
}
