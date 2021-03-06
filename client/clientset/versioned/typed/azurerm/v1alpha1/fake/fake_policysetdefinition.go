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

// FakePolicySetDefinitions implements PolicySetDefinitionInterface
type FakePolicySetDefinitions struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var policysetdefinitionsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "policysetdefinitions"}

var policysetdefinitionsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "PolicySetDefinition"}

// Get takes name of the policySetDefinition, and returns the corresponding policySetDefinition object, and an error if there is any.
func (c *FakePolicySetDefinitions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PolicySetDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(policysetdefinitionsResource, c.ns, name), &v1alpha1.PolicySetDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicySetDefinition), err
}

// List takes label and field selectors, and returns the list of PolicySetDefinitions that match those selectors.
func (c *FakePolicySetDefinitions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PolicySetDefinitionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(policysetdefinitionsResource, policysetdefinitionsKind, c.ns, opts), &v1alpha1.PolicySetDefinitionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PolicySetDefinitionList{ListMeta: obj.(*v1alpha1.PolicySetDefinitionList).ListMeta}
	for _, item := range obj.(*v1alpha1.PolicySetDefinitionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested policySetDefinitions.
func (c *FakePolicySetDefinitions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(policysetdefinitionsResource, c.ns, opts))

}

// Create takes the representation of a policySetDefinition and creates it.  Returns the server's representation of the policySetDefinition, and an error, if there is any.
func (c *FakePolicySetDefinitions) Create(ctx context.Context, policySetDefinition *v1alpha1.PolicySetDefinition, opts v1.CreateOptions) (result *v1alpha1.PolicySetDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(policysetdefinitionsResource, c.ns, policySetDefinition), &v1alpha1.PolicySetDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicySetDefinition), err
}

// Update takes the representation of a policySetDefinition and updates it. Returns the server's representation of the policySetDefinition, and an error, if there is any.
func (c *FakePolicySetDefinitions) Update(ctx context.Context, policySetDefinition *v1alpha1.PolicySetDefinition, opts v1.UpdateOptions) (result *v1alpha1.PolicySetDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(policysetdefinitionsResource, c.ns, policySetDefinition), &v1alpha1.PolicySetDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicySetDefinition), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePolicySetDefinitions) UpdateStatus(ctx context.Context, policySetDefinition *v1alpha1.PolicySetDefinition, opts v1.UpdateOptions) (*v1alpha1.PolicySetDefinition, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(policysetdefinitionsResource, "status", c.ns, policySetDefinition), &v1alpha1.PolicySetDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicySetDefinition), err
}

// Delete takes name of the policySetDefinition and deletes it. Returns an error if one occurs.
func (c *FakePolicySetDefinitions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(policysetdefinitionsResource, c.ns, name), &v1alpha1.PolicySetDefinition{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePolicySetDefinitions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(policysetdefinitionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PolicySetDefinitionList{})
	return err
}

// Patch applies the patch and returns the patched policySetDefinition.
func (c *FakePolicySetDefinitions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PolicySetDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(policysetdefinitionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PolicySetDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicySetDefinition), err
}
