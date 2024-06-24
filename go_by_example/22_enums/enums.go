package main

import "fmt"

// enumerated types (enums) area a special case of sum types
// https://en.wikipedia.org/wiki/Algebraic_data_type
// an enum ia type that has a fixed number of possible values
// each with a distinct name
// Go doesn't have an enum type as a distinct language feature
// but enums are simple to implement using existing idioms

// our enum type ServerState has an underlying int type
type ServerState int

// the possible values for ServerState are defined as constants
// the special keyword iota generates successive constant values automatically
// in this case 0, 1, 2 and so on
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

// by implementing fmt.Stringer interface
// values of ServerState can be printed out or converted to strings
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retying",
	// this can get cumbersome if there are many possible values
	// in such cases the stringer tool can be used in conjunction with go:generate to automate the process
	// https://pkg.go.dev/golang.org/x/tools/cmd/stringer
	// see this post for a longer explanation
	// https://eli.thegreenplace.net/2021/a-comprehensive-guide-to-go-generate
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	// if we have a value of type int
	// we cannot pass it to transition
	// the compiler will complain about type mismatch
	// this provides some degree of compile-time type safety for enums
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)
}

// transition emulates a state transition for a server
// it takes the existing state and returns a new state
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		// suppose we check some predicates here to determine the next state...
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
