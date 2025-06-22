package internal

import (
	"fmt"
	"net/url"

	"github.com/odrianoaliveira/loadbalancer/pkg/config"
	"go.uber.org/zap"
)

type LoadBalancer struct {
	logger    *zap.Logger
	backends  []Backend
	nextIndex int
}

func (l *LoadBalancer) Start() {
	panic("unimplemented")
}

func NewLoadBalancer(filePath string, log *zap.Logger) (*LoadBalancer, error) {
	cfg, err := config.ReadConfig(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read configuration: %w", err)
	}

	if len(cfg.LoadBalancerConfig.Backends) == 0 {
		return nil, fmt.Errorf("no backends configured in the load balancer configuration")
	}

	if bes, err := mapToBackends(cfg.LoadBalancerConfig.Backends); err == nil {
		return &LoadBalancer{
			logger:    log,
			backends:  bes,
			nextIndex: 0,
		}, nil
	} else {
		return nil, fmt.Errorf("failed to map backends: %w", err)
	}

}

func mapToBackends(backend []config.Backend) ([]Backend, error) {
	var backends []Backend
	for _, b := range backend {
		parsedURL, err := url.Parse(b.URL)
		if err != nil {
			return nil, fmt.Errorf("invalid URL in configuration: %s", b.URL)
		}
		backends = append(backends, Backend{
			URL:         parsedURL,
			IsAlive:     true,
			connections: 0,
		})
	}

	return backends, nil
}
