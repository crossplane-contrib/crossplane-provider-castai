package gkeclusterid

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/registry"
)

// deprecationNote becomes the GkeClusterId kind description. Upstream removed
// the castai_gke_cluster_id docs page when it deprecated the resource in
// 9.18.0, so there is no scraped description left to use, and the deprecation
// warning itself never reaches Crossplane users.
const deprecationNote = "Deprecated: CAST AI deprecated this resource in provider 9.18.0 and will remove it in a future version. " +
	"Use GkeCluster with targetServiceAccountEmail and castaiServiceAccountEmail instead. " +
	"Deleting a GkeClusterId whose cluster has credentials assigned, for example by a GkeCluster, no longer deregisters the cluster. " +
	"It stays registered in CAST AI until the GkeCluster that owns it is deleted."

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
		if r.MetaResource == nil {
			r.MetaResource = &registry.Resource{Name: "castai_gke_cluster_id"}
		}
		r.MetaResource.Description = deprecationNote
	})
}
