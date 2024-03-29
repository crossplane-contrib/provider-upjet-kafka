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

type TopicInitParameters struct {

	// A map of string k/v attributes.
	Config map[string]string `json:"config,omitempty" tf:"config,omitempty"`

	// Number of partitions.
	Partitions *int64 `json:"partitions,omitempty" tf:"partitions,omitempty"`

	// Number of replicas.
	ReplicationFactor *int64 `json:"replicationFactor,omitempty" tf:"replication_factor,omitempty"`
}

type TopicObservation struct {

	// A map of string k/v attributes.
	Config map[string]string `json:"config,omitempty" tf:"config,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the topic.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Number of partitions.
	Partitions *int64 `json:"partitions,omitempty" tf:"partitions,omitempty"`

	// Number of replicas.
	ReplicationFactor *int64 `json:"replicationFactor,omitempty" tf:"replication_factor,omitempty"`
}

type TopicParameters struct {

	// A map of string k/v attributes.
	// +kubebuilder:validation:Optional
	Config map[string]string `json:"config,omitempty" tf:"config,omitempty"`

	// The name of the topic.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Number of partitions.
	// +kubebuilder:validation:Optional
	Partitions *int64 `json:"partitions,omitempty" tf:"partitions,omitempty"`

	// Number of replicas.
	// +kubebuilder:validation:Optional
	ReplicationFactor *int64 `json:"replicationFactor,omitempty" tf:"replication_factor,omitempty"`
}

// TopicSpec defines the desired state of Topic
type TopicSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TopicParameters `json:"forProvider"`
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
	InitProvider TopicInitParameters `json:"initProvider,omitempty"`
}

// TopicStatus defines the observed state of Topic.
type TopicStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TopicObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Topic is the Schema for the Topics API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,kafka}
type Topic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.partitions) || (has(self.initProvider) && has(self.initProvider.partitions))",message="spec.forProvider.partitions is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.replicationFactor) || (has(self.initProvider) && has(self.initProvider.replicationFactor))",message="spec.forProvider.replicationFactor is a required parameter"
	Spec   TopicSpec   `json:"spec"`
	Status TopicStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TopicList contains a list of Topics
type TopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Topic `json:"items"`
}

// Repository type metadata.
var (
	Topic_Kind             = "Topic"
	Topic_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Topic_Kind}.String()
	Topic_KindAPIVersion   = Topic_Kind + "." + CRDGroupVersion.String()
	Topic_GroupVersionKind = CRDGroupVersion.WithKind(Topic_Kind)
)

func init() {
	SchemeBuilder.Register(&Topic{}, &TopicList{})
}
