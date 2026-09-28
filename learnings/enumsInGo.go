package learnings

import "fmt"

// Go doesn’t have an enum type as a distinct language feature

type ServerState uint8

// iota assigns successive numbers, starting at zero.
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

// serverState := ServerState(99) is valid even though the we have defined enums from 0 to 3
// So these constants give you enum-like names, but they don’t create a closed set of allowed values or require exhaustive switch cases.

// Enum to String
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

// It becomes method for ServerState
// It also satisfies fmt.Stringer
func (ss ServerState) String() string {
	return stateName[ss]
}

func LearnEnums() {
	serverState := StateConnected
	// Method
	fmt.Println(serverState.String())
	// fmt.Stringer
	// If map is not created, then it will print the uint8 value assigned to serverState
	fmt.Println(serverState)

}
