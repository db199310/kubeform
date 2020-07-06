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

// FakeCloudwatchLogSubscriptionFilters implements CloudwatchLogSubscriptionFilterInterface
type FakeCloudwatchLogSubscriptionFilters struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var cloudwatchlogsubscriptionfiltersResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "cloudwatchlogsubscriptionfilters"}

var cloudwatchlogsubscriptionfiltersKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "CloudwatchLogSubscriptionFilter"}

// Get takes name of the cloudwatchLogSubscriptionFilter, and returns the corresponding cloudwatchLogSubscriptionFilter object, and an error if there is any.
func (c *FakeCloudwatchLogSubscriptionFilters) Get(name string, options v1.GetOptions) (result *v1alpha1.CloudwatchLogSubscriptionFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cloudwatchlogsubscriptionfiltersResource, c.ns, name), &v1alpha1.CloudwatchLogSubscriptionFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchLogSubscriptionFilter), err
}

// List takes label and field selectors, and returns the list of CloudwatchLogSubscriptionFilters that match those selectors.
func (c *FakeCloudwatchLogSubscriptionFilters) List(opts v1.ListOptions) (result *v1alpha1.CloudwatchLogSubscriptionFilterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cloudwatchlogsubscriptionfiltersResource, cloudwatchlogsubscriptionfiltersKind, c.ns, opts), &v1alpha1.CloudwatchLogSubscriptionFilterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CloudwatchLogSubscriptionFilterList{ListMeta: obj.(*v1alpha1.CloudwatchLogSubscriptionFilterList).ListMeta}
	for _, item := range obj.(*v1alpha1.CloudwatchLogSubscriptionFilterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudwatchLogSubscriptionFilters.
func (c *FakeCloudwatchLogSubscriptionFilters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cloudwatchlogsubscriptionfiltersResource, c.ns, opts))

}

// Create takes the representation of a cloudwatchLogSubscriptionFilter and creates it.  Returns the server's representation of the cloudwatchLogSubscriptionFilter, and an error, if there is any.
func (c *FakeCloudwatchLogSubscriptionFilters) Create(cloudwatchLogSubscriptionFilter *v1alpha1.CloudwatchLogSubscriptionFilter) (result *v1alpha1.CloudwatchLogSubscriptionFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cloudwatchlogsubscriptionfiltersResource, c.ns, cloudwatchLogSubscriptionFilter), &v1alpha1.CloudwatchLogSubscriptionFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchLogSubscriptionFilter), err
}

// Update takes the representation of a cloudwatchLogSubscriptionFilter and updates it. Returns the server's representation of the cloudwatchLogSubscriptionFilter, and an error, if there is any.
func (c *FakeCloudwatchLogSubscriptionFilters) Update(cloudwatchLogSubscriptionFilter *v1alpha1.CloudwatchLogSubscriptionFilter) (result *v1alpha1.CloudwatchLogSubscriptionFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cloudwatchlogsubscriptionfiltersResource, c.ns, cloudwatchLogSubscriptionFilter), &v1alpha1.CloudwatchLogSubscriptionFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchLogSubscriptionFilter), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCloudwatchLogSubscriptionFilters) UpdateStatus(cloudwatchLogSubscriptionFilter *v1alpha1.CloudwatchLogSubscriptionFilter) (*v1alpha1.CloudwatchLogSubscriptionFilter, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cloudwatchlogsubscriptionfiltersResource, "status", c.ns, cloudwatchLogSubscriptionFilter), &v1alpha1.CloudwatchLogSubscriptionFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchLogSubscriptionFilter), err
}

// Delete takes name of the cloudwatchLogSubscriptionFilter and deletes it. Returns an error if one occurs.
func (c *FakeCloudwatchLogSubscriptionFilters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cloudwatchlogsubscriptionfiltersResource, c.ns, name), &v1alpha1.CloudwatchLogSubscriptionFilter{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudwatchLogSubscriptionFilters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cloudwatchlogsubscriptionfiltersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CloudwatchLogSubscriptionFilterList{})
	return err
}

// Patch applies the patch and returns the patched cloudwatchLogSubscriptionFilter.
func (c *FakeCloudwatchLogSubscriptionFilters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudwatchLogSubscriptionFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cloudwatchlogsubscriptionfiltersResource, c.ns, name, pt, data, subresources...), &v1alpha1.CloudwatchLogSubscriptionFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchLogSubscriptionFilter), err
}
