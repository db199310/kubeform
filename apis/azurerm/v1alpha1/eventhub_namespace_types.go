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

type EventhubNamespace_ struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventhubNamespace_Spec   `json:"spec,omitempty"`
	Status            EventhubNamespace_Status `json:"status,omitempty"`
}

type EventhubNamespace_SpecNetworkRulesetsIpRule struct {
	// +optional
	Action string `json:"action,omitempty" tf:"action,omitempty"`
	IpMask string `json:"ipMask" tf:"ip_mask"`
}

type EventhubNamespace_SpecNetworkRulesetsVirtualNetworkRule struct {
	// +optional
	IgnoreMissingVirtualNetworkServiceEndpoint bool   `json:"ignoreMissingVirtualNetworkServiceEndpoint,omitempty" tf:"ignore_missing_virtual_network_service_endpoint,omitempty"`
	SubnetID                                   string `json:"subnetID" tf:"subnet_id"`
}

type EventhubNamespace_SpecNetworkRulesets struct {
	DefaultAction string `json:"defaultAction" tf:"default_action"`
	// +optional
	// +kubebuilder:validation:MaxItems=128
	IpRule []EventhubNamespace_SpecNetworkRulesetsIpRule `json:"ipRule,omitempty" tf:"ip_rule,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=128
	VirtualNetworkRule []EventhubNamespace_SpecNetworkRulesetsVirtualNetworkRule `json:"virtualNetworkRule,omitempty" tf:"virtual_network_rule,omitempty"`
}

type EventhubNamespace_Spec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	AutoInflateEnabled bool `json:"autoInflateEnabled,omitempty" tf:"auto_inflate_enabled,omitempty"`
	// +optional
	Capacity int64 `json:"capacity,omitempty" tf:"capacity,omitempty"`
	// +optional
	DefaultPrimaryConnectionString string `json:"-" sensitive:"true" tf:"default_primary_connection_string,omitempty"`
	// +optional
	DefaultPrimaryConnectionStringAlias string `json:"-" sensitive:"true" tf:"default_primary_connection_string_alias,omitempty"`
	// +optional
	DefaultPrimaryKey string `json:"-" sensitive:"true" tf:"default_primary_key,omitempty"`
	// +optional
	DefaultSecondaryConnectionString string `json:"-" sensitive:"true" tf:"default_secondary_connection_string,omitempty"`
	// +optional
	DefaultSecondaryConnectionStringAlias string `json:"-" sensitive:"true" tf:"default_secondary_connection_string_alias,omitempty"`
	// +optional
	DefaultSecondaryKey string `json:"-" sensitive:"true" tf:"default_secondary_key,omitempty"`
	Location            string `json:"location" tf:"location"`
	// +optional
	MaximumThroughputUnits int64  `json:"maximumThroughputUnits,omitempty" tf:"maximum_throughput_units,omitempty"`
	Name                   string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	NetworkRulesets   []EventhubNamespace_SpecNetworkRulesets `json:"networkRulesets,omitempty" tf:"network_rulesets,omitempty"`
	ResourceGroupName string                                  `json:"resourceGroupName" tf:"resource_group_name"`
	Sku               string                                  `json:"sku" tf:"sku"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EventhubNamespace_Status struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *EventhubNamespace_Spec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EventhubNamespace_List is a list of EventhubNamespace_s
type EventhubNamespace_List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EventhubNamespace_ CRD objects
	Items []EventhubNamespace_ `json:"items,omitempty"`
}
