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

type DataFactoryTriggerSchedule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataFactoryTriggerScheduleSpec   `json:"spec,omitempty"`
	Status            DataFactoryTriggerScheduleStatus `json:"status,omitempty"`
}

type DataFactoryTriggerScheduleSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Annotations     []string `json:"annotations,omitempty" tf:"annotations,omitempty"`
	DataFactoryName string   `json:"dataFactoryName" tf:"data_factory_name"`
	// +optional
	EndTime string `json:"endTime,omitempty" tf:"end_time,omitempty"`
	// +optional
	Frequency string `json:"frequency,omitempty" tf:"frequency,omitempty"`
	// +optional
	Interval     int64  `json:"interval,omitempty" tf:"interval,omitempty"`
	Name         string `json:"name" tf:"name"`
	PipelineName string `json:"pipelineName" tf:"pipeline_name"`
	// +optional
	PipelineParameters map[string]string `json:"pipelineParameters,omitempty" tf:"pipeline_parameters,omitempty"`
	ResourceGroupName  string            `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	StartTime string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type DataFactoryTriggerScheduleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DataFactoryTriggerScheduleSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DataFactoryTriggerScheduleList is a list of DataFactoryTriggerSchedules
type DataFactoryTriggerScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DataFactoryTriggerSchedule CRD objects
	Items []DataFactoryTriggerSchedule `json:"items,omitempty"`
}
