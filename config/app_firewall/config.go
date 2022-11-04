package app_firewall

import (
	"github.com/clhain/provider-volterra/config/common"
	"github.com/upbound/upjet/pkg/config"
	ujconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("volterra_app_firewall", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
		r.References["namespace"] = config.Reference{
			Type:      "github.com/clhain/provider-volterra/apis/volterra/v1alpha1.VolterraNamespace",
			Extractor: common.ExtractResourceNameFuncPath,
		}
	})
}
