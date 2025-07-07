package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/odrianoaliveira/loadbalancer/internal"
)

func main() {
	cfgFile := flag.String("config", "", "Path to the configuration file")
	flag.Parse()

	if *cfgFile == "" {
		slog.Warn("Configuration file path is required")
		flag.Usage()
		return
	}

	if err := run(*cfgFile); err != nil {
		slog.Error("Error running load balancer", "error", err)
		os.Exit(1)
	}
}

func run(fileName string) error {
	lb, err := internal.NewLoadBalancer(fileName)
	if err != nil {
		return fmt.Errorf("failed to create load balancer: %w", err)
	}

	if err = lb.Start(); err != nil {
		return fmt.Errorf("failed to start load balancer: %w", err)
	}

	return nil
}
