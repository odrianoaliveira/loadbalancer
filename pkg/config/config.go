package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Backend struct {
	URL string `yaml:"url"`
}

type LoadBalancerConfig struct {
	Strategy string    `yaml:"strategy"`
	Backends []Backend `yaml:"backends"`
}

type Config struct {
	LoadBalancerConfig LoadBalancerConfig `yaml:"loadbalancer"`
}

func ReadConfig(filePath string) (*Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("configuration file not found: %s", filePath)
		}
		return nil, err
	}

	var config Config
	if err = yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
