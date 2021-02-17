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

type Snapshot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnapshotSpec   `json:"spec,omitempty"`
	Status            SnapshotStatus `json:"status,omitempty"`
}

type SnapshotSpecEncryptionSettingsDiskEncryptionKey struct {
	SecretURL     string `json:"secretURL" tf:"secret_url"`
	SourceVaultID string `json:"sourceVaultID" tf:"source_vault_id"`
}

type SnapshotSpecEncryptionSettingsKeyEncryptionKey struct {
	KeyURL        string `json:"keyURL" tf:"key_url"`
	SourceVaultID string `json:"sourceVaultID" tf:"source_vault_id"`
}

type SnapshotSpecEncryptionSettings struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DiskEncryptionKey []SnapshotSpecEncryptionSettingsDiskEncryptionKey `json:"diskEncryptionKey,omitempty" tf:"disk_encryption_key,omitempty"`
	Enabled           bool                                              `json:"enabled" tf:"enabled"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	KeyEncryptionKey []SnapshotSpecEncryptionSettingsKeyEncryptionKey `json:"keyEncryptionKey,omitempty" tf:"key_encryption_key,omitempty"`
}

type SnapshotSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	CreateOption string `json:"createOption" tf:"create_option"`
	// +optional
	DiskSizeGb int64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	EncryptionSettings []SnapshotSpecEncryptionSettings `json:"encryptionSettings,omitempty" tf:"encryption_settings,omitempty"`
	Location           string                           `json:"location" tf:"location"`
	Name               string                           `json:"name" tf:"name"`
	ResourceGroupName  string                           `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SourceResourceID string `json:"sourceResourceID,omitempty" tf:"source_resource_id,omitempty"`
	// +optional
	SourceURI string `json:"sourceURI,omitempty" tf:"source_uri,omitempty"`
	// +optional
	StorageAccountID string `json:"storageAccountID,omitempty" tf:"storage_account_id,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SnapshotStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SnapshotSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
	// +optional
	TerraformErrors []string `json:"terraformErrors,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SnapshotList is a list of Snapshots
type SnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Snapshot CRD objects
	Items []Snapshot `json:"items,omitempty"`
}
