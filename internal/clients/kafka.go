/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	terraform2 "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/config"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"

	"github.com/crossplane-contrib/provider-upjet-kafka/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig        = "no providerConfigRef provided"
	errGetProviderConfig       = "cannot get referenced ProviderConfig"
	errTrackUsage              = "cannot track ProviderConfig usage"
	errExtractCredentials      = "cannot extract credentials"
	errUnmarshalCredentials    = "cannot unmarshal kafka credentials as JSON"
	keyTfBootstrapBrokers      = "bootstrap_servers"
	keyTfCaCert                = "ca_cert"
	keyTfClientCert            = "client_cert"
	keyTfClientKey             = "client_key"
	keyTfClientKeyPassphrase   = "client_key_passphrase"
	keyTfSaslUsername          = "sasl_username"
	keyTfSaslPassword          = "sasl_password"
	keyTfSaslMechanism         = "sasl_mechanism"
	keyTfSkipTlsVerify         = "skip_tls_verify"
	keyTfTlsEnabled            = "tls_enabled"
	keyTfTimeout               = "timeout"
	keyXpBootstrapBrokers      = "bootstrapBrokers"
	keyXpBootstrapBrokerString = "bootstrapBrokerString"
	keyXpCaCert                = "caCert"
	keyXpClientCert            = "clientCert"
	keyXpClientKey             = "clientKey"
	keyXpClientKeyPassphrase   = "clientKeyPassphrase"
	keyXpSaslUsername          = "saslUsername"
	keyXpSaslPassword          = "saslPassword"
	keyXpSaslMechanism         = "saslMechanism"
	keyXpSkipTlsVerify         = "skipTlsVerify"
	keyXpTlsEnabled            = "tlsEnabled"
	keyXpTimeout               = "timeout"
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

		// Merge the ProviderConfig and credentials into a configuration map for the Terraform provider.
		tfConfig, err := parseToTfConfig(ctx, creds, pc)
		if err != nil {
			return ps, errors.Wrap(err, "failed to parse credentials to terraform configuration")
		}
		ps.Configuration = tfConfig
		if ujprovider.TerraformProvider != nil {
			diag := ujprovider.TerraformProvider.Configure(ctx, &terraform2.ResourceConfig{
				Config: tfConfig,
			})
			if diag != nil && diag.HasError() {
				return ps, errors.Errorf("failed to configure the provider: %v", diag)
			}
			ps.Meta = ujprovider.TerraformProvider.Meta()
		} else {
			fmt.Println("ERROR: Need to instantiate a terraform provider somehow")
			return ps, errors.Wrap(err, "Unable to set terraform provider setup.Meta")
		}
		return ps, nil
	}
}

