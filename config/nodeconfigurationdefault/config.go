package nodeconfigurationdefault

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("castai_node_configuration_default", func(r *config.Resource) {
		r.ShortGroup = ""
		r.Kind = "NodeConfigurationDefault"
		r.References = config.References{
			"cluster_id": {
				Type: "github.com/castai/crossplane-provider-castai/apis/castai/v1alpha1.EksClusterId",
			},
			"configuration_id": {
				Type: "github.com/castai/crossplane-provider-castai/apis/castai/v1alpha1.NodeConfiguration",
			},
		}
	})
}
