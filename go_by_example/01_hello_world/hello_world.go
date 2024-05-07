package main

import "fmt"

// a Go application needs a `package main` and a `func main` in order to execute
// this is the entrypoint of your application
func main() {
	// $ go run <filepath>
	// $ go build <filepath> && ./<binary>
	// fmt é o pacote da stdlib usado para lidar com formatação de I/O
	fmt.Println("hello world!")
}
