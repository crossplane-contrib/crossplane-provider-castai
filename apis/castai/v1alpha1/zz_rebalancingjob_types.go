/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RebalancingJobInitParameters struct {

	// (String) CAST AI cluster id.
	// CAST AI cluster id.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/crossplane-provider-castai/apis/castai/v1alpha1.EksClusterId
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a EksClusterId in castai to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a EksClusterId in castai to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// (Boolean) The job will only be executed if it's enabled.
	// The job will only be executed if it's enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) Rebalancing schedule of this job.
	// Rebalancing schedule of this job.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/crossplane-provider-castai/apis/castai/v1alpha1.RebalancingSchedule
	RebalancingScheduleID *string `json:"rebalancingScheduleId,omitempty" tf:"rebalancing_schedule_id,omitempty"`

	// Reference to a RebalancingSchedule in castai to populate rebalancingScheduleId.
	// +kubebuilder:validation:Optional
	RebalancingScheduleIDRef *v1.Reference `json:"rebalancingScheduleIdRef,omitempty" tf:"-"`

	// Selector for a RebalancingSchedule in castai to populate rebalancingScheduleId.
	// +kubebuilder:validation:Optional
	RebalancingScheduleIDSelector *v1.Selector `json:"rebalancingScheduleIdSelector,omitempty" tf:"-"`
}

type RebalancingJobObservation struct {

	// (String) CAST AI cluster id.
	// CAST AI cluster id.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// (Boolean) The job will only be executed if it's enabled.
	// The job will only be executed if it's enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Rebalancing schedule of this job.
	// Rebalancing schedule of this job.
	RebalancingScheduleID *string `json:"rebalancingScheduleId,omitempty" tf:"rebalancing_schedule_id,omitempty"`
}

type RebalancingJobParameters struct {

	// (String) CAST AI cluster id.
	// CAST AI cluster id.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/crossplane-provider-castai/apis/castai/v1alpha1.EksClusterId
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a EksClusterId in castai to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a EksClusterId in castai to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// (Boolean) The job will only be executed if it's enabled.
	// The job will only be executed if it's enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) Rebalancing schedule of this job.
	// Rebalancing schedule of this job.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/crossplane-provider-castai/apis/castai/v1alpha1.RebalancingSchedule
	// +kubebuilder:validation:Optional
	RebalancingScheduleID *string `json:"rebalancingScheduleId,omitempty" tf:"rebalancing_schedule_id,omitempty"`

	// Reference to a RebalancingSchedule in castai to populate rebalancingScheduleId.
	// +kubebuilder:validation:Optional
	RebalancingScheduleIDRef *v1.Reference `json:"rebalancingScheduleIdRef,omitempty" tf:"-"`

	// Selector for a RebalancingSchedule in castai to populate rebalancingScheduleId.
	// +kubebuilder:validation:Optional
	RebalancingScheduleIDSelector *v1.Selector `json:"rebalancingScheduleIdSelector,omitempty" tf:"-"`
}

// RebalancingJobSpec defines the desired state of RebalancingJob
type RebalancingJobSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RebalancingJobParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider RebalancingJobInitParameters `json:"initProvider,omitempty"`
}

// RebalancingJobStatus defines the observed state of RebalancingJob.
type RebalancingJobStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RebalancingJobObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RebalancingJob is the Schema for the RebalancingJobs API. Job assigns a rebalancing schedule to a cluster.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,castai}
type RebalancingJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RebalancingJobSpec   `json:"spec"`
	Status            RebalancingJobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RebalancingJobList contains a list of RebalancingJobs
type RebalancingJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RebalancingJob `json:"items"`
}

// Repository type metadata.
var (
	RebalancingJob_Kind             = "RebalancingJob"
	RebalancingJob_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RebalancingJob_Kind}.String()
	RebalancingJob_KindAPIVersion   = RebalancingJob_Kind + "." + CRDGroupVersion.String()
	RebalancingJob_GroupVersionKind = CRDGroupVersion.WithKind(RebalancingJob_Kind)
)

func init() {
	SchemeBuilder.Register(&RebalancingJob{}, &RebalancingJobList{})
}
