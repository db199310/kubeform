package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type AzurermPacketCapture struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermPacketCaptureSpec   `json:"spec,omitempty"`
	Status            AzurermPacketCaptureStatus `json:"status,omitempty"`
}

type AzurermPacketCaptureSpecFilter struct {
	RemoteIpAddress string `json:"remote_ip_address"`
	RemotePort      string `json:"remote_port"`
	LocalIpAddress  string `json:"local_ip_address"`
	LocalPort       string `json:"local_port"`
	Protocol        string `json:"protocol"`
}

type AzurermPacketCaptureSpecStorageLocation struct {
	StoragePath      string `json:"storage_path"`
	FilePath         string `json:"file_path"`
	StorageAccountId string `json:"storage_account_id"`
}

type AzurermPacketCaptureSpec struct {
	ResourceGroupName      string                     `json:"resource_group_name"`
	MaximumBytesPerPacket  int                        `json:"maximum_bytes_per_packet"`
	MaximumCaptureDuration int                        `json:"maximum_capture_duration"`
	Filter                 []AzurermPacketCaptureSpec `json:"filter"`
	Name                   string                     `json:"name"`
	NetworkWatcherName     string                     `json:"network_watcher_name"`
	TargetResourceId       string                     `json:"target_resource_id"`
	MaximumBytesPerSession int                        `json:"maximum_bytes_per_session"`
	StorageLocation        []AzurermPacketCaptureSpec `json:"storage_location"`
}

type AzurermPacketCaptureStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermPacketCaptureList is a list of AzurermPacketCaptures
type AzurermPacketCaptureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermPacketCapture CRD objects
	Items []AzurermPacketCapture `json:"items,omitempty"`
}
