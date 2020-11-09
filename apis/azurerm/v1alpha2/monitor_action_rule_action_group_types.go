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

type MonitorActionRuleActionGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorActionRuleActionGroupSpec   `json:"spec,omitempty"`
	Status            MonitorActionRuleActionGroupStatus `json:"status,omitempty"`
}

type MonitorActionRuleActionGroupSpecConditionAlertContext struct {
	Operator string   `json:"operator" tf:"operator"`
	Values   []string `json:"values" tf:"values"`
}

type MonitorActionRuleActionGroupSpecConditionAlertRuleID struct {
	Operator string   `json:"operator" tf:"operator"`
	Values   []string `json:"values" tf:"values"`
}

type MonitorActionRuleActionGroupSpecConditionDescription struct {
	Operator string   `json:"operator" tf:"operator"`
	Values   []string `json:"values" tf:"values"`
}

type MonitorActionRuleActionGroupSpecConditionMonitor struct {
	Operator string   `json:"operator" tf:"operator"`
	Values   []string `json:"values" tf:"values"`
}

type MonitorActionRuleActionGroupSpecConditionMonitorService struct {
	Operator string   `json:"operator" tf:"operator"`
	Values   []string `json:"values" tf:"values"`
}

type MonitorActionRuleActionGroupSpecConditionSeverity struct {
	Operator string   `json:"operator" tf:"operator"`
	Values   []string `json:"values" tf:"values"`
}

type MonitorActionRuleActionGroupSpecConditionTargetResourceType struct {
	Operator string   `json:"operator" tf:"operator"`
	Values   []string `json:"values" tf:"values"`
}

type MonitorActionRuleActionGroupSpecCondition struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AlertContext []MonitorActionRuleActionGroupSpecConditionAlertContext `json:"alertContext,omitempty" tf:"alert_context,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AlertRuleID []MonitorActionRuleActionGroupSpecConditionAlertRuleID `json:"alertRuleID,omitempty" tf:"alert_rule_id,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Description []MonitorActionRuleActionGroupSpecConditionDescription `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Monitor []MonitorActionRuleActionGroupSpecConditionMonitor `json:"monitor,omitempty" tf:"monitor,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	MonitorService []MonitorActionRuleActionGroupSpecConditionMonitorService `json:"monitorService,omitempty" tf:"monitor_service,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Severity []MonitorActionRuleActionGroupSpecConditionSeverity `json:"severity,omitempty" tf:"severity,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	TargetResourceType []MonitorActionRuleActionGroupSpecConditionTargetResourceType `json:"targetResourceType,omitempty" tf:"target_resource_type,omitempty"`
}

type MonitorActionRuleActionGroupSpecScope struct {
	ResourceIDS []string `json:"resourceIDS" tf:"resource_ids"`
	Type        string   `json:"type" tf:"type"`
}

type MonitorActionRuleActionGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ActionGroupID string `json:"actionGroupID" tf:"action_group_id"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Condition []MonitorActionRuleActionGroupSpecCondition `json:"condition,omitempty" tf:"condition,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Enabled           bool   `json:"enabled,omitempty" tf:"enabled,omitempty"`
	Name              string `json:"name" tf:"name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Scope []MonitorActionRuleActionGroupSpecScope `json:"scope,omitempty" tf:"scope,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MonitorActionRuleActionGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *MonitorActionRuleActionGroupSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MonitorActionRuleActionGroupList is a list of MonitorActionRuleActionGroups
type MonitorActionRuleActionGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MonitorActionRuleActionGroup CRD objects
	Items []MonitorActionRuleActionGroup `json:"items,omitempty"`
}
