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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:storageversion

type RouteTable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteTableSpec   `json:"spec,omitempty"`
	Status            RouteTableStatus `json:"status,omitempty"`
}

type RouteTableSpecRoute struct {
	AddressPrefix string `json:"addressPrefix" tf:"address_prefix"`
	Name          string `json:"name" tf:"name"`
	// +optional
	NextHopInIPAddress string `json:"nextHopInIPAddress,omitempty" tf:"next_hop_in_ip_address,omitempty"`
	NextHopType        string `json:"nextHopType" tf:"next_hop_type"`
}

type RouteTableSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	DisableBGPRoutePropagation bool   `json:"disableBGPRoutePropagation,omitempty" tf:"disable_bgp_route_propagation,omitempty"`
	Location                   string `json:"location" tf:"location"`
	Name                       string `json:"name" tf:"name"`
	ResourceGroupName          string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Route []RouteTableSpecRoute `json:"route,omitempty" tf:"route,omitempty"`
	// +optional
	Subnets []string `json:"subnets,omitempty" tf:"subnets,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RouteTableStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *RouteTableSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RouteTableList is a list of RouteTables
type RouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RouteTable CRD objects
	Items []RouteTable `json:"items,omitempty"`
}
