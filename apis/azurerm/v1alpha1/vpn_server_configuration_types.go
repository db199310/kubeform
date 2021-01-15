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

type VpnServerConfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpnServerConfigurationSpec   `json:"spec,omitempty"`
	Status            VpnServerConfigurationStatus `json:"status,omitempty"`
}

type VpnServerConfigurationSpecAzureActiveDirectoryAuthentication struct {
	Audience string `json:"audience" tf:"audience"`
	Issuer   string `json:"issuer" tf:"issuer"`
	Tenant   string `json:"tenant" tf:"tenant"`
}

type VpnServerConfigurationSpecClientRevokedCertificate struct {
	Name       string `json:"name" tf:"name"`
	Thumbprint string `json:"thumbprint" tf:"thumbprint"`
}

type VpnServerConfigurationSpecClientRootCertificate struct {
	Name           string `json:"name" tf:"name"`
	PublicCertData string `json:"publicCertData" tf:"public_cert_data"`
}

type VpnServerConfigurationSpecIpsecPolicy struct {
	DhGroup             string `json:"dhGroup" tf:"dh_group"`
	IkeEncryption       string `json:"ikeEncryption" tf:"ike_encryption"`
	IkeIntegrity        string `json:"ikeIntegrity" tf:"ike_integrity"`
	IpsecEncryption     string `json:"ipsecEncryption" tf:"ipsec_encryption"`
	IpsecIntegrity      string `json:"ipsecIntegrity" tf:"ipsec_integrity"`
	PfsGroup            string `json:"pfsGroup" tf:"pfs_group"`
	SaDataSizeKilobytes int64  `json:"saDataSizeKilobytes" tf:"sa_data_size_kilobytes"`
	SaLifetimeSeconds   int64  `json:"saLifetimeSeconds" tf:"sa_lifetime_seconds"`
}

type VpnServerConfigurationSpecRadiusServerClientRootCertificate struct {
	Name       string `json:"name" tf:"name"`
	Thumbprint string `json:"thumbprint" tf:"thumbprint"`
}

type VpnServerConfigurationSpecRadiusServerServerRootCertificate struct {
	Name           string `json:"name" tf:"name"`
	PublicCertData string `json:"publicCertData" tf:"public_cert_data"`
}

type VpnServerConfigurationSpecRadiusServer struct {
	Address string `json:"address" tf:"address"`
	// +optional
	ClientRootCertificate []VpnServerConfigurationSpecRadiusServerClientRootCertificate `json:"clientRootCertificate,omitempty" tf:"client_root_certificate,omitempty"`
	Secret                string                                                        `json:"-" sensitive:"true" tf:"secret"`
	ServerRootCertificate []VpnServerConfigurationSpecRadiusServerServerRootCertificate `json:"serverRootCertificate" tf:"server_root_certificate"`
}

type VpnServerConfigurationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	// +kubebuilder:validation:MinItems=1
	AzureActiveDirectoryAuthentication []VpnServerConfigurationSpecAzureActiveDirectoryAuthentication `json:"azureActiveDirectoryAuthentication,omitempty" tf:"azure_active_directory_authentication,omitempty"`
	// +optional
	ClientRevokedCertificate []VpnServerConfigurationSpecClientRevokedCertificate `json:"clientRevokedCertificate,omitempty" tf:"client_revoked_certificate,omitempty"`
	// +optional
	ClientRootCertificate []VpnServerConfigurationSpecClientRootCertificate `json:"clientRootCertificate,omitempty" tf:"client_root_certificate,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	IpsecPolicy []VpnServerConfigurationSpecIpsecPolicy `json:"ipsecPolicy,omitempty" tf:"ipsec_policy,omitempty"`
	Location    string                                  `json:"location" tf:"location"`
	Name        string                                  `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RadiusServer      []VpnServerConfigurationSpecRadiusServer `json:"radiusServer,omitempty" tf:"radius_server,omitempty"`
	ResourceGroupName string                                   `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	VpnAuthenticationTypes []string `json:"vpnAuthenticationTypes" tf:"vpn_authentication_types"`
	// +optional
	VpnProtocols []string `json:"vpnProtocols,omitempty" tf:"vpn_protocols,omitempty"`
}

type VpnServerConfigurationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *VpnServerConfigurationSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VpnServerConfigurationList is a list of VpnServerConfigurations
type VpnServerConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpnServerConfiguration CRD objects
	Items []VpnServerConfiguration `json:"items,omitempty"`
}
