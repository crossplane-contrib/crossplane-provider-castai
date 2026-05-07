/*
Copyright 2022 Upbound Inc.
*/

package namespaced

import "github.com/crossplane/upjet/v2/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"castai_ai_optimizer_hosted_model":           config.IdentifierFromProvider,
	"castai_ai_optimizer_model_registry":         config.IdentifierFromProvider,
	"castai_ai_optimizer_model_specs":            config.IdentifierFromProvider,
	"castai_aks_cluster":                         config.IdentifierFromProvider,
	"castai_allocation_group":                    config.IdentifierFromProvider,
	"castai_autoscaler":                          config.IdentifierFromProvider,
	"castai_cache_configuration":                 config.IdentifierFromProvider,
	"castai_cache_group":                         config.IdentifierFromProvider,
	"castai_cache_rule":                          config.IdentifierFromProvider,
	"castai_commitments":                         config.IdentifierFromProvider,
	"castai_eks_cluster":                         config.IdentifierFromProvider,
	"castai_eks_clusterid":                       config.IdentifierFromProvider,
	"castai_eks_user_arn":                        config.IdentifierFromProvider,
	"castai_enterprise_group":                    config.IdentifierFromProvider,
	"castai_enterprise_role_binding":             config.IdentifierFromProvider,
	"castai_evictor_advanced_config":             config.IdentifierFromProvider,
	"castai_gke_cluster":                         config.IdentifierFromProvider,
	"castai_gke_cluster_id":                      config.IdentifierFromProvider,
	"castai_hibernation_schedule":                config.IdentifierFromProvider,
	"castai_node_configuration":                  config.IdentifierFromProvider,
	"castai_node_configuration_default":          config.IdentifierFromProvider,
	"castai_node_template":                       config.IdentifierFromProvider,
	"castai_organization_group":                  config.IdentifierFromProvider,
	"castai_organization_members":                config.IdentifierFromProvider,
	"castai_pod_mutation":                        config.IdentifierFromProvider,
	"castai_rebalancing_job":                     config.IdentifierFromProvider,
	"castai_rebalancing_schedule":                config.IdentifierFromProvider,
	"castai_reservations":                        config.IdentifierFromProvider,
	"castai_role_bindings":                       config.IdentifierFromProvider,
	"castai_security_runtime_rule":               config.IdentifierFromProvider,
	"castai_service_account":                     config.IdentifierFromProvider,
	"castai_service_account_key":                 config.IdentifierFromProvider,
	"castai_sso_connection":                      config.IdentifierFromProvider,
	"castai_workload_custom_metrics_data_source": config.IdentifierFromProvider,
	"castai_workload_scaling_policy":             config.IdentifierFromProvider,
	"castai_workload_scaling_policy_order":       config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
