/*
Copyright 2021 Upbound Inc.
*/

package main

import (
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	xpcontroller "github.com/crossplane/crossplane-runtime/v2/pkg/controller"
	"github.com/crossplane/crossplane-runtime/v2/pkg/errors"
	"github.com/crossplane/crossplane-runtime/v2/pkg/feature"
	"github.com/crossplane/crossplane-runtime/v2/pkg/logging"
	"github.com/crossplane/crossplane-runtime/v2/pkg/ratelimiter"
	tjcontroller "github.com/crossplane/upjet/v2/pkg/controller"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	"gopkg.in/alecthomas/kingpin.v2"
	authorizationv1 "k8s.io/api/authorization/v1"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	"github.com/crossplane-contrib/crossplane-provider-castai/apis"
	configCluster "github.com/crossplane-contrib/crossplane-provider-castai/config/cluster"
	configNamespaced "github.com/crossplane-contrib/crossplane-provider-castai/config/namespaced"
	"github.com/crossplane-contrib/crossplane-provider-castai/internal/clients"
	controllerCluster "github.com/crossplane-contrib/crossplane-provider-castai/internal/controller"
	"github.com/crossplane-contrib/crossplane-provider-castai/internal/controller/cluster/providerconfig"
)

func main() {
	var (
		app              = kingpin.New(filepath.Base(os.Args[0]), "Terraform based Crossplane provider for CastAI").DefaultEnvars()
		debug            = app.Flag("debug", "Run with debug logging.").Short('d').Bool()
		syncInterval     = app.Flag("sync", "Sync interval controls how often all resources will be double checked for drift.").Short('s').Default("1h").Duration()
		pollInterval     = app.Flag("poll", "Poll interval controls how often an individual resource should be checked for drift.").Default("10m").Duration()
		leaderElection   = app.Flag("leader-election", "Use leader election for the controller manager.").Short('l').Default("false").OverrideDefaultFromEnvar("LEADER_ELECTION").Bool()
		terraformVersion = app.Flag("terraform-version", "Terraform version.").Required().Envar("TERRAFORM_VERSION").String()
		providerSource   = app.Flag("terraform-provider-source", "Terraform provider source.").Required().Envar("TERRAFORM_PROVIDER_SOURCE").String()
		providerVersion  = app.Flag("terraform-provider-version", "Terraform provider version.").Required().Envar("TERRAFORM_PROVIDER_VERSION").String()
		maxReconcileRate = app.Flag("max-reconcile-rate", "The global maximum rate per second at which resources may checked for drift from the desired state.").Default("10").Int()
	)

	kingpin.MustParse(app.Parse(os.Args[1:]))

	log.Default().SetOutput(io.Discard)
	ctrl.SetLogger(zap.New(zap.WriteTo(io.Discard)))

	zl := zap.New(zap.UseDevMode(*debug))
	logger := logging.NewLogrLogger(zl.WithName("crossplane-provider-castai"))

	if *debug {
		ctrl.SetLogger(zl)
	}

	pollJitter := time.Duration(float64(*pollInterval) * 0.05)
	logger.Debug("Starting", "sync-interval", syncInterval.String(),
		"poll-interval", pollInterval.String(), "poll-jitter", pollJitter, "max-reconcile-rate", *maxReconcileRate)

	cfg, err := ctrl.GetConfig()
	kingpin.FatalIfError(err, "Cannot get API server rest config")

	mgr, err := ctrl.NewManager(cfg, ctrl.Options{
		LeaderElection:   *leaderElection,
		LeaderElectionID: "crossplane-leader-election-crossplane-provider-castai",
		Cache: cache.Options{
			SyncPeriod: syncInterval,
		},
		LeaderElectionResourceLock: resourcelock.LeasesResourceLock,
		LeaseDuration:              func() *time.Duration { d := 60 * time.Second; return &d }(),
		RenewDeadline:              func() *time.Duration { d := 50 * time.Second; return &d }(),
	})
	kingpin.FatalIfError(err, "Cannot create controller manager")
	kingpin.FatalIfError(apis.AddToScheme(mgr.GetScheme()), "Cannot add CastAI APIs to scheme")

	// SafeStart: Check if we can watch CRDs
	ctx := context.Background()
	featureFlags := &feature.Flags{}
	
	canWatch, err := canWatchCRD(ctx, mgr)
	if err != nil {
		logger.Info("Unable to verify CRD watch permissions, assuming SafeStart is not needed", "error", err)
		canWatch = true // Assume we can watch if RBAC check fails
	}
	
	if !canWatch {
		logger.Info("CRD watch permissions not available, using SafeStart gating")
		featureFlags.Enable(feature.FlagSafeStart)
	}

	// Setup cluster-scoped resources
	oCluster := tjcontroller.Options{
		Options: xpcontroller.Options{
			Logger:                  logger,
			GlobalRateLimiter:       ratelimiter.NewGlobal(*maxReconcileRate),
			PollInterval:            *pollInterval,
			MaxConcurrentReconciles: 1,
			Features:                featureFlags,
		},
		Provider:       configCluster.GetProvider(),
		SetupFn:        clients.TerraformSetupBuilder(*terraformVersion, *providerSource, *providerVersion),
		WorkspaceStore: terraform.NewWorkspaceStore(logger, terraform.WithFeatures(featureFlags)),
		PollJitter:     pollJitter,
	}

	if featureFlags.Enabled(feature.FlagSafeStart) {
		kingpin.FatalIfError(providerconfig.SetupGated(mgr, oCluster), "Cannot setup gated cluster CastAI controllers")
	} else {
		kingpin.FatalIfError(controllerCluster.Setup(mgr, oCluster), "Cannot setup CastAI controllers")
	}

	// Setup namespaced-scoped resources
	oNamespaced := tjcontroller.Options{
		Options: xpcontroller.Options{
			Logger:                  logger,
			GlobalRateLimiter:       ratelimiter.NewGlobal(*maxReconcileRate),
			PollInterval:            *pollInterval,
			MaxConcurrentReconciles: 1,
			Features:                featureFlags,
		},
		Provider:       configNamespaced.GetProvider(),
		SetupFn:        clients.TerraformSetupBuilder(*terraformVersion, *providerSource, *providerVersion),
		WorkspaceStore: terraform.NewWorkspaceStore(logger, terraform.WithFeatures(featureFlags)),
		PollJitter:     pollJitter,
	}
	
	if featureFlags.Enabled(feature.FlagSafeStart) {
		kingpin.FatalIfError(providerconfig.SetupGated(mgr, oNamespaced), "Cannot setup gated namespaced CastAI controllers")
	}

	kingpin.FatalIfError(mgr.Start(ctrl.SetupSignalHandler()), "Cannot start controller manager")
}

// canWatchCRD checks if the provider has permission to watch CustomResourceDefinitions.
func canWatchCRD(ctx context.Context, mgr manager.Manager) (bool, error) {
	if err := authorizationv1.AddToScheme(mgr.GetScheme()); err != nil {
		return false, err
	}
	verbs := []string{"get", "list", "watch"}
	for _, verb := range verbs {
		sar := &authorizationv1.SelfSubjectAccessReview{
			Spec: authorizationv1.SelfSubjectAccessReviewSpec{
				ResourceAttributes: &authorizationv1.ResourceAttributes{
					Group:    "apiextensions.k8s.io",
					Resource: "customresourcedefinitions",
					Verb:     verb,
				},
			},
		}
		if err := mgr.GetClient().Create(ctx, sar); err != nil {
			return false, errors.Wrapf(err, "unable to perform RBAC check for verb %s on CustomResourceDefinitions", verb)
		}
		if !sar.Status.Allowed {
			return false, nil
		}
	}
	return true, nil
}
