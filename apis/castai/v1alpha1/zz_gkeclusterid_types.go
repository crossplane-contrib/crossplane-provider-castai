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

type GkeClusterIdInitParameters struct {

	// Service account email in cast project
	CastServiceAccount *string `json:"castServiceAccount,omitempty" tf:"cast_service_account,omitempty"`

	// Service account email in client project
	ClientServiceAccount *string `json:"clientServiceAccount,omitempty" tf:"client_service_account,omitempty"`

	// GCP cluster zone in case of zonal or region in case of regional cluster
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// GKE cluster name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// GCP project id
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

type GkeClusterIdObservation struct {

	// Service account email in cast project
	CastServiceAccount *string `json:"castServiceAccount,omitempty" tf:"cast_service_account,omitempty"`

	// Service account email in client project
	ClientServiceAccount *string `json:"clientServiceAccount,omitempty" tf:"client_service_account,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// GCP cluster zone in case of zonal or region in case of regional cluster
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// GKE cluster name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// GCP project id
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

type GkeClusterIdParameters struct {

	// Service account email in cast project
	// +kubebuilder:validation:Optional
	CastServiceAccount *string `json:"castServiceAccount,omitempty" tf:"cast_service_account,omitempty"`

	// Service account email in client project
	// +kubebuilder:validation:Optional
	ClientServiceAccount *string `json:"clientServiceAccount,omitempty" tf:"client_service_account,omitempty"`

	// GCP cluster zone in case of zonal or region in case of regional cluster
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// GKE cluster name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// GCP project id
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

// GkeClusterIdSpec defines the desired state of GkeClusterId
type GkeClusterIdSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GkeClusterIdParameters `json:"forProvider"`
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
	InitProvider GkeClusterIdInitParameters `json:"initProvider,omitempty"`
}

// GkeClusterIdStatus defines the observed state of GkeClusterId.
type GkeClusterIdStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GkeClusterIdObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// GkeClusterId is the Schema for the GkeClusterIds API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,castai}
type GkeClusterId struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.projectId) || (has(self.initProvider) && has(self.initProvider.projectId))",message="spec.forProvider.projectId is a required parameter"
	Spec   GkeClusterIdSpec   `json:"spec"`
	Status GkeClusterIdStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GkeClusterIdList contains a list of GkeClusterIds
type GkeClusterIdList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GkeClusterId `json:"items"`
}

// Repository type metadata.
var (
	GkeClusterId_Kind             = "GkeClusterId"
	GkeClusterId_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GkeClusterId_Kind}.String()
	GkeClusterId_KindAPIVersion   = GkeClusterId_Kind + "." + CRDGroupVersion.String()
	GkeClusterId_GroupVersionKind = CRDGroupVersion.WithKind(GkeClusterId_Kind)
)

func init() {
	SchemeBuilder.Register(&GkeClusterId{}, &GkeClusterIdList{})
}
