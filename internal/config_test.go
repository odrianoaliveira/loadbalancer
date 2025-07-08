package internal

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadConfig_Success(t *testing.T) {
	tmpDir := t.TempDir()
	configContent := `
loadbalancer:
  strategy: round-robin
  backends:
    - url: http://localhost:8081
    - url: http://localhost:8082
`
	configPath := filepath.Join(tmpDir, "config.yaml")
	require.NoError(t, os.WriteFile(configPath, []byte(configContent), 0644), "failed to write temp config file")

	cfg, err := ReadConfig(configPath)
	require.NoError(t, err, "expected no error")

	assert.Equal(t, StrategyRoundRobin, cfg.LoadBalancerConfig.Strategy, "unexpected strategy")

	assert.Len(t, cfg.LoadBalancerConfig.Backends, 2, "unexpected number of backends")
	expectedURLs := []string{
		"http://localhost:8081",
		"http://localhost:8082",
	}
	var actualURLs []string
	for _, backend := range cfg.LoadBalancerConfig.Backends {
		actualURLs = append(actualURLs, backend.URL)
	}

	assert.ElementsMatch(t, expectedURLs, actualURLs, "unexpected backend URLs")

}

func TestReadConfig_FileNotFound(t *testing.T) {
	_, err := ReadConfig("nonexistent.yaml")
	require.Error(t, err, "expected error for missing file")
	assert.EqualError(t, err, "configuration file not found: nonexistent.yaml")
}

func TestReadConfig_InvalidYAML(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "bad.yaml")
	badContent := "not: [valid, yaml"
	require.NoError(t, os.WriteFile(configPath, []byte(badContent), 0644), "failed to write temp config file")

	_, err := ReadConfig(configPath)
	require.Error(t, err, "expected error for invalid YAML")
}
