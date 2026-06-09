package edgelocation

import (
	"context"
	"strings"

	"github.com/crossplane/upjet/pkg/config"
)

// nonExistentUUID is a valid UUID format that will never match a real resource.
// The CAST AI provider returns 400 for empty IDs but 404 for valid UUIDs that
// don't exist. Terraform treats 404 as "not found" and proceeds to create the
// resource, whereas 400 is treated as a hard error.
const nonExistentUUID = "00000000-0000-0000-0000-000000000000"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("castai_edge_location", func(r *config.Resource) {
		r.ShortGroup = ""
		r.Kind = "EdgeLocation"
		r.References = config.References{
			"cluster_id": {
				Type: "github.com/crossplane-contrib/crossplane-provider-castai/apis/castai/v1alpha1.OmniCluster",
			},
		}
		// The CAST AI provider returns HTTP 400 when the resource ID is empty,
		// instead of returning an empty state (not found). Upjet reconstructs
		// the TF state with "id":"" for new resources, causing the provider to
		// fail with 400 on the first refresh. We override GetIDFn to return a
		// placeholder UUID when the external name is empty, so the provider
		// returns 404 (not found) instead of 400 (bad request), allowing
		// Terraform to proceed with resource creation.
		r.ExternalName.GetIDFn = func(_ context.Context, externalName string, _ map[string]any, _ map[string]any) (string, error) {
			if strings.TrimSpace(externalName) == "" {
				return nonExistentUUID, nil
			}
			return externalName, nil
		}
		// castai_edge_location uses the Terraform Plugin Framework with
		// nesting_mode=single blocks. Upjet v1 converts these to lists in the
		// CRD, but Terraform expects plain objects. AddSingletonListConversion
		// makes Upjet serialize them as objects in the generated main.tf.json.
		for _, field := range []string{
			"aws",
			"control_plane",
			"custom",
			"gcp",
			"networking",
			"networking.cni",
			"oci",
		} {
			r.AddSingletonListConversion(
				field,
				field,
			)
		}
	})
}
