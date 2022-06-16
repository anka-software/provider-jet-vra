package blueprint

import (
	"github.com/crossplane-contrib/provider-jet-vra/config/common"

	"github.com/crossplane/terrajet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vra_blueprint", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha1
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = "blueprint"

		r.References["project_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-jet-vra/apis/project/v1alpha1.Project",
		}
	})

	p.AddResourceConfigurator("vra_blueprint_version", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha1
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = "blueprint"

		r.References["blueprint_id"] = config.Reference{
			Type: "Blueprint", // "github.com/crossplane-contrib/provider-jet-vra/apis/blueprint/v1alpha1.Blueprint",
		}
	})
}
