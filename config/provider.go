/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/Mongey/terraform-provider-kafka/kafka"
	ujconfig "github.com/crossplane/upjet/pkg/config"
)

const (
	resourcePrefix = "kafka"
	modulePath     = "github.com/mbbush/provider-kafka-jet"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("jet.crossplane.io"),
		//ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithNoForkIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithTerraformProvider(kafka.Provider()),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
			GroupKindOverrides(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
