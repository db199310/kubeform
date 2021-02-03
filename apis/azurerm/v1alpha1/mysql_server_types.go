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

type MysqlServer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MysqlServerSpec   `json:"spec,omitempty"`
	Status            MysqlServerStatus `json:"status,omitempty"`
}

type MysqlServerSpecStorageProfile struct {
	// +optional
	// Deprecated
	AutoGrow string `json:"autoGrow,omitempty" tf:"auto_grow,omitempty"`
	// +optional
	// Deprecated
	BackupRetentionDays int64 `json:"backupRetentionDays,omitempty" tf:"backup_retention_days,omitempty"`
	// +optional
	// Deprecated
	GeoRedundantBackup string `json:"geoRedundantBackup,omitempty" tf:"geo_redundant_backup,omitempty"`
	// +optional
	// Deprecated
	StorageMb int64 `json:"storageMb,omitempty" tf:"storage_mb,omitempty"`
}

type MysqlServerSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	RemoteBackend *base.Backend `json:"remoteBackend,omitempty" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	AdministratorLogin string `json:"administratorLogin,omitempty" tf:"administrator_login,omitempty"`
	// +optional
	AdministratorLoginPassword string `json:"-" sensitive:"true" tf:"administrator_login_password,omitempty"`
	// +optional
	AutoGrowEnabled bool `json:"autoGrowEnabled,omitempty" tf:"auto_grow_enabled,omitempty"`
	// +optional
	BackupRetentionDays int64 `json:"backupRetentionDays,omitempty" tf:"backup_retention_days,omitempty"`
	// +optional
	CreateMode string `json:"createMode,omitempty" tf:"create_mode,omitempty"`
	// +optional
	CreationSourceServerID string `json:"creationSourceServerID,omitempty" tf:"creation_source_server_id,omitempty"`
	// +optional
	Fqdn string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`
	// +optional
	GeoRedundantBackupEnabled bool `json:"geoRedundantBackupEnabled,omitempty" tf:"geo_redundant_backup_enabled,omitempty"`
	// +optional
	InfrastructureEncryptionEnabled bool   `json:"infrastructureEncryptionEnabled,omitempty" tf:"infrastructure_encryption_enabled,omitempty"`
	Location                        string `json:"location" tf:"location"`
	Name                            string `json:"name" tf:"name"`
	// +optional
	PublicNetworkAccessEnabled bool   `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`
	ResourceGroupName          string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	RestorePointInTime string `json:"restorePointInTime,omitempty" tf:"restore_point_in_time,omitempty"`
	SkuName            string `json:"skuName" tf:"sku_name"`
	// +optional
	// Deprecated
	SslEnforcement string `json:"sslEnforcement,omitempty" tf:"ssl_enforcement,omitempty"`
	// +optional
	SslEnforcementEnabled bool `json:"sslEnforcementEnabled,omitempty" tf:"ssl_enforcement_enabled,omitempty"`
	// +optional
	SslMinimalTLSVersionEnforced string `json:"sslMinimalTLSVersionEnforced,omitempty" tf:"ssl_minimal_tls_version_enforced,omitempty"`
	// +optional
	StorageMb int64 `json:"storageMb,omitempty" tf:"storage_mb,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// Deprecated
	StorageProfile []MysqlServerSpecStorageProfile `json:"storageProfile,omitempty" tf:"storage_profile,omitempty"`
	// +optional
	Tags    map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	Version string            `json:"version" tf:"version"`
}

type MysqlServerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *MysqlServerSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MysqlServerList is a list of MysqlServers
type MysqlServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MysqlServer CRD objects
	Items []MysqlServer `json:"items,omitempty"`
}
