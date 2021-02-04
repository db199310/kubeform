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
	"encoding/json"

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

type SDPAzFnv1 struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SDPAzFnv1Spec   `json:"spec,omitempty"`
	Status            SDPAzFnv1Status `json:"status,omitempty"`
}

type SDPAzFnv1Spec struct {
	// +optional
	SecretRef     *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
	ProviderRef   core.LocalObjectReference  `json:"providerRef" tf:"-"`
	RemoteBackend *base.Backend              `json:"remoteBackend,omitempty" tf:"-"`

	// +optional
	Source string `json:"source" tf:"source"`

	// +optional
	// A map object for Active Directory. please refer https://www.terraform.io/docs/providers/azurerm/r/function_app.html
	ActiveDirectory json.RawMessage `json:"activeDirectory,omitempty" tf:"active_directory,omitempty"`
	// +optional
	// App Settings. Package deploy configured
	AppSettings map[string]string `json:"appSettings,omitempty" tf:"app_settings,omitempty"`
	// +optional
	// App insights type
	ApplicationInsightsType string `json:"applicationInsightsType,omitempty" tf:"application_insights_type,omitempty"`
	// +optional
	// Authentication Settings
	AuthSettings map[string]SDPAzFnv1AuthSettings `json:"authSettings,omitempty" tf:"auth_settings,omitempty"`
	// +optional
	// connection strings for fn app
	ConnectionStrings []SDPAzFnv1ConnectionStrings `json:"connectionStrings,omitempty" tf:"connection_strings,omitempty"`
	// Environment. Upto 5 character. For e.g. dev, dev01 , prd01
	Environment string `json:"environment" tf:"environment"`
	// +optional
	// Existing App Service plan name
	ExistingAspName string `json:"existingAspName,omitempty" tf:"existing_asp_name,omitempty"`
	// +optional
	// Existing App Service plan resource Group
	ExistingAspResourceGroupName string `json:"existingAspResourceGroupName,omitempty" tf:"existing_asp_resource_group_name,omitempty"`
	// +optional
	// Additional tags for the App Service resources, in addition to the resource group tags.
	FnAppAdditionalTags map[string]string `json:"fnAppAdditionalTags,omitempty" tf:"fn_app_additional_tags,omitempty"`
	// +optional
	// Hostname with the stratos.shell/stratos.shell.com suffix
	Host string `json:"host,omitempty" tf:"host,omitempty"`
	// +optional
	// Instance number
	Instance string `json:"instance,omitempty" tf:"instance,omitempty"`
	// +optional
	// Subnet IDS for VNet integration
	IntegrationSubnetID string `json:"integrationSubnetID,omitempty" tf:"integration_subnet_id,omitempty"`
	// +optional
	// Linux Docker image to use
	LinuxFxVersion string `json:"linuxFxVersion,omitempty" tf:"linux_fx_version,omitempty"`
	// name suffix for the function app
	NameSuffix string `json:"nameSuffix" tf:"nameSuffix"`
	// +optional
	// OS Type for the fn app. Should match with App Service plan
	OsType string `json:"osType,omitempty" tf:"os_type,omitempty"`
	// owner
	Owner string `json:"owner" tf:"owner"`
	// +optional
	// placement
	Placement string `json:"placement,omitempty" tf:"placement,omitempty"`
	// +optional
	// project name
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	//  4 character project stream name/code
	ProjectStream string `json:"projectStream" tf:"projectStream"`
	// region. Choose from australia, europe, asia, europe
	Region string `json:"region" tf:"region"`
	// +optional
	// releaseVersion
	ReleaseVersion string `json:"releaseVersion,omitempty" tf:"releaseVersion,omitempty"`
	// +optional
	// Resource Group name where the fn app needs to be created
	ResourceGroupName string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`
	// +optional
	// Run time version of the Fn app
	RuntimeVersion string `json:"runtimeVersion,omitempty" tf:"runtime_version,omitempty"`
	// +optional
	// Site config block for Fn app
	SiteConfig map[string]SDPAzFnv1SiteConfig `json:"siteConfig,omitempty" tf:"site_config,omitempty"`
	// +optional
	// Site config core parameters for Fn app
	SiteConfigCors map[string]SDPAzFnv1SiteConfigCors `json:"siteConfigCors,omitempty" tf:"site_config_cors,omitempty"`
	// +optional
	// site config ip restrictions block parameters for fn app
	SiteConfigIPRestrictions json.RawMessage `json:"siteConfigIPRestrictions,omitempty" tf:"site_config_ip_restrictions,omitempty"`
	//  3 character workstream name/code
	WorkStream string `json:"workStream" tf:"workStream"`
}

type SDPAzFnv1AuthSettings struct {
	// +optional
	AuthEnabled *bool `json:"authEnabled,omitempty" tf:"auth_enabled,omitempty"`
}

type SDPAzFnv1ConnectionStrings struct {
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
	// +optional
	Value string `json:"value,omitempty" tf:"value,omitempty"`
}

type SDPAzFnv1SiteConfig struct {
	// +optional
	AlwaysOn *bool `json:"alwaysOn,omitempty" tf:"always_on,omitempty"`
	// +optional
	FtpsState string `json:"ftpsState,omitempty" tf:"ftps_state,omitempty"`
	// +optional
	Http2Enabled *bool `json:"http2Enabled,omitempty" tf:"http2_enabled,omitempty"`
	// +optional
	Use32BitWorkerProcess *bool `json:"use32BitWorkerProcess,omitempty" tf:"use_32_bit_worker_process,omitempty"`
	// +optional
	WebsocketsEnabled *bool `json:"websocketsEnabled,omitempty" tf:"websockets_enabled,omitempty"`
}

type SDPAzFnv1SiteConfigCors struct {
	// +optional
	AllowedOrigins []string `json:"allowedOrigins,omitempty" tf:"allowed_origins,omitempty"`
	// +optional
	SupportCredentials *bool `json:"supportCredentials,omitempty" tf:"support_credentials,omitempty"`
}

type SDPAzFnv1Output struct{}

type SDPAzFnv1Status struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SDPAzFnv1Output `json:"output,omitempty"`
	// +optional
	State string `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SDPAzFnv1List is a list of SDPAzFnv1s
type SDPAzFnv1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SDPAzFnv1 CRD objects
	Items []SDPAzFnv1 `json:"items,omitempty"`
}
