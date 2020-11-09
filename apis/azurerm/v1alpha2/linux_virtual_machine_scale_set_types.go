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

type LinuxVirtualMachineScaleSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LinuxVirtualMachineScaleSetSpec   `json:"spec,omitempty"`
	Status            LinuxVirtualMachineScaleSetStatus `json:"status,omitempty"`
}

type LinuxVirtualMachineScaleSetSpecAdditionalCapabilities struct {
	// +optional
	UltraSsdEnabled bool `json:"ultraSsdEnabled,omitempty" tf:"ultra_ssd_enabled,omitempty"`
}

type LinuxVirtualMachineScaleSetSpecAdminSSHKey struct {
	PublicKey string `json:"publicKey" tf:"public_key"`
	Username  string `json:"username" tf:"username"`
}

type LinuxVirtualMachineScaleSetSpecAutomaticInstanceRepair struct {
	Enabled bool `json:"enabled" tf:"enabled"`
	// +optional
	GracePeriod string `json:"gracePeriod,omitempty" tf:"grace_period,omitempty"`
}

type LinuxVirtualMachineScaleSetSpecAutomaticOsUpgradePolicy struct {
	DisableAutomaticRollback bool `json:"disableAutomaticRollback" tf:"disable_automatic_rollback"`
	EnableAutomaticOsUpgrade bool `json:"enableAutomaticOsUpgrade" tf:"enable_automatic_os_upgrade"`
}

type LinuxVirtualMachineScaleSetSpecBootDiagnostics struct {
	StorageAccountURI string `json:"storageAccountURI" tf:"storage_account_uri"`
}

type LinuxVirtualMachineScaleSetSpecDataDisk struct {
	Caching string `json:"caching" tf:"caching"`
	// +optional
	CreateOption string `json:"createOption,omitempty" tf:"create_option,omitempty"`
	// +optional
	DiskEncryptionSetID string `json:"diskEncryptionSetID,omitempty" tf:"disk_encryption_set_id,omitempty"`
	DiskSizeGb          int64  `json:"diskSizeGb" tf:"disk_size_gb"`
	Lun                 int64  `json:"lun" tf:"lun"`
	StorageAccountType  string `json:"storageAccountType" tf:"storage_account_type"`
	// +optional
	WriteAcceleratorEnabled bool `json:"writeAcceleratorEnabled,omitempty" tf:"write_accelerator_enabled,omitempty"`
}

type LinuxVirtualMachineScaleSetSpecIdentity struct {
	// +optional
	IdentityIDS []string `json:"identityIDS,omitempty" tf:"identity_ids,omitempty"`
	// +optional
	PrincipalID string `json:"principalID,omitempty" tf:"principal_id,omitempty"`
	Type        string `json:"type" tf:"type"`
}

type LinuxVirtualMachineScaleSetSpecNetworkInterfaceIpConfigurationPublicIPAddressIpTag struct {
	Tag  string `json:"tag" tf:"tag"`
	Type string `json:"type" tf:"type"`
}

type LinuxVirtualMachineScaleSetSpecNetworkInterfaceIpConfigurationPublicIPAddress struct {
	// +optional
	DomainNameLabel string `json:"domainNameLabel,omitempty" tf:"domain_name_label,omitempty"`
	// +optional
	IdleTimeoutInMinutes int64 `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes,omitempty"`
	// +optional
	IpTag []LinuxVirtualMachineScaleSetSpecNetworkInterfaceIpConfigurationPublicIPAddressIpTag `json:"ipTag,omitempty" tf:"ip_tag,omitempty"`
	Name  string                                                                               `json:"name" tf:"name"`
	// +optional
	PublicIPPrefixID string `json:"publicIPPrefixID,omitempty" tf:"public_ip_prefix_id,omitempty"`
}

