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

type PrivateEndpoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateEndpointSpec   `json:"spec,omitempty"`
	Status            PrivateEndpointStatus `json:"status,omitempty"`
}

type PrivateEndpointSpecPrivateServiceConnection struct {
	IsManualConnection          bool   `json:"isManualConnection" tf:"is_manual_connection"`
	Name                        string `json:"name" tf:"name"`
	PrivateConnectionResourceID string `json:"privateConnectionResourceID" tf:"private_connection_resource_id"`
	// +optional
	PrivateIPAddress string `json:"privateIPAddress,omitempty" tf:"private_ip_address,omitempty"`
	// +optional
	RequestMessage string `json:"requestMessage,omitempty" tf:"request_message,omitempty"`
	// +optional
	SubresourceNames []string `json:"subresourceNames,omitempty" tf:"subresource_names,omitempty"`
}

type PrivateEndpointSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Location string `json:"location" tf:"location"`
	Name     string `json:"name" tf:"name"`
	// +kubebuilder:validation:MaxItems=1
	PrivateServiceConnection []PrivateEndpointSpecPrivateServiceConnection `json:"privateServiceConnection" tf:"private_service_connection"`
	ResourceGroupName        string                                        `json:"resourceGroupName" tf:"resource_group_name"`
	SubnetID                 string                                        `json:"subnetID" tf:"subnet_id"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PrivateEndpointStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *PrivateEndpointSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PrivateEndpointList is a list of PrivateEndpoints
type PrivateEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PrivateEndpoint CRD objects
	Items []PrivateEndpoint `json:"items,omitempty"`
}
