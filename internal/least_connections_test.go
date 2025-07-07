package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLeastConnectionsLoadBalancer(t *testing.T) {
	b := baseLoadBalancer{
		strategy: StrategyLeastConn,
		backends: []Backend{},
	}
	lb := NewLeastConnectionsLoadBalancer(b)
	assert.NotNil(t, lb, "Expected non-nil LeastConnectionsLoadBalancer")
}

func TestLeastConnectionsLoadBalancer_Start(t *testing.T) {
	b := baseLoadBalancer{
		strategy: StrategyLeastConn,
		backends: []Backend{},
	}
	lb := NewLeastConnectionsLoadBalancer(b)
	err := lb.Start()
	assert.Error(t, err, "Expected error from Start")
	assert.EqualError(t, err, "LeastConnectionsLoadBalancer does not implement Start method")
}
