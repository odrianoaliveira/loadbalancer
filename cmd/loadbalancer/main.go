package main

import "fmt"

type LoadBalancer struct {
	servers   []string
	nextIndex int
}

func main() {
	// Initialize the load balancer
	lb := NewLoadBalancer()

	// Start the load balancer
	go lb.Start()

	// Simulate adding servers
	lb.AddServer("http://server1.com")
	lb.AddServer("http://server2.com")

	// Simulate handling requests
	for i := 0; i < 10; i++ {
		go func(reqID int) {
			server := lb.GetNextServer()
			fmt.Printf("Request %d is being handled by %s\n", reqID, server)
		}(i)
	}

	// Keep the main goroutine alive
	select {}
}
