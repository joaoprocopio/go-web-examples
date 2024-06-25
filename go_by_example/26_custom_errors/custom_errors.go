package main

import (
	"errors"
	"fmt"
)

// it's possible to use custom types as errors by implementing the Error() method on them
// here's a variant on the example above that uses a custom type to explicitly represent an argument error

// a custom error type usually has the suffix "Error"
type argError struct {
	arg     int
	message string
}

// adding this Error method makes argError implement error interface
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {
		// return our custom error
		return -1, &argError{arg, "can't work with it"}
	}

	return arg + 3, nil
}

func main() {
	// errors.As is a more advanced version of errors.Is
	// it checks that a given error (or any error in its chain)
	// matches a specific error type and converts
	// to a value of that type, returning true
	_, err := f(42)
	var ae *argError
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}
