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

type SentinelAlertRuleMsSecurityIncident struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SentinelAlertRuleMsSecurityIncidentSpec   `json:"spec,omitempty"`
	Status            SentinelAlertRuleMsSecurityIncidentStatus `json:"status,omitempty"`
}

type SentinelAlertRuleMsSecurityIncidentSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	RemoteBackend *base.Backend `json:"remoteBackend,omitempty" tf:"-"`

	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	DisplayName string `json:"displayName" tf:"display_name"`
	// +optional
	Enabled                 bool   `json:"enabled,omitempty" tf:"enabled,omitempty"`
	LogAnalyticsWorkspaceID string `json:"logAnalyticsWorkspaceID" tf:"log_analytics_workspace_id"`
	Name                    string `json:"name" tf:"name"`
	ProductFilter           string `json:"productFilter" tf:"product_filter"`
	// +kubebuilder:validation:MinItems=1
	SeverityFilter []string `json:"severityFilter" tf:"severity_filter"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	TextWhitelist []string `json:"textWhitelist,omitempty" tf:"text_whitelist,omitempty"`
}

type SentinelAlertRuleMsSecurityIncidentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SentinelAlertRuleMsSecurityIncidentSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SentinelAlertRuleMsSecurityIncidentList is a list of SentinelAlertRuleMsSecurityIncidents
type SentinelAlertRuleMsSecurityIncidentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SentinelAlertRuleMsSecurityIncident CRD objects
	Items []SentinelAlertRuleMsSecurityIncident `json:"items,omitempty"`
}
