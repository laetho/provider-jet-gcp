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

type FirewallPolicyRuleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	RuleTupleCount *float64 `json:"ruleTupleCount,omitempty" tf:"rule_tuple_count,omitempty"`
}

type FirewallPolicyRuleParameters struct {

	// The Action to perform when the client connection triggers the rule. Can currently be either "allow" or "deny()" where valid values for status are 403, 404, and 502.
	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// An optional description for this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The direction in which this rule applies. Possible values: INGRESS, EGRESS
	// +kubebuilder:validation:Required
	Direction *string `json:"direction" tf:"direction,omitempty"`

	// Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "goto_next" rules.
	// +kubebuilder:validation:Optional
	EnableLogging *bool `json:"enableLogging,omitempty" tf:"enable_logging,omitempty"`

	// The firewall policy of the resource.
	// +kubebuilder:validation:Required
	FirewallPolicy *string `json:"firewallPolicy" tf:"firewall_policy,omitempty"`

	// A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
	// +kubebuilder:validation:Required
	Match []MatchParameters `json:"match" tf:"match,omitempty"`

	// An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
	// +kubebuilder:validation:Required
	Priority *float64 `json:"priority" tf:"priority,omitempty"`

	// A list of network resource URLs to which this rule applies. This field allows you to control which network's VMs get this rule. If this field is left blank, all VMs within the organization will receive the rule.
	// +kubebuilder:validation:Optional
	TargetResources []*string `json:"targetResources,omitempty" tf:"target_resources,omitempty"`

	// A list of service accounts indicating the sets of instances that are applied with this rule.
	// +kubebuilder:validation:Optional
	TargetServiceAccounts []*string `json:"targetServiceAccounts,omitempty" tf:"target_service_accounts,omitempty"`
}

type Layer4ConfigsObservation struct {
}

type Layer4ConfigsParameters struct {

	// The IP protocol to which this rule applies. The protocol type is required when creating a firewall rule. This value can either be one of the following well known protocol strings (`tcp`, `udp`, `icmp`, `esp`, `ah`, `ipip`, `sctp`), or the IP protocol number.
	// +kubebuilder:validation:Required
	IPProtocol *string `json:"ipProtocol" tf:"ip_protocol,omitempty"`

	// An optional list of ports to which this rule applies. This field is only applicable for UDP or TCP protocol. Each entry must be either an integer or a range. If not specified, this rule applies to connections through any port. Example inputs include: ``.
	// +kubebuilder:validation:Optional
	Ports []*string `json:"ports,omitempty" tf:"ports,omitempty"`
}

type MatchObservation struct {
}

type MatchParameters struct {

	// CIDR IP address range. Maximum number of destination CIDR IP ranges allowed is 256.
	// +kubebuilder:validation:Optional
	DestIPRanges []*string `json:"destIpRanges,omitempty" tf:"dest_ip_ranges,omitempty"`

	// Pairs of IP protocols and ports that the rule should match.
	// +kubebuilder:validation:Required
	Layer4Configs []Layer4ConfigsParameters `json:"layer4Configs" tf:"layer4_configs,omitempty"`

	// CIDR IP address range. Maximum number of source CIDR IP ranges allowed is 256.
	// +kubebuilder:validation:Optional
	SrcIPRanges []*string `json:"srcIpRanges,omitempty" tf:"src_ip_ranges,omitempty"`
}

// FirewallPolicyRuleSpec defines the desired state of FirewallPolicyRule
type FirewallPolicyRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallPolicyRuleParameters `json:"forProvider"`
}

// FirewallPolicyRuleStatus defines the observed state of FirewallPolicyRule.
type FirewallPolicyRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallPolicyRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallPolicyRule is the Schema for the FirewallPolicyRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type FirewallPolicyRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallPolicyRuleSpec   `json:"spec"`
	Status            FirewallPolicyRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallPolicyRuleList contains a list of FirewallPolicyRules
type FirewallPolicyRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirewallPolicyRule `json:"items"`
}

// Repository type metadata.
var (
	FirewallPolicyRule_Kind             = "FirewallPolicyRule"
	FirewallPolicyRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FirewallPolicyRule_Kind}.String()
	FirewallPolicyRule_KindAPIVersion   = FirewallPolicyRule_Kind + "." + CRDGroupVersion.String()
	FirewallPolicyRule_GroupVersionKind = CRDGroupVersion.WithKind(FirewallPolicyRule_Kind)
)

func init() {
	SchemeBuilder.Register(&FirewallPolicyRule{}, &FirewallPolicyRuleList{})
}