type LinuxVirtualMachineScaleSetSpecNetworkInterfaceIpConfiguration struct {
	// +optional
	ApplicationGatewayBackendAddressPoolIDS []string `json:"applicationGatewayBackendAddressPoolIDS,omitempty" tf:"application_gateway_backend_address_pool_ids,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ApplicationSecurityGroupIDS []string `json:"applicationSecurityGroupIDS,omitempty" tf:"application_security_group_ids,omitempty"`
	// +optional
	LoadBalancerBackendAddressPoolIDS []string `json:"loadBalancerBackendAddressPoolIDS,omitempty" tf:"load_balancer_backend_address_pool_ids,omitempty"`
	// +optional
	LoadBalancerInboundNATRulesIDS []string `json:"loadBalancerInboundNATRulesIDS,omitempty" tf:"load_balancer_inbound_nat_rules_ids,omitempty"`
	Name                           string   `json:"name" tf:"name"`
	// +optional
	Primary bool `json:"primary,omitempty" tf:"primary,omitempty"`
	// +optional
	PublicIPAddress []LinuxVirtualMachineScaleSetSpecNetworkInterfaceIpConfigurationPublicIPAddress `json:"publicIPAddress,omitempty" tf:"public_ip_address,omitempty"`
	// +optional
	SubnetID string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
	// +optional
	Version string `json:"version,omitempty" tf:"version,omitempty"`
}

type LinuxVirtualMachineScaleSetSpecNetworkInterface struct {
	// +optional
	DnsServers []string `json:"dnsServers,omitempty" tf:"dns_servers,omitempty"`
	// +optional
	EnableAcceleratedNetworking bool `json:"enableAcceleratedNetworking,omitempty" tf:"enable_accelerated_networking,omitempty"`
	// +optional
	EnableIPForwarding bool                                                             `json:"enableIPForwarding,omitempty" tf:"enable_ip_forwarding,omitempty"`
	IpConfiguration    []LinuxVirtualMachineScaleSetSpecNetworkInterfaceIpConfiguration `json:"ipConfiguration" tf:"ip_configuration"`
	Name               string                                                           `json:"name" tf:"name"`
	// +optional
	NetworkSecurityGroupID string `json:"networkSecurityGroupID,omitempty" tf:"network_security_group_id,omitempty"`
	// +optional
	Primary bool `json:"primary,omitempty" tf:"primary,omitempty"`
}

type LinuxVirtualMachineScaleSetSpecOsDiskDiffDiskSettings struct {
	Option string `json:"option" tf:"option"`
}

type LinuxVirtualMachineScaleSetSpecOsDisk struct {
	Caching string `json:"caching" tf:"caching"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DiffDiskSettings []LinuxVirtualMachineScaleSetSpecOsDiskDiffDiskSettings `json:"diffDiskSettings,omitempty" tf:"diff_disk_settings,omitempty"`
	// +optional
	DiskEncryptionSetID string `json:"diskEncryptionSetID,omitempty" tf:"disk_encryption_set_id,omitempty"`
	// +optional
	DiskSizeGb         int64  `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`
	StorageAccountType string `json:"storageAccountType" tf:"storage_account_type"`
	// +optional
	WriteAcceleratorEnabled bool `json:"writeAcceleratorEnabled,omitempty" tf:"write_accelerator_enabled,omitempty"`
}

type LinuxVirtualMachineScaleSetSpecPlan struct {
	Name      string `json:"name" tf:"name"`
	Product   string `json:"product" tf:"product"`
	Publisher string `json:"publisher" tf:"publisher"`
}

type LinuxVirtualMachineScaleSetSpecRollingUpgradePolicy struct {
	MaxBatchInstancePercent             int64  `json:"maxBatchInstancePercent" tf:"max_batch_instance_percent"`
	MaxUnhealthyInstancePercent         int64  `json:"maxUnhealthyInstancePercent" tf:"max_unhealthy_instance_percent"`
	MaxUnhealthyUpgradedInstancePercent int64  `json:"maxUnhealthyUpgradedInstancePercent" tf:"max_unhealthy_upgraded_instance_percent"`
	PauseTimeBetweenBatches             string `json:"pauseTimeBetweenBatches" tf:"pause_time_between_batches"`
}

type LinuxVirtualMachineScaleSetSpecSecretCertificate struct {
	Url string `json:"url" tf:"url"`
}

type LinuxVirtualMachineScaleSetSpecSecret struct {
	// +kubebuilder:validation:MinItems=1
	Certificate []LinuxVirtualMachineScaleSetSpecSecretCertificate `json:"certificate" tf:"certificate"`
	KeyVaultID  string                                             `json:"keyVaultID" tf:"key_vault_id"`
}

type LinuxVirtualMachineScaleSetSpecSourceImageReference struct {
	Offer     string `json:"offer" tf:"offer"`
	Publisher string `json:"publisher" tf:"publisher"`
	Sku       string `json:"sku" tf:"sku"`
	Version   string `json:"version" tf:"version"`
}

