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

type RecoveryReplicatedVm struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RecoveryReplicatedVmSpec   `json:"spec,omitempty"`
	Status            RecoveryReplicatedVmStatus `json:"status,omitempty"`
}

type RecoveryReplicatedVmSpecManagedDisk struct {
	DiskID                  string `json:"diskID" tf:"disk_id"`
	StagingStorageAccountID string `json:"stagingStorageAccountID" tf:"staging_storage_account_id"`
	TargetDiskType          string `json:"targetDiskType" tf:"target_disk_type"`
	TargetReplicaDiskType   string `json:"targetReplicaDiskType" tf:"target_replica_disk_type"`
	TargetResourceGroupID   string `json:"targetResourceGroupID" tf:"target_resource_group_id"`
}

type RecoveryReplicatedVmSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ManagedDisk                           []RecoveryReplicatedVmSpecManagedDisk `json:"managedDisk,omitempty" tf:"managed_disk,omitempty"`
	Name                                  string                                `json:"name" tf:"name"`
	RecoveryReplicationPolicyID           string                                `json:"recoveryReplicationPolicyID" tf:"recovery_replication_policy_id"`
	RecoveryVaultName                     string                                `json:"recoveryVaultName" tf:"recovery_vault_name"`
	ResourceGroupName                     string                                `json:"resourceGroupName" tf:"resource_group_name"`
	SourceRecoveryFabricName              string                                `json:"sourceRecoveryFabricName" tf:"source_recovery_fabric_name"`
	SourceRecoveryProtectionContainerName string                                `json:"sourceRecoveryProtectionContainerName" tf:"source_recovery_protection_container_name"`
	SourceVmID                            string                                `json:"sourceVmID" tf:"source_vm_id"`
	// +optional
	TargetAvailabilitySetID             string `json:"targetAvailabilitySetID,omitempty" tf:"target_availability_set_id,omitempty"`
	TargetRecoveryFabricID              string `json:"targetRecoveryFabricID" tf:"target_recovery_fabric_id"`
	TargetRecoveryProtectionContainerID string `json:"targetRecoveryProtectionContainerID" tf:"target_recovery_protection_container_id"`
	TargetResourceGroupID               string `json:"targetResourceGroupID" tf:"target_resource_group_id"`
}

type RecoveryReplicatedVmStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *RecoveryReplicatedVmSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
	// +optional
	TerraformErrors []string `json:"terraformErrors,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RecoveryReplicatedVmList is a list of RecoveryReplicatedVms
type RecoveryReplicatedVmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RecoveryReplicatedVm CRD objects
	Items []RecoveryReplicatedVm `json:"items,omitempty"`
}
