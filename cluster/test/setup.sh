#!/usr/bin/env bash
set -aeuo pipefail

echo "Running setup.sh"
pwd
echo "Installing kafka in the cluster"
helm upgrade kafka /home/matt/git/provider-upjet-kafka/cluster/test/local-kafka-dev -n kafka --create-namespace --wait --install

echo "Setting kafka admin user password"

kubectl exec -n kafka svc/kafka -- kafka-configs --zookeeper zookeeper:2181 --alter --add-config "SCRAM-SHA-256=[password=kafka]" --entity-type users --entity-name kafka
kubectl exec -n kafka svc/kafka -- kafka-configs --zookeeper zookeeper:2181 --alter --add-config "SCRAM-SHA-256=[password=client-secret]" --entity-type users --entity-name client

echo "Creating cloud credential secret..."
PROVIDER_CONFIG=$(cat <<EOL
{
	"bootstrapBrokers": [
		"kafka.kafka.svc.cluster.local:9092"
	],
	"saslUsername": "client",
	"saslPassword": "client-secret",
	"saslMechanism": "SCRAM-SHA256",
	"tlsEnabled": false
}
EOL
)
kubectl -n upbound-system create secret generic provider-secret --from-literal=config="$PROVIDER_CONFIG" --dry-run=client -o yaml | kubectl apply -f -

echo "Waiting until provider is healthy..."
kubectl wait provider.pkg --all --for condition=Healthy --timeout 5m

echo "Waiting for all pods to come online..."
kubectl -n upbound-system wait --for=condition=Available deployment --all --timeout=5m

echo "Creating a default provider config..."
cat <<EOF | kubectl apply -f -
apiVersion: kafka.upjet.crossplane.io/v1beta1
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
EOF
