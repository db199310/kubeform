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

type MonitorMetricAlertrule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorMetricAlertruleSpec   `json:"spec,omitempty"`
	Status            MonitorMetricAlertruleStatus `json:"status,omitempty"`
}

type MonitorMetricAlertruleSpecEmailAction struct {
	// +optional
	CustomEmails []string `json:"customEmails,omitempty" tf:"custom_emails,omitempty"`
	// +optional
	SendToServiceOwners bool `json:"sendToServiceOwners,omitempty" tf:"send_to_service_owners,omitempty"`
}

type MonitorMetricAlertruleSpecWebhookAction struct {
	// +optional
	Properties map[string]string `json:"properties,omitempty" tf:"properties,omitempty"`
	ServiceURI string            `json:"serviceURI" tf:"service_uri"`
}

type MonitorMetricAlertruleSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Aggregation string `json:"aggregation" tf:"aggregation"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	EmailAction []MonitorMetricAlertruleSpecEmailAction `json:"emailAction,omitempty" tf:"email_action,omitempty"`
	// +optional
	Enabled           bool   `json:"enabled,omitempty" tf:"enabled,omitempty"`
	Location          string `json:"location" tf:"location"`
	MetricName        string `json:"metricName" tf:"metric_name"`
	Name              string `json:"name" tf:"name"`
	Operator          string `json:"operator" tf:"operator"`
	Period            string `json:"period" tf:"period"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	ResourceID        string `json:"resourceID" tf:"resource_id"`
	// +optional
	Tags      map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	Threshold float64           `json:"threshold" tf:"threshold"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	WebhookAction []MonitorMetricAlertruleSpecWebhookAction `json:"webhookAction,omitempty" tf:"webhook_action,omitempty"`
}

type MonitorMetricAlertruleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *MonitorMetricAlertruleSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
	// +optional
	TerraFormLogs *base.TerraFormLogs `json:"terraformLogs,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MonitorMetricAlertruleList is a list of MonitorMetricAlertrules
type MonitorMetricAlertruleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MonitorMetricAlertrule CRD objects
	Items []MonitorMetricAlertrule `json:"items,omitempty"`
}
