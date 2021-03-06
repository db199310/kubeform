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

type AppServiceCertificateOrder struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppServiceCertificateOrderSpec   `json:"spec,omitempty"`
	Status            AppServiceCertificateOrderStatus `json:"status,omitempty"`
}

type AppServiceCertificateOrderSpecCertificates struct {
	// +optional
	CertificateName string `json:"certificateName,omitempty" tf:"certificate_name,omitempty"`
	// +optional
	KeyVaultID string `json:"keyVaultID,omitempty" tf:"key_vault_id,omitempty"`
	// +optional
	KeyVaultSecretName string `json:"keyVaultSecretName,omitempty" tf:"key_vault_secret_name,omitempty"`
	// +optional
	ProvisioningState string `json:"provisioningState,omitempty" tf:"provisioning_state,omitempty"`
}

type AppServiceCertificateOrderSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AppServiceCertificateNotRenewableReasons []string `json:"appServiceCertificateNotRenewableReasons,omitempty" tf:"app_service_certificate_not_renewable_reasons,omitempty"`
	// +optional
	AutoRenew bool `json:"autoRenew,omitempty" tf:"auto_renew,omitempty"`
	// +optional
	Certificates []AppServiceCertificateOrderSpecCertificates `json:"certificates,omitempty" tf:"certificates,omitempty"`
	// +optional
	Csr string `json:"csr,omitempty" tf:"csr,omitempty"`
	// +optional
	DistinguishedName string `json:"distinguishedName,omitempty" tf:"distinguished_name,omitempty"`
	// +optional
	DomainVerificationToken string `json:"domainVerificationToken,omitempty" tf:"domain_verification_token,omitempty"`
	// +optional
	ExpirationTime string `json:"expirationTime,omitempty" tf:"expiration_time,omitempty"`
	// +optional
	IntermediateThumbprint string `json:"intermediateThumbprint,omitempty" tf:"intermediate_thumbprint,omitempty"`
	// +optional
	IsPrivateKeyExternal bool `json:"isPrivateKeyExternal,omitempty" tf:"is_private_key_external,omitempty"`
	// +optional
	KeySize  int64  `json:"keySize,omitempty" tf:"key_size,omitempty"`
	Location string `json:"location" tf:"location"`
	Name     string `json:"name" tf:"name"`
	// +optional
	ProductType       string `json:"productType,omitempty" tf:"product_type,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	RootThumbprint string `json:"rootThumbprint,omitempty" tf:"root_thumbprint,omitempty"`
	// +optional
	SignedCertificateThumbprint string `json:"signedCertificateThumbprint,omitempty" tf:"signed_certificate_thumbprint,omitempty"`
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	ValidityInYears int64 `json:"validityInYears,omitempty" tf:"validity_in_years,omitempty"`
}

type AppServiceCertificateOrderStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AppServiceCertificateOrderSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
	// +optional
	TerraformErrors []string `json:"terraformErrors,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppServiceCertificateOrderList is a list of AppServiceCertificateOrders
type AppServiceCertificateOrderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppServiceCertificateOrder CRD objects
	Items []AppServiceCertificateOrder `json:"items,omitempty"`
}
