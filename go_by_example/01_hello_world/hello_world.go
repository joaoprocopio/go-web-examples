package main

import "fmt"

// a Go application needs a `package main` and a `func main` in order to execute
// it is the entrypoint of your application
func main() {
	// this command is used to compiled and directly run a Go file
	// $ go run hello_world.go

	// these commands are used to build and execute Go code
	// $ go build hello_world.go
	// $ ./hello_world

	// fmt is a standard library package that handles I/O formattings
	fmt.Println("hello world!")
}
