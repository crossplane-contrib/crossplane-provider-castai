/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	akscluster "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/akscluster"
	autoscaler "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/autoscaler"
	ekscluster "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/ekscluster"
	eksclusterid "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/eksclusterid"
	eksuserarn "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/eksuserarn"
	evictoradvancedconfig "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/evictoradvancedconfig"
	gkecluster "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/gkecluster"
	nodeconfiguration "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/nodeconfiguration"
	nodeconfigurationdefault "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/nodeconfigurationdefault"
	nodetemplate "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/nodetemplate"
	organizationmembers "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/organizationmembers"
	rebalancingjob "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/rebalancingjob"
	rebalancingschedule "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/rebalancingschedule"
	reservations "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/reservations"
	ssoconnection "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/castai/ssoconnection"
	providerconfig "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/providerconfig"
	scalingpolicy "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/workload/scalingpolicy"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		akscluster.Setup,
		autoscaler.Setup,
		ekscluster.Setup,
		eksclusterid.Setup,
		eksuserarn.Setup,
		evictoradvancedconfig.Setup,
		gkecluster.Setup,
		nodeconfiguration.Setup,
		nodeconfigurationdefault.Setup,
		nodetemplate.Setup,
		organizationmembers.Setup,
		rebalancingjob.Setup,
		rebalancingschedule.Setup,
		reservations.Setup,
		ssoconnection.Setup,
		providerconfig.Setup,
		scalingpolicy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
