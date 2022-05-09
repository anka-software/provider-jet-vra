package blueprintversion

import "github.com/crossplane/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vra_blueprint_version", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = "blueprintversion"

		r.References["blueprint"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-jet-vra/apis/blueprint/v1alpha1.Blueprint",
		}
	})
}
