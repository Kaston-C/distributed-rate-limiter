# 🚦 Distributed Rate Limiter

A scalable and observable rate limiter built in **Go**, using **Redis** for centralized tracking, **Docker** for containerization, **Kubernetes** for orchestration, and **Prometheus** for metrics.
> Handles traffic limits per IP with Prometheus-powered monitoring.

---

## 🚀 Features

- ⚡ High-performance HTTP rate limiting
- 📊 Real-time monitoring via Prometheus
- 🐳 Containerized with Docker & Docker Compose
- ☸️ Kubernetes-ready for production environments
- 🔁 Configurable request limits and intervals

---

## 💡 Why I Built This

I built this project as a learning exercise to get hands-on experience with **Redis**, **Docker**, **Kubernetes**, and **Prometheus**. While implementing a rate limiter provided a practical example to understand the core concepts, the primary goal was to dive into containerization, orchestration, and monitoring.

---

## 🧱 Architecture Overview
```
+--------+          +--------------+          +-------+
| Client |  ----->  | Rate Limiter |  ----->  | Redis |
+--------+          +--------------+          +-------+
                           |
                           v
                     +------------+
                     | Prometheus |
                     +------------+
```

---

## ⚙️ Getting Started

### Prerequisites

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- [Docker Desktop with Kubernetes](https://docs.docker.com/desktop/kubernetes/)

---

### 🐳 Run Locally with Docker Compose

```bash
docker-compose up --build
````

Once running:

* **Rate limiter endpoint**: [http://localhost:8080](http://localhost:8080)
* **Prometheus metrics**: [http://localhost:9090](http://localhost:9090)

---

### ☸️ Run with Kubernetes

1. **(Optional) Build the image**
   If you haven’t already built the Docker image using `docker-compose up --build`, run:

```bash
docker build -t rate-limiter .
```

2. **Tag the image for Kubernetes**
   Tag the image so Kubernetes can find it locally:

```bash
docker tag rate-limiter-rate-limiter:latest rate-limiter:latest
```

3. **Apply Kubernetes resources:**

```bash
kubectl apply -f k8s/redis-deployment.yaml
kubectl apply -f k8s/rate-limiter-deployment.yaml
kubectl apply -f k8s/prometheus-config.yaml
kubectl apply -f k8s/prometheus-deployment.yaml
```

---

### 📈 View Metrics in Prometheus

* Visit: [http://localhost:9090](http://localhost:9090)
* Try querying:

  * `total_requests`
  * `rate_limited_requests`

---

## 📂 Project Structure

```
.
├── internal/
│   ├── handlers/      # HTTP route handlers and middleware
│   ├── metrics/       # Prometheus metrics setup
│   ├── ratelimiter    # Rate limiter logic (Go)
├── k8s/               # Kubernetes manifests
├── Dockerfile
├── docker-compose.yaml
├── main.go
└── go.mod
```

---

## 🧪 Example Request (cURL)

```bash
curl http://localhost:8080/
```

To simulate multiple requests:

```bash
for i in {1..20}; do curl -s -o /dev/null -w "Request $i: %{http_code}\n" http://localhost:8080 done
```

---

## 🔧 Potential Enhancements

### 🔄 More Precise Rate Limiting Algorithms

* **Sliding Window Log**: Tracks individual request timestamps for fine-grained control, but is memory-intensive.
* **Sliding Window Counter**: Uses interpolated counters for efficient sliding window approximation.
* **Token Bucket / Leaky Bucket**: Smoother enforcement and better handling of bursts than fixed windows.

### 📈 Dynamic Limits Per User/IP

* Integrate role- or subscription-based rate limits.
* Fetch rate limit rules from a config service or database.

### 📊 Improved Observability

* Add more Prometheus metrics (e.g., per-endpoint stats).
* Integrate **Grafana** for visualization.

### 🔒 Security Enhancements

* Add IP validation or authentication.
* Rate-limit based on API keys or user IDs.

---

## 📋 License

MIT License — see [LICENSE](LICENSE)

---

## 🙌 Acknowledgments

* [Go Redis Client](https://github.com/redis/go-redis)
* [Prometheus Go Client](https://github.com/prometheus/client_golang)
* [Docker & Kubernetes](https://www.docker.com/)