package internal

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type BackendConfig struct {
	URL string `yaml:"url"`
}

type LoadBalancerConfig struct {
	Strategy LBStrategy      `yaml:"strategy"`
	Backends []BackendConfig `yaml:"backends"`
}

type Config struct {
	LoadBalancerConfig LoadBalancerConfig `yaml:"loadbalancer"`
}

func ReadConfig(filePath string) (Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return Config{}, fmt.Errorf("configuration file not found: %s", filePath)
		}
		return Config{}, err
	}

	var config Config
	if err = yaml.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}

	if len(config.LoadBalancerConfig.Backends) == 0 {
		return Config{}, fmt.Errorf("no backends configured in the load balancer configuration")
	}

	return config, nil
}
