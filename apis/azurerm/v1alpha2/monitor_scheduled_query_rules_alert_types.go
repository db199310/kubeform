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

type MonitorScheduledQueryRulesAlert struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorScheduledQueryRulesAlertSpec   `json:"spec,omitempty"`
	Status            MonitorScheduledQueryRulesAlertStatus `json:"status,omitempty"`
}

type MonitorScheduledQueryRulesAlertSpecAction struct {
	ActionGroup []string `json:"actionGroup" tf:"action_group"`
	// +optional
	CustomWebhookPayload string `json:"customWebhookPayload,omitempty" tf:"custom_webhook_payload,omitempty"`
	// +optional
	EmailSubject string `json:"emailSubject,omitempty" tf:"email_subject,omitempty"`
}

type MonitorScheduledQueryRulesAlertSpecTriggerMetricTrigger struct {
	MetricColumn      string  `json:"metricColumn" tf:"metric_column"`
	MetricTriggerType string  `json:"metricTriggerType" tf:"metric_trigger_type"`
	Operator          string  `json:"operator" tf:"operator"`
	Threshold         float64 `json:"threshold" tf:"threshold"`
}

type MonitorScheduledQueryRulesAlertSpecTrigger struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	MetricTrigger []MonitorScheduledQueryRulesAlertSpecTriggerMetricTrigger `json:"metricTrigger,omitempty" tf:"metric_trigger,omitempty"`
	Operator      string                                                    `json:"operator" tf:"operator"`
	Threshold     float64                                                   `json:"threshold" tf:"threshold"`
}

type MonitorScheduledQueryRulesAlertSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:MaxItems=1
	Action []MonitorScheduledQueryRulesAlertSpecAction `json:"action" tf:"action"`
	// +optional
	// +kubebuilder:validation:MaxItems=100
	AuthorizedResourceIDS []string `json:"authorizedResourceIDS,omitempty" tf:"authorized_resource_ids,omitempty"`
	DataSourceID          string   `json:"dataSourceID" tf:"data_source_id"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Enabled   bool   `json:"enabled,omitempty" tf:"enabled,omitempty"`
	Frequency int64  `json:"frequency" tf:"frequency"`
	Location  string `json:"location" tf:"location"`
	Name      string `json:"name" tf:"name"`
	Query     string `json:"query" tf:"query"`
	// +optional
	QueryType         string `json:"queryType,omitempty" tf:"query_type,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Severity int64 `json:"severity,omitempty" tf:"severity,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	Throttling int64 `json:"throttling,omitempty" tf:"throttling,omitempty"`
	TimeWindow int64 `json:"timeWindow" tf:"time_window"`
	// +kubebuilder:validation:MaxItems=1
	Trigger []MonitorScheduledQueryRulesAlertSpecTrigger `json:"trigger" tf:"trigger"`
}

type MonitorScheduledQueryRulesAlertStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *MonitorScheduledQueryRulesAlertSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MonitorScheduledQueryRulesAlertList is a list of MonitorScheduledQueryRulesAlerts
type MonitorScheduledQueryRulesAlertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MonitorScheduledQueryRulesAlert CRD objects
	Items []MonitorScheduledQueryRulesAlert `json:"items,omitempty"`
}
