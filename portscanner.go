package main

import (
	"fmt"
	"net"
	"sort"
)

// Configurations
var target = "192.168.0.144"
var maxPort = 1024
var protocol = "tcp"

// Worker function
func worker(ports, results chan int) {
	for port := range ports {
		address := fmt.Sprintf("%v:%d", target, port)
		conn, err := net.Dial(protocol, address)
		if err != nil {
			// port is closed or filtered
			results <- 0
			continue
		}
		// port is open
		conn.Close()
		results <- port
	}
}

// Portscanner
func main() {
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= maxPort; i++ {
			ports <- i
		}
	}()

	for i := 0; i < maxPort; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}
	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("Port Open: %d\n", port)
	}
}
