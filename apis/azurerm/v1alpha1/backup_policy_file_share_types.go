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

type BackupPolicyFileShare struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupPolicyFileShareSpec   `json:"spec,omitempty"`
	Status            BackupPolicyFileShareStatus `json:"status,omitempty"`
}

type BackupPolicyFileShareSpecBackup struct {
	Frequency string `json:"frequency" tf:"frequency"`
	Time      string `json:"time" tf:"time"`
}

type BackupPolicyFileShareSpecRetentionDaily struct {
	Count int64 `json:"count" tf:"count"`
}

type BackupPolicyFileShareSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:MaxItems=1
	Backup            []BackupPolicyFileShareSpecBackup `json:"backup" tf:"backup"`
	Name              string                            `json:"name" tf:"name"`
	RecoveryVaultName string                            `json:"recoveryVaultName" tf:"recovery_vault_name"`
	ResourceGroupName string                            `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	RetentionDaily []BackupPolicyFileShareSpecRetentionDaily `json:"retentionDaily" tf:"retention_daily"`
	// +optional
	Timezone string `json:"timezone,omitempty" tf:"timezone,omitempty"`
}

type BackupPolicyFileShareStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *BackupPolicyFileShareSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase           base.Phase `json:"phase,omitempty"`
	TerraformErrors []string   `json:"terraformErrors,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BackupPolicyFileShareList is a list of BackupPolicyFileShares
type BackupPolicyFileShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BackupPolicyFileShare CRD objects
	Items []BackupPolicyFileShare `json:"items,omitempty"`
}
