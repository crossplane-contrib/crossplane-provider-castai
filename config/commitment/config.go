package commitment

import (
	"context"
	"strings"

	"github.com/crossplane/upjet/pkg/config"
)

// nonExistentUUID is a valid UUID format that will never match a real resource.
// Mirrors config/edgelocation: Plugin Framework resources in this provider
// return 400 for an empty id but 404 for a well-formed id that does not exist,
// and only 404 lets Terraform proceed to create the resource.
const nonExistentUUID = "00000000-0000-0000-0000-000000000000"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("castai_commitment", func(r *config.Resource) {
		r.ShortGroup = ""
		// The plural kind Commitments (castai_commitments) already owns the
		// CRD commitments.castai.upbound.io, so the singular resource needs a
		// kind whose plural does not collide.
		r.Kind = "CommitmentEntry"
		r.ExternalName.GetIDFn = func(_ context.Context, externalName string, _ map[string]any, _ map[string]any) (string, error) {
			if strings.TrimSpace(externalName) == "" {
				return nonExistentUUID, nil
			}
			return externalName, nil
		}
		// castai_commitment uses the Terraform Plugin Framework with
		// nesting_mode=single detail blocks, one per commitment type. Upjet v1
		// converts these to lists in the CRD, but Terraform expects plain
		// objects, so serialize them as objects in the generated main.tf.json.
		for _, field := range []string{
			"aws_capacity_block_details",
			"aws_odcr_details",
			"aws_reserved_instances_details",
			"aws_savings_plan_details",
			"azure_reservation_details",
			"azure_savings_plan_details",
			"gcp_capacity_reservation_details",
			"gcp_capacity_reservation_details.share_settings",
			"gcp_flex_cud_details",
			"gcp_resource_cud_details",
		} {
			r.AddSingletonListConversion(field, field)
		}
	})
}
