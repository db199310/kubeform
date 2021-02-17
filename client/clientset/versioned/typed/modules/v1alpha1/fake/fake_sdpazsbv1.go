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
	v1alpha1 "kubeform.dev/kubeform/apis/modules/v1alpha1"
)

// FakeSDPAzsbv1s implements SDPAzsbv1Interface
type FakeSDPAzsbv1s struct {
	Fake *FakeModulesV1alpha1
	ns   string
}

var sdpazsbv1sResource = schema.GroupVersionResource{Group: "modules.kubeform.com", Version: "v1alpha1", Resource: "sdpazsbv1s"}

var sdpazsbv1sKind = schema.GroupVersionKind{Group: "modules.kubeform.com", Version: "v1alpha1", Kind: "SDPAzsbv1"}

// Get takes name of the sDPAzsbv1, and returns the corresponding sDPAzsbv1 object, and an error if there is any.
func (c *FakeSDPAzsbv1s) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SDPAzsbv1, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sdpazsbv1sResource, c.ns, name), &v1alpha1.SDPAzsbv1{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SDPAzsbv1), err
}

// List takes label and field selectors, and returns the list of SDPAzsbv1s that match those selectors.
func (c *FakeSDPAzsbv1s) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SDPAzsbv1List, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sdpazsbv1sResource, sdpazsbv1sKind, c.ns, opts), &v1alpha1.SDPAzsbv1List{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SDPAzsbv1List{ListMeta: obj.(*v1alpha1.SDPAzsbv1List).ListMeta}
	for _, item := range obj.(*v1alpha1.SDPAzsbv1List).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sDPAzsbv1s.
func (c *FakeSDPAzsbv1s) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sdpazsbv1sResource, c.ns, opts))

}

// Create takes the representation of a sDPAzsbv1 and creates it.  Returns the server's representation of the sDPAzsbv1, and an error, if there is any.
func (c *FakeSDPAzsbv1s) Create(ctx context.Context, sDPAzsbv1 *v1alpha1.SDPAzsbv1, opts v1.CreateOptions) (result *v1alpha1.SDPAzsbv1, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sdpazsbv1sResource, c.ns, sDPAzsbv1), &v1alpha1.SDPAzsbv1{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SDPAzsbv1), err
}

// Update takes the representation of a sDPAzsbv1 and updates it. Returns the server's representation of the sDPAzsbv1, and an error, if there is any.
func (c *FakeSDPAzsbv1s) Update(ctx context.Context, sDPAzsbv1 *v1alpha1.SDPAzsbv1, opts v1.UpdateOptions) (result *v1alpha1.SDPAzsbv1, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sdpazsbv1sResource, c.ns, sDPAzsbv1), &v1alpha1.SDPAzsbv1{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SDPAzsbv1), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSDPAzsbv1s) UpdateStatus(ctx context.Context, sDPAzsbv1 *v1alpha1.SDPAzsbv1, opts v1.UpdateOptions) (*v1alpha1.SDPAzsbv1, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sdpazsbv1sResource, "status", c.ns, sDPAzsbv1), &v1alpha1.SDPAzsbv1{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SDPAzsbv1), err
}

// Delete takes name of the sDPAzsbv1 and deletes it. Returns an error if one occurs.
func (c *FakeSDPAzsbv1s) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sdpazsbv1sResource, c.ns, name), &v1alpha1.SDPAzsbv1{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSDPAzsbv1s) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sdpazsbv1sResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SDPAzsbv1List{})
	return err
}

// Patch applies the patch and returns the patched sDPAzsbv1.
func (c *FakeSDPAzsbv1s) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SDPAzsbv1, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sdpazsbv1sResource, c.ns, name, pt, data, subresources...), &v1alpha1.SDPAzsbv1{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SDPAzsbv1), err
}
