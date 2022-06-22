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

type NetworkEndpointGroupObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`
}

type NetworkEndpointGroupParameters struct {

	// The default port used if the port number is not specified in the
	// network endpoint.
	// +kubebuilder:validation:Optional
	DefaultPort *float64 `json:"defaultPort,omitempty" tf:"default_port,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of the resource; provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The network to which all network endpoints in the NEG belong.
	// Uses "default" project network if unspecified.
	// +kubebuilder:validation:Required
	Network *string `json:"network" tf:"network,omitempty"`

	// Type of network endpoints in this network endpoint group.
	// NON_GCP_PRIVATE_IP_PORT is used for hybrid connectivity network
	// endpoint groups (see https://cloud.google.com/load-balancing/docs/hybrid).
	// Note that NON_GCP_PRIVATE_IP_PORT can only be used with Backend Services
	// that 1) have the following load balancing schemes: EXTERNAL, EXTERNAL_MANAGED,
	// INTERNAL_MANAGED, and INTERNAL_SELF_MANAGED and 2) support the RATE or
	// CONNECTION balancing modes. Default value: "GCE_VM_IP_PORT" Possible values: ["GCE_VM_IP_PORT", "NON_GCP_PRIVATE_IP_PORT"]
	// +kubebuilder:validation:Optional
	NetworkEndpointType *string `json:"networkEndpointType,omitempty" tf:"network_endpoint_type,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Optional subnetwork to which all network endpoints in the NEG belong.
	// +kubebuilder:validation:Optional
	Subnetwork *string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`

	// Zone where the network endpoint group is located.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// NetworkEndpointGroupSpec defines the desired state of NetworkEndpointGroup
type NetworkEndpointGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkEndpointGroupParameters `json:"forProvider"`
}

// NetworkEndpointGroupStatus defines the observed state of NetworkEndpointGroup.
type NetworkEndpointGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkEndpointGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkEndpointGroup is the Schema for the NetworkEndpointGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type NetworkEndpointGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkEndpointGroupSpec   `json:"spec"`
	Status            NetworkEndpointGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkEndpointGroupList contains a list of NetworkEndpointGroups
type NetworkEndpointGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkEndpointGroup `json:"items"`
}

// Repository type metadata.
var (
	NetworkEndpointGroup_Kind             = "NetworkEndpointGroup"
	NetworkEndpointGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkEndpointGroup_Kind}.String()
	NetworkEndpointGroup_KindAPIVersion   = NetworkEndpointGroup_Kind + "." + CRDGroupVersion.String()
	NetworkEndpointGroup_GroupVersionKind = CRDGroupVersion.WithKind(NetworkEndpointGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkEndpointGroup{}, &NetworkEndpointGroupList{})
}
