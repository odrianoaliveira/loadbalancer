package internal

import (
	"fmt"
	"net/url"

	"github.com/odrianoaliveira/loadbalancer/pkg/config"
)

type LoadBalancer struct {
	backends  []Backend
	nextIndex int
}

func (l *LoadBalancer) Start() {
	panic("unimplemented")
}

func NewLoadBalancer(filePath string) *LoadBalancer {
	cfg, err := config.ReadConfig(filePath)
	if err != nil {
		panic("Failed to read configuration: " + err.Error())
	}

	if len(cfg.LoadBalancerConfig.Backends) == 0 {
		panic("No backends defined in configuration")
	}

	if bes, err := mapToBackends(cfg.LoadBalancerConfig.Backends); err == nil {
		return &LoadBalancer{
			backends:  bes,
			nextIndex: 0,
		}
	} else {
		panic("Failed to map backends: " + err.Error())
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
