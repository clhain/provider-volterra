package volterra

import (
	"github.com/clhain/provider-volterra/config/common"
	"github.com/upbound/upjet/pkg/config"
	ujconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("volterra_app_firewall", func(r *ujconfig.Resource) {
		r.References["namespace"] = config.Reference{
			Type:      "github.com/clhain/provider-volterra/apis/volterra/v1alpha1.VolterraNamespace",
			Extractor: common.ExtractResourceNameFuncPath,
		}
	})
	p.AddResourceConfigurator("volterra_aws_vpc_site", func(r *ujconfig.Resource) {
		r.Kind = "AWSVPCSite"
		r.References["cloud_credentials.name"] = config.Reference{
			Type:      "github.com/clhain/provider-volterra/apis/volterra/v1alpha1.CloudCredentials",
			Extractor: common.ExtractResourceNameFuncPath,
		}
	})
	p.AddResourceConfigurator("volterra_cloud_credentials", func(r *ujconfig.Resource) {
	})
	p.AddResourceConfigurator("volterra_gcp_vpc_site", func(r *ujconfig.Resource) {
		r.Kind = "GCPVPCSite"
		r.References["cloud_credentials.name"] = config.Reference{
			Type:      "github.com/clhain/provider-volterra/apis/volterra/v1alpha1.CloudCredentials",
			Extractor: common.ExtractResourceNameFuncPath,
		}
	})
	p.AddResourceConfigurator("volterra_http_loadbalancer", func(r *ujconfig.Resource) {
		// r.References["app_firewall"] = config.Reference{
		// 	Type: "github.com/clhain/provider-volterra/apis/volterra/v1alpha1.AppFirewall",
		// }
		r.References["namespace"] = config.Reference{
			Type:      "github.com/clhain/provider-volterra/apis/volterra/v1alpha1.VolterraNamespace",
			Extractor: common.ExtractResourceNameFuncPath,
		}
		// r.References["default_route_pools.pool"] = config.Reference{
		// 	Type: "github.com/clhain/provider-volterra/apis/origin_pool/v1alpha1.OriginPool",
		// }
	})
	p.AddResourceConfigurator("volterra_namespace", func(r *ujconfig.Resource) {
		r.Kind = "VolterraNamespace"
	})
	p.AddResourceConfigurator("volterra_origin_pool", func(r *ujconfig.Resource) {
		r.References["namespace"] = config.Reference{
			Type:      "github.com/clhain/provider-volterra/apis/volterra/v1alpha1.VolterraNamespace",
			Extractor: common.ExtractResourceNameFuncPath,
		}
	})
	p.AddResourceConfigurator("volterra_tf_params_action", func(r *ujconfig.Resource) {
		r.Kind = "TFParamsAction"
		r.References["site_name"] = config.Reference{
			Type:      "github.com/clhain/provider-volterra/apis/volterra/v1alpha1.GCPVPCSite",
			Extractor: common.ExtractResourceNameFuncPath,
		}
	})
}
