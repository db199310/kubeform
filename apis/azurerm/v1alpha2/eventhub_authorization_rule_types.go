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

type EventhubAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventhubAuthorizationRuleSpec   `json:"spec,omitempty"`
	Status            EventhubAuthorizationRuleStatus `json:"status,omitempty"`
}

type EventhubAuthorizationRuleSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	EventhubName string `json:"eventhubName" tf:"eventhub_name"`
	// +optional
	Listen bool `json:"listen,omitempty" tf:"listen,omitempty"`
	// +optional
	Manage        bool   `json:"manage,omitempty" tf:"manage,omitempty"`
	Name          string `json:"name" tf:"name"`
	NamespaceName string `json:"namespaceName" tf:"namespace_name"`
	// +optional
	PrimaryConnectionString string `json:"-" sensitive:"true" tf:"primary_connection_string,omitempty"`
	// +optional
	PrimaryConnectionStringAlias string `json:"-" sensitive:"true" tf:"primary_connection_string_alias,omitempty"`
	// +optional
	PrimaryKey        string `json:"-" sensitive:"true" tf:"primary_key,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SecondaryConnectionString string `json:"-" sensitive:"true" tf:"secondary_connection_string,omitempty"`
	// +optional
	SecondaryConnectionStringAlias string `json:"-" sensitive:"true" tf:"secondary_connection_string_alias,omitempty"`
	// +optional
	SecondaryKey string `json:"-" sensitive:"true" tf:"secondary_key,omitempty"`
	// +optional
	Send bool `json:"send,omitempty" tf:"send,omitempty"`
}

type EventhubAuthorizationRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *EventhubAuthorizationRuleSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EventhubAuthorizationRuleList is a list of EventhubAuthorizationRules
type EventhubAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EventhubAuthorizationRule CRD objects
	Items []EventhubAuthorizationRule `json:"items,omitempty"`
}
