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

type Eventhub struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventhubSpec   `json:"spec,omitempty"`
	Status            EventhubStatus `json:"status,omitempty"`
}

type EventhubSpecCaptureDescriptionDestination struct {
	ArchiveNameFormat string `json:"archiveNameFormat" tf:"archive_name_format"`
	BlobContainerName string `json:"blobContainerName" tf:"blob_container_name"`
	Name              string `json:"name" tf:"name"`
	StorageAccountID  string `json:"storageAccountID" tf:"storage_account_id"`
}

type EventhubSpecCaptureDescription struct {
	// +kubebuilder:validation:MaxItems=1
	Destination []EventhubSpecCaptureDescriptionDestination `json:"destination" tf:"destination"`
	Enabled     bool                                        `json:"enabled" tf:"enabled"`
	Encoding    string                                      `json:"encoding" tf:"encoding"`
	// +optional
	IntervalInSeconds int64 `json:"intervalInSeconds,omitempty" tf:"interval_in_seconds,omitempty"`
	// +optional
	SizeLimitInBytes int64 `json:"sizeLimitInBytes,omitempty" tf:"size_limit_in_bytes,omitempty"`
	// +optional
	SkipEmptyArchives bool `json:"skipEmptyArchives,omitempty" tf:"skip_empty_archives,omitempty"`
}

type EventhubSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	RemoteBackend *base.Backend `json:"remoteBackend,omitempty" tf:"-"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	CaptureDescription []EventhubSpecCaptureDescription `json:"captureDescription,omitempty" tf:"capture_description,omitempty"`
	MessageRetention   int64                            `json:"messageRetention" tf:"message_retention"`
	Name               string                           `json:"name" tf:"name"`
	NamespaceName      string                           `json:"namespaceName" tf:"namespace_name"`
	PartitionCount     int64                            `json:"partitionCount" tf:"partition_count"`
	// +optional
	PartitionIDS      []string `json:"partitionIDS,omitempty" tf:"partition_ids,omitempty"`
	ResourceGroupName string   `json:"resourceGroupName" tf:"resource_group_name"`
}

type EventhubStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *EventhubSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EventhubList is a list of Eventhubs
type EventhubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Eventhub CRD objects
	Items []Eventhub `json:"items,omitempty"`
}
