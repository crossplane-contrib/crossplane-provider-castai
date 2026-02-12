package workloadscalingpolicyorder

import (
	"context"
	"reflect"

	"sigs.k8s.io/controller-runtime/pkg/client"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("castai_workload_scaling_policy_order", func(r *config.Resource) {
		r.ShortGroup = ""
		r.Kind = "ScalingPolicyOrder"
		r.References = config.References{
			"policy_ids": {
				Type: "github.com/crossplane-contrib/crossplane-provider-castai/apis/castai/v1alpha1.ScalingPolicy",
				// Use status.atProvider.id instead of external-name for reference resolution
				Extractor: `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
			},
		}
		r.InitializerFns = append(r.InitializerFns, newPolicyRefsAlwaysResolver)
	})
}

func newPolicyRefsAlwaysResolver(_ client.Client) managed.Initializer {
	return managed.InitializerFn(func(_ context.Context, mg xpresource.Managed) error {
		always := v1.ResolvePolicyAlways
		setAlways := func(refs []v1.Reference) {
			for i := range refs {
				if refs[i].Policy == nil {
					refs[i].Policy = &v1.Policy{}
				}
				refs[i].Policy.Resolve = &always
			}
		}

		val := reflect.ValueOf(mg).Elem()
		for _, path := range [2][3]string{
			{"Spec", "ForProvider", "PolicyIdsRefs"},
			{"Spec", "InitProvider", "PolicyIdsRefs"},
		} {
			field := val
			for _, name := range path {
				field = field.FieldByName(name)
				if !field.IsValid() {
					break
				}
			}
			if !field.IsValid() {
				continue
			}
			refs, ok := field.Interface().([]v1.Reference)
			if !ok {
				continue
			}
			setAlways(refs)
		}
		return nil
	})
}
