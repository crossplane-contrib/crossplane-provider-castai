/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	aioptimizerhostedmodel "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/aioptimizerhostedmodel"
	aioptimizermodelregistry "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/aioptimizermodelregistry"
	aioptimizermodelspecs "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/aioptimizermodelspecs"
	akscluster "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/akscluster"
	allocationgroup "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/allocationgroup"
	autoscaler "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/autoscaler"
	cacheconfiguration "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/cacheconfiguration"
	cachegroup "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/cachegroup"
	cacherule "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/cacherule"
	commitments "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/commitments"
	ekscluster "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/ekscluster"
	eksclusterid "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/eksclusterid"
	eksuserarn "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/eksuserarn"
	enterprisegroup "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/enterprisegroup"
	enterpriserolebinding "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/enterpriserolebinding"
	evictoradvancedconfig "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/evictoradvancedconfig"
	gkecluster "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/gkecluster"
	gkeclusterid "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/gkeclusterid"
	hibernationschedule "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/hibernationschedule"
	nodeconfiguration "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/nodeconfiguration"
	nodeconfigurationdefault "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/nodeconfigurationdefault"
	nodetemplate "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/nodetemplate"
	organizationgroup "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/organizationgroup"
	organizationmembers "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/organizationmembers"
	podmutation "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/podmutation"
	rebalancingjob "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/rebalancingjob"
	rebalancingschedule "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/rebalancingschedule"
	reservations "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/reservations"
	rolebindings "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/rolebindings"
	scalingpolicy "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/scalingpolicy"
	scalingpolicyorder "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/scalingpolicyorder"
	securityruntimerule "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/securityruntimerule"
	serviceaccount "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/serviceaccount"
	serviceaccountkey "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/serviceaccountkey"
	ssoconnection "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/ssoconnection"
	workloadcustommetricsdatasource "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/castai/workloadcustommetricsdatasource"
	providerconfig "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		aioptimizerhostedmodel.Setup,
		aioptimizermodelregistry.Setup,
		aioptimizermodelspecs.Setup,
		akscluster.Setup,
		allocationgroup.Setup,
		autoscaler.Setup,
		cacheconfiguration.Setup,
		cachegroup.Setup,
		cacherule.Setup,
		commitments.Setup,
		ekscluster.Setup,
		eksclusterid.Setup,
		eksuserarn.Setup,
		enterprisegroup.Setup,
		enterpriserolebinding.Setup,
		evictoradvancedconfig.Setup,
		gkecluster.Setup,
		gkeclusterid.Setup,
		hibernationschedule.Setup,
		nodeconfiguration.Setup,
		nodeconfigurationdefault.Setup,
		nodetemplate.Setup,
		organizationgroup.Setup,
		organizationmembers.Setup,
		podmutation.Setup,
		rebalancingjob.Setup,
		rebalancingschedule.Setup,
		reservations.Setup,
		rolebindings.Setup,
		scalingpolicy.Setup,
		scalingpolicyorder.Setup,
		securityruntimerule.Setup,
		serviceaccount.Setup,
		serviceaccountkey.Setup,
		ssoconnection.Setup,
		workloadcustommetricsdatasource.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
