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

type KubernetesCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubernetesClusterSpec   `json:"spec,omitempty"`
	Status            KubernetesClusterStatus `json:"status,omitempty"`
}

type KubernetesClusterSpecAddonProfileAciConnectorLinux struct {
	Enabled bool `json:"enabled" tf:"enabled"`
	// +optional
	SubnetName string `json:"subnetName,omitempty" tf:"subnet_name,omitempty"`
}

type KubernetesClusterSpecAddonProfileAzurePolicy struct {
	Enabled bool `json:"enabled" tf:"enabled"`
}

type KubernetesClusterSpecAddonProfileHttpApplicationRouting struct {
	Enabled bool `json:"enabled" tf:"enabled"`
	// +optional
	HttpApplicationRoutingZoneName string `json:"httpApplicationRoutingZoneName,omitempty" tf:"http_application_routing_zone_name,omitempty"`
}

type KubernetesClusterSpecAddonProfileKubeDashboard struct {
	Enabled bool `json:"enabled" tf:"enabled"`
}

type KubernetesClusterSpecAddonProfileOmsAgent struct {
	Enabled bool `json:"enabled" tf:"enabled"`
	// +optional
	LogAnalyticsWorkspaceID string `json:"logAnalyticsWorkspaceID,omitempty" tf:"log_analytics_workspace_id,omitempty"`
}

type KubernetesClusterSpecAddonProfile struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AciConnectorLinux []KubernetesClusterSpecAddonProfileAciConnectorLinux `json:"aciConnectorLinux,omitempty" tf:"aci_connector_linux,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AzurePolicy []KubernetesClusterSpecAddonProfileAzurePolicy `json:"azurePolicy,omitempty" tf:"azure_policy,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HttpApplicationRouting []KubernetesClusterSpecAddonProfileHttpApplicationRouting `json:"httpApplicationRouting,omitempty" tf:"http_application_routing,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	KubeDashboard []KubernetesClusterSpecAddonProfileKubeDashboard `json:"kubeDashboard,omitempty" tf:"kube_dashboard,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	OmsAgent []KubernetesClusterSpecAddonProfileOmsAgent `json:"omsAgent,omitempty" tf:"oms_agent,omitempty"`
}

type KubernetesClusterSpecAgentPoolProfile struct {
	// +optional
	AvailabilityZones []string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`
	// +optional
	Count int64 `json:"count,omitempty" tf:"count,omitempty"`
	// +optional
	// Deprecated
	DnsPrefix string `json:"dnsPrefix,omitempty" tf:"dns_prefix,omitempty"`
	// +optional
	EnableAutoScaling bool `json:"enableAutoScaling,omitempty" tf:"enable_auto_scaling,omitempty"`
	// +optional
	EnableNodePublicIP bool `json:"enableNodePublicIP,omitempty" tf:"enable_node_public_ip,omitempty"`
	// +optional
	// Deprecated
	Fqdn string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`
	// +optional
	MaxCount int64 `json:"maxCount,omitempty" tf:"max_count,omitempty"`
	// +optional
	MaxPods int64 `json:"maxPods,omitempty" tf:"max_pods,omitempty"`
	// +optional
	MinCount int64  `json:"minCount,omitempty" tf:"min_count,omitempty"`
	Name     string `json:"name" tf:"name"`
	// +optional
	NodeTaints []string `json:"nodeTaints,omitempty" tf:"node_taints,omitempty"`
	// +optional
	OsDiskSizeGb int64 `json:"osDiskSizeGb,omitempty" tf:"os_disk_size_gb,omitempty"`
	// +optional
	OsType string `json:"osType,omitempty" tf:"os_type,omitempty"`
	// +optional
	Type   string `json:"type,omitempty" tf:"type,omitempty"`
	VmSize string `json:"vmSize" tf:"vm_size"`
	// +optional
	VnetSubnetID string `json:"vnetSubnetID,omitempty" tf:"vnet_subnet_id,omitempty"`
}

type KubernetesClusterSpecDefaultNodePool struct {
	// +optional
	AvailabilityZones []string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`
	// +optional
	EnableAutoScaling bool `json:"enableAutoScaling,omitempty" tf:"enable_auto_scaling,omitempty"`
	// +optional
	EnableNodePublicIP bool `json:"enableNodePublicIP,omitempty" tf:"enable_node_public_ip,omitempty"`
	// +optional
	MaxCount int64 `json:"maxCount,omitempty" tf:"max_count,omitempty"`
	// +optional
	MaxPods int64 `json:"maxPods,omitempty" tf:"max_pods,omitempty"`
	// +optional
	MinCount int64  `json:"minCount,omitempty" tf:"min_count,omitempty"`
	Name     string `json:"name" tf:"name"`
	// +optional
	NodeCount int64 `json:"nodeCount,omitempty" tf:"node_count,omitempty"`
	// +optional
	NodeTaints []string `json:"nodeTaints,omitempty" tf:"node_taints,omitempty"`
	// +optional
	OsDiskSizeGb int64 `json:"osDiskSizeGb,omitempty" tf:"os_disk_size_gb,omitempty"`
	// +optional
	Type   string `json:"type,omitempty" tf:"type,omitempty"`
	VmSize string `json:"vmSize" tf:"vm_size"`
	// +optional
	VnetSubnetID string `json:"vnetSubnetID,omitempty" tf:"vnet_subnet_id,omitempty"`
}

