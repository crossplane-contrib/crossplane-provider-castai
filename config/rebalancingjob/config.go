package rebalancingjob

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("castai_rebalancing_job", func(r *config.Resource) {
		r.ShortGroup = ""
		r.Kind = "RebalancingJob"
		r.References = config.References{
			"cluster_id": {
				Type: "github.com/crossplane-contrib/crossplane-provider-castai/apis/castai/v1alpha1.EksClusterId",
			},
			"rebalancing_schedule_id": {
				Type: "github.com/crossplane-contrib/crossplane-provider-castai/apis/castai/v1alpha1.RebalancingSchedule",
			},
		}
	})
}
