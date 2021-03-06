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

type KustoEventhubDataConnection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KustoEventhubDataConnectionSpec   `json:"spec,omitempty"`
	Status            KustoEventhubDataConnectionStatus `json:"status,omitempty"`
}

type KustoEventhubDataConnectionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ClusterName   string `json:"clusterName" tf:"cluster_name"`
	ConsumerGroup string `json:"consumerGroup" tf:"consumer_group"`
	// +optional
	DataFormat   string `json:"dataFormat,omitempty" tf:"data_format,omitempty"`
	DatabaseName string `json:"databaseName" tf:"database_name"`
	EventhubID   string `json:"eventhubID" tf:"eventhub_id"`
	Location     string `json:"location" tf:"location"`
	// +optional
	MappingRuleName   string `json:"mappingRuleName,omitempty" tf:"mapping_rule_name,omitempty"`
	Name              string `json:"name" tf:"name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	TableName string `json:"tableName,omitempty" tf:"table_name,omitempty"`
}

type KustoEventhubDataConnectionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *KustoEventhubDataConnectionSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
	// +optional
	TerraformErrors []string `json:"terraformErrors,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// KustoEventhubDataConnectionList is a list of KustoEventhubDataConnections
type KustoEventhubDataConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KustoEventhubDataConnection CRD objects
	Items []KustoEventhubDataConnection `json:"items,omitempty"`
}
