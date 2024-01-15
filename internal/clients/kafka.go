/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Mongey/terraform-provider-kafka/kafka"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/config"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"strings"

	"github.com/crossplane/upjet/pkg/terraform"

	"github.com/mbbush/provider-kafka-jet/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal kafka credentials as JSON"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string, ujprovider *config.Provider) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}
		fmt.Println("running terraform setup builder")

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]any{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		fmt.Println(creds)
		//configKeys := []string{
		//	"ca_cert",
		//	"tls_enabled",
		//	"client_cert",
		//	"bootstrap_servers",
		//	"client_key",
		//	"client_key_passphrase",
		//	"skip_tls_verify",
		//	"sasl_username",
		//	"sasl_password",
		//	"sasl_mechanism",
		//}

		// TODO: See if there's somewhere in here where provider-aws configures the client.
		// It looks like the terraform.Setup in type noForkExternal is unset. How to set it?

		// Set credentials in Terraform provider configuration.
		// This Crossplane provider config schema exactly matches the schema of the terraform provider config.
		// See https://github.com/Mongey/terraform-provider-kafka#provider-configuration

		// I think this is only used for the cli provider?
		ps.Configuration = creds

		// It seems like there should be a way to just call the ConfigureFunc defined on the provider, but I couldn't figure it out.
		ps.Meta, err = providerConfigure(creds)
		if err != nil {
			return ps, errors.Wrap(err, "Unable to set terraform provider setup.Meta")
		}
		return ps, nil
	}
}
func providerConfigure(config map[string]any) (interface{}, error) {
	saslMechanism, ok := config["sasl_mechanism"].(string)
	if !ok {
		saslMechanism = "plain"
	}
	switch saslMechanism {
	case "scram-sha512", "scram-sha256", "plain":
	default:
		return nil, fmt.Errorf("[ERROR] Invalid sasl mechanism \"%s\": can only be \"scram-sha256\", \"scram-sha512\" or \"plain\"", saslMechanism)
	}
	caCert, ok := config["ca_cert"].(string)
	if !ok {
		caCert = ""
	}
	clientCert, ok := config["client_cert"].(string)
	if !ok {
		clientCert = ""
	}
	clientKey, ok := config["client_key"].(string)
	if !ok {
		clientKey = ""
	}
	clientKeyPassphrase, ok := config["client_key_passphrase"].(string)
	if !ok {
		clientKeyPassphrase = ""
	}
	var brokers []string
	iBrokers, ok := config["bootstrap_servers"].([]interface{})
	if !ok || iBrokers == nil || len(iBrokers) == 0 {
		sBrokers, ok := config["bootstrap_broker_string"].(string)
		if !ok || sBrokers == "" {
			return nil, errors.New("Can't deserialize bootstrap_servers")
		}
		brokers = strings.Split(sBrokers, ",")
	} else {
		brokers = make([]string, len(iBrokers))
		for idx, ifce := range iBrokers {
			if ifce != nil {
				brokers[idx] = ifce.(string)
			}
		}
	}

	saslUsername, ok := config["sasl_username"].(string)
	if !ok {
		saslUsername = ""
	}
	saslPassword, ok := config["sasl_password"].(string)
	if !ok {
		saslPassword = ""
	}
	skipTlsVerify, ok := config["skip_tls_verify"].(bool)
	if !ok {
		skipTlsVerify = false
	}
	tlsEnabled, ok := config["tls_enabled"].(bool)
	if !ok {
		tlsEnabled = true
	}
	timeout, ok := config["timeout"].(int)
	if !ok {
		timeout = 3 // 3 seconds is the default in the underlying sarama library's Admin client
	}

	tfProviderConfig := &kafka.Config{
		BootstrapServers:        &brokers,
		CACert:                  caCert,
		ClientCert:              clientCert,
		ClientCertKey:           clientKey,
		ClientCertKeyPassphrase: clientKeyPassphrase,
		SkipTLSVerify:           skipTlsVerify,
		SASLUsername:            saslUsername,
		SASLPassword:            saslPassword,
		SASLMechanism:           saslMechanism,
		TLSEnabled:              tlsEnabled,
		Timeout:                 timeout,
	}

	return &kafka.LazyClient{
		Config: tfProviderConfig,
	}, nil
}
