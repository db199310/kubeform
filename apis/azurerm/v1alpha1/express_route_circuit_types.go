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

type ExpressRouteCircuit struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExpressRouteCircuitSpec   `json:"spec,omitempty"`
	Status            ExpressRouteCircuitStatus `json:"status,omitempty"`
}

type ExpressRouteCircuitSpecSku struct {
	Family string `json:"family" tf:"family"`
	Tier   string `json:"tier" tf:"tier"`
}

type ExpressRouteCircuitSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	AllowClassicOperations bool   `json:"allowClassicOperations,omitempty" tf:"allow_classic_operations,omitempty"`
	BandwidthInMbps        int64  `json:"bandwidthInMbps" tf:"bandwidth_in_mbps"`
	Location               string `json:"location" tf:"location"`
	Name                   string `json:"name" tf:"name"`
	PeeringLocation        string `json:"peeringLocation" tf:"peering_location"`
	ResourceGroupName      string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	ServiceKey          string `json:"-" sensitive:"true" tf:"service_key,omitempty"`
	ServiceProviderName string `json:"serviceProviderName" tf:"service_provider_name"`
	// +optional
	ServiceProviderProvisioningState string `json:"serviceProviderProvisioningState,omitempty" tf:"service_provider_provisioning_state,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Sku []ExpressRouteCircuitSpecSku `json:"sku" tf:"sku"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ExpressRouteCircuitStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ExpressRouteCircuitSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
	// +optional
	TerraFormLogs *base.TerraFormLogs `json:"terraformLogs,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ExpressRouteCircuitList is a list of ExpressRouteCircuits
type ExpressRouteCircuitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ExpressRouteCircuit CRD objects
	Items []ExpressRouteCircuit `json:"items,omitempty"`
}
