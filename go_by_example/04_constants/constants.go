package main

import (
	"fmt"
	"math"
)

// Go supports constants of character, string, boolean and numeric values

// const declares a constant value (can't be changed)
// a const statement can appear anywhere a var statement can

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500_000_000

	// constant expressions perform arithmetic with arbitrary precision
	// a numeric constant has no type until it's given one, such as by an explicit conversion
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))

	// a number can be given a type by using it in a context that requires one
	// such as a variable assignment or function call

	// in this example, math.Sin expects a float64
	fmt.Println(math.Sin(n))
}
