# grpc-health

This repo contains a sample app exposing a health endpoint by implementing grpc_health_v1. Usecase is to demo the gRPC readiness and liveness probes introduced in Kubernetes 1.23.

```bash
cat <<EOF | kubectl apply -f -
apiVersion: apps/v1
kind: Deployment
metadata:
  creationTimestamp: null
  labels:
    app: grpc-health
  name: grpc-health
spec:
  replicas: 1
  selector:
    matchLabels:
      app: grpc-health
  strategy: {}
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: grpc-health
    spec:
      containers:
      - image: ghcr.io/nmeisenzahl/grpc-health/grpc-health:latest
        imagePullPolicy: Always
        name: grpc-health
        ports:
        - containerPort: 5000
        resources: {}
        livenessProbe:
            grpc:
                port: 5000
            initialDelaySeconds: 10
        readinessProbe:
            grpc:
                port: 5000
            initialDelaySeconds: 10
status: {}
EOF
```
