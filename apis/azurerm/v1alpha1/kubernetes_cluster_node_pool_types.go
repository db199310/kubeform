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

type KubernetesClusterNodePool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubernetesClusterNodePoolSpec   `json:"spec,omitempty"`
	Status            KubernetesClusterNodePoolStatus `json:"status,omitempty"`
}

type KubernetesClusterNodePoolSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AvailabilityZones []string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`
	// +optional
	EnableAutoScaling bool `json:"enableAutoScaling,omitempty" tf:"enable_auto_scaling,omitempty"`
	// +optional
	EnableNodePublicIP bool `json:"enableNodePublicIP,omitempty" tf:"enable_node_public_ip,omitempty"`
	// +optional
	EvictionPolicy      string `json:"evictionPolicy,omitempty" tf:"eviction_policy,omitempty"`
	KubernetesClusterID string `json:"kubernetesClusterID" tf:"kubernetes_cluster_id"`
	// +optional
	MaxCount int64 `json:"maxCount,omitempty" tf:"max_count,omitempty"`
	// +optional
	MaxPods int64 `json:"maxPods,omitempty" tf:"max_pods,omitempty"`
	// +optional
	MinCount int64 `json:"minCount,omitempty" tf:"min_count,omitempty"`
	// +optional
	Mode string `json:"mode,omitempty" tf:"mode,omitempty"`
	Name string `json:"name" tf:"name"`
	// +optional
	NodeCount int64 `json:"nodeCount,omitempty" tf:"node_count,omitempty"`
	// +optional
	NodeLabels map[string]string `json:"nodeLabels,omitempty" tf:"node_labels,omitempty"`
	// +optional
	NodeTaints []string `json:"nodeTaints,omitempty" tf:"node_taints,omitempty"`
	// +optional
	OrchestratorVersion string `json:"orchestratorVersion,omitempty" tf:"orchestrator_version,omitempty"`
	// +optional
	OsDiskSizeGb int64 `json:"osDiskSizeGb,omitempty" tf:"os_disk_size_gb,omitempty"`
	// +optional
	OsType string `json:"osType,omitempty" tf:"os_type,omitempty"`
	// +optional
	Priority string `json:"priority,omitempty" tf:"priority,omitempty"`
	// +optional
	SpotMaxPrice float64 `json:"spotMaxPrice,omitempty" tf:"spot_max_price,omitempty"`
	// +optional
	Tags   map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	VmSize string            `json:"vmSize" tf:"vm_size"`
	// +optional
	VnetSubnetID string `json:"vnetSubnetID,omitempty" tf:"vnet_subnet_id,omitempty"`
}

type KubernetesClusterNodePoolStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *KubernetesClusterNodePoolSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
	// +optional
	TerraFormLogs *base.TerraFormLogs `json:"terraformLogs,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// KubernetesClusterNodePoolList is a list of KubernetesClusterNodePools
type KubernetesClusterNodePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KubernetesClusterNodePool CRD objects
	Items []KubernetesClusterNodePool `json:"items,omitempty"`
}
