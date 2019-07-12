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

type AwsNetworkAcl struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsNetworkAclSpec   `json:"spec,omitempty"`
	Status            AwsNetworkAclStatus `json:"status,omitempty"`
}

type AwsNetworkAclSpecEgress struct {
	Ipv6CidrBlock string `json:"ipv6_cidr_block"`
	IcmpCode      int    `json:"icmp_code"`
	FromPort      int    `json:"from_port"`
	ToPort        int    `json:"to_port"`
	RuleNo        int    `json:"rule_no"`
	Action        string `json:"action"`
	CidrBlock     string `json:"cidr_block"`
	Protocol      string `json:"protocol"`
	IcmpType      int    `json:"icmp_type"`
}

type AwsNetworkAclSpecIngress struct {
	FromPort      int    `json:"from_port"`
	ToPort        int    `json:"to_port"`
	Ipv6CidrBlock string `json:"ipv6_cidr_block"`
	IcmpType      int    `json:"icmp_type"`
	RuleNo        int    `json:"rule_no"`
	Action        string `json:"action"`
	Protocol      string `json:"protocol"`
	CidrBlock     string `json:"cidr_block"`
	IcmpCode      int    `json:"icmp_code"`
}

type AwsNetworkAclSpec struct {
	Egress    []AwsNetworkAclSpec `json:"egress"`
	Tags      map[string]string   `json:"tags"`
	OwnerId   string              `json:"owner_id"`
	VpcId     string              `json:"vpc_id"`
	SubnetId  string              `json:"subnet_id"`
	SubnetIds []string            `json:"subnet_ids"`
	Ingress   []AwsNetworkAclSpec `json:"ingress"`
}

type AwsNetworkAclStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsNetworkAclList is a list of AwsNetworkAcls
type AwsNetworkAclList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsNetworkAcl CRD objects
	Items []AwsNetworkAcl `json:"items,omitempty"`
}
