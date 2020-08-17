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
	"encoding/json"

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

type F4dpAzStgv1 struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              F4dpAzStgv1Spec   `json:"spec,omitempty"`
	Status            F4dpAzStgv1Status `json:"status,omitempty"`
}

type F4dpAzStgv1Spec struct {
	// +optional
	SecretRef   *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
	ProviderRef core.LocalObjectReference  `json:"providerRef" tf:"-"`
	// +optional
	Source string `json:"source" tf:"source"`

	// +optional
	// Account Kind
	AccountKind json.RawMessage `json:"accountKind,omitempty" tf:"account_kind,omitempty"`
	// +optional
	// Account tier default is Standard
	AccountTier string `json:"accountTier,omitempty" tf:"account_tier,omitempty"`
	// +optional
	// List of Blobs
	Blobs []F4dpAzStgv1Blobs `json:"blobs,omitempty" tf:"blobs,omitempty"`
	// +optional
	// List of services to bypass network rules
	Bypass []string `json:"bypass,omitempty" tf:"bypass,omitempty"`
	// +optional
	// The name of the resource group in which to create the storage account. Changing this forces a new resource to be created.
	Containers []string `json:"containers,omitempty" tf:"containers,omitempty"`
	// +optional
	// Ip Rules
	IpRules []string `json:"ipRules,omitempty" tf:"ip_rules,omitempty"`
	// +optional
	IsHnsEnabled *bool `json:"isHnsEnabled,omitempty" tf:"is_hns_enabled,omitempty"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location string `json:"location" tf:"location"`
	// +optional
	// List of Storage Queues
	Queues []string `json:"queues,omitempty" tf:"queues,omitempty"`
	// +optional
	// Sepcify replication type default is LRS
	ReplicationType string `json:"replicationType,omitempty" tf:"replication_type,omitempty"`
	// The name of the resource group in which to create the storage account. Changing this forces a new resource to be created.
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	// List of Storage shares
	Shares []F4dpAzStgv1Shares `json:"shares,omitempty" tf:"shares,omitempty"`
	// Specifies the storage account in which to create the storage container. Changing this forces a new resource to be created.
	StorageAccountName string `json:"storageAccountName" tf:"storage_account_name"`
	// +optional
	// The vnet subnet id
	SubnetID []json.RawMessage `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
	// +optional
	// List of storage tables
	Tables []string `json:"tables,omitempty" tf:"tables,omitempty"`
	// +optional
	// The tags to associate with assets
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type F4dpAzStgv1Blobs struct {
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
}

type F4dpAzStgv1Shares struct {
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	Quota json.Number `json:"quota,omitempty" tf:"quota,omitempty"`
}

type F4dpAzStgv1Output struct {
	//
	// +optional
	StorageAccountContainerName string `json:"storageAccountContainerName" tf:"storage_account_container_name"`
	//
	// +optional
	StorageAccountID string `json:"storageAccountID" tf:"storage_account_id"`
	//
	// +optional
	StorageAccountName string `json:"storageAccountName" tf:"storage_account_name"`
}

type F4dpAzStgv1Status struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *F4dpAzStgv1Output `json:"output,omitempty"`
	// +optional
	State string `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// F4dpAzStgv1List is a list of F4dpAzStgv1s
type F4dpAzStgv1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of F4dpAzStgv1 CRD objects
	Items []F4dpAzStgv1 `json:"items,omitempty"`
}
