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

type SearchService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SearchServiceSpec   `json:"spec,omitempty"`
	Status            SearchServiceStatus `json:"status,omitempty"`
}

type SearchServiceSpecQueryKeys struct {
	// +optional
	Key string `json:"key,omitempty" tf:"key,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
}

type SearchServiceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Location string `json:"location" tf:"location"`
	Name     string `json:"name" tf:"name"`
	// +optional
	PartitionCount int64 `json:"partitionCount,omitempty" tf:"partition_count,omitempty"`
	// +optional
	PrimaryKey string `json:"primaryKey,omitempty" tf:"primary_key,omitempty"`
	// +optional
	QueryKeys []SearchServiceSpecQueryKeys `json:"queryKeys,omitempty" tf:"query_keys,omitempty"`
	// +optional
	ReplicaCount      int64  `json:"replicaCount,omitempty" tf:"replica_count,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SecondaryKey string `json:"secondaryKey,omitempty" tf:"secondary_key,omitempty"`
	Sku          string `json:"sku" tf:"sku"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SearchServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SearchServiceSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
	// +optional
	TerraformErrors []string `json:"terraformErrors,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SearchServiceList is a list of SearchServices
type SearchServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SearchService CRD objects
	Items []SearchService `json:"items,omitempty"`
}
