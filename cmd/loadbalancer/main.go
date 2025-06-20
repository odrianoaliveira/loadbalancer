package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/odrianoaliveira/mini-loadbalancer/internal"
)

func main() {
	cfgFile := flag.String("config", "", "Path to the configuration file")
	flag.Parse()

	if *cfgFile == "" {
		fmt.Println("Configuration file path is required")
		flag.Usage()
		os.Exit(1)
	}

	lb := internal.NewLoadBalancer(*cfgFile)
	lb.Start()
}
