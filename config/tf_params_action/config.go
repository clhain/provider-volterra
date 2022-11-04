package tf_params_action

import (
	"github.com/clhain/provider-volterra/config/common"
	"github.com/upbound/upjet/pkg/config"
	ujconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("volterra_tf_params_action", func(r *ujconfig.Resource) {
		r.Kind = "TFParamsAction"
		r.ShortGroup = ""
		r.References["site_name"] = config.Reference{
			Type:      "github.com/clhain/provider-volterra/apis/volterra/v1alpha1.GCPVPCSite",
			Extractor: common.ExtractResourceNameFuncPath,
		}
	})
}
