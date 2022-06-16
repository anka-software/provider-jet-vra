package network

import (
	"github.com/crossplane-contrib/provider-jet-vra/config/common"

	"github.com/crossplane/terrajet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vra_network", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha1
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = "network"
	})

	p.AddResourceConfigurator("network_ip_range", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha1
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = "network"

		// r.References["fabric_network_id"]
	})

	p.AddResourceConfigurator("vra_network_profile", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha1
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = "network"

		// r.References["region_id"]
		// r.References["fabric_network_ids"]
	})
}
