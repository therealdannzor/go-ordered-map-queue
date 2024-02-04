package command

import "fmt"

// Command represents a valid server instruction
type Command struct {
	// Valid operation
	Operation OperationType `json:"operation"`

	// Optional key and value parameters
	Params []string `json:"params"`
}

func ParseCommand(o OperationType, params []string) (*Command, error) {
	if o == Add {
		if len(params) != 2 {
			return nil, fmt.Errorf(
				"unable to parse %s operation, requires two parameters (key and value)",
				Add.String(),
			)
		}
	} else if o == Delete {
		if len(params) != 1 {
			return nil, fmt.Errorf(
				"unable to parse %s operation, requires a parameter (key)",
				Delete.String(),
			)
		}
	} else if o == Lookup {
		if len(params) != 1 {
			return nil, fmt.Errorf(
				"unable to parse %s operation, requires a paramter (key)",
				Lookup.String(),
			)
		}
	}

	return &Command{Operation: o, Params: params}, nil
}
