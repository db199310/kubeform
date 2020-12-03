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

type PrivateLinkService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateLinkServiceSpec   `json:"spec,omitempty"`
	Status            PrivateLinkServiceStatus `json:"status,omitempty"`
}

type PrivateLinkServiceSpecNatIPConfiguration struct {
	Name    string `json:"name" tf:"name"`
	Primary bool   `json:"primary" tf:"primary"`
	// +optional
	PrivateIPAddress string `json:"privateIPAddress,omitempty" tf:"private_ip_address,omitempty"`
	// +optional
	PrivateIPAddressVersion string `json:"privateIPAddressVersion,omitempty" tf:"private_ip_address_version,omitempty"`
	SubnetID                string `json:"subnetID" tf:"subnet_id"`
}

type PrivateLinkServiceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Alias string `json:"alias,omitempty" tf:"alias,omitempty"`
	// +optional
	AutoApprovalSubscriptionIDS []string `json:"autoApprovalSubscriptionIDS,omitempty" tf:"auto_approval_subscription_ids,omitempty"`
	// +optional
	EnableProxyProtocol                    bool     `json:"enableProxyProtocol,omitempty" tf:"enable_proxy_protocol,omitempty"`
	LoadBalancerFrontendIPConfigurationIDS []string `json:"loadBalancerFrontendIPConfigurationIDS" tf:"load_balancer_frontend_ip_configuration_ids"`
	Location                               string   `json:"location" tf:"location"`
	Name                                   string   `json:"name" tf:"name"`
	// +kubebuilder:validation:MaxItems=8
	NatIPConfiguration []PrivateLinkServiceSpecNatIPConfiguration `json:"natIPConfiguration" tf:"nat_ip_configuration"`
	ResourceGroupName  string                                     `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	VisibilitySubscriptionIDS []string `json:"visibilitySubscriptionIDS,omitempty" tf:"visibility_subscription_ids,omitempty"`
}

type PrivateLinkServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *PrivateLinkServiceSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
	// +optional
	TerraFormLogs *base.TerraFormLogs `json:"terraformLogs,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PrivateLinkServiceList is a list of PrivateLinkServices
type PrivateLinkServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PrivateLinkService CRD objects
	Items []PrivateLinkService `json:"items,omitempty"`
}
