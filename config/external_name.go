/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"fmt"
	"github.com/crossplane/upjet/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Don't rely on metadata.name as the name of the topic, because managed resources are cluster-scoped
	// so MRs for topics on multiple clusters could have a name collision.
	"kafka_topic": config.TemplatedStringAsIdentifier("", "{{ .parameters.name }}"),
	"kafka_acl":   config.TemplatedStringAsIdentifier("", "{{ .parameters.acl_principal }}|{{ .parameters.acl_host }}|{{ .parameters.acl_operation }}|{{ .parameters.acl_permission_type }}|{{ .parameters.resource_type }}|{{ .parameters.resource_name }}|{{ .parameters.resource_pattern_type_filter }}"),
	// I don't know what the id is. Let's find out.
	"kafka_quota":                 config.IdentifierFromProvider,
	"kafka_user_scram_credential": config.IdentifierFromProvider,
}

func GroupKindOverrides() config.ResourceOption {
	return func(r *config.Resource) {
		if r.Name == "kafka_user_scram_credential" {
			r.ShortGroup = "kafka"
			r.Kind = "UserScramCredential"
		}
	}
}

// NoAsync disables the async-by-default behavior of upjet. Unlike many resources in cloud providers, Kafka resources
// are fast to create. None of them need to be done asynchronously.
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
			fmt.Println(r.Name, r.ShouldUseNoForkClient())
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
