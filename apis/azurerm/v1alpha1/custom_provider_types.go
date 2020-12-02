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

type CustomProvider struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CustomProviderSpec   `json:"spec,omitempty"`
	Status            CustomProviderStatus `json:"status,omitempty"`
}

type CustomProviderSpecAction struct {
	Endpoint string `json:"endpoint" tf:"endpoint"`
	Name     string `json:"name" tf:"name"`
}

type CustomProviderSpecResourceType struct {
	Endpoint string `json:"endpoint" tf:"endpoint"`
	Name     string `json:"name" tf:"name"`
	// +optional
	RoutingType string `json:"routingType,omitempty" tf:"routing_type,omitempty"`
}

type CustomProviderSpecValidation struct {
	Specification string `json:"specification" tf:"specification"`
}

type CustomProviderSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Action            []CustomProviderSpecAction `json:"action,omitempty" tf:"action,omitempty"`
	Location          string                     `json:"location" tf:"location"`
	Name              string                     `json:"name" tf:"name"`
	ResourceGroupName string                     `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	ResourceType []CustomProviderSpecResourceType `json:"resourceType,omitempty" tf:"resource_type,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	Validation []CustomProviderSpecValidation `json:"validation,omitempty" tf:"validation,omitempty"`
}

type CustomProviderStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CustomProviderSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
	// +optional
	TerraFormState string `json:"terraformState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CustomProviderList is a list of CustomProviders
type CustomProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CustomProvider CRD objects
	Items []CustomProvider `json:"items,omitempty"`
}
