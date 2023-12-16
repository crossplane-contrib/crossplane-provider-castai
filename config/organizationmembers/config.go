package organizationmembers

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("castai_organization_members", func(r *config.Resource) {
		r.ShortGroup = ""
		r.Kind = "OrganizationMembers"
	})
}
