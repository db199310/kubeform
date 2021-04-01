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
// +kubebuilder:storageversion

type SDPAzAppv1 struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SDPAzAppv1Spec   `json:"spec,omitempty"`
	Status            SDPAzAppv1Status `json:"status,omitempty"`
}

type SDPAzAppv1Spec struct {
	// +optional
	SecretRef   *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
	ProviderRef core.LocalObjectReference  `json:"providerRef" tf:"-"`
	// +optional
	Source string `json:"source" tf:"source"`

	// +optional
	// A map object for Active Directory. please refer https://www.terraform.io/docs/providers/azurerm/r/function_app.html
	ActiveDirectory json.RawMessage `json:"activeDirectory,omitempty" tf:"active_directory,omitempty"`
	// +optional
	// App command line
	AppCommandLine string `json:"appCommandLine,omitempty" tf:"app_command_line,omitempty"`
	// +optional
	// App Settings
	AppSettings map[string]string `json:"appSettings,omitempty" tf:"app_settings,omitempty"`
	// +optional
	// App insights type
	ApplicationInsightsType string `json:"applicationInsightsType,omitempty" tf:"application_insights_type,omitempty"`
	// +optional
	// Additional tags for the App Service resources, in addition to the resource group tags.
	AppsvcAdditionalTags map[string]string `json:"appsvcAdditionalTags,omitempty" tf:"appsvc_additional_tags,omitempty"`
	// +optional
	// Authentication Settings
	AuthSettings map[string]SDPAzAppv1AuthSettings `json:"authSettings,omitempty" tf:"auth_settings,omitempty"`
	// +optional
	// connection strings for appsvc app
	ConnectionStrings []SDPAzAppv1ConnectionStrings `json:"connectionStrings,omitempty" tf:"connection_strings,omitempty"`
	// +optional
	// Dotnet framework version
	DotnetFrameworkVersion string `json:"dotnetFrameworkVersion,omitempty" tf:"dotnet_framework_version,omitempty"`
	// Environment. upto 5 Character. For e.g. dev, dev01, prd01
	Environment string `json:"environment" tf:"environment"`
	// +optional
	// Existing App Services plan name.
	ExistingAspName string `json:"existingAspName,omitempty" tf:"existing_asp_name,omitempty"`
	// +optional
	// Existing App Service plan resource Group
	ExistingAspResourceGroupName string `json:"existingAspResourceGroupName,omitempty" tf:"existing_asp_resource_group_name,omitempty"`
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
	// Java container - JAVA ,JETTY, TOMCAT
	JavaContainer string `json:"javaContainer,omitempty" tf:"java_container,omitempty"`
	// +optional
	// Java Container Version
	JavaContainerVersion string `json:"javaContainerVersion,omitempty" tf:"java_container_version,omitempty"`
	// +optional
	// Java Version
	JavaVersion string `json:"javaVersion,omitempty" tf:"java_version,omitempty"`
	// +optional
	// Linux Docker container image
	LinuxFxVersion string `json:"linuxFxVersion,omitempty" tf:"linux_fx_version,omitempty"`
	// App Service resourcess name prefix.
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
	// project stream name
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// project stream name
	ProjectStream string `json:"projectStream" tf:"projectStream"`
	// +optional
	// Python Version
	PythonVersion string `json:"pythonVersion,omitempty" tf:"python_version,omitempty"`
	// region
	Region string `json:"region" tf:"region"`
	// +optional
	// releaseVersion
	ReleaseVersion string `json:"releaseVersion,omitempty" tf:"releaseVersion,omitempty"`
	// +optional
	// The App Service resources group name.
	ResourceGroupName string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`
	// +optional
	// Site config block for appsvc
	SiteConfig map[string]SDPAzAppv1SiteConfig `json:"siteConfig,omitempty" tf:"site_config,omitempty"`
	// +optional
	// site config cors parameters for appsvc
	SiteConfigCors map[string]SDPAzAppv1SiteConfigCors `json:"siteConfigCors,omitempty" tf:"site_config_cors,omitempty"`
	// +optional
	// site config ip restrctions block parameters for appsvc
	SiteConfigIPRestrictions json.RawMessage `json:"siteConfigIPRestrictions,omitempty" tf:"site_config_ip_restrictions,omitempty"`
	// +optional
	// Storage account to store app logs
	StorageAccount json.RawMessage `json:"storageAccount,omitempty" tf:"storage_account,omitempty"`
	// +optional
	// Windows Docker container image
	WindowsFxVersion string `json:"windowsFxVersion,omitempty" tf:"windows_fx_version,omitempty"`
	//  3 character workstream name/code
	WorkStream string `json:"workStream" tf:"workStream"`
}

type SDPAzAppv1AuthSettings struct {
	// +optional
	AuthEnabled *bool `json:"authEnabled,omitempty" tf:"auth_enabled,omitempty"`
}

type SDPAzAppv1ConnectionStrings struct {
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
	// +optional
	Value string `json:"value,omitempty" tf:"value,omitempty"`
}

type SDPAzAppv1SiteConfig struct {
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

type SDPAzAppv1SiteConfigCors struct {
	// +optional
	AllowedOrigins []string `json:"allowedOrigins,omitempty" tf:"allowed_origins,omitempty"`
	// +optional
	SupportCredentials string `json:"supportCredentials,omitempty" tf:"support_credentials,omitempty"`
}

type SDPAzAppv1Output struct{}

type SDPAzAppv1Status struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SDPAzAppv1Output `json:"output,omitempty"`
	// +optional
	State string `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
	// +optional
	TerraformErrors []string `json:"terraformErrors,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SDPAzAppv1List is a list of SDPAzAppv1s
type SDPAzAppv1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SDPAzAppv1 CRD objects
	Items []SDPAzAppv1 `json:"items,omitempty"`
}