type KubernetesClusterSpecIdentity struct {
	// +optional
	PrincipalID string `json:"principalID,omitempty" tf:"principal_id,omitempty"`
	// +optional
	TenantID string `json:"tenantID,omitempty" tf:"tenant_id,omitempty"`
	Type     string `json:"type" tf:"type"`
}

type KubernetesClusterSpecKubeAdminConfig struct {
	// +optional
	ClientCertificate string `json:"clientCertificate,omitempty" tf:"client_certificate,omitempty"`
	// +optional
	ClientKey string `json:"-" sensitive:"true" tf:"client_key,omitempty"`
	// +optional
	ClusterCaCertificate string `json:"clusterCaCertificate,omitempty" tf:"cluster_ca_certificate,omitempty"`
	// +optional
	Host string `json:"host,omitempty" tf:"host,omitempty"`
	// +optional
	Password string `json:"-" sensitive:"true" tf:"password,omitempty"`
	// +optional
	Username string `json:"username,omitempty" tf:"username,omitempty"`
}

type KubernetesClusterSpecKubeConfig struct {
	// +optional
	ClientCertificate string `json:"clientCertificate,omitempty" tf:"client_certificate,omitempty"`
	// +optional
	ClientKey string `json:"-" sensitive:"true" tf:"client_key,omitempty"`
	// +optional
	ClusterCaCertificate string `json:"clusterCaCertificate,omitempty" tf:"cluster_ca_certificate,omitempty"`
	// +optional
	Host string `json:"host,omitempty" tf:"host,omitempty"`
	// +optional
	Password string `json:"-" sensitive:"true" tf:"password,omitempty"`
	// +optional
	Username string `json:"username,omitempty" tf:"username,omitempty"`
}

type KubernetesClusterSpecLinuxProfileSshKey struct {
	KeyData string `json:"keyData" tf:"key_data"`
}

type KubernetesClusterSpecLinuxProfile struct {
	AdminUsername string `json:"adminUsername" tf:"admin_username"`
	// +kubebuilder:validation:MaxItems=1
	SshKey []KubernetesClusterSpecLinuxProfileSshKey `json:"sshKey" tf:"ssh_key"`
}

type KubernetesClusterSpecNetworkProfileLoadBalancerProfile struct {
	// +optional
	EffectiveOutboundIPS []string `json:"effectiveOutboundIPS,omitempty" tf:"effective_outbound_ips,omitempty"`
	// +optional
	ManagedOutboundIPCount int64 `json:"managedOutboundIPCount,omitempty" tf:"managed_outbound_ip_count,omitempty"`
	// +optional
	OutboundIPAddressIDS []string `json:"outboundIPAddressIDS,omitempty" tf:"outbound_ip_address_ids,omitempty"`
	// +optional
	OutboundIPPrefixIDS []string `json:"outboundIPPrefixIDS,omitempty" tf:"outbound_ip_prefix_ids,omitempty"`
}

