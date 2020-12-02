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

type EventgridTopic struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventgridTopicSpec   `json:"spec,omitempty"`
	Status            EventgridTopicStatus `json:"status,omitempty"`
}

type EventgridTopicSpecInputMappingDefaultValues struct {
	// +optional
	DataVersion string `json:"dataVersion,omitempty" tf:"data_version,omitempty"`
	// +optional
	EventType string `json:"eventType,omitempty" tf:"event_type,omitempty"`
	// +optional
	Subject string `json:"subject,omitempty" tf:"subject,omitempty"`
}

type EventgridTopicSpecInputMappingFields struct {
	// +optional
	DataVersion string `json:"dataVersion,omitempty" tf:"data_version,omitempty"`
	// +optional
	EventTime string `json:"eventTime,omitempty" tf:"event_time,omitempty"`
	// +optional
	EventType string `json:"eventType,omitempty" tf:"event_type,omitempty"`
	// +optional
	ID string `json:"ID,omitempty" tf:"id,omitempty"`
	// +optional
	Subject string `json:"subject,omitempty" tf:"subject,omitempty"`
	// +optional
	Topic string `json:"topic,omitempty" tf:"topic,omitempty"`
}

type EventgridTopicSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	Endpoint string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	InputMappingDefaultValues []EventgridTopicSpecInputMappingDefaultValues `json:"inputMappingDefaultValues,omitempty" tf:"input_mapping_default_values,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	InputMappingFields []EventgridTopicSpecInputMappingFields `json:"inputMappingFields,omitempty" tf:"input_mapping_fields,omitempty"`
	// +optional
	InputSchema string `json:"inputSchema,omitempty" tf:"input_schema,omitempty"`
	Location    string `json:"location" tf:"location"`
	Name        string `json:"name" tf:"name"`
	// +optional
	PrimaryAccessKey  string `json:"-" sensitive:"true" tf:"primary_access_key,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SecondaryAccessKey string `json:"-" sensitive:"true" tf:"secondary_access_key,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EventgridTopicStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *EventgridTopicSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
	// +optional
	TerraFormState string `json:"terraformState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EventgridTopicList is a list of EventgridTopics
type EventgridTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EventgridTopic CRD objects
	Items []EventgridTopic `json:"items,omitempty"`
}
