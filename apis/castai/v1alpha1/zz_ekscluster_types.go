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

type EksClusterInitParameters struct {

	// (String) ID of AWS account
	// ID of AWS account
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (String) AWS IAM role ARN that will be assumed by CAST AI user. This role should allow sts:AssumeRole action for CAST AI user that can be retrieved using castai_eks_user_arn data source
	// AWS IAM role ARN that will be assumed by CAST AI user. This role should allow `sts:AssumeRole` action for CAST AI user that can be retrieved using `castai_eks_user_arn` data source
	AssumeRoleArn *string `json:"assumeRoleArn,omitempty" tf:"assume_role_arn,omitempty"`

	// (Boolean) Should CAST AI remove nodes managed by CAST AI on disconnect
	// Should CAST AI remove nodes managed by CAST AI on disconnect
	DeleteNodesOnDisconnect *bool `json:"deleteNodesOnDisconnect,omitempty" tf:"delete_nodes_on_disconnect,omitempty"`

	// (String) name of your EKS cluster
	// name of your EKS cluster
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) AWS region where the cluster is placed
	// AWS region where the cluster is placed
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type EksClusterObservation struct {

	// (String) ID of AWS account
	// ID of AWS account
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (String) AWS IAM role ARN that will be assumed by CAST AI user. This role should allow sts:AssumeRole action for CAST AI user that can be retrieved using castai_eks_user_arn data source
	// AWS IAM role ARN that will be assumed by CAST AI user. This role should allow `sts:AssumeRole` action for CAST AI user that can be retrieved using `castai_eks_user_arn` data source
	AssumeRoleArn *string `json:"assumeRoleArn,omitempty" tf:"assume_role_arn,omitempty"`

	// (String) CAST AI internal credentials ID
	// CAST AI internal credentials ID
	CredentialsID *string `json:"credentialsId,omitempty" tf:"credentials_id,omitempty"`

	// (Boolean) Should CAST AI remove nodes managed by CAST AI on disconnect
	// Should CAST AI remove nodes managed by CAST AI on disconnect
	DeleteNodesOnDisconnect *bool `json:"deleteNodesOnDisconnect,omitempty" tf:"delete_nodes_on_disconnect,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) name of your EKS cluster
	// name of your EKS cluster
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) AWS region where the cluster is placed
	// AWS region where the cluster is placed
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type EksClusterParameters struct {

	// (String) ID of AWS account
	// ID of AWS account
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (String) AWS IAM role ARN that will be assumed by CAST AI user. This role should allow sts:AssumeRole action for CAST AI user that can be retrieved using castai_eks_user_arn data source
	// AWS IAM role ARN that will be assumed by CAST AI user. This role should allow `sts:AssumeRole` action for CAST AI user that can be retrieved using `castai_eks_user_arn` data source
	// +kubebuilder:validation:Optional
	AssumeRoleArn *string `json:"assumeRoleArn,omitempty" tf:"assume_role_arn,omitempty"`

	// (Boolean) Should CAST AI remove nodes managed by CAST AI on disconnect
	// Should CAST AI remove nodes managed by CAST AI on disconnect
	// +kubebuilder:validation:Optional
	DeleteNodesOnDisconnect *bool `json:"deleteNodesOnDisconnect,omitempty" tf:"delete_nodes_on_disconnect,omitempty"`

	// (String) name of your EKS cluster
	// name of your EKS cluster
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) AWS region where the cluster is placed
	// AWS region where the cluster is placed
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// EksClusterSpec defines the desired state of EksCluster
type EksClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EksClusterParameters `json:"forProvider"`
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
	InitProvider EksClusterInitParameters `json:"initProvider,omitempty"`
}

// EksClusterStatus defines the observed state of EksCluster.
type EksClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EksClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// EksCluster is the Schema for the EksClusters API. EKS cluster resource allows connecting an existing EKS cluster to CAST AI.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,castai}
type EksCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.accountId) || (has(self.initProvider) && has(self.initProvider.accountId))",message="spec.forProvider.accountId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.region) || (has(self.initProvider) && has(self.initProvider.region))",message="spec.forProvider.region is a required parameter"
	Spec   EksClusterSpec   `json:"spec"`
	Status EksClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EksClusterList contains a list of EksClusters
type EksClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EksCluster `json:"items"`
}

// Repository type metadata.
var (
	EksCluster_Kind             = "EksCluster"
	EksCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EksCluster_Kind}.String()
	EksCluster_KindAPIVersion   = EksCluster_Kind + "." + CRDGroupVersion.String()
	EksCluster_GroupVersionKind = CRDGroupVersion.WithKind(EksCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&EksCluster{}, &EksClusterList{})
}
