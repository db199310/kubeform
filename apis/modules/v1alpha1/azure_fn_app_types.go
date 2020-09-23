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

type AzureFnApp struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzureFnAppSpec   `json:"spec,omitempty"`
	Status            AzureFnAppStatus `json:"status,omitempty"`
}

type AzureFnAppSpec struct {
	// +optional
	SecretRef   *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
	ProviderRef core.LocalObjectReference  `json:"providerRef" tf:"-"`
	// +optional
	Source string `json:"source" tf:"source"`

	// +optional
	// A map object for Active Directory. please refer https://www.terraform.io/docs/providers/azurerm/r/function_app.html
	ActiveDirectory json.RawMessage `json:"activeDirectory,omitempty" tf:"active_directory,omitempty"`
	// App resourcess name prefix.
	AppPrefix string `json:"appPrefix" tf:"app_prefix"`
	// +optional
	// App Settings. Package deploy configured
	AppSettings map[string]string `json:"appSettings,omitempty" tf:"app_settings,omitempty"`
	// +optional
	// App insights type
	ApplicationInsightsType string `json:"applicationInsightsType,omitempty" tf:"application_insights_type,omitempty"`
	// +optional
	// App service plan kind. Should be able to accomodate the fn app
	AspKind string `json:"aspKind,omitempty" tf:"asp_kind,omitempty"`
	// +optional
	// App Service plan max worker count
	AspMaxWorkerCnt json.Number `json:"aspMaxWorkerCnt,omitempty" tf:"asp_max_worker_cnt,omitempty"`
	// +optional
	// Reserved field for App Service plan (Linux). Boolean
	AspReserved *bool `json:"aspReserved,omitempty" tf:"asp_reserved,omitempty"`
	// +optional
	// App Service plan capacity
	AspSkuCap json.Number `json:"aspSkuCap,omitempty" tf:"asp_sku_cap,omitempty"`
	// +optional
	// App service plan size
	AspSkuSize string `json:"aspSkuSize,omitempty" tf:"asp_sku_size,omitempty"`
	// +optional
	// Tier of the app service plan
	AspSkuTier string `json:"aspSkuTier,omitempty" tf:"asp_sku_tier,omitempty"`
	// +optional
	// Authentication Settings
	AuthSettings map[string]AzureFnAppAuthSettings `json:"authSettings,omitempty" tf:"auth_settings,omitempty"`
	// +optional
	// Should client affinity be enabled?
	ClientAffinityEnabled *bool `json:"clientAffinityEnabled,omitempty" tf:"client_affinity_enabled,omitempty"`
	// +optional
	// connection strings for fn app
	ConnectionStrings []AzureFnAppConnectionStrings `json:"connectionStrings,omitempty" tf:"connection_strings,omitempty"`
	// +optional
	// Require application insights resource?
	CreateApplicationInsightsResource *bool `json:"createApplicationInsightsResource,omitempty" tf:"create_application_insights_resource,omitempty"`
	// Environment
	Environment string `json:"environment" tf:"environment"`
	// +optional
	// Existing App Service plan name
	ExistingAspName string `json:"existingAspName,omitempty" tf:"existing_asp_name,omitempty"`
	// +optional
	// Existing App Service plan resource Group
	ExistingAspResGrpName string `json:"existingAspResGrpName,omitempty" tf:"existing_asp_res_grp_name,omitempty"`
	// +optional
	// Existing service plan enabled?
	ExistingServicePlanEnabled *bool `json:"existingServicePlanEnabled,omitempty" tf:"existing_service_plan_enabled,omitempty"`
	// +optional
	// Additional tags for the App Service resources, in addition to the resource group tags.
	FnAppAdditionalTags map[string]string `json:"fnAppAdditionalTags,omitempty" tf:"fn_app_additional_tags,omitempty"`
	// +optional
	// Function App location
	FnAppLocation string `json:"fnAppLocation,omitempty" tf:"fn_app_location,omitempty"`
	// +optional
	// Should fn app be enabled?
	FnEnabled *bool `json:"fnEnabled,omitempty" tf:"fn_enabled,omitempty"`
	// +optional
	// Is Fn app required?
	FnRequired *bool `json:"fnRequired,omitempty" tf:"fn_required,omitempty"`
	// +optional
	// Run time version of the Fn app
	FnappVersion string `json:"fnappVersion,omitempty" tf:"fnapp_version,omitempty"`
	// +optional
	// identity for fn app. please refer https://www.terraform.io/docs/providers/azurerm/r/function_app.html
	Identity json.RawMessage `json:"identity,omitempty" tf:"identity,omitempty"`
	// +optional
	// Subnet IDS for VNet integration
	IntegrationSubnetID string `json:"integrationSubnetID,omitempty" tf:"integration_subnet_id,omitempty"`
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
	ProjectStream string `json:"projectStream,omitempty" tf:"projectStream,omitempty"`
	// region
	Region string `json:"region" tf:"region"`
	// +optional
	// releaseVersion
	ReleaseVersion json.RawMessage `json:"releaseVersion,omitempty" tf:"releaseVersion,omitempty"`
	// Resource Group name where the fn app needs to be created
	ResGrpName string `json:"resGrpName" tf:"res_grp_name"`
	// +optional
	// Site config block for Fn app
	SiteConfig map[string]AzureFnAppSiteConfig `json:"siteConfig,omitempty" tf:"site_config,omitempty"`
	// +optional
	// Site config core parameters for Fn app
	SiteConfigCors map[string]AzureFnAppSiteConfigCors `json:"siteConfigCors,omitempty" tf:"site_config_cors,omitempty"`
	// +optional
	// site config ip restrictions block parameters for fn app
	SiteConfigIPRestrictions json.RawMessage `json:"siteConfigIPRestrictions,omitempty" tf:"site_config_ip_restrictions,omitempty"`
	// +optional
	// Zip file location to be used to do the deployment. Should be publicly accessible
	Sourcezip string `json:"sourcezip,omitempty" tf:"sourcezip,omitempty"`
	// +optional
	// Vnet integration required for the function app?
	VnetIntegrationRequired *bool `json:"vnetIntegrationRequired,omitempty" tf:"vnet_integration_required,omitempty"`
}

