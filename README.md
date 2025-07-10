
# 🌐 L7 Load Balancer
A lightweight and concurrent L7 load balancer written in Go, designed to distribute HTTP requests across multiple downstream services. This project implements core load balancing strategies and supports configuration via a simple file-based approach.

<img src="https://golang.org/doc/gopher/frontpage.png" alt="Gopher" width="80"/>

## Features

* [x] Load distribution across downstream services
* [x] Round Robin strategy
* [x] Configurable downstream members via a configuration file
* [ ] Least Connections strategy
* [ ] Sticky Sessions based on configurable key (e.g., user ID, IP address)
* [ ] Health checking for downstream services
* [ ] Detect backend failure and redirect traffic seamlessly
* [ ] Metrics/Observability

## Load Balancing Strategies

### Round Robin

It even distributes requests downstream in a circular fashion.

### Least Connections

Routes traffic to the downstream instance with the fewest active connections.

### Sticky Sessions

Routes requests based on a sticky key (e.g., cookie, header, query parameter). The same client consistently reaches the same downstream instance.

## Configuration

The service is configured via a YAML or JSON file listing downstream service URLs and optionally their weights, sticky session parameters, and health check settings.

Example (YAML):

```yaml
loadbalancer:
  strategy: round_robin  # Options: round_robin, least_connections, sticky
  port: 9090
  sticky_key: X-Session-ID
  backends:
    - url: http://localhost:8081
    - url: http://localhost:8082
```

## Getting Started

```bash
git clone https://github.com/yourorg/loadbalancer.git
cd loadbalancer
make build
make run
```

## Development Best Practices

* Follows Go standard project layout
* Linting with `golangci-lint`
* Race detection with `go test -race`
* CI/CD integration via GitHub Actions
