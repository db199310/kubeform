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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakeKubernetesClusterNodePools implements KubernetesClusterNodePoolInterface
type FakeKubernetesClusterNodePools struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var kubernetesclusternodepoolsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "kubernetesclusternodepools"}

var kubernetesclusternodepoolsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "KubernetesClusterNodePool"}

// Get takes name of the kubernetesClusterNodePool, and returns the corresponding kubernetesClusterNodePool object, and an error if there is any.
func (c *FakeKubernetesClusterNodePools) Get(name string, options v1.GetOptions) (result *v1alpha1.KubernetesClusterNodePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kubernetesclusternodepoolsResource, c.ns, name), &v1alpha1.KubernetesClusterNodePool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubernetesClusterNodePool), err
}

// List takes label and field selectors, and returns the list of KubernetesClusterNodePools that match those selectors.
func (c *FakeKubernetesClusterNodePools) List(opts v1.ListOptions) (result *v1alpha1.KubernetesClusterNodePoolList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kubernetesclusternodepoolsResource, kubernetesclusternodepoolsKind, c.ns, opts), &v1alpha1.KubernetesClusterNodePoolList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KubernetesClusterNodePoolList{ListMeta: obj.(*v1alpha1.KubernetesClusterNodePoolList).ListMeta}
	for _, item := range obj.(*v1alpha1.KubernetesClusterNodePoolList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kubernetesClusterNodePools.
func (c *FakeKubernetesClusterNodePools) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kubernetesclusternodepoolsResource, c.ns, opts))

}

// Create takes the representation of a kubernetesClusterNodePool and creates it.  Returns the server's representation of the kubernetesClusterNodePool, and an error, if there is any.
func (c *FakeKubernetesClusterNodePools) Create(kubernetesClusterNodePool *v1alpha1.KubernetesClusterNodePool) (result *v1alpha1.KubernetesClusterNodePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kubernetesclusternodepoolsResource, c.ns, kubernetesClusterNodePool), &v1alpha1.KubernetesClusterNodePool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubernetesClusterNodePool), err
}

// Update takes the representation of a kubernetesClusterNodePool and updates it. Returns the server's representation of the kubernetesClusterNodePool, and an error, if there is any.
func (c *FakeKubernetesClusterNodePools) Update(kubernetesClusterNodePool *v1alpha1.KubernetesClusterNodePool) (result *v1alpha1.KubernetesClusterNodePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kubernetesclusternodepoolsResource, c.ns, kubernetesClusterNodePool), &v1alpha1.KubernetesClusterNodePool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubernetesClusterNodePool), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKubernetesClusterNodePools) UpdateStatus(kubernetesClusterNodePool *v1alpha1.KubernetesClusterNodePool) (*v1alpha1.KubernetesClusterNodePool, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(kubernetesclusternodepoolsResource, "status", c.ns, kubernetesClusterNodePool), &v1alpha1.KubernetesClusterNodePool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubernetesClusterNodePool), err
}

// Delete takes name of the kubernetesClusterNodePool and deletes it. Returns an error if one occurs.
func (c *FakeKubernetesClusterNodePools) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(kubernetesclusternodepoolsResource, c.ns, name), &v1alpha1.KubernetesClusterNodePool{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKubernetesClusterNodePools) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kubernetesclusternodepoolsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.KubernetesClusterNodePoolList{})
	return err
}

// Patch applies the patch and returns the patched kubernetesClusterNodePool.
func (c *FakeKubernetesClusterNodePools) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KubernetesClusterNodePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kubernetesclusternodepoolsResource, c.ns, name, pt, data, subresources...), &v1alpha1.KubernetesClusterNodePool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubernetesClusterNodePool), err
}
