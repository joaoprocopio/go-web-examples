package main

import (
	"fmt"
)

// in Go, variables are explicitly declared and used by the compiler
// for example, the type declarations of a variable are used to check the type-correctness of function calls

func main() {
	// var keyword declares 1 or more variables
	var a = "initial"
	fmt.Println(a)

	// you can declare multiple variables at once
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go will infer the type of initialized valies (extremely important)
	// because Go will infer the type of a variable by its initialized value it makes this variable below be valid
	var d = true
	fmt.Println(d)

	// variables declared without a corresponding initialization are zero-valued
	// for example, the zero value for an int is 0
	var e int
	fmt.Println(e)

	// and as a example to zero-valued variables
	// the zero value for an string is "" (empty string)
	var f string
	fmt.Println(f, len(f))

	// the := syntax is a shorthand for declaring and initializing a variable
	// for exemple the code below is equivalent to var f string = "apple"
	// this syntax is only available inside functions (extremely important)
	g := "apple"
	fmt.Println(g)
}
