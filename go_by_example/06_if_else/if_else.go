package main

import "fmt"

// branching with if and else in Go is straight-forward

func main() {
	// here's a basic example
	if 7%2 == 0 {
		fmt.Println("7 é par")
	} else {
		fmt.Println("7 é ímpar")
	}

	// you can have an if statement without and else
	if 8%4 == 0 {
		fmt.Println("8 é divisível por 4")
	}

	// logical operators like && and || are often useful in conditions
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("ou 8 ou 7 é par")
	}

	// a statement can precede conditionals
	// any variables declared in this statement are available in the current and all subsequent branches
	if num := 9; num < 0 {
		fmt.Println(num, "é negativo")
	} else if num < 10 {
		fmt.Println(num, "tem 1 dígito")
	} else {
		fmt.Println(num, "tem múltiplos dígitos")
	}

	// note that you don't need parentheses around conditions in Go, but that the braces are required

	// there is no ternary if in Go, so you'll need to use a full if statement even for a basic condition
}
