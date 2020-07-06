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
	v1alpha1 "kubeform.dev/kubeform/apis/linode/v1alpha1"
)

// FakeSshkeys implements SshkeyInterface
type FakeSshkeys struct {
	Fake *FakeLinodeV1alpha1
	ns   string
}

var sshkeysResource = schema.GroupVersionResource{Group: "linode.kubeform.com", Version: "v1alpha1", Resource: "sshkeys"}

var sshkeysKind = schema.GroupVersionKind{Group: "linode.kubeform.com", Version: "v1alpha1", Kind: "Sshkey"}

// Get takes name of the sshkey, and returns the corresponding sshkey object, and an error if there is any.
func (c *FakeSshkeys) Get(name string, options v1.GetOptions) (result *v1alpha1.Sshkey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sshkeysResource, c.ns, name), &v1alpha1.Sshkey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Sshkey), err
}

// List takes label and field selectors, and returns the list of Sshkeys that match those selectors.
func (c *FakeSshkeys) List(opts v1.ListOptions) (result *v1alpha1.SshkeyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sshkeysResource, sshkeysKind, c.ns, opts), &v1alpha1.SshkeyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SshkeyList{ListMeta: obj.(*v1alpha1.SshkeyList).ListMeta}
	for _, item := range obj.(*v1alpha1.SshkeyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sshkeys.
func (c *FakeSshkeys) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sshkeysResource, c.ns, opts))

}

// Create takes the representation of a sshkey and creates it.  Returns the server's representation of the sshkey, and an error, if there is any.
func (c *FakeSshkeys) Create(sshkey *v1alpha1.Sshkey) (result *v1alpha1.Sshkey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sshkeysResource, c.ns, sshkey), &v1alpha1.Sshkey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Sshkey), err
}

// Update takes the representation of a sshkey and updates it. Returns the server's representation of the sshkey, and an error, if there is any.
func (c *FakeSshkeys) Update(sshkey *v1alpha1.Sshkey) (result *v1alpha1.Sshkey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sshkeysResource, c.ns, sshkey), &v1alpha1.Sshkey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Sshkey), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSshkeys) UpdateStatus(sshkey *v1alpha1.Sshkey) (*v1alpha1.Sshkey, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sshkeysResource, "status", c.ns, sshkey), &v1alpha1.Sshkey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Sshkey), err
}

// Delete takes name of the sshkey and deletes it. Returns an error if one occurs.
func (c *FakeSshkeys) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sshkeysResource, c.ns, name), &v1alpha1.Sshkey{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSshkeys) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sshkeysResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SshkeyList{})
	return err
}

// Patch applies the patch and returns the patched sshkey.
func (c *FakeSshkeys) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Sshkey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sshkeysResource, c.ns, name, pt, data, subresources...), &v1alpha1.Sshkey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Sshkey), err
}
