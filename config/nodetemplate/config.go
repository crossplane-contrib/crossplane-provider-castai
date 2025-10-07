package nodetemplate

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("castai_node_template", func(r *config.Resource) {
		r.ShortGroup = ""
		r.Kind = "NodeTemplate"
		r.References = config.References{
			"cluster_id": {
				Type: "github.com/crossplane-contrib/crossplane-provider-castai/apis/castai/v1alpha1.EksClusterId",
			},
			"configuration_id": {
				Type: "github.com/crossplane-contrib/crossplane-provider-castai/apis/castai/v1alpha1.NodeConfiguration",
			},
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"custom_labels"},
		}
	})
}
