// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type QuotaInitParameters struct {

	// A map of string k/v properties.
	Config map[string]float64 `json:"config,omitempty" tf:"config,omitempty"`
}

type QuotaObservation struct {

	// A map of string k/v properties.
	Config map[string]float64 `json:"config,omitempty" tf:"config,omitempty"`

	// The name of the entity
	EntityName *string `json:"entityName,omitempty" tf:"entity_name,omitempty"`

	// The type of the entity (client-id, user, ip)
	EntityType *string `json:"entityType,omitempty" tf:"entity_type,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type QuotaParameters struct {

	// A map of string k/v properties.
	// +kubebuilder:validation:Optional
	Config map[string]float64 `json:"config,omitempty" tf:"config,omitempty"`

	// The name of the entity
	// +kubebuilder:validation:Required
	EntityName *string `json:"entityName" tf:"entity_name,omitempty"`

	// The type of the entity (client-id, user, ip)
	// +kubebuilder:validation:Required
	EntityType *string `json:"entityType" tf:"entity_type,omitempty"`
}

// QuotaSpec defines the desired state of Quota
type QuotaSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     QuotaParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider QuotaInitParameters `json:"initProvider,omitempty"`
}

// QuotaStatus defines the observed state of Quota.
type QuotaStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        QuotaObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Quota is the Schema for the Quotas API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,kafka}
type Quota struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              QuotaSpec   `json:"spec"`
	Status            QuotaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QuotaList contains a list of Quotas
type QuotaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Quota `json:"items"`
}

// Repository type metadata.
var (
	Quota_Kind             = "Quota"
	Quota_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Quota_Kind}.String()
	Quota_KindAPIVersion   = Quota_Kind + "." + CRDGroupVersion.String()
	Quota_GroupVersionKind = CRDGroupVersion.WithKind(Quota_Kind)
)

func init() {
	SchemeBuilder.Register(&Quota{}, &QuotaList{})
}
