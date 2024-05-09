package main

import "fmt"

// go supports anonymous functions, which can form closures
// anonymous functions are useful when you want to define a function inline without having to name it

// the function intSeq returns another function
// which we define anonymously in the body of intSeq
// the returned function closes over a variable i to form a closure
func intSeq() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

func main() {
	// we call intSeq, assigning the result (a function) to nextInt
	// this function value captures its own i value
	// which will be updated time we call nextInt
	nextInt := intSeq()

	// see the effect of the closure by calling nextInt a few times
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// to confirm that the state is unique to that particular function
	// create and test a new one
	newInts := intSeq()
	fmt.Println(newInts())
}
