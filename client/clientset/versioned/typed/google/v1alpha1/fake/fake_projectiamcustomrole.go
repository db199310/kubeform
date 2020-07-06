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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakeProjectIamCustomRoles implements ProjectIamCustomRoleInterface
type FakeProjectIamCustomRoles struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var projectiamcustomrolesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "projectiamcustomroles"}

var projectiamcustomrolesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "ProjectIamCustomRole"}

// Get takes name of the projectIamCustomRole, and returns the corresponding projectIamCustomRole object, and an error if there is any.
func (c *FakeProjectIamCustomRoles) Get(name string, options v1.GetOptions) (result *v1alpha1.ProjectIamCustomRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(projectiamcustomrolesResource, c.ns, name), &v1alpha1.ProjectIamCustomRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProjectIamCustomRole), err
}

// List takes label and field selectors, and returns the list of ProjectIamCustomRoles that match those selectors.
func (c *FakeProjectIamCustomRoles) List(opts v1.ListOptions) (result *v1alpha1.ProjectIamCustomRoleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(projectiamcustomrolesResource, projectiamcustomrolesKind, c.ns, opts), &v1alpha1.ProjectIamCustomRoleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ProjectIamCustomRoleList{ListMeta: obj.(*v1alpha1.ProjectIamCustomRoleList).ListMeta}
	for _, item := range obj.(*v1alpha1.ProjectIamCustomRoleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested projectIamCustomRoles.
func (c *FakeProjectIamCustomRoles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(projectiamcustomrolesResource, c.ns, opts))

}

// Create takes the representation of a projectIamCustomRole and creates it.  Returns the server's representation of the projectIamCustomRole, and an error, if there is any.
func (c *FakeProjectIamCustomRoles) Create(projectIamCustomRole *v1alpha1.ProjectIamCustomRole) (result *v1alpha1.ProjectIamCustomRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(projectiamcustomrolesResource, c.ns, projectIamCustomRole), &v1alpha1.ProjectIamCustomRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProjectIamCustomRole), err
}

// Update takes the representation of a projectIamCustomRole and updates it. Returns the server's representation of the projectIamCustomRole, and an error, if there is any.
func (c *FakeProjectIamCustomRoles) Update(projectIamCustomRole *v1alpha1.ProjectIamCustomRole) (result *v1alpha1.ProjectIamCustomRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(projectiamcustomrolesResource, c.ns, projectIamCustomRole), &v1alpha1.ProjectIamCustomRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProjectIamCustomRole), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeProjectIamCustomRoles) UpdateStatus(projectIamCustomRole *v1alpha1.ProjectIamCustomRole) (*v1alpha1.ProjectIamCustomRole, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(projectiamcustomrolesResource, "status", c.ns, projectIamCustomRole), &v1alpha1.ProjectIamCustomRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProjectIamCustomRole), err
}

// Delete takes name of the projectIamCustomRole and deletes it. Returns an error if one occurs.
func (c *FakeProjectIamCustomRoles) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(projectiamcustomrolesResource, c.ns, name), &v1alpha1.ProjectIamCustomRole{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProjectIamCustomRoles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(projectiamcustomrolesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ProjectIamCustomRoleList{})
	return err
}

// Patch applies the patch and returns the patched projectIamCustomRole.
func (c *FakeProjectIamCustomRoles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ProjectIamCustomRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(projectiamcustomrolesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ProjectIamCustomRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProjectIamCustomRole), err
}
