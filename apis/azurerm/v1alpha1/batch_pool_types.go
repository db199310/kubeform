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

type BatchPool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BatchPoolSpec   `json:"spec,omitempty"`
	Status            BatchPoolStatus `json:"status,omitempty"`
}

type BatchPoolSpecAutoScale struct {
	// +optional
	EvaluationInterval string `json:"evaluationInterval,omitempty" tf:"evaluation_interval,omitempty"`
	Formula            string `json:"formula" tf:"formula"`
}

type BatchPoolSpecCertificate struct {
	ID            string `json:"ID" tf:"id"`
	StoreLocation string `json:"storeLocation" tf:"store_location"`
	// +optional
	StoreName string `json:"storeName,omitempty" tf:"store_name,omitempty"`
	// +optional
	Visibility []string `json:"visibility,omitempty" tf:"visibility,omitempty"`
}

type BatchPoolSpecContainerConfigurationContainerRegistries struct {
	Password       string `json:"-" sensitive:"true" tf:"password"`
	RegistryServer string `json:"registryServer" tf:"registry_server"`
	UserName       string `json:"userName" tf:"user_name"`
}

type BatchPoolSpecContainerConfiguration struct {
	// +optional
	ContainerRegistries []BatchPoolSpecContainerConfigurationContainerRegistries `json:"containerRegistries,omitempty" tf:"container_registries,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
}

type BatchPoolSpecFixedScale struct {
	// +optional
	ResizeTimeout string `json:"resizeTimeout,omitempty" tf:"resize_timeout,omitempty"`
	// +optional
	TargetDedicatedNodes int64 `json:"targetDedicatedNodes,omitempty" tf:"target_dedicated_nodes,omitempty"`
	// +optional
	TargetLowPriorityNodes int64 `json:"targetLowPriorityNodes,omitempty" tf:"target_low_priority_nodes,omitempty"`
}

type BatchPoolSpecNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRules struct {
	Access              string `json:"access" tf:"access"`
	Priority            int64  `json:"priority" tf:"priority"`
	SourceAddressPrefix string `json:"sourceAddressPrefix" tf:"source_address_prefix"`
}

type BatchPoolSpecNetworkConfigurationEndpointConfiguration struct {
	BackendPort       int64  `json:"backendPort" tf:"backend_port"`
	FrontendPortRange string `json:"frontendPortRange" tf:"frontend_port_range"`
	Name              string `json:"name" tf:"name"`
	// +optional
	NetworkSecurityGroupRules []BatchPoolSpecNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRules `json:"networkSecurityGroupRules,omitempty" tf:"network_security_group_rules,omitempty"`
	Protocol                  string                                                                            `json:"protocol" tf:"protocol"`
}

type BatchPoolSpecNetworkConfiguration struct {
	// +optional
	EndpointConfiguration []BatchPoolSpecNetworkConfigurationEndpointConfiguration `json:"endpointConfiguration,omitempty" tf:"endpoint_configuration,omitempty"`
	SubnetID              string                                                   `json:"subnetID" tf:"subnet_id"`
}

type BatchPoolSpecStartTaskResourceFile struct {
	// +optional
	AutoStorageContainerName string `json:"autoStorageContainerName,omitempty" tf:"auto_storage_container_name,omitempty"`
	// +optional
	BlobPrefix string `json:"blobPrefix,omitempty" tf:"blob_prefix,omitempty"`
	// +optional
	FileMode string `json:"fileMode,omitempty" tf:"file_mode,omitempty"`
	// +optional
	FilePath string `json:"filePath,omitempty" tf:"file_path,omitempty"`
	// +optional
	HttpURL string `json:"httpURL,omitempty" tf:"http_url,omitempty"`
	// +optional
	StorageContainerURL string `json:"storageContainerURL,omitempty" tf:"storage_container_url,omitempty"`
}

type BatchPoolSpecStartTaskUserIdentityAutoUser struct {
	// +optional
	ElevationLevel string `json:"elevationLevel,omitempty" tf:"elevation_level,omitempty"`
	// +optional
	Scope string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type BatchPoolSpecStartTaskUserIdentity struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AutoUser []BatchPoolSpecStartTaskUserIdentityAutoUser `json:"autoUser,omitempty" tf:"auto_user,omitempty"`
	// +optional
	UserName string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

type BatchPoolSpecStartTask struct {
	CommandLine string `json:"commandLine" tf:"command_line"`
	// +optional
	Environment map[string]string `json:"environment,omitempty" tf:"environment,omitempty"`
	// +optional
	MaxTaskRetryCount int64 `json:"maxTaskRetryCount,omitempty" tf:"max_task_retry_count,omitempty"`
	// +optional
	ResourceFile []BatchPoolSpecStartTaskResourceFile `json:"resourceFile,omitempty" tf:"resource_file,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	UserIdentity []BatchPoolSpecStartTaskUserIdentity `json:"userIdentity" tf:"user_identity"`
	// +optional
	WaitForSuccess bool `json:"waitForSuccess,omitempty" tf:"wait_for_success,omitempty"`
}

type BatchPoolSpecStorageImageReference struct {
	// +optional
	ID string `json:"ID,omitempty" tf:"id,omitempty"`
	// +optional
	Offer string `json:"offer,omitempty" tf:"offer,omitempty"`
	// +optional
	Publisher string `json:"publisher,omitempty" tf:"publisher,omitempty"`
	// +optional
	Sku string `json:"sku,omitempty" tf:"sku,omitempty"`
	// +optional
	Version string `json:"version,omitempty" tf:"version,omitempty"`
}

type BatchPoolSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	AccountName string `json:"accountName" tf:"account_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AutoScale []BatchPoolSpecAutoScale `json:"autoScale,omitempty" tf:"auto_scale,omitempty"`
	// +optional
	Certificate []BatchPoolSpecCertificate `json:"certificate,omitempty" tf:"certificate,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	ContainerConfiguration []BatchPoolSpecContainerConfiguration `json:"containerConfiguration,omitempty" tf:"container_configuration,omitempty"`
	// +optional
	DisplayName string `json:"displayName,omitempty" tf:"display_name,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	FixedScale []BatchPoolSpecFixedScale `json:"fixedScale,omitempty" tf:"fixed_scale,omitempty"`
	// +optional
	MaxTasksPerNode int64 `json:"maxTasksPerNode,omitempty" tf:"max_tasks_per_node,omitempty"`
	// +optional
	Metadata map[string]string `json:"metadata,omitempty" tf:"metadata,omitempty"`
	Name     string            `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	NetworkConfiguration []BatchPoolSpecNetworkConfiguration `json:"networkConfiguration,omitempty" tf:"network_configuration,omitempty"`
	NodeAgentSkuID       string                              `json:"nodeAgentSkuID" tf:"node_agent_sku_id"`
	ResourceGroupName    string                              `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	StartTask []BatchPoolSpecStartTask `json:"startTask,omitempty" tf:"start_task,omitempty"`
	// +optional
	StopPendingResizeOperation bool `json:"stopPendingResizeOperation,omitempty" tf:"stop_pending_resize_operation,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	StorageImageReference []BatchPoolSpecStorageImageReference `json:"storageImageReference" tf:"storage_image_reference"`
	VmSize                string                               `json:"vmSize" tf:"vm_size"`
}

type BatchPoolStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *BatchPoolSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
	// +optional
	TerraformErrors []string `json:"terraformErrors,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BatchPoolList is a list of BatchPools
type BatchPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BatchPool CRD objects
	Items []BatchPool `json:"items,omitempty"`
}
