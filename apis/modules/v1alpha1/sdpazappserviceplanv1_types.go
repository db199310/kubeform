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

type SDPAzappserviceplanv1 struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SDPAzappserviceplanv1Spec   `json:"spec,omitempty"`
	Status            SDPAzappserviceplanv1Status `json:"status,omitempty"`
}

type SDPAzappserviceplanv1Spec struct {
	// +optional
	SecretRef   *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
	ProviderRef core.LocalObjectReference  `json:"providerRef" tf:"-"`
	// +optional
	Source string `json:"source" tf:"source"`

	// Client name/account used in naming
	ClientName string `json:"clientName" tf:"client_name"`
	// +optional
	// Name of the App Service Plan, generated if not set.
	CustomName string `json:"customName,omitempty" tf:"custom_name,omitempty"`
	// Project environment
	Environment string `json:"environment" tf:"environment"`
	// +optional
	// Extra tags to add
	ExtraTags map[string]string `json:"extraTags,omitempty" tf:"extra_tags,omitempty"`
	// The kind of the App Service Plan to create. See documentation https://www.terraform.io/docs/providers/azurerm/r/app_service_plan.html#kind
	Kind string `json:"kind" tf:"kind"`
	// Azure location.
	Location string `json:"location" tf:"location"`
	// Short string for Azure location.
	LocationShort string `json:"locationShort" tf:"location_short"`
	// +optional
	// Optional prefix for the generated name
	NamePrefix string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`
	// owner
	Owner string `json:"owner" tf:"owner"`
	// +optional
	// placement
	Placement string `json:"placement,omitempty" tf:"placement,omitempty"`
	// project stream name
	ProjectStream string `json:"projectStream" tf:"projectStream"`
	// region
	Region string `json:"region" tf:"region"`
	// +optional
	// releaseVersion
	ReleaseVersion string `json:"releaseVersion,omitempty" tf:"releaseVersion,omitempty"`
	// +optional
	// Flag indicating if App Service Plan should be reserved. Forced to true if "kind" is "Linux".
	Reserved string `json:"reserved,omitempty" tf:"reserved,omitempty"`
	// Resource group name
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// A sku block. See documentation https://www.terraform.io/docs/providers/azurerm/r/app_service_plan.html#sku
	Sku map[string]string `json:"sku" tf:"sku"`
	// Project stack name
	Stack string `json:"stack" tf:"stack"`
	//  4 character project stream name/code
	WorkStream string `json:"workStream" tf:"workStream"`
}

type SDPAzappserviceplanv1Output struct {
	// Id of the created App Service Plan
	// +optional
	AppServicePlanID string `json:"appServicePlanID" tf:"app_service_plan_id"`
	// Azure location of the created App Service Plan
	// +optional
	AppServicePlanLocation string `json:"appServicePlanLocation" tf:"app_service_plan_location"`
	// Maximum number of workers for the created App Service Plan
	// +optional
	AppServicePlanMaxWorkers string `json:"appServicePlanMaxWorkers" tf:"app_service_plan_max_workers"`
	// Name of the created App Service Plan
	// +optional
	AppServicePlanName string `json:"appServicePlanName" tf:"app_service_plan_name"`
}

type SDPAzappserviceplanv1Status struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SDPAzappserviceplanv1Output `json:"output,omitempty"`
	// +optional
	State string `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SDPAzappserviceplanv1List is a list of SDPAzappserviceplanv1s
type SDPAzappserviceplanv1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SDPAzappserviceplanv1 CRD objects
	Items []SDPAzappserviceplanv1 `json:"items,omitempty"`
}
