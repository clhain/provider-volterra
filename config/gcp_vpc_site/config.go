package gcp_vpc_site

import (
	"github.com/clhain/provider-volterra/config/common"
	"github.com/upbound/upjet/pkg/config"
	ujconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("volterra_gcp_vpc_site", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
		r.Kind = "GCPVPCSite"
		r.References["cloud_credentials.name"] = config.Reference{
			Type:      "github.com/clhain/provider-volterra/apis/volterra/v1alpha1.CloudCredentials",
			Extractor: common.ExtractResourceNameFuncPath,
		}
	})
}
