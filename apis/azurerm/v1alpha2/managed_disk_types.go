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

type ManagedDisk struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagedDiskSpec   `json:"spec,omitempty"`
	Status            ManagedDiskStatus `json:"status,omitempty"`
}

type ManagedDiskSpecEncryptionSettingsDiskEncryptionKey struct {
	SecretURL     string `json:"secretURL" tf:"secret_url"`
	SourceVaultID string `json:"sourceVaultID" tf:"source_vault_id"`
}

type ManagedDiskSpecEncryptionSettingsKeyEncryptionKey struct {
	KeyURL        string `json:"keyURL" tf:"key_url"`
	SourceVaultID string `json:"sourceVaultID" tf:"source_vault_id"`
}

type ManagedDiskSpecEncryptionSettings struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DiskEncryptionKey []ManagedDiskSpecEncryptionSettingsDiskEncryptionKey `json:"diskEncryptionKey,omitempty" tf:"disk_encryption_key,omitempty"`
	Enabled           bool                                                 `json:"enabled" tf:"enabled"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	KeyEncryptionKey []ManagedDiskSpecEncryptionSettingsKeyEncryptionKey `json:"keyEncryptionKey,omitempty" tf:"key_encryption_key,omitempty"`
}

type ManagedDiskSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	CreateOption string `json:"createOption" tf:"create_option"`
	// +optional
	DiskEncryptionSetID string `json:"diskEncryptionSetID,omitempty" tf:"disk_encryption_set_id,omitempty"`
	// +optional
	DiskIopsReadWrite int64 `json:"diskIopsReadWrite,omitempty" tf:"disk_iops_read_write,omitempty"`
	// +optional
	DiskMbpsReadWrite int64 `json:"diskMbpsReadWrite,omitempty" tf:"disk_mbps_read_write,omitempty"`
	// +optional
	DiskSizeGb int64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	EncryptionSettings []ManagedDiskSpecEncryptionSettings `json:"encryptionSettings,omitempty" tf:"encryption_settings,omitempty"`
	// +optional
	ImageReferenceID string `json:"imageReferenceID,omitempty" tf:"image_reference_id,omitempty"`
	Location         string `json:"location" tf:"location"`
	Name             string `json:"name" tf:"name"`
	// +optional
	OsType            string `json:"osType,omitempty" tf:"os_type,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SourceResourceID string `json:"sourceResourceID,omitempty" tf:"source_resource_id,omitempty"`
	// +optional
	SourceURI string `json:"sourceURI,omitempty" tf:"source_uri,omitempty"`
	// +optional
	StorageAccountID   string `json:"storageAccountID,omitempty" tf:"storage_account_id,omitempty"`
	StorageAccountType string `json:"storageAccountType" tf:"storage_account_type"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Zones []string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type ManagedDiskStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ManagedDiskSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ManagedDiskList is a list of ManagedDisks
type ManagedDiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ManagedDisk CRD objects
	Items []ManagedDisk `json:"items,omitempty"`
}
