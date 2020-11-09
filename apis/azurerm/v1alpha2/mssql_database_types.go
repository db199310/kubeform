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

type MssqlDatabase struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MssqlDatabaseSpec   `json:"spec,omitempty"`
	Status            MssqlDatabaseStatus `json:"status,omitempty"`
}

type MssqlDatabaseSpecExtendedAuditingPolicy struct {
	// +optional
	RetentionInDays         int64  `json:"retentionInDays,omitempty" tf:"retention_in_days,omitempty"`
	StorageAccountAccessKey string `json:"-" sensitive:"true" tf:"storage_account_access_key"`
	// +optional
	StorageAccountAccessKeyIsSecondary bool   `json:"storageAccountAccessKeyIsSecondary,omitempty" tf:"storage_account_access_key_is_secondary,omitempty"`
	StorageEndpoint                    string `json:"storageEndpoint" tf:"storage_endpoint"`
}

type MssqlDatabaseSpecThreatDetectionPolicy struct {
	// +optional
	DisabledAlerts []string `json:"disabledAlerts,omitempty" tf:"disabled_alerts,omitempty"`
	// +optional
	EmailAccountAdmins string `json:"emailAccountAdmins,omitempty" tf:"email_account_admins,omitempty"`
	// +optional
	EmailAddresses []string `json:"emailAddresses,omitempty" tf:"email_addresses,omitempty"`
	// +optional
	RetentionDays int64 `json:"retentionDays,omitempty" tf:"retention_days,omitempty"`
	// +optional
	State string `json:"state,omitempty" tf:"state,omitempty"`
	// +optional
	StorageAccountAccessKey string `json:"-" sensitive:"true" tf:"storage_account_access_key,omitempty"`
	// +optional
	StorageEndpoint string `json:"storageEndpoint,omitempty" tf:"storage_endpoint,omitempty"`
	// +optional
	UseServerDefault string `json:"useServerDefault,omitempty" tf:"use_server_default,omitempty"`
}

type MssqlDatabaseSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	AutoPauseDelayInMinutes int64 `json:"autoPauseDelayInMinutes,omitempty" tf:"auto_pause_delay_in_minutes,omitempty"`
	// +optional
	Collation string `json:"collation,omitempty" tf:"collation,omitempty"`
	// +optional
	CreateMode string `json:"createMode,omitempty" tf:"create_mode,omitempty"`
	// +optional
	CreationSourceDatabaseID string `json:"creationSourceDatabaseID,omitempty" tf:"creation_source_database_id,omitempty"`
	// +optional
	ElasticPoolID string `json:"elasticPoolID,omitempty" tf:"elastic_pool_id,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ExtendedAuditingPolicy []MssqlDatabaseSpecExtendedAuditingPolicy `json:"extendedAuditingPolicy,omitempty" tf:"extended_auditing_policy,omitempty"`
	// +optional
	LicenseType string `json:"licenseType,omitempty" tf:"license_type,omitempty"`
	// +optional
	MaxSizeGb int64 `json:"maxSizeGb,omitempty" tf:"max_size_gb,omitempty"`
	// +optional
	MinCapacity float64 `json:"minCapacity,omitempty" tf:"min_capacity,omitempty"`
	Name        string  `json:"name" tf:"name"`
	// +optional
	ReadReplicaCount int64 `json:"readReplicaCount,omitempty" tf:"read_replica_count,omitempty"`
	// +optional
	ReadScale bool `json:"readScale,omitempty" tf:"read_scale,omitempty"`
	// +optional
	RestorePointInTime string `json:"restorePointInTime,omitempty" tf:"restore_point_in_time,omitempty"`
	// +optional
	SampleName string `json:"sampleName,omitempty" tf:"sample_name,omitempty"`
	ServerID   string `json:"serverID" tf:"server_id"`
	// +optional
	SkuName string `json:"skuName,omitempty" tf:"sku_name,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ThreatDetectionPolicy []MssqlDatabaseSpecThreatDetectionPolicy `json:"threatDetectionPolicy,omitempty" tf:"threat_detection_policy,omitempty"`
	// +optional
	ZoneRedundant bool `json:"zoneRedundant,omitempty" tf:"zone_redundant,omitempty"`
}

type MssqlDatabaseStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *MssqlDatabaseSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MssqlDatabaseList is a list of MssqlDatabases
type MssqlDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MssqlDatabase CRD objects
	Items []MssqlDatabase `json:"items,omitempty"`
}
