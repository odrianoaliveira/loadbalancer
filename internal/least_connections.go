package internal

import "fmt"

type LeastConnectionsLoadBalancer struct {
	baseLoadBalancer
}

func NewLeastConnectionsLoadBalancer(b baseLoadBalancer) LoadBalancer {
	return &LeastConnectionsLoadBalancer{
		baseLoadBalancer: b,
	}
}

func (l *LeastConnectionsLoadBalancer) Start() error {
	return fmt.Errorf("LeastConnectionsLoadBalancer does not implement Start method")
}
