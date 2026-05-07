/*
Copyright 2022 Upbound Inc.
*/

package providerconfig

import (
	"context"

	"github.com/crossplane/crossplane-runtime/v2/pkg/event"
	"github.com/crossplane/crossplane-runtime/v2/pkg/meta"
	"github.com/crossplane/crossplane-runtime/v2/pkg/reconciler/providerconfig"
	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/controller"
	"github.com/pkg/errors"
	admissionregistration "k8s.io/api/admissionregistration/v1"
	authorizationv1 "k8s.io/api/authorization/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane-contrib/crossplane-provider-castai/apis/namespaced/v1beta1"
)

// SetupGated registers the controller with the manager's gated watcher.
// It ensures controllers don't start until needed CRDs are present.
func SetupGated(mgr ctrl.Manager, o controller.Options, starter func(func(), ...schema.GroupVersionKind)) error {
	zk := v1beta1.ProviderConfigGroupVersionKind.String() + "/" + v1beta1.ProviderConfigKind
	k := v1beta1.ProviderConfigGroupVersionKind
	starter(func() {
		if err := Setup(mgr, o); err != nil {
			mgr.GetLogger().Error(err, "unable to create controller", "gvk", zk)
		}
	}, k, v1beta1.ProviderConfigUsageGroupVersionKind)
	return nil
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
			return false, errors.Wrapf(err, "unable to perform RBAC check for verb %s on CustomResourceDefinitions", verbs)
		}
		if !sar.Status.Allowed {
			return false, nil
		}
	}
	return true, nil
}
