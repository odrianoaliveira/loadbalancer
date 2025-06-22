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

	if *cfgFile == "" {
		fmt.Println("Configuration file path is required")
		flag.Usage()
		os.Exit(1)
	}

	logger := internal.NewLogger()

	lb, err := internal.NewLoadBalancer(*cfgFile, logger)
	if err != nil {
		logger.Error("Failed to create load balancer", zap.Error(err))
		os.Exit(1)
	}
	lb.Start()
}
