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

package v1alpha2

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha2"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type WebApplicationFirewallPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WebApplicationFirewallPolicySpec   `json:"spec,omitempty"`
	Status            WebApplicationFirewallPolicyStatus `json:"status,omitempty"`
}

type WebApplicationFirewallPolicySpecCustomRulesMatchConditionsMatchVariables struct {
	// +optional
	Selector     string `json:"selector,omitempty" tf:"selector,omitempty"`
	VariableName string `json:"variableName" tf:"variable_name"`
}

type WebApplicationFirewallPolicySpecCustomRulesMatchConditions struct {
	MatchValues    []string                                                                   `json:"matchValues" tf:"match_values"`
	MatchVariables []WebApplicationFirewallPolicySpecCustomRulesMatchConditionsMatchVariables `json:"matchVariables" tf:"match_variables"`
	// +optional
	NegationCondition bool   `json:"negationCondition,omitempty" tf:"negation_condition,omitempty"`
	Operator          string `json:"operator" tf:"operator"`
}

type WebApplicationFirewallPolicySpecCustomRules struct {
	Action          string                                                       `json:"action" tf:"action"`
	MatchConditions []WebApplicationFirewallPolicySpecCustomRulesMatchConditions `json:"matchConditions" tf:"match_conditions"`
	// +optional
	Name     string `json:"name,omitempty" tf:"name,omitempty"`
	Priority int64  `json:"priority" tf:"priority"`
	RuleType string `json:"ruleType" tf:"rule_type"`
}

type WebApplicationFirewallPolicySpecManagedRulesExclusion struct {
	MatchVariable         string `json:"matchVariable" tf:"match_variable"`
	Selector              string `json:"selector" tf:"selector"`
	SelectorMatchOperator string `json:"selectorMatchOperator" tf:"selector_match_operator"`
}

type WebApplicationFirewallPolicySpecManagedRulesManagedRuleSetRuleGroupOverride struct {
	DisabledRules []string `json:"disabledRules" tf:"disabled_rules"`
	RuleGroupName string   `json:"ruleGroupName" tf:"rule_group_name"`
}

type WebApplicationFirewallPolicySpecManagedRulesManagedRuleSet struct {
	// +optional
	RuleGroupOverride []WebApplicationFirewallPolicySpecManagedRulesManagedRuleSetRuleGroupOverride `json:"ruleGroupOverride,omitempty" tf:"rule_group_override,omitempty"`
	// +optional
	Type    string `json:"type,omitempty" tf:"type,omitempty"`
	Version string `json:"version" tf:"version"`
}

type WebApplicationFirewallPolicySpecManagedRules struct {
	// +optional
	Exclusion      []WebApplicationFirewallPolicySpecManagedRulesExclusion      `json:"exclusion,omitempty" tf:"exclusion,omitempty"`
	ManagedRuleSet []WebApplicationFirewallPolicySpecManagedRulesManagedRuleSet `json:"managedRuleSet" tf:"managed_rule_set"`
}

type WebApplicationFirewallPolicySpecPolicySettings struct {
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	Mode string `json:"mode,omitempty" tf:"mode,omitempty"`
}

type WebApplicationFirewallPolicySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CustomRules []WebApplicationFirewallPolicySpecCustomRules `json:"customRules,omitempty" tf:"custom_rules,omitempty"`
	Location    string                                        `json:"location" tf:"location"`
	// +kubebuilder:validation:MaxItems=1
	ManagedRules []WebApplicationFirewallPolicySpecManagedRules `json:"managedRules" tf:"managed_rules"`
	Name         string                                         `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	PolicySettings    []WebApplicationFirewallPolicySpecPolicySettings `json:"policySettings,omitempty" tf:"policy_settings,omitempty"`
	ResourceGroupName string                                           `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type WebApplicationFirewallPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *WebApplicationFirewallPolicySpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WebApplicationFirewallPolicyList is a list of WebApplicationFirewallPolicys
type WebApplicationFirewallPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WebApplicationFirewallPolicy CRD objects
	Items []WebApplicationFirewallPolicy `json:"items,omitempty"`
}
