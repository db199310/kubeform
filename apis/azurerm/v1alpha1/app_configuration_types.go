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

type AppConfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppConfigurationSpec   `json:"spec,omitempty"`
	Status            AppConfigurationStatus `json:"status,omitempty"`
}

type AppConfigurationSpecPrimaryReadKey struct {
	// +optional
	ConnectionString string `json:"-" sensitive:"true" tf:"connection_string,omitempty"`
	// +optional
	ID string `json:"-" sensitive:"true" tf:"id,omitempty"`
	// +optional
	Secret string `json:"-" sensitive:"true" tf:"secret,omitempty"`
}

type AppConfigurationSpecPrimaryWriteKey struct {
	// +optional
	ConnectionString string `json:"-" sensitive:"true" tf:"connection_string,omitempty"`
	// +optional
	ID string `json:"-" sensitive:"true" tf:"id,omitempty"`
	// +optional
	Secret string `json:"-" sensitive:"true" tf:"secret,omitempty"`
}

type AppConfigurationSpecSecondaryReadKey struct {
	// +optional
	ConnectionString string `json:"-" sensitive:"true" tf:"connection_string,omitempty"`
	// +optional
	ID string `json:"-" sensitive:"true" tf:"id,omitempty"`
	// +optional
	Secret string `json:"-" sensitive:"true" tf:"secret,omitempty"`
}

type AppConfigurationSpecSecondaryWriteKey struct {
	// +optional
	ConnectionString string `json:"-" sensitive:"true" tf:"connection_string,omitempty"`
	// +optional
	ID string `json:"-" sensitive:"true" tf:"id,omitempty"`
	// +optional
	Secret string `json:"-" sensitive:"true" tf:"secret,omitempty"`
}

type AppConfigurationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	Endpoint string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`
	Location string `json:"location" tf:"location"`
	Name     string `json:"name" tf:"name"`
	// +optional
	PrimaryReadKey []AppConfigurationSpecPrimaryReadKey `json:"primaryReadKey,omitempty" tf:"primary_read_key,omitempty"`
	// +optional
	PrimaryWriteKey   []AppConfigurationSpecPrimaryWriteKey `json:"primaryWriteKey,omitempty" tf:"primary_write_key,omitempty"`
	ResourceGroupName string                                `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SecondaryReadKey []AppConfigurationSpecSecondaryReadKey `json:"secondaryReadKey,omitempty" tf:"secondary_read_key,omitempty"`
	// +optional
	SecondaryWriteKey []AppConfigurationSpecSecondaryWriteKey `json:"secondaryWriteKey,omitempty" tf:"secondary_write_key,omitempty"`
	// +optional
	Sku string `json:"sku,omitempty" tf:"sku,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AppConfigurationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AppConfigurationSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppConfigurationList is a list of AppConfigurations
type AppConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppConfiguration CRD objects
	Items []AppConfiguration `json:"items,omitempty"`
}
