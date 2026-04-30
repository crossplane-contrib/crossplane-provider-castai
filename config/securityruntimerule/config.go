package securityruntimerule

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("castai_security_runtime_rule", func(r *config.Resource) {
		r.ShortGroup = ""
		r.Kind = "SecurityRuntimeRule"
	})
}
