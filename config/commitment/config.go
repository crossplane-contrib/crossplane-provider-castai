package commitment

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("castai_commitment", func(r *config.Resource) {
		r.ShortGroup = ""
		r.Kind = "Commitment"
		// The existing Kind "Commitments" (castai_commitments) already owns the
		// plural "commitments", so the CRD name commitments.castai.upbound.io
		// is taken. Override the API path/plural to avoid the collision.
		r.Path = "cloudcommitments"
		// castai_commitment uses the Terraform Plugin Framework with
		// nesting_mode=single blocks for its per-cloud detail attributes. Upjet
		// v1 converts these to lists in the CRD, but Terraform expects plain
		// objects. AddSingletonListConversion makes Upjet serialize them as
		// objects in the generated main.tf.json.
		for _, field := range []string{
			"aws_capacity_block_details",
			"aws_odcr_details",
			"aws_reserved_instances_details",
			"aws_savings_plan_details",
			"azure_reservation_details",
			"azure_savings_plan_details",
			"gcp_capacity_reservation_details",
			"gcp_flex_cud_details",
			"gcp_resource_cud_details",
		} {
			r.AddSingletonListConversion(
				field,
				field,
			)
		}
	})
}
