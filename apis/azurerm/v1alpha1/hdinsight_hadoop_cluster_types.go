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

type HdinsightHadoopCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HdinsightHadoopClusterSpec   `json:"spec,omitempty"`
	Status            HdinsightHadoopClusterStatus `json:"status,omitempty"`
}

type HdinsightHadoopClusterSpecComponentVersion struct {
	Hadoop string `json:"hadoop" tf:"hadoop"`
}

type HdinsightHadoopClusterSpecGateway struct {
	Enabled  bool   `json:"enabled" tf:"enabled"`
	Password string `json:"-" sensitive:"true" tf:"password"`
	Username string `json:"username" tf:"username"`
}

type HdinsightHadoopClusterSpecMetastoresAmbari struct {
	DatabaseName string `json:"databaseName" tf:"database_name"`
	Password     string `json:"-" sensitive:"true" tf:"password"`
	Server       string `json:"server" tf:"server"`
	Username     string `json:"username" tf:"username"`
}

type HdinsightHadoopClusterSpecMetastoresHive struct {
	DatabaseName string `json:"databaseName" tf:"database_name"`
	Password     string `json:"-" sensitive:"true" tf:"password"`
	Server       string `json:"server" tf:"server"`
	Username     string `json:"username" tf:"username"`
}

type HdinsightHadoopClusterSpecMetastoresOozie struct {
	DatabaseName string `json:"databaseName" tf:"database_name"`
	Password     string `json:"-" sensitive:"true" tf:"password"`
	Server       string `json:"server" tf:"server"`
	Username     string `json:"username" tf:"username"`
}

type HdinsightHadoopClusterSpecMetastores struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Ambari []HdinsightHadoopClusterSpecMetastoresAmbari `json:"ambari,omitempty" tf:"ambari,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Hive []HdinsightHadoopClusterSpecMetastoresHive `json:"hive,omitempty" tf:"hive,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Oozie []HdinsightHadoopClusterSpecMetastoresOozie `json:"oozie,omitempty" tf:"oozie,omitempty"`
}

type HdinsightHadoopClusterSpecRolesEdgeNodeInstallScriptAction struct {
	Name string `json:"name" tf:"name"`
	Uri  string `json:"uri" tf:"uri"`
}

type HdinsightHadoopClusterSpecRolesEdgeNode struct {
	// +kubebuilder:validation:MinItems=1
	InstallScriptAction []HdinsightHadoopClusterSpecRolesEdgeNodeInstallScriptAction `json:"installScriptAction" tf:"install_script_action"`
	TargetInstanceCount int64                                                        `json:"targetInstanceCount" tf:"target_instance_count"`
	VmSize              string                                                       `json:"vmSize" tf:"vm_size"`
}

type HdinsightHadoopClusterSpecRolesHeadNode struct {
	// +optional
	Password string `json:"-" sensitive:"true" tf:"password,omitempty"`
	// +optional
	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`
	// +optional
	SubnetID string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
	Username string `json:"username" tf:"username"`
	// +optional
	VirtualNetworkID string `json:"virtualNetworkID,omitempty" tf:"virtual_network_id,omitempty"`
	VmSize           string `json:"vmSize" tf:"vm_size"`
}

type HdinsightHadoopClusterSpecRolesWorkerNode struct {
	// +optional
	MinInstanceCount int64 `json:"minInstanceCount,omitempty" tf:"min_instance_count,omitempty"`
	// +optional
	Password string `json:"-" sensitive:"true" tf:"password,omitempty"`
	// +optional
	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`
	// +optional
	SubnetID            string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
	TargetInstanceCount int64  `json:"targetInstanceCount" tf:"target_instance_count"`
	Username            string `json:"username" tf:"username"`
	// +optional
	VirtualNetworkID string `json:"virtualNetworkID,omitempty" tf:"virtual_network_id,omitempty"`
	VmSize           string `json:"vmSize" tf:"vm_size"`
}

