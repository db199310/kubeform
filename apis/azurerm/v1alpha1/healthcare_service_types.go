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
// +kubebuilder:storageversion

type HealthcareService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HealthcareServiceSpec   `json:"spec,omitempty"`
	Status            HealthcareServiceStatus `json:"status,omitempty"`
}

type HealthcareServiceSpecAuthenticationConfiguration struct {
	// +optional
	Audience string `json:"audience,omitempty" tf:"audience,omitempty"`
	// +optional
	Authority string `json:"authority,omitempty" tf:"authority,omitempty"`
	// +optional
	SmartProxyEnabled bool `json:"smartProxyEnabled,omitempty" tf:"smart_proxy_enabled,omitempty"`
}

type HealthcareServiceSpecCorsConfiguration struct {
	// +optional
	AllowCredentials bool `json:"allowCredentials,omitempty" tf:"allow_credentials,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=64
	AllowedHeaders []string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=64
	AllowedMethods []string `json:"allowedMethods,omitempty" tf:"allowed_methods,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=64
	AllowedOrigins []string `json:"allowedOrigins,omitempty" tf:"allowed_origins,omitempty"`
	// +optional
	MaxAgeInSeconds int64 `json:"maxAgeInSeconds,omitempty" tf:"max_age_in_seconds,omitempty"`
}

type HealthcareServiceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:MinItems=1
	AccessPolicyObjectIDS []string `json:"accessPolicyObjectIDS" tf:"access_policy_object_ids"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AuthenticationConfiguration []HealthcareServiceSpecAuthenticationConfiguration `json:"authenticationConfiguration,omitempty" tf:"authentication_configuration,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CorsConfiguration []HealthcareServiceSpecCorsConfiguration `json:"corsConfiguration,omitempty" tf:"cors_configuration,omitempty"`
	// +optional
	CosmosdbThroughput int64 `json:"cosmosdbThroughput,omitempty" tf:"cosmosdb_throughput,omitempty"`
	// +optional
	Kind              string `json:"kind,omitempty" tf:"kind,omitempty"`
	Location          string `json:"location" tf:"location"`
	Name              string `json:"name" tf:"name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type HealthcareServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *HealthcareServiceSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase           base.Phase `json:"phase,omitempty"`
	TerraformErrors []string   `json:"terraformErrors,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// HealthcareServiceList is a list of HealthcareServices
type HealthcareServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of HealthcareService CRD objects
	Items []HealthcareService `json:"items,omitempty"`
}
