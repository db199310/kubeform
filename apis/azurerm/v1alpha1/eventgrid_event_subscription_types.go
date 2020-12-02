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

type EventgridEventSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventgridEventSubscriptionSpec   `json:"spec,omitempty"`
	Status            EventgridEventSubscriptionStatus `json:"status,omitempty"`
}

type EventgridEventSubscriptionSpecAdvancedFilterBoolEquals struct {
	Key   string `json:"key" tf:"key"`
	Value bool   `json:"value" tf:"value"`
}

type EventgridEventSubscriptionSpecAdvancedFilterNumberGreaterThan struct {
	Key   string  `json:"key" tf:"key"`
	Value float64 `json:"value" tf:"value"`
}

type EventgridEventSubscriptionSpecAdvancedFilterNumberGreaterThanOrEquals struct {
	Key   string  `json:"key" tf:"key"`
	Value float64 `json:"value" tf:"value"`
}

type EventgridEventSubscriptionSpecAdvancedFilterNumberIn struct {
	Key string `json:"key" tf:"key"`
	// +kubebuilder:validation:MaxItems=5
	Values []float64 `json:"values" tf:"values"`
}

type EventgridEventSubscriptionSpecAdvancedFilterNumberLessThan struct {
	Key   string  `json:"key" tf:"key"`
	Value float64 `json:"value" tf:"value"`
}

type EventgridEventSubscriptionSpecAdvancedFilterNumberLessThanOrEquals struct {
	Key   string  `json:"key" tf:"key"`
	Value float64 `json:"value" tf:"value"`
}

type EventgridEventSubscriptionSpecAdvancedFilterNumberNotIn struct {
	Key string `json:"key" tf:"key"`
	// +kubebuilder:validation:MaxItems=5
	Values []float64 `json:"values" tf:"values"`
}

type EventgridEventSubscriptionSpecAdvancedFilterStringBeginsWith struct {
	Key string `json:"key" tf:"key"`
	// +kubebuilder:validation:MaxItems=5
	Values []string `json:"values" tf:"values"`
}

type EventgridEventSubscriptionSpecAdvancedFilterStringContains struct {
	Key string `json:"key" tf:"key"`
	// +kubebuilder:validation:MaxItems=5
	Values []string `json:"values" tf:"values"`
}

type EventgridEventSubscriptionSpecAdvancedFilterStringEndsWith struct {
	Key string `json:"key" tf:"key"`
	// +kubebuilder:validation:MaxItems=5
	Values []string `json:"values" tf:"values"`
}

type EventgridEventSubscriptionSpecAdvancedFilterStringIn struct {
	Key string `json:"key" tf:"key"`
	// +kubebuilder:validation:MaxItems=5
	Values []string `json:"values" tf:"values"`
}

type EventgridEventSubscriptionSpecAdvancedFilterStringNotIn struct {
	Key string `json:"key" tf:"key"`
	// +kubebuilder:validation:MaxItems=5
	Values []string `json:"values" tf:"values"`
}

