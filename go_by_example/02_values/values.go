package main

import "fmt"

func main() {
	// Go has various value types including strings, integers, floats, booleans, etc
	// here are some basic types

	// strings, can be concatenated using with the add operator (+)
	fmt.Println("go" + "lang")

	// integers
	fmt.Println("1+1 =", 1+1)

	// floats
	fmt.Println("7.0/3.0", 7.0/3.0)

	// booleans, and their logical operators
	fmt.Println(true && false) // and
	fmt.Println(true || false) // or
	fmt.Println(!true)         // not
}
