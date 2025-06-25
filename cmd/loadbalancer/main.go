package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/odrianoaliveira/loadbalancer/internal"
	"go.uber.org/zap"
)

func main() {
	cfgFile := flag.String("config", "", "Path to the configuration file")
	flag.Parse()

	logger := internal.NewLogger()

	if *cfgFile == "" {
		logger.Warn("Configuration file path is required")
		flag.Usage()
		return
	}

	if err := run(logger, *cfgFile); err != nil {
		logger.Error("Error running load balancer", zap.Error(err))
		os.Exit(1)
	}
}

func run(logger *zap.Logger, fileName string) error {
	lb, err := internal.NewLoadBalancer(fileName, logger)
	if err != nil {
		return fmt.Errorf("failed to create load balancer: %w", err)
	}

	if err = lb.Start(); err != nil {
		return fmt.Errorf("failed to start load balancer: %w", err)
	}

	return nil
}
