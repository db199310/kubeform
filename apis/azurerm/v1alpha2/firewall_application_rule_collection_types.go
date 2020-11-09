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

package v1alpha2

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha2"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type FirewallApplicationRuleCollection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallApplicationRuleCollectionSpec   `json:"spec,omitempty"`
	Status            FirewallApplicationRuleCollectionStatus `json:"status,omitempty"`
}

type FirewallApplicationRuleCollectionSpecRuleProtocol struct {
	// +optional
	Port int64  `json:"port,omitempty" tf:"port,omitempty"`
	Type string `json:"type" tf:"type"`
}

type FirewallApplicationRuleCollectionSpecRule struct {
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	FqdnTags []string `json:"fqdnTags,omitempty" tf:"fqdn_tags,omitempty"`
	Name     string   `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	Protocol        []FirewallApplicationRuleCollectionSpecRuleProtocol `json:"protocol,omitempty" tf:"protocol,omitempty"`
	SourceAddresses []string                                            `json:"sourceAddresses" tf:"source_addresses"`
	// +optional
	TargetFqdns []string `json:"targetFqdns,omitempty" tf:"target_fqdns,omitempty"`
}

type FirewallApplicationRuleCollectionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Action            string `json:"action" tf:"action"`
	AzureFirewallName string `json:"azureFirewallName" tf:"azure_firewall_name"`
	Name              string `json:"name" tf:"name"`
	Priority          int64  `json:"priority" tf:"priority"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MinItems=1
	Rule []FirewallApplicationRuleCollectionSpecRule `json:"rule" tf:"rule"`
}

type FirewallApplicationRuleCollectionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *FirewallApplicationRuleCollectionSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// FirewallApplicationRuleCollectionList is a list of FirewallApplicationRuleCollections
type FirewallApplicationRuleCollectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of FirewallApplicationRuleCollection CRD objects
	Items []FirewallApplicationRuleCollection `json:"items,omitempty"`
}
