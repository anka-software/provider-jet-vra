//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefinitionObservation) DeepCopyInto(out *DefinitionObservation) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IconID != nil {
		in, out := &in.IconID, &out.IconID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NumberOfItems != nil {
		in, out := &in.NumberOfItems, &out.NumberOfItems
		*out = new(float64)
		**out = **in
	}
	if in.SourceName != nil {
		in, out := &in.SourceName, &out.SourceName
		*out = new(string)
		**out = **in
	}
	if in.SourceType != nil {
		in, out := &in.SourceType, &out.SourceType
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefinitionObservation.
func (in *DefinitionObservation) DeepCopy() *DefinitionObservation {
	if in == nil {
		return nil
	}
	out := new(DefinitionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefinitionParameters) DeepCopyInto(out *DefinitionParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefinitionParameters.
func (in *DefinitionParameters) DeepCopy() *DefinitionParameters {
	if in == nil {
		return nil
	}
	out := new(DefinitionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ItemEntitlement) DeepCopyInto(out *ItemEntitlement) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ItemEntitlement.
func (in *ItemEntitlement) DeepCopy() *ItemEntitlement {
	if in == nil {
		return nil
	}
	out := new(ItemEntitlement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ItemEntitlement) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ItemEntitlementList) DeepCopyInto(out *ItemEntitlementList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ItemEntitlement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ItemEntitlementList.
func (in *ItemEntitlementList) DeepCopy() *ItemEntitlementList {
	if in == nil {
		return nil
	}
	out := new(ItemEntitlementList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ItemEntitlementList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ItemEntitlementObservation) DeepCopyInto(out *ItemEntitlementObservation) {
	*out = *in
	if in.Definition != nil {
		in, out := &in.Definition, &out.Definition
		*out = make([]DefinitionObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ItemEntitlementObservation.
func (in *ItemEntitlementObservation) DeepCopy() *ItemEntitlementObservation {
	if in == nil {
		return nil
	}
	out := new(ItemEntitlementObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ItemEntitlementParameters) DeepCopyInto(out *ItemEntitlementParameters) {
	*out = *in
	if in.CatalogItemID != nil {
		in, out := &in.CatalogItemID, &out.CatalogItemID
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ItemEntitlementParameters.
func (in *ItemEntitlementParameters) DeepCopy() *ItemEntitlementParameters {
	if in == nil {
		return nil
	}
	out := new(ItemEntitlementParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ItemEntitlementSpec) DeepCopyInto(out *ItemEntitlementSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ItemEntitlementSpec.
func (in *ItemEntitlementSpec) DeepCopy() *ItemEntitlementSpec {
	if in == nil {
		return nil
	}
	out := new(ItemEntitlementSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ItemEntitlementStatus) DeepCopyInto(out *ItemEntitlementStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ItemEntitlementStatus.
func (in *ItemEntitlementStatus) DeepCopy() *ItemEntitlementStatus {
	if in == nil {
		return nil
	}
	out := new(ItemEntitlementStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceBlueprint) DeepCopyInto(out *SourceBlueprint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceBlueprint.
func (in *SourceBlueprint) DeepCopy() *SourceBlueprint {
	if in == nil {
		return nil
	}
	out := new(SourceBlueprint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SourceBlueprint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceBlueprintList) DeepCopyInto(out *SourceBlueprintList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SourceBlueprint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceBlueprintList.
func (in *SourceBlueprintList) DeepCopy() *SourceBlueprintList {
	if in == nil {
		return nil
	}
	out := new(SourceBlueprintList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SourceBlueprintList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceBlueprintObservation) DeepCopyInto(out *SourceBlueprintObservation) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.CreatedBy != nil {
		in, out := &in.CreatedBy, &out.CreatedBy
		*out = new(string)
		**out = **in
	}
	if in.Global != nil {
		in, out := &in.Global, &out.Global
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ItemsFound != nil {
		in, out := &in.ItemsFound, &out.ItemsFound
		*out = new(string)
		**out = **in
	}
	if in.ItemsImported != nil {
		in, out := &in.ItemsImported, &out.ItemsImported
		*out = new(string)
		**out = **in
	}
	if in.LastImportCompletedAt != nil {
		in, out := &in.LastImportCompletedAt, &out.LastImportCompletedAt
		*out = new(string)
		**out = **in
	}
	if in.LastImportErrors != nil {
		in, out := &in.LastImportErrors, &out.LastImportErrors
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.LastImportStartedAt != nil {
		in, out := &in.LastImportStartedAt, &out.LastImportStartedAt
		*out = new(string)
		**out = **in
	}
	if in.LastUpdatedBy != nil {
		in, out := &in.LastUpdatedBy, &out.LastUpdatedBy
		*out = new(string)
		**out = **in
	}
	if in.TypeID != nil {
		in, out := &in.TypeID, &out.TypeID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceBlueprintObservation.
func (in *SourceBlueprintObservation) DeepCopy() *SourceBlueprintObservation {
	if in == nil {
		return nil
	}
	out := new(SourceBlueprintObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceBlueprintParameters) DeepCopyInto(out *SourceBlueprintParameters) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceBlueprintParameters.
func (in *SourceBlueprintParameters) DeepCopy() *SourceBlueprintParameters {
	if in == nil {
		return nil
	}
	out := new(SourceBlueprintParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceBlueprintSpec) DeepCopyInto(out *SourceBlueprintSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceBlueprintSpec.
func (in *SourceBlueprintSpec) DeepCopy() *SourceBlueprintSpec {
	if in == nil {
		return nil
	}
	out := new(SourceBlueprintSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceBlueprintStatus) DeepCopyInto(out *SourceBlueprintStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceBlueprintStatus.
func (in *SourceBlueprintStatus) DeepCopy() *SourceBlueprintStatus {
	if in == nil {
		return nil
	}
	out := new(SourceBlueprintStatus)
	in.DeepCopyInto(out)
	return out
}
