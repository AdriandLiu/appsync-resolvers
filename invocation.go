package resolvers

import (
	"encoding/json"
	"fmt"
)

type context struct {
	Arguments json.RawMessage `json:"arguments"`
	Source    json.RawMessage `json:"source"`
}

type invocation struct {
	Resolve string  `json:"resolve"`
	Context context `json:"context"`
}

func (in invocation) isRoot() bool {
	return in.Context.Source == nil || string(in.Context.Source) == "null"
}

func (in invocation) payload() json.RawMessage {
	if in.isRoot() {
		fmt.Println("!!!!")
		return in.Context.Arguments
	}

	return in.Context.Source
}
