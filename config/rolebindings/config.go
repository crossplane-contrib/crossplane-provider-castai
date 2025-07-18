package rolebindings

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("castai_role_bindings", func(r *config.Resource) {
		r.ShortGroup = ""
		r.Kind = "RoleBindings"
	})
}
