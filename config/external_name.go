/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/upbound/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"volterra_app_firewall":      config.IdentifierFromProvider,
	"volterra_cloud_credentials": config.IdentifierFromProvider,
	"volterra_gcp_vpc_site":      config.IdentifierFromProvider,
	"volterra_aws_vpc_site":      config.IdentifierFromProvider,
	"volterra_healthcheck":       config.IdentifierFromProvider,
	"volterra_http_loadbalancer": config.IdentifierFromProvider,
	"volterra_namespace":         config.IdentifierFromProvider,
	"volterra_origin_pool":       config.IdentifierFromProvider,
	"volterra_tf_params_action":  config.IdentifierFromProvider,
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
