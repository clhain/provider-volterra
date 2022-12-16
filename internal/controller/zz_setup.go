/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	providerconfig "github.com/clhain/provider-volterra/internal/controller/providerconfig"
	appfirewall "github.com/clhain/provider-volterra/internal/controller/volterra/appfirewall"
	awsvpcsite "github.com/clhain/provider-volterra/internal/controller/volterra/awsvpcsite"
	cloudcredentials "github.com/clhain/provider-volterra/internal/controller/volterra/cloudcredentials"
	gcpvpcsite "github.com/clhain/provider-volterra/internal/controller/volterra/gcpvpcsite"
	healthcheck "github.com/clhain/provider-volterra/internal/controller/volterra/healthcheck"
	httploadbalancer "github.com/clhain/provider-volterra/internal/controller/volterra/httploadbalancer"
	originpool "github.com/clhain/provider-volterra/internal/controller/volterra/originpool"
	tfparamsaction "github.com/clhain/provider-volterra/internal/controller/volterra/tfparamsaction"
	volterranamespace "github.com/clhain/provider-volterra/internal/controller/volterra/volterranamespace"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerconfig.Setup,
		appfirewall.Setup,
		awsvpcsite.Setup,
		cloudcredentials.Setup,
		gcpvpcsite.Setup,
		healthcheck.Setup,
		httploadbalancer.Setup,
		originpool.Setup,
		tfparamsaction.Setup,
		volterranamespace.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
