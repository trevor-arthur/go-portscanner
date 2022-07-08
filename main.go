package main

import (
	"fmt"

	"github.com/trevor-arthur/go-portscanner/port"
)

func main() {
	fmt.Println("[*] Go Port Scanner")

	open := port.ScanPort("tcp", "192.168.0.169", 22)
	fmt.Printf("Port Open: %t\n", open)

	results := port.InitialScan("192.168.0.169")
	fmt.Println(results)
}
