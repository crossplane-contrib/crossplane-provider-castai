/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	aioptimizerhostedmodel "github.com/crossplane-contrib/crossplane-provider-castai/config/aioptimizerhostedmodel"
	aioptimizermodelregistry "github.com/crossplane-contrib/crossplane-provider-castai/config/aioptimizermodelregistry"
	aioptimizermodelspecs "github.com/crossplane-contrib/crossplane-provider-castai/config/aioptimizermodelspecs"
	akscluster "github.com/crossplane-contrib/crossplane-provider-castai/config/akscluster"
	allocationgroup "github.com/crossplane-contrib/crossplane-provider-castai/config/allocationgroup"
	autoscaler "github.com/crossplane-contrib/crossplane-provider-castai/config/autoscaler"
	cacheconfiguration "github.com/crossplane-contrib/crossplane-provider-castai/config/cacheconfiguration"
	cachegroup "github.com/crossplane-contrib/crossplane-provider-castai/config/cachegroup"
	cacherule "github.com/crossplane-contrib/crossplane-provider-castai/config/cacherule"
	commitments "github.com/crossplane-contrib/crossplane-provider-castai/config/commitments"
	ekscluster "github.com/crossplane-contrib/crossplane-provider-castai/config/ekscluster"
	eksclusterid "github.com/crossplane-contrib/crossplane-provider-castai/config/eksclusterid"
	eksuserarn "github.com/crossplane-contrib/crossplane-provider-castai/config/eksuserarn"
	enterprisegroup "github.com/crossplane-contrib/crossplane-provider-castai/config/enterprisegroup"
	enterpriserolebinding "github.com/crossplane-contrib/crossplane-provider-castai/config/enterpriserolebinding"
	evictoradvancedconfig "github.com/crossplane-contrib/crossplane-provider-castai/config/evictoradvancedconfig"
	gkecluster "github.com/crossplane-contrib/crossplane-provider-castai/config/gkecluster"
	gkeclusterid "github.com/crossplane-contrib/crossplane-provider-castai/config/gkeclusterid"
	hibernationschedule "github.com/crossplane-contrib/crossplane-provider-castai/config/hibernationschedule"
	nodeconfiguration "github.com/crossplane-contrib/crossplane-provider-castai/config/nodeconfiguration"
	nodeconfigurationdefault "github.com/crossplane-contrib/crossplane-provider-castai/config/nodeconfigurationdefault"
	nodetemplate "github.com/crossplane-contrib/crossplane-provider-castai/config/nodetemplate"
	organizationgroup "github.com/crossplane-contrib/crossplane-provider-castai/config/organizationgroup"
	organizationmembers "github.com/crossplane-contrib/crossplane-provider-castai/config/organizationmembers"
	podmutation "github.com/crossplane-contrib/crossplane-provider-castai/config/podmutation"
	rebalancingjob "github.com/crossplane-contrib/crossplane-provider-castai/config/rebalancingjob"
	rebalancingschedule "github.com/crossplane-contrib/crossplane-provider-castai/config/rebalancingschedule"
	reservations "github.com/crossplane-contrib/crossplane-provider-castai/config/reservations"
	rolebindings "github.com/crossplane-contrib/crossplane-provider-castai/config/rolebindings"
	securityruntimerule "github.com/crossplane-contrib/crossplane-provider-castai/config/securityruntimerule"
	serviceaccount "github.com/crossplane-contrib/crossplane-provider-castai/config/serviceaccount"
	serviceaccountkey "github.com/crossplane-contrib/crossplane-provider-castai/config/serviceaccountkey"
	ssoconnation "github.com/crossplane-contrib/crossplane-provider-castai/config/ssoconnection"
	workloadcustommetricsdatasource "github.com/crossplane-contrib/crossplane-provider-castai/config/workloadcustommetricsdatasource"
	workloadscalingpolicy "github.com/crossplane-contrib/crossplane-provider-castai/config/workloadscalingpolicy"
	workloadscalingpolicyorder "github.com/crossplane-contrib/crossplane-provider-castai/config/workloadscalingpolicyorder"
)

const (
	resourcePrefix = "castai"
	modulePath     = "github.com/crossplane-contrib/crossplane-provider-castai"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
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
		ssoconnation.Configure,
		workloadcustommetricsdatasource.Configure,
		workloadscalingpolicy.Configure,
		workloadscalingpolicyorder.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
