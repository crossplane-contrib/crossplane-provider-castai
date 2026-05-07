/*
Copyright 2021 Upbound Inc.
*/

package namespaced

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	aioptimizerhostedmodel "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/aioptimizerhostedmodel"
	aioptimizermodelregistry "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/aioptimizermodelregistry"
	aioptimizermodelspecs "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/aioptimizermodelspecs"
	akscluster "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/akscluster"
	allocationgroup "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/allocationgroup"
	autoscaler "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/autoscaler"
	cacheconfiguration "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/cacheconfiguration"
	cachegroup "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/cachegroup"
	cacherule "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/cacherule"
	commitments "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/commitments"
	ekscluster "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/ekscluster"
	eksclusterid "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/eksclusterid"
	eksuserarn "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/eksuserarn"
	enterprisegroup "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/enterprisegroup"
	enterpriserolebinding "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/enterpriserolebinding"
	evictoradvancedconfig "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/evictoradvancedconfig"
	gkecluster "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/gkecluster"
	gkeclusterid "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/gkeclusterid"
	hibernationschedule "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/hibernationschedule"
	nodeconfiguration "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/nodeconfiguration"
	nodeconfigurationdefault "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/nodeconfigurationdefault"
	nodetemplate "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/nodetemplate"
	organizationgroup "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/organizationgroup"
	organizationmembers "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/organizationmembers"
	podmutation "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/podmutation"
	rebalancingjob "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/rebalancingjob"
	rebalancingschedule "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/rebalancingschedule"
	reservations "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/reservations"
	rolebindings "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/rolebindings"
	securityruntimerule "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/securityruntimerule"
	serviceaccount "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/serviceaccount"
	serviceaccountkey "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/serviceaccountkey"
	ssoconnection "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/ssoconnection"
	workloadcustommetricsdatasource "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/workloadcustommetricsdatasource"
	workloadscalingpolicy "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/workloadscalingpolicy"
	workloadscalingpolicyorder "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced/workloadscalingpolicyorder"
)

const (
	resourcePrefix = "castai"
	modulePath     = "github.com/crossplane-contrib/crossplane-provider-castai"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration for namespaced resources with root group castai.m.upbound.io
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("castai.m.upbound.io"),
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
