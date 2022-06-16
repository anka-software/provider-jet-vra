/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	tjconfig "github.com/crossplane/terrajet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	blockdevice "github.com/crossplane-contrib/provider-jet-vra/config/block_device"
	"github.com/crossplane-contrib/provider-jet-vra/config/blueprint"
	catalog_item_entitlement "github.com/crossplane-contrib/provider-jet-vra/config/catalog_item_entitlement"
	catalog_source_blueprint "github.com/crossplane-contrib/provider-jet-vra/config/catalog_source_blueprint"
	catalog_source_entitlement "github.com/crossplane-contrib/provider-jet-vra/config/catalog_source_entitlement"
	cloudaccount "github.com/crossplane-contrib/provider-jet-vra/config/cloud_account"
	content_source "github.com/crossplane-contrib/provider-jet-vra/config/content_source"
	"github.com/crossplane-contrib/provider-jet-vra/config/deployment"
	"github.com/crossplane-contrib/provider-jet-vra/config/fabric"
	flavorprofile "github.com/crossplane-contrib/provider-jet-vra/config/flavor_profile"
	imageprofile "github.com/crossplane-contrib/provider-jet-vra/config/image_profile"
	"github.com/crossplane-contrib/provider-jet-vra/config/loadbalancer"
	"github.com/crossplane-contrib/provider-jet-vra/config/machine"
	"github.com/crossplane-contrib/provider-jet-vra/config/network"
	"github.com/crossplane-contrib/provider-jet-vra/config/project"
	"github.com/crossplane-contrib/provider-jet-vra/config/storage"
	"github.com/crossplane-contrib/provider-jet-vra/config/zone"
)

const (
	resourcePrefix = "vra"
	modulePath     = "github.com/crossplane-contrib/provider-jet-vra"
)

//go:embed schema.json
var providerSchema string

// IncludedResources s
var IncludedResources = []string{
	"vra_deployment$",
	"vra_blueprint$",
	"vra_blueprint_version$",

	"vra_catalog_item_entitlement$",
	"vra_catalog_source_blueprint$",
	"vra_catalog_source_entitlement$",
	"vra_content_source$",

	"vra_project$",

	"vra_block_device$",
	"vra_block_device_snapshot$",

	"vra_fabric_compute$",
	"vra_fabric_datastore_vsphere$",
	"vra_fabric_network_vsphere$",

	"vra_network$",
	"vra_network_ip_range$",
	"vra_network_profile$",
	"vra_flavor_profile$",
	"vra_image_profile$",

	"vra_machine$",
	"vra_load_balancer$",
	"vra_zone$",

	"vra_cloud_account_aws$",
	"vra_cloud_account_azure$",
	"vra_cloud_account_nsxt$",
	"vra_cloud_account_gcp$",
	"vra_cloud_account_nsxv$",
	"vra_cloud_account_vmc$",
	"vra_cloud_account_vsphere$",

	"vra_storage_profile$",
	"vra_storage_profile_aws$",
	"vra_storage_profile_azure$",
	"vra_storage_profile_vsphere$",
}

// skipList
var skipList = []string{}

// GetProvider returns provider configuration
func GetProvider() *tjconfig.Provider {
	defaultResourceFn := func(name string, terraformResource *schema.Resource, opts ...tjconfig.ResourceOption) *tjconfig.Resource {
		r := tjconfig.DefaultResource(name, terraformResource)
		// Add any provider-specific defaulting here. For example:
		//   r.ExternalName = tjconfig.IdentifierFromProvider
		return r
	}

	pc := tjconfig.NewProviderWithSchema([]byte(providerSchema), resourcePrefix, modulePath,
		tjconfig.WithShortName("vrajet"),
		tjconfig.WithRootGroup("vra.jet.crossplane.io"),
		tjconfig.WithDefaultResourceFn(defaultResourceFn),
		tjconfig.WithIncludeList(IncludedResources),
		tjconfig.WithSkipList(skipList))

	for _, configure := range []func(provider *tjconfig.Provider){
		// add custom config functions
		blockdevice.Configure,
		blueprint.Configure,
		catalog_item_entitlement.Configure,
		catalog_source_blueprint.Configure,
		catalog_source_entitlement.Configure,
		cloudaccount.Configure,
		content_source.Configure,
		deployment.Configure,
		fabric.Configure,
		flavorprofile.Configure,
		imageprofile.Configure,
		loadbalancer.Configure,
		machine.Configure,
		network.Configure,
		project.Configure,
		storage.Configure,
		zone.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
