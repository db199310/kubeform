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

type Frontdoor struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FrontdoorSpec   `json:"spec,omitempty"`
	Status            FrontdoorStatus `json:"status,omitempty"`
}

type FrontdoorSpecBackendPoolBackend struct {
	Address string `json:"address" tf:"address"`
	// +optional
	Enabled    bool   `json:"enabled,omitempty" tf:"enabled,omitempty"`
	HostHeader string `json:"hostHeader" tf:"host_header"`
	HttpPort   int64  `json:"httpPort" tf:"http_port"`
	HttpsPort  int64  `json:"httpsPort" tf:"https_port"`
	// +optional
	Priority int64 `json:"priority,omitempty" tf:"priority,omitempty"`
	// +optional
	Weight int64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type FrontdoorSpecBackendPool struct {
	// +kubebuilder:validation:MaxItems=100
	Backend         []FrontdoorSpecBackendPoolBackend `json:"backend" tf:"backend"`
	HealthProbeName string                            `json:"healthProbeName" tf:"health_probe_name"`
	// +optional
	ID                string `json:"ID,omitempty" tf:"id,omitempty"`
	LoadBalancingName string `json:"loadBalancingName" tf:"load_balancing_name"`
	Name              string `json:"name" tf:"name"`
}

type FrontdoorSpecBackendPoolHealthProbe struct {
	// +optional
	ID string `json:"ID,omitempty" tf:"id,omitempty"`
	// +optional
	IntervalInSeconds int64  `json:"intervalInSeconds,omitempty" tf:"interval_in_seconds,omitempty"`
	Name              string `json:"name" tf:"name"`
	// +optional
	Path string `json:"path,omitempty" tf:"path,omitempty"`
	// +optional
	Protocol string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type FrontdoorSpecBackendPoolLoadBalancing struct {
	// +optional
	AdditionalLatencyMilliseconds int64 `json:"additionalLatencyMilliseconds,omitempty" tf:"additional_latency_milliseconds,omitempty"`
	// +optional
	ID   string `json:"ID,omitempty" tf:"id,omitempty"`
	Name string `json:"name" tf:"name"`
	// +optional
	SampleSize int64 `json:"sampleSize,omitempty" tf:"sample_size,omitempty"`
	// +optional
	SuccessfulSamplesRequired int64 `json:"successfulSamplesRequired,omitempty" tf:"successful_samples_required,omitempty"`
}

type FrontdoorSpecFrontendEndpointCustomHTTPSConfiguration struct {
	// +optional
	AzureKeyVaultCertificateSecretName string `json:"azureKeyVaultCertificateSecretName,omitempty" tf:"azure_key_vault_certificate_secret_name,omitempty"`
	// +optional
	AzureKeyVaultCertificateSecretVersion string `json:"azureKeyVaultCertificateSecretVersion,omitempty" tf:"azure_key_vault_certificate_secret_version,omitempty"`
	// +optional
	AzureKeyVaultCertificateVaultID string `json:"azureKeyVaultCertificateVaultID,omitempty" tf:"azure_key_vault_certificate_vault_id,omitempty"`
	// +optional
	CertificateSource string `json:"certificateSource,omitempty" tf:"certificate_source,omitempty"`
	// +optional
	MinimumTLSVersion string `json:"minimumTLSVersion,omitempty" tf:"minimum_tls_version,omitempty"`
	// +optional
	ProvisioningState string `json:"provisioningState,omitempty" tf:"provisioning_state,omitempty"`
	// +optional
	ProvisioningSubstate string `json:"provisioningSubstate,omitempty" tf:"provisioning_substate,omitempty"`
}

type FrontdoorSpecFrontendEndpoint struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CustomHTTPSConfiguration       []FrontdoorSpecFrontendEndpointCustomHTTPSConfiguration `json:"customHTTPSConfiguration,omitempty" tf:"custom_https_configuration,omitempty"`
	CustomHTTPSProvisioningEnabled bool                                                    `json:"customHTTPSProvisioningEnabled" tf:"custom_https_provisioning_enabled"`
	HostName                       string                                                  `json:"hostName" tf:"host_name"`
	// +optional
	ID   string `json:"ID,omitempty" tf:"id,omitempty"`
	Name string `json:"name" tf:"name"`
	// +optional
	SessionAffinityEnabled bool `json:"sessionAffinityEnabled,omitempty" tf:"session_affinity_enabled,omitempty"`
	// +optional
	SessionAffinityTtlSeconds int64 `json:"sessionAffinityTtlSeconds,omitempty" tf:"session_affinity_ttl_seconds,omitempty"`
	// +optional
	WebApplicationFirewallPolicyLinkID string `json:"webApplicationFirewallPolicyLinkID,omitempty" tf:"web_application_firewall_policy_link_id,omitempty"`
}

type FrontdoorSpecRoutingRuleForwardingConfiguration struct {
	BackendPoolName string `json:"backendPoolName" tf:"backend_pool_name"`
	// +optional
	CacheEnabled bool `json:"cacheEnabled,omitempty" tf:"cache_enabled,omitempty"`
	// +optional
	CacheQueryParameterStripDirective string `json:"cacheQueryParameterStripDirective,omitempty" tf:"cache_query_parameter_strip_directive,omitempty"`
	// +optional
	CacheUseDynamicCompression bool `json:"cacheUseDynamicCompression,omitempty" tf:"cache_use_dynamic_compression,omitempty"`
	// +optional
	CustomForwardingPath string `json:"customForwardingPath,omitempty" tf:"custom_forwarding_path,omitempty"`
	// +optional
	ForwardingProtocol string `json:"forwardingProtocol,omitempty" tf:"forwarding_protocol,omitempty"`
}

type FrontdoorSpecRoutingRuleRedirectConfiguration struct {
	// +optional
	CustomFragment string `json:"customFragment,omitempty" tf:"custom_fragment,omitempty"`
	// +optional
	CustomHost string `json:"customHost,omitempty" tf:"custom_host,omitempty"`
	// +optional
	CustomPath string `json:"customPath,omitempty" tf:"custom_path,omitempty"`
	// +optional
	CustomQueryString string `json:"customQueryString,omitempty" tf:"custom_query_string,omitempty"`
	RedirectProtocol  string `json:"redirectProtocol" tf:"redirect_protocol"`
	RedirectType      string `json:"redirectType" tf:"redirect_type"`
}

type FrontdoorSpecRoutingRule struct {
	// +kubebuilder:validation:MaxItems=2
	AcceptedProtocols []string `json:"acceptedProtocols" tf:"accepted_protocols"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ForwardingConfiguration []FrontdoorSpecRoutingRuleForwardingConfiguration `json:"forwardingConfiguration,omitempty" tf:"forwarding_configuration,omitempty"`
	// +kubebuilder:validation:MaxItems=100
	FrontendEndpoints []string `json:"frontendEndpoints" tf:"frontend_endpoints"`
	// +optional
	ID   string `json:"ID,omitempty" tf:"id,omitempty"`
	Name string `json:"name" tf:"name"`
	// +kubebuilder:validation:MaxItems=25
	PatternsToMatch []string `json:"patternsToMatch" tf:"patterns_to_match"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RedirectConfiguration []FrontdoorSpecRoutingRuleRedirectConfiguration `json:"redirectConfiguration,omitempty" tf:"redirect_configuration,omitempty"`
}

type FrontdoorSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:MaxItems=50
	BackendPool []FrontdoorSpecBackendPool `json:"backendPool" tf:"backend_pool"`
	// +kubebuilder:validation:MaxItems=5000
	BackendPoolHealthProbe []FrontdoorSpecBackendPoolHealthProbe `json:"backendPoolHealthProbe" tf:"backend_pool_health_probe"`
	// +kubebuilder:validation:MaxItems=5000
	BackendPoolLoadBalancing []FrontdoorSpecBackendPoolLoadBalancing `json:"backendPoolLoadBalancing" tf:"backend_pool_load_balancing"`
	// +optional
	Cname                                   string `json:"cname,omitempty" tf:"cname,omitempty"`
	EnforceBackendPoolsCertificateNameCheck bool   `json:"enforceBackendPoolsCertificateNameCheck" tf:"enforce_backend_pools_certificate_name_check"`
	// +optional
	FriendlyName string `json:"friendlyName,omitempty" tf:"friendly_name,omitempty"`
	// +kubebuilder:validation:MaxItems=100
	FrontendEndpoint []FrontdoorSpecFrontendEndpoint `json:"frontendEndpoint" tf:"frontend_endpoint"`
	// +optional
	LoadBalancerEnabled bool   `json:"loadBalancerEnabled,omitempty" tf:"load_balancer_enabled,omitempty"`
	Location            string `json:"location" tf:"location"`
	Name                string `json:"name" tf:"name"`
	ResourceGroupName   string `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=100
	RoutingRule []FrontdoorSpecRoutingRule `json:"routingRule" tf:"routing_rule"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type FrontdoorStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *FrontdoorSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
	// +optional
	TerraformErrors []string `json:"terraformErrors,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// FrontdoorList is a list of Frontdoors
type FrontdoorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Frontdoor CRD objects
	Items []Frontdoor `json:"items,omitempty"`
}
