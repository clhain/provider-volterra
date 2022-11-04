package namespace

import (
	ujconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("volterra_namespace", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
		r.Kind = "VolterraNamespace"
	})
}
