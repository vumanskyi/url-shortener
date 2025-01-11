# 🚀 URL Shortener Deployment on Minikube with Kubernetes

This guide will walk you through deploying your URL Shortener application on Minikube using Kubernetes. It includes setting up Redis as a backend service and ensuring the app runs smoothly.

### 📦 1. Prerequisites

Minikube installed: https://minikube.sigs.k8s.io/docs/start/

Kubectl installed: https://kubernetes.io/docs/tasks/tools/

Docker image of the application pushed to a registry (e.g., Docker Hub).

Ensure Minikube is running:
```
minikube start
```

### 📂 2. Project Structure
```
.
├── redis-deployment.yaml
├── redis-service.yaml
├── url-shortener-deployment.yaml
└── url-shortener-service.yaml
```
### 🛠 3. Apply Kubernetes Resources

Run the following commands to apply the deployments and services:

```
kubectl apply -f redis-deployment.yaml
kubectl apply -f redis-service.yaml
kubectl apply -f url-shortener-deployment.yaml
kubectl apply -f url-shortener-service.yaml
```
Check the status of your pods:
```
kubectl get pods
```
Check the services:
```
kubectl get services
```
### 🌐 4. Access the Application

Use Minikube to expose the url-shortener service:

```
minikube service url-shortener
```
This command will open the application in your default browser. Alternatively, you can check the URL with:

```
minikube service list
```

### 📄 5. Cleaning Up

To delete all resources:
```
kubectl delete -f redis-deployment.yaml
kubectl delete -f redis-service.yaml
kubectl delete -f url-shortener-deployment.yaml
kubectl delete -f url-shortener-service.yaml
```
Or delete all resources in the namespace:

```
kubectl delete all --all
```
