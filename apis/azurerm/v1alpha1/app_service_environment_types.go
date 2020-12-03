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

type AppServiceEnvironment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppServiceEnvironmentSpec   `json:"spec,omitempty"`
	Status            AppServiceEnvironmentStatus `json:"status,omitempty"`
}

type AppServiceEnvironmentSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	FrontEndScaleFactor int64 `json:"frontEndScaleFactor,omitempty" tf:"front_end_scale_factor,omitempty"`
	// +optional
	InternalLoadBalancingMode string `json:"internalLoadBalancingMode,omitempty" tf:"internal_load_balancing_mode,omitempty"`
	// +optional
	Location string `json:"location,omitempty" tf:"location,omitempty"`
	Name     string `json:"name" tf:"name"`
	// +optional
	PricingTier string `json:"pricingTier,omitempty" tf:"pricing_tier,omitempty"`
	// +optional
	ResourceGroupName string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`
	SubnetID          string `json:"subnetID" tf:"subnet_id"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AppServiceEnvironmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AppServiceEnvironmentSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
	// +optional
	TerraFormLogs *base.TerraFormLogs `json:"terraformLogs,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppServiceEnvironmentList is a list of AppServiceEnvironments
type AppServiceEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppServiceEnvironment CRD objects
	Items []AppServiceEnvironment `json:"items,omitempty"`
}
