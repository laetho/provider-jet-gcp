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

type RegionTargetHTTPProxyObservation struct {
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	ProxyID *float64 `json:"proxyId,omitempty" tf:"proxy_id,omitempty"`

	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type RegionTargetHTTPProxyParameters struct {

	// An optional description of this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The Region in which the created target https proxy should reside.
	// If it is not provided, the provider region is used.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A reference to the RegionUrlMap resource that defines the mapping from URL
	// to the BackendService.
	// +kubebuilder:validation:Required
	URLMap *string `json:"urlMap" tf:"url_map,omitempty"`
}

// RegionTargetHTTPProxySpec defines the desired state of RegionTargetHTTPProxy
type RegionTargetHTTPProxySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegionTargetHTTPProxyParameters `json:"forProvider"`
}

// RegionTargetHTTPProxyStatus defines the observed state of RegionTargetHTTPProxy.
type RegionTargetHTTPProxyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegionTargetHTTPProxyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RegionTargetHTTPProxy is the Schema for the RegionTargetHTTPProxys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type RegionTargetHTTPProxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RegionTargetHTTPProxySpec   `json:"spec"`
	Status            RegionTargetHTTPProxyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegionTargetHTTPProxyList contains a list of RegionTargetHTTPProxys
type RegionTargetHTTPProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegionTargetHTTPProxy `json:"items"`
}

// Repository type metadata.
var (
	RegionTargetHTTPProxy_Kind             = "RegionTargetHTTPProxy"
	RegionTargetHTTPProxy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RegionTargetHTTPProxy_Kind}.String()
	RegionTargetHTTPProxy_KindAPIVersion   = RegionTargetHTTPProxy_Kind + "." + CRDGroupVersion.String()
	RegionTargetHTTPProxy_GroupVersionKind = CRDGroupVersion.WithKind(RegionTargetHTTPProxy_Kind)
)

func init() {
	SchemeBuilder.Register(&RegionTargetHTTPProxy{}, &RegionTargetHTTPProxyList{})
}
