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

type VirtualMachineScaleSetExtension struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualMachineScaleSetExtensionSpec   `json:"spec,omitempty"`
	Status            VirtualMachineScaleSetExtensionStatus `json:"status,omitempty"`
}

type VirtualMachineScaleSetExtensionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	AutoUpgradeMinorVersion bool `json:"autoUpgradeMinorVersion,omitempty" tf:"auto_upgrade_minor_version,omitempty"`
	// +optional
	ForceUpdateTag string `json:"forceUpdateTag,omitempty" tf:"force_update_tag,omitempty"`
	Name           string `json:"name" tf:"name"`
	// +optional
	ProtectedSettings string `json:"-" sensitive:"true" tf:"protected_settings,omitempty"`
	// +optional
	ProvisionAfterExtensions []string `json:"provisionAfterExtensions,omitempty" tf:"provision_after_extensions,omitempty"`
	Publisher                string   `json:"publisher" tf:"publisher"`
	// +optional
	Settings                 string `json:"settings,omitempty" tf:"settings,omitempty"`
	Type                     string `json:"type" tf:"type"`
	TypeHandlerVersion       string `json:"typeHandlerVersion" tf:"type_handler_version"`
	VirtualMachineScaleSetID string `json:"virtualMachineScaleSetID" tf:"virtual_machine_scale_set_id"`
}

type VirtualMachineScaleSetExtensionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *VirtualMachineScaleSetExtensionSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VirtualMachineScaleSetExtensionList is a list of VirtualMachineScaleSetExtensions
type VirtualMachineScaleSetExtensionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VirtualMachineScaleSetExtension CRD objects
	Items []VirtualMachineScaleSetExtension `json:"items,omitempty"`
}