package catalogsourceblueprint

import (
	"github.com/crossplane-contrib/provider-jet-vra/config/common"

	"github.com/crossplane/terrajet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vra_catalog_source_blueprint", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha1
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = "catalogsourceblueprint"

		r.References["project_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-jet-vra/apis/project/v1alpha1.Project",
		}
	})

}
