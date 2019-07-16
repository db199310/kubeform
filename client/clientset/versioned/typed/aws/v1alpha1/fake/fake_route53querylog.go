/*
Copyright 2019 The Kubeform Authors.

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

// FakeRoute53QueryLogs implements Route53QueryLogInterface
type FakeRoute53QueryLogs struct {
	Fake *FakeAwsV1alpha1
}

var route53querylogsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "route53querylogs"}

var route53querylogsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "Route53QueryLog"}

// Get takes name of the route53QueryLog, and returns the corresponding route53QueryLog object, and an error if there is any.
func (c *FakeRoute53QueryLogs) Get(name string, options v1.GetOptions) (result *v1alpha1.Route53QueryLog, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(route53querylogsResource, name), &v1alpha1.Route53QueryLog{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Route53QueryLog), err
}

// List takes label and field selectors, and returns the list of Route53QueryLogs that match those selectors.
func (c *FakeRoute53QueryLogs) List(opts v1.ListOptions) (result *v1alpha1.Route53QueryLogList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(route53querylogsResource, route53querylogsKind, opts), &v1alpha1.Route53QueryLogList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.Route53QueryLogList{ListMeta: obj.(*v1alpha1.Route53QueryLogList).ListMeta}
	for _, item := range obj.(*v1alpha1.Route53QueryLogList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested route53QueryLogs.
func (c *FakeRoute53QueryLogs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(route53querylogsResource, opts))
}

// Create takes the representation of a route53QueryLog and creates it.  Returns the server's representation of the route53QueryLog, and an error, if there is any.
func (c *FakeRoute53QueryLogs) Create(route53QueryLog *v1alpha1.Route53QueryLog) (result *v1alpha1.Route53QueryLog, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(route53querylogsResource, route53QueryLog), &v1alpha1.Route53QueryLog{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Route53QueryLog), err
}

// Update takes the representation of a route53QueryLog and updates it. Returns the server's representation of the route53QueryLog, and an error, if there is any.
func (c *FakeRoute53QueryLogs) Update(route53QueryLog *v1alpha1.Route53QueryLog) (result *v1alpha1.Route53QueryLog, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(route53querylogsResource, route53QueryLog), &v1alpha1.Route53QueryLog{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Route53QueryLog), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRoute53QueryLogs) UpdateStatus(route53QueryLog *v1alpha1.Route53QueryLog) (*v1alpha1.Route53QueryLog, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(route53querylogsResource, "status", route53QueryLog), &v1alpha1.Route53QueryLog{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Route53QueryLog), err
}

// Delete takes name of the route53QueryLog and deletes it. Returns an error if one occurs.
func (c *FakeRoute53QueryLogs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(route53querylogsResource, name), &v1alpha1.Route53QueryLog{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRoute53QueryLogs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(route53querylogsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.Route53QueryLogList{})
	return err
}

// Patch applies the patch and returns the patched route53QueryLog.
func (c *FakeRoute53QueryLogs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Route53QueryLog, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(route53querylogsResource, name, pt, data, subresources...), &v1alpha1.Route53QueryLog{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Route53QueryLog), err
}