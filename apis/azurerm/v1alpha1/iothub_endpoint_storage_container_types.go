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

type IothubEndpointStorageContainer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IothubEndpointStorageContainerSpec   `json:"spec,omitempty"`
	Status            IothubEndpointStorageContainerStatus `json:"status,omitempty"`
}

type IothubEndpointStorageContainerSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	BatchFrequencyInSeconds int64  `json:"batchFrequencyInSeconds,omitempty" tf:"batch_frequency_in_seconds,omitempty"`
	ConnectionString        string `json:"-" sensitive:"true" tf:"connection_string"`
	ContainerName           string `json:"containerName" tf:"container_name"`
	// +optional
	Encoding string `json:"encoding,omitempty" tf:"encoding,omitempty"`
	// +optional
	FileNameFormat string `json:"fileNameFormat,omitempty" tf:"file_name_format,omitempty"`
	IothubName     string `json:"iothubName" tf:"iothub_name"`
	// +optional
	MaxChunkSizeInBytes int64  `json:"maxChunkSizeInBytes,omitempty" tf:"max_chunk_size_in_bytes,omitempty"`
	Name                string `json:"name" tf:"name"`
	ResourceGroupName   string `json:"resourceGroupName" tf:"resource_group_name"`
}

type IothubEndpointStorageContainerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *IothubEndpointStorageContainerSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
	// +optional
	TerraFormState string `json:"terraformState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IothubEndpointStorageContainerList is a list of IothubEndpointStorageContainers
type IothubEndpointStorageContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IothubEndpointStorageContainer CRD objects
	Items []IothubEndpointStorageContainer `json:"items,omitempty"`
}
