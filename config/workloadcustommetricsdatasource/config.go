package workloadcustommetricsdatasource

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("castai_workload_custom_metrics_data_source", func(r *config.Resource) {
		r.ShortGroup = ""
		r.Kind = "WorkloadCustomMetricsDataSource"
	})
}
