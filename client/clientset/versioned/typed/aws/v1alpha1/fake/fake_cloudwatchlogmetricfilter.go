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

// FakeCloudwatchLogMetricFilters implements CloudwatchLogMetricFilterInterface
type FakeCloudwatchLogMetricFilters struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var cloudwatchlogmetricfiltersResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "cloudwatchlogmetricfilters"}

var cloudwatchlogmetricfiltersKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "CloudwatchLogMetricFilter"}

// Get takes name of the cloudwatchLogMetricFilter, and returns the corresponding cloudwatchLogMetricFilter object, and an error if there is any.
func (c *FakeCloudwatchLogMetricFilters) Get(name string, options v1.GetOptions) (result *v1alpha1.CloudwatchLogMetricFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cloudwatchlogmetricfiltersResource, c.ns, name), &v1alpha1.CloudwatchLogMetricFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchLogMetricFilter), err
}

// List takes label and field selectors, and returns the list of CloudwatchLogMetricFilters that match those selectors.
func (c *FakeCloudwatchLogMetricFilters) List(opts v1.ListOptions) (result *v1alpha1.CloudwatchLogMetricFilterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cloudwatchlogmetricfiltersResource, cloudwatchlogmetricfiltersKind, c.ns, opts), &v1alpha1.CloudwatchLogMetricFilterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CloudwatchLogMetricFilterList{ListMeta: obj.(*v1alpha1.CloudwatchLogMetricFilterList).ListMeta}
	for _, item := range obj.(*v1alpha1.CloudwatchLogMetricFilterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudwatchLogMetricFilters.
func (c *FakeCloudwatchLogMetricFilters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cloudwatchlogmetricfiltersResource, c.ns, opts))

}

// Create takes the representation of a cloudwatchLogMetricFilter and creates it.  Returns the server's representation of the cloudwatchLogMetricFilter, and an error, if there is any.
func (c *FakeCloudwatchLogMetricFilters) Create(cloudwatchLogMetricFilter *v1alpha1.CloudwatchLogMetricFilter) (result *v1alpha1.CloudwatchLogMetricFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cloudwatchlogmetricfiltersResource, c.ns, cloudwatchLogMetricFilter), &v1alpha1.CloudwatchLogMetricFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchLogMetricFilter), err
}

// Update takes the representation of a cloudwatchLogMetricFilter and updates it. Returns the server's representation of the cloudwatchLogMetricFilter, and an error, if there is any.
func (c *FakeCloudwatchLogMetricFilters) Update(cloudwatchLogMetricFilter *v1alpha1.CloudwatchLogMetricFilter) (result *v1alpha1.CloudwatchLogMetricFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cloudwatchlogmetricfiltersResource, c.ns, cloudwatchLogMetricFilter), &v1alpha1.CloudwatchLogMetricFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchLogMetricFilter), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCloudwatchLogMetricFilters) UpdateStatus(cloudwatchLogMetricFilter *v1alpha1.CloudwatchLogMetricFilter) (*v1alpha1.CloudwatchLogMetricFilter, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cloudwatchlogmetricfiltersResource, "status", c.ns, cloudwatchLogMetricFilter), &v1alpha1.CloudwatchLogMetricFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchLogMetricFilter), err
}

// Delete takes name of the cloudwatchLogMetricFilter and deletes it. Returns an error if one occurs.
func (c *FakeCloudwatchLogMetricFilters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cloudwatchlogmetricfiltersResource, c.ns, name), &v1alpha1.CloudwatchLogMetricFilter{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudwatchLogMetricFilters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cloudwatchlogmetricfiltersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CloudwatchLogMetricFilterList{})
	return err
}

// Patch applies the patch and returns the patched cloudwatchLogMetricFilter.
func (c *FakeCloudwatchLogMetricFilters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudwatchLogMetricFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cloudwatchlogmetricfiltersResource, c.ns, name, pt, data, subresources...), &v1alpha1.CloudwatchLogMetricFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchLogMetricFilter), err
}
