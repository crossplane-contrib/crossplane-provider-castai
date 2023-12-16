package ssoconnation

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("castai_sso_connection", func(r *config.Resource) {
		r.ShortGroup = ""
		r.Kind = "SSOConnection"
	})
}