func parseToTfConfig(ctx context.Context, creds map[string]any, pc *v1beta1.ProviderConfig) (map[string]any, error) {
	tfConfig := make(map[string]any)

	// brokers
	// First read from creds string, then from creds array, then from provider config. Last one takes precedence.
	sBrokers, ok := creds[keyXpBootstrapBrokerString].(string)
	if ok && sBrokers != "" {
		//fmt.Println("Using bootstrap broker string from credentials")
		tfConfig[keyTfBootstrapBrokers] = strings.Split(sBrokers, ",")
	}
	iCredsBrokers, ok := creds[keyXpBootstrapBrokers].([]interface{})
	if ok && iCredsBrokers != nil {
		//fmt.Println("Using bootstrap brokers from credentials")
		brokers := make([]string, len(iCredsBrokers))
		for idx, ifce := range iCredsBrokers {
			if ifce != nil {
				brokers[idx] = ifce.(string)
			}
		}
		tfConfig[keyTfBootstrapBrokers] = brokers
	}

	if pcBrokerString := pc.Spec.BootstrapBrokerString; pcBrokerString != nil {
		//fmt.Println("Using bootstrap broker string from provider config")
		tfConfig[keyTfBootstrapBrokers] = strings.Split(*pcBrokerString, ",")
	}
	if pcBrokers := pc.Spec.BootstrapBrokers; pcBrokers != nil {
		//fmt.Println("Using bootstrap brokers from provider config")
		tfConfig[keyTfBootstrapBrokers] = pc.Spec.BootstrapBrokers
	}

	// tls config
	// creds
	if cCaCert, ok := creds[keyXpCaCert].(string); ok {
		//fmt.Println("Using ca cert from credentials")
		tfConfig[keyTfCaCert] = cCaCert
	}
	if cClientCert, ok := creds[keyXpClientCert].(string); ok {
		//fmt.Println("Using client cert from credentials")
		tfConfig[keyTfClientCert] = cClientCert
	}
	if cClientKey, ok := creds[keyXpClientKey].(string); ok {
		//fmt.Println("Using client key from credentials")
		tfConfig[keyTfClientKey] = cClientKey
	}
	if cClientKeyPassphrase, ok := creds[keyXpClientKeyPassphrase].(string); ok {
		//fmt.Println("Using client key passphrase from credentials")
		tfConfig[keyTfClientKeyPassphrase] = cClientKeyPassphrase
	}
	if cSkipTlsVerify, ok := creds[keyXpSkipTlsVerify].(bool); ok {
		//fmt.Println("Using skip tls verify from credentials")
		tfConfig[keyTfSkipTlsVerify] = cSkipTlsVerify
	}
	if cTlsEnabled, ok := creds[keyXpTlsEnabled].(bool); ok {
		//fmt.Println("Using tls enabled from credentials")
		tfConfig[keyTfTlsEnabled] = cTlsEnabled
	}
	if pc.Spec.TlsConfig != nil {
		// provider config
		if pc.Spec.TlsConfig.CaCert != nil {
			//fmt.Println("Using ca cert from provider config")
			tfConfig[keyTfCaCert] = *pc.Spec.TlsConfig.CaCert
		}
		if pc.Spec.TlsConfig.ClientCert != nil {
			//fmt.Println("Using client cert from provider config")
			tfConfig[keyTfClientCert] = *pc.Spec.TlsConfig.ClientCert
		}
		if pc.Spec.TlsConfig.SkipTlsVerify != nil {
			//fmt.Println("Using skip tls verify from provider config")
			tfConfig[keyTfSkipTlsVerify] = *pc.Spec.TlsConfig.SkipTlsVerify
		}
		if pc.Spec.TlsConfig.TlsEnabled != nil {
			//fmt.Println("Using tls enabled from provider config")
			tfConfig[keyTfTlsEnabled] = *pc.Spec.TlsConfig.TlsEnabled
		}
	}

	// sasl
	// creds
	if cSaslUsername, ok := creds[keyXpSaslUsername].(string); ok {
		//fmt.Println("Using sasl username from credentials")
		tfConfig[keyTfSaslUsername] = cSaslUsername
	}
	if cSaslPassword, ok := creds[keyXpSaslPassword].(string); ok {
		//fmt.Println("Using sasl password from credentials")
		tfConfig[keyTfSaslPassword] = cSaslPassword
	}
	if cSaslMechanism, ok := creds[keyXpSaslMechanism].(string); ok {
		//fmt.Println("Using sasl mechanism from credentials")
		// terraform provider uses lowercase sasl mechanism, unlike most of the rest of the kafka ecosystem
		tfConfig[keyTfSaslMechanism] = strings.ToLower(cSaslMechanism)
	}

	// provider config
	if pc.Spec.SaslMechanism != nil {
		//fmt.Println("Using sasl mechanism from provider config")
		tfConfig[keyTfSaslMechanism] = strings.ToLower(*pc.Spec.SaslMechanism)
	}

	// timeout
	if cTimeout, ok := creds[keyXpTimeout]; ok {
		//fmt.Println("Using timeout from credentials")
		tfConfig[keyTfTimeout] = cTimeout
	}
	if pc.Spec.Timeout != nil {
		//fmt.Println("Using timeout from provider config")
		tfConfig[keyTfTimeout] = *pc.Spec.Timeout
	}

	if brokers, ok := tfConfig[keyTfBootstrapBrokers].([]string); !ok || len(brokers) == 0 {
		return nil, errors.New("no bootstrap brokers provided")
	}

	return tfConfig, nil
}
