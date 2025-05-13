package gkeclusterid

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("castai_gke_cluster_id", func(r *config.Resource) {
		r.ShortGroup = ""
		r.Kind = "GkeClusterId"
		r.References = config.References{
			"cluster_id": {
				Type: "github.com/crossplane-contrib/crossplane-provider-castai/apis/castai/v1alpha1.GkeCluster",
			},
		}
	})
}
