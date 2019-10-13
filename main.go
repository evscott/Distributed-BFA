package main

import (
	"fmt"
	"github.com/evscott/Distributed-BFA/Models"
	"github.com/evscott/Distributed-BFA/Node"
	"net"
	"strings"
	"time"
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

func printNeighbourPaths(node *Node.Info) {
	for dest, cost := range node.LG {
		fmt.Printf("%s --> %s, cost: %d\n", node.Port, dest, cost)
	}
}

func printShortestPaths(node *Node.Info) {
	for dest, cost := range node.Length {
		fmt.Printf("%s --> %s, cost: %d\n", node.Port, dest, cost)
	}
}

// runExample creates a distributed system and finds the shortest path between all of its nodes.
//
// This example follows the graph in `assets/example.png` with the following name assignments:
// `1` = `8001`
// `2` = `8002`
// `2` = `8003`
// ... etc
func runExample(ip string) {
	allNodes := []string{"8001", "8002", "8003", "8004", "8005"}

	n1 := Node.Create(
		ip,
		"8001",
		[]Models.Edge{
			{"8002", 2},
			{"8004", 5},
		},
		allNodes)

	n2 := Node.Create(
		ip,
		"8002",
		[]Models.Edge{
			{"8001", 2},
			{"8003", 14},
			{"8004", 5},
			{"8005", 4},
		},
		allNodes,
	)

	n3 := Node.Create(
		ip,
		"8003",
		[]Models.Edge{
			{"8002", 14},
			{"8005", 34},
		},
		allNodes,
	)

	n4 := Node.Create(
		ip,
		"8004",
		[]Models.Edge{
			{"8001", 5},
			{"8002", 5},
			{"8005", 58},
		},
		allNodes,
	)

	n5 := Node.Create(
		ip,
		"8005",
		[]Models.Edge{
			{"8002", 4},
			{"8003", 34},
			{"8004", 58},
		},
		allNodes,
	)

	go n1.ListenOnPort()
	go n2.ListenOnPort()
	go n3.ListenOnPort()
	go n4.ListenOnPort()
	go n5.ListenOnPort()

	time.Sleep(time.Second / 10)

	fmt.Println()

	n1.Start()

	time.Sleep(time.Second)

	fmt.Println("Neighbour paths:")

	fmt.Println()
	printNeighbourPaths(n1)

	fmt.Println()
	printNeighbourPaths(n2)

	fmt.Println()
	printNeighbourPaths(n3)

	fmt.Println()
	printNeighbourPaths(n4)

	fmt.Println()
	printNeighbourPaths(n5)

	fmt.Println("\nShortest paths:")

	fmt.Println()
	printShortestPaths(n1)

	fmt.Println()
	printShortestPaths(n2)

	fmt.Println()
	printShortestPaths(n3)

	fmt.Println()
	printShortestPaths(n4)

	fmt.Println()
	printShortestPaths(n5)
}
