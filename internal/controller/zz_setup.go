/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	akscluster "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/akscluster"
	allocationgroup "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/allocationgroup"
	autoscaler "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/autoscaler"
	commitments "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/commitments"
	ekscluster "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/ekscluster"
	eksclusterid "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/eksclusterid"
	eksuserarn "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/eksuserarn"
	evictoradvancedconfig "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/evictoradvancedconfig"
	gkecluster "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/gkecluster"
	gkeclusterid "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/gkeclusterid"
	hibernationschedule "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/hibernationschedule"
	nodeconfiguration "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/nodeconfiguration"
	nodeconfigurationdefault "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/nodeconfigurationdefault"
	nodetemplate "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/nodetemplate"
	organizationgroup "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/organizationgroup"
	organizationmembers "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/organizationmembers"
	rebalancingjob "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/rebalancingjob"
	rebalancingschedule "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/rebalancingschedule"
	reservations "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/reservations"
	rolebindings "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/rolebindings"
	scalingpolicy "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/scalingpolicy"
	scalingpolicyorder "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/scalingpolicyorder"
	securityruntimerule "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/securityruntimerule"
	serviceaccount "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/serviceaccount"
	serviceaccountkey "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/serviceaccountkey"
	ssoconnection "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/ssoconnection"
	providerconfig "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		akscluster.Setup,
		allocationgroup.Setup,
		autoscaler.Setup,
		commitments.Setup,
		ekscluster.Setup,
		eksclusterid.Setup,
		eksuserarn.Setup,
		evictoradvancedconfig.Setup,
		gkecluster.Setup,
		gkeclusterid.Setup,
		hibernationschedule.Setup,
		nodeconfiguration.Setup,
		nodeconfigurationdefault.Setup,
		nodetemplate.Setup,
		organizationgroup.Setup,
		organizationmembers.Setup,
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
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
