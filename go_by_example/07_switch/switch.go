package main

import (
	"fmt"
	"time"
)

// switch statements express conditionals across many branches

func main() {
	i := 2
	fmt.Println("escreva", i, " como ")
	// here's a basic switch
	switch i {
	case 1:
		fmt.Println("um")
	case 2:
		fmt.Println("dois")
	case 3:
		fmt.Println("três")
	}

	// you can use commas to separate multiple expressions in the same case statement
	// we use the optional default case in this example too
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("é fim de semana")
	default:
		fmt.Println("é dia de semana")
	}

	t := time.Now()

	// switch without an expression is an alternate way to express if/else logic
	// here we also show how the case expressions can be non-constants
	switch {
	case t.Hour() < 12:
		fmt.Println("é antes de meio dia")
	default:
		fmt.Println("é depois de meio dia")
	}

	// a type switch compares types instead of values
	// you can use this to discover the type of an interface values
	// in this example, the variable t will have the type corresponding to its clause
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("sou bool")
		case int:
			fmt.Println("sou int")
		default:
			fmt.Printf("não conheço o tipo %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
