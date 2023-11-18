/*
Copyright 2022 Upbound Inc.
*/

package v1alpha1

import (
	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// AWSUserARNParameters are the parameters used for AWSUserARN
type AWSUserARNParameters struct {
	// CAST AI cluster id
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/crossplane-provider-castai/apis/castai/v1alpha1.EksClusterId
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty"`

	// Reference to a EksClusterId in castai to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty"`

	// Selector for a EksClusterId in castai to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty"`
}

// AWSUserARNObservation are the parameters used for AWSUserARN status
type AWSUserARNObservation struct {
	ARN                 *string `json:"arn,omitempty"`
	ClusterID           *string `json:"clusterId,omitempty"`
	ManagementAccountID *string `json:"managementAccountId,omitempty"`
	UserPrefix          *string `json:"userPrefix,omitempty"`
}

// AWSUserARNSpec defines the desired state of AWSUserARN
type AWSUserARNSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AWSUserARNParameters `json:"forProvider"`
}

// AWSUserARNStatus defines the observed state of AWSUserARN.
type AWSUserARNStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AWSUserARNObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AWSUserARN is used to retrieve ARN of given EKS cluster ID.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,castai}
type AWSUserARN struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AWSUserARNSpec   `json:"spec"`
	Status            AWSUserARNStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AWSUserARNList contains a list of AWSUserARNs
type AWSUserARNList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AWSUserARN `json:"items"`
}

// Repository type metadata.
var (
	AWSUserARNKind             = "AWSUserARN"
	AWSUserARNGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AWSUserARNKind}.String()
	AWSUserARNKindAPIVersion   = AWSUserARNKind + "." + CRDGroupVersion.String()
	AWSUserARNGroupVersionKind = CRDGroupVersion.WithKind(AWSUserARNKind)
)

func init() {
	SchemeBuilder.Register(&AWSUserARN{}, &AWSUserARNList{})
}
