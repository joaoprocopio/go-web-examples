package main

import (
	"fmt"
	"slices"
)

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
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// slices can aso be copy'd
	// here we create an empty slice c of the same length as s and copy into c from s
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// slices support a "slice" operator with the syntax slice[low:high]
	// for example, this gets a slice of elements s[2], s[3] and s[4]
	l := s[2:5]
	fmt.Println("sl1:", l)

	// the slices up to (but excluding s[5])
	l = s[:5]
	fmt.Println("sl2:", l)

	// and slices up from (and including) s[2]
	l = s[2:]
	fmt.Println("sl3:", l)

	// we can declare and initialize a variable for slice in a single line as well
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// the slices package contains a number of useful utility functions for slices
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// slices can be composed into multi-dimensional data structures
	// the length of the inner slices can vary, unlike multi-dimensional arrays
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	// note that while slices are different types than arrays
	// they are rendered similarly by fmt.Println
	// checkout this blog post about slices https://go.dev/blog/slices-intro
}
