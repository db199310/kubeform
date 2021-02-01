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

type DevTestLinuxVirtualMachine struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DevTestLinuxVirtualMachineSpec   `json:"spec,omitempty"`
	Status            DevTestLinuxVirtualMachineStatus `json:"status,omitempty"`
}

type DevTestLinuxVirtualMachineSpecGalleryImageReference struct {
	Offer     string `json:"offer" tf:"offer"`
	Publisher string `json:"publisher" tf:"publisher"`
	Sku       string `json:"sku" tf:"sku"`
	Version   string `json:"version" tf:"version"`
}

type DevTestLinuxVirtualMachineSpecInboundNATRule struct {
	BackendPort int64 `json:"backendPort" tf:"backend_port"`
	// +optional
	FrontendPort int64  `json:"frontendPort,omitempty" tf:"frontend_port,omitempty"`
	Protocol     string `json:"protocol" tf:"protocol"`
}

type DevTestLinuxVirtualMachineSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	AllowClaim bool `json:"allowClaim,omitempty" tf:"allow_claim,omitempty"`
	// +optional
	DisallowPublicIPAddress bool `json:"disallowPublicIPAddress,omitempty" tf:"disallow_public_ip_address,omitempty"`
	// +optional
	Fqdn string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	GalleryImageReference []DevTestLinuxVirtualMachineSpecGalleryImageReference `json:"galleryImageReference" tf:"gallery_image_reference"`
	// +optional
	InboundNATRule      []DevTestLinuxVirtualMachineSpecInboundNATRule `json:"inboundNATRule,omitempty" tf:"inbound_nat_rule,omitempty"`
	LabName             string                                         `json:"labName" tf:"lab_name"`
	LabSubnetName       string                                         `json:"labSubnetName" tf:"lab_subnet_name"`
	LabVirtualNetworkID string                                         `json:"labVirtualNetworkID" tf:"lab_virtual_network_id"`
	Location            string                                         `json:"location" tf:"location"`
	Name                string                                         `json:"name" tf:"name"`
	// +optional
	Notes string `json:"notes,omitempty" tf:"notes,omitempty"`
	// +optional
	Password          string `json:"-" sensitive:"true" tf:"password,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	Size              string `json:"size" tf:"size"`
	// +optional
	SshKey      string `json:"sshKey,omitempty" tf:"ssh_key,omitempty"`
	StorageType string `json:"storageType" tf:"storage_type"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	UniqueIdentifier string `json:"uniqueIdentifier,omitempty" tf:"unique_identifier,omitempty"`
	Username         string `json:"username" tf:"username"`
}

type DevTestLinuxVirtualMachineStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DevTestLinuxVirtualMachineSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase           base.Phase `json:"phase,omitempty"`
	TerraformErrors []string   `json:"terraformErrors,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DevTestLinuxVirtualMachineList is a list of DevTestLinuxVirtualMachines
type DevTestLinuxVirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DevTestLinuxVirtualMachine CRD objects
	Items []DevTestLinuxVirtualMachine `json:"items,omitempty"`
}
