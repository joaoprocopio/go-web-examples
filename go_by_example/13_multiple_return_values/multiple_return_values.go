package main

import "fmt"

// Go has built-in support for multiple return values
// this feature is used often in idiomatic Go
// for example to return both result and error values from a function

// the (int, int) in this function signature shows that the function returns 2 integers
func vals() (int, int) {
	return 3, 7
}

func main() {
	// here we use the 2 different return values from the call with multiple assignment
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// if you only want a subset of the returned values
	// use the blank identifier _
	_, c := vals()
	fmt.Println(c)

	// accepting a variable number of arguments is another feature of Go functions
}
