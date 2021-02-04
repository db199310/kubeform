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

type VirtualWAN struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualWANSpec   `json:"spec,omitempty"`
	Status            VirtualWANStatus `json:"status,omitempty"`
}

type VirtualWANSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	RemoteBackend *base.Backend `json:"remoteBackend,omitempty" tf:"-"`

	// +optional
	AllowBranchToBranchTraffic bool `json:"allowBranchToBranchTraffic,omitempty" tf:"allow_branch_to_branch_traffic,omitempty"`
	// +optional
	AllowVnetToVnetTraffic bool `json:"allowVnetToVnetTraffic,omitempty" tf:"allow_vnet_to_vnet_traffic,omitempty"`
	// +optional
	DisableVPNEncryption bool   `json:"disableVPNEncryption,omitempty" tf:"disable_vpn_encryption,omitempty"`
	Location             string `json:"location" tf:"location"`
	Name                 string `json:"name" tf:"name"`
	// +optional
	Office365LocalBreakoutCategory string `json:"office365LocalBreakoutCategory,omitempty" tf:"office365_local_breakout_category,omitempty"`
	ResourceGroupName              string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
}

type VirtualWANStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *VirtualWANSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VirtualWANList is a list of VirtualWANs
type VirtualWANList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VirtualWAN CRD objects
	Items []VirtualWAN `json:"items,omitempty"`
}
