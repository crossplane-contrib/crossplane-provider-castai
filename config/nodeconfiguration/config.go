package nodeconfiguration

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("castai_node_configuration", func(r *config.Resource) {
		r.ShortGroup = ""
		r.Kind = "NodeConfiguration"
		r.References = config.References{
			"cluster_id": {
				Type: "github.com/dkb-bank/provider-castai/apis/castai/v1alpha1.EksClusterId",
			},
		}
	})
}
