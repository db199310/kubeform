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

type MonitorActionRuleSuppression struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorActionRuleSuppressionSpec   `json:"spec,omitempty"`
	Status            MonitorActionRuleSuppressionStatus `json:"status,omitempty"`
}

type MonitorActionRuleSuppressionSpecConditionAlertContext struct {
	Operator string   `json:"operator" tf:"operator"`
	Values   []string `json:"values" tf:"values"`
}

type MonitorActionRuleSuppressionSpecConditionAlertRuleID struct {
	Operator string   `json:"operator" tf:"operator"`
	Values   []string `json:"values" tf:"values"`
}

type MonitorActionRuleSuppressionSpecConditionDescription struct {
	Operator string   `json:"operator" tf:"operator"`
	Values   []string `json:"values" tf:"values"`
}

type MonitorActionRuleSuppressionSpecConditionMonitor struct {
	Operator string   `json:"operator" tf:"operator"`
	Values   []string `json:"values" tf:"values"`
}

type MonitorActionRuleSuppressionSpecConditionMonitorService struct {
	Operator string   `json:"operator" tf:"operator"`
	Values   []string `json:"values" tf:"values"`
}

type MonitorActionRuleSuppressionSpecConditionSeverity struct {
	Operator string   `json:"operator" tf:"operator"`
	Values   []string `json:"values" tf:"values"`
}

type MonitorActionRuleSuppressionSpecConditionTargetResourceType struct {
	Operator string   `json:"operator" tf:"operator"`
	Values   []string `json:"values" tf:"values"`
}

type MonitorActionRuleSuppressionSpecCondition struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AlertContext []MonitorActionRuleSuppressionSpecConditionAlertContext `json:"alertContext,omitempty" tf:"alert_context,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AlertRuleID []MonitorActionRuleSuppressionSpecConditionAlertRuleID `json:"alertRuleID,omitempty" tf:"alert_rule_id,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Description []MonitorActionRuleSuppressionSpecConditionDescription `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Monitor []MonitorActionRuleSuppressionSpecConditionMonitor `json:"monitor,omitempty" tf:"monitor,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	MonitorService []MonitorActionRuleSuppressionSpecConditionMonitorService `json:"monitorService,omitempty" tf:"monitor_service,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Severity []MonitorActionRuleSuppressionSpecConditionSeverity `json:"severity,omitempty" tf:"severity,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	TargetResourceType []MonitorActionRuleSuppressionSpecConditionTargetResourceType `json:"targetResourceType,omitempty" tf:"target_resource_type,omitempty"`
}

type MonitorActionRuleSuppressionSpecScope struct {
	ResourceIDS []string `json:"resourceIDS" tf:"resource_ids"`
	Type        string   `json:"type" tf:"type"`
}

type MonitorActionRuleSuppressionSpecSuppressionSchedule struct {
	EndDateUtc string `json:"endDateUtc" tf:"end_date_utc"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	RecurrenceMonthly []int64 `json:"recurrenceMonthly,omitempty" tf:"recurrence_monthly,omitempty"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	RecurrenceWeekly []string `json:"recurrenceWeekly,omitempty" tf:"recurrence_weekly,omitempty"`
	StartDateUtc     string   `json:"startDateUtc" tf:"start_date_utc"`
}

type MonitorActionRuleSuppressionSpecSuppression struct {
	RecurrenceType string `json:"recurrenceType" tf:"recurrence_type"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Schedule []MonitorActionRuleSuppressionSpecSuppressionSchedule `json:"schedule,omitempty" tf:"schedule,omitempty"`
}

type MonitorActionRuleSuppressionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	Condition []MonitorActionRuleSuppressionSpecCondition `json:"condition,omitempty" tf:"condition,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Enabled           bool   `json:"enabled,omitempty" tf:"enabled,omitempty"`
	Name              string `json:"name" tf:"name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Scope []MonitorActionRuleSuppressionSpecScope `json:"scope,omitempty" tf:"scope,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Suppression []MonitorActionRuleSuppressionSpecSuppression `json:"suppression" tf:"suppression"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MonitorActionRuleSuppressionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *MonitorActionRuleSuppressionSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MonitorActionRuleSuppressionList is a list of MonitorActionRuleSuppressions
type MonitorActionRuleSuppressionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MonitorActionRuleSuppression CRD objects
	Items []MonitorActionRuleSuppression `json:"items,omitempty"`
}
