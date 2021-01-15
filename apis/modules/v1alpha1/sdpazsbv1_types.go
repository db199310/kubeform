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
	"encoding/json"

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
// +kubebuilder:storageversion

type SDPAzsbv1 struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SDPAzsbv1Spec   `json:"spec,omitempty"`
	Status            SDPAzsbv1Status `json:"status,omitempty"`
}

type SDPAzsbv1Spec struct {
	// +optional
	SecretRef   *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
	ProviderRef core.LocalObjectReference  `json:"providerRef" tf:"-"`
	// +optional
	Source string `json:"source" tf:"source"`

	// +optional
	// The number of message units.
	Capacity json.Number `json:"capacity,omitempty" tf:"capacity,omitempty"`
	// Environment. Upto 5 character. For e.g. dev, dev01 , prd01
	Environment string `json:"environment" tf:"environment"`
	// +optional
	// Instance number
	Instance string `json:"instance,omitempty" tf:"instance,omitempty"`
	// The name of the namespace.
	Name string `json:"name" tf:"name"`
	// +optional
	// Namespace Authorization rules
	NamespaceAuthRules []SDPAzsbv1NamespaceAuthRules `json:"namespaceAuthRules,omitempty" tf:"namespace_auth_rules,omitempty"`
	// owner
	Owner string `json:"owner" tf:"owner"`
	// +optional
	// placement
	Placement string `json:"placement,omitempty" tf:"placement,omitempty"`
	// +optional
	// project stream name
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	//  4 character project stream name/code
	ProjectStream string `json:"projectStream" tf:"projectStream"`
	// +optional
	// Queue authorization rules
	QueueAuthRule []SDPAzsbv1QueueAuthRule `json:"queueAuthRule,omitempty" tf:"queue_auth_rule,omitempty"`
	// +optional
	// Queues for the ServiceBus namespace
	Queues []SDPAzsbv1Queues `json:"queues,omitempty" tf:"queues,omitempty"`
	// region. Choose from australia, europe, asia, europe
	Region string `json:"region" tf:"region"`
	// +optional
	// releaseVersion
	ReleaseVersion string `json:"releaseVersion,omitempty" tf:"releaseVersion,omitempty"`
	// The name of an existing resource group.
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	// The SKU of the namespace. The options are: `Basic`, `Standard`, `Premium`.
	Sku string `json:"sku,omitempty" tf:"sku,omitempty"`
	// +optional
	SubnetIDS []string `json:"subnetIDS,omitempty" tf:"subnet_ids,omitempty"`
	// +optional
	// Subscription rules.
	SubscriptionRules json.RawMessage `json:"subscriptionRules,omitempty" tf:"subscription_rules,omitempty"`
	// +optional
	// subscription for the Topic in the namespace
	Subscriptions json.RawMessage `json:"subscriptions,omitempty" tf:"subscriptions,omitempty"`
	// +optional
	//  Map of tags to assign to the resources.
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	// Topic Authorization Rule
	TopicAuthRule []SDPAzsbv1TopicAuthRule `json:"topicAuthRule,omitempty" tf:"topic_auth_rule,omitempty"`
	// +optional
	// List of topics.
	Topics []SDPAzsbv1Topics `json:"topics,omitempty" tf:"topics,omitempty"`
	//  4 character project stream name/code
	WorkStream string `json:"workStream" tf:"workStream"`
}

type SDPAzsbv1NamespaceAuthRules struct {
	// +optional
	Listen *bool `json:"listen,omitempty" tf:"listen,omitempty"`
	// +optional
	Manage *bool `json:"manage,omitempty" tf:"manage,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	Send *bool `json:"send,omitempty" tf:"send,omitempty"`
}

type SDPAzsbv1QueueAuthRule struct {
	// +optional
	Listen *bool `json:"listen,omitempty" tf:"listen,omitempty"`
	// +optional
	Manage *bool `json:"manage,omitempty" tf:"manage,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	QueueName string `json:"queueName,omitempty" tf:"queue_name,omitempty"`
	// +optional
	Send *bool `json:"send,omitempty" tf:"send,omitempty"`
}

