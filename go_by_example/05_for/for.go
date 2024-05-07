package main

import "fmt"

// for is Go's only looping construct
// here are some basic types of for loops

func main() {
	// the most basic type, with a single condition
	i := 0
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	// a classic initial/condition/after for loop
	// if you're familiar with C family languages and JavaScript, you surely know this
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// another way of accomplishing the basic "do this N times" iteration is range over an integer
	for i := range 3 {
		fmt.Println("range", i)
	}

	// for without a condition will loop repeatedly until you break out of the loop or return from enclosing function
	for {
		fmt.Println("loop")
		break
	}

	// you can also continue to the next iteration of the loop
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	// break and continue constructs are used for control flow
	// return inside a for loop enclose a function
}
