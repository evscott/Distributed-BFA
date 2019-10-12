package Models

import (
	"github.com/evscott/Distributed-BFA/constants"
)

// Pair representing a pair of node ID's, I and J.
type Edge struct {
	Node     string `json:"i"`
	Distance int    `json:"j"`
}

// TODO
type Message struct {
	Source string           `json:"source"`
	Intent constants.Intent `json:"intent"`
	Length map[string]int   `json:"length"`
}

// Just for pretty printing Request/Response info.
func (req Message) String() string {
	return "Message:{ Origin:" + req.Source + ", Intent: " + string(req.Intent) + " }\n"
}
