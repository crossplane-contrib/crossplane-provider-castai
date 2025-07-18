package workloadscalingpolicyorder

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("castai_workload_scaling_policy_order", func(r *config.Resource) {
		r.ShortGroup = ""
		r.Kind = "ScalingPolicyOrder"
	})
}
