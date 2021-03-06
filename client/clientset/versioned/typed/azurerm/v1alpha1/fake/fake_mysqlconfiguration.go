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

// FakeMysqlConfigurations implements MysqlConfigurationInterface
type FakeMysqlConfigurations struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var mysqlconfigurationsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "mysqlconfigurations"}

var mysqlconfigurationsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "MysqlConfiguration"}

// Get takes name of the mysqlConfiguration, and returns the corresponding mysqlConfiguration object, and an error if there is any.
func (c *FakeMysqlConfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MysqlConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mysqlconfigurationsResource, c.ns, name), &v1alpha1.MysqlConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlConfiguration), err
}

// List takes label and field selectors, and returns the list of MysqlConfigurations that match those selectors.
func (c *FakeMysqlConfigurations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MysqlConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mysqlconfigurationsResource, mysqlconfigurationsKind, c.ns, opts), &v1alpha1.MysqlConfigurationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MysqlConfigurationList{ListMeta: obj.(*v1alpha1.MysqlConfigurationList).ListMeta}
	for _, item := range obj.(*v1alpha1.MysqlConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mysqlConfigurations.
func (c *FakeMysqlConfigurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mysqlconfigurationsResource, c.ns, opts))

}

// Create takes the representation of a mysqlConfiguration and creates it.  Returns the server's representation of the mysqlConfiguration, and an error, if there is any.
func (c *FakeMysqlConfigurations) Create(ctx context.Context, mysqlConfiguration *v1alpha1.MysqlConfiguration, opts v1.CreateOptions) (result *v1alpha1.MysqlConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mysqlconfigurationsResource, c.ns, mysqlConfiguration), &v1alpha1.MysqlConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlConfiguration), err
}

// Update takes the representation of a mysqlConfiguration and updates it. Returns the server's representation of the mysqlConfiguration, and an error, if there is any.
func (c *FakeMysqlConfigurations) Update(ctx context.Context, mysqlConfiguration *v1alpha1.MysqlConfiguration, opts v1.UpdateOptions) (result *v1alpha1.MysqlConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mysqlconfigurationsResource, c.ns, mysqlConfiguration), &v1alpha1.MysqlConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlConfiguration), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMysqlConfigurations) UpdateStatus(ctx context.Context, mysqlConfiguration *v1alpha1.MysqlConfiguration, opts v1.UpdateOptions) (*v1alpha1.MysqlConfiguration, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(mysqlconfigurationsResource, "status", c.ns, mysqlConfiguration), &v1alpha1.MysqlConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlConfiguration), err
}

// Delete takes name of the mysqlConfiguration and deletes it. Returns an error if one occurs.
func (c *FakeMysqlConfigurations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(mysqlconfigurationsResource, c.ns, name), &v1alpha1.MysqlConfiguration{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMysqlConfigurations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mysqlconfigurationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MysqlConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched mysqlConfiguration.
func (c *FakeMysqlConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MysqlConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mysqlconfigurationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.MysqlConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlConfiguration), err
}
