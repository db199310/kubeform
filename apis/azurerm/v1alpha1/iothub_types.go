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

type Iothub struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IothubSpec   `json:"spec,omitempty"`
	Status            IothubStatus `json:"status,omitempty"`
}

type IothubSpecEndpoint struct {
	// +optional
	BatchFrequencyInSeconds int64  `json:"batchFrequencyInSeconds,omitempty" tf:"batch_frequency_in_seconds,omitempty"`
	ConnectionString        string `json:"-" sensitive:"true" tf:"connection_string"`
	// +optional
	ContainerName string `json:"containerName,omitempty" tf:"container_name,omitempty"`
	// +optional
	Encoding string `json:"encoding,omitempty" tf:"encoding,omitempty"`
	// +optional
	FileNameFormat string `json:"fileNameFormat,omitempty" tf:"file_name_format,omitempty"`
	// +optional
	MaxChunkSizeInBytes int64  `json:"maxChunkSizeInBytes,omitempty" tf:"max_chunk_size_in_bytes,omitempty"`
	Name                string `json:"name" tf:"name"`
	Type                string `json:"type" tf:"type"`
}

type IothubSpecFallbackRoute struct {
	// +optional
	Condition string `json:"condition,omitempty" tf:"condition,omitempty"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	EndpointNames []string `json:"endpointNames,omitempty" tf:"endpoint_names,omitempty"`
	// +optional
	Source string `json:"source,omitempty" tf:"source,omitempty"`
}

type IothubSpecFileUpload struct {
	ConnectionString string `json:"-" sensitive:"true" tf:"connection_string"`
	ContainerName    string `json:"containerName" tf:"container_name"`
	// +optional
	DefaultTtl string `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`
	// +optional
	LockDuration string `json:"lockDuration,omitempty" tf:"lock_duration,omitempty"`
	// +optional
	MaxDeliveryCount int64 `json:"maxDeliveryCount,omitempty" tf:"max_delivery_count,omitempty"`
	// +optional
	Notifications bool `json:"notifications,omitempty" tf:"notifications,omitempty"`
	// +optional
	SasTtl string `json:"sasTtl,omitempty" tf:"sas_ttl,omitempty"`
}

type IothubSpecIpFilterRule struct {
	Action string `json:"action" tf:"action"`
	IpMask string `json:"ipMask" tf:"ip_mask"`
	Name   string `json:"name" tf:"name"`
}

type IothubSpecRoute struct {
	// +optional
	Condition     string   `json:"condition,omitempty" tf:"condition,omitempty"`
	Enabled       bool     `json:"enabled" tf:"enabled"`
	EndpointNames []string `json:"endpointNames" tf:"endpoint_names"`
	Name          string   `json:"name" tf:"name"`
	Source        string   `json:"source" tf:"source"`
}

type IothubSpecSharedAccessPolicy struct {
	// +optional
	KeyName string `json:"keyName,omitempty" tf:"key_name,omitempty"`
	// +optional
	Permissions string `json:"permissions,omitempty" tf:"permissions,omitempty"`
	// +optional
	PrimaryKey string `json:"-" sensitive:"true" tf:"primary_key,omitempty"`
	// +optional
	SecondaryKey string `json:"-" sensitive:"true" tf:"secondary_key,omitempty"`
}

type IothubSpecSku struct {
	Capacity int64  `json:"capacity" tf:"capacity"`
	Name     string `json:"name" tf:"name"`
	// +optional
	// Deprecated
	Tier string `json:"tier,omitempty" tf:"tier,omitempty"`
}

type IothubSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	Endpoint []IothubSpecEndpoint `json:"endpoint,omitempty" tf:"endpoint,omitempty"`
	// +optional
	EventHubEventsEndpoint string `json:"eventHubEventsEndpoint,omitempty" tf:"event_hub_events_endpoint,omitempty"`
	// +optional
	EventHubEventsPath string `json:"eventHubEventsPath,omitempty" tf:"event_hub_events_path,omitempty"`
	// +optional
	EventHubOperationsEndpoint string `json:"eventHubOperationsEndpoint,omitempty" tf:"event_hub_operations_endpoint,omitempty"`
	// +optional
	EventHubOperationsPath string `json:"eventHubOperationsPath,omitempty" tf:"event_hub_operations_path,omitempty"`
	// +optional
	EventHubPartitionCount int64 `json:"eventHubPartitionCount,omitempty" tf:"event_hub_partition_count,omitempty"`
	// +optional
	EventHubRetentionInDays int64 `json:"eventHubRetentionInDays,omitempty" tf:"event_hub_retention_in_days,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	FallbackRoute []IothubSpecFallbackRoute `json:"fallbackRoute,omitempty" tf:"fallback_route,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	FileUpload []IothubSpecFileUpload `json:"fileUpload,omitempty" tf:"file_upload,omitempty"`
	// +optional
	Hostname string `json:"hostname,omitempty" tf:"hostname,omitempty"`
	// +optional
	IpFilterRule      []IothubSpecIpFilterRule `json:"ipFilterRule,omitempty" tf:"ip_filter_rule,omitempty"`
	Location          string                   `json:"location" tf:"location"`
	Name              string                   `json:"name" tf:"name"`
	ResourceGroupName string                   `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Route []IothubSpecRoute `json:"route,omitempty" tf:"route,omitempty"`
	// +optional
	SharedAccessPolicy []IothubSpecSharedAccessPolicy `json:"sharedAccessPolicy,omitempty" tf:"shared_access_policy,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Sku []IothubSpecSku `json:"sku" tf:"sku"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
}

type IothubStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *IothubSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
	// +optional
	TerraformErrors []string `json:"terraformErrors,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IothubList is a list of Iothubs
type IothubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Iothub CRD objects
	Items []Iothub `json:"items,omitempty"`
}
