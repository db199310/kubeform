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

type NetappVolume struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetappVolumeSpec   `json:"spec,omitempty"`
	Status            NetappVolumeStatus `json:"status,omitempty"`
}

type NetappVolumeSpecExportPolicyRule struct {
	AllowedClients []string `json:"allowedClients" tf:"allowed_clients"`
	// +optional
	// Deprecated
	CifsEnabled bool `json:"cifsEnabled,omitempty" tf:"cifs_enabled,omitempty"`
	// +optional
	// Deprecated
	Nfsv3Enabled bool `json:"nfsv3Enabled,omitempty" tf:"nfsv3_enabled,omitempty"`
	// +optional
	// Deprecated
	Nfsv4Enabled bool `json:"nfsv4Enabled,omitempty" tf:"nfsv4_enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	ProtocolsEnabled []string `json:"protocolsEnabled,omitempty" tf:"protocols_enabled,omitempty"`
	RuleIndex        int64    `json:"ruleIndex" tf:"rule_index"`
	// +optional
	UnixReadOnly bool `json:"unixReadOnly,omitempty" tf:"unix_read_only,omitempty"`
	// +optional
	UnixReadWrite bool `json:"unixReadWrite,omitempty" tf:"unix_read_write,omitempty"`
}

type NetappVolumeSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	RemoteBackend *base.Backend `json:"remoteBackend,omitempty" tf:"-"`

	AccountName string `json:"accountName" tf:"account_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=5
	ExportPolicyRule []NetappVolumeSpecExportPolicyRule `json:"exportPolicyRule,omitempty" tf:"export_policy_rule,omitempty"`
	Location         string                             `json:"location" tf:"location"`
	// +optional
	MountIPAddresses []string `json:"mountIPAddresses,omitempty" tf:"mount_ip_addresses,omitempty"`
	Name             string   `json:"name" tf:"name"`
	PoolName         string   `json:"poolName" tf:"pool_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=2
	Protocols         []string `json:"protocols,omitempty" tf:"protocols,omitempty"`
	ResourceGroupName string   `json:"resourceGroupName" tf:"resource_group_name"`
	ServiceLevel      string   `json:"serviceLevel" tf:"service_level"`
	StorageQuotaInGb  int64    `json:"storageQuotaInGb" tf:"storage_quota_in_gb"`
	SubnetID          string   `json:"subnetID" tf:"subnet_id"`
	// +optional
	Tags       map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	VolumePath string            `json:"volumePath" tf:"volume_path"`
}

type NetappVolumeStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *NetappVolumeSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NetappVolumeList is a list of NetappVolumes
type NetappVolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NetappVolume CRD objects
	Items []NetappVolume `json:"items,omitempty"`
}
