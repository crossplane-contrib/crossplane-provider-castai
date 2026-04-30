package aioptimizermodelregistry

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("castai_ai_optimizer_model_registry", func(r *config.Resource) {
		r.ShortGroup = ""
		r.Kind = "AIOptimizerModelRegistry"
	})
}