type SDPAzsbv1Queues struct {
	// +optional
	AutoDeleteOnIdle string `json:"autoDeleteOnIdle,omitempty" tf:"auto_delete_on_idle,omitempty"`
	// +optional
	DeadLetteringOnMessageExpiration *bool `json:"deadLetteringOnMessageExpiration,omitempty" tf:"dead_lettering_on_message_expiration,omitempty"`
	// +optional
	DefaultMessageTtl string `json:"defaultMessageTtl,omitempty" tf:"default_message_ttl,omitempty"`
	// +optional
	DuplicateDetectionHistoryTimeWindow string `json:"duplicateDetectionHistoryTimeWindow,omitempty" tf:"duplicate_detection_history_time_window,omitempty"`
	// +optional
	EnableExpress *bool `json:"enableExpress,omitempty" tf:"enable_express,omitempty"`
	// +optional
	EnablePartitioning *bool `json:"enablePartitioning,omitempty" tf:"enable_partitioning,omitempty"`
	// +optional
	LockDuration string `json:"lockDuration,omitempty" tf:"lock_duration,omitempty"`
	// +optional
	MaxDeliveryCount json.Number `json:"maxDeliveryCount,omitempty" tf:"max_delivery_count,omitempty"`
	// +optional
	MaxSizeInMegabytes json.Number `json:"maxSizeInMegabytes,omitempty" tf:"max_size_in_megabytes,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	RequiresDuplicateDetection *bool `json:"requiresDuplicateDetection,omitempty" tf:"requires_duplicate_detection,omitempty"`
	// +optional
	RequiresSession *bool `json:"requiresSession,omitempty" tf:"requires_session,omitempty"`
}

type SDPAzsbv1TopicAuthRule struct {
	// +optional
	Listen *bool `json:"listen,omitempty" tf:"listen,omitempty"`
	// +optional
	Manage *bool `json:"manage,omitempty" tf:"manage,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	Send *bool `json:"send,omitempty" tf:"send,omitempty"`
	// +optional
	TopicName string `json:"topicName,omitempty" tf:"topic_name,omitempty"`
}

type SDPAzsbv1Topics struct {
	// +optional
	AutoDeleteOnIdle string `json:"autoDeleteOnIdle,omitempty" tf:"auto_delete_on_idle,omitempty"`
	// +optional
	DefaultMessageTtl string `json:"defaultMessageTtl,omitempty" tf:"default_message_ttl,omitempty"`
	// +optional
	DuplicateDetectionHistoryTimeWindow string `json:"duplicateDetectionHistoryTimeWindow,omitempty" tf:"duplicate_detection_history_time_window,omitempty"`
	// +optional
	EnableBatchedOperations string `json:"enableBatchedOperations,omitempty" tf:"enable_batched_operations,omitempty"`
	// +optional
	EnableExpress *bool `json:"enableExpress,omitempty" tf:"enable_express,omitempty"`
	// +optional
	EnablePartitioning *bool `json:"enablePartitioning,omitempty" tf:"enable_partitioning,omitempty"`
	// +optional
	MaxSizeInMegabytes json.Number `json:"maxSizeInMegabytes,omitempty" tf:"max_size_in_megabytes,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	RequiresDuplicateDetection *bool `json:"requiresDuplicateDetection,omitempty" tf:"requires_duplicate_detection,omitempty"`
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
	// +optional
	SupportOrdering *bool `json:"supportOrdering,omitempty" tf:"support_ordering,omitempty"`
}

type SDPAzsbv1Output struct{}

type SDPAzsbv1Status struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SDPAzsbv1Output `json:"output,omitempty"`
	// +optional
	State string `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SDPAzsbv1List is a list of SDPAzsbv1s
type SDPAzsbv1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SDPAzsbv1 CRD objects
	Items []SDPAzsbv1 `json:"items,omitempty"`
}
