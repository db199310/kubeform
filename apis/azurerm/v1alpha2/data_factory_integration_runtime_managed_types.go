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

type DataFactoryIntegrationRuntimeManaged struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataFactoryIntegrationRuntimeManagedSpec   `json:"spec,omitempty"`
	Status            DataFactoryIntegrationRuntimeManagedStatus `json:"status,omitempty"`
}

type DataFactoryIntegrationRuntimeManagedSpecCatalogInfo struct {
	AdministratorLogin    string `json:"administratorLogin" tf:"administrator_login"`
	AdministratorPassword string `json:"-" sensitive:"true" tf:"administrator_password"`
	// +optional
	PricingTier    string `json:"pricingTier,omitempty" tf:"pricing_tier,omitempty"`
	ServerEndpoint string `json:"serverEndpoint" tf:"server_endpoint"`
}

type DataFactoryIntegrationRuntimeManagedSpecCustomSetupScript struct {
	BlobContainerURI string `json:"blobContainerURI" tf:"blob_container_uri"`
	SasToken         string `json:"-" sensitive:"true" tf:"sas_token"`
}

type DataFactoryIntegrationRuntimeManagedSpecVnetIntegration struct {
	SubnetName string `json:"subnetName" tf:"subnet_name"`
	VnetID     string `json:"vnetID" tf:"vnet_id"`
}

type DataFactoryIntegrationRuntimeManagedSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	CatalogInfo []DataFactoryIntegrationRuntimeManagedSpecCatalogInfo `json:"catalogInfo,omitempty" tf:"catalog_info,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CustomSetupScript []DataFactoryIntegrationRuntimeManagedSpecCustomSetupScript `json:"customSetupScript,omitempty" tf:"custom_setup_script,omitempty"`
	DataFactoryName   string                                                      `json:"dataFactoryName" tf:"data_factory_name"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Edition string `json:"edition,omitempty" tf:"edition,omitempty"`
	// +optional
	LicenseType string `json:"licenseType,omitempty" tf:"license_type,omitempty"`
	Location    string `json:"location" tf:"location"`
	// +optional
	MaxParallelExecutionsPerNode int64  `json:"maxParallelExecutionsPerNode,omitempty" tf:"max_parallel_executions_per_node,omitempty"`
	Name                         string `json:"name" tf:"name"`
	NodeSize                     string `json:"nodeSize" tf:"node_size"`
	// +optional
	NumberOfNodes     int64  `json:"numberOfNodes,omitempty" tf:"number_of_nodes,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	VnetIntegration []DataFactoryIntegrationRuntimeManagedSpecVnetIntegration `json:"vnetIntegration,omitempty" tf:"vnet_integration,omitempty"`
}

type DataFactoryIntegrationRuntimeManagedStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DataFactoryIntegrationRuntimeManagedSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DataFactoryIntegrationRuntimeManagedList is a list of DataFactoryIntegrationRuntimeManageds
type DataFactoryIntegrationRuntimeManagedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DataFactoryIntegrationRuntimeManaged CRD objects
	Items []DataFactoryIntegrationRuntimeManaged `json:"items,omitempty"`
}