type KubernetesClusterSpecNetworkProfile struct {
	// +optional
	DnsServiceIP string `json:"dnsServiceIP,omitempty" tf:"dns_service_ip,omitempty"`
	// +optional
	DockerBridgeCIDR string `json:"dockerBridgeCIDR,omitempty" tf:"docker_bridge_cidr,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LoadBalancerProfile []KubernetesClusterSpecNetworkProfileLoadBalancerProfile `json:"loadBalancerProfile,omitempty" tf:"load_balancer_profile,omitempty"`
	// +optional
	LoadBalancerSku string `json:"loadBalancerSku,omitempty" tf:"load_balancer_sku,omitempty"`
	NetworkPlugin   string `json:"networkPlugin" tf:"network_plugin"`
	// +optional
	NetworkPolicy string `json:"networkPolicy,omitempty" tf:"network_policy,omitempty"`
	// +optional
	PodCIDR string `json:"podCIDR,omitempty" tf:"pod_cidr,omitempty"`
	// +optional
	ServiceCIDR string `json:"serviceCIDR,omitempty" tf:"service_cidr,omitempty"`
}

type KubernetesClusterSpecRoleBasedAccessControlAzureActiveDirectory struct {
	ClientAppID     string `json:"clientAppID" tf:"client_app_id"`
	ServerAppID     string `json:"serverAppID" tf:"server_app_id"`
	ServerAppSecret string `json:"-" sensitive:"true" tf:"server_app_secret"`
	// +optional
	TenantID string `json:"tenantID,omitempty" tf:"tenant_id,omitempty"`
}

type KubernetesClusterSpecRoleBasedAccessControl struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AzureActiveDirectory []KubernetesClusterSpecRoleBasedAccessControlAzureActiveDirectory `json:"azureActiveDirectory,omitempty" tf:"azure_active_directory,omitempty"`
	Enabled              bool                                                              `json:"enabled" tf:"enabled"`
}

type KubernetesClusterSpecServicePrincipal struct {
	ClientID     string `json:"clientID" tf:"client_id"`
	ClientSecret string `json:"-" sensitive:"true" tf:"client_secret"`
}

type KubernetesClusterSpecWindowsProfile struct {
	// +optional
	AdminPassword string `json:"-" sensitive:"true" tf:"admin_password,omitempty"`
	AdminUsername string `json:"adminUsername" tf:"admin_username"`
}

type KubernetesClusterSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	AddonProfile []KubernetesClusterSpecAddonProfile `json:"addonProfile,omitempty" tf:"addon_profile,omitempty"`
	// +optional
	// Deprecated
	AgentPoolProfile []KubernetesClusterSpecAgentPoolProfile `json:"agentPoolProfile,omitempty" tf:"agent_pool_profile,omitempty"`
	// +optional
	ApiServerAuthorizedIPRanges []string `json:"apiServerAuthorizedIPRanges,omitempty" tf:"api_server_authorized_ip_ranges,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DefaultNodePool []KubernetesClusterSpecDefaultNodePool `json:"defaultNodePool,omitempty" tf:"default_node_pool,omitempty"`
	DnsPrefix       string                                 `json:"dnsPrefix" tf:"dns_prefix"`
	// +optional
	EnablePodSecurityPolicy bool `json:"enablePodSecurityPolicy,omitempty" tf:"enable_pod_security_policy,omitempty"`
	// +optional
	Fqdn string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Identity []KubernetesClusterSpecIdentity `json:"identity,omitempty" tf:"identity,omitempty"`
	// +optional
	KubeAdminConfig []KubernetesClusterSpecKubeAdminConfig `json:"kubeAdminConfig,omitempty" tf:"kube_admin_config,omitempty"`
	// +optional
	KubeAdminConfigRaw string `json:"-" sensitive:"true" tf:"kube_admin_config_raw,omitempty"`
	// +optional
	KubeConfig []KubernetesClusterSpecKubeConfig `json:"kubeConfig,omitempty" tf:"kube_config,omitempty"`
	// +optional
	KubeConfigRaw string `json:"-" sensitive:"true" tf:"kube_config_raw,omitempty"`
	// +optional
	KubernetesVersion string `json:"kubernetesVersion,omitempty" tf:"kubernetes_version,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LinuxProfile []KubernetesClusterSpecLinuxProfile `json:"linuxProfile,omitempty" tf:"linux_profile,omitempty"`
	Location     string                              `json:"location" tf:"location"`
	Name         string                              `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	NetworkProfile []KubernetesClusterSpecNetworkProfile `json:"networkProfile,omitempty" tf:"network_profile,omitempty"`
	// +optional
	NodeResourceGroup string `json:"nodeResourceGroup,omitempty" tf:"node_resource_group,omitempty"`
	// +optional
	PrivateFqdn string `json:"privateFqdn,omitempty" tf:"private_fqdn,omitempty"`
	// +optional
	PrivateLinkEnabled bool   `json:"privateLinkEnabled,omitempty" tf:"private_link_enabled,omitempty"`
	ResourceGroupName  string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RoleBasedAccessControl []KubernetesClusterSpecRoleBasedAccessControl `json:"roleBasedAccessControl,omitempty" tf:"role_based_access_control,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	ServicePrincipal []KubernetesClusterSpecServicePrincipal `json:"servicePrincipal" tf:"service_principal"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	WindowsProfile []KubernetesClusterSpecWindowsProfile `json:"windowsProfile,omitempty" tf:"windows_profile,omitempty"`
}

type KubernetesClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *KubernetesClusterSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase           base.Phase `json:"phase,omitempty"`
	TerraformErrors []string   `json:"terraformErrors,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// KubernetesClusterList is a list of KubernetesClusters
type KubernetesClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KubernetesCluster CRD objects
	Items []KubernetesCluster `json:"items,omitempty"`
}
