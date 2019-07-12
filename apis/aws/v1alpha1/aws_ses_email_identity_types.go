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

type AwsSesEmailIdentity struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSesEmailIdentitySpec   `json:"spec,omitempty"`
	Status            AwsSesEmailIdentityStatus `json:"status,omitempty"`
}

type AwsSesEmailIdentitySpec struct {
	Email string `json:"email"`
	Arn   string `json:"arn"`
}

type AwsSesEmailIdentityStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSesEmailIdentityList is a list of AwsSesEmailIdentitys
type AwsSesEmailIdentityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSesEmailIdentity CRD objects
	Items []AwsSesEmailIdentity `json:"items,omitempty"`
}
