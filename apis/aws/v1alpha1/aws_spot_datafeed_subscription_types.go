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

type AwsSpotDatafeedSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSpotDatafeedSubscriptionSpec   `json:"spec,omitempty"`
	Status            AwsSpotDatafeedSubscriptionStatus `json:"status,omitempty"`
}

type AwsSpotDatafeedSubscriptionSpec struct {
	Bucket string `json:"bucket"`
	Prefix string `json:"prefix"`
}



type AwsSpotDatafeedSubscriptionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSpotDatafeedSubscriptionList is a list of AwsSpotDatafeedSubscriptions
type AwsSpotDatafeedSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSpotDatafeedSubscription CRD objects
	Items []AwsSpotDatafeedSubscription `json:"items,omitempty"`
}