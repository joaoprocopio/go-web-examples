package main

import "fmt"

func main() {
	// go tem diversos tipos de valores, incluindo strings, integers, floats, booleans e etc
	// os básicos são

	// strings, podem ser concatenadas usando o operador de soma (+)
	fmt.Println("go" + "lang")

	// integers
	fmt.Println("1+1 =", 1+1)

	// floats
	fmt.Println("7.0/3.0", 7.0/3.0)

	// booleans, e seus operadores
	fmt.Println(true && false) // and
	fmt.Println(true || false) // or
	fmt.Println(!true)         // not
}
