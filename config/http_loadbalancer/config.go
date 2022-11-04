package http_loadbalancer

import (
	"github.com/clhain/provider-volterra/config/common"
	"github.com/upbound/upjet/pkg/config"
	ujconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("volterra_http_loadbalancer", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
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
}

// resource "volterra_http_loadbalancer" "this" {
// 	name                            = "${var.name_prefix}-loadbalancer"
// 	namespace                       = var.namespace
// 	description                     = "HTTPS loadbalancer object for external service"
// 	domains                         = [var.app_fqdn]
// 	advertise_on_public_default_vip = true
// 	default_route_pools {
// 	  pool {
// 		name      = volterra_origin_pool.this.name
// 		namespace = var.namespace
// 	  }
// 	}
// 	https_auto_cert {
// 	  add_hsts      = true
// 	  http_redirect = true
// 	  no_mtls       = true
// 	}
// 	app_firewall {
// 	  name      = "${var.name_prefix}-app-firewall"
// 	  namespace = var.namespace
// 	}
// 	disable_rate_limit              = true
// 	round_robin                     = true
// 	service_policies_from_namespace = true
// 	no_challenge                    = true
//   }
