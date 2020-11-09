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

type BotChannelSlack struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BotChannelSlackSpec   `json:"spec,omitempty"`
	Status            BotChannelSlackStatus `json:"status,omitempty"`
}

type BotChannelSlackSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BotName      string `json:"botName" tf:"bot_name"`
	ClientID     string `json:"clientID" tf:"client_id"`
	ClientSecret string `json:"-" sensitive:"true" tf:"client_secret"`
	// +optional
	LandingPageURL    string `json:"landingPageURL,omitempty" tf:"landing_page_url,omitempty"`
	Location          string `json:"location" tf:"location"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	VerificationToken string `json:"-" sensitive:"true" tf:"verification_token"`
}

type BotChannelSlackStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *BotChannelSlackSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BotChannelSlackList is a list of BotChannelSlacks
type BotChannelSlackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BotChannelSlack CRD objects
	Items []BotChannelSlack `json:"items,omitempty"`
}
