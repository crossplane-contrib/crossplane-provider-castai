package omnicluster

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("castai_omni_cluster", func(r *config.Resource) {
		r.ShortGroup = ""
		r.Kind = "OmniCluster"
		// castai_omni_cluster uses the Terraform Plugin Framework with
		// nesting_mode=single blocks that must be serialized as objects.
		r.AddSingletonListConversion("status", "status")
	})
}
