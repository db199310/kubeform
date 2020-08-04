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

type LinuxVirtualMachine struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LinuxVirtualMachineSpec   `json:"spec,omitempty"`
	Status            LinuxVirtualMachineStatus `json:"status,omitempty"`
}

type LinuxVirtualMachineSpecAdditionalCapabilities struct {
	// +optional
	UltraSsdEnabled bool `json:"ultraSsdEnabled,omitempty" tf:"ultra_ssd_enabled,omitempty"`
}

type LinuxVirtualMachineSpecAdminSSHKey struct {
	PublicKey string `json:"publicKey" tf:"public_key"`
	Username  string `json:"username" tf:"username"`
}

type LinuxVirtualMachineSpecBootDiagnostics struct {
	StorageAccountURI string `json:"storageAccountURI" tf:"storage_account_uri"`
}

type LinuxVirtualMachineSpecIdentity struct {
	// +optional
	IdentityIDS []string `json:"identityIDS,omitempty" tf:"identity_ids,omitempty"`
	// +optional
	PrincipalID string `json:"principalID,omitempty" tf:"principal_id,omitempty"`
	// +optional
	TenantID string `json:"tenantID,omitempty" tf:"tenant_id,omitempty"`
	Type     string `json:"type" tf:"type"`
}

type LinuxVirtualMachineSpecOsDiskDiffDiskSettings struct {
	Option string `json:"option" tf:"option"`
}

type LinuxVirtualMachineSpecOsDisk struct {
	Caching string `json:"caching" tf:"caching"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DiffDiskSettings []LinuxVirtualMachineSpecOsDiskDiffDiskSettings `json:"diffDiskSettings,omitempty" tf:"diff_disk_settings,omitempty"`
	// +optional
	DiskEncryptionSetID string `json:"diskEncryptionSetID,omitempty" tf:"disk_encryption_set_id,omitempty"`
	// +optional
	DiskSizeGb int64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`
	// +optional
	Name               string `json:"name,omitempty" tf:"name,omitempty"`
	StorageAccountType string `json:"storageAccountType" tf:"storage_account_type"`
	// +optional
	WriteAcceleratorEnabled bool `json:"writeAcceleratorEnabled,omitempty" tf:"write_accelerator_enabled,omitempty"`
}

type LinuxVirtualMachineSpecPlan struct {
	Name      string `json:"name" tf:"name"`
	Product   string `json:"product" tf:"product"`
	Publisher string `json:"publisher" tf:"publisher"`
}

type LinuxVirtualMachineSpecSecretCertificate struct {
	Url string `json:"url" tf:"url"`
}

type LinuxVirtualMachineSpecSecret struct {
	// +kubebuilder:validation:MinItems=1
	Certificate []LinuxVirtualMachineSpecSecretCertificate `json:"certificate" tf:"certificate"`
	KeyVaultID  string                                     `json:"keyVaultID" tf:"key_vault_id"`
}

type LinuxVirtualMachineSpecSourceImageReference struct {
	Offer     string `json:"offer" tf:"offer"`
	Publisher string `json:"publisher" tf:"publisher"`
	Sku       string `json:"sku" tf:"sku"`
	Version   string `json:"version" tf:"version"`
}

type LinuxVirtualMachineSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	AdditionalCapabilities []LinuxVirtualMachineSpecAdditionalCapabilities `json:"additionalCapabilities,omitempty" tf:"additional_capabilities,omitempty"`
	// +optional
	AdminPassword string `json:"-" sensitive:"true" tf:"admin_password,omitempty"`
	// +optional
	AdminSSHKey   []LinuxVirtualMachineSpecAdminSSHKey `json:"adminSSHKey,omitempty" tf:"admin_ssh_key,omitempty"`
	AdminUsername string                               `json:"adminUsername" tf:"admin_username"`
	// +optional
	AllowExtensionOperations bool `json:"allowExtensionOperations,omitempty" tf:"allow_extension_operations,omitempty"`
	// +optional
	AvailabilitySetID string `json:"availabilitySetID,omitempty" tf:"availability_set_id,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	BootDiagnostics []LinuxVirtualMachineSpecBootDiagnostics `json:"bootDiagnostics,omitempty" tf:"boot_diagnostics,omitempty"`
	// +optional
	ComputerName string `json:"computerName,omitempty" tf:"computer_name,omitempty"`
	// +optional
	CustomData string `json:"-" sensitive:"true" tf:"custom_data,omitempty"`
	// +optional
	DedicatedHostID string `json:"dedicatedHostID,omitempty" tf:"dedicated_host_id,omitempty"`
	// +optional
	DisablePasswordAuthentication bool `json:"disablePasswordAuthentication,omitempty" tf:"disable_password_authentication,omitempty"`
	// +optional
	EvictionPolicy string `json:"evictionPolicy,omitempty" tf:"eviction_policy,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Identity []LinuxVirtualMachineSpecIdentity `json:"identity,omitempty" tf:"identity,omitempty"`
	Location string                            `json:"location" tf:"location"`
	// +optional
	MaxBidPrice float64 `json:"maxBidPrice,omitempty" tf:"max_bid_price,omitempty"`
	Name        string  `json:"name" tf:"name"`
	// +kubebuilder:validation:MinItems=1
	NetworkInterfaceIDS []string `json:"networkInterfaceIDS" tf:"network_interface_ids"`
	// +kubebuilder:validation:MaxItems=1
	OsDisk []LinuxVirtualMachineSpecOsDisk `json:"osDisk" tf:"os_disk"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Plan []LinuxVirtualMachineSpecPlan `json:"plan,omitempty" tf:"plan,omitempty"`
	// +optional
	Priority string `json:"priority,omitempty" tf:"priority,omitempty"`
	// +optional
	PrivateIPAddress string `json:"privateIPAddress,omitempty" tf:"private_ip_address,omitempty"`
	// +optional
	PrivateIPAddresses []string `json:"privateIPAddresses,omitempty" tf:"private_ip_addresses,omitempty"`
	// +optional
	ProvisionVmAgent bool `json:"provisionVmAgent,omitempty" tf:"provision_vm_agent,omitempty"`
	// +optional
	ProximityPlacementGroupID string `json:"proximityPlacementGroupID,omitempty" tf:"proximity_placement_group_id,omitempty"`
	// +optional
	PublicIPAddress string `json:"publicIPAddress,omitempty" tf:"public_ip_address,omitempty"`
	// +optional
	PublicIPAddresses []string `json:"publicIPAddresses,omitempty" tf:"public_ip_addresses,omitempty"`
	ResourceGroupName string   `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Secret []LinuxVirtualMachineSpecSecret `json:"secret,omitempty" tf:"secret,omitempty"`
	Size   string                          `json:"size" tf:"size"`
	// +optional
	SourceImageID string `json:"sourceImageID,omitempty" tf:"source_image_id,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SourceImageReference []LinuxVirtualMachineSpecSourceImageReference `json:"sourceImageReference,omitempty" tf:"source_image_reference,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	VirtualMachineID string `json:"virtualMachineID,omitempty" tf:"virtual_machine_id,omitempty"`
	// +optional
	VirtualMachineScaleSetID string `json:"virtualMachineScaleSetID,omitempty" tf:"virtual_machine_scale_set_id,omitempty"`
	// +optional
	Zone string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type LinuxVirtualMachineStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *LinuxVirtualMachineSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LinuxVirtualMachineList is a list of LinuxVirtualMachines
type LinuxVirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LinuxVirtualMachine CRD objects
	Items []LinuxVirtualMachine `json:"items,omitempty"`
}
