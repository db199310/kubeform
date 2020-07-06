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

// FakePubsubTopics implements PubsubTopicInterface
type FakePubsubTopics struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var pubsubtopicsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "pubsubtopics"}

var pubsubtopicsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "PubsubTopic"}

// Get takes name of the pubsubTopic, and returns the corresponding pubsubTopic object, and an error if there is any.
func (c *FakePubsubTopics) Get(name string, options v1.GetOptions) (result *v1alpha1.PubsubTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pubsubtopicsResource, c.ns, name), &v1alpha1.PubsubTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PubsubTopic), err
}

// List takes label and field selectors, and returns the list of PubsubTopics that match those selectors.
func (c *FakePubsubTopics) List(opts v1.ListOptions) (result *v1alpha1.PubsubTopicList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pubsubtopicsResource, pubsubtopicsKind, c.ns, opts), &v1alpha1.PubsubTopicList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PubsubTopicList{ListMeta: obj.(*v1alpha1.PubsubTopicList).ListMeta}
	for _, item := range obj.(*v1alpha1.PubsubTopicList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pubsubTopics.
func (c *FakePubsubTopics) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pubsubtopicsResource, c.ns, opts))

}

// Create takes the representation of a pubsubTopic and creates it.  Returns the server's representation of the pubsubTopic, and an error, if there is any.
func (c *FakePubsubTopics) Create(pubsubTopic *v1alpha1.PubsubTopic) (result *v1alpha1.PubsubTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pubsubtopicsResource, c.ns, pubsubTopic), &v1alpha1.PubsubTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PubsubTopic), err
}

// Update takes the representation of a pubsubTopic and updates it. Returns the server's representation of the pubsubTopic, and an error, if there is any.
func (c *FakePubsubTopics) Update(pubsubTopic *v1alpha1.PubsubTopic) (result *v1alpha1.PubsubTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pubsubtopicsResource, c.ns, pubsubTopic), &v1alpha1.PubsubTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PubsubTopic), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePubsubTopics) UpdateStatus(pubsubTopic *v1alpha1.PubsubTopic) (*v1alpha1.PubsubTopic, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(pubsubtopicsResource, "status", c.ns, pubsubTopic), &v1alpha1.PubsubTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PubsubTopic), err
}

// Delete takes name of the pubsubTopic and deletes it. Returns an error if one occurs.
func (c *FakePubsubTopics) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(pubsubtopicsResource, c.ns, name), &v1alpha1.PubsubTopic{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePubsubTopics) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pubsubtopicsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.PubsubTopicList{})
	return err
}

// Patch applies the patch and returns the patched pubsubTopic.
func (c *FakePubsubTopics) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PubsubTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pubsubtopicsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PubsubTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PubsubTopic), err
}
