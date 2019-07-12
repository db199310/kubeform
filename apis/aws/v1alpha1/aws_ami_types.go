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

type AwsAmi struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAmiSpec   `json:"spec,omitempty"`
	Status            AwsAmiStatus `json:"status,omitempty"`
}

type AwsAmiSpecEphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
}

type AwsAmiSpecEbsBlockDevice struct {
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
	Encrypted           bool   `json:"encrypted"`
	Iops                int    `json:"iops"`
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
}

type AwsAmiSpec struct {
	VirtualizationType   string            `json:"virtualization_type"`
	Description          string            `json:"description"`
	EphemeralBlockDevice []AwsAmiSpec      `json:"ephemeral_block_device"`
	RamdiskId            string            `json:"ramdisk_id"`
	RootDeviceName       string            `json:"root_device_name"`
	RootSnapshotId       string            `json:"root_snapshot_id"`
	SriovNetSupport      string            `json:"sriov_net_support"`
	Tags                 map[string]string `json:"tags"`
	EnaSupport           bool              `json:"ena_support"`
	ManageEbsSnapshots   bool              `json:"manage_ebs_snapshots"`
	Name                 string            `json:"name"`
	ImageLocation        string            `json:"image_location"`
	EbsBlockDevice       []AwsAmiSpec      `json:"ebs_block_device"`
	KernelId             string            `json:"kernel_id"`
	Architecture         string            `json:"architecture"`
}

type AwsAmiStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsAmiList is a list of AwsAmis
type AwsAmiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAmi CRD objects
	Items []AwsAmi `json:"items,omitempty"`
}
