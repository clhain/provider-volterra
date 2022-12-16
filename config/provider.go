/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	"github.com/clhain/provider-volterra/config/volterra"
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
		volterra.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
