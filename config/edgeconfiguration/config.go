package edgeconfiguration

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
	p.AddResourceConfigurator("castai_edge_configuration", func(r *config.Resource) {
		r.ShortGroup = ""
		r.Kind = "EdgeConfiguration"
		r.References = config.References{
			"cluster_id": {
				Type: "github.com/crossplane-contrib/crossplane-provider-castai/apis/castai/v1alpha1.OmniCluster",
			},
			"edge_location_id": {
				Type: "github.com/crossplane-contrib/crossplane-provider-castai/apis/castai/v1alpha1.EdgeLocation",
			},
		}
		// Same as castai_edge_location: the provider answers 400 to the empty
		// id Upjet uses for a new resource, so refresh with a placeholder UUID
		// that answers 404 and lets Terraform create the resource.
		r.ExternalName.GetIDFn = func(_ context.Context, externalName string, _ map[string]any, _ map[string]any) (string, error) {
			if strings.TrimSpace(externalName) == "" {
				return nonExistentUUID, nil
			}
			return externalName, nil
		}
		// castai_edge_configuration uses the Terraform Plugin Framework with
		// nesting_mode=single blocks that must be serialized as objects.
		for _, field := range []string{"aws", "cri", "custom", "gcp", "nebius", "oci"} {
			r.AddSingletonListConversion(field, field)
		}
	})
}
