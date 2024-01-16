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

The topic resource has a working automated test pipeline. The other three resources do not yet, and are provided on an as-is basis. They _should_ work just as well as they do in terraform, but I have not confirmed that.

The configuration interface is likely to undergo breaking changes to make it more convenient to use from inside crossplane. Currently it is an exact copy of the terraform provider's config.

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/crossplane-contrib/provider-upjet-kafka/issues).

I would welcome assistance and collaboration.