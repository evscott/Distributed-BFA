package Node

import (
	"encoding/json"
	"fmt"
	"github.com/evscott/Distributed-BFA/Models"
	"github.com/evscott/Distributed-BFA/constants"
	"net"
	"time"
)

type Info struct {
	IP         string            `json:"IP"`
	Port       string            `json:"port"`
	Neighbours []string          `json:"neighbors"`
	Nodes      []string          `json:"nodes"`
	LG         map[string]int    `json:"lg"`
	Length     map[string]int    `json:"length"`
	RoutingTo  map[string]string `json:"routingTo"`
}

// Just for pretty printing Node info
func (i Info) String() string {
	return "NodeInfo:{IP:" + i.IP + ", Port:" + i.Port + " }"
}

// Create is used a constructor that instantiates a new node using it's initial knowledge.
//
// A node must be created with initial knowledge of it's network IP, ID, and the IDs of it's neighbors.
func Create(ip, port string, neighbours []Models.Edge, nodes []string) *Info {
	newNode := Info{
		IP:        ip,
		Port:      port,
		Nodes:     nodes,
		LG:        make(map[string]int),
		Length:    make(map[string]int),
		RoutingTo: make(map[string]string),
	}

	for _, neighbour := range neighbours {
		newNode.Neighbours = append(newNode.Neighbours, neighbour.Node)
		newNode.LG[neighbour.Node] = neighbour.Distance
	}

	for _, node := range newNode.Nodes {
		newNode.Length[node] = 1000
	}
	newNode.Length[newNode.Port] = 0

	return &newNode
}

// TODO
func (i *Info) Start() {
	for _, neighbor := range i.Neighbours {
		msgOut := Models.Message{
			Source: i.Port,
			Intent: constants.IntentSendUpdate,
			Length: i.Length,
		}
		if err := i.SendMsg(msgOut, neighbor); err != nil {
			fmt.Printf("Error sending message")
			return // failure, terminate
		}
	}
}

// TODO
func (i *Info) ReceiveUpdate(msgIn Models.Message) {
	updated := false
	j := msgIn.Source

	for _, k := range i.Nodes {
		if k != i.Port {
			if i.Length[k] > (i.LG[j] + msgIn.Length[k]) {
				i.Length[k] = i.LG[j] + msgIn.Length[k]
				i.RoutingTo[k] = j
				updated = true
			}
		}
	}

	if updated {
		for _, neighbour := range i.Neighbours {
			msgOut := Models.Message{
				Source: i.Port,
				Intent: constants.IntentSendUpdate,
				Length: i.Length,
			}
			if err := i.SendMsg(msgOut, neighbour); err != nil {
				fmt.Println("Error sending message")
				return // failure, terminate
			}
		}
	}
}

// TODO
func (i *Info) SendMsg(msg Models.Message, dest string) error {
	connOut, err := net.DialTimeout("tcp", i.IP+":"+dest, time.Duration(10)*time.Second)
	if err != nil {
		if _, ok := err.(net.Error); ok {
			fmt.Printf("Couldn't send go to %s:%s \n", i.IP, dest)
			return err
		}
	}

	if err := json.NewEncoder(connOut).Encode(&msg); err != nil {
		fmt.Printf("Couldn't enncode message %v \n", msg)
		return err
	}
	return nil
}

// ListenOnPort is the communication satellite for a node that listens for incoming messages.
// Incoming messages are marshalled into a `Message` struct, and are directed to a handler
// depending on the messages `Intent`.
// Incoming messages that cannot be marshalled into a `Message` may cause erroneous behaviour.
func (i *Info) ListenOnPort() {
	ln, err := net.Listen("tcp", fmt.Sprint(":"+i.Port))
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("Staring node on %s:%s...\n", i.IP, i.Port)

	for {
		connIn, err := ln.Accept()
		if err != nil {
			if _, ok := err.(net.Error); ok {
				fmt.Printf("Error received while listening %s:%s \n", i.IP, i.Port)
			}
		}

		var msg Models.Message
		if err := json.NewDecoder(connIn).Decode(&msg); err != nil {
			fmt.Printf("Error decoding %v\n", err)
		}

		switch msg.Intent {
		case constants.IntentSendUpdate:
			i.ReceiveUpdate(msg)
		}
	}
}
