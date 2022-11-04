package common

import (
	jresource "github.com/upbound/upjet/pkg/resource"

	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
)

const (
	// SelfPackagePath is the golang path for this package.
	SelfPackagePath = "github.com/clhain/provider-volterra/config/common"

	// ExtractResourceIDFuncPath holds the GCP resource ID extractor func name
	ExtractResourceNameFuncPath = "github.com/clhain/provider-volterra/config/common.ExtractResourceName()"
)

// GetField returns the value of field as a string in a map[string]interface{},
//  fails properly otherwise.
func GetField(from map[string]interface{}, path string) (string, error) {
	return fieldpath.Pave(from).GetString(path)
}

func ExtractResourceName() reference.ExtractValueFn {
	return func(mr resource.Managed) string {
		tr, ok := mr.(jresource.Terraformed)
		if !ok {
			return ""
		}
		p, err := tr.GetParameters()
		if err != nil {
			return ""
		}
		n, ok := p["name"]
		if !ok {
			return ""
		}
		return n.(string)
	}
}
