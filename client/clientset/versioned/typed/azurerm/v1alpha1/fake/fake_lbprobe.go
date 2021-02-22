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

// FakeLbProbes implements LbProbeInterface
type FakeLbProbes struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var lbprobesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "lbprobes"}

var lbprobesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "LbProbe"}

// Get takes name of the lbProbe, and returns the corresponding lbProbe object, and an error if there is any.
func (c *FakeLbProbes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LbProbe, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(lbprobesResource, c.ns, name), &v1alpha1.LbProbe{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbProbe), err
}

// List takes label and field selectors, and returns the list of LbProbes that match those selectors.
func (c *FakeLbProbes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LbProbeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(lbprobesResource, lbprobesKind, c.ns, opts), &v1alpha1.LbProbeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LbProbeList{ListMeta: obj.(*v1alpha1.LbProbeList).ListMeta}
	for _, item := range obj.(*v1alpha1.LbProbeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested lbProbes.
func (c *FakeLbProbes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(lbprobesResource, c.ns, opts))

}

// Create takes the representation of a lbProbe and creates it.  Returns the server's representation of the lbProbe, and an error, if there is any.
func (c *FakeLbProbes) Create(ctx context.Context, lbProbe *v1alpha1.LbProbe, opts v1.CreateOptions) (result *v1alpha1.LbProbe, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(lbprobesResource, c.ns, lbProbe), &v1alpha1.LbProbe{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbProbe), err
}

// Update takes the representation of a lbProbe and updates it. Returns the server's representation of the lbProbe, and an error, if there is any.
func (c *FakeLbProbes) Update(ctx context.Context, lbProbe *v1alpha1.LbProbe, opts v1.UpdateOptions) (result *v1alpha1.LbProbe, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(lbprobesResource, c.ns, lbProbe), &v1alpha1.LbProbe{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbProbe), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLbProbes) UpdateStatus(ctx context.Context, lbProbe *v1alpha1.LbProbe, opts v1.UpdateOptions) (*v1alpha1.LbProbe, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(lbprobesResource, "status", c.ns, lbProbe), &v1alpha1.LbProbe{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbProbe), err
}

// Delete takes name of the lbProbe and deletes it. Returns an error if one occurs.
func (c *FakeLbProbes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(lbprobesResource, c.ns, name), &v1alpha1.LbProbe{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLbProbes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(lbprobesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LbProbeList{})
	return err
}

// Patch applies the patch and returns the patched lbProbe.
func (c *FakeLbProbes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LbProbe, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(lbprobesResource, c.ns, name, pt, data, subresources...), &v1alpha1.LbProbe{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbProbe), err
}
