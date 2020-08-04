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

// FakeLogicAppTriggerCustoms implements LogicAppTriggerCustomInterface
type FakeLogicAppTriggerCustoms struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var logicapptriggercustomsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "logicapptriggercustoms"}

var logicapptriggercustomsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "LogicAppTriggerCustom"}

// Get takes name of the logicAppTriggerCustom, and returns the corresponding logicAppTriggerCustom object, and an error if there is any.
func (c *FakeLogicAppTriggerCustoms) Get(name string, options v1.GetOptions) (result *v1alpha1.LogicAppTriggerCustom, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(logicapptriggercustomsResource, c.ns, name), &v1alpha1.LogicAppTriggerCustom{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogicAppTriggerCustom), err
}

// List takes label and field selectors, and returns the list of LogicAppTriggerCustoms that match those selectors.
func (c *FakeLogicAppTriggerCustoms) List(opts v1.ListOptions) (result *v1alpha1.LogicAppTriggerCustomList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(logicapptriggercustomsResource, logicapptriggercustomsKind, c.ns, opts), &v1alpha1.LogicAppTriggerCustomList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LogicAppTriggerCustomList{ListMeta: obj.(*v1alpha1.LogicAppTriggerCustomList).ListMeta}
	for _, item := range obj.(*v1alpha1.LogicAppTriggerCustomList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested logicAppTriggerCustoms.
func (c *FakeLogicAppTriggerCustoms) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(logicapptriggercustomsResource, c.ns, opts))

}

// Create takes the representation of a logicAppTriggerCustom and creates it.  Returns the server's representation of the logicAppTriggerCustom, and an error, if there is any.
func (c *FakeLogicAppTriggerCustoms) Create(logicAppTriggerCustom *v1alpha1.LogicAppTriggerCustom) (result *v1alpha1.LogicAppTriggerCustom, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(logicapptriggercustomsResource, c.ns, logicAppTriggerCustom), &v1alpha1.LogicAppTriggerCustom{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogicAppTriggerCustom), err
}

// Update takes the representation of a logicAppTriggerCustom and updates it. Returns the server's representation of the logicAppTriggerCustom, and an error, if there is any.
func (c *FakeLogicAppTriggerCustoms) Update(logicAppTriggerCustom *v1alpha1.LogicAppTriggerCustom) (result *v1alpha1.LogicAppTriggerCustom, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(logicapptriggercustomsResource, c.ns, logicAppTriggerCustom), &v1alpha1.LogicAppTriggerCustom{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogicAppTriggerCustom), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLogicAppTriggerCustoms) UpdateStatus(logicAppTriggerCustom *v1alpha1.LogicAppTriggerCustom) (*v1alpha1.LogicAppTriggerCustom, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(logicapptriggercustomsResource, "status", c.ns, logicAppTriggerCustom), &v1alpha1.LogicAppTriggerCustom{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogicAppTriggerCustom), err
}

// Delete takes name of the logicAppTriggerCustom and deletes it. Returns an error if one occurs.
func (c *FakeLogicAppTriggerCustoms) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(logicapptriggercustomsResource, c.ns, name), &v1alpha1.LogicAppTriggerCustom{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLogicAppTriggerCustoms) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(logicapptriggercustomsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.LogicAppTriggerCustomList{})
	return err
}

// Patch applies the patch and returns the patched logicAppTriggerCustom.
func (c *FakeLogicAppTriggerCustoms) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LogicAppTriggerCustom, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(logicapptriggercustomsResource, c.ns, name, pt, data, subresources...), &v1alpha1.LogicAppTriggerCustom{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogicAppTriggerCustom), err
}
