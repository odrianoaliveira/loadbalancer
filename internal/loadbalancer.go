package internal

import (
	"fmt"
	"log/slog"
	"net/url"
)

type LBStrategy string

const (
	StrategyRoundRobin LBStrategy = "round-robin"
	StrategyLeastConn  LBStrategy = "least-connections"
)

type baseLoadBalancer struct {
	port     int
	strategy LBStrategy
	backends []Backend
}

type LoadBalancer interface {
	Start() error
}

func NewLoadBalancer(filePath string) (LoadBalancer, error) {
	cfg, err := ReadConfig(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read configuration: %w", err)
	}

	bes, err := mapToBackends(cfg.LoadBalancerConfig.Backends)
	if err != nil {
		return nil, fmt.Errorf("failed to map backends: %w", err)
	}

	switch cfg.LoadBalancerConfig.Strategy {
	case StrategyRoundRobin:
		slog.Info("Using Round Robin strategy for load balancing")
		lb := baseLoadBalancer{
			port:     cfg.LoadBalancerConfig.Port,
			strategy: StrategyRoundRobin,
			backends: bes,
		}
		return NewRoundRobinLoadBalancer(lb), nil
	case StrategyLeastConn:
		slog.Info("Using Least Connections strategy for load balancing")
		lb := baseLoadBalancer{
			port:     cfg.LoadBalancerConfig.Port,
			strategy: StrategyLeastConn,
			backends: bes,
		}
		return NewLeastConnectionsLoadBalancer(lb), nil
	default:
		return nil, fmt.Errorf("unsupported load balancing strategy: %s", cfg.LoadBalancerConfig.Strategy)
	}

}

func mapToBackends(backend []BackendConfig) ([]Backend, error) {
	var backends []Backend
	for _, b := range backend {
		parsedURL, err := url.Parse(b.URL)
		if err != nil {
			return nil, fmt.Errorf("invalid URL in configuration: %s", b.URL)
		}

		backends = append(backends, Backend{
			URL:         *parsedURL,
			IsAlive:     true,
			connections: 0,
		})
	}

	return backends, nil
}
