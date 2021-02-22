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

// FakeSqlServers implements SqlServerInterface
type FakeSqlServers struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var sqlserversResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "sqlservers"}

var sqlserversKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "SqlServer"}

// Get takes name of the sqlServer, and returns the corresponding sqlServer object, and an error if there is any.
func (c *FakeSqlServers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SqlServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sqlserversResource, c.ns, name), &v1alpha1.SqlServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlServer), err
}

// List takes label and field selectors, and returns the list of SqlServers that match those selectors.
func (c *FakeSqlServers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SqlServerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sqlserversResource, sqlserversKind, c.ns, opts), &v1alpha1.SqlServerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SqlServerList{ListMeta: obj.(*v1alpha1.SqlServerList).ListMeta}
	for _, item := range obj.(*v1alpha1.SqlServerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sqlServers.
func (c *FakeSqlServers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sqlserversResource, c.ns, opts))

}

// Create takes the representation of a sqlServer and creates it.  Returns the server's representation of the sqlServer, and an error, if there is any.
func (c *FakeSqlServers) Create(ctx context.Context, sqlServer *v1alpha1.SqlServer, opts v1.CreateOptions) (result *v1alpha1.SqlServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sqlserversResource, c.ns, sqlServer), &v1alpha1.SqlServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlServer), err
}

// Update takes the representation of a sqlServer and updates it. Returns the server's representation of the sqlServer, and an error, if there is any.
func (c *FakeSqlServers) Update(ctx context.Context, sqlServer *v1alpha1.SqlServer, opts v1.UpdateOptions) (result *v1alpha1.SqlServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sqlserversResource, c.ns, sqlServer), &v1alpha1.SqlServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlServer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSqlServers) UpdateStatus(ctx context.Context, sqlServer *v1alpha1.SqlServer, opts v1.UpdateOptions) (*v1alpha1.SqlServer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sqlserversResource, "status", c.ns, sqlServer), &v1alpha1.SqlServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlServer), err
}

// Delete takes name of the sqlServer and deletes it. Returns an error if one occurs.
func (c *FakeSqlServers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sqlserversResource, c.ns, name), &v1alpha1.SqlServer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSqlServers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sqlserversResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SqlServerList{})
	return err
}

// Patch applies the patch and returns the patched sqlServer.
func (c *FakeSqlServers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SqlServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sqlserversResource, c.ns, name, pt, data, subresources...), &v1alpha1.SqlServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlServer), err
}
