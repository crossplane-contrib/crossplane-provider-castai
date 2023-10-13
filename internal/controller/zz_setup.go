/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	autoscaler "github.com/castai/crossplane-provider-castai/internal/controller/castai/autoscaler"
	ekscluster "github.com/castai/crossplane-provider-castai/internal/controller/castai/ekscluster"
	eksclusterid "github.com/castai/crossplane-provider-castai/internal/controller/castai/eksclusterid"
	nodeconfiguration "github.com/castai/crossplane-provider-castai/internal/controller/castai/nodeconfiguration"
	nodeconfigurationdefault "github.com/castai/crossplane-provider-castai/internal/controller/castai/nodeconfigurationdefault"
	nodetemplate "github.com/castai/crossplane-provider-castai/internal/controller/castai/nodetemplate"
	rebalancingjob "github.com/castai/crossplane-provider-castai/internal/controller/castai/rebalancingjob"
	rebalancingschedule "github.com/castai/crossplane-provider-castai/internal/controller/castai/rebalancingschedule"
	providerconfig "github.com/castai/crossplane-provider-castai/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		autoscaler.Setup,
		ekscluster.Setup,
		eksclusterid.Setup,
		nodeconfiguration.Setup,
		nodeconfigurationdefault.Setup,
		nodetemplate.Setup,
		rebalancingjob.Setup,
		rebalancingschedule.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
