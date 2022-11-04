/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	"github.com/clhain/provider-volterra/config/app_firewall"
	"github.com/clhain/provider-volterra/config/cloud_credentials"
	"github.com/clhain/provider-volterra/config/gcp_vpc_site"
	"github.com/clhain/provider-volterra/config/http_loadbalancer"
	"github.com/clhain/provider-volterra/config/namespace"
	"github.com/clhain/provider-volterra/config/null"
	"github.com/clhain/provider-volterra/config/origin_pool"
	"github.com/clhain/provider-volterra/config/tf_params_action"
)

const (
	resourcePrefix = "volterra"
	modulePath     = "github.com/clhain/provider-volterra"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
			groupOverrides(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		null.Configure,
		namespace.Configure,
		origin_pool.Configure,
		app_firewall.Configure,
		tf_params_action.Configure,
		cloud_credentials.Configure,
		gcp_vpc_site.Configure,
		http_loadbalancer.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
