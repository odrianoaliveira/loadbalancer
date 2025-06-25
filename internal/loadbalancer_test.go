package internal

import (
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
