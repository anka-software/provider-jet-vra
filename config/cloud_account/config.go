package cloudaccount

import (
	"github.com/crossplane-contrib/provider-jet-vra/config/common"

	"github.com/crossplane/terrajet/pkg/config"
)

const shortGroup = "cloudaccount"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vra_cloud_account_aws", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha1
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
	})

	p.AddResourceConfigurator("vra_cloud_account_azure", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha1
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
	})

	p.AddResourceConfigurator("vra_cloud_account_gcp", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha1
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
	})

	p.AddResourceConfigurator("vra_cloud_account_nsxt", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha1
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
	})

	p.AddResourceConfigurator("vra_cloud_account_nsxv", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha1
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
	})

	p.AddResourceConfigurator("vra_cloud_account_vmc", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha1
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
	})

	p.AddResourceConfigurator("vra_cloud_account_vsphere", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha1
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup

		// r.References["associated_cloud_account_ids"]
	})
}
