package Models

import (
	"github.com/evscott/Distributed-BFA/constants"
)

// Pair representing a pair of node ID's, I and J.
type Edge struct {
	Node     string `json:"i"`
	Distance int    `json:"j"`
}

// The format for a Request/Response in finding a shortest path across a distributed system.
//
// Source represents the Message sender.
// Intent represents the messages intent; i.e., whether it is to be handled by `Update` or some other handler.
// Length represents the `Source`'s lengths spanning across the distributed system; i.e.: how far away node Z is from Y.
type Message struct {
	Source string           `json:"source"`
	Intent constants.Intent `json:"intent"`
	Length map[string]int   `json:"length"`
}

// Just for pretty printing Request/Response info.
func (req Message) String() string {
	return "Message:{ Origin:" + req.Source + ", Intent: " + string(req.Intent) + " }\n"
}
