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

type CosmosdbGremlinGraph struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CosmosdbGremlinGraphSpec   `json:"spec,omitempty"`
	Status            CosmosdbGremlinGraphStatus `json:"status,omitempty"`
}

type CosmosdbGremlinGraphSpecConflictResolutionPolicy struct {
	// +optional
	ConflictResolutionPath string `json:"conflictResolutionPath,omitempty" tf:"conflict_resolution_path,omitempty"`
	// +optional
	ConflictResolutionProcedure string `json:"conflictResolutionProcedure,omitempty" tf:"conflict_resolution_procedure,omitempty"`
	Mode                        string `json:"mode" tf:"mode"`
}

type CosmosdbGremlinGraphSpecIndexPolicy struct {
	// +optional
	Automatic bool `json:"automatic,omitempty" tf:"automatic,omitempty"`
	// +optional
	ExcludedPaths []string `json:"excludedPaths,omitempty" tf:"excluded_paths,omitempty"`
	// +optional
	IncludedPaths []string `json:"includedPaths,omitempty" tf:"included_paths,omitempty"`
	IndexingMode  string   `json:"indexingMode" tf:"indexing_mode"`
}

type CosmosdbGremlinGraphSpecUniqueKey struct {
	Paths []string `json:"paths" tf:"paths"`
}

type CosmosdbGremlinGraphSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AccountName              string                                             `json:"accountName" tf:"account_name"`
	ConflictResolutionPolicy []CosmosdbGremlinGraphSpecConflictResolutionPolicy `json:"conflictResolutionPolicy" tf:"conflict_resolution_policy"`
	DatabaseName             string                                             `json:"databaseName" tf:"database_name"`
	IndexPolicy              []CosmosdbGremlinGraphSpecIndexPolicy              `json:"indexPolicy" tf:"index_policy"`
	Name                     string                                             `json:"name" tf:"name"`
	// +optional
	PartitionKeyPath  string `json:"partitionKeyPath,omitempty" tf:"partition_key_path,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Throughput int64 `json:"throughput,omitempty" tf:"throughput,omitempty"`
	// +optional
	UniqueKey []CosmosdbGremlinGraphSpecUniqueKey `json:"uniqueKey,omitempty" tf:"unique_key,omitempty"`
}

type CosmosdbGremlinGraphStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CosmosdbGremlinGraphSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
	// +optional
	TerraFormState string `json:"terraformState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CosmosdbGremlinGraphList is a list of CosmosdbGremlinGraphs
type CosmosdbGremlinGraphList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CosmosdbGremlinGraph CRD objects
	Items []CosmosdbGremlinGraph `json:"items,omitempty"`
}
