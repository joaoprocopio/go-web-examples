package main

import "fmt"

// um programa em go precisa de um package main uma func main, esse é o entrypoint da aplicação
func main() {
	// $ go run <filepath>
	// $ go build <filepath> && ./<binary>
	// fmt é o pacote da stdlib usado para lidar com formatação de I/O
	fmt.Println("hello world!")
}
