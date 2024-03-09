# Provider Kafka

`provider-upjet-kafka` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
Kafka API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/crossplane-contrib/provider-upjet-kafka):
```
up ctp provider install crossplane-contrib/provider-upjet-kafka:v0.1.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-upjet-kafka
spec:
  package: crossplane-contrib/provider-upjet-kafka:v0.1.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/crossplane-contrib/provider-upjet-kafka).

## Configuration
Non-sensitive values can either be set in the `ProviderConfig`, or from a credential source such as a kubernetes secret.
Sensitive values (private keys, usernames and passwords) must come from a credentials source. If defined in both, the `ProviderConfig` takes precedence.

The format of a kubernetes secret used as a credentials source should be a JSON object, in a single key. Specify that key in the ProviderConfig.

To aid platform developers working with limited tooling, the kafka brokers to connect to may be specified either as an array or a string.
Only one is necessary. If both are specified, the array will take precedence over the string.

In both the `ProviderConfig` and the credentials, unused configuration options should be left null/unset, as setting them to empty values may overwrite defaults defined elsewhere.

Supported SASL mechanisms are PLAIN, SCRAM-SHA256, and SCRAM-SHA512

Example credentials json with all options filled:

```json
{
  bootstrapBrokers: [
    "broker1:9092",
    "broker2:9092"
  ],
  "bootstrapBrokerString": "broker1:9092,broker2:9092",
  "caCert": "string that is passed directly to the underlying terraform provider",
  "clientCert": "string that is passed directly to the underlying terraform provider",
  "clientKey": "string that is passed directly to the underlying terraform provider",
  "clientKeyPassphrase": "string that is passed directly to the underlying terraform provider",
  "saslUsername": "crossplane",
  "saslPassword": "BetterBeMoreSecureThanThisOne",
  "saslMechanism": "SCRAM-SHA256",
  "skipTlsVerify": false,
  "tlsEnabled": true,
  "timeout": 10
}
```

## Developing

Run code-generation pipeline:
```console
make generate
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build, deploy locally in kind, and run tests
```console
make e2e UPTEST_EXAMPLE_LIST=examples/topic/topic.yaml
```

Build binary:

```console
make build
```

## What's different from the existing [provider-kafka](https://github.com/crossplane-contrib/provider-kafka)?
This provider uses upjet with the no-fork architecture to build on the existing [Mongey/terraform-provider-kafka](https://github.com/Mongey/terraform-provider-kafka), adding support for crossplane management policies, kafka quotas, and kafka user scram credentials.

The kafka library underlying the terraform provider is [sarama](https://github.com/IBM/sarama). This library does not currently support AWS MSK SASL-IAM auth. 

## How stable/well-tested is this?

The topic and quota resources have a working automated test pipeline. ACL and UserScramCredential resources do not yet, and are provided on an as-is basis. They _should_ work just as well as they do in terraform, but I have not confirmed that.

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/crossplane-contrib/provider-upjet-kafka/issues).

I would welcome assistance and collaboration.