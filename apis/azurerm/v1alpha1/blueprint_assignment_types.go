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

type BlueprintAssignment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BlueprintAssignmentSpec   `json:"spec,omitempty"`
	Status            BlueprintAssignmentStatus `json:"status,omitempty"`
}

type BlueprintAssignmentSpecIdentity struct {
	// +kubebuilder:validation:MinItems=1
	IdentityIDS []string `json:"identityIDS" tf:"identity_ids"`
	// +optional
	PrincipalID string `json:"principalID,omitempty" tf:"principal_id,omitempty"`
	// +optional
	TenantID string `json:"tenantID,omitempty" tf:"tenant_id,omitempty"`
	Type     string `json:"type" tf:"type"`
}

type BlueprintAssignmentSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	RemoteBackend *base.Backend `json:"remoteBackend,omitempty" tf:"-"`

	// +optional
	BlueprintName string `json:"blueprintName,omitempty" tf:"blueprint_name,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	DisplayName string `json:"displayName,omitempty" tf:"display_name,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Identity []BlueprintAssignmentSpecIdentity `json:"identity,omitempty" tf:"identity,omitempty"`
	Location string                            `json:"location" tf:"location"`
	// +optional
	// +kubebuilder:validation:MaxItems=5
	LockExcludePrincipals []string `json:"lockExcludePrincipals,omitempty" tf:"lock_exclude_principals,omitempty"`
	// +optional
	LockMode string `json:"lockMode,omitempty" tf:"lock_mode,omitempty"`
	Name     string `json:"name" tf:"name"`
	// +optional
	ParameterValues string `json:"parameterValues,omitempty" tf:"parameter_values,omitempty"`
	// +optional
	ResourceGroups       string `json:"resourceGroups,omitempty" tf:"resource_groups,omitempty"`
	TargetSubscriptionID string `json:"targetSubscriptionID" tf:"target_subscription_id"`
	// +optional
	Type      string `json:"type,omitempty" tf:"type,omitempty"`
	VersionID string `json:"versionID" tf:"version_id"`
}

type BlueprintAssignmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *BlueprintAssignmentSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BlueprintAssignmentList is a list of BlueprintAssignments
type BlueprintAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BlueprintAssignment CRD objects
	Items []BlueprintAssignment `json:"items,omitempty"`
}
