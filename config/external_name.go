/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"strings"

	"github.com/crossplane/upjet/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider. Every one of them is based entirely on the parameters, so even when observing existing resources,
// you should never need to set the external-name annotation manually.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Don't rely on metadata.name as the name of the topic, because managed resources are cluster-scoped
	// so MRs for topics on multiple clusters could have a name collision.
	"kafka_topic":                 PipeSeparatedParametersAsIdentifier("name"),
	"kafka_acl":                   PipeSeparatedParametersAsIdentifier("acl_principal", "acl_host", "acl_operation", "acl_permission_type", "resource_type", "resource_name", "resource_pattern_type_filter"),
	"kafka_quota":                 PipeSeparatedParametersAsIdentifier("entity_name", "entity_type"),
	"kafka_user_scram_credential": PipeSeparatedParametersAsIdentifier("username", "scram_mechanism"),
}

// PipeSeparatedParametersAsIdentifier configures the external-name config for each resource in this provider to use a
// value for both its terraform id and external name that consists of the relevant kafka query api parameters, separated
// by pipe characters (|).
func PipeSeparatedParametersAsIdentifier(params ...string) config.ExternalName {
	template := "{{ .parameters." + strings.Join(params, " }}|{{ .parameters.") + " }}"
	en := config.TemplatedStringAsIdentifier("", template)
	en.DisableNameInitializer = true
	return en
}

// GroupKindOverrides keeps all four resources in this small provider in a single group.
func GroupKindOverrides() config.ResourceOption {
	return func(r *config.Resource) {
		if r.Name == "kafka_user_scram_credential" {
			r.ShortGroup = "kafka"
			r.Kind = "UserScramCredential"
		}
	}
}

// NoAsync disables the async-by-default behavior of upjet. Unlike many resources in cloud providers, Kafka resources
// are fast to create. None of them need to be done asynchronously. However, it seems that upjet has much better
// error handling when running in async mode, so we may to be async just for that reason.
func NoAsync() config.ResourceOption {
	return func(r *config.Resource) {
		r.UseAsync = false
	}
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
