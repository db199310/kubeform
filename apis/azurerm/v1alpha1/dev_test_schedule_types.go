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

type DevTestSchedule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DevTestScheduleSpec   `json:"spec,omitempty"`
	Status            DevTestScheduleStatus `json:"status,omitempty"`
}

type DevTestScheduleSpecDailyRecurrence struct {
	Time string `json:"time" tf:"time"`
}

type DevTestScheduleSpecHourlyRecurrence struct {
	Minute int64 `json:"minute" tf:"minute"`
}

type DevTestScheduleSpecNotificationSettings struct {
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
	// +optional
	TimeInMinutes int64 `json:"timeInMinutes,omitempty" tf:"time_in_minutes,omitempty"`
	// +optional
	WebhookURL string `json:"webhookURL,omitempty" tf:"webhook_url,omitempty"`
}

type DevTestScheduleSpecWeeklyRecurrence struct {
	Time string `json:"time" tf:"time"`
	// +optional
	WeekDays []string `json:"weekDays,omitempty" tf:"week_days,omitempty"`
}

type DevTestScheduleSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	DailyRecurrence []DevTestScheduleSpecDailyRecurrence `json:"dailyRecurrence,omitempty" tf:"daily_recurrence,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HourlyRecurrence []DevTestScheduleSpecHourlyRecurrence `json:"hourlyRecurrence,omitempty" tf:"hourly_recurrence,omitempty"`
	LabName          string                                `json:"labName" tf:"lab_name"`
	Location         string                                `json:"location" tf:"location"`
	Name             string                                `json:"name" tf:"name"`
	// +kubebuilder:validation:MaxItems=1
	NotificationSettings []DevTestScheduleSpecNotificationSettings `json:"notificationSettings" tf:"notification_settings"`
	ResourceGroupName    string                                    `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
	// +optional
	Tags       map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	TaskType   string            `json:"taskType" tf:"task_type"`
	TimeZoneID string            `json:"timeZoneID" tf:"time_zone_id"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	WeeklyRecurrence []DevTestScheduleSpecWeeklyRecurrence `json:"weeklyRecurrence,omitempty" tf:"weekly_recurrence,omitempty"`
}

type DevTestScheduleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DevTestScheduleSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
	// +optional
	TerraFormState string `json:"terraformState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DevTestScheduleList is a list of DevTestSchedules
type DevTestScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DevTestSchedule CRD objects
	Items []DevTestSchedule `json:"items,omitempty"`
}
