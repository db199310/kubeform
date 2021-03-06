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

type EventgridEventSubscriptionSpecEventhubEndpoint struct {
	EventhubID string `json:"eventhubID" tf:"eventhub_id"`
}

type EventgridEventSubscriptionSpecHybridConnectionEndpoint struct {
	HybridConnectionID string `json:"hybridConnectionID" tf:"hybrid_connection_id"`
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
	Url string `json:"url" tf:"url"`
}

type EventgridEventSubscriptionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	EventDeliverySchema string `json:"eventDeliverySchema,omitempty" tf:"event_delivery_schema,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	EventhubEndpoint []EventgridEventSubscriptionSpecEventhubEndpoint `json:"eventhubEndpoint,omitempty" tf:"eventhub_endpoint,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HybridConnectionEndpoint []EventgridEventSubscriptionSpecHybridConnectionEndpoint `json:"hybridConnectionEndpoint,omitempty" tf:"hybrid_connection_endpoint,omitempty"`
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
	TerraformErrors []string `json:"terraformErrors,omitempty"`
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
