package internal

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRoundRobinLoadBalancer(t *testing.T) {
	b := baseLoadBalancer{
		strategy: StrategyRoundRobin,
		backends: []Backend{},
	}
	lb := NewRoundRobinLoadBalancer(b)
	assert.NotNil(t, lb, "Expected non-nil RoundRobinLoadBalancer")
}

func TestRoundRobinLoadBalancer_nextBE(t *testing.T) {
	be1, _ := url.Parse("http://localhost:8081")
	be2, _ := url.Parse("http://localhost:8082")
	b := baseLoadBalancer{
		strategy: StrategyRoundRobin,
		backends: []Backend{
			{URL: *be1, IsAlive: true},
			{URL: *be2, IsAlive: true},
		},
	}
	lb := NewRoundRobinLoadBalancer(b)
	rr, ok := lb.(*RoundRobinLoadBalancer)
	assert.True(t, ok, "Expected type *RoundRobinLoadBalancer")

	// First call
	u1, err1 := rr.nextBE()
	assert.NoError(t, err1)
	assert.Equal(t, *be2, u1)

	// Second call
	u2, err2 := rr.nextBE()
	assert.NoError(t, err2)
	assert.Equal(t, *be1, u2)
}

func TestRoundRobinLoadBalancer_nextBE_noBackends(t *testing.T) {
	b := baseLoadBalancer{
		strategy: StrategyRoundRobin,
		backends: []Backend{},
	}
	lb := NewRoundRobinLoadBalancer(b)
	rr, ok := lb.(*RoundRobinLoadBalancer)
	assert.True(t, ok, "Expected type *RoundRobinLoadBalancer")

	_, err := rr.nextBE()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no backends available")
}
