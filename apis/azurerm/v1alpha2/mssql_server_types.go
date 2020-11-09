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

package v1alpha2

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha2"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type MssqlServer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MssqlServerSpec   `json:"spec,omitempty"`
	Status            MssqlServerStatus `json:"status,omitempty"`
}

type MssqlServerSpecAzureadAdministrator struct {
	LoginUsername string `json:"loginUsername" tf:"login_username"`
	ObjectID      string `json:"objectID" tf:"object_id"`
	// +optional
	TenantID string `json:"tenantID,omitempty" tf:"tenant_id,omitempty"`
}

type MssqlServerSpecExtendedAuditingPolicy struct {
	// +optional
	RetentionInDays         int64  `json:"retentionInDays,omitempty" tf:"retention_in_days,omitempty"`
	StorageAccountAccessKey string `json:"-" sensitive:"true" tf:"storage_account_access_key"`
	// +optional
	StorageAccountAccessKeyIsSecondary bool   `json:"storageAccountAccessKeyIsSecondary,omitempty" tf:"storage_account_access_key_is_secondary,omitempty"`
	StorageEndpoint                    string `json:"storageEndpoint" tf:"storage_endpoint"`
}

type MssqlServerSpecIdentity struct {
	// +optional
	PrincipalID string `json:"principalID,omitempty" tf:"principal_id,omitempty"`
	// +optional
	TenantID string `json:"tenantID,omitempty" tf:"tenant_id,omitempty"`
	Type     string `json:"type" tf:"type"`
}

type MssqlServerSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	AdministratorLogin         string `json:"administratorLogin" tf:"administrator_login"`
	AdministratorLoginPassword string `json:"-" sensitive:"true" tf:"administrator_login_password"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	AzureadAdministrator []MssqlServerSpecAzureadAdministrator `json:"azureadAdministrator,omitempty" tf:"azuread_administrator,omitempty"`
	// +optional
	ConnectionPolicy string `json:"connectionPolicy,omitempty" tf:"connection_policy,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ExtendedAuditingPolicy []MssqlServerSpecExtendedAuditingPolicy `json:"extendedAuditingPolicy,omitempty" tf:"extended_auditing_policy,omitempty"`
	// +optional
	FullyQualifiedDomainName string `json:"fullyQualifiedDomainName,omitempty" tf:"fully_qualified_domain_name,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Identity []MssqlServerSpecIdentity `json:"identity,omitempty" tf:"identity,omitempty"`
	Location string                    `json:"location" tf:"location"`
	Name     string                    `json:"name" tf:"name"`
	// +optional
	PublicNetworkAccessEnabled bool   `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`
	ResourceGroupName          string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags    map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	Version string            `json:"version" tf:"version"`
}

type MssqlServerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *MssqlServerSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MssqlServerList is a list of MssqlServers
type MssqlServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MssqlServer CRD objects
	Items []MssqlServer `json:"items,omitempty"`
}
