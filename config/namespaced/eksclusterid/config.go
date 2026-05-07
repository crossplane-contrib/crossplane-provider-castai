package eksclusterid

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("castai_eks_clusterid", func(r *config.Resource) {
		r.ShortGroup = ""
		r.Kind = "EksClusterId"
		r.References = config.References{
			"cluster_id": {
				Type: "github.com/crossplane-contrib/crossplane-provider-castai/apis/cluster/castai/v1alpha1.EksCluster",
			},
		}
	})
}
