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

type ContainerRegistry struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContainerRegistrySpec   `json:"spec,omitempty"`
	Status            ContainerRegistryStatus `json:"status,omitempty"`
}

type ContainerRegistrySpecNetworkRuleSetIpRule struct {
	Action  string `json:"action" tf:"action"`
	IpRange string `json:"ipRange" tf:"ip_range"`
}

type ContainerRegistrySpecNetworkRuleSetVirtualNetwork struct {
	Action   string `json:"action" tf:"action"`
	SubnetID string `json:"subnetID" tf:"subnet_id"`
}

type ContainerRegistrySpecNetworkRuleSet struct {
	// +optional
	DefaultAction string `json:"defaultAction,omitempty" tf:"default_action,omitempty"`
	// +optional
	IpRule []ContainerRegistrySpecNetworkRuleSetIpRule `json:"ipRule,omitempty" tf:"ip_rule,omitempty"`
	// +optional
	VirtualNetwork []ContainerRegistrySpecNetworkRuleSetVirtualNetwork `json:"virtualNetwork,omitempty" tf:"virtual_network,omitempty"`
}

type ContainerRegistrySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	AdminEnabled bool `json:"adminEnabled,omitempty" tf:"admin_enabled,omitempty"`
	// +optional
	AdminPassword string `json:"-" sensitive:"true" tf:"admin_password,omitempty"`
	// +optional
	AdminUsername string `json:"adminUsername,omitempty" tf:"admin_username,omitempty"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	GeoreplicationLocations []string `json:"georeplicationLocations,omitempty" tf:"georeplication_locations,omitempty"`
	Location                string   `json:"location" tf:"location"`
	// +optional
	LoginServer string `json:"loginServer,omitempty" tf:"login_server,omitempty"`
	Name        string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	NetworkRuleSet    []ContainerRegistrySpecNetworkRuleSet `json:"networkRuleSet,omitempty" tf:"network_rule_set,omitempty"`
	ResourceGroupName string                                `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Sku string `json:"sku,omitempty" tf:"sku,omitempty"`
	// +optional
	StorageAccountID string `json:"storageAccountID,omitempty" tf:"storage_account_id,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ContainerRegistryStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ContainerRegistrySpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ContainerRegistryList is a list of ContainerRegistrys
type ContainerRegistryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ContainerRegistry CRD objects
	Items []ContainerRegistry `json:"items,omitempty"`
}
