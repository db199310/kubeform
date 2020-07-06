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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeEc2TransitGatewayRouteTables implements Ec2TransitGatewayRouteTableInterface
type FakeEc2TransitGatewayRouteTables struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var ec2transitgatewayroutetablesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "ec2transitgatewayroutetables"}

var ec2transitgatewayroutetablesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "Ec2TransitGatewayRouteTable"}

// Get takes name of the ec2TransitGatewayRouteTable, and returns the corresponding ec2TransitGatewayRouteTable object, and an error if there is any.
func (c *FakeEc2TransitGatewayRouteTables) Get(name string, options v1.GetOptions) (result *v1alpha1.Ec2TransitGatewayRouteTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ec2transitgatewayroutetablesResource, c.ns, name), &v1alpha1.Ec2TransitGatewayRouteTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ec2TransitGatewayRouteTable), err
}

// List takes label and field selectors, and returns the list of Ec2TransitGatewayRouteTables that match those selectors.
func (c *FakeEc2TransitGatewayRouteTables) List(opts v1.ListOptions) (result *v1alpha1.Ec2TransitGatewayRouteTableList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ec2transitgatewayroutetablesResource, ec2transitgatewayroutetablesKind, c.ns, opts), &v1alpha1.Ec2TransitGatewayRouteTableList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.Ec2TransitGatewayRouteTableList{ListMeta: obj.(*v1alpha1.Ec2TransitGatewayRouteTableList).ListMeta}
	for _, item := range obj.(*v1alpha1.Ec2TransitGatewayRouteTableList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ec2TransitGatewayRouteTables.
func (c *FakeEc2TransitGatewayRouteTables) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ec2transitgatewayroutetablesResource, c.ns, opts))

}

// Create takes the representation of a ec2TransitGatewayRouteTable and creates it.  Returns the server's representation of the ec2TransitGatewayRouteTable, and an error, if there is any.
func (c *FakeEc2TransitGatewayRouteTables) Create(ec2TransitGatewayRouteTable *v1alpha1.Ec2TransitGatewayRouteTable) (result *v1alpha1.Ec2TransitGatewayRouteTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ec2transitgatewayroutetablesResource, c.ns, ec2TransitGatewayRouteTable), &v1alpha1.Ec2TransitGatewayRouteTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ec2TransitGatewayRouteTable), err
}

// Update takes the representation of a ec2TransitGatewayRouteTable and updates it. Returns the server's representation of the ec2TransitGatewayRouteTable, and an error, if there is any.
func (c *FakeEc2TransitGatewayRouteTables) Update(ec2TransitGatewayRouteTable *v1alpha1.Ec2TransitGatewayRouteTable) (result *v1alpha1.Ec2TransitGatewayRouteTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ec2transitgatewayroutetablesResource, c.ns, ec2TransitGatewayRouteTable), &v1alpha1.Ec2TransitGatewayRouteTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ec2TransitGatewayRouteTable), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEc2TransitGatewayRouteTables) UpdateStatus(ec2TransitGatewayRouteTable *v1alpha1.Ec2TransitGatewayRouteTable) (*v1alpha1.Ec2TransitGatewayRouteTable, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ec2transitgatewayroutetablesResource, "status", c.ns, ec2TransitGatewayRouteTable), &v1alpha1.Ec2TransitGatewayRouteTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ec2TransitGatewayRouteTable), err
}

// Delete takes name of the ec2TransitGatewayRouteTable and deletes it. Returns an error if one occurs.
func (c *FakeEc2TransitGatewayRouteTables) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ec2transitgatewayroutetablesResource, c.ns, name), &v1alpha1.Ec2TransitGatewayRouteTable{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEc2TransitGatewayRouteTables) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ec2transitgatewayroutetablesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.Ec2TransitGatewayRouteTableList{})
	return err
}

// Patch applies the patch and returns the patched ec2TransitGatewayRouteTable.
func (c *FakeEc2TransitGatewayRouteTables) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Ec2TransitGatewayRouteTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ec2transitgatewayroutetablesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Ec2TransitGatewayRouteTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ec2TransitGatewayRouteTable), err
}
