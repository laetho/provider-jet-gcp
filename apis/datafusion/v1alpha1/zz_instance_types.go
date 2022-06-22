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

type InstanceObservation struct {
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	GcsBucket *string `json:"gcsBucket,omitempty" tf:"gcs_bucket,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	ServiceEndpoint *string `json:"serviceEndpoint,omitempty" tf:"service_endpoint,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`

	StateMessage *string `json:"stateMessage,omitempty" tf:"state_message,omitempty"`

	TenantProjectID *string `json:"tenantProjectId,omitempty" tf:"tenant_project_id,omitempty"`

	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type InstanceParameters struct {

	// User-managed service account to set on Dataproc when Cloud Data Fusion creates Dataproc to run data processing pipelines.
	// +kubebuilder:validation:Optional
	DataprocServiceAccount *string `json:"dataprocServiceAccount,omitempty" tf:"dataproc_service_account,omitempty"`

	// An optional description of the instance.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Option to enable Stackdriver Logging.
	// +kubebuilder:validation:Optional
	EnableStackdriverLogging *bool `json:"enableStackdriverLogging,omitempty" tf:"enable_stackdriver_logging,omitempty"`

	// Option to enable Stackdriver Monitoring.
	// +kubebuilder:validation:Optional
	EnableStackdriverMonitoring *bool `json:"enableStackdriverMonitoring,omitempty" tf:"enable_stackdriver_monitoring,omitempty"`

	// The resource labels for instance to use to annotate any related underlying resources,
	// such as Compute Engine VMs.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The ID of the instance or a fully qualified identifier for the instance.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Network configuration options. These are required when a private Data Fusion instance is to be created.
	// +kubebuilder:validation:Optional
	NetworkConfig []NetworkConfigParameters `json:"networkConfig,omitempty" tf:"network_config,omitempty"`

	// Map of additional options used to configure the behavior of Data Fusion instance.
	// +kubebuilder:validation:Optional
	Options map[string]*string `json:"options,omitempty" tf:"options,omitempty"`

	// Specifies whether the Data Fusion instance should be private. If set to
	// true, all Data Fusion nodes will have private IP addresses and will not be
	// able to access the public internet.
	// +kubebuilder:validation:Optional
	PrivateInstance *bool `json:"privateInstance,omitempty" tf:"private_instance,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the Data Fusion instance.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Represents the type of Data Fusion instance. Each type is configured with
	// the default settings for processing and memory.
	// - BASIC: Basic Data Fusion instance. In Basic type, the user will be able to create data pipelines
	// using point and click UI. However, there are certain limitations, such as fewer number
	// of concurrent pipelines, no support for streaming pipelines, etc.
	// - ENTERPRISE: Enterprise Data Fusion instance. In Enterprise type, the user will have more features
	// available, such as support for streaming pipelines, higher number of concurrent pipelines, etc.
	// - DEVELOPER: Developer Data Fusion instance. In Developer type, the user will have all features available but
	// with restrictive capabilities. This is to help enterprises design and develop their data ingestion and integration
	// pipelines at low cost. Possible values: ["BASIC", "ENTERPRISE", "DEVELOPER"]
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// Current version of the Data Fusion.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type NetworkConfigObservation struct {
}

type NetworkConfigParameters struct {

	// The IP range in CIDR notation to use for the managed Data Fusion instance
	// nodes. This range must not overlap with any other ranges used in the Data Fusion instance network.
	// +kubebuilder:validation:Required
	IPAllocation *string `json:"ipAllocation" tf:"ip_allocation,omitempty"`

	// Name of the network in the project with which the tenant project
	// will be peered for executing pipelines. In case of shared VPC where the network resides in another host
	// project the network should specified in the form of projects/{host-project-id}/global/networks/{network}
	// +kubebuilder:validation:Required
	Network *string `json:"network" tf:"network,omitempty"`
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Instance is the Schema for the Instances API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec"`
	Status            InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// Repository type metadata.
var (
	Instance_Kind             = "Instance"
	Instance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Instance_Kind}.String()
	Instance_KindAPIVersion   = Instance_Kind + "." + CRDGroupVersion.String()
	Instance_GroupVersionKind = CRDGroupVersion.WithKind(Instance_Kind)
)

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}
