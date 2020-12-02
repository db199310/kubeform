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

type LogAnalyticsDatasourceWindowsPerformanceCounter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogAnalyticsDatasourceWindowsPerformanceCounterSpec   `json:"spec,omitempty"`
	Status            LogAnalyticsDatasourceWindowsPerformanceCounterStatus `json:"status,omitempty"`
}

type LogAnalyticsDatasourceWindowsPerformanceCounterSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	CounterName       string `json:"counterName" tf:"counter_name"`
	InstanceName      string `json:"instanceName" tf:"instance_name"`
	IntervalSeconds   int64  `json:"intervalSeconds" tf:"interval_seconds"`
	Name              string `json:"name" tf:"name"`
	ObjectName        string `json:"objectName" tf:"object_name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	WorkspaceName     string `json:"workspaceName" tf:"workspace_name"`
}

type LogAnalyticsDatasourceWindowsPerformanceCounterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *LogAnalyticsDatasourceWindowsPerformanceCounterSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
	// +optional
	TerraFormState string `json:"terraformState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LogAnalyticsDatasourceWindowsPerformanceCounterList is a list of LogAnalyticsDatasourceWindowsPerformanceCounters
type LogAnalyticsDatasourceWindowsPerformanceCounterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LogAnalyticsDatasourceWindowsPerformanceCounter CRD objects
	Items []LogAnalyticsDatasourceWindowsPerformanceCounter `json:"items,omitempty"`
}
