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
// +kubebuilder:storageversion

type SDPAzadfv1 struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SDPAzadfv1Spec   `json:"spec,omitempty"`
	Status            SDPAzadfv1Status `json:"status,omitempty"`
}

type SDPAzadfv1Spec struct {
	// +optional
	SecretRef   *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
	ProviderRef core.LocalObjectReference  `json:"providerRef" tf:"-"`
	// +optional
	Source string `json:"source" tf:"source"`

	// +optional
	// The tags to associate with your resource group.
	AdditionalTags map[string]string `json:"additionalTags,omitempty" tf:"additional_tags,omitempty"`
	// Environment. Upto 5 character. For e.g. dev, dev01 , prd01
	Environment string `json:"environment" tf:"environment"`
	// +optional
	// List of Github Configuration
	GithubConfiguration []SDPAzadfv1GithubConfiguration `json:"githubConfiguration,omitempty" tf:"github_configuration,omitempty"`
	// +optional
	// Instance number
	Instance string `json:"instance,omitempty" tf:"instance,omitempty"`
	// NamePrefix for the data factory
	NamePrefix string `json:"namePrefix" tf:"namePrefix"`
	// owner
	Owner string `json:"owner" tf:"owner"`
	// +optional
	// placement
	Placement string `json:"placement,omitempty" tf:"placement,omitempty"`
	// +optional
	// project stream name
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	//  4 character project stream name/code
	ProjectStream string `json:"projectStream" tf:"projectStream"`
	// region. Choose from australia, europe, asia, europe
	Region string `json:"region" tf:"region"`
	// +optional
	// releaseVersion
	ReleaseVersion string `json:"releaseVersion,omitempty" tf:"releaseVersion,omitempty"`
	// +optional
	// Resource Group name where the Data Factory needs to be created
	ResourceGroupName string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`
	// +optional
	// List of VSTS Configuration
	VstsConfiguration []SDPAzadfv1VstsConfiguration `json:"vstsConfiguration,omitempty" tf:"vsts_configuration,omitempty"`
	//  4 character project stream name/code
	WorkStream string `json:"workStream" tf:"workStream"`
}

type SDPAzadfv1GithubConfiguration struct {
	// +optional
	AccountName string `json:"accountName,omitempty" tf:"account_name,omitempty"`
	// +optional
	BranchName string `json:"branchName,omitempty" tf:"branch_name,omitempty"`
	// +optional
	GitURL string `json:"gitURL,omitempty" tf:"git_url,omitempty"`
	// +optional
	RepositoryName string `json:"repositoryName,omitempty" tf:"repository_name,omitempty"`
	// +optional
	RootFolder string `json:"rootFolder,omitempty" tf:"root_folder,omitempty"`
}

type SDPAzadfv1VstsConfiguration struct {
	// +optional
	AccountName string `json:"accountName,omitempty" tf:"account_name,omitempty"`
	// +optional
	BranchName string `json:"branchName,omitempty" tf:"branch_name,omitempty"`
	// +optional
	ProjectName string `json:"projectName,omitempty" tf:"project_name,omitempty"`
	// +optional
	RepositoryName string `json:"repositoryName,omitempty" tf:"repository_name,omitempty"`
	// +optional
	RootFolder string `json:"rootFolder,omitempty" tf:"root_folder,omitempty"`
	// +optional
	TenantID string `json:"tenantID,omitempty" tf:"tenant_id,omitempty"`
}

type SDPAzadfv1Output struct {
	// ID of the DataFactory
	// +optional
	DatafactoryID string `json:"datafactoryID" tf:"datafactory_id"`
	// Name of the DataFactory
	// +optional
	DatafactoryName string `json:"datafactoryName" tf:"datafactory_name"`
	//
	// +optional
	Identity string `json:"identity" tf:"identity"`
}

type SDPAzadfv1Status struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SDPAzadfv1Output `json:"output,omitempty"`
	// +optional
	State string `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
	// +optional
	TerraformErrors []string `json:"terraformErrors,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SDPAzadfv1List is a list of SDPAzadfv1s
type SDPAzadfv1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SDPAzadfv1 CRD objects
	Items []SDPAzadfv1 `json:"items,omitempty"`
}
