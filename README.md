# URL Shortener API 🚀

A simple and scalable URL Shortener API built with **Golang** and deployed using **Docker** and **Kubernetes**. This project showcases modern **cloud-native development** practices and includes automated **CI/CD pipelines** using **GitHub Actions**.

## 🛠️ Tech Stack

- **Golang** - The core programming language for the API.
- **Redis** - In-memory data store for caching.
- **Docker** - Containerization for easy deployment.
- **Kubernetes** - Container orchestration for scaling and management.
- **GitHub Actions** - CI/CD for automated Docker builds, pushes, and deployments.
- **Minikube** - Local Kubernetes for testing deployments.

## ⚙️ Features

- Shorten long URLs into easily shareable, custom URLs.
- Custom alias for shortened URLs.
- Support for redirection to original URLs.
- Cache for frequently accessed URLs using Redis.
- Scalable with Kubernetes for easy deployment.

## 🧪 Getting Started

### Prerequisites

- **Docker**: For containerization.
- **Kubernetes (Minikube)**: For local Kubernetes setup.
- **Helm**: For Kubernetes deployments.
- **Redis**: Running locally or in your Kubernetes cluster.

### 🐋 Running Locally with Docker Compose

To run the project locally using Docker Compose, follow these steps:

1. Clone this repository:

   ```bash
   git clone https://github.com/vladumanskyi/url-shortener-api.git
   cd url-shortener-api
   ```
2. Build and run the containers:
    ```bash
    docker-compose up --build
    ```

3. For kubernetes details you can find it [here](https://github.com/vumanskyi/url-shortener/blob/main/infrastructure/k8s/README.md)
