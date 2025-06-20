package internal

import "github.com/odrianoaliveira/mini-loadbalancer/pkg/config"

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

	return &LoadBalancer{
		backends:  mapToBackends(cfg.LoadBalancerConfig.Backends),
		nextIndex: 0,
	}
}

func mapToBackends(backend []config.Backend) []Backend {
	var backends []Backend
	for _, b := range backend {
		backends = append(backends, Backend{
			URL:         b.URL,
			IsAlive:     true,
			connections: 0,
		})
	}

	return backends
}
