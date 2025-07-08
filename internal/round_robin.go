package internal

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)

type RoundRobinLoadBalancer struct {
	baseLoadBalancer
	tcounter uint64
}

func NewRoundRobinLoadBalancer(b baseLoadBalancer) LoadBalancer {
	return &RoundRobinLoadBalancer{
		baseLoadBalancer: b,
		tcounter:         0,
	}
}

func (r *RoundRobinLoadBalancer) Start() error {
	port := fmt.Sprintf(":%d", r.port)
	slog.Info("Starting load balancer...", "port", port)
	listenAddr := fmt.Sprintf(":%d", r.port)

	slog.Info("Load balancer started", "address", listenAddr)
	if err := http.ListenAndServe(listenAddr, r.withProxy()); err != nil {
		return fmt.Errorf("failed to ListenAndServe the load balancer: %w", err)
	}

	return nil
}

func (r *RoundRobinLoadBalancer) withProxy() *httputil.ReverseProxy {
	director := func(req *http.Request) {
		target, err := r.nextBE()
		if err != nil {
			slog.Error("reverse proxy error", "error", err)
			req.URL = nil
			req.Header.Set("X-Reverse-Proxy-Error", "backend unavailable")
			return
		}
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.Host = target.Host
	}

	return &httputil.ReverseProxy{Director: director}
}

func (r *RoundRobinLoadBalancer) nextBE() (url.URL, error) {
	if len(r.backends) == 0 {
		return url.URL{}, fmt.Errorf("no backends available")
	}
	atomic.AddUint64(&r.tcounter, 1)

	n := r.tcounter % uint64(len(r.backends))
	return r.backends[n].URL, nil
}
