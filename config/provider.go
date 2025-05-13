/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/crossplane-contrib/crossplane-provider-castai/config/akscluster"
	autoscaler "github.com/crossplane-contrib/crossplane-provider-castai/config/autoscaler"
	ekscluster "github.com/crossplane-contrib/crossplane-provider-castai/config/ekscluster"
	eksclusterid "github.com/crossplane-contrib/crossplane-provider-castai/config/eksclusterid"
	eksuserarn "github.com/crossplane-contrib/crossplane-provider-castai/config/eksuserarn"
	"github.com/crossplane-contrib/crossplane-provider-castai/config/evictoradvancedconfig"
	"github.com/crossplane-contrib/crossplane-provider-castai/config/gkecluster"
	gkeclusterid "github.com/crossplane-contrib/crossplane-provider-castai/config/gkeclusterid"
	nodeconfiguration "github.com/crossplane-contrib/crossplane-provider-castai/config/nodeconfiguration"
	nodeconfigurationdefault "github.com/crossplane-contrib/crossplane-provider-castai/config/nodeconfigurationdefault"
	nodetemplate "github.com/crossplane-contrib/crossplane-provider-castai/config/nodetemplate"
	"github.com/crossplane-contrib/crossplane-provider-castai/config/organizationmembers"
	rebalancingjob "github.com/crossplane-contrib/crossplane-provider-castai/config/rebalancingjob"
	rebalancingschedule "github.com/crossplane-contrib/crossplane-provider-castai/config/rebalancingschedule"
	"github.com/crossplane-contrib/crossplane-provider-castai/config/reservations"
	ssoconnation "github.com/crossplane-contrib/crossplane-provider-castai/config/ssoconnection"
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
		ekscluster.Configure,
		autoscaler.Configure,
		nodeconfiguration.Configure,
		nodeconfigurationdefault.Configure,
		nodetemplate.Configure,
		rebalancingjob.Configure,
		rebalancingschedule.Configure,
		eksclusterid.Configure,
		eksuserarn.Configure,
		akscluster.Configure,
		evictoradvancedconfig.Configure,
		gkecluster.Configure,
		gkeclusterid.Configure,
		organizationmembers.Configure,
		reservations.Configure,
		ssoconnation.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
