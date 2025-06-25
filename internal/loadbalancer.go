package internal

import (
	"fmt"
	"net/http"
	"net/url"

	"go.uber.org/zap"
)

type LoadBalancer struct {
	logger    *zap.Logger
	backends  []Backend
	nextIndex int
}

// TODO: cover with tests
func (l *LoadBalancer) Start() error {
	l.logger.Info("Starting load balancer...")
	listenAddr := ":9090" //TODO: make this configurable
	rrLb := NewRoundRobinReverseProxy(l.backends, l.logger)

	l.logger.Info("Load balancer started", zap.String("address", listenAddr))

	if err := http.ListenAndServe(listenAddr, rrLb.WithProxy()); err != nil {
		return fmt.Errorf("failed to ListenAndServe the load balancer: %w", err)
	}

	return nil
}

func NewLoadBalancer(filePath string, log *zap.Logger) (LoadBalancer, error) {
	cfg, err := ReadConfig(filePath)
	if err != nil {
		return LoadBalancer{}, fmt.Errorf("failed to read configuration: %w", err)
	}

	bes, err := mapToBackends(cfg.LoadBalancerConfig.Backends)
	if err != nil {
		return LoadBalancer{}, fmt.Errorf("failed to map backends: %w", err)
	}

	return LoadBalancer{
		logger:    log,
		backends:  bes,
		nextIndex: 0,
	}, nil
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
