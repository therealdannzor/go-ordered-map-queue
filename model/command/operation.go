package command

import "fmt"

// OperationType denotes a valid operation for the server to execute
type OperationType string

const (
	Add    OperationType = "ADD"
	Delete OperationType = "DELETE"
	Lookup OperationType = "LOOKUP"
	GetAll OperationType = "GET_ALL"
)

func (o OperationType) String() string {
	return string(o)
}

func ParseOperation(s string) (OperationType, error) {
	allowed := map[string]bool{
		"ADD":     true,
		"DELETE":  true,
		"LOOKUP":  true,
		"GET_ALL": true,
	}

	if allowed[s] {
		op := OperationType(s)
		return op, nil
	}

	return "", fmt.Errorf("invalid operation: %s", s)
}
