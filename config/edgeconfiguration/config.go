package edgeconfiguration

import (
	"github.com/crossplane/upjet/pkg/config"
)

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
		// castai_edge_configuration uses the Terraform Plugin Framework with
		// nesting_mode=single blocks that must be serialized as objects.
		for _, field := range []string{"aws", "cri", "custom", "gcp", "oci"} {
			r.AddSingletonListConversion(field, field)
		}
	})
}
