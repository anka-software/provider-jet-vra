package storage

import (
	"github.com/crossplane-contrib/provider-jet-vra/config/common"

	"github.com/crossplane/terrajet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vra_storage_profile", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha1
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = "storage"

		// r.References["region_id"]
		// r.References["disk_target_properties"]
	})

	p.AddResourceConfigurator("vra_storage_profile_aws", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha1
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = "storage"

		// r.References["region_id"]
	})

	p.AddResourceConfigurator("vra_storage_profile_azure", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha1
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = "storage"

		// r.References["region_id"]
	})

	p.AddResourceConfigurator("vra_storage_profile_vsphere", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha1
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = "storage"

		// r.References["region_id"]
		// r.References["datastore_id"]
		// r.References["storage_policy_id"]
	})
}
