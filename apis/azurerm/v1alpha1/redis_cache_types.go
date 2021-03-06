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

type RedisCache struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisCacheSpec   `json:"spec,omitempty"`
	Status            RedisCacheStatus `json:"status,omitempty"`
}

type RedisCacheSpecPatchSchedule struct {
	DayOfWeek string `json:"dayOfWeek" tf:"day_of_week"`
	// +optional
	StartHourUtc int64 `json:"startHourUtc,omitempty" tf:"start_hour_utc,omitempty"`
}

type RedisCacheSpecRedisConfiguration struct {
	// +optional
	AofBackupEnabled bool `json:"aofBackupEnabled,omitempty" tf:"aof_backup_enabled,omitempty"`
	// +optional
	AofStorageConnectionString0 string `json:"-" sensitive:"true" tf:"aof_storage_connection_string_0,omitempty"`
	// +optional
	AofStorageConnectionString1 string `json:"-" sensitive:"true" tf:"aof_storage_connection_string_1,omitempty"`
	// +optional
	EnableAuthentication bool `json:"enableAuthentication,omitempty" tf:"enable_authentication,omitempty"`
	// +optional
	Maxclients int64 `json:"maxclients,omitempty" tf:"maxclients,omitempty"`
	// +optional
	MaxfragmentationmemoryReserved int64 `json:"maxfragmentationmemoryReserved,omitempty" tf:"maxfragmentationmemory_reserved,omitempty"`
	// +optional
	MaxmemoryDelta int64 `json:"maxmemoryDelta,omitempty" tf:"maxmemory_delta,omitempty"`
	// +optional
	MaxmemoryPolicy string `json:"maxmemoryPolicy,omitempty" tf:"maxmemory_policy,omitempty"`
	// +optional
	MaxmemoryReserved int64 `json:"maxmemoryReserved,omitempty" tf:"maxmemory_reserved,omitempty"`
	// +optional
	NotifyKeyspaceEvents string `json:"notifyKeyspaceEvents,omitempty" tf:"notify_keyspace_events,omitempty"`
	// +optional
	RdbBackupEnabled bool `json:"rdbBackupEnabled,omitempty" tf:"rdb_backup_enabled,omitempty"`
	// +optional
	RdbBackupFrequency int64 `json:"rdbBackupFrequency,omitempty" tf:"rdb_backup_frequency,omitempty"`
	// +optional
	RdbBackupMaxSnapshotCount int64 `json:"rdbBackupMaxSnapshotCount,omitempty" tf:"rdb_backup_max_snapshot_count,omitempty"`
	// +optional
	RdbStorageConnectionString string `json:"-" sensitive:"true" tf:"rdb_storage_connection_string,omitempty"`
}

type RedisCacheSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	Capacity int64 `json:"capacity" tf:"capacity"`
	// +optional
	EnableNonSSLPort bool   `json:"enableNonSSLPort,omitempty" tf:"enable_non_ssl_port,omitempty"`
	Family           string `json:"family" tf:"family"`
	// +optional
	Hostname string `json:"hostname,omitempty" tf:"hostname,omitempty"`
	Location string `json:"location" tf:"location"`
	// +optional
	MinimumTLSVersion string `json:"minimumTLSVersion,omitempty" tf:"minimum_tls_version,omitempty"`
	Name              string `json:"name" tf:"name"`
	// +optional
	PatchSchedule []RedisCacheSpecPatchSchedule `json:"patchSchedule,omitempty" tf:"patch_schedule,omitempty"`
	// +optional
	Port int64 `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	PrimaryAccessKey string `json:"-" sensitive:"true" tf:"primary_access_key,omitempty"`
	// +optional
	PrivateStaticIPAddress string `json:"privateStaticIPAddress,omitempty" tf:"private_static_ip_address,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RedisConfiguration []RedisCacheSpecRedisConfiguration `json:"redisConfiguration,omitempty" tf:"redis_configuration,omitempty"`
	ResourceGroupName  string                             `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SecondaryAccessKey string `json:"-" sensitive:"true" tf:"secondary_access_key,omitempty"`
	// +optional
	ShardCount int64  `json:"shardCount,omitempty" tf:"shard_count,omitempty"`
	SkuName    string `json:"skuName" tf:"sku_name"`
	// +optional
	SslPort int64 `json:"sslPort,omitempty" tf:"ssl_port,omitempty"`
	// +optional
	SubnetID string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Zones []string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type RedisCacheStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *RedisCacheSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
	// +optional
	TerraformErrors []string `json:"terraformErrors,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RedisCacheList is a list of RedisCaches
type RedisCacheList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RedisCache CRD objects
	Items []RedisCache `json:"items,omitempty"`
}
