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

type SpringCloudService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpringCloudServiceSpec   `json:"spec,omitempty"`
	Status            SpringCloudServiceStatus `json:"status,omitempty"`
}

type SpringCloudServiceSpecConfigServerGitSettingHttpBasicAuth struct {
	Password string `json:"-" sensitive:"true" tf:"password"`
	Username string `json:"username" tf:"username"`
}

type SpringCloudServiceSpecConfigServerGitSettingRepositoryHttpBasicAuth struct {
	Password string `json:"-" sensitive:"true" tf:"password"`
	Username string `json:"username" tf:"username"`
}

type SpringCloudServiceSpecConfigServerGitSettingRepositorySshAuth struct {
	// +optional
	HostKey string `json:"-" sensitive:"true" tf:"host_key,omitempty"`
	// +optional
	HostKeyAlgorithm string `json:"hostKeyAlgorithm,omitempty" tf:"host_key_algorithm,omitempty"`
	PrivateKey       string `json:"-" sensitive:"true" tf:"private_key"`
	// +optional
	StrictHostKeyCheckingEnabled bool `json:"strictHostKeyCheckingEnabled,omitempty" tf:"strict_host_key_checking_enabled,omitempty"`
}

type SpringCloudServiceSpecConfigServerGitSettingRepository struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HttpBasicAuth []SpringCloudServiceSpecConfigServerGitSettingRepositoryHttpBasicAuth `json:"httpBasicAuth,omitempty" tf:"http_basic_auth,omitempty"`
	// +optional
	Label string `json:"label,omitempty" tf:"label,omitempty"`
	Name  string `json:"name" tf:"name"`
	// +optional
	Pattern []string `json:"pattern,omitempty" tf:"pattern,omitempty"`
	// +optional
	SearchPaths []string `json:"searchPaths,omitempty" tf:"search_paths,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SshAuth []SpringCloudServiceSpecConfigServerGitSettingRepositorySshAuth `json:"sshAuth,omitempty" tf:"ssh_auth,omitempty"`
	Uri     string                                                          `json:"uri" tf:"uri"`
}

type SpringCloudServiceSpecConfigServerGitSettingSshAuth struct {
	// +optional
	HostKey string `json:"-" sensitive:"true" tf:"host_key,omitempty"`
	// +optional
	HostKeyAlgorithm string `json:"hostKeyAlgorithm,omitempty" tf:"host_key_algorithm,omitempty"`
	PrivateKey       string `json:"-" sensitive:"true" tf:"private_key"`
	// +optional
	StrictHostKeyCheckingEnabled bool `json:"strictHostKeyCheckingEnabled,omitempty" tf:"strict_host_key_checking_enabled,omitempty"`
}

type SpringCloudServiceSpecConfigServerGitSetting struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HttpBasicAuth []SpringCloudServiceSpecConfigServerGitSettingHttpBasicAuth `json:"httpBasicAuth,omitempty" tf:"http_basic_auth,omitempty"`
	// +optional
	Label string `json:"label,omitempty" tf:"label,omitempty"`
	// +optional
	Repository []SpringCloudServiceSpecConfigServerGitSettingRepository `json:"repository,omitempty" tf:"repository,omitempty"`
	// +optional
	SearchPaths []string `json:"searchPaths,omitempty" tf:"search_paths,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SshAuth []SpringCloudServiceSpecConfigServerGitSettingSshAuth `json:"sshAuth,omitempty" tf:"ssh_auth,omitempty"`
	Uri     string                                                `json:"uri" tf:"uri"`
}

type SpringCloudServiceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	ConfigServerGitSetting []SpringCloudServiceSpecConfigServerGitSetting `json:"configServerGitSetting,omitempty" tf:"config_server_git_setting,omitempty"`
	Location               string                                         `json:"location" tf:"location"`
	Name                   string                                         `json:"name" tf:"name"`
	ResourceGroupName      string                                         `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SpringCloudServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SpringCloudServiceSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SpringCloudServiceList is a list of SpringCloudServices
type SpringCloudServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SpringCloudService CRD objects
	Items []SpringCloudService `json:"items,omitempty"`
}
