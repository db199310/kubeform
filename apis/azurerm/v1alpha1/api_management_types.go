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

type ApiManagement struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementSpec   `json:"spec,omitempty"`
	Status            ApiManagementStatus `json:"status,omitempty"`
}

type ApiManagementSpecAdditionalLocation struct {
	// +optional
	GatewayRegionalURL string `json:"gatewayRegionalURL,omitempty" tf:"gateway_regional_url,omitempty"`
	Location           string `json:"location" tf:"location"`
	// +optional
	PublicIPAddresses []string `json:"publicIPAddresses,omitempty" tf:"public_ip_addresses,omitempty"`
}

type ApiManagementSpecCertificate struct {
	CertificatePassword string `json:"-" sensitive:"true" tf:"certificate_password"`
	EncodedCertificate  string `json:"-" sensitive:"true" tf:"encoded_certificate"`
	StoreName           string `json:"storeName" tf:"store_name"`
}

type ApiManagementSpecHostnameConfigurationDeveloperPortal struct {
	// +optional
	Certificate string `json:"-" sensitive:"true" tf:"certificate,omitempty"`
	// +optional
	CertificatePassword string `json:"-" sensitive:"true" tf:"certificate_password,omitempty"`
	HostName            string `json:"hostName" tf:"host_name"`
	// +optional
	KeyVaultID string `json:"keyVaultID,omitempty" tf:"key_vault_id,omitempty"`
	// +optional
	NegotiateClientCertificate bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty"`
}

type ApiManagementSpecHostnameConfigurationManagement struct {
	// +optional
	Certificate string `json:"-" sensitive:"true" tf:"certificate,omitempty"`
	// +optional
	CertificatePassword string `json:"-" sensitive:"true" tf:"certificate_password,omitempty"`
	HostName            string `json:"hostName" tf:"host_name"`
	// +optional
	KeyVaultID string `json:"keyVaultID,omitempty" tf:"key_vault_id,omitempty"`
	// +optional
	NegotiateClientCertificate bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty"`
}

type ApiManagementSpecHostnameConfigurationPortal struct {
	// +optional
	Certificate string `json:"-" sensitive:"true" tf:"certificate,omitempty"`
	// +optional
	CertificatePassword string `json:"-" sensitive:"true" tf:"certificate_password,omitempty"`
	HostName            string `json:"hostName" tf:"host_name"`
	// +optional
	KeyVaultID string `json:"keyVaultID,omitempty" tf:"key_vault_id,omitempty"`
	// +optional
	NegotiateClientCertificate bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty"`
}

type ApiManagementSpecHostnameConfigurationProxy struct {
	// +optional
	Certificate string `json:"-" sensitive:"true" tf:"certificate,omitempty"`
	// +optional
	CertificatePassword string `json:"-" sensitive:"true" tf:"certificate_password,omitempty"`
	// +optional
	DefaultSSLBinding bool   `json:"defaultSSLBinding,omitempty" tf:"default_ssl_binding,omitempty"`
	HostName          string `json:"hostName" tf:"host_name"`
	// +optional
	KeyVaultID string `json:"keyVaultID,omitempty" tf:"key_vault_id,omitempty"`
	// +optional
	NegotiateClientCertificate bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty"`
}

type ApiManagementSpecHostnameConfigurationScm struct {
	// +optional
	Certificate string `json:"-" sensitive:"true" tf:"certificate,omitempty"`
	// +optional
	CertificatePassword string `json:"-" sensitive:"true" tf:"certificate_password,omitempty"`
	HostName            string `json:"hostName" tf:"host_name"`
	// +optional
	KeyVaultID string `json:"keyVaultID,omitempty" tf:"key_vault_id,omitempty"`
	// +optional
	NegotiateClientCertificate bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty"`
}

type ApiManagementSpecHostnameConfiguration struct {
	// +optional
	DeveloperPortal []ApiManagementSpecHostnameConfigurationDeveloperPortal `json:"developerPortal,omitempty" tf:"developer_portal,omitempty"`
	// +optional
	Management []ApiManagementSpecHostnameConfigurationManagement `json:"management,omitempty" tf:"management,omitempty"`
	// +optional
	Portal []ApiManagementSpecHostnameConfigurationPortal `json:"portal,omitempty" tf:"portal,omitempty"`
	// +optional
	Proxy []ApiManagementSpecHostnameConfigurationProxy `json:"proxy,omitempty" tf:"proxy,omitempty"`
	// +optional
	Scm []ApiManagementSpecHostnameConfigurationScm `json:"scm,omitempty" tf:"scm,omitempty"`
}