type EventgridEventSubscriptionSpecAdvancedFilter struct {
	// +optional
	BoolEquals []EventgridEventSubscriptionSpecAdvancedFilterBoolEquals `json:"boolEquals,omitempty" tf:"bool_equals,omitempty"`
	// +optional
	NumberGreaterThan []EventgridEventSubscriptionSpecAdvancedFilterNumberGreaterThan `json:"numberGreaterThan,omitempty" tf:"number_greater_than,omitempty"`
	// +optional
	NumberGreaterThanOrEquals []EventgridEventSubscriptionSpecAdvancedFilterNumberGreaterThanOrEquals `json:"numberGreaterThanOrEquals,omitempty" tf:"number_greater_than_or_equals,omitempty"`
	// +optional
	NumberIn []EventgridEventSubscriptionSpecAdvancedFilterNumberIn `json:"numberIn,omitempty" tf:"number_in,omitempty"`
	// +optional
	NumberLessThan []EventgridEventSubscriptionSpecAdvancedFilterNumberLessThan `json:"numberLessThan,omitempty" tf:"number_less_than,omitempty"`
	// +optional
	NumberLessThanOrEquals []EventgridEventSubscriptionSpecAdvancedFilterNumberLessThanOrEquals `json:"numberLessThanOrEquals,omitempty" tf:"number_less_than_or_equals,omitempty"`
	// +optional
	NumberNotIn []EventgridEventSubscriptionSpecAdvancedFilterNumberNotIn `json:"numberNotIn,omitempty" tf:"number_not_in,omitempty"`
	// +optional
	StringBeginsWith []EventgridEventSubscriptionSpecAdvancedFilterStringBeginsWith `json:"stringBeginsWith,omitempty" tf:"string_begins_with,omitempty"`
	// +optional
	StringContains []EventgridEventSubscriptionSpecAdvancedFilterStringContains `json:"stringContains,omitempty" tf:"string_contains,omitempty"`
	// +optional
	StringEndsWith []EventgridEventSubscriptionSpecAdvancedFilterStringEndsWith `json:"stringEndsWith,omitempty" tf:"string_ends_with,omitempty"`
	// +optional
	StringIn []EventgridEventSubscriptionSpecAdvancedFilterStringIn `json:"stringIn,omitempty" tf:"string_in,omitempty"`
	// +optional
	StringNotIn []EventgridEventSubscriptionSpecAdvancedFilterStringNotIn `json:"stringNotIn,omitempty" tf:"string_not_in,omitempty"`
}

type EventgridEventSubscriptionSpecAzureFunctionEndpoint struct {
	FunctionID string `json:"functionID" tf:"function_id"`
	// +optional
	MaxEventsPerBatch int64 `json:"maxEventsPerBatch,omitempty" tf:"max_events_per_batch,omitempty"`
	// +optional
	PreferredBatchSizeInKilobytes int64 `json:"preferredBatchSizeInKilobytes,omitempty" tf:"preferred_batch_size_in_kilobytes,omitempty"`
}

type EventgridEventSubscriptionSpecEventhubEndpoint struct {
	// +optional
	EventhubID string `json:"eventhubID,omitempty" tf:"eventhub_id,omitempty"`
}

type EventgridEventSubscriptionSpecHybridConnectionEndpoint struct {
	// +optional
	HybridConnectionID string `json:"hybridConnectionID,omitempty" tf:"hybrid_connection_id,omitempty"`
}

type EventgridEventSubscriptionSpecRetryPolicy struct {
	EventTimeToLive     int64 `json:"eventTimeToLive" tf:"event_time_to_live"`
	MaxDeliveryAttempts int64 `json:"maxDeliveryAttempts" tf:"max_delivery_attempts"`
}

type EventgridEventSubscriptionSpecStorageBlobDeadLetterDestination struct {
	StorageAccountID         string `json:"storageAccountID" tf:"storage_account_id"`
	StorageBlobContainerName string `json:"storageBlobContainerName" tf:"storage_blob_container_name"`
}

type EventgridEventSubscriptionSpecStorageQueueEndpoint struct {
	QueueName        string `json:"queueName" tf:"queue_name"`
	StorageAccountID string `json:"storageAccountID" tf:"storage_account_id"`
}

type EventgridEventSubscriptionSpecSubjectFilter struct {
	// +optional
	CaseSensitive bool `json:"caseSensitive,omitempty" tf:"case_sensitive,omitempty"`
	// +optional
	SubjectBeginsWith string `json:"subjectBeginsWith,omitempty" tf:"subject_begins_with,omitempty"`
	// +optional
	SubjectEndsWith string `json:"subjectEndsWith,omitempty" tf:"subject_ends_with,omitempty"`
}

