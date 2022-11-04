/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	common "github.com/clhain/provider-volterra/config/common"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this AppFirewall.
func (mg *AppFirewall) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Namespace),
		Extract:      common.ExtractResourceName(),
		Reference:    mg.Spec.ForProvider.NamespaceRef,
		Selector:     mg.Spec.ForProvider.NamespaceSelector,
		To: reference.To{
			List:    &VolterraNamespaceList{},
			Managed: &VolterraNamespace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Namespace")
	}
	mg.Spec.ForProvider.Namespace = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NamespaceRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this GCPVPCSite.
func (mg *GCPVPCSite) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.CloudCredentials); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CloudCredentials[i3].Name),
			Extract:      common.ExtractResourceName(),
			Reference:    mg.Spec.ForProvider.CloudCredentials[i3].NameRef,
			Selector:     mg.Spec.ForProvider.CloudCredentials[i3].NameSelector,
			To: reference.To{
				List:    &CloudCredentialsList{},
				Managed: &CloudCredentials{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.CloudCredentials[i3].Name")
		}
		mg.Spec.ForProvider.CloudCredentials[i3].Name = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.CloudCredentials[i3].NameRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this HTTPLoadbalancer.
func (mg *HTTPLoadbalancer) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Namespace),
		Extract:      common.ExtractResourceName(),
		Reference:    mg.Spec.ForProvider.NamespaceRef,
		Selector:     mg.Spec.ForProvider.NamespaceSelector,
		To: reference.To{
			List:    &VolterraNamespaceList{},
			Managed: &VolterraNamespace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Namespace")
	}
	mg.Spec.ForProvider.Namespace = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NamespaceRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this OriginPool.
func (mg *OriginPool) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Namespace),
		Extract:      common.ExtractResourceName(),
		Reference:    mg.Spec.ForProvider.NamespaceRef,
		Selector:     mg.Spec.ForProvider.NamespaceSelector,
		To: reference.To{
			List:    &VolterraNamespaceList{},
			Managed: &VolterraNamespace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Namespace")
	}
	mg.Spec.ForProvider.Namespace = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NamespaceRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TFParamsAction.
func (mg *TFParamsAction) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SiteName),
		Extract:      common.ExtractResourceName(),
		Reference:    mg.Spec.ForProvider.SiteNameRef,
		Selector:     mg.Spec.ForProvider.SiteNameSelector,
		To: reference.To{
			List:    &GCPVPCSiteList{},
			Managed: &GCPVPCSite{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SiteName")
	}
	mg.Spec.ForProvider.SiteName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SiteNameRef = rsp.ResolvedReference

	return nil
}