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

type FollowupIntentInfoObservation struct {
	FollowupIntentName *string `json:"followupIntentName,omitempty" tf:"followup_intent_name,omitempty"`

	ParentFollowupIntentName *string `json:"parentFollowupIntentName,omitempty" tf:"parent_followup_intent_name,omitempty"`
}

type FollowupIntentInfoParameters struct {
}

type IntentObservation struct {
	FollowupIntentInfo []FollowupIntentInfoObservation `json:"followupIntentInfo,omitempty" tf:"followup_intent_info,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	RootFollowupIntentName *string `json:"rootFollowupIntentName,omitempty" tf:"root_followup_intent_name,omitempty"`
}

type IntentParameters struct {

	// The name of the action associated with the intent.
	// Note: The action name must not contain whitespaces.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The list of platforms for which the first responses will be copied from the messages in PLATFORM_UNSPECIFIED
	// (i.e. default platform). Possible values: ["FACEBOOK", "SLACK", "TELEGRAM", "KIK", "SKYPE", "LINE", "VIBER", "ACTIONS_ON_GOOGLE", "GOOGLE_HANGOUTS"]
	// +kubebuilder:validation:Optional
	DefaultResponsePlatforms []*string `json:"defaultResponsePlatforms,omitempty" tf:"default_response_platforms,omitempty"`

	// The name of this intent to be displayed on the console.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// The collection of event names that trigger the intent. If the collection of input contexts is not empty, all of
	// the contexts must be present in the active user session for an event to trigger this intent. See the
	// [events reference](https://cloud.google.com/dialogflow/docs/events-overview) for more details.
	// +kubebuilder:validation:Optional
	Events []*string `json:"events,omitempty" tf:"events,omitempty"`

	// The list of context names required for this intent to be triggered.
	// Format: projects/<Project ID>/agent/sessions/-/contexts/<Context ID>.
	// +kubebuilder:validation:Optional
	InputContextNames []*string `json:"inputContextNames,omitempty" tf:"input_context_names,omitempty"`

	// Indicates whether this is a fallback intent.
	// +kubebuilder:validation:Optional
	IsFallback *bool `json:"isFallback,omitempty" tf:"is_fallback,omitempty"`

	// Indicates whether Machine Learning is disabled for the intent.
	// Note: If mlDisabled setting is set to true, then this intent is not taken into account during inference in ML
	// ONLY match mode. Also, auto-markup in the UI is turned off.
	// +kubebuilder:validation:Optional
	MLDisabled *bool `json:"mlDisabled,omitempty" tf:"ml_disabled,omitempty"`

	// The unique identifier of the parent intent in the chain of followup intents.
	// Format: projects/<Project ID>/agent/intents/<Intent ID>.
	// +kubebuilder:validation:Optional
	ParentFollowupIntentName *string `json:"parentFollowupIntentName,omitempty" tf:"parent_followup_intent_name,omitempty"`

	// The priority of this intent. Higher numbers represent higher priorities.
	// - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds
	// to the Normal priority in the console.
	// - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Indicates whether to delete all contexts in the current session when this intent is matched.
	// +kubebuilder:validation:Optional
	ResetContexts *bool `json:"resetContexts,omitempty" tf:"reset_contexts,omitempty"`

	// Indicates whether webhooks are enabled for the intent.
	// * WEBHOOK_STATE_ENABLED: Webhook is enabled in the agent and in the intent.
	// * WEBHOOK_STATE_ENABLED_FOR_SLOT_FILLING: Webhook is enabled in the agent and in the intent. Also, each slot
	// filling prompt is forwarded to the webhook. Possible values: ["WEBHOOK_STATE_ENABLED", "WEBHOOK_STATE_ENABLED_FOR_SLOT_FILLING"]
	// +kubebuilder:validation:Optional
	WebhookState *string `json:"webhookState,omitempty" tf:"webhook_state,omitempty"`
}

// IntentSpec defines the desired state of Intent
type IntentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IntentParameters `json:"forProvider"`
}

// IntentStatus defines the observed state of Intent.
type IntentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IntentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Intent is the Schema for the Intents API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type Intent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IntentSpec   `json:"spec"`
	Status            IntentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IntentList contains a list of Intents
type IntentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Intent `json:"items"`
}

// Repository type metadata.
var (
	Intent_Kind             = "Intent"
	Intent_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Intent_Kind}.String()
	Intent_KindAPIVersion   = Intent_Kind + "." + CRDGroupVersion.String()
	Intent_GroupVersionKind = CRDGroupVersion.WithKind(Intent_Kind)
)

func init() {
	SchemeBuilder.Register(&Intent{}, &IntentList{})
}
