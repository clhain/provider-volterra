/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/upbound/upjet/pkg/terraform"

	"github.com/clhain/provider-volterra/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal volterra credentials as JSON"
	keyAPIP12File           = "api_p12_file"
	keyAPICert              = "api_cert"
	keyAPIKey               = "api_key"
	keyURL                  = "url"
	keyTimeout              = "timeout"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		// Since volterra uses paths to files rather than secret values, this is useless at the moment.
		// data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		// if err != nil {
		// 	return ps, errors.Wrap(err, errExtractCredentials)
		// }
		// creds := map[string]string{}
		// if err := json.Unmarshal(data, &creds); err != nil {
		// 	return ps, errors.Wrap(err, errUnmarshalCredentials)
		// }

		// Set credentials in Terraform provider configuration.
		ps.Configuration = map[string]interface{}{
			keyAPIP12File: pc.Spec.APIP12File,
			keyAPICert:    pc.Spec.APICert,
			keyAPIKey:     pc.Spec.APIKey,
			keyURL:        pc.Spec.URL,
			keyTimeout:    pc.Spec.Timeout,
		}
		return ps, nil
	}
}
