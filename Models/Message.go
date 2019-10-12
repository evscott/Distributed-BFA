package Models

import (
	"github.com/evscott/Distributed-BFA/constants"
)

// Pair representing a pair of node ID's, I and J.
type Pair struct {
	I string `json:"i"`
	J string `json:"j"`
}

// TODO
type Message struct {
	Source     string           `json:"source"`
	Intent     constants.Intent `json:"intent"`
	Data       string           `json:"data"`
}

// Just for pretty printing Request/Response info.
func (req Message) String() string {
	return "Message:{ Origin:" + req.Source + ", Intent: " + string(req.Intent) + " }\n"
}