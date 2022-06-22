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

type GarbageCollectionPolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type GarbageCollectionPolicyParameters struct {

	// The name of the column family.
	// +kubebuilder:validation:Required
	ColumnFamily *string `json:"columnFamily" tf:"column_family,omitempty"`

	// Serialized JSON string for garbage collection policy. Conflicts with "mode", "max_age" and "max_version".
	// +kubebuilder:validation:Optional
	GcRules *string `json:"gcRules,omitempty" tf:"gc_rules,omitempty"`

	// The name of the Bigtable instance.
	// +kubebuilder:validation:Required
	InstanceName *string `json:"instanceName" tf:"instance_name,omitempty"`

	// GC policy that applies to all cells older than the given age.
	// +kubebuilder:validation:Optional
	MaxAge []MaxAgeParameters `json:"maxAge,omitempty" tf:"max_age,omitempty"`

	// GC policy that applies to all versions of a cell except for the most recent.
	// +kubebuilder:validation:Optional
	MaxVersion []MaxVersionParameters `json:"maxVersion,omitempty" tf:"max_version,omitempty"`

	// If multiple policies are set, you should choose between UNION OR INTERSECTION.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The name of the table.
	// +kubebuilder:validation:Required
	Table *string `json:"table" tf:"table,omitempty"`
}

type MaxAgeObservation struct {
}

type MaxAgeParameters struct {

	// Number of days before applying GC policy.
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// Duration before applying GC policy
	// +kubebuilder:validation:Optional
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`
}

type MaxVersionObservation struct {
}

type MaxVersionParameters struct {

	// Number of version before applying the GC policy.
	// +kubebuilder:validation:Required
	Number *float64 `json:"number" tf:"number,omitempty"`
}

// GarbageCollectionPolicySpec defines the desired state of GarbageCollectionPolicy
type GarbageCollectionPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GarbageCollectionPolicyParameters `json:"forProvider"`
}

// GarbageCollectionPolicyStatus defines the observed state of GarbageCollectionPolicy.
type GarbageCollectionPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GarbageCollectionPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GarbageCollectionPolicy is the Schema for the GarbageCollectionPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type GarbageCollectionPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GarbageCollectionPolicySpec   `json:"spec"`
	Status            GarbageCollectionPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GarbageCollectionPolicyList contains a list of GarbageCollectionPolicys
type GarbageCollectionPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GarbageCollectionPolicy `json:"items"`
}

// Repository type metadata.
var (
	GarbageCollectionPolicy_Kind             = "GarbageCollectionPolicy"
	GarbageCollectionPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GarbageCollectionPolicy_Kind}.String()
	GarbageCollectionPolicy_KindAPIVersion   = GarbageCollectionPolicy_Kind + "." + CRDGroupVersion.String()
	GarbageCollectionPolicy_GroupVersionKind = CRDGroupVersion.WithKind(GarbageCollectionPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&GarbageCollectionPolicy{}, &GarbageCollectionPolicyList{})
}
