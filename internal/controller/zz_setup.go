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

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/terrajet/pkg/controller"

	blockdevice "github.com/crossplane-contrib/provider-jet-vra/internal/controller/blockdevice/blockdevice"
	blockdevicesnapshot "github.com/crossplane-contrib/provider-jet-vra/internal/controller/blockdevice/blockdevicesnapshot"
	blueprint "github.com/crossplane-contrib/provider-jet-vra/internal/controller/blueprint/blueprint"
	version "github.com/crossplane-contrib/provider-jet-vra/internal/controller/blueprint/version"
	itementitlement "github.com/crossplane-contrib/provider-jet-vra/internal/controller/catalogitementitlement/itementitlement"
	sourceblueprint "github.com/crossplane-contrib/provider-jet-vra/internal/controller/catalogsourceblueprint/sourceblueprint"
	sourceentitlement "github.com/crossplane-contrib/provider-jet-vra/internal/controller/catalogsourceentitlement/sourceentitlement"
	accountaws "github.com/crossplane-contrib/provider-jet-vra/internal/controller/cloudaccount/accountaws"
	accountazure "github.com/crossplane-contrib/provider-jet-vra/internal/controller/cloudaccount/accountazure"
	accountgcp "github.com/crossplane-contrib/provider-jet-vra/internal/controller/cloudaccount/accountgcp"
	accountnsxt "github.com/crossplane-contrib/provider-jet-vra/internal/controller/cloudaccount/accountnsxt"
	accountnsxv "github.com/crossplane-contrib/provider-jet-vra/internal/controller/cloudaccount/accountnsxv"
	accountvmc "github.com/crossplane-contrib/provider-jet-vra/internal/controller/cloudaccount/accountvmc"
	accountvsphere "github.com/crossplane-contrib/provider-jet-vra/internal/controller/cloudaccount/accountvsphere"
	source "github.com/crossplane-contrib/provider-jet-vra/internal/controller/contentsource/source"
	deployment "github.com/crossplane-contrib/provider-jet-vra/internal/controller/deployment/deployment"
	compute "github.com/crossplane-contrib/provider-jet-vra/internal/controller/fabric/compute"
	datastorevsphere "github.com/crossplane-contrib/provider-jet-vra/internal/controller/fabric/datastorevsphere"
	networkvsphere "github.com/crossplane-contrib/provider-jet-vra/internal/controller/fabric/networkvsphere"
	profile "github.com/crossplane-contrib/provider-jet-vra/internal/controller/flavorprofile/profile"
	profileimageprofile "github.com/crossplane-contrib/provider-jet-vra/internal/controller/imageprofile/profile"
	balancer "github.com/crossplane-contrib/provider-jet-vra/internal/controller/loadbalancer/balancer"
	machine "github.com/crossplane-contrib/provider-jet-vra/internal/controller/machine/machine"
	iprange "github.com/crossplane-contrib/provider-jet-vra/internal/controller/network/iprange"
	network "github.com/crossplane-contrib/provider-jet-vra/internal/controller/network/network"
	profilenetwork "github.com/crossplane-contrib/provider-jet-vra/internal/controller/network/profile"
	project "github.com/crossplane-contrib/provider-jet-vra/internal/controller/project/project"
	providerconfig "github.com/crossplane-contrib/provider-jet-vra/internal/controller/providerconfig"
	profilestorage "github.com/crossplane-contrib/provider-jet-vra/internal/controller/storage/profile"
	profileaws "github.com/crossplane-contrib/provider-jet-vra/internal/controller/storage/profileaws"
	profileazure "github.com/crossplane-contrib/provider-jet-vra/internal/controller/storage/profileazure"
	profilevsphere "github.com/crossplane-contrib/provider-jet-vra/internal/controller/storage/profilevsphere"
	zone "github.com/crossplane-contrib/provider-jet-vra/internal/controller/zone/zone"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		blockdevice.Setup,
		blockdevicesnapshot.Setup,
		blueprint.Setup,
		version.Setup,
		itementitlement.Setup,
		sourceblueprint.Setup,
		sourceentitlement.Setup,
		accountaws.Setup,
		accountazure.Setup,
		accountgcp.Setup,
		accountnsxt.Setup,
		accountnsxv.Setup,
		accountvmc.Setup,
		accountvsphere.Setup,
		source.Setup,
		deployment.Setup,
		compute.Setup,
		datastorevsphere.Setup,
		networkvsphere.Setup,
		profile.Setup,
		profileimageprofile.Setup,
		balancer.Setup,
		machine.Setup,
		iprange.Setup,
		network.Setup,
		profilenetwork.Setup,
		project.Setup,
		providerconfig.Setup,
		profilestorage.Setup,
		profileaws.Setup,
		profileazure.Setup,
		profilevsphere.Setup,
		zone.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
