/*
Copyright 2021 Upbound Inc.
*/

package cluster

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	aioptimizerhostedmodel "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/aioptimizerhostedmodel"
	aioptimizermodelregistry "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/aioptimizermodelregistry"
	aioptimizermodelspecs "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/aioptimizermodelspecs"
	akscluster "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/akscluster"
	allocationgroup "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/allocationgroup"
	autoscaler "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/autoscaler"
	cacheconfiguration "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/cacheconfiguration"
	cachegroup "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/cachegroup"
	cacherule "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/cacherule"
	commitments "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/commitments"
	ekscluster "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/ekscluster"
	eksclusterid "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/eksclusterid"
	eksuserarn "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/eksuserarn"
	enterprisegroup "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/enterprisegroup"
	enterpriserolebinding "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/enterpriserolebinding"
	evictoradvancedconfig "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/evictoradvancedconfig"
	gkecluster "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/gkecluster"
	gkeclusterid "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/gkeclusterid"
	hibernationschedule "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/hibernationschedule"
	nodeconfiguration "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/nodeconfiguration"
	nodeconfigurationdefault "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/nodeconfigurationdefault"
	nodetemplate "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/nodetemplate"
	organizationgroup "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/organizationgroup"
	organizationmembers "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/organizationmembers"
	podmutation "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/podmutation"
	rebalancingjob "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/rebalancingjob"
	rebalancingschedule "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/rebalancingschedule"
	reservations "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/reservations"
	rolebindings "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/rolebindings"
	securityruntimerule "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/securityruntimerule"
	serviceaccount "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/serviceaccount"
	serviceaccountkey "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/serviceaccountkey"
	ssoconnection "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/ssoconnection"
	workloadcustommetricsdatasource "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/workloadcustommetricsdatasource"
	workloadscalingpolicy "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/workloadscalingpolicy"
	workloadscalingpolicyorder "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster/workloadscalingpolicyorder"
)

const (
	resourcePrefix = "castai"
	modulePath     = "github.com/crossplane-contrib/crossplane-provider-castai"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration for cluster-scoped resources
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		aioptimizerhostedmodel.Configure,
		aioptimizermodelregistry.Configure,
		aioptimizermodelspecs.Configure,
		akscluster.Configure,
		allocationgroup.Configure,
		autoscaler.Configure,
		cacheconfiguration.Configure,
		cachegroup.Configure,
		cacherule.Configure,
		commitments.Configure,
		ekscluster.Configure,
		eksclusterid.Configure,
		eksuserarn.Configure,
		enterprisegroup.Configure,
		enterpriserolebinding.Configure,
		evictoradvancedconfig.Configure,
		gkecluster.Configure,
		gkeclusterid.Configure,
		hibernationschedule.Configure,
		nodeconfiguration.Configure,
		nodeconfigurationdefault.Configure,
		nodetemplate.Configure,
		organizationgroup.Configure,
		organizationmembers.Configure,
		podmutation.Configure,
		rebalancingjob.Configure,
		rebalancingschedule.Configure,
		reservations.Configure,
		rolebindings.Configure,
		securityruntimerule.Configure,
		serviceaccount.Configure,
		serviceaccountkey.Configure,
		ssoconnection.Configure,
		workloadcustommetricsdatasource.Configure,
		workloadscalingpolicy.Configure,
		workloadscalingpolicyorder.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
