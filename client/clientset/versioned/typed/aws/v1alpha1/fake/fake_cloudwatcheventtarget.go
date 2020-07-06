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

// FakeCloudwatchEventTargets implements CloudwatchEventTargetInterface
type FakeCloudwatchEventTargets struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var cloudwatcheventtargetsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "cloudwatcheventtargets"}

var cloudwatcheventtargetsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "CloudwatchEventTarget"}

// Get takes name of the cloudwatchEventTarget, and returns the corresponding cloudwatchEventTarget object, and an error if there is any.
func (c *FakeCloudwatchEventTargets) Get(name string, options v1.GetOptions) (result *v1alpha1.CloudwatchEventTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cloudwatcheventtargetsResource, c.ns, name), &v1alpha1.CloudwatchEventTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchEventTarget), err
}

// List takes label and field selectors, and returns the list of CloudwatchEventTargets that match those selectors.
func (c *FakeCloudwatchEventTargets) List(opts v1.ListOptions) (result *v1alpha1.CloudwatchEventTargetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cloudwatcheventtargetsResource, cloudwatcheventtargetsKind, c.ns, opts), &v1alpha1.CloudwatchEventTargetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CloudwatchEventTargetList{ListMeta: obj.(*v1alpha1.CloudwatchEventTargetList).ListMeta}
	for _, item := range obj.(*v1alpha1.CloudwatchEventTargetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudwatchEventTargets.
func (c *FakeCloudwatchEventTargets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cloudwatcheventtargetsResource, c.ns, opts))

}

// Create takes the representation of a cloudwatchEventTarget and creates it.  Returns the server's representation of the cloudwatchEventTarget, and an error, if there is any.
func (c *FakeCloudwatchEventTargets) Create(cloudwatchEventTarget *v1alpha1.CloudwatchEventTarget) (result *v1alpha1.CloudwatchEventTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cloudwatcheventtargetsResource, c.ns, cloudwatchEventTarget), &v1alpha1.CloudwatchEventTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchEventTarget), err
}

// Update takes the representation of a cloudwatchEventTarget and updates it. Returns the server's representation of the cloudwatchEventTarget, and an error, if there is any.
func (c *FakeCloudwatchEventTargets) Update(cloudwatchEventTarget *v1alpha1.CloudwatchEventTarget) (result *v1alpha1.CloudwatchEventTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cloudwatcheventtargetsResource, c.ns, cloudwatchEventTarget), &v1alpha1.CloudwatchEventTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchEventTarget), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCloudwatchEventTargets) UpdateStatus(cloudwatchEventTarget *v1alpha1.CloudwatchEventTarget) (*v1alpha1.CloudwatchEventTarget, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cloudwatcheventtargetsResource, "status", c.ns, cloudwatchEventTarget), &v1alpha1.CloudwatchEventTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchEventTarget), err
}

// Delete takes name of the cloudwatchEventTarget and deletes it. Returns an error if one occurs.
func (c *FakeCloudwatchEventTargets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cloudwatcheventtargetsResource, c.ns, name), &v1alpha1.CloudwatchEventTarget{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudwatchEventTargets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cloudwatcheventtargetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CloudwatchEventTargetList{})
	return err
}

// Patch applies the patch and returns the patched cloudwatchEventTarget.
func (c *FakeCloudwatchEventTargets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudwatchEventTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cloudwatcheventtargetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.CloudwatchEventTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchEventTarget), err
}