type ApiManagementSpecIdentity struct {
	// +optional
	// +kubebuilder:validation:MinItems=1
	IdentityIDS []string `json:"identityIDS,omitempty" tf:"identity_ids,omitempty"`
	// +optional
	PrincipalID string `json:"principalID,omitempty" tf:"principal_id,omitempty"`
	// +optional
	TenantID string `json:"tenantID,omitempty" tf:"tenant_id,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
}

type ApiManagementSpecPolicy struct {
	// +optional
	XmlContent string `json:"xmlContent,omitempty" tf:"xml_content,omitempty"`
	// +optional
	XmlLink string `json:"xmlLink,omitempty" tf:"xml_link,omitempty"`
}

type ApiManagementSpecProtocols struct {
	// +optional
	EnableHttp2 bool `json:"enableHttp2,omitempty" tf:"enable_http2,omitempty"`
}

type ApiManagementSpecSecurity struct {
	// +optional
	EnableBackendSSL30 bool `json:"enableBackendSSL30,omitempty" tf:"enable_backend_ssl30,omitempty"`
	// +optional
	EnableBackendTLS10 bool `json:"enableBackendTLS10,omitempty" tf:"enable_backend_tls10,omitempty"`
	// +optional
	EnableBackendTLS11 bool `json:"enableBackendTLS11,omitempty" tf:"enable_backend_tls11,omitempty"`
	// +optional
	EnableFrontendSSL30 bool `json:"enableFrontendSSL30,omitempty" tf:"enable_frontend_ssl30,omitempty"`
	// +optional
	EnableFrontendTLS10 bool `json:"enableFrontendTLS10,omitempty" tf:"enable_frontend_tls10,omitempty"`
	// +optional
	EnableFrontendTLS11 bool `json:"enableFrontendTLS11,omitempty" tf:"enable_frontend_tls11,omitempty"`
	// +optional
	EnableTripleDESCiphers bool `json:"enableTripleDESCiphers,omitempty" tf:"enable_triple_des_ciphers,omitempty"`
}

type ApiManagementSpecSignIn struct {
	Enabled bool `json:"enabled" tf:"enabled"`
}

type ApiManagementSpecSignUpTermsOfService struct {
	ConsentRequired bool `json:"consentRequired" tf:"consent_required"`
	Enabled         bool `json:"enabled" tf:"enabled"`
	// +optional
	Text string `json:"text,omitempty" tf:"text,omitempty"`
}

type ApiManagementSpecSignUp struct {
	Enabled bool `json:"enabled" tf:"enabled"`
	// +kubebuilder:validation:MaxItems=1
	TermsOfService []ApiManagementSpecSignUpTermsOfService `json:"termsOfService" tf:"terms_of_service"`
}

type ApiManagementSpecVirtualNetworkConfiguration struct {
	SubnetID string `json:"subnetID" tf:"subnet_id"`
}

type ApiManagementSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	AdditionalLocation []ApiManagementSpecAdditionalLocation `json:"additionalLocation,omitempty" tf:"additional_location,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=10
	Certificate []ApiManagementSpecCertificate `json:"certificate,omitempty" tf:"certificate,omitempty"`
	// +optional
	GatewayRegionalURL string `json:"gatewayRegionalURL,omitempty" tf:"gateway_regional_url,omitempty"`
	// +optional
	GatewayURL string `json:"gatewayURL,omitempty" tf:"gateway_url,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HostnameConfiguration []ApiManagementSpecHostnameConfiguration `json:"hostnameConfiguration,omitempty" tf:"hostname_configuration,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Identity []ApiManagementSpecIdentity `json:"identity,omitempty" tf:"identity,omitempty"`
	Location string                      `json:"location" tf:"location"`
	// +optional
	ManagementAPIURL string `json:"managementAPIURL,omitempty" tf:"management_api_url,omitempty"`
	Name             string `json:"name" tf:"name"`
	// +optional
	NotificationSenderEmail string `json:"notificationSenderEmail,omitempty" tf:"notification_sender_email,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Policy []ApiManagementSpecPolicy `json:"policy,omitempty" tf:"policy,omitempty"`
	// +optional
	PortalURL string `json:"portalURL,omitempty" tf:"portal_url,omitempty"`
	// +optional
	PrivateIPAddresses []string `json:"privateIPAddresses,omitempty" tf:"private_ip_addresses,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Protocols []ApiManagementSpecProtocols `json:"protocols,omitempty" tf:"protocols,omitempty"`
	// +optional
	PublicIPAddresses []string `json:"publicIPAddresses,omitempty" tf:"public_ip_addresses,omitempty"`
	PublisherEmail    string   `json:"publisherEmail" tf:"publisher_email"`
	PublisherName     string   `json:"publisherName" tf:"publisher_name"`
	ResourceGroupName string   `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	ScmURL string `json:"scmURL,omitempty" tf:"scm_url,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Security []ApiManagementSpecSecurity `json:"security,omitempty" tf:"security,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SignIn []ApiManagementSpecSignIn `json:"signIn,omitempty" tf:"sign_in,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SignUp  []ApiManagementSpecSignUp `json:"signUp,omitempty" tf:"sign_up,omitempty"`
	SkuName string                    `json:"skuName" tf:"sku_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	VirtualNetworkConfiguration []ApiManagementSpecVirtualNetworkConfiguration `json:"virtualNetworkConfiguration,omitempty" tf:"virtual_network_configuration,omitempty"`
	// +optional
	VirtualNetworkType string `json:"virtualNetworkType,omitempty" tf:"virtual_network_type,omitempty"`
}

type ApiManagementStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ApiManagementSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
	// +optional
	TerraFormLogs *base.TerraFormLogs `json:"terraformLogs,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiManagementList is a list of ApiManagements
type ApiManagementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiManagement CRD objects
	Items []ApiManagement `json:"items,omitempty"`
}
