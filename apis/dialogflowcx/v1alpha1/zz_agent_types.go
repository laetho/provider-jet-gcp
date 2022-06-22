/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AgentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	StartFlow *string `json:"startFlow,omitempty" tf:"start_flow,omitempty"`
}

type AgentParameters struct {

	// The URI of the agent's avatar. Avatars are used throughout the Dialogflow console and in the self-hosted Web Demo integration.
	// +kubebuilder:validation:Optional
	AvatarURI *string `json:"avatarUri,omitempty" tf:"avatar_uri,omitempty"`

	// The default language of the agent as a language tag. [See Language Support](https://cloud.google.com/dialogflow/cx/docs/reference/language)
	// for a list of the currently supported language codes. This field cannot be updated after creation.
	// +kubebuilder:validation:Required
	DefaultLanguageCode *string `json:"defaultLanguageCode" tf:"default_language_code,omitempty"`

	// The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The human-readable name of the agent, unique within the location.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// Indicates if automatic spell correction is enabled in detect intent requests.
	// +kubebuilder:validation:Optional
	EnableSpellCorrection *bool `json:"enableSpellCorrection,omitempty" tf:"enable_spell_correction,omitempty"`

	// Determines whether this agent should log conversation queries.
	// +kubebuilder:validation:Optional
	EnableStackdriverLogging *bool `json:"enableStackdriverLogging,omitempty" tf:"enable_stackdriver_logging,omitempty"`

	// The name of the location this agent is located in.
	//
	// ~> **Note:** The first time you are deploying an Agent in your project you must configure location settings.
	// This is a one time step but at the moment you can only [configure location settings](https://cloud.google.com/dialogflow/cx/docs/concept/region#location-settings) via the Dialogflow CX console.
	// Another options is to use global location so you don't need to manually configure location settings.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Name of the SecuritySettings reference for the agent. Format: projects/<Project ID>/locations/<Location ID>/securitySettings/<Security Settings ID>.
	// +kubebuilder:validation:Optional
	SecuritySettings *string `json:"securitySettings,omitempty" tf:"security_settings,omitempty"`

	// Settings related to speech recognition.
	// +kubebuilder:validation:Optional
	SpeechToTextSettings []SpeechToTextSettingsParameters `json:"speechToTextSettings,omitempty" tf:"speech_to_text_settings,omitempty"`

	// The list of all languages supported by this agent (except for the default_language_code).
	// +kubebuilder:validation:Optional
	SupportedLanguageCodes []*string `json:"supportedLanguageCodes,omitempty" tf:"supported_language_codes,omitempty"`

	// The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York,
	// Europe/Paris.
	// +kubebuilder:validation:Required
	TimeZone *string `json:"timeZone" tf:"time_zone,omitempty"`
}

type SpeechToTextSettingsObservation struct {
}

type SpeechToTextSettingsParameters struct {

	// Whether to use speech adaptation for speech recognition.
	// +kubebuilder:validation:Optional
	EnableSpeechAdaptation *bool `json:"enableSpeechAdaptation,omitempty" tf:"enable_speech_adaptation,omitempty"`
}

// AgentSpec defines the desired state of Agent
type AgentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AgentParameters `json:"forProvider"`
}

// AgentStatus defines the observed state of Agent.
type AgentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AgentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Agent is the Schema for the Agents API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type Agent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AgentSpec   `json:"spec"`
	Status            AgentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AgentList contains a list of Agents
type AgentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Agent `json:"items"`
}

// Repository type metadata.
var (
	Agent_Kind             = "Agent"
	Agent_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Agent_Kind}.String()
	Agent_KindAPIVersion   = Agent_Kind + "." + CRDGroupVersion.String()
	Agent_GroupVersionKind = CRDGroupVersion.WithKind(Agent_Kind)
)

func init() {
	SchemeBuilder.Register(&Agent{}, &AgentList{})
}
