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

type GoogleLoggingProjectExclusion struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleLoggingProjectExclusionSpec   `json:"spec,omitempty"`
	Status            GoogleLoggingProjectExclusionStatus `json:"status,omitempty"`
}

type GoogleLoggingProjectExclusionSpec struct {
	Disabled    bool   `json:"disabled"`
	Filter      string `json:"filter"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Project     string `json:"project"`
}

type GoogleLoggingProjectExclusionStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleLoggingProjectExclusionList is a list of GoogleLoggingProjectExclusions
type GoogleLoggingProjectExclusionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleLoggingProjectExclusion CRD objects
	Items []GoogleLoggingProjectExclusion `json:"items,omitempty"`
}