type LinuxVirtualMachineScaleSetSpecTerminateNotification struct {
	Enabled bool `json:"enabled" tf:"enabled"`
	// +optional
	Timeout string `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

type LinuxVirtualMachineScaleSetSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	AdditionalCapabilities []LinuxVirtualMachineScaleSetSpecAdditionalCapabilities `json:"additionalCapabilities,omitempty" tf:"additional_capabilities,omitempty"`
	// +optional
	AdminPassword string `json:"-" sensitive:"true" tf:"admin_password,omitempty"`
	// +optional
	AdminSSHKey   []LinuxVirtualMachineScaleSetSpecAdminSSHKey `json:"adminSSHKey,omitempty" tf:"admin_ssh_key,omitempty"`
	AdminUsername string                                       `json:"adminUsername" tf:"admin_username"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AutomaticInstanceRepair []LinuxVirtualMachineScaleSetSpecAutomaticInstanceRepair `json:"automaticInstanceRepair,omitempty" tf:"automatic_instance_repair,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AutomaticOsUpgradePolicy []LinuxVirtualMachineScaleSetSpecAutomaticOsUpgradePolicy `json:"automaticOsUpgradePolicy,omitempty" tf:"automatic_os_upgrade_policy,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	BootDiagnostics []LinuxVirtualMachineScaleSetSpecBootDiagnostics `json:"bootDiagnostics,omitempty" tf:"boot_diagnostics,omitempty"`
	// +optional
	ComputerNamePrefix string `json:"computerNamePrefix,omitempty" tf:"computer_name_prefix,omitempty"`
	// +optional
	CustomData string `json:"-" sensitive:"true" tf:"custom_data,omitempty"`
	// +optional
	DataDisk []LinuxVirtualMachineScaleSetSpecDataDisk `json:"dataDisk,omitempty" tf:"data_disk,omitempty"`
	// +optional
	DisablePasswordAuthentication bool `json:"disablePasswordAuthentication,omitempty" tf:"disable_password_authentication,omitempty"`
	// +optional
	DoNotRunExtensionsOnOverprovisionedMachines bool `json:"doNotRunExtensionsOnOverprovisionedMachines,omitempty" tf:"do_not_run_extensions_on_overprovisioned_machines,omitempty"`
	// +optional
	EvictionPolicy string `json:"evictionPolicy,omitempty" tf:"eviction_policy,omitempty"`
	// +optional
	HealthProbeID string `json:"healthProbeID,omitempty" tf:"health_probe_id,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Identity  []LinuxVirtualMachineScaleSetSpecIdentity `json:"identity,omitempty" tf:"identity,omitempty"`
	Instances int64                                     `json:"instances" tf:"instances"`
	Location  string                                    `json:"location" tf:"location"`
	// +optional
	MaxBidPrice      float64                                           `json:"maxBidPrice,omitempty" tf:"max_bid_price,omitempty"`
	Name             string                                            `json:"name" tf:"name"`
	NetworkInterface []LinuxVirtualMachineScaleSetSpecNetworkInterface `json:"networkInterface" tf:"network_interface"`
	// +kubebuilder:validation:MaxItems=1
	OsDisk []LinuxVirtualMachineScaleSetSpecOsDisk `json:"osDisk" tf:"os_disk"`
	// +optional
	Overprovision bool `json:"overprovision,omitempty" tf:"overprovision,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Plan []LinuxVirtualMachineScaleSetSpecPlan `json:"plan,omitempty" tf:"plan,omitempty"`
	// +optional
	Priority string `json:"priority,omitempty" tf:"priority,omitempty"`
	// +optional
	ProvisionVmAgent bool `json:"provisionVmAgent,omitempty" tf:"provision_vm_agent,omitempty"`
	// +optional
	ProximityPlacementGroupID string `json:"proximityPlacementGroupID,omitempty" tf:"proximity_placement_group_id,omitempty"`
	ResourceGroupName         string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RollingUpgradePolicy []LinuxVirtualMachineScaleSetSpecRollingUpgradePolicy `json:"rollingUpgradePolicy,omitempty" tf:"rolling_upgrade_policy,omitempty"`
	// +optional
	ScaleInPolicy string `json:"scaleInPolicy,omitempty" tf:"scale_in_policy,omitempty"`
	// +optional
	Secret []LinuxVirtualMachineScaleSetSpecSecret `json:"secret,omitempty" tf:"secret,omitempty"`
	// +optional
	SinglePlacementGroup bool   `json:"singlePlacementGroup,omitempty" tf:"single_placement_group,omitempty"`
	Sku                  string `json:"sku" tf:"sku"`
	// +optional
	SourceImageID string `json:"sourceImageID,omitempty" tf:"source_image_id,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SourceImageReference []LinuxVirtualMachineScaleSetSpecSourceImageReference `json:"sourceImageReference,omitempty" tf:"source_image_reference,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	TerminateNotification []LinuxVirtualMachineScaleSetSpecTerminateNotification `json:"terminateNotification,omitempty" tf:"terminate_notification,omitempty"`
	// +optional
	UniqueID string `json:"uniqueID,omitempty" tf:"unique_id,omitempty"`
	// +optional
	UpgradeMode string `json:"upgradeMode,omitempty" tf:"upgrade_mode,omitempty"`
	// +optional
	ZoneBalance bool `json:"zoneBalance,omitempty" tf:"zone_balance,omitempty"`
	// +optional
	Zones []string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type LinuxVirtualMachineScaleSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *LinuxVirtualMachineScaleSetSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LinuxVirtualMachineScaleSetList is a list of LinuxVirtualMachineScaleSets
type LinuxVirtualMachineScaleSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LinuxVirtualMachineScaleSet CRD objects
	Items []LinuxVirtualMachineScaleSet `json:"items,omitempty"`
}
