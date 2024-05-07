package main

import "fmt"

// slices are an important data type in Go
// giving a more powerful interface to sequences than arrays

func main() {
	// unlike arrays, slices are typed only by the elements they contain (not the number of elements)
	// an uninitialized slice equals to nil and has length 0
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// to create an empty slice with non-zero length, use the builtin make
	// here we make a slice of strings of length 3 (initially zero-valued)
	// by default a new slice's capacity is equal to its length
	// if we know the slice is going to grow ahead of time
	// it's possible to pass a capacity explicitly as an additional parameter to make
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// we can set and get just like arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// len returns the length of the slice as expected
	fmt.Println("len:", len(s))

	// in addition to these basic operations
	// slices support several more that make them richer than arrays
	// one is the builtin append, which returns a slice containing one or more new values
	// note that we need to accept a return value from append as we may get a new slice value
}
