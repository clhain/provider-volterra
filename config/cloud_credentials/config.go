package cloud_credentials

import (
	ujconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("volterra_cloud_credentials", func(r *ujconfig.Resource) {
		r.ShortGroup = ""
	})
}

// resource "volterra_cloud_credentials" "gcp" {
// 	name        = format("%s-cred", var.site_name)
// 	description = format("GCP credential will be used to create site %s", var.site_name)
// 	namespace   = "system"
// 	gcp_cred_file {
// 	  credential_file{
// 		clear_secret_info {
// 		  url = format("string:///%s", filebase64(var.gcp_cred_file_path))
// 		}
// 	  }
// 	}
//   }