type EventgridEventSubscriptionSpecWebhookEndpoint struct {
	// +optional
	ActiveDirectoryAppIDOrURI string `json:"activeDirectoryAppIDOrURI,omitempty" tf:"active_directory_app_id_or_uri,omitempty"`
	// +optional
	ActiveDirectoryTenantID string `json:"activeDirectoryTenantID,omitempty" tf:"active_directory_tenant_id,omitempty"`
	// +optional
	BaseURL string `json:"baseURL,omitempty" tf:"base_url,omitempty"`
	// +optional
	MaxEventsPerBatch int64 `json:"maxEventsPerBatch,omitempty" tf:"max_events_per_batch,omitempty"`
	// +optional
	PreferredBatchSizeInKilobytes int64  `json:"preferredBatchSizeInKilobytes,omitempty" tf:"preferred_batch_size_in_kilobytes,omitempty"`
	Url                           string `json:"url" tf:"url"`
}

type EventgridEventSubscriptionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	AdvancedFilter []EventgridEventSubscriptionSpecAdvancedFilter `json:"advancedFilter,omitempty" tf:"advanced_filter,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AzureFunctionEndpoint []EventgridEventSubscriptionSpecAzureFunctionEndpoint `json:"azureFunctionEndpoint,omitempty" tf:"azure_function_endpoint,omitempty"`
	// +optional
	EventDeliverySchema string `json:"eventDeliverySchema,omitempty" tf:"event_delivery_schema,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// Deprecated
	EventhubEndpoint []EventgridEventSubscriptionSpecEventhubEndpoint `json:"eventhubEndpoint,omitempty" tf:"eventhub_endpoint,omitempty"`
	// +optional
	EventhubEndpointID string `json:"eventhubEndpointID,omitempty" tf:"eventhub_endpoint_id,omitempty"`
	// +optional
	ExpirationTimeUtc string `json:"expirationTimeUtc,omitempty" tf:"expiration_time_utc,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// Deprecated
	HybridConnectionEndpoint []EventgridEventSubscriptionSpecHybridConnectionEndpoint `json:"hybridConnectionEndpoint,omitempty" tf:"hybrid_connection_endpoint,omitempty"`
	// +optional
	HybridConnectionEndpointID string `json:"hybridConnectionEndpointID,omitempty" tf:"hybrid_connection_endpoint_id,omitempty"`
	// +optional
	IncludedEventTypes []string `json:"includedEventTypes,omitempty" tf:"included_event_types,omitempty"`
	// +optional
	Labels []string `json:"labels,omitempty" tf:"labels,omitempty"`
	Name   string   `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RetryPolicy []EventgridEventSubscriptionSpecRetryPolicy `json:"retryPolicy,omitempty" tf:"retry_policy,omitempty"`
	Scope       string                                      `json:"scope" tf:"scope"`
	// +optional
	ServiceBusQueueEndpointID string `json:"serviceBusQueueEndpointID,omitempty" tf:"service_bus_queue_endpoint_id,omitempty"`
	// +optional
	ServiceBusTopicEndpointID string `json:"serviceBusTopicEndpointID,omitempty" tf:"service_bus_topic_endpoint_id,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	StorageBlobDeadLetterDestination []EventgridEventSubscriptionSpecStorageBlobDeadLetterDestination `json:"storageBlobDeadLetterDestination,omitempty" tf:"storage_blob_dead_letter_destination,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	StorageQueueEndpoint []EventgridEventSubscriptionSpecStorageQueueEndpoint `json:"storageQueueEndpoint,omitempty" tf:"storage_queue_endpoint,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SubjectFilter []EventgridEventSubscriptionSpecSubjectFilter `json:"subjectFilter,omitempty" tf:"subject_filter,omitempty"`
	// +optional
	TopicName string `json:"topicName,omitempty" tf:"topic_name,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	WebhookEndpoint []EventgridEventSubscriptionSpecWebhookEndpoint `json:"webhookEndpoint,omitempty" tf:"webhook_endpoint,omitempty"`
}

type EventgridEventSubscriptionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *EventgridEventSubscriptionSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
	// +optional
	TerraFormState string `json:"terraformState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EventgridEventSubscriptionList is a list of EventgridEventSubscriptions
type EventgridEventSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EventgridEventSubscription CRD objects
	Items []EventgridEventSubscription `json:"items,omitempty"`
}
