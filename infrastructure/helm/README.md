# URL Shortener Service - Helm Deployment

This guide helps you to deploy the URL Shortener service on Kubernetes using Helm.

## Prerequisites

- [Minikube](https://minikube.sigs.k8s.io/docs/) — to run a local Kubernetes cluster.
- [Helm](https://helm.sh/) — for managing Kubernetes charts.

## Setup and Installation

1. **Start Minikube**:
   If Minikube is not already running, start it with the following command:
   ```
   minikube start
   ```
2. **Install the Helm Chart**: Navigate to the Helm chart directory and install the chart with the following command:

```
helm install url-shortener ./url-shortener -f ./url-shortener/values.yaml
```

3. **Access the Service**: Once the installation is complete, you can access the service through Minikube by running:
```
minikube service url-shortener
```

**Uninstall the Helm Chart**:
```
helm uninstall url-shortener
```