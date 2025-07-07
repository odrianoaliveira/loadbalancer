package internal

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMapToBackends(t *testing.T) {
	cfg := &Config{
		LoadBalancerConfig: LoadBalancerConfig{
			Strategy: "round-robin",
			Backends: []BackendConfig{
				{URL: "http://localhost:8081"},
				{URL: "http://localhost:8082"},
			},
		},
	}

	mappedBackends, err := mapToBackends(cfg.LoadBalancerConfig.Backends)
	require.NoError(t, err, "expected no error when mapping backends")
	require.Len(t, mappedBackends, 2, "expected two backends to be mapped")

	var expectedURLs = []string{
		"http://localhost:8081",
		"http://localhost:8082",
	}

	var actualURLs []string
	for _, backend := range mappedBackends {
		actualURLs = append(actualURLs, backend.URL.String())
	}
	assert.ElementsMatch(t, expectedURLs, actualURLs, "backend URLs do not match")
}

func TestMapToBackends_InvalidURL(t *testing.T) {
	cfg := &Config{
		LoadBalancerConfig: LoadBalancerConfig{
			Strategy: "round-robin",
			Backends: []BackendConfig{
				{URL: "://invalid-url"},
			},
		},
	}
	_, err := mapToBackends(cfg.LoadBalancerConfig.Backends)
	assert.Error(t, err, "expected error for invalid URL")
}

func TestNewLoadBalancer_Success(t *testing.T) {
	// Create a temporary YAML config file
	configContent := `loadbalancer:
  strategy: round-robin
  backends:
    - url: http://localhost:8081
    - url: http://localhost:8082
`
	tmpFile, err := os.CreateTemp("", "lbconfig-*.yaml")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name())
	_, err = tmpFile.Write([]byte(configContent))
	require.NoError(t, err)
	tmpFile.Close()

	lb, err := NewLoadBalancer(tmpFile.Name())
	require.NoError(t, err)
	require.NotNil(t, lb, "expected load balancer to be created")
	rr, ok := lb.(*RoundRobinLoadBalancer)
	require.True(t, ok, "expected lb to be of type *RoundRobinLoadBalancer")
	require.Len(t, rr.backends, 2, "expected two backends in RoundRobinLoadBalancer")
}

func TestNewLoadBalancer_InvalidConfig(t *testing.T) {
	// Invalid YAML
	tmpFile, err := os.CreateTemp("", "lbconfig-*.yaml")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name())
	_, err = tmpFile.Write([]byte("invalid: [yaml"))
	require.NoError(t, err)
	tmpFile.Close()

	_, err = NewLoadBalancer(tmpFile.Name())
	assert.Error(t, err, "expected error for invalid config")
}

func TestNewLoadBalancer_NoBackends(t *testing.T) {
	// Config with no backends
	configContent := `loadbalancer:
  strategy: round-robin
  backends: []
`
	tmpFile, err := os.CreateTemp("", "lbconfig-*.yaml")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name())
	_, err = tmpFile.Write([]byte(configContent))
	require.NoError(t, err)
	tmpFile.Close()

	_, err = NewLoadBalancer(tmpFile.Name())
	assert.Error(t, err, "expected error for no backends")
}
