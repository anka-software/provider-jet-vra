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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DefinitionObservation struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IconID *string `json:"iconId,omitempty" tf:"icon_id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	NumberOfItems *float64 `json:"numberOfItems,omitempty" tf:"number_of_items,omitempty"`

	SourceName *string `json:"sourceName,omitempty" tf:"source_name,omitempty"`

	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DefinitionParameters struct {
}

type SourceEntitlementObservation struct {
	Definition []DefinitionObservation `json:"definition,omitempty" tf:"definition,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SourceEntitlementParameters struct {

	// Catalog source id.
	// +kubebuilder:validation:Required
	CatalogSourceID *string `json:"catalogSourceId" tf:"catalog_source_id,omitempty"`

	// Project id.
	// +kubebuilder:validation:Required
	ProjectID *string `json:"projectId" tf:"project_id,omitempty"`
}

// SourceEntitlementSpec defines the desired state of SourceEntitlement
type SourceEntitlementSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SourceEntitlementParameters `json:"forProvider"`
}

// SourceEntitlementStatus defines the observed state of SourceEntitlement.
type SourceEntitlementStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SourceEntitlementObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SourceEntitlement is the Schema for the SourceEntitlements API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vrajet}
type SourceEntitlement struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SourceEntitlementSpec   `json:"spec"`
	Status            SourceEntitlementStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SourceEntitlementList contains a list of SourceEntitlements
type SourceEntitlementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SourceEntitlement `json:"items"`
}

// Repository type metadata.
var (
	SourceEntitlement_Kind             = "SourceEntitlement"
	SourceEntitlement_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SourceEntitlement_Kind}.String()
	SourceEntitlement_KindAPIVersion   = SourceEntitlement_Kind + "." + CRDGroupVersion.String()
	SourceEntitlement_GroupVersionKind = CRDGroupVersion.WithKind(SourceEntitlement_Kind)
)

func init() {
	SchemeBuilder.Register(&SourceEntitlement{}, &SourceEntitlementList{})
}
