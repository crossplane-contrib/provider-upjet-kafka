/*
Copyright 2022 Upbound Inc.
*/

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// A ProviderConfigSpec defines the desired state of a ProviderConfig.
type ProviderConfigSpec struct {
	// An array of host:port strings for the kafka brokers to connect to. E.g. ["broker-0:9092", "broker-1:9092"]
	// If both are defined, this takes precedence over BootstrapBrokerString.
	// +optional
	BootstrapBrokers *[]string `json:"bootstrapBrokers,omitempty"`
	// A comma-separated string of host:port strings for the kafka brokers to connect to.
	// E.g. "broker-0:9092,broker-1:9092"
	// +optional
	BootstrapBrokerString *string `json:"bootstrapBrokerString,omitempty"`
	// Credentials required to authenticate to this provider.
	// Non-sensitive configuration parameters may either be specified in the credentials secret or in the ProviderConfig
	// See the project's readme on github for schema details.
	Credentials *ProviderCredentials `json:"credentials"`

	// The SASL mechanism to use to authenticate with kafka. Supported values are PLAIN, SCRAM-SHA256, SCRAM-SHA512.
	// Leave unset if not using SASL.
	// +optional
	// +kubebuilder:validation:Enum=PLAIN;SCRAM-SHA-256;SCRAM-SHA-512
	SaslMechanism *string `json:"saslMechanism,omitempty"`
	// Timeout in seconds to configure on the underlying kafka client.
	// +optional
	Timeout *int `json:"timeout,omitempty"`

	// +optional
	TlsConfig *TlsConfig `json:"tls,omitempty"`
}

type TlsConfig struct {
	// CA Certificate to validate the server's certificate
	// +optional
	CaCert *string `json:"caCert,omitempty"`
	// Client certificate
	// +optional
	ClientCert *string `json:"clientCert,omitempty"`
	// Enable communications with the target kafka cluster over TLS
	// +optional
	TlsEnabled *bool `json:"tlsEnabled,omitempty"`
	// Disable basic TLS verification. This should only be true if the kafka server is an insecure development instance.
	// +optional
	SkipTlsVerify *bool `json:"skipTlsVerify,omitempty"`
}

// ProviderCredentials required to authenticate.
type ProviderCredentials struct {
	// Source of the provider credentials.
	// +kubebuilder:validation:Enum=None;Secret;InjectedIdentity;Environment;Filesystem
	Source xpv1.CredentialsSource `json:"source"`

	xpv1.CommonCredentialSelectors `json:",inline"`
}

// A ProviderConfigStatus reflects the observed state of a ProviderConfig.
type ProviderConfigStatus struct {
	xpv1.ProviderConfigStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// A ProviderConfig configures a Kafka provider.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="SECRET-NAME",type="string",JSONPath=".spec.credentials.secretRef.name",priority=1
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:resource:scope=Cluster,categories={crossplane,provider,kafka}
type ProviderConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProviderConfigSpec   `json:"spec"`
	Status ProviderConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProviderConfigList contains a list of ProviderConfig.
type ProviderConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProviderConfig `json:"items"`
}

// +kubebuilder:object:root=true

// A ProviderConfigUsage indicates that a resource is using a ProviderConfig.
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="CONFIG-NAME",type="string",JSONPath=".providerConfigRef.name"
// +kubebuilder:printcolumn:name="RESOURCE-KIND",type="string",JSONPath=".resourceRef.kind"
// +kubebuilder:printcolumn:name="RESOURCE-NAME",type="string",JSONPath=".resourceRef.name"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,provider,kafka}
type ProviderConfigUsage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	xpv1.ProviderConfigUsage `json:",inline"`
}

// +kubebuilder:object:root=true

// ProviderConfigUsageList contains a list of ProviderConfigUsage
type ProviderConfigUsageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProviderConfigUsage `json:"items"`
}