type AzureFnAppAuthSettings struct {
	// +optional
	AuthEnabled *bool `json:"authEnabled,omitempty" tf:"auth_enabled,omitempty"`
}

type AzureFnAppConnectionStrings struct {
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
	// +optional
	Value string `json:"value,omitempty" tf:"value,omitempty"`
}

type AzureFnAppSiteConfig struct {
	// +optional
	AlwaysOn *bool `json:"alwaysOn,omitempty" tf:"always_on,omitempty"`
	// +optional
	FtpsState string `json:"ftpsState,omitempty" tf:"ftps_state,omitempty"`
	// +optional
	Http2Enabled *bool `json:"http2Enabled,omitempty" tf:"http2_enabled,omitempty"`
	// +optional
	LinuxFxVersion string `json:"linuxFxVersion,omitempty" tf:"linux_fx_version,omitempty"`
	// +optional
	Use32BitWorkerProcess *bool `json:"use32BitWorkerProcess,omitempty" tf:"use_32_bit_worker_process,omitempty"`
	// +optional
	WebsocketsEnabled *bool `json:"websocketsEnabled,omitempty" tf:"websockets_enabled,omitempty"`
}

type AzureFnAppSiteConfigCors struct {
	// +optional
	AllowedOrigins []string `json:"allowedOrigins,omitempty" tf:"allowed_origins,omitempty"`
	// +optional
	SupportCredentials string `json:"supportCredentials,omitempty" tf:"support_credentials,omitempty"`
}

type AzureFnAppOutput struct {
	// Map output of the App Service Plans
	// +optional
	AppServicePlans string `json:"appServicePlans" tf:"app_service_plans"`
	// Map output of the Fnapp Services
	// +optional
	FnApp string `json:"fnApp" tf:"fn_app"`
}

type AzureFnAppStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AzureFnAppOutput `json:"output,omitempty"`
	// +optional
	State string `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzureFnAppList is a list of AzureFnApps
type AzureFnAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzureFnApp CRD objects
	Items []AzureFnApp `json:"items,omitempty"`
}