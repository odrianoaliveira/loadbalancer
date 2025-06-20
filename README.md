
# Load Balancer
A lightweight and concurrent load balancer written in Go, designed to distribute HTTP requests across multiple downstream services. This project implements core load balancing strategies and supports configuration via a simple file-based approach.

<img src="https://golang.org/doc/gopher/frontpage.png" alt="Gopher" width="80"/>

## ✨ Features

* [ ] Load distribution across downstream services
* [ ] Round Robin strategy
* [ ] Least Connections strategy
* [ ] Sticky Sessions based on configurable key (e.g., user ID, IP address)
* [ ] Configurable downstream members via a configuration file
* [ ] Health checking for downstream services (planned)
* [ ] Prometheus metrics

## 🌎 Load Balancing Strategies

### Round Robin

Evenly distributes requests to downstreams in a circular fashion.

### Least Connections

Routes traffic to the downstream instance with the fewest active connections.

### Sticky Sessions

Routes requests based on a sticky key (e.g., cookie, header, query parameter). The same client consistently reaches the same downstream instance.

## 📂 Configuration

The service is configured via a YAML or JSON file listing downstream service URLs and optionally their weights, sticky session parameters, and health check settings.

Example (YAML):

```yaml
loadbalancer:
  strategy: round_robin  # Options: round_robin, least_connections, sticky
  sticky_key: X-Session-ID
  backends:
    - url: http://localhost:8081
    - url: http://localhost:8082
```

## ⚙️ Getting Started

```bash
git clone https://github.com/yourorg/loadbalancer.git
cd loadbalancer
make build
./loadbalancer -config ./configs/config.yaml
```

## 🛠️ Development Best Practices

* Follows Go standard project layout
* Linting with `golangci-lint`
* Race detection with `go test -race`
* Modular packages for extensibility
* CI/CD integration via GitHub Actions

## 📆 Roadmap

* [ ] Backend health checks
* [ ] Timeout and retry mechanisms
* [ ] Metrics & observability (Prometheus, OpenTelemetry)
* [ ] Dynamic configuration reload
