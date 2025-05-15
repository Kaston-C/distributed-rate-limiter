# 🚦 Distributed Rate Limiter

A scalable and observable rate limiter built in **Go**, using **Redis** for centralized tracking, **Docker** for containerization, **Kubernetes** for orchestration, and **Prometheus** for metrics.
> Handles traffic limits per IP with Prometheus-powered monitoring.

---

## 🚀 Features

- ⚡ High-performance HTTP rate limiting
- 🐳 Containerized with Docker & Docker Compose
- ☸️ Kubernetes-ready for production environments
- 📊 Real-time monitoring via Prometheus
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

The rate limiter is **distributed** by using **Redis as a centralized store** for request counts. Multiple instances or pods share this Redis backend and use atomic operations to update counters consistently. This ensures a **global rate limit** enforced across all running instances, whether on Docker ports or Kubernetes pods.


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

* **Rate limiter endpoints**: [http://localhost:8081](http://localhost:8080) and [http://localhost:8082](http://localhost:8080)

---

### ☸️ Run with Kubernetes

1. **Build the image:**

```bash
docker build -t rate-limiter .
```

2. **Apply Kubernetes resources:**

```bash
kubectl apply -f k8s/redis-deployment.yaml
kubectl apply -f k8s/rate-limiter-deployment.yaml
kubectl apply -f k8s/prometheus-config.yaml
kubectl apply -f k8s/prometheus-deployment.yaml
```

Once running:

* **Rate limiter endpoint**: [http://localhost:8080](http://localhost:8080)

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
├── prometheus.yml
├── main.go
└── go.mod
```

---

## 🧪 Example Request (cURL)

Simulate multiple requests to test the rate-limiting behavior. You should see a mix of `200 OK` and `429 Too Many Requests` responses once the limit is exceeded.


---

### 🐳 Docker:

```bash
curl http://localhost:8081/

curl http://localhost:8082/
```

To simulate multiple requests:

```bash
for i in {1..20}; do curl -s -o /dev/null -w "Request $i: %{http_code}\n" http://localhost:8081; done
```

```bash
for i in {1..20}; do curl -s -o /dev/null -w "Request $i: %{http_code}\n" http://localhost:8082; done
```

---

### ☸️ Kubernetes:

```bash
curl http://localhost:8080/
```

To simulate multiple requests:

```bash
for i in {1..20}; do curl -s -o /dev/null -w "Request $i: %{http_code}\n" http://localhost:8080; done
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