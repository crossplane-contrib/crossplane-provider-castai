package gkecluster

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("castai_gke_cluster", func(r *config.Resource) {
		r.ShortGroup = ""
		r.Kind = "GkeCluster"

		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["id"].(string); ok {
				conn["attribute.cluster_id"] = []byte(a)
			}
			return conn, nil
		}
	})
}
