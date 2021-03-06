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

// FakeCosmosdbGremlinGraphs implements CosmosdbGremlinGraphInterface
type FakeCosmosdbGremlinGraphs struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var cosmosdbgremlingraphsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "cosmosdbgremlingraphs"}

var cosmosdbgremlingraphsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "CosmosdbGremlinGraph"}

// Get takes name of the cosmosdbGremlinGraph, and returns the corresponding cosmosdbGremlinGraph object, and an error if there is any.
func (c *FakeCosmosdbGremlinGraphs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CosmosdbGremlinGraph, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cosmosdbgremlingraphsResource, c.ns, name), &v1alpha1.CosmosdbGremlinGraph{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CosmosdbGremlinGraph), err
}

// List takes label and field selectors, and returns the list of CosmosdbGremlinGraphs that match those selectors.
func (c *FakeCosmosdbGremlinGraphs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CosmosdbGremlinGraphList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cosmosdbgremlingraphsResource, cosmosdbgremlingraphsKind, c.ns, opts), &v1alpha1.CosmosdbGremlinGraphList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CosmosdbGremlinGraphList{ListMeta: obj.(*v1alpha1.CosmosdbGremlinGraphList).ListMeta}
	for _, item := range obj.(*v1alpha1.CosmosdbGremlinGraphList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cosmosdbGremlinGraphs.
func (c *FakeCosmosdbGremlinGraphs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cosmosdbgremlingraphsResource, c.ns, opts))

}

// Create takes the representation of a cosmosdbGremlinGraph and creates it.  Returns the server's representation of the cosmosdbGremlinGraph, and an error, if there is any.
func (c *FakeCosmosdbGremlinGraphs) Create(ctx context.Context, cosmosdbGremlinGraph *v1alpha1.CosmosdbGremlinGraph, opts v1.CreateOptions) (result *v1alpha1.CosmosdbGremlinGraph, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cosmosdbgremlingraphsResource, c.ns, cosmosdbGremlinGraph), &v1alpha1.CosmosdbGremlinGraph{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CosmosdbGremlinGraph), err
}

// Update takes the representation of a cosmosdbGremlinGraph and updates it. Returns the server's representation of the cosmosdbGremlinGraph, and an error, if there is any.
func (c *FakeCosmosdbGremlinGraphs) Update(ctx context.Context, cosmosdbGremlinGraph *v1alpha1.CosmosdbGremlinGraph, opts v1.UpdateOptions) (result *v1alpha1.CosmosdbGremlinGraph, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cosmosdbgremlingraphsResource, c.ns, cosmosdbGremlinGraph), &v1alpha1.CosmosdbGremlinGraph{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CosmosdbGremlinGraph), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCosmosdbGremlinGraphs) UpdateStatus(ctx context.Context, cosmosdbGremlinGraph *v1alpha1.CosmosdbGremlinGraph, opts v1.UpdateOptions) (*v1alpha1.CosmosdbGremlinGraph, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cosmosdbgremlingraphsResource, "status", c.ns, cosmosdbGremlinGraph), &v1alpha1.CosmosdbGremlinGraph{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CosmosdbGremlinGraph), err
}

// Delete takes name of the cosmosdbGremlinGraph and deletes it. Returns an error if one occurs.
func (c *FakeCosmosdbGremlinGraphs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cosmosdbgremlingraphsResource, c.ns, name), &v1alpha1.CosmosdbGremlinGraph{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCosmosdbGremlinGraphs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cosmosdbgremlingraphsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CosmosdbGremlinGraphList{})
	return err
}

// Patch applies the patch and returns the patched cosmosdbGremlinGraph.
func (c *FakeCosmosdbGremlinGraphs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CosmosdbGremlinGraph, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cosmosdbgremlingraphsResource, c.ns, name, pt, data, subresources...), &v1alpha1.CosmosdbGremlinGraph{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CosmosdbGremlinGraph), err
}
