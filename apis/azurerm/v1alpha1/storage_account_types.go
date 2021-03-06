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

type StorageAccount struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageAccountSpec   `json:"spec,omitempty"`
	Status            StorageAccountStatus `json:"status,omitempty"`
}

type StorageAccountSpecBlobPropertiesDeleteRetentionPolicy struct {
	// +optional
	Days int64 `json:"days,omitempty" tf:"days,omitempty"`
}

type StorageAccountSpecBlobProperties struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DeleteRetentionPolicy []StorageAccountSpecBlobPropertiesDeleteRetentionPolicy `json:"deleteRetentionPolicy,omitempty" tf:"delete_retention_policy,omitempty"`
}

type StorageAccountSpecCustomDomain struct {
	Name string `json:"name" tf:"name"`
	// +optional
	UseSubdomain bool `json:"useSubdomain,omitempty" tf:"use_subdomain,omitempty"`
}

type StorageAccountSpecIdentity struct {
	// +optional
	PrincipalID string `json:"principalID,omitempty" tf:"principal_id,omitempty"`
	// +optional
	TenantID string `json:"tenantID,omitempty" tf:"tenant_id,omitempty"`
	Type     string `json:"type" tf:"type"`
}

type StorageAccountSpecNetworkRules struct {
	// +optional
	Bypass        []string `json:"bypass,omitempty" tf:"bypass,omitempty"`
	DefaultAction string   `json:"defaultAction" tf:"default_action"`
	// +optional
	IpRules []string `json:"ipRules,omitempty" tf:"ip_rules,omitempty"`
	// +optional
	VirtualNetworkSubnetIDS []string `json:"virtualNetworkSubnetIDS,omitempty" tf:"virtual_network_subnet_ids,omitempty"`
}

type StorageAccountSpecQueuePropertiesCorsRule struct {
	// +kubebuilder:validation:MaxItems=64
	AllowedHeaders []string `json:"allowedHeaders" tf:"allowed_headers"`
	// +kubebuilder:validation:MaxItems=64
	AllowedMethods []string `json:"allowedMethods" tf:"allowed_methods"`
	// +kubebuilder:validation:MaxItems=64
	AllowedOrigins []string `json:"allowedOrigins" tf:"allowed_origins"`
	// +kubebuilder:validation:MaxItems=64
	ExposedHeaders  []string `json:"exposedHeaders" tf:"exposed_headers"`
	MaxAgeInSeconds int64    `json:"maxAgeInSeconds" tf:"max_age_in_seconds"`
}

type StorageAccountSpecQueuePropertiesHourMetrics struct {
	Enabled bool `json:"enabled" tf:"enabled"`
	// +optional
	IncludeApis bool `json:"includeApis,omitempty" tf:"include_apis,omitempty"`
	// +optional
	RetentionPolicyDays int64  `json:"retentionPolicyDays,omitempty" tf:"retention_policy_days,omitempty"`
	Version             string `json:"version" tf:"version"`
}

type StorageAccountSpecQueuePropertiesLogging struct {
	Delete bool `json:"delete" tf:"delete"`
	Read   bool `json:"read" tf:"read"`
	// +optional
	RetentionPolicyDays int64  `json:"retentionPolicyDays,omitempty" tf:"retention_policy_days,omitempty"`
	Version             string `json:"version" tf:"version"`
	Write               bool   `json:"write" tf:"write"`
}

type StorageAccountSpecQueuePropertiesMinuteMetrics struct {
	Enabled bool `json:"enabled" tf:"enabled"`
	// +optional
	IncludeApis bool `json:"includeApis,omitempty" tf:"include_apis,omitempty"`
	// +optional
	RetentionPolicyDays int64  `json:"retentionPolicyDays,omitempty" tf:"retention_policy_days,omitempty"`
	Version             string `json:"version" tf:"version"`
}

type StorageAccountSpecQueueProperties struct {
	// +optional
	// +kubebuilder:validation:MaxItems=5
	CorsRule []StorageAccountSpecQueuePropertiesCorsRule `json:"corsRule,omitempty" tf:"cors_rule,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HourMetrics []StorageAccountSpecQueuePropertiesHourMetrics `json:"hourMetrics,omitempty" tf:"hour_metrics,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Logging []StorageAccountSpecQueuePropertiesLogging `json:"logging,omitempty" tf:"logging,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	MinuteMetrics []StorageAccountSpecQueuePropertiesMinuteMetrics `json:"minuteMetrics,omitempty" tf:"minute_metrics,omitempty"`
}

type StorageAccountSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	AccessTier string `json:"accessTier,omitempty" tf:"access_tier,omitempty"`
	// +optional
	AccountEncryptionSource string `json:"accountEncryptionSource,omitempty" tf:"account_encryption_source,omitempty"`
	// +optional
	AccountKind            string `json:"accountKind,omitempty" tf:"account_kind,omitempty"`
	AccountReplicationType string `json:"accountReplicationType" tf:"account_replication_type"`
	AccountTier            string `json:"accountTier" tf:"account_tier"`
	// +optional
	// Deprecated
	AccountType string `json:"accountType,omitempty" tf:"account_type,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	BlobProperties []StorageAccountSpecBlobProperties `json:"blobProperties,omitempty" tf:"blob_properties,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CustomDomain []StorageAccountSpecCustomDomain `json:"customDomain,omitempty" tf:"custom_domain,omitempty"`
	// +optional
	// Deprecated
	EnableAdvancedThreatProtection bool `json:"enableAdvancedThreatProtection,omitempty" tf:"enable_advanced_threat_protection,omitempty"`
	// +optional
	EnableBlobEncryption bool `json:"enableBlobEncryption,omitempty" tf:"enable_blob_encryption,omitempty"`
	// +optional
	EnableFileEncryption bool `json:"enableFileEncryption,omitempty" tf:"enable_file_encryption,omitempty"`
	// +optional
	EnableHTTPSTrafficOnly bool `json:"enableHTTPSTrafficOnly,omitempty" tf:"enable_https_traffic_only,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Identity []StorageAccountSpecIdentity `json:"identity,omitempty" tf:"identity,omitempty"`
	// +optional
	IsHnsEnabled bool   `json:"isHnsEnabled,omitempty" tf:"is_hns_enabled,omitempty"`
	Location     string `json:"location" tf:"location"`
	Name         string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	NetworkRules []StorageAccountSpecNetworkRules `json:"networkRules,omitempty" tf:"network_rules,omitempty"`
	// +optional
	PrimaryAccessKey string `json:"-" sensitive:"true" tf:"primary_access_key,omitempty"`
	// +optional
	PrimaryBlobConnectionString string `json:"-" sensitive:"true" tf:"primary_blob_connection_string,omitempty"`
	// +optional
	PrimaryBlobEndpoint string `json:"primaryBlobEndpoint,omitempty" tf:"primary_blob_endpoint,omitempty"`
	// +optional
	PrimaryBlobHost string `json:"primaryBlobHost,omitempty" tf:"primary_blob_host,omitempty"`
	// +optional
	PrimaryConnectionString string `json:"-" sensitive:"true" tf:"primary_connection_string,omitempty"`
	// +optional
	PrimaryDfsEndpoint string `json:"primaryDfsEndpoint,omitempty" tf:"primary_dfs_endpoint,omitempty"`
	// +optional
	PrimaryDfsHost string `json:"primaryDfsHost,omitempty" tf:"primary_dfs_host,omitempty"`
	// +optional
	PrimaryFileEndpoint string `json:"primaryFileEndpoint,omitempty" tf:"primary_file_endpoint,omitempty"`
	// +optional
	PrimaryFileHost string `json:"primaryFileHost,omitempty" tf:"primary_file_host,omitempty"`
	// +optional
	PrimaryLocation string `json:"primaryLocation,omitempty" tf:"primary_location,omitempty"`
	// +optional
	PrimaryQueueEndpoint string `json:"primaryQueueEndpoint,omitempty" tf:"primary_queue_endpoint,omitempty"`
	// +optional
	PrimaryQueueHost string `json:"primaryQueueHost,omitempty" tf:"primary_queue_host,omitempty"`
	// +optional
	PrimaryTableEndpoint string `json:"primaryTableEndpoint,omitempty" tf:"primary_table_endpoint,omitempty"`
	// +optional
	PrimaryTableHost string `json:"primaryTableHost,omitempty" tf:"primary_table_host,omitempty"`
	// +optional
	PrimaryWebEndpoint string `json:"primaryWebEndpoint,omitempty" tf:"primary_web_endpoint,omitempty"`
	// +optional
	PrimaryWebHost string `json:"primaryWebHost,omitempty" tf:"primary_web_host,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	QueueProperties   []StorageAccountSpecQueueProperties `json:"queueProperties,omitempty" tf:"queue_properties,omitempty"`
	ResourceGroupName string                              `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SecondaryAccessKey string `json:"-" sensitive:"true" tf:"secondary_access_key,omitempty"`
	// +optional
	SecondaryBlobConnectionString string `json:"-" sensitive:"true" tf:"secondary_blob_connection_string,omitempty"`
	// +optional
	SecondaryBlobEndpoint string `json:"secondaryBlobEndpoint,omitempty" tf:"secondary_blob_endpoint,omitempty"`
	// +optional
	SecondaryBlobHost string `json:"secondaryBlobHost,omitempty" tf:"secondary_blob_host,omitempty"`
	// +optional
	SecondaryConnectionString string `json:"-" sensitive:"true" tf:"secondary_connection_string,omitempty"`
	// +optional
	SecondaryDfsEndpoint string `json:"secondaryDfsEndpoint,omitempty" tf:"secondary_dfs_endpoint,omitempty"`
	// +optional
	SecondaryDfsHost string `json:"secondaryDfsHost,omitempty" tf:"secondary_dfs_host,omitempty"`
	// +optional
	SecondaryFileEndpoint string `json:"secondaryFileEndpoint,omitempty" tf:"secondary_file_endpoint,omitempty"`
	// +optional
	SecondaryFileHost string `json:"secondaryFileHost,omitempty" tf:"secondary_file_host,omitempty"`
	// +optional
	SecondaryLocation string `json:"secondaryLocation,omitempty" tf:"secondary_location,omitempty"`
	// +optional
	SecondaryQueueEndpoint string `json:"secondaryQueueEndpoint,omitempty" tf:"secondary_queue_endpoint,omitempty"`
	// +optional
	SecondaryQueueHost string `json:"secondaryQueueHost,omitempty" tf:"secondary_queue_host,omitempty"`
	// +optional
	SecondaryTableEndpoint string `json:"secondaryTableEndpoint,omitempty" tf:"secondary_table_endpoint,omitempty"`
	// +optional
	SecondaryTableHost string `json:"secondaryTableHost,omitempty" tf:"secondary_table_host,omitempty"`
	// +optional
	SecondaryWebEndpoint string `json:"secondaryWebEndpoint,omitempty" tf:"secondary_web_endpoint,omitempty"`
	// +optional
	SecondaryWebHost string `json:"secondaryWebHost,omitempty" tf:"secondary_web_host,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type StorageAccountStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *StorageAccountSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
	// +optional
	TerraformErrors []string `json:"terraformErrors,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StorageAccountList is a list of StorageAccounts
type StorageAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageAccount CRD objects
	Items []StorageAccount `json:"items,omitempty"`
}
