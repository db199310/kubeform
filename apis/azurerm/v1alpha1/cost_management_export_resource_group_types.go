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

type CostManagementExportResourceGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CostManagementExportResourceGroupSpec   `json:"spec,omitempty"`
	Status            CostManagementExportResourceGroupStatus `json:"status,omitempty"`
}

type CostManagementExportResourceGroupSpecDeliveryInfo struct {
	ContainerName    string `json:"containerName" tf:"container_name"`
	RootFolderPath   string `json:"rootFolderPath" tf:"root_folder_path"`
	StorageAccountID string `json:"storageAccountID" tf:"storage_account_id"`
}

type CostManagementExportResourceGroupSpecQuery struct {
	TimeFrame string `json:"timeFrame" tf:"time_frame"`
	Type      string `json:"type" tf:"type"`
}

type CostManagementExportResourceGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Active bool `json:"active,omitempty" tf:"active,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	DeliveryInfo []CostManagementExportResourceGroupSpecDeliveryInfo `json:"deliveryInfo" tf:"delivery_info"`
	Name         string                                              `json:"name" tf:"name"`
	// +kubebuilder:validation:MaxItems=1
	Query                 []CostManagementExportResourceGroupSpecQuery `json:"query" tf:"query"`
	RecurrencePeriodEnd   string                                       `json:"recurrencePeriodEnd" tf:"recurrence_period_end"`
	RecurrencePeriodStart string                                       `json:"recurrencePeriodStart" tf:"recurrence_period_start"`
	RecurrenceType        string                                       `json:"recurrenceType" tf:"recurrence_type"`
	ResourceGroupID       string                                       `json:"resourceGroupID" tf:"resource_group_id"`
}

type CostManagementExportResourceGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CostManagementExportResourceGroupSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
	// +optional
	TerraFormState string `json:"terraformState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CostManagementExportResourceGroupList is a list of CostManagementExportResourceGroups
type CostManagementExportResourceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CostManagementExportResourceGroup CRD objects
	Items []CostManagementExportResourceGroup `json:"items,omitempty"`
}
