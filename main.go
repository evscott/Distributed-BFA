package main

import (
	"fmt"
	"net"
	"strings"
)

// main is the entry point for this distributed system.
func main() {
	ipAddr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Print(err)
		return
	}
	ip := strings.Split(ipAddr[0].String(), "/")[0]

	runExample(ip)
}

func runExample(ip string) {

}