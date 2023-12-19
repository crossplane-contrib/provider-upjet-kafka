#!/usr/bin/env bash
set -aeuo pipefail

echo "Running setup.sh"
echo "Installing kafka in the cluster"
helm upgrade --install --wait kafka oci://registry-1.docker.io/bitnamicharts/kafka -n kafka --create-namespace

echo "Creating cloud credential secret..."
KAFKA_PASSWORD=$(kubectl get secret -n kafka kafka-user-passwords -o jsonpath='{.data.client-passwords}' | base64 -d | cut -d , -f 1)
PROVIDER_CONFIG=$(cat <<EOL
{
	"bootstrap_servers": [
		"kafka-controller-0.kafka-controller-headless.kafka.svc.cluster.local:9092",
		"kafka-controller-1.kafka-controller-headless.kafka.svc.cluster.local:9092",
		"kafka-controller-2.kafka-controller-headless.kafka.svc.cluster.local:9092"
	],
	"sasl_username": "user1",
	"sasl_password": "$KAFKA_PASSWORD",
	"sasl_mechanism": "scram-sha256",
	"tls_enabled": false
}
EOL
)
${KUBECTL} -n upbound-system create secret generic provider-secret --from-literal=config="$PROVIDER_CONFIG" --dry-run=client -o yaml | ${KUBECTL} apply -f -

echo "Waiting until provider is healthy..."
${KUBECTL} wait provider.pkg --all --for condition=Healthy --timeout 5m

echo "Waiting for all pods to come online..."
${KUBECTL} -n upbound-system wait --for=condition=Available deployment --all --timeout=5m

echo "Creating a default provider config..."
cat <<EOF | ${KUBECTL} apply -f -
apiVersion: kafka.jet.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: provider-secret
      namespace: upbound-system
      key: config
