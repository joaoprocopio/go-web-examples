package main

import (
	"errors"
	"fmt"
)

// in Go it's idiomatic to communicate errors via an explicit, separate return value
// this contrasts with the exceptions used in languages like Java and Ruby and the overloaded single result / error value used in C
// Go's approach makes it easy to see which functions returns and to handle them using the same language constructs employed for other, no-error tasks

// errors package https://pkg.go.dev/errors
// blogpost https://go.dev/blog/go1.13-errors

// by convention, errors are the last return value and have type error, a built-in interface
func f(arg int) (int, error) {
	if arg == 42 {
		// errors.New constructs a basic error value with the given error messaage
		return -1, errors.New("can't work with 42")
	}

	// a nil value in the error position indicates that there was no error
	return arg + 3, nil
}

// a sentinel error is a predeclared variable that is used to signify a specific error condition
var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		// we can wrap errors with higher-level errors to add context
		// the simples way to do this is with the %w verb in fmt.Errorf
		// wrapped errors create a logical chain (A wraps B, which wraps C, etc.)
		// that can be queried functions like errors.Is and errors.As
		return fmt.Errorf("making tea: %w", ErrPower)
	}

	return nil
}

func main() {
	for _, i := range []int{7, 42} {
		// it's common to use an inline error check in the if line
		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			// errors.Is checks that a given error (or any error in its chain) matches a specific error value
			// this is specially useful with wrapped or nested errors
			// allowing you to identify specific error types or sentinel errors in a chain of errors
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("we should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("now it is dark")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		}

		fmt.Println("tea is ready!")
	}
}
