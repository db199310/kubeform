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

type GoogleOrganizationPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleOrganizationPolicySpec   `json:"spec,omitempty"`
	Status            GoogleOrganizationPolicyStatus `json:"status,omitempty"`
}

type GoogleOrganizationPolicySpecListPolicyAllow struct {
	Values []string `json:"values"`
	All    bool     `json:"all"`
}

type GoogleOrganizationPolicySpecListPolicyDeny struct {
	All    bool     `json:"all"`
	Values []string `json:"values"`
}

type GoogleOrganizationPolicySpecListPolicy struct {
	Allow          []GoogleOrganizationPolicySpecListPolicy `json:"allow"`
	Deny           []GoogleOrganizationPolicySpecListPolicy `json:"deny"`
	SuggestedValue string                                   `json:"suggested_value"`
}

type GoogleOrganizationPolicySpecRestorePolicy struct {
	Default bool `json:"default"`
}

type GoogleOrganizationPolicySpecBooleanPolicy struct {
	Enforced bool `json:"enforced"`
}

type GoogleOrganizationPolicySpec struct {
	ListPolicy    []GoogleOrganizationPolicySpec `json:"list_policy"`
	Version       int                            `json:"version"`
	Etag          string                         `json:"etag"`
	UpdateTime    string                         `json:"update_time"`
	RestorePolicy []GoogleOrganizationPolicySpec `json:"restore_policy"`
	OrgId         string                         `json:"org_id"`
	Constraint    string                         `json:"constraint"`
	BooleanPolicy []GoogleOrganizationPolicySpec `json:"boolean_policy"`
}

type GoogleOrganizationPolicyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleOrganizationPolicyList is a list of GoogleOrganizationPolicys
type GoogleOrganizationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleOrganizationPolicy CRD objects
	Items []GoogleOrganizationPolicy `json:"items,omitempty"`
}