type HdinsightHadoopClusterSpecRolesZookeeperNode struct {
	// +optional
	Password string `json:"-" sensitive:"true" tf:"password,omitempty"`
	// +optional
	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`
	// +optional
	SubnetID string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
	Username string `json:"username" tf:"username"`
	// +optional
	VirtualNetworkID string `json:"virtualNetworkID,omitempty" tf:"virtual_network_id,omitempty"`
	VmSize           string `json:"vmSize" tf:"vm_size"`
}

type HdinsightHadoopClusterSpecRoles struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	EdgeNode []HdinsightHadoopClusterSpecRolesEdgeNode `json:"edgeNode,omitempty" tf:"edge_node,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	HeadNode []HdinsightHadoopClusterSpecRolesHeadNode `json:"headNode" tf:"head_node"`
	// +kubebuilder:validation:MaxItems=1
	WorkerNode []HdinsightHadoopClusterSpecRolesWorkerNode `json:"workerNode" tf:"worker_node"`
	// +kubebuilder:validation:MaxItems=1
	ZookeeperNode []HdinsightHadoopClusterSpecRolesZookeeperNode `json:"zookeeperNode" tf:"zookeeper_node"`
}

type HdinsightHadoopClusterSpecStorageAccount struct {
	IsDefault          bool   `json:"isDefault" tf:"is_default"`
	StorageAccountKey  string `json:"-" sensitive:"true" tf:"storage_account_key"`
	StorageContainerID string `json:"storageContainerID" tf:"storage_container_id"`
}

type HdinsightHadoopClusterSpecStorageAccountGen2 struct {
	FilesystemID              string `json:"filesystemID" tf:"filesystem_id"`
	IsDefault                 bool   `json:"isDefault" tf:"is_default"`
	ManagedIdentityResourceID string `json:"managedIdentityResourceID" tf:"managed_identity_resource_id"`
	StorageResourceID         string `json:"storageResourceID" tf:"storage_resource_id"`
}

type HdinsightHadoopClusterSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	ClusterVersion string `json:"clusterVersion" tf:"cluster_version"`
	// +kubebuilder:validation:MaxItems=1
	ComponentVersion []HdinsightHadoopClusterSpecComponentVersion `json:"componentVersion" tf:"component_version"`
	// +kubebuilder:validation:MaxItems=1
	Gateway []HdinsightHadoopClusterSpecGateway `json:"gateway" tf:"gateway"`
	// +optional
	HttpsEndpoint string `json:"httpsEndpoint,omitempty" tf:"https_endpoint,omitempty"`
	Location      string `json:"location" tf:"location"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Metastores        []HdinsightHadoopClusterSpecMetastores `json:"metastores,omitempty" tf:"metastores,omitempty"`
	Name              string                                 `json:"name" tf:"name"`
	ResourceGroupName string                                 `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Roles []HdinsightHadoopClusterSpecRoles `json:"roles" tf:"roles"`
	// +optional
	SshEndpoint string `json:"sshEndpoint,omitempty" tf:"ssh_endpoint,omitempty"`
	// +optional
	StorageAccount []HdinsightHadoopClusterSpecStorageAccount `json:"storageAccount,omitempty" tf:"storage_account,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	StorageAccountGen2 []HdinsightHadoopClusterSpecStorageAccountGen2 `json:"storageAccountGen2,omitempty" tf:"storage_account_gen2,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	Tier string            `json:"tier" tf:"tier"`
	// +optional
	TlsMinVersion string `json:"tlsMinVersion,omitempty" tf:"tls_min_version,omitempty"`
}

type HdinsightHadoopClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *HdinsightHadoopClusterSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
	// +optional
	TerraFormLogs *base.TerraFormLogs `json:"terraformLogs,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// HdinsightHadoopClusterList is a list of HdinsightHadoopClusters
type HdinsightHadoopClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of HdinsightHadoopCluster CRD objects
	Items []HdinsightHadoopCluster `json:"items,omitempty"`
}
