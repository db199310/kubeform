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
// +kubebuilder:storageversion

type StreamAnalyticsJob struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StreamAnalyticsJobSpec   `json:"spec,omitempty"`
	Status            StreamAnalyticsJobStatus `json:"status,omitempty"`
}

type StreamAnalyticsJobSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CompatibilityLevel string `json:"compatibilityLevel,omitempty" tf:"compatibility_level,omitempty"`
	// +optional
	DataLocale string `json:"dataLocale,omitempty" tf:"data_locale,omitempty"`
	// +optional
	EventsLateArrivalMaxDelayInSeconds int64 `json:"eventsLateArrivalMaxDelayInSeconds,omitempty" tf:"events_late_arrival_max_delay_in_seconds,omitempty"`
	// +optional
	EventsOutOfOrderMaxDelayInSeconds int64 `json:"eventsOutOfOrderMaxDelayInSeconds,omitempty" tf:"events_out_of_order_max_delay_in_seconds,omitempty"`
	// +optional
	EventsOutOfOrderPolicy string `json:"eventsOutOfOrderPolicy,omitempty" tf:"events_out_of_order_policy,omitempty"`
	// +optional
	JobID    string `json:"jobID,omitempty" tf:"job_id,omitempty"`
	Location string `json:"location" tf:"location"`
	Name     string `json:"name" tf:"name"`
	// +optional
	OutputErrorPolicy string `json:"outputErrorPolicy,omitempty" tf:"output_error_policy,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	StreamingUnits    int64  `json:"streamingUnits" tf:"streaming_units"`
	// +optional
	Tags                map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	TransformationQuery string            `json:"transformationQuery" tf:"transformation_query"`
}

type StreamAnalyticsJobStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *StreamAnalyticsJobSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase           base.Phase `json:"phase,omitempty"`
	TerraformErrors []string   `json:"terraformErrors,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StreamAnalyticsJobList is a list of StreamAnalyticsJobs
type StreamAnalyticsJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StreamAnalyticsJob CRD objects
	Items []StreamAnalyticsJob `json:"items,omitempty"`
}
