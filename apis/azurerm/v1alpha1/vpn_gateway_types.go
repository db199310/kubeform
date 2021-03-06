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

type VpnGateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpnGatewaySpec   `json:"spec,omitempty"`
	Status            VpnGatewayStatus `json:"status,omitempty"`
}

type VpnGatewaySpecBgpSettings struct {
	Asn int64 `json:"asn" tf:"asn"`
	// +optional
	BgpPeeringAddress string `json:"bgpPeeringAddress,omitempty" tf:"bgp_peering_address,omitempty"`
	PeerWeight        int64  `json:"peerWeight" tf:"peer_weight"`
}

type VpnGatewaySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	BgpSettings       []VpnGatewaySpecBgpSettings `json:"bgpSettings,omitempty" tf:"bgp_settings,omitempty"`
	Location          string                      `json:"location" tf:"location"`
	Name              string                      `json:"name" tf:"name"`
	ResourceGroupName string                      `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	ScaleUnit int64 `json:"scaleUnit,omitempty" tf:"scale_unit,omitempty"`
	// +optional
	Tags         map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	VirtualHubID string            `json:"virtualHubID" tf:"virtual_hub_id"`
}

type VpnGatewayStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *VpnGatewaySpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
	// +optional
	TerraformErrors []string `json:"terraformErrors,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VpnGatewayList is a list of VpnGateways
type VpnGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpnGateway CRD objects
	Items []VpnGateway `json:"items,omitempty"`
}
