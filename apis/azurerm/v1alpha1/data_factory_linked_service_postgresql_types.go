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

type DataFactoryLinkedServicePostgresql struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataFactoryLinkedServicePostgresqlSpec   `json:"spec,omitempty"`
	Status            DataFactoryLinkedServicePostgresqlStatus `json:"status,omitempty"`
}

type DataFactoryLinkedServicePostgresqlSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AdditionalProperties map[string]string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`
	// +optional
	Annotations      []string `json:"annotations,omitempty" tf:"annotations,omitempty"`
	ConnectionString string   `json:"connectionString" tf:"connection_string"`
	DataFactoryName  string   `json:"dataFactoryName" tf:"data_factory_name"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	IntegrationRuntimeName string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`
	Name                   string `json:"name" tf:"name"`
	// +optional
	Parameters        map[string]string `json:"parameters,omitempty" tf:"parameters,omitempty"`
	ResourceGroupName string            `json:"resourceGroupName" tf:"resource_group_name"`
}

type DataFactoryLinkedServicePostgresqlStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DataFactoryLinkedServicePostgresqlSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
	// +optional
	TerraformErrors []string `json:"terraformErrors,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DataFactoryLinkedServicePostgresqlList is a list of DataFactoryLinkedServicePostgresqls
type DataFactoryLinkedServicePostgresqlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DataFactoryLinkedServicePostgresql CRD objects
	Items []DataFactoryLinkedServicePostgresql `json:"items,omitempty"`
}
