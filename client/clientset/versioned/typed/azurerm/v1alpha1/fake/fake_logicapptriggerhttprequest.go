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

// FakeLogicAppTriggerHTTPRequests implements LogicAppTriggerHTTPRequestInterface
type FakeLogicAppTriggerHTTPRequests struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var logicapptriggerhttprequestsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "logicapptriggerhttprequests"}

var logicapptriggerhttprequestsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "LogicAppTriggerHTTPRequest"}

// Get takes name of the logicAppTriggerHTTPRequest, and returns the corresponding logicAppTriggerHTTPRequest object, and an error if there is any.
func (c *FakeLogicAppTriggerHTTPRequests) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LogicAppTriggerHTTPRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(logicapptriggerhttprequestsResource, c.ns, name), &v1alpha1.LogicAppTriggerHTTPRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogicAppTriggerHTTPRequest), err
}

// List takes label and field selectors, and returns the list of LogicAppTriggerHTTPRequests that match those selectors.
func (c *FakeLogicAppTriggerHTTPRequests) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LogicAppTriggerHTTPRequestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(logicapptriggerhttprequestsResource, logicapptriggerhttprequestsKind, c.ns, opts), &v1alpha1.LogicAppTriggerHTTPRequestList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LogicAppTriggerHTTPRequestList{ListMeta: obj.(*v1alpha1.LogicAppTriggerHTTPRequestList).ListMeta}
	for _, item := range obj.(*v1alpha1.LogicAppTriggerHTTPRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested logicAppTriggerHTTPRequests.
func (c *FakeLogicAppTriggerHTTPRequests) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(logicapptriggerhttprequestsResource, c.ns, opts))

}

// Create takes the representation of a logicAppTriggerHTTPRequest and creates it.  Returns the server's representation of the logicAppTriggerHTTPRequest, and an error, if there is any.
func (c *FakeLogicAppTriggerHTTPRequests) Create(ctx context.Context, logicAppTriggerHTTPRequest *v1alpha1.LogicAppTriggerHTTPRequest, opts v1.CreateOptions) (result *v1alpha1.LogicAppTriggerHTTPRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(logicapptriggerhttprequestsResource, c.ns, logicAppTriggerHTTPRequest), &v1alpha1.LogicAppTriggerHTTPRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogicAppTriggerHTTPRequest), err
}

// Update takes the representation of a logicAppTriggerHTTPRequest and updates it. Returns the server's representation of the logicAppTriggerHTTPRequest, and an error, if there is any.
func (c *FakeLogicAppTriggerHTTPRequests) Update(ctx context.Context, logicAppTriggerHTTPRequest *v1alpha1.LogicAppTriggerHTTPRequest, opts v1.UpdateOptions) (result *v1alpha1.LogicAppTriggerHTTPRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(logicapptriggerhttprequestsResource, c.ns, logicAppTriggerHTTPRequest), &v1alpha1.LogicAppTriggerHTTPRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogicAppTriggerHTTPRequest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLogicAppTriggerHTTPRequests) UpdateStatus(ctx context.Context, logicAppTriggerHTTPRequest *v1alpha1.LogicAppTriggerHTTPRequest, opts v1.UpdateOptions) (*v1alpha1.LogicAppTriggerHTTPRequest, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(logicapptriggerhttprequestsResource, "status", c.ns, logicAppTriggerHTTPRequest), &v1alpha1.LogicAppTriggerHTTPRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogicAppTriggerHTTPRequest), err
}

// Delete takes name of the logicAppTriggerHTTPRequest and deletes it. Returns an error if one occurs.
func (c *FakeLogicAppTriggerHTTPRequests) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(logicapptriggerhttprequestsResource, c.ns, name), &v1alpha1.LogicAppTriggerHTTPRequest{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLogicAppTriggerHTTPRequests) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(logicapptriggerhttprequestsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LogicAppTriggerHTTPRequestList{})
	return err
}

// Patch applies the patch and returns the patched logicAppTriggerHTTPRequest.
func (c *FakeLogicAppTriggerHTTPRequests) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LogicAppTriggerHTTPRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(logicapptriggerhttprequestsResource, c.ns, name, pt, data, subresources...), &v1alpha1.LogicAppTriggerHTTPRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogicAppTriggerHTTPRequest), err
}
