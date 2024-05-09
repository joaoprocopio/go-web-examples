package main

import "fmt"

// in Go, an array is a numbered sequence of elements of a specific length
// in typical Go code, slices are much more common
// arrays are useful in some specific scenarios

func main() {
	// here we create an array a that hold exactly 5 integers
	// the type of elements are both part of the array's type
	// by default an array is zero-valued, which for int's means zeros
	var a [5]int
	fmt.Println("emp:", a)

	// we can set a value at an index using the array[index] = value syntax
	// and get a value with array[index]
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// the builtin len returns the length of an array
	fmt.Println("len:", len(a))

	// use this syntax to declare and initialize an array in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// you can also have the compiler count the number of elements for you with ...
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// if you specify the index with :, the elements in between will be zeroed
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	// array types are one-dimensional, but you can compose types to build multi-dimensional data structures
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d:", twoD)
}
